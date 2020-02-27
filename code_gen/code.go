package code_gen

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"yum/ast"
)

var typeMap = map[string]byte{
	"num": 0x7c,
}

var operandMap = map[string]byte{
	"+":             0xa0,
	"-":             0xa1,
	"*":             0xa2,
	"/":             0xa3,
	"min":           0xa4,
	"max":           0xa5,
	"sqrt":          0x9f,
	"ceil":          0x9b,
	"floor":         0x9c,
	"trunc":         0x9d,
	"nearest":       0x9e,
	"abs":           0x99,
	"neg":           0x9a,
	"copysign":      0xa6,
	"eq":            0x61,
	"ne":            0x62,
	"lt":            0x63,
	"gt":            0x64,
	"le":            0x65,
	"ge":            0x66,
	"f64.load":      0x2b,
	"f64.store":     0x39,
	"memory.grow":   0x40,
	"nop":           0x01,
	"drop":          0x1a,
	"i32.const":     0x41,
	"i64.const":     0x42,
	"f32.const":     0x43,
	"f64.const":     0x44,
	"get_local":     0x20,
	"set_local":     0x21,
	"tee_local":     0x22,
	"get_global":    0x23,
	"set_global":    0x24,
	"select":        0x1b,
	"call":          0x10,
	"call_indirect": 0x11,
}

const FUNC_END = 0x0b
const LOCAL_GET = 0x20     // local.get
const DECLARE_CONST = 0x41 // delcare const

type FuncData struct {
	Index  int
	Name   string
	Params map[string]FuncParam
}

type FuncParam struct {
	Index int
	Name  string
	Type  string
}

func float64ToByte(f float64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, f)
	if err != nil {
		panic("binary.Write failed: " + err.Error())
	}
	return buf.Bytes()
}

func operate(buffer *bytes.Buffer, funcName string, funcSymbolTable map[string]FuncData, operation *string, l *ast.Literal, r *ast.Expression) error {
	if l != nil && l.Num != nil {
		buffer.WriteByte(operandMap["f64.const"])
		buffer.Write(float64ToByte(*l.Num))
	}
	if r.Left.Num != nil {
		buffer.WriteByte(operandMap["f64.const"])
		buffer.Write(float64ToByte(*r.Left.Num))
		buffer.WriteByte(operandMap[*operation])
	}
	if l != nil && l.Reference != nil {
		if operation != nil {
			ref, ok := funcSymbolTable[funcName].Params[*l.Reference]
			if !ok {
				return fmt.Errorf("func" + funcName + " parameter " + *l.Reference + " does not exist")
			}
			buffer.WriteByte(operandMap["get_local"])
			buffer.WriteByte(byte(ref.Index))
		} else {
			// funcCALL
			callFunc, ok := funcSymbolTable[*l.Reference]
			if !ok {
				return fmt.Errorf("func" + funcName + " trying to call non-existing func '" + *l.Reference + "'")
			}
			if len(callFunc.Params) != len(r.Left.Params) {
				return fmt.Errorf("func'" + funcName + "' trying to call function '" + callFunc.Name + "' with extra parameters")
			}
			for _, v := range r.Left.Params {
				if v.Num != nil {
					buffer.WriteByte(operandMap["f64.const"])
					buffer.Write(float64ToByte(*v.Num))
				}
				if v.Reference != nil {
					vref, ok := funcSymbolTable[funcName].Params[*v.Reference]
					if !ok {
						return fmt.Errorf("func'" + funcName + "' trying to call '" + callFunc.Name + "' with non-existing variable " + *v.Reference)
					}
					buffer.WriteByte(operandMap["get_local"])
					buffer.WriteByte(byte(vref.Index))
				}
			}
			buffer.WriteByte(operandMap["call"])
			buffer.WriteByte(byte(callFunc.Index))
			if r.Operator != nil {
				err := operate(buffer, funcName, funcSymbolTable, r.Operator, nil, r.Right)
				if err != nil {
					return err
				}
			}
		}
	}
	if r.Left != nil && r.Left.Reference != nil {
		ref, ok := funcSymbolTable[funcName].Params[*r.Left.Reference]
		if !ok {
			return fmt.Errorf("func'" + funcName + "' parameter '" + *r.Left.Reference + "' does not exist")
		}
		buffer.WriteByte(operandMap["get_local"])
		buffer.WriteByte(byte(ref.Index))
		buffer.WriteByte(operandMap[*operation])
	}
	if r.Operator != nil {
		err := operate(buffer, funcName, funcSymbolTable, r.Operator, nil, r.Right)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateCode(ast *ast.Ast) (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write([]byte{0x00, 0x61, 0x73, 0x6d}) // WASM_BINARY_MAGIC
	buffer.Write([]byte{0x01, 0x00, 0x00, 0x00}) // WASM_BINARY_VERSION

	types := []byte{} // num types, func, num params
	funcs := []byte{}
	exports := []byte{}
	funcBodys := []byte{}
	funcSymbolTable := map[string]FuncData{}
	for _, p := range ast.Modules {
		for i, fun := range p.Funs {
			funcSymbolTable[fun.Name] = FuncData{
				Index:  i,
				Name:   fun.Name,
				Params: map[string]FuncParam{},
			}
			types = append(types, 0x60, byte(len(fun.Parameters))) // func, num params
			for pi, param := range fun.Parameters {                // i32, i32
				t, ok := typeMap[param.Type.Name]
				if !ok {
					return nil, fmt.Errorf("func'" + fun.Name + "' parameter '" + param.Name + "' type '" + param.Type.Name + "' does not exist")
				}
				types = append(types, t)
				funcSymbolTable[fun.Name].Params[param.Name] = FuncParam{
					Index: pi,
					Name:  param.Name,
					Type:  param.Type.Name,
				}
			}
			// TODO: we need to remove the last item since '\n' gets captured somehow needs to be fixed later
			fun.ReturnTypes = fun.ReturnTypes[:len(fun.ReturnTypes)-1]
			if len(fun.ReturnTypes) > 0 {
				t, ok := typeMap[fun.ReturnTypes[0]]
				if !ok {
					return nil, fmt.Errorf("func'" + fun.Name + "' return type '" + fun.ReturnTypes[0] + "' does not exist")
				}
				types = append(types, 0x01) // num results
				types = append(types, t)    // i32
			} else {
				types = append(types, 0x00) // num results
			}
			funcs = append(funcs, byte(i))                 // function 0 signature index
			exports = append(exports, byte(len(fun.Name))) // name length
			exports = append(exports, []byte(fun.Name)...) // name
			exports = append(exports, 0x00, byte(i))       // export kind, export funcindex

			funcbody := bytes.NewBuffer([]byte{
				0x00, // local decl count
			})
			for i, s := range fun.Body {
				err := operate(funcbody, fun.Name, funcSymbolTable, s.Exp.Operator, s.Exp.Left, s.Exp.Right)
				if err != nil {
					return nil, err
				}
				if i != len(fun.Body)-1 {
					funcbody.WriteByte(operandMap["drop"])
				}
			}
			if len(fun.ReturnTypes) == 0 {
				funcbody.WriteByte(operandMap["drop"])
			}
			funcBodys = append(funcBodys, byte(len(funcbody.Bytes())+1)) // funcbody size
			funcBodys = append(funcBodys, funcbody.Bytes()...)
			funcBodys = append(funcBodys, FUNC_END)
			// if i == 0 {
			// 	return fmt.Errorf(fmt.Sprintf("%+v", funcbody.Bytes()))
			// }
		}
		buffer.Write([]byte{0x01, byte(len(types) + 1), byte(len(p.Funs))}) // Type section code, section size, num types, type data
		buffer.Write(types)
		buffer.Write([]byte{0x03, byte(len(funcs) + 1), byte(len(p.Funs))}) // funcsig section code, section size, num types, type data
		buffer.Write(funcs)
		buffer.Write([]byte{0x07, byte(len(exports) + 1), byte(len(p.Funs))}) // exports section code, section size, num exports
		buffer.Write(exports)
		buffer.Write([]byte{0x0a, byte(len(funcBodys) + 1), byte(len(p.Funs))}) // funcbody section code, section size, num functions
		buffer.Write(funcBodys)
	}
	return buffer, nil
}

// ;; A function that takes an `i64` and returns
// ;; three `i32`s.
// (func(param i64) (result i32 i32 i32)
//   ...)

// ;; A loop that consumes an `i32` stack value
// ;; at the start of each iteration.
// loop (param i32)
//   ...
// end

// ;; A block that produces two `i32` stack values.
// block (result i32 i32)
//   ...
// end

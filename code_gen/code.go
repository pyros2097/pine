package code_gen

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"yum/ast"
)

var typeMap = map[string]byte{
	"num": 0x7c,
	"i32": 0x7f,
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
	Index      int
	Name       string
	Params     map[string]FuncParam
	ReturnType string
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

func emitExpression(buffer *bytes.Buffer, funcName string, funcSymbolTable map[string]FuncData, operation *string, l *ast.Literal, r *ast.Expression) error {
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
				err := emitExpression(buffer, funcName, funcSymbolTable, r.Operator, nil, r.Right)
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
		err := emitExpression(buffer, funcName, funcSymbolTable, r.Operator, nil, r.Right)
		if err != nil {
			return err
		}
	}
	return nil
}

func emitTypes(buffer *bytes.Buffer, fun FuncData) error {
	buffer.WriteByte(0x60)                  // type func
	buffer.WriteByte(byte(len(fun.Params))) // num params
	for _, param := range fun.Params {      // i32, i32
		t, ok := typeMap[param.Type]
		if !ok {
			return fmt.Errorf("func'" + fun.Name + "' parameter '" + param.Name + "' type '" + param.Type + "' does not exist")
		}
		buffer.WriteByte(t)
	}
	if fun.ReturnType != "" {
		t, ok := typeMap[fun.ReturnType]
		if !ok {
			return fmt.Errorf("func'" + fun.Name + "' return type '" + fun.ReturnType + "' does not exist")
		}
		buffer.WriteByte(0x01) // num results
		buffer.WriteByte(t)    // i32
	} else {
		buffer.WriteByte(0x00) // num results
	}
	// panic(repr.String(buffer.Bytes()))
	return nil
}

func GenerateCode(tree *ast.Ast) (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write([]byte{0x00, 0x61, 0x73, 0x6d}) // WASM_BINARY_MAGIC
	buffer.Write([]byte{0x01, 0x00, 0x00, 0x00}) // WASM_BINARY_VERSION

	typesBuffer := bytes.NewBuffer(nil)
	funcs := []byte{}
	exports := []byte{}
	funcBodys := []byte{}
	externFuncsTable := map[string]FuncData{}
	funcSymbolTable := map[string]FuncData{}
	for _, p := range tree.Modules {
		for _, a := range p.ExternFuncs {
			returnType := ""
			// TODO: we need to remove the last item since '\n' gets captured somehow needs to be fixed later
			if len(a.ReturnTypes) > 1 {
				returnType = a.ReturnTypes[0]
			}
			externFuncsTable[a.Name] = FuncData{
				Index:      len(externFuncsTable),
				Name:       a.Name,
				Params:     map[string]FuncParam{},
				ReturnType: returnType,
			}
			for pi, param := range a.Parameters {
				externFuncsTable[a.Name].Params[param.Name] = FuncParam{
					Index: pi,
					Name:  param.Name,
					Type:  param.Type.Name,
				}
			}
			err := emitTypes(typesBuffer, externFuncsTable[a.Name])
			if err != nil {
				return nil, fmt.Errorf("Failed to emitTypes %v", err)
			}
		}
		for _, fun := range p.Funs {
			fnIndex := len(funcSymbolTable)
			returnType := ""
			if len(fun.ReturnTypes) > 1 {
				returnType = fun.ReturnTypes[0]
			}
			funcData := FuncData{
				Index:      fnIndex,
				Name:       fun.Name,
				Params:     map[string]FuncParam{},
				ReturnType: returnType,
			}
			funcSymbolTable[fun.Name] = funcData
			for pi, param := range fun.Parameters {
				funcSymbolTable[fun.Name].Params[param.Name] = FuncParam{
					Index: pi,
					Name:  param.Name,
					Type:  param.Type.Name,
				}
			}
			err := emitTypes(typesBuffer, funcSymbolTable[fun.Name])
			if err != nil {
				return nil, fmt.Errorf("Failed to emitTypes %v", err)
			}
			funcs = append(funcs, byte(len(externFuncsTable)+fnIndex)) // function 0 signature index
			exports = append(exports, byte(len(fun.Name)))             // name length
			exports = append(exports, []byte(fun.Name)...)             // name
			exports = append(exports, 0x00, byte(fnIndex))             // export kind, export funcindex

			funcbody := bytes.NewBuffer([]byte{
				0x00, // local decl count
			})
			for i, s := range fun.Body {
				err := emitExpression(funcbody, fun.Name, funcSymbolTable, s.Exp.Operator, s.Exp.Left, s.Exp.Right)
				if err != nil {
					return nil, err
				}
				if i != len(fun.Body)-1 {
					funcbody.WriteByte(operandMap["drop"])
				}
			}
			if funcData.ReturnType == "" {
				funcbody.WriteByte(operandMap["drop"])
			}
			funcBodys = append(funcBodys, byte(len(funcbody.Bytes())+1)) // funcbody size
			funcBodys = append(funcBodys, funcbody.Bytes()...)
			funcBodys = append(funcBodys, FUNC_END)
			// if fnIndex == 1 {
			// 	return nil, fmt.Errorf(fmt.Sprintf("%+v", funcbody.Bytes()))
			// }
		}
	}
	buffer.Write([]byte{0x01, byte(typesBuffer.Len() + 1), byte(len(externFuncsTable) + len(funcSymbolTable))}) // Type section code, section size, num types, type data
	buffer.Write(typesBuffer.Bytes())
	// return nil, fmt.Errorf(fmt.Sprintf("%x", typesBuffer.Bytes()))
	buffer.Write([]byte{0x03, byte(len(funcs) + 1), byte(len(funcSymbolTable))}) // funcsig section code, section size, num types, type data
	buffer.Write(funcs)
	buffer.Write([]byte{0x07, byte(len(exports) + 1), byte(len(funcSymbolTable))}) // exports section code, section size, num exports
	buffer.Write(exports)
	buffer.Write([]byte{0x0a, byte(len(funcBodys) + 1), byte(len(funcSymbolTable))}) // funcbody section code, section size, num functions
	buffer.Write(funcBodys)
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

// ; section "Import" (2)
// 0000016: 02                                        ; section code
// 0000017: 00                                        ; section size (guess)
// 0000018: 01                                        ; num imports
// ; import header 0
// 0000019: 0d                                        ; string length
// 000001a: 7761 7369 5f75 6e73 7461 626c 65         wasi_unstable  ; import module name
// 0000027: 08                                        ; string length
// 0000028: 6664 5f77 7269 7465                      fd_write  ; import field name
// 0000030: 00                                        ; import kind
// 0000031: 00                                        ; import signature index
// 0000017: 1a                                        ; FIXUP section size

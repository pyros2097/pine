package code_gen

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"yum/ast"

	"github.com/alecthomas/repr"
)

func EncodeSleb128(b *bytes.Buffer, v int32) {
	for {
		c := uint8(v & 0x7f)
		s := uint8(v & 0x40)
		v >>= 7
		if (v != -1 || s == 0) && (v != 0 || s != 0) {
			c |= 0x80
		}
		b.WriteByte(c)
		if c&0x80 == 0 {
			break
		}
	}
}

func writeMemoryData(b *bytes.Buffer, address byte, v int32) {
	b.WriteByte(operandMap["i32.const"])
	b.WriteByte(address)
	b.WriteByte(operandMap["i32.const"])
	EncodeSleb128(b, v)
	b.WriteByte(operandMap["i32.store"])
	b.WriteByte(byte(0x02))
	b.WriteByte(byte(0x00))
}

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
	"i32.load":      0x28,
	"f64.load":      0x2b,
	"i32.store":     0x36,
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
	if l != nil && l.Str != nil {
		// (i32.store (i32.const 0) (i32.const 8))  ;; iov.iov_base - This is a pointer to the start of the 'hello world\n' string
		// (i32.store (i32.const 4) (i32.const 12))  ;; iov.iov_len - The length of the 'hello world\n' string
		// (i32.store (i32.const 8) (i32.const 0x6c6c6568))
		// (i32.store (i32.const 12) (i32.const 0x6f77206f))
		// (i32.store (i32.const 16) (i32.const 0x0a646c72))
		// buffer.WriteByte(operandMap["i32.const"])
		// buffer.WriteByte(byte(0))
	}
	// TODO: remove r
	if r != nil && r.Left != nil && r.Left.Num != nil {
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
				if v.Str != nil {
					str := *v.Str
					data := []byte(str)
					if len(data)%4 == 1 {
						data = append(data, 0)
					}
					if len(data)%4 == 2 {
						data = append(data, 0)
					}
					if len(data)%4 == 3 {
						data = append(data, '\n')
					}

					// (i32.store (i32.const 0) (i32.const 8))  ;; iov.iov_base - This is a pointer to the start of the 'hello world\n' string
					// (i32.store (i32.const 4) (i32.const 12))  ;; iov.iov_len - The length of the 'hello world\n' string
					// (i32.store (i32.const 8) (i32.const 0x6c6c6568))
					// (i32.store (i32.const 12) (i32.const 0x6f77206f))
					// (i32.store (i32.const 16) (i32.const 0x0a646c72))
					writeMemoryData(buffer, 0, 8)
					writeMemoryData(buffer, 4, int32(len(string(data))))
					startAddress := 0
					startData := 0
					for i := range data {
						index := i + 1
						if index%4 == 0 {
							startAddress = index + 4
							startData = index - 4
							remainingData := data[startData:index]
							writeMemoryData(buffer, byte(startAddress), int32(binary.LittleEndian.Uint32(remainingData)))
						}
					}
				}
			}
			for _, v := range r.Left.Params {
				if v.Num != nil {
					buffer.WriteByte(0x41)
					// TODO: fix this
					// this has to be float64 or int32
					buffer.WriteByte(byte(int32(*v.Num)))
				}
				if v.Reference != nil {
					vref, ok := funcSymbolTable[funcName].Params[*v.Reference]
					if !ok {
						return fmt.Errorf("func'" + funcName + "' trying to call '" + callFunc.Name + "' with non-existing variable " + *v.Reference)
					}
					buffer.WriteByte(operandMap["get_local"])
					buffer.WriteByte(byte(vref.Index))
				}
				if v.Str != nil {
					buffer.WriteByte(operandMap["i32.const"])
					buffer.WriteByte(byte(0))
				}
			}
			buffer.WriteByte(operandMap["call"])
			buffer.WriteByte(byte(callFunc.Index))
			// return fmt.Errorf("%s %d", callFunc.Name, callFunc.Index)
			if r.Operator != nil {
				err := emitExpression(buffer, funcName, funcSymbolTable, r.Operator, nil, r.Right)
				if err != nil {
					return err
				}
			}
		}
	}
	if r != nil && r.Left != nil && r.Left.Reference != nil {
		ref, ok := funcSymbolTable[funcName].Params[*r.Left.Reference]
		if !ok {
			return fmt.Errorf("func'" + funcName + "' parameter '" + *r.Left.Reference + "' does not exist")
		}
		buffer.WriteByte(operandMap["get_local"])
		buffer.WriteByte(byte(ref.Index))
		buffer.WriteByte(operandMap[*operation])
	}
	if r != nil && r.Operator != nil {
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

// _pointer = a: string -> i32
// (i32.store (i32.const 0) (i32.const 8)) // iov.iov_base - This is a pointer to the start of the 'hello world\n' string

// _length = a: string -> i32
// (i32.store (i32.const 4) (i32.const 12)) // iov.iov_len - The length of the 'hello world\n' string

func GenerateCode(tree *ast.Ast) (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write([]byte{0x00, 0x61, 0x73, 0x6d}) // WASM_BINARY_MAGIC
	buffer.Write([]byte{0x01, 0x00, 0x00, 0x00}) // WASM_BINARY_VERSION

	typesSection := bytes.NewBuffer(nil)
	importsSection := bytes.NewBuffer(nil)
	funcs := []byte{}
	exports := []byte{}
	funcBodys := []byte{}
	externFuncsCount := 0
	funcsCount := 0
	funcSymbolTable := map[string]FuncData{}
	for _, p := range tree.Modules {
		for _, a := range p.ExternFuncs {
			funcSymbolTable[a.Name] = FuncData{
				Index:      externFuncsCount,
				Name:       a.Name,
				Params:     map[string]FuncParam{},
				ReturnType: a.ReturnType,
			}
			for pi, param := range a.Parameters {
				funcSymbolTable[a.Name].Params[param.Name] = FuncParam{
					Index: pi,
					Name:  param.Name,
					Type:  param.Type.Name,
				}
			}
			err := emitTypes(typesSection, funcSymbolTable[a.Name])
			if err != nil {
				return nil, fmt.Errorf("Failed to emitTypes %v", err)
			}
			importsSection.WriteByte(byte(len(p.Name)))      // module name length
			importsSection.Write([]byte(p.Name))             // module name
			importsSection.WriteByte(byte(len(a.Name)))      // fn name length
			importsSection.Write([]byte(a.Name))             // fn name
			importsSection.WriteByte(0x00)                   // import kind
			importsSection.WriteByte(byte(externFuncsCount)) // import signature index
			externFuncsCount += 1
		}
		for _, fun := range p.Funs {
			returnType := ""
			if len(fun.ReturnTypes) > 1 {
				returnType = fun.ReturnTypes[0]
			}
			funcSymbolTable[fun.Name] = FuncData{
				Index:      externFuncsCount + funcsCount,
				Name:       fun.Name,
				Params:     map[string]FuncParam{},
				ReturnType: returnType,
			}
			for pi, param := range fun.Parameters {
				funcSymbolTable[fun.Name].Params[param.Name] = FuncParam{
					Index: pi,
					Name:  param.Name,
					Type:  param.Type.Name,
				}
			}
			err := emitTypes(typesSection, funcSymbolTable[fun.Name])
			if err != nil {
				return nil, fmt.Errorf("Failed to emitTypes %v", err)
			}
			funcs = append(funcs, byte(externFuncsCount+funcsCount))           // function 0 signature index
			exports = append(exports, byte(len(fun.Name)))                     // name length
			exports = append(exports, []byte(fun.Name)...)                     // name
			exports = append(exports, 0x00, byte(externFuncsCount+funcsCount)) // export kind, export funcindex

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
			if funcSymbolTable[fun.Name].ReturnType == "" {
				funcbody.WriteByte(operandMap["drop"])
			}
			funcBodys = append(funcBodys, byte(len(funcbody.Bytes())+1)) // funcbody size
			funcBodys = append(funcBodys, funcbody.Bytes()...)
			funcBodys = append(funcBodys, FUNC_END)
			// if fnIndex == 1 {
			// 	return nil, fmt.Errorf(fmt.Sprintf("%+v", funcbody.Bytes()))
			// }
			funcsCount += 1
		}
	}
	repr.Println(funcSymbolTable)
	buffer.Write([]byte{0x01, byte(typesSection.Len() + 1), byte(externFuncsCount + funcsCount)}) // Type section code, section size, num types, type data
	buffer.Write(typesSection.Bytes())
	buffer.Write([]byte{0x02, byte(importsSection.Len() + 1), byte(externFuncsCount)}) // Imports section, section size, num imports, type data
	buffer.Write(importsSection.Bytes())
	buffer.Write([]byte{0x03, byte(len(funcs) + 1), byte(funcsCount)}) // Func Sig section code, section size, num types, type data
	buffer.Write(funcs)
	buffer.Write([]byte{0x05, 0x03, 0x01, 0x00, 0x01})                   // Memory section code, section size, num memories, flags, initial (1 page 64KB)
	buffer.Write([]byte{0x07, byte(len(exports) + 1), byte(funcsCount)}) // exports section code, section size, num exports
	buffer.Write(exports)
	buffer.Write([]byte{0x0a, byte(len(funcBodys) + 1), byte(funcsCount)}) // funcbody section code, section size, num functions
	buffer.Write(funcBodys)
	// buffer.Write([]byte{0x0b, byte(len(dataBodys) + 1), byte(len(datas))}) // Data section code, section size, num data segments
	// buffer.Write(dataBodys)
	return buffer, nil
}

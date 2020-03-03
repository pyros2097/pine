package code_gen

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"yum/ast"
	"yum/code_gen/op"

	"github.com/alecthomas/repr"
)

func emitString(b *bytes.Buffer, address byte, v int32) {
	b.WriteByte(op.I32_CONST)
	b.WriteByte(address)
	b.WriteByte(op.I32_CONST)
	encodeSleb128(b, v)
	b.WriteByte(op.I32_STORE)
	b.WriteByte(byte(0x02))
	b.WriteByte(byte(0x00))
}

var typeMap = map[string]byte{
	"num": op.F64,
	"i32": op.I32,
}

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

func emitOperation(buffer *bytes.Buffer, funcName, operation string) error {
	switch operation {
	case "+":
		buffer.WriteByte(op.F64_ADD)
	case "-":
		buffer.WriteByte(op.F64_SUB)
	case "*":
		buffer.WriteByte(op.F64_MUL)
	case "/":
		buffer.WriteByte(op.F64_DIV)
	default:
		return fmt.Errorf("func '%s' operation '%s' is invalid", funcName, operation)
	}
	return nil
}

func emitExpression(buffer *bytes.Buffer, funcName string, funcSymbolTable map[string]FuncData, operation *string, l *ast.Literal, r *ast.Expression) error {
	if l != nil && l.Num != nil {
		buffer.WriteByte(op.F64_CONST)
		buffer.Write(float64ToByte(*l.Num))
	}
	if l != nil && l.Str != nil {
		// (i32.store (i32.const 0) (i32.const 8))  ;; iov.iov_base - This is a pointer to the start of the 'hello world\n' string
		// (i32.store (i32.const 4) (i32.const 12))  ;; iov.iov_len - The length of the 'hello world\n' string
		// (i32.store (i32.const 8) (i32.const 0x6c6c6568))
		// (i32.store (i32.const 12) (i32.const 0x6f77206f))
		// (i32.store (i32.const 16) (i32.const 0x0a646c72))
		// buffer.WriteByte(op.I32_CONST)
		// buffer.WriteByte(byte(0))
	}
	// TODO: remove r
	if r != nil && r.Left != nil && r.Left.Num != nil {
		buffer.WriteByte(op.F64_CONST)
		buffer.Write(float64ToByte(*r.Left.Num))
		emitOperation(buffer, funcName, *operation)
	}
	if l != nil && l.Reference != nil {
		if operation != nil {
			ref, ok := funcSymbolTable[funcName].Params[*l.Reference]
			if !ok {
				return fmt.Errorf("func '%s' parameter '%s' does not exist", funcName, *l.Reference)
			}
			buffer.WriteByte(op.GET_LOCAL)
			buffer.WriteByte(byte(ref.Index))
		} else {
			// funcCALL
			callFunc, ok := funcSymbolTable[*l.Reference]
			if !ok {
				return fmt.Errorf("func '%s' trying to call non-existing func '%s'", funcName, *l.Reference)
			}
			if len(callFunc.Params) != len(r.Left.Params) {
				return fmt.Errorf("func '%s' trying to call function '%s' with extra parameters", funcName, callFunc.Name)
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
					emitString(buffer, 0, 8)
					emitString(buffer, 4, int32(len(string(data))))
					startAddress := 0
					startData := 0
					for i := range data {
						index := i + 1
						if index%4 == 0 {
							startAddress = index + 4
							startData = index - 4
							remainingData := data[startData:index]
							emitString(buffer, byte(startAddress), int32(binary.LittleEndian.Uint32(remainingData)))
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
					buffer.WriteByte(op.GET_LOCAL)
					buffer.WriteByte(byte(vref.Index))
				}
				if v.Str != nil {
					buffer.WriteByte(op.I32_CONST)
					buffer.WriteByte(byte(0))
				}
			}
			buffer.WriteByte(op.CALL)
			buffer.WriteByte(byte(callFunc.Index))
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
		buffer.WriteByte(op.GET_LOCAL)
		buffer.WriteByte(byte(ref.Index))
		emitOperation(buffer, funcName, *operation)
	}
	if r != nil && r.Operator != nil {
		err := emitExpression(buffer, funcName, funcSymbolTable, r.Operator, nil, r.Right)
		if err != nil {
			return err
		}
	}
	return nil
}

// _pointer = a: string -> i32
// (i32.store (i32.const 0) (i32.const 8)) // iov.iov_base - This is a pointer to the start of the 'hello world\n' string

// _length = a: string -> i32
// (i32.store (i32.const 4) (i32.const 12)) // iov.iov_len - The length of the 'hello world\n' string

type Emitter struct {
	TypesSection     *bytes.Buffer
	ImportsSection   *bytes.Buffer
	FuncsSection     *bytes.Buffer
	ExportsSection   *bytes.Buffer
	FuncsBodySection *bytes.Buffer
	Tree             *ast.Ast
	funcSymbolTable  map[string]FuncData
}

func NewEmitter(tree *ast.Ast) Emitter {
	return Emitter{
		TypesSection:     bytes.NewBuffer(nil),
		ImportsSection:   bytes.NewBuffer(nil),
		FuncsSection:     bytes.NewBuffer(nil),
		ExportsSection:   bytes.NewBuffer(nil),
		FuncsBodySection: bytes.NewBuffer(nil),
		funcSymbolTable:  map[string]FuncData{},
		Tree:             tree,
	}
}

func (e Emitter) EmitTypes(fun FuncData) error {
	e.TypesSection.WriteByte(0x60)                  // type func
	e.TypesSection.WriteByte(byte(len(fun.Params))) // num params
	for _, param := range fun.Params {              // i32, i32
		t, ok := typeMap[param.Type]
		if !ok {
			return fmt.Errorf("func'" + fun.Name + "' parameter '" + param.Name + "' type '" + param.Type + "' does not exist")
		}
		e.TypesSection.WriteByte(t)
	}
	if fun.ReturnType != "" {
		t, ok := typeMap[fun.ReturnType]
		if !ok {
			return fmt.Errorf("func'" + fun.Name + "' return type '" + fun.ReturnType + "' does not exist")
		}
		e.TypesSection.WriteByte(0x01) // num results
		e.TypesSection.WriteByte(t)    // i32
	} else {
		e.TypesSection.WriteByte(0x00) // num results
	}
	// panic(repr.String(buffer.Bytes()))
	return nil
}

func (e Emitter) EmitImports(moduleName, funcName string, typeIndex int) {
	e.ImportsSection.WriteByte(byte(len(moduleName))) // module name length
	e.ImportsSection.Write([]byte(moduleName))        // module name
	e.ImportsSection.WriteByte(byte(len(funcName)))   // fn name length
	e.ImportsSection.Write([]byte(funcName))          // fn name
	e.ImportsSection.WriteByte(0x00)                  // import kind
	e.ImportsSection.WriteByte(byte(typeIndex))       // import signature index
}

func (e Emitter) EmitFuncs(moduleName, funcName string, typeIndex int) {
}

func (e Emitter) EmitAll() (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer(nil)
	buffer.Write([]byte{0x00, 0x61, 0x73, 0x6d}) // WASM_BINARY_MAGIC
	buffer.Write([]byte{0x01, 0x00, 0x00, 0x00}) // WASM_BINARY_VERSION

	exports := []byte{}
	funcBodys := []byte{}
	externFuncsCount := 0
	funcsCount := 0
	funcSymbolTable := map[string]FuncData{}
	for _, p := range e.Tree.Modules {
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
			err := e.EmitTypes(funcSymbolTable[a.Name])
			if err != nil {
				return nil, fmt.Errorf("Failed to emitTypes %v", err)
			}
			e.EmitImports(p.Name, a.Name, externFuncsCount)
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
			err := e.EmitTypes(funcSymbolTable[fun.Name])
			if err != nil {
				return nil, fmt.Errorf("Failed to emitTypes %v", err)
			}
			e.FuncsSection.WriteByte(byte(externFuncsCount + funcsCount))      // function 0 signature index
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
					funcbody.WriteByte(op.DROP)
				}
			}
			if funcSymbolTable[fun.Name].ReturnType == "" {
				funcbody.WriteByte(op.DROP)
			}
			funcBodys = append(funcBodys, byte(len(funcbody.Bytes())+1)) // funcbody size
			funcBodys = append(funcBodys, funcbody.Bytes()...)
			funcBodys = append(funcBodys, op.FUNC_END)
			// if fnIndex == 1 {
			// 	return nil, fmt.Errorf(fmt.Sprintf("%+v", funcbody.Bytes()))
			// }
			funcsCount += 1
		}
	}
	repr.Println(funcSymbolTable)
	buffer.Write([]byte{0x01, byte(e.TypesSection.Len() + 1), byte(externFuncsCount + funcsCount)}) // Type section code, section size, num types, type data
	buffer.Write(e.TypesSection.Bytes())
	buffer.Write([]byte{0x02, byte(e.ImportsSection.Len() + 1), byte(externFuncsCount)}) // Imports section, section size, num imports, type data
	buffer.Write(e.ImportsSection.Bytes())
	buffer.Write([]byte{0x03, byte(e.FuncsSection.Len() + 1), byte(funcsCount)}) // Func Sig section code, section size, num types, type data
	buffer.Write(e.FuncsSection.Bytes())
	buffer.Write([]byte{0x05, 0x03, 0x01, 0x00, 0x01})                   // Memory section code, section size, num memories, flags, initial (1 page 64KB)
	buffer.Write([]byte{0x07, byte(len(exports) + 1), byte(funcsCount)}) // exports section code, section size, num exports
	buffer.Write(exports)
	buffer.Write([]byte{0x0a, byte(len(funcBodys) + 1), byte(funcsCount)}) // funcbody section code, section size, num functions
	buffer.Write(funcBodys)
	return buffer, nil
}

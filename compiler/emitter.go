package compiler

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/pyros2097/pine/compiler/op"
)

type FuncData struct {
	Index      int
	Name       string
	Params     map[string]*FuncParam
	ReturnType string
}

type FuncParam struct {
	Index int
	Name  string
	Type  string
	ptr   int
}

type TypeData struct {
	Name  string
	Type  string
	Value string
}

type Emitter struct {
	TypesSection     *bytes.Buffer
	ImportsSection   *bytes.Buffer
	FuncsSection     *bytes.Buffer
	MemorySection    *bytes.Buffer
	GlobalsSection   *bytes.Buffer
	ExportsSection   *bytes.Buffer
	FuncsBodySection *bytes.Buffer
	Module           *Module
	funcs            map[string]*FuncData
	types            map[string]*TypeData
	typeOpCode       map[string]byte
	externFuncsCount int
	funcsCount       int
	globalsCount     int
	initial          bool
	emitJs           bool
}

func NewEmitter(module *Module) *Emitter {
	return &Emitter{
		TypesSection:     bytes.NewBuffer(nil),
		ImportsSection:   bytes.NewBuffer(nil),
		FuncsSection:     bytes.NewBuffer(nil),
		MemorySection:    bytes.NewBuffer(nil),
		GlobalsSection:   bytes.NewBuffer(nil),
		ExportsSection:   bytes.NewBuffer(nil),
		FuncsBodySection: bytes.NewBuffer(nil),
		funcs:            map[string]*FuncData{},
		types: map[string]*TypeData{
			"int": &TypeData{ // type int i32
				Name:  "int",
				Type:  "alias",
				Value: "i32",
			},
			"float": &TypeData{ // type float f32
				Name:  "float",
				Type:  "alias",
				Value: "f32",
			},
			"string": &TypeData{ // type i32 for reference types
				Name:  "string",
				Type:  "type",
				Value: "i32",
			},
		},
		typeOpCode: map[string]byte{
			"i32": op.I32,
			"f32": op.F32,
		},
		Module:           module,
		externFuncsCount: 0,
		funcsCount:       0,
		initial:          false,
	}
}

func (e *Emitter) EmitTypes(name string) error {
	fun := e.funcs[name]
	e.TypesSection.WriteByte(0x60)                  // type func
	e.TypesSection.WriteByte(byte(len(fun.Params))) // num params
	for _, param := range fun.Params {              // i32, i32
		t, ok := e.types[param.Type]
		if !ok {
			return fmt.Errorf("function '" + fun.Name + "' parameter '" + param.Name + "' type '" + param.Type + "' does not exist")
		}
		opCode, ok := e.typeOpCode[t.Value]
		if !ok {
			return fmt.Errorf("function '" + fun.Name + "' parameter '" + param.Name + "' opcode does not exist")
		}
		e.TypesSection.WriteByte(opCode)
	}
	if fun.ReturnType != "" {
		t, ok := e.types[fun.ReturnType]
		if !ok {
			return fmt.Errorf("function '" + fun.Name + "' return type '" + fun.ReturnType + "' does not exist")
		}
		opCode, ok := e.typeOpCode[t.Value]
		if !ok {
			return fmt.Errorf("function '" + fun.Name + "' return type '" + fun.ReturnType + "' opcode does not exist")
		}
		e.TypesSection.WriteByte(0x01) // num results
		e.TypesSection.WriteByte(opCode)
	} else {
		e.TypesSection.WriteByte(0x00) // num results
	}
	return nil
}

func (e *Emitter) EmitImports(moduleName, funcName string, typeIndex int) {
	e.ImportsSection.WriteByte(byte(len(moduleName))) // module name length
	e.ImportsSection.Write([]byte(moduleName))        // module name
	e.ImportsSection.WriteByte(byte(len(funcName)))   // fn name length
	e.ImportsSection.Write([]byte(funcName))          // fn name
	e.ImportsSection.WriteByte(0x00)                  // import kind
	e.ImportsSection.WriteByte(byte(typeIndex))       // import signature index
}

func (e *Emitter) EmitFuncs(name string, typeIndex int) {
	e.FuncsSection.WriteByte(byte(typeIndex))
	e.ExportsSection.WriteByte(byte(len(name))) // name length
	e.ExportsSection.Write([]byte(name))
	e.ExportsSection.WriteByte(0x00)            // export kind
	e.ExportsSection.WriteByte(byte(typeIndex)) // export funcindex
}

func (e *Emitter) EmitExports(typeIndex int) {
	e.FuncsSection.WriteByte(byte(typeIndex))
}

func (e *Emitter) EmitFuncBody(name string, body []*Statement) error {
	// buf := bytes.NewBuffer([]byte{
	// 	0x00, // local decl count
	// })
	// for i, s := range body {
	// 	err := e.emitExpression(buf, name, s.Exp.Operator, s.Exp.Left, s.Exp.Right)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if i != len(body)-1 {
	// 		buf.WriteByte(op.DROP)
	// 	}
	// }
	// if e.funcs[name].ReturnType == "" {
	// 	buf.WriteByte(op.DROP)
	// }
	// encodeSleb128(e.FuncsBodySection, int32(buf.Len()+1)) // funcbody size
	// e.FuncsBodySection.Write(buf.Bytes())
	// e.FuncsBodySection.WriteByte(op.END)
	return nil
}

func (e *Emitter) emitExpression(buffer *bytes.Buffer, funcName string, operation *string, l *Literal, r *Expression) error {
	if l != nil {
		if l.Int != nil {
			buffer.WriteByte(op.I32_CONST)
			encodeSleb128(buffer, int32(*l.Int))
		}
		if l.Float != nil {
			buffer.WriteByte(op.F32_CONST)
			buffer.Write(float32ToByte(*l.Float))
		}
		if l.Str != nil {
			// (i32.store (i32.const 0) (i32.const 8))  ;; iov.iov_base - This is a pointer to the start of the 'hello world\n' string
			// (i32.store (i32.const 4) (i32.const 12))  ;; iov.iov_len - The length of the 'hello world\n' string
			// (i32.store (i32.const 8) (i32.const 0x6c6c6568))
			// (i32.store (i32.const 12) (i32.const 0x6f77206f))
			// (i32.store (i32.const 16) (i32.const 0x0a646c72))
			// buffer.WriteByte(op.I32_CONST)
			// buffer.WriteByte(byte(0))
		}
	}
	// TODO: remove r
	if r != nil && r.Left != nil {
		if r.Left.Int != nil {
			buffer.WriteByte(op.I32_CONST)
			encodeSleb128(buffer, int32(*r.Left.Int))
			e.emitOperation(buffer, funcName, *operation, "int")
		}
		if r.Left.Float != nil {
			buffer.WriteByte(op.F32_CONST)
			buffer.Write(float32ToByte(*r.Left.Float))
			e.emitOperation(buffer, funcName, *operation, "float")
		}
	}
	if l != nil && l.Reference != nil {
		if operation != nil {
			ref, ok := e.funcs[funcName].Params[*l.Reference]
			if !ok {
				return fmt.Errorf("func '%s' parameter '%s' does not exist", funcName, *l.Reference)
			}
			buffer.WriteByte(op.GET_LOCAL)
			buffer.WriteByte(byte(ref.Index))
		} else {
			// funcCALL
			callFunc, ok := e.funcs[*l.Reference]
			if !ok {
				return fmt.Errorf("func '%s' trying to call non-existing func '%s'", funcName, *l.Reference)
			}
			if len(callFunc.Params) != len(r.Left.Params) {
				return fmt.Errorf("func '%s' trying to call function '%s' with extra parameters", funcName, callFunc.Name)
			}
			for i, v := range r.Left.Params {
				for _, p := range callFunc.Params {
					if i == p.Index {
						if v.Str != nil {
							p.ptr = e.emitString(buffer, *v.Str)
						}
						break
					}
				}
			}
			for i, v := range r.Left.Params {
				if v.Reference != nil {
					vref, ok := e.funcs[funcName].Params[*v.Reference]
					if !ok {
						return fmt.Errorf("function '" + funcName + "' trying to call '" + callFunc.Name + "' with non-existing variable " + *v.Reference)
					}
					buffer.WriteByte(op.GET_LOCAL)
					buffer.WriteByte(byte(vref.Index))
				}
				for _, p := range callFunc.Params {
					if i == p.Index {
						if v.Int != nil {
							buffer.WriteByte(op.I32_CONST)
							encodeSleb128(buffer, int32(*v.Int))
						}
						if v.Float != nil {
							buffer.WriteByte(op.F32_CONST)
							buffer.Write(float32ToByte(*v.Float))
						}
						if v.Str != nil {
							buffer.WriteByte(op.I32_CONST)
							buffer.WriteByte(byte(p.ptr))
						}
						break
					}
				}
			}
			buffer.WriteByte(op.CALL)
			buffer.WriteByte(byte(callFunc.Index))
			if r.Operator != nil {
				err := e.emitExpression(buffer, funcName, r.Operator, nil, r.Right)
				if err != nil {
					return err
				}
			}
		}
	}
	if r != nil && r.Left != nil && r.Left.Reference != nil {
		ref, ok := e.funcs[funcName].Params[*r.Left.Reference]
		if !ok {
			return fmt.Errorf("function '" + funcName + "' parameter '" + *r.Left.Reference + "' does not exist")
		}
		buffer.WriteByte(op.GET_LOCAL)
		buffer.WriteByte(byte(ref.Index))
		e.emitOperation(buffer, funcName, *operation, ref.Type)
	}
	if r != nil && r.Operator != nil {
		err := e.emitExpression(buffer, funcName, r.Operator, nil, r.Right)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Emitter) emitOperation(buffer *bytes.Buffer, funcName, operation, ttype string) error {
	switch ttype {
	case "int":
		switch operation {
		case "+":
			buffer.WriteByte(op.I32_ADD)
		case "-":
			buffer.WriteByte(op.I32_SUB)
		case "*":
			buffer.WriteByte(op.I32_MUL)
		case "/":
			buffer.WriteByte(op.I32_DIV_S)
		default:
			return fmt.Errorf("func '%s' operation '%s' is invalid for '%s'", funcName, operation, ttype)
		}
	case "float":
		switch operation {
		case "+":
			buffer.WriteByte(op.F32_ADD)
		case "-":
			buffer.WriteByte(op.F32_SUB)
		case "*":
			buffer.WriteByte(op.F32_MUL)
		case "/":
			buffer.WriteByte(op.F32_DIV)
		default:
			return fmt.Errorf("func '%s' operation '%s' is invalid for '%s'", funcName, operation, ttype)
		}
	default:
		return fmt.Errorf("func '%s' operation '%s' is invalid for '%s'", funcName, operation, ttype)
	}
	return nil
}

func (e *Emitter) emitHeapNext(b *bytes.Buffer) {
	// (set_global $heap-next (i32.add (get_global $heap-next) (i32.const 4)))
	// (get_global $heap-next)
	size := 4
	if e.initial == false {
		e.initial = true
		size = 0
	}
	b.WriteByte(op.GET_GLOBAL)
	b.WriteByte(byte(0x00)) // global heap_index
	b.WriteByte(op.I32_CONST)
	b.WriteByte(byte(size))
	b.WriteByte(op.I32_ADD)
	b.WriteByte(byte(op.SET_GLOBAL))
	b.WriteByte(byte(0x00))
	b.WriteByte(byte(op.GET_GLOBAL))
	b.WriteByte(byte(0x00))

}

func (e *Emitter) emitStore(b *bytes.Buffer, v int32) {
	//  (i32.store (get_global $heap-next) (i32.const 8))
	e.emitHeapNext(b)
	// b.WriteByte(op.I32_CONST)
	// b.WriteByte(address)
	b.WriteByte(op.I32_CONST)
	encodeSleb128(b, v)
	b.WriteByte(op.I32_STORE)
	b.WriteByte(byte(0x02))
	b.WriteByte(byte(0x00))
}

// (i32.store (i32.const 0) (i32.const 8))  ;; iov.iov_base - This is a pointer to the start of the 'hello world\n' string
// (i32.store (i32.const 4) (i32.const 12))  ;; iov.iov_len - The length of the 'hello world\n' string
// (i32.store (i32.const 8) (i32.const 0x6c6c6568))
// (i32.store (i32.const 12) (i32.const 0x6f77206f))
// (i32.store (i32.const 16) (i32.const 0x0a646c72))
func (e *Emitter) emitString(b *bytes.Buffer, str string) int {
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

	e.emitStore(b, 8)                        // start of the string (i32.store (i32.const 0) (i32.const 8))
	e.emitStore(b, int32(len(string(data)))) // length of the string (i32.store (i32.const 4) (i32.const 12))
	startData := 0
	for i := range data {
		index := i + 1
		if index%4 == 0 {
			startData = index - 4
			remainingData := data[startData:index]
			e.emitStore(b, int32(binary.LittleEndian.Uint32(remainingData)))
		}
	}

	return 0
}

func (e *Emitter) EmitMemory() {
	e.MemorySection.Write([]byte{0x00, 0x01}) // flags, initial (1 page 64KB)
}

func (e *Emitter) EmitRuntime() {
	// type, global mutability, value
	e.GlobalsSection.Write([]byte{op.I32, 0x01, op.I32_CONST, 0x00, op.END}) // heap-next
	e.globalsCount = 1
}
func (e *Emitter) BuildLookup() {
	for _, t := range e.Module.Types {
		e.types[t.Name] = &TypeData{
			Name:  t.Name,
			Value: "i32",
		}
	}
	for i, fun := range e.Module.Functions {
		_, exists := e.funcs[fun.Name]
		if exists {
			panic(fmt.Sprintf("Function '%s' is already defined", fun.Name))
		}
		e.funcs[fun.Name] = &FuncData{
			Index:  i,
			Name:   fun.Name,
			Params: map[string]*FuncParam{},
		}
		if fun.ReturnType != nil {
			e.funcs[fun.Name].ReturnType = fun.ReturnType.Name
		}
		for pi, param := range fun.Parameters {
			if param.ValueField != nil {
				e.funcs[fun.Name].Params[param.ValueField.Name] = &FuncParam{
					Index: pi,
					Name:  param.ValueField.Name,
					Type:  param.ValueField.Type,
				}
			}
			if param.FuncField != nil {
				e.funcs[fun.Name].Params[param.FuncField.Name] = &FuncParam{
					Index: pi,
					Name:  param.FuncField.Name,
					// Type:  param.FuncField.Parameters,
				}
			}
		}
	}
}

// React.createElement(View, { name: "test" },
//     React.createElement(View, null,
//         React.createElement(Text, null, "Hello")),
//     React.createElement(View, null,
//         React.createElement(Text, null, "Hello")));

func (e *Emitter) EmitJS() (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer(nil)
	e.BuildLookup()
	for _, fun := range e.Module.Functions {
		params := []string{}
		for _, param := range fun.Parameters {
			params = append(params, param.toJS())
		}
		buffer.WriteString(fmt.Sprintf("function %s(%s) {\n", fun.Name, strings.Join(params, ", ")))
		for _, st := range fun.Statements {
			if st.Assignment != nil {
				buffer.WriteString(fmt.Sprintf("  const %s = %s;\n", st.Assignment.Name, st.Assignment.Expression.toJS()))
			}
			if st.EchoStatement != nil {
				buffer.WriteString(fmt.Sprintf("  console.log(\"%s\");\n", st.EchoStatement.Expression.toJS()))
			}
			if st.ReturnStatement != nil {
				buffer.WriteString(fmt.Sprintf("  return %s;\n", st.ReturnStatement.Expression.toJS()))
			}
		}
		buffer.WriteString("}\n\n")
	}
	return buffer, nil
}

func (e *Emitter) EmitWASM() (*bytes.Buffer, error) {
	// buffer := bytes.NewBuffer(nil)
	// buffer.Write([]byte{0x00, 0x61, 0x73, 0x6d}) // WASM_BINARY_MAGIC
	// buffer.Write([]byte{0x01, 0x00, 0x00, 0x00}) // WASM_BINARY_VERSION
	// e.EmitMemory()
	// e.EmitRuntime()
	e.BuildLookup()

	// for _, t := range e.Module.Types {
	// 	e.types[t.Name] = &TypeData{
	// 		Name:  t.Name,
	// 		Value: "i32",
	// 	}
	// }
	// for i, fun := range e.Module.Functions {
	// 	err := e.EmitTypes(fun.Name)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("Failed to emitTypes %v", err)
	// 	}
	// 	// if fun.Type == "extern" {
	// 	// 	e.EmitImports(module.Name, fun.Name, e.externFuncsCount)
	// 	// 	e.externFuncsCount += 1
	// 	// }
	// 	e.EmitFuncs(fun.Name, e.externFuncsCount+e.funcsCount) // function 0 signature index

	// 	err = e.EmitFuncBody(fun.Name, fun.Body)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("Failed to EmitFuncBody %v", err)
	// 	}
	// 	e.funcsCount += 1
	// }
	// repr.Println(e.types)
	// repr.Println(e.funcs)
	// buffer.Write([]byte{op.SECTION_TYPES, byte(e.TypesSection.Len() + 1), byte(e.externFuncsCount + e.funcsCount)}) // Type section code, section size, num types, type data
	// buffer.Write(e.TypesSection.Bytes())
	// buffer.Write([]byte{op.SECTION_IMPORTS, byte(e.ImportsSection.Len() + 1), byte(e.externFuncsCount)}) // Imports section, section size, num imports, type data
	// buffer.Write(e.ImportsSection.Bytes())
	// buffer.Write([]byte{op.SECTION_FUNCS, byte(e.FuncsSection.Len() + 1), byte(e.funcsCount)}) // Func Sig section code, section size, num types, type data
	// buffer.Write(e.FuncsSection.Bytes())
	// buffer.Write([]byte{op.SECTION_MEMORY, byte(e.MemorySection.Len() + 1), 0x01}) // Memory section code, section size, num memories
	// buffer.Write(e.MemorySection.Bytes())
	// buffer.Write([]byte{op.SECTION_GLOBALS, byte(e.GlobalsSection.Len() + 1), byte(e.globalsCount)}) // globals section code, section size, num globals
	// buffer.Write(e.GlobalsSection.Bytes())
	// buffer.Write([]byte{op.SECTION_EXPORTS, byte(e.ExportsSection.Len() + 1), byte(e.funcsCount)}) // exports section code, section size, num exports
	// buffer.Write(e.ExportsSection.Bytes())
	// buffer.WriteByte(op.SECTION_FUNCS_BODY)                  // funcbody section code
	// encodeSleb128(buffer, int32(e.FuncsBodySection.Len()+1)) // section size
	// buffer.WriteByte(byte(e.funcsCount))                     // num functions
	// buffer.Write(e.FuncsBodySection.Bytes())
	// return buffer, nil
	return nil, nil
}

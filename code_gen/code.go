package code_gen

import (
	"bytes"
	"io"
	"yum/ast"
)

var typeMap = map[string]byte{
	"int":   0x7f,
	"float": 0x7c,
}

var operandMap = map[string]map[string]byte{
	"int": map[string]byte{
		"+": 0x6a,
		"-": 0x6b,
		"*": 0x6c,
		"/": 0x6d,
	},
	"float": map[string]byte{
		"add": 0xa0,
	},
}

const FUNC_END = 0x0b
const LOCAL_GET = 0x20     // local.get
const DECLARE_CONST = 0x41 // delcare const

type FuncParam struct {
	Index int
	Ref   string
	Type  string
}

func operate(buffer *bytes.Buffer, operation string, l *ast.Literal, r *ast.Expression) {
	if l != nil && l.Int != nil {
		buffer.WriteByte(DECLARE_CONST)
		buffer.WriteByte(byte(*l.Int))
	}
	if r.Left.Int != nil {
		buffer.WriteByte(DECLARE_CONST)
		buffer.WriteByte(byte(*r.Left.Int))
		buffer.WriteByte(operandMap["int"][operation])
	}
	if r.Operator != nil {
		operate(buffer, *r.Operator, nil, r.Right)
	}
	// if s.Left.Reference != nil {
	// 	ref, ok := funcParameterMap[*s.Left.Reference]
	// 	if !ok {
	// 		panic("Func " + fun.Name + " parameter " + *s.Left.Reference + " doesn't exist")
	// 	}
	// 	funcbody = append(funcbody, LOCAL_GET, byte(ref.Index))
	// }
	// if s.Right.Left.Reference != nil {
	// 	ref, ok := funcParameterMap[*s.Right.Left.Reference]
	// 	if !ok {
	// 		panic("Func " + fun.Name + " parameter " + *s.Right.Left.Reference + " doesn't exist")
	// 	}
	// 	funcbody = append(funcbody, LOCAL_GET, byte(ref.Index))
	// 	funcbody = append(funcbody, operandMap[ref.Type]["add"])
	// }
}

func GenerateCode(w io.Writer, ast *ast.Ast) {
	w.Write([]byte{0x00, 0x61, 0x73, 0x6d}) // WASM_BINARY_MAGIC
	w.Write([]byte{0x01, 0x00, 0x00, 0x00}) // WASM_BINARY_VERSION

	types := []byte{} // num types, func, num params
	funcs := []byte{}
	exports := []byte{}
	funcBodys := []byte{}
	for _, p := range ast.Modules {
		for i, fun := range p.Funs {
			funcParameterMap := map[string]FuncParam{}
			types = append(types, 0x60, byte(len(fun.Parameters))) // func, num params
			for i, v := range fun.Parameters {                     // i32, i32
				types = append(types, typeMap[v.Type.Name])
				funcParameterMap[v.ID] = FuncParam{
					Index: i,
					Ref:   v.ID,
					Type:  v.Type.Name,
				}
			}
			types = append(types, 0x01)                        // num results
			types = append(types, typeMap[fun.ReturnTypes[0]]) // i32

			funcs = append(funcs, byte(i))                 // function 0 signature index
			exports = append(exports, byte(len(fun.Name))) // name length
			exports = append(exports, []byte(fun.Name)...) // name
			exports = append(exports, 0x00, byte(i))       // export kind, export func index

			funcbody := bytes.NewBuffer([]byte{
				0x00, // local decl count
			})

			for _, s := range fun.Body {
				if op := s.Exp.Operator; op != nil {
					if *op == "+" {
						operate(funcbody, *op, s.Exp.Left, s.Exp.Right)
					}
				}
			}
			funcbody.WriteByte(FUNC_END)
			funcBodys = append(funcBodys, byte(len(funcbody.Bytes()))) // func body size
			funcBodys = append(funcBodys, funcbody.Bytes()...)
			// panic(fmt.Sprintf("%+v", funcbody.Bytes()))
		}
		w.Write([]byte{0x01, byte(len(types) + 1), byte(len(p.Funs))}) // Type section code, section size, num types, type data
		w.Write(types)
		w.Write([]byte{0x03, byte(len(funcs) + 1), byte(len(p.Funs))}) // Func sig section code, section size, num types, type data
		w.Write(funcs)
		w.Write([]byte{0x07, byte(len(exports) + 1), byte(len(p.Funs))}) // exports section code, section size, num exports
		w.Write(exports)
		w.Write([]byte{0x0a, byte(len(funcBodys) + 1), byte(len(p.Funs))}) // func body section code, section size, num functions
		w.Write(funcBodys)
	}
}

// ;; A function that takes an `i64` and returns
// ;; three `i32`s.
// (func (param i64) (result i32 i32 i32)
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

package compiler

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

type Ast struct {
	Pos     lexer.Position
	Modules []*Module `{ @@ }`
}

type Module struct {
	Pos       lexer.Position
	Name      string      `"module" @Ident`
	End       *string     `@NewLine`
	Imports   []*Import   `{ @@ }`
	Types     []*Type     `{ @@ }`
	Functions []*Function `{ @@ }`
}

type Import struct {
	Pos   lexer.Position
	Start *string `@NewLine`
	Url   string  `"import" @String`
	End   *string `@NewLine`
}

// type Fun struct {
// 	Pos       lexer.Position
// 	Decorator *DecoratorDecl `{ @@ }`
// 	Fun       *FunDecl       `{ @@ }`
// }

// type DecoratorDecl struct {
// 	Pos    lexer.Position
// 	Name   string     `"@" @Ident`
// 	Values []*Literal `"(" [ @@ { "," @@ } ] ")"`
// }

type Type struct {
	Pos   lexer.Position
	Start string `@NewLine`
	Name  string `"type" @Ident`
	Alias string `@Ident`
	End   string `@NewLine`
}

type Function struct {
	Pos        lexer.Position
	Start      string          `@NewLine`
	Type       string          `@("extern" | "proc" | "method")`
	Name       string          `@Ident`
	Parameters []FuncParameter `"(" [ @@ { "," @@ } ] ")"`
	ReturnType string          `"-"">" { @Ident }`
	EndReturn  string          `@NewLine`
	Body       []*Block        `{ @@ }`
	End        string          `@NewLine`
}

type FuncParameter struct {
	Pos  lexer.Position
	Name string            `@Ident`
	Type FuncParameterType `@@`
}

type FuncParameterType struct {
	Pos  lexer.Position
	Name string `":" @Ident`
}

type Block struct {
	Exp *Expression `@@`
	End *string     `@NewLine`
}

type Statement struct {
	Pos             lexer.Position
	FuncCall        *FuncCallStatement   `@@`
	Assignment      *AssignmentStatement `| @@`
	If              *IfStatement         `| @@`
	For             *ForStatement        `| @@`
	AssertStatement *AssertStatement     `| @@`
	ReturnStatement *Expression          `| @@`
	XmlDecl         *XmlDecl             `| @@`
}

type XmlDecl struct {
	Pos        lexer.Position
	Name       string         `"<"@Ident`
	Parameters []*KeyValue    `[ @@ { @@ } ]`
	EndName    lexer.Position `">"`
	Children   []*XmlDecl     `{ @@ }`
	Value      string         `{ @Ident }` // Todo make this match with @String or Literal
	Close      string         `"<""/"@Ident">"`
}

type KeyValue struct {
	Pos   lexer.Position
	Key   string   `@Ident`
	Value *Literal `"=" @@`
}

type AssignmentStatement struct {
	Pos   lexer.Position
	Name  string             `@Ident`
	Value *AssignmentLiteral `":""=" @@`
}

type FuncCallStatement struct {
	Pos    lexer.Position
	Name   string     `@Ident`
	Values []*Literal `"(" [ @@ { "," @@ } ] ")"`
}

type IfStatement struct {
	Pos       lexer.Position
	Condition string         `"if" @Ident`
	End       lexer.Position `"end"`
}

type ForStatement struct {
	Pos         lexer.Position
	Infinite    *ForInfiniteStatement    `@@`
	Conditional *ForConditionalStatement `| @@`
	Iterator    *ForIteratorStatement    `| @@`
}

type ForInfiniteStatement struct {
	Pos lexer.Position `"for"`
	End lexer.Position `"end"`
}

type ForConditionalStatement struct {
	Pos        lexer.Position `"for"`
	Identifier string         `@Ident`
	End        lexer.Position `"end"`
}

type ForIteratorStatement struct {
	Pos      lexer.Position `"for"`
	Variable string         `@Ident`
	// Identifier string         `"in" @Ident`
	End lexer.Position `"end"`
}

type Expression struct {
	Left     *Literal    `@@`
	Operator *string     `{ @("+" | "-" | "*" | "/" | "<=" | ">=" | "=""=" | "<" | ">" | "!=" | ":""=" | "=") }`
	Right    *Expression `{ @@ }`
}

type ReturnStatement struct {
	Pos        lexer.Position `"return"`
	Expression *Expression    `{ @@ }`
}

type AssertStatement struct {
	Pos        lexer.Position `"assert"`
	Expression *Expression    `@@`
}

type MapStatement struct {
	Pos   lexer.Position
	Value *MapLiteral `"{" { @@ [ "," ] } "}"`
}

type MapLiteral struct {
	Pos   lexer.Position
	Key   *string   `@String ":"`
	Value *MapValue `@@`
}

type MapValue struct {
	Pos     lexer.Position
	Literal *AssignmentLiteral `@@`
	Value   *MapStatement      `| @@`
}

type AssignmentLiteral struct {
	Pos     lexer.Position
	Call    *FuncCallStatement `@@`
	Literal *Literal           `| @@`
	Map     *MapStatement      `| @@`
}

// Literal is a "union" type, where only one matching value will be present.
type Literal struct {
	Pos       lexer.Position
	Nil       bool        `@"nil"`
	Str       *string     `| @String`
	Int       *int        `| @Int`
	Float     *float32    `| @Float`
	Bool      *string     `| @( "true" | "false" )`
	Reference *string     `| @Ident { @"." @Ident }`
	List      []*Literal  `| "[" { @@ [ "," ] } "]"`
	Params    []*Literal  `| "(" [ @@ { "," @@ } ] ")"`
	Sub       *Expression `| "(" @@ ")"`
	// Map       []*MapItem `| "{" { @@ [ "," ] } "}"`
}

// func (l *Literal) GoString() string {
// 	switch {
// 	case l.Str != nil:
// 		return fmt.Sprintf("%q", *l.Str)
// 	case l.Float != nil:
// 		return fmt.Sprintf("%v", *l.Float)
// 	case l.Int != nil:
// 		return fmt.Sprintf("%v", *l.Int)
// 	case l.Bool != nil:
// 		return fmt.Sprintf("%v", *l.Bool)
// 	case l.Reference != nil:
// 		return fmt.Sprintf("%s", *l.Reference)
// 	case l.Minus != nil:
// 		return fmt.Sprintf("-%v", l.Minus)
// 	case l.List != nil:
// 		parts := []string{}
// 		for _, e := range l.List {
// 			parts = append(parts, e.GoString())
// 		}
// 		return fmt.Sprintf("[%s]", strings.Join(parts, ", "))
// 		// case l.Map != nil:
// 		// 	parts := map[string]string{}
// 		// 	for _, e := range l.Map {
// 		// 		parts[e.Key.GoString()] = e.Value.GoString()
// 		// 	}
// 		// 	return fmt.Sprintf("%#v", parts)
// 	}
// 	panic("unsupported?")
// }

// type MapItem struct {
// 	Pos   lexer.Position
// 	Key   *Literal `@@ ":"`
// 	Value *Literal `@@`
// }

// func (m *MapItem) GoString() string {
// 	return fmt.Sprintf("%v: %v", m.Key, m.Value)
// }

// type CommentStatement struct {
// 	Pos  lexer.Position
// 	Text string `@Comment`
// }

var parser = participle.MustBuild(&Ast{},
	participle.Lexer(DefaultDefinition),
)

func ParseFile(filename string) (*Ast, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	tree := &Ast{}
	err = parser.ParseBytes(data, tree)
	if err != nil {
		return nil, err
	}
	for _, p := range tree.Modules {
		for _, imp := range p.Imports {
			parts := strings.Split(imp.Url, "/")
			moduleName := parts[len(parts)-1]
			println(moduleName)
			subTree, err := ParseFile("../examples/" + moduleName + ".yum")
			if err != nil {
				return nil, fmt.Errorf("failed to parse module: %s -> %v", moduleName, err)
			}
			tree.Modules = append(subTree.Modules, tree.Modules...)
		}
	}
	return tree, nil
}

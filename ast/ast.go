package ast

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

var parser = participle.MustBuild(&Ast{})

func ParseFile(filename string) (*Ast, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	ast := &Ast{}
	err = parser.ParseBytes(data, ast)
	if err != nil {
		return nil, err
	}
	return ast, nil
}

type Ast struct {
	Pos      lexer.Position
	Packages []*Package `{ @@ }`
}

type Package struct {
	Pos     lexer.Position
	Name    *string            `"package" @Ident`
	Imports []*ImportStatement `{ @@ }`
	Structs []*Struct          `{ @@ }`
	Funs    []*FunDecl         `{ @@ }`
	// Enums    []*Enum            `| @@`
	Typedefs []*Typedef `{ @@ }`
	// Consts     []*ConstAssignment `| @@`
}

type ImportStatement struct {
	Pos  lexer.Position
	Name string `"import" @Ident`
}

type Type struct {
	Pos  lexer.Position
	Name string `":" @Ident`
}

type KeyValue struct {
	Pos   lexer.Position
	Key   string   `@Ident`
	Value *Literal `"=" @@`
}

type ReturnType struct {
	Pos  lexer.Position
	Name string `"-"">" @Ident`
}

type ClassField struct {
	Pos  lexer.Position
	Name string `@Ident`
	Type *Type  `@@`
	// Default     *Literal      `[ "=" @@ ]`
	// Annotations []*Annotation `[ "(" @@ { "," @@ } ")" ] [ ";" ]`
}

type Struct struct {
	Pos     lexer.Position
	Name    string         `"class" @Ident`
	Traits  []*ClassTrait  `{ "(" [ @@ { "," @@ } ] ")" }`
	Fields  []*ClassField  `{ @@ }`
	Methods []*Fun         `{ @@ }`
	End     lexer.Position `"end"`
}

type ClassTrait struct {
	Pos lexer.Position
	ID  string `@Ident`
}

type Enum struct {
	Pos   lexer.Position
	Name  string   `"enum" @Ident "{"w`
	Cases []string `{ @Ident } "}"`
}

type Typedef struct {
	Pos  lexer.Position
	Type *Type  `"typedef" @@`
	Name string `@Ident`
}

type MethodParameter struct {
	Pos  lexer.Position
	ID   string `@Ident`
	Type *Type  `@@`
}

type MethodCallParameter struct {
	Pos  lexer.Position
	ID   string `@Ident`
	Type *Type  `@@`
}

type DecoratorDecl struct {
	Pos    lexer.Position
	Name   string     `"@" @Ident`
	Values []*Literal `"(" [ @@ { "," @@ } ] ")"`
}

type FunDecl struct {
	Pos        lexer.Position
	Name       string             `"func" @Ident`
	Parameters []*MethodParameter `"(" [ @@ { "," @@ } ] ")"`
	ReturnType *ReturnType        `{ @@ }`
	Body       []*Statement       `{ @@ }`
	End        lexer.Position     `"end"`
}

type Fun struct {
	Pos       lexer.Position
	Decorator *DecoratorDecl `{ @@ }`
	Fun       *FunDecl       `{ @@ }`
}

type Block struct {
	Pos        lexer.Position
	Statements []*Statement `{ @@ }`
}

type Statement struct {
	Pos             lexer.Position
	FuncCall        *FuncCallStatement   `@@`
	Assignment      *AssignmentStatement `| @@`
	If              *IfStatement         `| @@`
	For             *ForStatement        `| @@`
	AssertStatement *AssertStatement     `| @@`
	ReturnStatement *ReturnStatement     `| @@`
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

type ReturnStatement struct {
	Pos           lexer.Position `"return"`
	ReturnLiteral *ReturnLiteral `@@`
}

type ReturnLiteral struct {
	Pos     lexer.Position
	Literal *Literal `@@`
	XmlDecl *XmlDecl `| @@`
}

type AssertStatement struct {
	Pos      lexer.Position     `"assert"`
	LHS      *AssignmentLiteral `@@`
	Operator *Operator          `@@`
	RHS      *AssignmentLiteral `@@`
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

type Operator struct {
	Pos   lexer.Position
	Value string `@( "<=" | ">=" | "=""=" | "<" | ">" | "!=" )`
}

// Literal is a "union" type, where only one matching value will be present.
type Literal struct {
	Pos       lexer.Position
	Nil       bool       `@"nil"`
	Str       *string    `| @String`
	Float     *float64   `| @Float`
	Int       *int64     `| @Int`
	Bool      *string    `| @( "true" | "false" )`
	Reference *string    `| @Ident { @"." @Ident }`
	Minus     *Literal   `| "-" @@`
	List      []*Literal `| "[" { @@ [ "," ] } "]"`
	// Map       []*MapItem `| "{" { @@ [ "," ] } "}"`
}

func (l *Literal) GoString() string {
	switch {
	case l.Str != nil:
		return fmt.Sprintf("%q", *l.Str)
	case l.Float != nil:
		return fmt.Sprintf("%v", *l.Float)
	case l.Int != nil:
		return fmt.Sprintf("%v", *l.Int)
	case l.Bool != nil:
		return fmt.Sprintf("%v", *l.Bool)
	case l.Reference != nil:
		return fmt.Sprintf("%s", *l.Reference)
	case l.Minus != nil:
		return fmt.Sprintf("-%v", l.Minus)
	case l.List != nil:
		parts := []string{}
		for _, e := range l.List {
			parts = append(parts, e.GoString())
		}
		return fmt.Sprintf("[%s]", strings.Join(parts, ", "))
		// case l.Map != nil:
		// 	parts := map[string]string{}
		// 	for _, e := range l.Map {
		// 		parts[e.Key.GoString()] = e.Value.GoString()
		// 	}
		// 	return fmt.Sprintf("%#v", parts)
	}
	panic("unsupported?")
}

// type MapItem struct {
// 	Pos   lexer.Position
// 	Key   *Literal `@@ ":"`
// 	Value *Literal `@@`
// }

// func (m *MapItem) GoString() string {
// 	return fmt.Sprintf("%v: %v", m.Key, m.Value)
// }

func CheckAst(ast *Ast) []error {
	return nil
}

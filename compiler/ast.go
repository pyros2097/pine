//golint:ignore

package compiler

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

type Module struct {
	Pos       lexer.Position
	Name      string      `parser:"\"module\" @Ident"`
	Imports   []*Import   `parser:"{ @@ }"`
	Enums     []*Enum     `parser:"{ @@ }"`
	Types     []*Type     `parser:"{ @@ }"`
	Externs   []*Extern   `parser:"{ @@ }"`
	Functions []*Function `parser:"{ @@ }"`
	Tests     []*Test     `parser:"{ @@ }"`
}

type Import struct {
	Pos lexer.Position
	URL string `parser:"\"import\" @String"`
}

type Enum struct {
	Pos   lexer.Position
	Name  string   `parser:"\"enum\" @Ident \"{\""`
	Value []string `parser:"[ @Ident { @Ident } ] \"}\""`
}

type Type struct {
	Pos    lexer.Position
	Name   string      `parser:"\"type\" @Ident \"{\""`
	Fields []TypeField `parser:"[ @@ { @@ } ] \"}\""`
}

type TypeField struct {
	Pos   lexer.Position
	Name  string `parser:"@Ident"`
	Value string `parser:"@Ident"`
}

type ReturnType struct {
	Pos  lexer.Position
	Name string `parser:"@Ident"`
}

type Extern struct {
	Pos        lexer.Position
	Name       string      `parser:"\"extern\" @Ident"`
	Parameters []TypeField `parser:"\"(\" [ @@ { \",\"  @@ } ] \")\""`
	ReturnType *ReturnType `parser:"{ @@ }"`
}

type Function struct {
	Pos        lexer.Position
	Type       string       `parser:"@(\"proc\" | \"method\")"`
	Name       string       `parser:"@Ident"`
	Parameters []TypeField  `parser:"\"(\" [ @@ { \",\"  @@ } ] \")\""`
	ReturnType *ReturnType  `parser:"{ @@ }"`
	Statements []*Statement `parser:"\"{\" { @@ } \"}\""`
}

type Test struct {
	Pos        lexer.Position
	Name       string       `parser:"\"test\" @String"`
	Statements []*Statement `parser:"\"{\" { @@ } \"}\""`
}

type Statement struct {
	Pos             lexer.Position
	FuncCall        *FuncCallStatement   `parser:"@@"`
	Assignment      *AssignmentStatement `parser:"| @@"`
	If              *IfStatement         `parser:"| @@"`
	For             *ForStatement        `parser:"| @@"`
	EchoStatement   *EchoStatement       `parser:"| @@"`
	AssertStatement *AssertStatement     `parser:"| @@"`
	ReturnStatement *ReturnStatement     `parser:"| @@"`
}

type KeyValue struct {
	Pos   lexer.Position
	Key   string   `parser:"@Ident"`
	Value *Literal `parser:"\"=\" @@"`
}

type AssignmentStatement struct {
	Pos        lexer.Position
	Name       string      `parser:"@Ident"`
	Expression *Expression `parser:"\":\"\"=\" @@"`
}

type FuncCallStatement struct {
	Pos    lexer.Position
	Name   string     `parser:"@Ident"`
	Values []*Literal `parser:"\"(\" [ @@ { \",\" @@ } ] \")\""`
}

type IfStatement struct {
	Pos       lexer.Position
	Condition []*Expression `parser:"\"if\" @@"`
	Result    []*Statement  `parser:"\"{\" { @@ } \"}\""`
	Alternate []*Statement  `parser:"| \"else\" \"{\" { @@ } \"}\""`
}

type ForStatement struct {
	Pos         lexer.Position
	Infinite    *ForInfiniteStatement    `parser:"@@"`
	Conditional *ForConditionalStatement `parser:"| @@"`
	Iterator    *ForIteratorStatement    `parser:"| @@"`
}

type ForInfiniteStatement struct {
	Pos        lexer.Position `parser:"\"for\""`
	Statements []*Statement   `parser:"\"{\" { @@ } \"}\""`
}

type ForConditionalStatement struct {
	Pos        lexer.Position `parser:"\"for\""`
	Condition  []*Expression  `parser:"@@"`
	Statements []*Statement   `parser:"\"{\" { @@ } \"}\""`
}

type ForIteratorStatement struct {
	Pos        lexer.Position `parser:"\"for\""`
	Item       string         `parser:"@Ident"`
	Reference  *Literal       `parser:"\"in\" @@"`
	Statements []*Statement   `parser:"\"{\" { @@ } \"}\""`
}

type Expression struct {
	Left     *Literal    `parser:"@@"`
	Operator *string     `parser:"{ @(\"+\" | \"-\" | \"*\" | \"/\" | \"<=\" | \">=\" | \"=\"\"=\" | \"<\" | \">\" | \":\"\"=\" |\"!\"\"=\" | \"=\") }"`
	Right    *Expression `parser:"{ @@ }"`
}

func (e *Expression) GoString() string {
	s := ""
	if e.Left != nil {
		s += e.Left.GoString()
	}
	if e.Operator != nil {
		s += " " + *e.Operator
	}
	if e.Right != nil {
		s += " " + e.Right.GoString()
	}
	return s
}

type ReturnStatement struct {
	Pos        lexer.Position `parser:"\"return\""`
	Expression *Expression    `parser:"{ @@ }"`
}

type EchoStatement struct {
	Pos        lexer.Position `parser:"\"echo\""`
	Expression *Expression    `parser:"@@"`
}

type AssertStatement struct {
	Pos        lexer.Position `parser:"\"@\"\"assert\""`
	Expression *Expression    `parser:"@@"`
}

type MapLiteral struct {
	Pos   lexer.Position
	Key   *string   `parser:"@String \":\""`
	Value *MapValue `parser:"@@"`
}

type MapValue struct {
	Pos     lexer.Position
	Literal *Literal    `parser:"@@"`
	Value   *MapLiteral `parser:"| @@"`
}

type XmlLiteral struct {
	Pos        lexer.Position
	Name       string         `parser:"\"<\"@Ident"`
	Parameters []*KeyValue    `parser:"[ @@ { @@ } ]"`
	EndName    lexer.Position `parser:"\">\""`
	Children   []*XmlLiteral  `parser:"{ @@ }"`
	Value      string         `parser:"{ @Ident }"` // Todo make this match with @String or Literal
	Close      string         `parser:"\"<\"\"/\"@Ident\">\""`
}

func (x *XmlLiteral) toJS(space string) string {
	s := space + "React.createElement(\n" + space + space + x.Name + ",\n" + space + space + "{"
	for _, param := range x.Parameters {
		s += `"` + param.Key + `"` + "=" + `"` + param.Value.GoString() + `", `
	}
	s += "},\n"
	for _, c := range x.Children {
		s += c.toJS(space+space) + ", "
	}
	s += space + ")\n"
	return s
}

// Literal is a "union" type, where only one matching value will be present.
type Literal struct {
	Pos       lexer.Position
	Nil       bool               `parser:"@\"nil\""`
	Str       *string            `parser:"| @String"`
	Int       *int               `parser:"| @Int"`
	Float     *float32           `parser:"| @Float"`
	Bool      *string            `parser:"| @( \"true\" | \"false\" )"`
	Reference *string            `parser:"| @Ident { @\".\" @Ident }"`
	FuncCall  *FuncCallStatement `parser:"| @@"`
	List      []*Literal         `parser:"| \"[\" { @@ [ \",\" ] } \"]\""`
	Map       []*MapLiteral      `| "{" { @@ [ "," ] } "}"`
	Params    []*Literal         `parser:"| \"(\" [ @@ { \",\" @@ } ] \")\""`
	Xml       *XmlLiteral        `parser:"| @@ "`
	// Sub       *Expression `parser:"| \"(\" @@ \")\""`
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
	// case l.Minus != nil:
	// 	return fmt.Sprintf("-%v", l.Minus)
	case l.List != nil:
		parts := []string{}
		for _, e := range l.List {
			parts = append(parts, e.GoString())
		}
		return fmt.Sprintf("[%s]", strings.Join(parts, ", "))
	case l.Map != nil:
		return ""
		// parts := map[string]string{}
		// for _, e := range l.Map {
		// 	parts[e.Value.Key.GoString()] = e.Value.Value.GoString()
		// }
		// return fmt.Sprintf("%#v", parts)
	case l.Xml != nil:
		return l.Xml.toJS("  ")
	}
	panic("unsupported? literal")
}

// type MapItem struct {
// 	Pos   lexer.Position
// 	Key   *Literal `@@ ":"`
// 	Value *Literal `@@`
// }

// func (m *MapItem) GoString() string {
// 	return fmt.Sprintf("%v: %v", m.Key, m.Value)
// }

var parser = participle.MustBuild(&Module{})

func ParseFile(filename string) (*Module, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	tree := &Module{}
	err = parser.ParseBytes(data, tree)
	if err != nil {
		return nil, err
	}
	// for _, imp := range tree.ImportSection.Imports {
	// 	parts := strings.Split(imp.Url, "/")
	// 	moduleName := parts[len(parts)-1]
	// 	println(moduleName)
	// 	// subTree, err := ParseFile("../examples/" + moduleName + ".yum")
	// 	// if err != nil {
	// 	// 	return nil, fmt.Errorf("failed to parse module: %s -> %v", moduleName, err)
	// 	// }
	// 	// tree.Modules = append(subTree.Modules, tree.Modules...)
	// }
	return tree, nil
}

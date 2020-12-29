package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/repr"
)

type Import struct {
	Name string `"(""use" @Ident ")"`
}

type Enum struct {
	Key    string   `"(""enum" @Ident`
	Values []string `[ ":"@Ident { ":"@Ident } ] ")"`
}

type Def struct {
	Key   string   `"(""def" @Ident`
	Value *Literal `@@ ")"`
}

type Defn struct {
	Name   string   `"(""defn" @Ident @{ "/" Ident }`
	Params []string `"[" [ @Ident { @Ident } ] "]"`
	Body   []Expr   `{ @@ } ")"`
}

type Expr struct {
	Op  *string  `"(" @(Ident { "/" Ident} | "+" | "-" | "*" | "/" | "<=" | ">=" | "=""=" | "<" | ">" | ":""=" |"!""=" | "=")`
	LHS *Literal `@@`
	RHS *Literal `{ @@ } ")"`
}

func (e *Expr) JS(indent string) string {
	switch *e.Op {
	case "+", "-", "*", "/", "^", "%", ">", "<", ">=", "<=", "==", "!=":
		return "(" + e.LHS.JS("") + " " + *e.Op + " " + e.RHS.JS("") + ")"
	case "if":
		return indent + "if " + e.LHS.JS("") + " {\n" + e.RHS.JS(indent+"  ") + "\n" + indent + "}"
	case "var":
		return indent + "var " + e.LHS.JS("") + " = " + e.RHS.JS("")
	default:
		// assume func call
		// all functions support named parameters by default
		return indent + strings.Replace(*e.Op, "/", "_", -1) + "(" + e.LHS.JS(indent) + "" + e.RHS.JS(indent) + ")"
	}
}

type Field struct {
	Name string `":" @Ident`
	Type string "@Ident"
	Tags []*Tag `[ @@ { @@ } ]`
}

type Tag struct {
	Name   string     `@Ident "("`
	Params []TagValue `[ @@ { @@ } ] ")"`
}

type TagValue struct {
	String *string  `@String`
	Number *float64 `| @(Float | Int)`
}

type Struct struct {
	Name   string   `"(""struct" @Ident`
	Fields []*Field `[ @@ { @@ } ] ")"`
}

type Literal struct {
	Nil       bool     `@"nil"`
	Reference *string  `| @Ident`
	String    *string  `| @String`
	Number    *float64 `| @(Float | Int)`
	Expr      *Expr    `| @@`
}

func (l *Literal) JS(indent string) string {
	if l != nil {
		if l.Nil {
			return "null"
		} else if l.Reference != nil {
			return *l.Reference
		} else if l.String != nil {
			return *l.String
		} else if l.Number != nil {
			return fmt.Sprintf("%f", *l.Number)
		} else {
			return indent + l.Expr.JS(indent+"  ")
		}
	}
	return ""
}

type Node struct {
	Import *Import `@@`
	Def    *Def    `| @@ `
	Defn   *Defn   `| @@ `
	Struct *Struct `| @@ `
	Enum   *Enum   `| @@`
}

type Module struct {
	Pos lexer.Position
	Ast []Node `{ @@ }`
}

func main() {
	parser := participle.MustBuild(&Module{})
	tree := &Module{}
	err := parser.ParseBytes([]byte(`
	(use jaylib)
	(use std)
	(use lala)

	(enum direction
	  :left
	  :right)

	(struct user
	  :name str min(10) max(20) range(10 120)
	  :age  float min(0) max(150))

	(defn min [s v msg]
	  (if (> (string/length s) v)
	  	(raise msg)))

	(def user1 "hello")
	(def user2 123.12)
	(def user3 144)
	(defn show/dir [dir]
		(var index (usestate 10))
		(html/div dir))
	`), tree)
	if err != nil {
		panic(err)
	}
	js := bytes.NewBuffer(nil)
	for _, a := range tree.Ast {
		if i := a.Import; i != nil {
			js.WriteString("import './" + i.Name + ".js'\n")
		}
		if a.Def != nil {
		}
		if d := a.Defn; d != nil {
			params := ""
			for _, p := range d.Params {
				params += p + ", "
			}
			js.WriteString("\n")
			js.WriteString("const " + d.Name + " = (" + params + ") => {\n")
			for i, b := range d.Body {
				if i == len(d.Body)-1 && *b.Op != "if" {
					js.WriteString("  return ")
				}
				js.WriteString(b.JS("  "))
				js.WriteString("\n")
			}
			js.WriteString("}\n")
		}
		if s := a.Struct; s != nil {
			params := ""
			for _, f := range s.Fields {
				params += f.Name + ", "
			}
			js.WriteString("\n")
			js.WriteString("const " + s.Name + "new = (" + params + ") => {\n")
			js.WriteString("  return {\n")
			for _, f := range s.Fields {
				js.WriteString("    " + f.Name + ": " + f.Name + ", \n")
			}
			js.WriteString("  }")
			js.WriteString("\n}\n")
		}
		if a.Enum != nil {
		}
	}
	println(repr.String(tree, repr.Indent("  ")))
	println(js.String())
}
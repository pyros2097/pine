package main

import (
	"bytes"

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
	  (if (+ (string/length s) v)
	  	(raise msg)))

	(def user1 "hello")
	(def user2 123.12)
	(def user3 144)
	(defn show/dir [dir]
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
			js.WriteString("\n}\n")
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

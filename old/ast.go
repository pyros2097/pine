package compiler

import (
	"fmt"
	"strings"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

var parser = participle.MustBuild(&Ast{})
	
// func (p *parser) parseFile() *ast.File {
// 	if p.trace {
// 		defer un(trace(p, "File"))
// 	}

// 	// Don't bother parsing the rest if we had errors scanning the first token.
// 	// Likely not a Go source file at all.
// 	if p.errors.Len() != 0 {
// 		return nil
// 	}

// 	// package clause
// 	doc := p.leadComment
// 	pos := p.expect(token.PACKAGE)
// 	// Go spec: The package clause is not a declaration;
// 	// the package name does not appear in any scope.
// 	ident := p.parseIdent()
// 	if ident.Name == "_" && p.mode&DeclarationErrors != 0 {
// 		p.error(p.pos, "invalid package name _")
// 	}
// 	p.expectSemi()

// 	// Don't bother parsing the rest if we had errors parsing the package clause.
// 	// Likely not a Go source file at all.
// 	if p.errors.Len() != 0 {
// 		return nil
// 	}

// 	p.openScope()
// 	p.pkgScope = p.topScope
// 	var decls []ast.Decl
// 	if p.mode&PackageClauseOnly == 0 {
// 		// import decls
// 		for p.tok == token.IMPORT {
// 			decls = append(decls, p.parseGenDecl(token.IMPORT, p.parseImportSpec))
// 		}

// 		if p.mode&ImportsOnly == 0 {
// 			// rest of package body
// 			for p.tok != token.EOF {
// 				decls = append(decls, p.parseDecl(syncDecl))
// 			}
// 		}
// 	}
// 	p.closeScope()
// 	assert(p.topScope == nil, "unbalanced scopes")
// 	assert(p.labelScope == nil, "unbalanced label scopes")

// 	// resolve global identifiers within the same file
// 	i := 0
// 	for _, ident := range p.unresolved {
// 		// i <= index for current ident
// 		assert(ident.Obj == unresolved, "object already resolved")
// 		ident.Obj = p.pkgScope.Lookup(ident.Name) // also removes unresolved sentinel
// 		if ident.Obj == nil {
// 			p.unresolved[i] = ident
// 			i++
// 		}
// 	}

// 	return &ast.File{
// 		Doc:        doc,
// 		Package:    pos,
// 		Name:       ident,
// 		Decls:      decls,
// 		Scope:      p.pkgScope,
// 		Imports:    p.imports,
// 		Unresolved: p.unresolved[0:i],
// 		Comments:   p.comments,
// 	}
// }

type Ast struct {
	Pos     lexer.Position
	Entries []*Entry `{ @@ }`
}

type Entry struct {
	Pos     lexer.Position
	Package *string            `"package" @Ident`
	Imports []*ImportStatement `| @@`
	Classes []*Class           `| @@`
	// Funcs   []*Func            `| @@`
	// Enums    []*Enum            `| @@`s
	Typedefs []*Typedef `| @@`
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

type ClassField struct {
	Pos  lexer.Position
	Name string `@Ident`
	Type *Type  `@@`
	// Default     *Literal      `[ "=" @@ ]`
	// Annotations []*Annotation `[ "(" @@ { "," @@ } ")" ] [ ";" ]`
}

type Class struct {
	Pos     lexer.Position
	Name    string         `"class" @Ident`
	Traits  []*ClassTrait  `{ "(" [ @@ { "," @@ } ] ")" }`
	Fields  []*ClassField  `{ @@ }`
	Methods []*ClassMethod `{ @@ }`
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

type ClassMethod struct {
	Pos        lexer.Position
	Name       string             `"fun" @Ident`
	Parameters []*MethodParameter `"(" [ @@ { "," @@ } ] ")"`
	ReturnType *Type              `{ @@ }`
	Body       []*Statement       `{ @@ }`
	End        lexer.Position     `"end"`
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
	Pos   lexer.Position `"return"`
	Liter *Literal       `@@`
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

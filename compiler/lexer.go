package compiler

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/alecthomas/participle/lexer"
)

// TextScannerLexer is a lexer that uses the text/scanner module.
var (
	TextScannerLexer lexer.Definition = &defaultDefinition{}

	// DefaultDefinition defines properties for the default lexer.
	DefaultDefinition = TextScannerLexer
)

type defaultDefinition struct{}

func (d *defaultDefinition) Lex(r io.Reader) (lexer.Lexer, error) {
	return Lex(r), nil
}

func (d *defaultDefinition) Symbols() map[string]rune {
	return map[string]rune{
		"EOF":       EOF,
		"Char":      Char,
		"Ident":     Ident,
		"Int":       Int,
		"Float":     Float,
		"String":    String,
		"RawString": RawString,
		"Comment":   Comment,
		"NewLine":   NewLine,
		"Indent":    Indent,
		"Dedent":    Dedent,
	}
}

// textScannerLexer is a Lexer based on text/Scanner
type textScannerLexer struct {
	scanner  *Scanner
	filename string
	err      error
	stack    int
}

// Lex an io.Reader with text/
//
// This provides very fast lexing of source code compatible with Go tokens.
//
// Note that this differs from text/Scanner in that string tokens will be unquoted.
func Lex(r io.Reader) lexer.Lexer {
	s := &Scanner{}
	s.Init(r)
	l := lexWithScanner(r, s)
	l.scanner.Error = func(s *Scanner, msg string) {
		// This is to support single quoted strings. Hacky.
		if !strings.HasSuffix(msg, "char literal") {
			l.err = lexer.Errorf(lexer.Position(l.scanner.Pos()), msg)
		}
	}
	return l
}

// LexWithScanner creates a Lexer from a user-provided
//
// Useful if you need to customise the
func LexWithScanner(r io.Reader, scan *Scanner) lexer.Lexer {
	return lexWithScanner(r, scan)
}

func lexWithScanner(r io.Reader, scan *Scanner) *textScannerLexer {
	lexer := &textScannerLexer{
		filename: lexer.NameOfReader(r),
		scanner:  scan,
	}
	return lexer
}

// LexBytes returns a new default lexer over bytes.
func LexBytes(b []byte) lexer.Lexer {
	return Lex(bytes.NewReader(b))
}

// LexString returns a new default lexer over a string.
func LexString(s string) lexer.Lexer {
	return Lex(strings.NewReader(s))
}

func (t *textScannerLexer) Next() (lexer.Token, error) {
	typ := t.scanner.Scan()
	text := t.scanner.TokenText()
	pos := lexer.Position(t.scanner.Position)
	pos.Filename = t.filename
	if t.err != nil {
		return lexer.Token{}, t.err
	}
	// skip all whitespaces in between expressions/statements
	if text == " " {
		return t.Next()
	}

	// start looking at whitespaces that appear after a new line
	if text == "\n" {
		typ = t.scanner.Scan()
		text = t.scanner.TokenText()
		pos = lexer.Position(t.scanner.Position)
		pos.Filename = t.filename
		if text == " " {
			typ = t.scanner.Scan()
			text = t.scanner.TokenText()
			pos = lexer.Position(t.scanner.Position)
			pos.Filename = t.filename
			if text == " " {
				if t.stack == 2 {
					// skip INDENT LEVEL 2
					return t.Next()
				}
				t.stack += 2
				return textScannerTransform(lexer.Token{
					Type:  Indent,
					Value: "III",
					Pos:   pos,
				})
			} else {
				return textScannerTransform(lexer.Token{
					Type:  typ,
					Value: text,
					Pos:   pos,
				})
			}
		} else {
			if t.stack == 2 {
				t.stack -= 2
				return textScannerTransform(lexer.Token{
					Type:  Dedent,
					Value: "III",
					Pos:   pos,
				})
			} else {
				// skip \n
				return t.Next()
			}
		}
	}
	return textScannerTransform(lexer.Token{
		Type:  typ,
		Value: text,
		Pos:   pos,
	})
}

func textScannerTransform(token lexer.Token) (lexer.Token, error) {
	// Unquote strings.
	switch token.Type {
	case Char:
		// FIXME(alec): This is pretty hacky...we convert a single quoted char into a double
		// quoted string in order to support single quoted strings.
		token.Value = fmt.Sprintf("\"%s\"", token.Value[1:len(token.Value)-1])
		fallthrough
	case String:
		s, err := strconv.Unquote(token.Value)
		if err != nil {
			return lexer.Token{}, lexer.Errorf(token.Pos, "%s: %q", err.Error(), token.Value)
		}
		token.Value = s
		if token.Type == Char && utf8.RuneCountInString(s) > 1 {
			token.Type = String
		}
	case RawString:
		token.Value = token.Value[1 : len(token.Value)-1]
	}
	return token, nil
}

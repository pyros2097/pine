# :yum: yum
Early draft/designs of a programming language I want to create

Frontend: PEG.js, ANTLR 4, YACC
// Maybe Yacc+ since go supports it
// Maybe implement in pony as pony has will have some what a similar syntax

Backend: LLVM, Go (lisp implemented in go)

## Imports/Packages
```js
use gui::Color
use std::Text
use std.Collections::{List, Hash}

List()
require {
  colors
  text
  crypto
  http
  log
}
```

## Package Management
Something like npm, Project Specific

## Functions/Closures
1. camelCase
2. Currying
3. Partials

## Interfaces (#Important)
Composition over direct inheritance

## Classes
1. Single Inheritance Level
2. Static Methods
3. Decorators
4. Annotations
5. Can be used Declaratively like golang
6. Unit Tests within methods
7. Contracts
8. Strong typing method signatures
9. Loose typing local variables
10. No @ symbol for this, works like java
11. No automatic returns/ Make it explicit
12. Pascal Case
```coffee
class Example
  let x: String
  let y: String
  let z: String
  let r: Number
  let k: Number|None

  constructor: (_a: Number, _b : Number) ->
    a = _a  # wrapped in getters and setters automatically this can be overridden
    b = _b
    
  calculate: (min: int) :int ->
    require min > 0 # contracts
    result = Math.random(min)
    ensure result > 0
    return result
    
Example{a: 10, b: 10} or Example(a, b)

# So we can have a declarative interface to almost all libraries like this
Window{
  x: 0,
  y: 0
  w: 400
  h: 300
  children: [
    VBox{
      children: [
        Label{
          text: 'This is my title'
        },
        Button{
          text: 'Click me'
          onClick: (event: Event) ->
            log.debug event.type
        }
      ]
    }
  ]
}

Window{x: 0, y: 0, w: 400, h: 300, children{
  Window{x: 0, y: 0, children: {
      Box{}
    }}
  }}
```

## Match
Must be easy to use
```coffee
match code
| 1 -> log('1')
| 4, 5 -> log('test)
| 7 -> {
|   logs
| }
else
  error
end

try
catch Exception
end
for cond
end
for i in range
end
```
# Standard Library

## JS
1. Cross platform GL Layer (glium) and NUI based on libGDX or kivy
2. Convert transpile classes/to js classes or create a js runtime similar to scala.js or Dart or that Intellij lang or clojure
3. Defacto JS Library Simulacra.js for rendering to web,desktop,mobile (Ex: GWT) using structs/classes mapped to html templates
4. Defacto Flexbox with css for layouting for GUI components (facebook-css implementation)

## Type System
0. None
1. Dec64 -> includes all math function no need to math package and convert to u16,u32
2. String -> includes all string functions including regex
3. Array
4. Map
5. Class -> Includes enums/FSM
6. Primitives ->
7. Type alias -> Unions
8. Attributes
9. JSON support

Maybe primitves Number, String map should be lowercase like num, string, map, list, none.

## Streams
A major part in std lib. All long operations must use streams. Need to improve error handling and sync streams. highlanderjs. Concurrent Streams.
Streams are good for functional programming and channels and async and gui single event loop.
Runtime use epoll or coroutines or MIO like event mechanism. Functional

## Generators (Maybe not needed)
Use them for default async behavior or await or make the language follow nodejs style callback support but the language encapsulates this with syntax like iced coffeescript or livescript

Reduce Capital Letters (unlike go)
Reduce using Shift Key

Grammar (Taken from pony)
```antlr
// ANTLR v3 grammar
grammar pony;

options
{
  output = AST;
  k = 1;
}

// Parser

module
  : STRING? use* class_def* 
  ;

use
  : 'use' (ID '=')? (STRING | use_ffi) ('if' infix)?
  ;

use_ffi
  : '@' (ID | STRING) typeargs ('(' | LPAREN_NEW) params? ')' '?'?
  ;

class_def
  : ('type' | 'interface' | 'trait' | 'primitive' | 'struct' | 'class' | 'actor') '@'? cap? ID typeparams? ('is' type)? STRING? members
  ;

members
  : field* method*
  ;

field
  : ('var' | 'let' | 'embed') ID ':' type ('delegate' type)? ('=' infix)?
  ;

method
  : ('fun' | 'be' | 'new') cap? ID typeparams? ('(' | LPAREN_NEW) params? ')' (':' type)? '?'? STRING? ('if' rawseq)? ('=>' rawseq)?
  ;

rawseq
  : exprseq
  | jump
  ;

exprseq
  : assignment (semiexpr | nosemi)?
  ;

nextexprseq
  : nextassignment (semiexpr | nosemi)?
  ;

nosemi
  : nextexprseq
  | jump
  ;

semiexpr
  : ';' (exprseq | jump)
  ;

jump
  : ('return' | 'break' | 'continue' | 'error' | 'compile_intrinsic' | 'compile_error') rawseq?
  ;

nextassignment
  : nextinfix ('=' assignment)?
  ;

assignment
  : infix ('=' assignment)?
  ;

nextinfix
  : nextterm antlr_0*
  ;

infix
  : term antlr_1*
  ;

binop
  : ('and' | 'or' | 'xor' | '+' | '-' | '*' | '/' | '%' | '<<' | '>>' | 'is' | 'isnt' | '==' | '!=' | '<' | '<=' | '>=' | '>') term
  ;

nextterm
  : 'if' rawseq 'then' rawseq (elseif | ('else' rawseq))? 'end'
  | 'ifdef' infix 'then' rawseq (elseifdef | ('else' rawseq))? 'end'
  | 'match' rawseq caseexpr* ('else' rawseq)? 'end'
  | 'while' rawseq 'do' rawseq ('else' rawseq)? 'end'
  | 'repeat' rawseq 'until' rawseq ('else' rawseq)? 'end'
  | 'for' idseq 'in' rawseq 'do' rawseq ('else' rawseq)? 'end'
  | 'with' (withelem (',' withelem)*) 'do' rawseq ('else' rawseq)? 'end'
  | 'try' rawseq ('else' rawseq)? ('then' rawseq)? 'end'
  | 'recover' cap? rawseq 'end'
  | 'consume' cap? term
  | nextpattern
  ;

term
  : 'if' rawseq 'then' rawseq (elseif | ('else' rawseq))? 'end'
  | 'ifdef' infix 'then' rawseq (elseifdef | ('else' rawseq))? 'end'
  | 'match' rawseq caseexpr* ('else' rawseq)? 'end'
  | 'while' rawseq 'do' rawseq ('else' rawseq)? 'end'
  | 'repeat' rawseq 'until' rawseq ('else' rawseq)? 'end'
  | 'for' idseq 'in' rawseq 'do' rawseq ('else' rawseq)? 'end'
  | 'with' (withelem (',' withelem)*) 'do' rawseq ('else' rawseq)? 'end'
  | 'try' rawseq ('else' rawseq)? ('then' rawseq)? 'end'
  | 'recover' cap? rawseq 'end'
  | 'consume' cap? term
  | pattern
  ;

withelem
  : idseq '=' rawseq
  ;

caseexpr
  : '|' pattern? ('if' rawseq)? ('=>' rawseq)?
  ;

elseifdef
  : 'elseif' infix 'then' rawseq (elseifdef | ('else' rawseq))?
  ;

elseif
  : 'elseif' rawseq 'then' rawseq (elseif | ('else' rawseq))?
  ;

idseq
  : ID
  | '_'
  | ('(' | LPAREN_NEW) idseq (',' idseq)* ')'
  ;

nextpattern
  : ('var' | 'let' | 'embed') ID (':' type)?
  | nextparampattern
  ;

pattern
  : ('var' | 'let' | 'embed') ID (':' type)?
  | parampattern
  ;

nextparampattern
  : ('not' | 'addressof' | MINUS_NEW | 'identityof') parampattern
  | nextpostfix
  ;

parampattern
  : ('not' | 'addressof' | '-' | MINUS_NEW | 'identityof') parampattern
  | postfix
  ;

nextpostfix
  : nextatom antlr_2*
  ;

postfix
  : atom antlr_3*
  ;

call
  : '(' positional? named? ')'
  ;

tilde
  : '~' ID
  ;

dot
  : '.' ID
  ;

nextatom
  : ID
  | literal
  | LPAREN_NEW (rawseq | '_') tuple? ')'
  | LSQUARE_NEW ('as' type ':')? rawseq (',' rawseq)* ']'
  | 'object' cap? ('is' type)? members 'end'
  | 'lambda' cap? typeparams? ('(' | LPAREN_NEW) params? ')' lambdacaptures? (':' type)? '?'? '=>' rawseq 'end'
  | '@' (ID | STRING) typeargs? ('(' | LPAREN_NEW) positional? named? ')' '?'?
  ;

atom
  : ID
  | literal
  | ('(' | LPAREN_NEW) (rawseq | '_') tuple? ')'
  | ('[' | LSQUARE_NEW) ('as' type ':')? rawseq (',' rawseq)* ']'
  | 'object' cap? ('is' type)? members 'end'
  | 'lambda' cap? typeparams? ('(' | LPAREN_NEW) params? ')' lambdacaptures? (':' type)? '?'? '=>' rawseq 'end'
  | '@' (ID | STRING) typeargs? ('(' | LPAREN_NEW) positional? named? ')' '?'?
  ;

literal
  : 'this'
  | 'true'
  | 'false'
  | INT
  | FLOAT
  | STRING
  ;

tuple
  : ',' (rawseq | '_') (',' (rawseq | '_'))*
  ;

lambdacaptures
  : ('(' | LPAREN_NEW) lambdacapture (',' lambdacapture)* ')'
  ;

lambdacapture
  : ID (':' type)? ('=' infix)?
  ;

positional
  : rawseq (',' rawseq)*
  ;

named
  : 'where' namedarg (',' namedarg)*
  ;

namedarg
  : ID '=' rawseq
  ;

type
  : atomtype ('->' type)?
  ;

atomtype
  : 'this'
  | 'box'
  | ('(' | LPAREN_NEW) (infixtype | '_') tupletype? ')'
  | nominal
  ;

tupletype
  : ',' (infixtype | '_') (',' (infixtype | '_'))*
  ;

infixtype
  : type antlr_4*
  ;

isecttype
  : '&' type
  ;

uniontype
  : '|' type
  ;

nominal
  : ID ('.' ID)? typeargs? (cap | gencap)? ('^' | '!')?
  ;

gencap
  : '#read'
  | '#send'
  | '#share'
  | '#any'
  ;

cap
  : 'iso'
  | 'trn'
  | 'ref'
  | 'val'
  | 'box'
  | 'tag'
  ;

typeargs
  : '[' type (',' type)* ']'
  ;

typeparams
  : ('[' | LSQUARE_NEW) typeparam (',' typeparam)* ']'
  ;

params
  : (param | '...') (',' (param | '...'))*
  ;

typeparam
  : ID (':' type)? ('=' type)?
  ;

param
  : (parampattern | '_') (':' type)? ('=' infix)?
  ;

antlr_0
  : binop
  | 'as' type
  ;

antlr_1
  : binop
  | 'as' type
  ;

antlr_2
  : dot
  | tilde
  | typeargs
  | call
  ;

antlr_3
  : dot
  | tilde
  | typeargs
  | call
  ;

antlr_4
  : uniontype
  | isecttype
  ;

// Rules of the form antlr_* are only present to avoid a bug in the
// interpreter

/* Precedence

Value:
1. postfix
2. unop
3. binop
4. =
5. seq
6. ,

Type:
1. ->
2. & |
3. ,
*/

// Lexer

ID
  : LETTER (LETTER | DIGIT | '_' | '\'')*
  | '_' (LETTER | DIGIT | '_' | '\'')+
  ;

INT
  : DIGIT (DIGIT | '_')*
  | '0' 'x' (HEX | '_')+
  | '0' 'b' (BINARY | '_')+
  | '\'' CHAR_CHAR* '\''
  ;

FLOAT
  : DIGIT (DIGIT | '_')* ('.' DIGIT (DIGIT | '_')*)? EXP?
  ;

STRING
  : '"' STRING_CHAR* '"'
  | '"""' (('"' | '""') ? ~'"')* '"""' '"'*
  ;

LPAREN_NEW
  : NEWLINE '('
  ;

LSQUARE_NEW
  : NEWLINE '['
  ;

MINUS_NEW
  : NEWLINE '-'
  ;

LINECOMMENT
  : '//' ~('\n')* {$channel = HIDDEN;}
  ;

NESTEDCOMMENT
  : '/*' (NESTEDCOMMENT | '/' ~'*' | ~('*' | '/') | ('*'+ ~('*' | '/')))* '*'+ '/' {$channel = HIDDEN;}
  ;

WS
  : (' ' | '\t' | '\r')+ {$channel = HIDDEN;}
  ;

NEWLINE
  : '\n' (' ' | '\t' | '\r')* {$channel = HIDDEN;}
  ;

fragment
CHAR_CHAR
  : ESC
  | ~('\'' | '\\')
  ;

fragment
STRING_CHAR
  : ESC
  | ~('"' | '\\')
  ;

fragment
EXP
  : ('e' | 'E') ('+' | '-')? (DIGIT | '_')+
  ;

fragment
LETTER
  : 'a'..'z' | 'A'..'Z'
  ;

fragment
BINARY
  : '0'..'1'
  ;

fragment
DIGIT
  : '0'..'9'
  ;

fragment
HEX
  : DIGIT | 'a'..'f' | 'A'..'F' // Only small letters
  ;

fragment
ESC
  : '\\' ('a' | 'b' | 'e' | 'f' | 'n' | 'r' | 't' | 'v' | '\"' | '\\' | '0')
  | HEX_ESC
  | UNICODE_ESC
  | UNICODE2_ESC
  ;

fragment
HEX_ESC
  : '\\' 'x' HEX HEX
  ;

fragment
UNICODE_ESC
  : '\\' 'u' HEX HEX HEX HEX
  ;

fragment
UNICODE2_ESC
  : '\\' 'U' HEX HEX HEX HEX HEX HEX
  ;
```

## License

The MIT License (MIT)

Copyright (c) 2015 pyros2097

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to u-se, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


# :yum: yum
Early draft/designs of a programming language I want to create

No Package Management (KISS, DRY), only one way to do things. You do something do it well don't reproduce stuff. Its Highly opionated. ex: cli, no-term, http. In go we mostly only need to use the std package. All new packages will need to request to be added.
All std packages will be in a separate repo in the same organization.
Third party packages will also have the same convention and the oyrganization will act as the registry and optimize it using some generated json file or something.
This is so that it is modular and that each package can be updated separately.

Maybe integrated into vim completely.

Fully Object Oriented no primitive types only objects to work on. Everything is an object.
Whether to have functions as first class citizens.

No Globals(maybe)
Reduce Capital Letters (unlike go)
Reduce using Shift Key

No Math Library, Everything should be encapsulated within Number class like pony.

No functions only classes and objects. Maybe static.

Scripting support and can be compiled also. But default it should be a lot like power shell (i.e. shell scripting features inbuilt (ls, grep, sed, cut) and also possible to make complete console ui's using it.(only console for now)

## Syntax
PEG Parser with integrated lexing and maybe incremental compiling.
The compiler will format all code and also on lint errors fail. They may be separate processes though. The standard case is camelCase and no other. 2 Spaces or tab width = 2 for indentation.

For faster compilation maybe the compiler will only check the current package
source and does not check whether the function exists in other packages.
This type checking should be done by the gocode like daemon within the editor
this will speed up compilation time (maybe).

The linter, formatter, compiler all will be within itself. The language should be very strict and highly opionated to and must be like a single person use case. It cannot be diverged from the de-facto standard layout.

## Type System
Try to implement much of this as classes like java, pony makes it easier to understand
0. None  
1. Number -> (dec64) includes all math function no need to have a separate math package and has functions for type convertion to u16, u32 (maybe as this might not be needed for FFI we can and should be able to automatically find its value)  
2. String -> includes all string functions including regex, strings  
3. Array(Maybe streams or like golang io.Reader/Writer)  
4. Map  
5. Class -> Includes enums/FSM 
6. Primitives ->  
7. Unions  
8. Attributes/tags (maybe)  
9. Arrow functions  
10. JSON as first class data storage format (maybe)  
11. Byte  
12. YOB/TOB (similar to gob)  

## Traits
1. Composition over direct inheritance
2. Currying
3. Partials

## Classes
1. Single Inheritance Level
2. Static Methods (Maybe no primitives seem better)
3. Decorators //maybe
4. Annotations // maybe
5. Should be used Declaratively like golang
6. Unit Tests within methods/or next to methods (maybe) or classes like pony
7. Contracts
8. Strong typing method signatures
9. Loose typing local variables
10. No @ symbol for this, works like java
11. Pascal Case

## Variables
1. Snake Case or Camel Case decide (pony looks good)
2. Mutable/Immutable
3. let or var
4. local/ global (maybe no globals at all only const)
```pony
import sort
import math

const PI = 3.14
cont Theta = 360
cont E = 2.7

enum Result<T, E>
  OK(T)
  Err(E)
  
enum Option<T>
  None
  Some(T)

trait Animal
  fun HasLegs(): Bool

class Example has Animal
  var x: String
  var y: String
  var z: String
  var r: Number
  var k: Number | None

  fun init(_a: Number, _b : Number) =>
    """
    wrapped in getters and setters automatically this can be overridden
    """
    a = _a
    b = _b
    
  fun calculate(min: Number): Number =>
    '''
    Test Doc
    '''
    require min > 0 # contracts
    result = math.Random(min)
    ensure result > 0
    return result
    
  nextPos(hasNext: Bool): Number =>
    return 4
    
Example{a: 10, b: 10}

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
```

## Match
Must be easy to use
```pony
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
Inbuilt Support for graphql and relay including a single store (Haxl caching)
Maybe integrate math features like github.com/non/spire

## Integrated Testing and Benchmarking methods (think cobra)

## JS
1. Cross platform GL Layer (glium) and NUI based on libGDX or kivy
2. Convert transpile classes/to js classes or create a js runtime similar to scala.js or Dart or that Intellij lang or clojure
3. Defacto JS Library Simulacra.js for rendering to web,desktop,mobile (Ex: GWT) using structs/classes mapped to html templates
4. Defacto Flexbox with css for layouting for GUI components (facebook-css implementation)
5. Implement DOM model for graphics rendering for easier cross platform coding.
6.python zodb or atom or graphene

## Concurrency
1. Single Event Loop (message queue or streams)
2. Async await or yeild resume

## Streams
A major part in std lib. All long operations must use streams. Need to improve error handling and sync streams. highlanderjs. Concurrent Streams.
Streams are good for functional programming and channels and async and gui single event loop.
Runtime use epoll or coroutines or MIO like event mechanism. Functional

## List of STD packages
1. json (maybe integrated into class)
2. http
3. websocket
4. jwt
5. files
6. term
7. ui
8. app
9. opengl
10. scrypt
11. gzip
12. csv
13. gob
14. cli
15. 

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
  | BYTE
  | NUMBER
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

Byte
  : DIGIT (DIGIT | '_')*
  | '0' 'x' (HEX | '_')+
  ;

Number
  : DIGIT (DIGIT | '_')* ('.' DIGIT (DIGIT | '_')*)?
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
STRING_CHAR
  : ESC
  | ~('"' | '\\')
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


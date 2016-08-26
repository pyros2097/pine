{
  var TYPES_TO_PROPERTY_NAMES = {
    CallExpression:   "callee",
    MemberExpression: "object",
  };

  function filledArray(count, value) {
    var result = new Array(count), i;

    for (i = 0; i < count; i++) {
      result[i] = value;
    }

    return result;
  }

  function extractOptional(optional, index) {
    return optional ? optional[index] : null;
  }

  function extractList(list, index) {
    var result = new Array(list.length), i;

    for (i = 0; i < list.length; i++) {
      result[i] = list[i][index];
    }

    return result;
  }

  function buildList(head, tail, index) {
    return [head].concat(extractList(tail, index));
  }

  function buildTree(head, tail, builder) {
    var result = head, i;

    for (i = 0; i < tail.length; i++) {
      result = builder(result, tail[i]);
    }

    return result;
  }

  function buildBinaryExpression(head, tail) {
    return buildTree(head, tail, function(result, element) {
      return {
        type:     "BinaryExpression",
        operator: element[1],
        left:     result,
        right:    element[3]
      };
    });
  }

  function buildLogicalExpression(head, tail) {
    return buildTree(head, tail, function(result, element) {
      return {
        type:     "LogicalExpression",
        operator: element[1],
        left:     result,
        right:    element[3]
      };
    });
  }

  function optionalList(value) {
    return value !== null ? value : [];
  }
}

Start
  = __ program:Program __ { return program; }

/* ----- A.1 Lexical Grammar ----- */

SourceCharacter
  = .

WhiteSpace "whitespace"
  = "\t"
  / "\v"
  / "\f"
  / " "

LineTerminator
  = [\n\r\u2028\u2029]

LineTerminatorSequence "end of line"
  = "\n"
  / "\r\n"
  / "\r"

Comment "comment"
  = MultiLineComment
  / SingleLineComment

MultiLineComment
  = "---" (!"---" SourceCharacter)* "---"

MultiLineCommentNoLineTerminator
  = "---" (!("---" / LineTerminator) SourceCharacter)* "---"

SingleLineComment
  = "#" (!LineTerminator SourceCharacter)*

Identifier
  = !ReservedWord name:IdentifierName { return name; }

IdentifierName "identifier"
  = head:IdentifierStart tail:IdentifierPart* {
      return {
        type: "Identifier",
        name: head + tail.join(""),
        public: true,
      };
    }

IdentifierStart
  = [_a-z]

IdentifierPart
  = [a-zA-Z0-9]

ReservedWord
  = Keyword
  / NullLiteral
  / BooleanLiteral

Keyword
  = BreakToken
  / CaseToken
  / CatchToken
  / ContinueToken
  / DebuggerToken
  / DefaultToken
  / DoToken
  / ElseToken
  / FinallyToken
  / ForToken
  / FunctionToken
  / IfToken
  / InstanceofToken
  / InToken
  / NewToken
  / ReturnToken
  / SwitchToken
  / ThisToken
  / ThrowToken
  / TryToken
  / TypeofToken
  / VarToken
  / WhileToken
  / WithToken
  / ConstToken
  / EnumToken
  / ExtendsToken
  / ImportToken
  / SuperToken
  / ClassToken

Literal
  = NullLiteral
  / BooleanLiteral
  / NumericLiteral
  / StringLiteral

NullLiteral
  = NullToken { return { type: "NullLiteral", value: null }; }

BooleanLiteral
  = TrueToken  { return { type: "BooleanLiteral", value: true  }; }
  / FalseToken { return { type: "BooleanLiteral", value: false }; }

NumericLiteral "number"
  = DecimalIntegerLiteral "."? DecimalDigit* {
      return { type: "NumericLiteral", value: parseFloat(text()) };
    }

DecimalIntegerLiteral
  = "0"
  / NonZeroDigit DecimalDigit*

DecimalDigit
  = [0-9]

NonZeroDigit
  = [1-9]

SignedInteger
  = [+-]? DecimalDigit+

StringLiteral "string"
  =  "'" chars:SingleStringCharacter* "'" {
      return { type: "StringLiteral", value: chars.join("") };
    }

SingleStringCharacter
  = !("'" / "\\" / LineTerminator) SourceCharacter { return text(); }
  / "\\" sequence:EscapeSequence { return sequence; }
  / LineContinuation

LineContinuation
  = "\\" LineTerminatorSequence { return ""; }

EscapeSequence
  = CharacterEscapeSequence
  / "0" !DecimalDigit { return "\0"; }

CharacterEscapeSequence
  = SingleEscapeCharacter
  / NonEscapeCharacter

SingleEscapeCharacter
  = "'"
  / '"'
  / "\\"
  / "b"  { return "\b";   }
  / "f"  { return "\f";   }
  / "n"  { return "\n";   }
  / "r"  { return "\r";   }
  / "t"  { return "\t";   }
  / "v"  { return "\x0B"; }   // IE does not recognize "\v".

NonEscapeCharacter
  = !(EscapeCharacter / LineTerminator) SourceCharacter { return text(); }

EscapeCharacter
  = SingleEscapeCharacter
  / DecimalDigit

/* Tokens */

BreakToken      = "break"      !IdentifierPart
CaseToken       = "case"       !IdentifierPart
CatchToken      = "catch"      !IdentifierPart
ClassToken      = "class"      !IdentifierPart
ConstToken      = "const"      !IdentifierPart
ContinueToken   = "continue"   !IdentifierPart
DebuggerToken   = "debugger"   !IdentifierPart
DefaultToken    = "default"    !IdentifierPart
DoToken         = "do"         !IdentifierPart
ElseToken       = "else"       !IdentifierPart
EndToken        = "end"        !IdentifierPart
EnumToken       = "enum"       !IdentifierPart
ExtendsToken    = "extends"    !IdentifierPart
FalseToken      = "false"      !IdentifierPart
FinallyToken    = "finally"    !IdentifierPart
ForToken        = "for"        !IdentifierPart
FunctionToken   = "function"   !IdentifierPart
GetToken        = "get"        !IdentifierPart
IfToken         = "if"         !IdentifierPart
ImportToken     = "import"     !IdentifierPart
InstanceofToken = "instanceof" !IdentifierPart
InToken         = "in"         !IdentifierPart
NewToken        = "new"        !IdentifierPart
NullToken       = "null"       !IdentifierPart
ReturnToken     = "return"     !IdentifierPart
SetToken        = "set"        !IdentifierPart
SuperToken      = "super"      !IdentifierPart
SwitchToken     = "switch"     !IdentifierPart
ThisToken       = "this"       !IdentifierPart
ThrowToken      = "throw"      !IdentifierPart
TrueToken       = "true"       !IdentifierPart
TryToken        = "try"        !IdentifierPart
TypeofToken     = "typeof"     !IdentifierPart
VarToken        = "var"        !IdentifierPart
WhileToken      = "while"      !IdentifierPart
WithToken       = "with"       !IdentifierPart

/* Skipped */

__
  = (WhiteSpace / LineTerminatorSequence / Comment)*

_
  = (WhiteSpace / MultiLineCommentNoLineTerminator)*

/* Automatic Semicolon Insertion */

EOS
  = __ ";"
  / _ SingleLineComment? LineTerminatorSequence
  / _ &"}"
  / __ EOF

EOF
  = !.

/* ----- A.3 Expressions ----- */

PrimaryExpression
  = ThisToken { return { type: "ThisExpression" }; }
  / Identifier
  / Literal
  / ArrayLiteral
  / ObjectLiteral
  / "(" __ expression:Expression __ ")" { return expression; }

ArrayLiteral
  = "[" __ elision:(Elision __)? "]" {
      return {
        type:     "ArrayExpression",
        elements: optionalList(extractOptional(elision, 0))
      };
    }
  / "[" __ elements:ElementList __ "]" {
      return {
        type:     "ArrayExpression",
        elements: elements
      };
    }
  / "[" __ elements:ElementList __ "," __ elision:(Elision __)? "]" {
      return {
        type:     "ArrayExpression",
        elements: elements.concat(optionalList(extractOptional(elision, 0)))
      };
    }

ElementList
  = head:(
      elision:(Elision __)? element:AssignmentExpression {
        return optionalList(extractOptional(elision, 0)).concat(element);
      }
    )
    tail:(
      __ "," __ elision:(Elision __)? element:AssignmentExpression {
        return optionalList(extractOptional(elision, 0)).concat(element);
      }
    )*
    { return Array.prototype.concat.apply(head, tail); }

Elision
  = "," commas:(__ ",")* { return filledArray(commas.length + 1, null); }

ObjectLiteral
  = "{" __ "}" { return { type: "ObjectExpression", properties: [] }; }
  / "{" __ properties:PropertyNameAndValueList __ "}" {
       return { type: "ObjectExpression", properties: properties };
     }
  / "{" __ properties:PropertyNameAndValueList __ "," __ "}" {
       return { type: "ObjectExpression", properties: properties };
     }
PropertyNameAndValueList
  = head:PropertyAssignment tail:(__ "," __ PropertyAssignment)* {
      return buildList(head, tail, 3);
    }

PropertyAssignment
  = key:PropertyName __ ":" __ value:AssignmentExpression {
      return { key: key, value: value, kind: "init" };
    }
  / GetToken __ key:PropertyName __
    "(" __ ")" __
    "{" __ body:FunctionBody __ "}"
    {
      return {
        key:   key,
        value: {
          type:   "FunctionExpression",
          id:     null,
          params: [],
          body:   body
        },
        kind:  "get"
      };
    }
  / SetToken __ key:PropertyName __
    "(" __ params:PropertySetParameterList __ ")" __
    "{" __ body:FunctionBody __ "}"
    {
      return {
        key:   key,
        value: {
          type:   "FunctionExpression",
          id:     null,
          params: params,
          body:   body
        },
        kind:  "set"
      };
    }

PropertyName
  = IdentifierName
  / StringLiteral
  / NumericLiteral

PropertySetParameterList
  = id:Identifier { return [id]; }

MemberExpression
  = head:(
        PrimaryExpression
      / FunctionExpression
      / NewToken __ callee:MemberExpression __ args:Arguments {
          return { type: "NewExpression", callee: callee, arguments: args };
        }
    )
    tail:(
        __ "[" __ property:Expression __ "]" {
          return { property: property, computed: true };
        }
      / __ "." __ property:IdentifierName {
          return { property: property, computed: false };
        }
    )*
    {
      return buildTree(head, tail, function(result, element) {
        return {
          type:     "MemberExpression",
          object:   result,
          property: element.property,
          computed: element.computed
        };
      });
    }

NewExpression
  = MemberExpression
  / NewToken __ callee:NewExpression {
      return { type: "NewExpression", callee: callee, arguments: [] };
    }

CallExpression
  = head:(
      callee:MemberExpression __ args:Arguments {
        return { type: "CallExpression", callee: callee, arguments: args };
      }
    )
    tail:(
        __ args:Arguments {
          return { type: "CallExpression", arguments: args };
        }
      / __ "[" __ property:Expression __ "]" {
          return {
            type:     "MemberExpression",
            property: property,
            computed: true
          };
        }
      / __ "." __ property:IdentifierName {
          return {
            type:     "MemberExpression",
            property: property,
            computed: false
          };
        }
    )*
    {
      return buildTree(head, tail, function(result, element) {
        element[TYPES_TO_PROPERTY_NAMES[element.type]] = result;

        return element;
      });
    }

Arguments
  = "(" __ args:(ArgumentList __)? ")" {
      return optionalList(extractOptional(args, 0));
    }

ArgumentList
  = head:AssignmentExpression tail:(__ "," __ AssignmentExpression)* {
      return buildList(head, tail, 3);
    }

LeftHandSideExpression
  = CallExpression
  / NewExpression

PostfixExpression
  = argument:LeftHandSideExpression _ operator:PostfixOperator {
      return {
        type:     "UpdateExpression",
        operator: operator,
        argument: argument,
        prefix:   false
      };
    }
  / LeftHandSideExpression

PostfixOperator
  = "++"
  / "--"

UnaryExpression
  = PostfixExpression
  / operator:UnaryOperator __ argument:UnaryExpression {
      var type = (operator === "++" || operator === "--")
        ? "UpdateExpression"
        : "UnaryExpression";

      return {
        type:     type,
        operator: operator,
        argument: argument,
        prefix:   true
      };
    }

UnaryOperator
  = $TypeofToken
  / "++"
  / "--"
  / $("+" !"=")
  / $("-" !"=")
  / "~"
  / "!"

MultiplicativeExpression
  = head:UnaryExpression
    tail:(__ MultiplicativeOperator __ UnaryExpression)*
    { return buildBinaryExpression(head, tail); }

MultiplicativeOperator
  = $("*" !"=")
  / $("/" !"=")
  / $("%" !"=")

AdditiveExpression
  = head:MultiplicativeExpression
    tail:(__ AdditiveOperator __ MultiplicativeExpression)*
    { return buildBinaryExpression(head, tail); }

AdditiveOperator
  = $("+" ![+=])
  / $("-" ![-=])

ShiftExpression
  = head:AdditiveExpression
    tail:(__ ShiftOperator __ AdditiveExpression)*
    { return buildBinaryExpression(head, tail); }

ShiftOperator
  = $("<<"  !"=")
  / $(">>>" !"=")
  / $(">>"  !"=")

RelationalExpression
  = head:ShiftExpression
    tail:(__ RelationalOperator __ ShiftExpression)*
    { return buildBinaryExpression(head, tail); }

RelationalOperator
  = "<="
  / ">="
  / $("<" !"<")
  / $(">" !">")
  / $InstanceofToken
  / $InToken

RelationalExpressionNoIn
  = head:ShiftExpression
    tail:(__ RelationalOperatorNoIn __ ShiftExpression)*
    { return buildBinaryExpression(head, tail); }

RelationalOperatorNoIn
  = "<="
  / ">="
  / $("<" !"<")
  / $(">" !">")
  / $InstanceofToken

EqualityExpression
  = head:RelationalExpression
    tail:(__ EqualityOperator __ RelationalExpression)*
    { return buildBinaryExpression(head, tail); }

EqualityExpressionNoIn
  = head:RelationalExpressionNoIn
    tail:(__ EqualityOperator __ RelationalExpressionNoIn)*
    { return buildBinaryExpression(head, tail); }

EqualityOperator
  = "==="
  / "!=="
  / "=="
  / "!="

BitwiseANDExpression
  = head:EqualityExpression
    tail:(__ BitwiseANDOperator __ EqualityExpression)*
    { return buildBinaryExpression(head, tail); }

BitwiseANDExpressionNoIn
  = head:EqualityExpressionNoIn
    tail:(__ BitwiseANDOperator __ EqualityExpressionNoIn)*
    { return buildBinaryExpression(head, tail); }

BitwiseANDOperator
  = $("&" ![&=])

BitwiseXORExpression
  = head:BitwiseANDExpression
    tail:(__ BitwiseXOROperator __ BitwiseANDExpression)*
    { return buildBinaryExpression(head, tail); }

BitwiseXORExpressionNoIn
  = head:BitwiseANDExpressionNoIn
    tail:(__ BitwiseXOROperator __ BitwiseANDExpressionNoIn)*
    { return buildBinaryExpression(head, tail); }

BitwiseXOROperator
  = $("^" !"=")

BitwiseORExpression
  = head:BitwiseXORExpression
    tail:(__ BitwiseOROperator __ BitwiseXORExpression)*
    { return buildBinaryExpression(head, tail); }

BitwiseORExpressionNoIn
  = head:BitwiseXORExpressionNoIn
    tail:(__ BitwiseOROperator __ BitwiseXORExpressionNoIn)*
    { return buildBinaryExpression(head, tail); }

BitwiseOROperator
  = $("|" ![|=])

LogicalANDExpression
  = head:BitwiseORExpression
    tail:(__ LogicalANDOperator __ BitwiseORExpression)*
    { return buildBinaryExpression(head, tail); }

LogicalANDExpressionNoIn
  = head:BitwiseORExpressionNoIn
    tail:(__ LogicalANDOperator __ BitwiseORExpressionNoIn)*
    { return buildBinaryExpression(head, tail); }

LogicalANDOperator
  = "and"

LogicalORExpression
  = head:LogicalANDExpression
    tail:(__ LogicalOROperator __ LogicalANDExpression)*
    { return buildBinaryExpression(head, tail); }

LogicalORExpressionNoIn
  = head:LogicalANDExpressionNoIn
    tail:(__ LogicalOROperator __ LogicalANDExpressionNoIn)*
    { return buildBinaryExpression(head, tail); }

LogicalOROperator
  = "or"

ConditionalExpression
  = test:LogicalORExpression __
    "?" __ consequent:AssignmentExpression __
    ":" __ alternate:AssignmentExpression
    {
      return {
        type:       "ConditionalExpression",
        test:       test,
        consequent: consequent,
        alternate:  alternate
      };
    }
  / LogicalORExpression

ConditionalExpressionNoIn
  = test:LogicalORExpressionNoIn __
    "?" __ consequent:AssignmentExpression __
    ":" __ alternate:AssignmentExpressionNoIn
    {
      return {
        type:       "ConditionalExpression",
        test:       test,
        consequent: consequent,
        alternate:  alternate
      };
    }
  / LogicalORExpressionNoIn

AssignmentExpression
  = left:LeftHandSideExpression __
    "=" !"=" __
    right:AssignmentExpression
    {
      return {
        type:     "AssignmentExpression",
        operator: "=",
        left:     left,
        right:    right
      };
    }
  / ConditionalExpression

AssignmentExpressionNoIn
  = left:LeftHandSideExpression __
    "=" !"=" __
    right:AssignmentExpressionNoIn
    {
      return {
        type:     "AssignmentExpression",
        operator: "=",
        left:     left,
        right:    right
      };
    }
  / ConditionalExpressionNoIn

Expression
  = head:AssignmentExpression tail:(__ "," __ AssignmentExpression)* {
      return tail.length > 0
        ? { type: "SequenceExpression", expressions: buildList(head, tail, 3) }
        : head;
    }

ExpressionNoIn
  = head:AssignmentExpressionNoIn tail:(__ "," __ AssignmentExpressionNoIn)* {
      return tail.length > 0
        ? { type: "SequenceExpression", expressions: buildList(head, tail, 3) }
        : head;
    }

/* ----- A.4 Statements ----- */

Statement
  = Block
  / VariableStatement
  / ExpressionStatement
  / IfStatement
  / ForStatement
  / ContinueStatement
  / BreakStatement
  / ReturnStatement
  / WithStatement
  / SwitchStatement
  / ThrowStatement
  / TryStatement
  / DebuggerStatement
  / EnumStatement
  / ClassStatement

Block
  = "{" __ body:(StatementList __)? "}" {
      return {
        type: "BlockStatement",
        body: optionalList(extractOptional(body, 0))
      };
    }

StatementList
  = head:Statement tail:(__ Statement)* { return buildList(head, tail, 1); }

VariableStatement
  = VarToken __ declarations:VariableDeclarationList EOS {
      return {
        type:         "VariableDeclaration",
        declarations: declarations
      };
    }

VariableDeclarationList
  = head:VariableDeclaration tail:(__ "," __ VariableDeclaration)* {
      return buildList(head, tail, 3);
    }

VariableDeclarationListNoIn
  = head:VariableDeclarationNoIn tail:(__ "," __ VariableDeclarationNoIn)* {
      return buildList(head, tail, 3);
    }

VariableDeclaration
  = id:Identifier init:(__ Initialiser)? {
      return {
        type: "VariableDeclarator",
        id:   id,
        init: extractOptional(init, 1)
      };
    }

VariableDeclarationNoIn
  = id:Identifier init:(__ InitialiserNoIn)? {
      return {
        type: "VariableDeclarator",
        id:   id,
        init: extractOptional(init, 1)
      };
    }

Initialiser
  = "=" !"=" __ expression:AssignmentExpression { return expression; }

InitialiserNoIn
  = "=" !"=" __ expression:AssignmentExpressionNoIn { return expression; }

ExpressionStatement
  = !("{" / FunctionToken) expression:Expression EOS {
      return {
        type:       "ExpressionStatement",
        expression: expression
      };
    }

EnumStatement
  = EnumToken __ id:Identifier __ expression:Expression __
    {
      return {
        type: "EnumStatement",
        id: id,
        expression: expression
      };
    }

ClassStatement
  = ClassToken __ id:Identifier __ expression:Expression __
    {
      return {
        type: "ClassStatement",
        id: id,
        expression: expression
      };
    }

IfStatement
  = IfToken __  test:LogicalORExpression __ "then" __
    consequent:Statement __
    ElseToken __
    alternate:Statement
    EndToken __
    {
      return {
        type:       "IfStatement",
        test:       test,
        consequent: consequent,
        alternate:  alternate
      };
    }
  / IfToken __ test:LogicalORExpression __ "then" __
    consequent:Statement __
    EndToken __
    {
      return {
        type:       "IfStatement",
        test:       test,
        consequent: consequent,
        alternate:  null
      };
    }

ForStatement
  = ForToken __ "do" __
    body:Statement
    EndToken __
    { return { type: "ForInfiniteStatement", body: body }; }
  / ForToken __ test:LogicalORExpression __ "do" __
    body:Statement
    EndToken __
    { return { type: "ForStatement", test: test, body: body }; }
  / ForToken __ left:Identifier __ InToken __ right:Expression __ "do" __
    body:Statement
    EndToken __
    {
      return {
        type:  "ForInStatement",
        left:  left,
        right: right,
        body:  body
      };
    }

ContinueStatement
  = ContinueToken EOS {
      return { type: "ContinueStatement"};
    }

BreakStatement
  = BreakToken EOS {
      return { type: "BreakStatement"};
    }

ReturnStatement
  = ReturnToken EOS {
      return { type: "ReturnStatement", argument: null };
    }
  / ReturnToken _ argument:Expression EOS {
      return { type: "ReturnStatement", argument: argument };
    }

WithStatement
  = WithToken __ "(" __ object:Expression __ ")" __
    body:Statement
    { return { type: "WithStatement", object: object, body: body }; }

SwitchStatement
  = SwitchToken __ "(" __ discriminant:Expression __ ")" __
    cases:CaseBlock
    {
      return {
        type:         "SwitchStatement",
        discriminant: discriminant,
        cases:        cases
      };
    }

CaseBlock
  = "{" __ clauses:(CaseClauses __)? "}" {
      return optionalList(extractOptional(clauses, 0));
    }
  / "{" __
    before:(CaseClauses __)?
    default_:DefaultClause __
    after:(CaseClauses __)? "}"
    {
      return optionalList(extractOptional(before, 0))
        .concat(default_)
        .concat(optionalList(extractOptional(after, 0)));
    }

CaseClauses
  = head:CaseClause tail:(__ CaseClause)* { return buildList(head, tail, 1); }

CaseClause
  = CaseToken __ test:Expression __ ":" consequent:(__ StatementList)? {
      return {
        type:       "SwitchCase",
        test:       test,
        consequent: optionalList(extractOptional(consequent, 1))
      };
    }

DefaultClause
  = DefaultToken __ ":" consequent:(__ StatementList)? {
      return {
        type:       "SwitchCase",
        test:       null,
        consequent: optionalList(extractOptional(consequent, 1))
      };
    }

ThrowStatement
  = ThrowToken _ argument:Expression EOS {
      return { type: "ThrowStatement", argument: argument };
    }

TryStatement
  = TryToken __ block:Block __ handler:Catch __ finalizer:Finally {
      return {
        type:      "TryStatement",
        block:     block,
        handler:   handler,
        finalizer: finalizer
      };
    }
  / TryToken __ block:Block __ handler:Catch {
      return {
        type:      "TryStatement",
        block:     block,
        handler:   handler,
        finalizer: null
      };
    }
  / TryToken __ block:Block __ finalizer:Finally {
      return {
        type:      "TryStatement",
        block:     block,
        handler:   null,
        finalizer: finalizer
      };
    }

Catch
  = CatchToken __ "(" __ param:Identifier __ ")" __ body:Block {
      return {
        type:  "CatchClause",
        param: param,
        body:  body
      };
    }

Finally
  = FinallyToken __ block:Block { return block; }

DebuggerStatement
  = DebuggerToken EOS { return { type: "DebuggerStatement" }; }

/* ----- A.5 Functions and Programs ----- */

FunctionDeclaration
  = FunctionToken __ id:Identifier __
    "(" __ params:(FormalParameterList __)? ")" __
    "{" __ body:FunctionBody __ "}"
    {
      return {
        type:   "FunctionDeclaration",
        id:     id,
        params: optionalList(extractOptional(params, 0)),
        body:   body
      };
    }

FunctionExpression
  = FunctionToken __ id:(Identifier __)?
    "(" __ params:(FormalParameterList __)? ")" __
    "{" __ body:FunctionBody __ "}"
    {
      return {
        type:   "FunctionExpression",
        id:     extractOptional(id, 0),
        params: optionalList(extractOptional(params, 0)),
        body:   body
      };
    }

FormalParameterList
  = head:Identifier tail:(__ "," __ Identifier)* {
      return buildList(head, tail, 3);
    }

FunctionBody
  = body:SourceElements? {
      return {
        type: "BlockStatement",
        body: optionalList(body)
      };
    }

Program
  = body:SourceElements? {
      return {
        type: "Program",
        body: optionalList(body)
      };
    }

SourceElements
  = head:SourceElement tail:(__ SourceElement)* {
      return buildList(head, tail, 1);
    }

SourceElement
  = Statement
  / FunctionDeclaration

/*
# hellodasd
---
Comment allez vous?
class Hello
---

var _asdi = 'd'
var bb = 4.0
var kk = 23 + 12.0 + 0.5
continue

function hello() {
  println('dasd')
}

if kk == 0 then {
}
else {
}end

for i in aa do {
}end

for a and b do {
}end

for do {
}end

enum _abc {
}
*/
# :yum: yum
Early draft/designs of a programming language I want to create

Frontend: PEG.js, ANTLR 4, YACC
Backend: LLVM

## Imports/Packages
```js
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

## Interfaces
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
```coffee
class Example

  constructor: (_a, _b) ->
    a = _a  # wrapped in getters and setters automatically this can be overridden
    b = _b
    
  calculate: (min: int) :int ->
    require min > 0 // contracts
    result = Math.random(min)
    ensure result > 0
    return result
```

## Streams
A major part in std lib. All long operations must use streams. Need to improve error handling and sync streams

## Generators
Use them for default async behavior or await or make the language follow nodejs style callback support but the language encapsulates this with syntax like iced coffeescript or livescript

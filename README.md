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

## Functions/Closures
1. camelCase
2. Currying
3. Partials

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
12. Pascal Case
```coffee
class Example

  constructor: (_a :int, _b : int) ->
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
| default -> error
```
# Standard Library

## Math

## Http

## Crypto

## Collections
1. Arrays
2. Map
3. Sets
4. WeakMaps
5. Lists
6. or everything can be a map

## Net

## Dns

## OpenGL

## Cross platform GL layer and UI based on libGDX


## Streams
A major part in std lib. All long operations must use streams. Need to improve error handling and sync streams

## Generators
Use them for default async behavior or await or make the language follow nodejs style callback support but the language encapsulates this with syntax like iced coffeescript or livescript

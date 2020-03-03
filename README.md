# The Yum Programming Language

# First class dependent types

## Example

```js
import (
    "std/routes"
    "std/html"
)

@route(/about/{name}")
func about(name string) -> Html {
    return (
        <Window x="1" y="2" width="480" height="800">
            <VBox>
                <Text>This is my title<Text>
            </VBox>
        </Window>
    )
}

router(host: "localhost", port: 1234)
```

```js
const name = "Bob"
age := 20
const large_number = 9999999999
println(name)
println(age)
println(large_number)
```

## Features

### Power println

```groovy
println complexObject
```

### Power Assert

```groovy
assert x == 2


// Output:
//
// Assertion failed:
// assert x == 2
//        | |
//        1 false
```

### Power Mocks

### Power Stubs

### Null checking

All values are default nullable/nilable you can make them non-nullable by using an exclamation mark after the type;

```go
func area(a int!, b float) {
}
```

### Multiple returns and errors

Documentation parses these errors in the func and lists them as errors the
function can throw (something like throws in java)

```go
func area(x int) (string, error) {
    if x < 0 {
        return nil, error("invalid", "x")
    }
    return "${x+1}", nil
}

func main() error {
    v := area() // if not handled the error is returned to the calling function like try() in golang

    v, err := area()
    if err != nil {
        if err.name == "" {
            println(err.message)
        }
        return err
    }
    println(v)
    return nil
}

func main() error {
    // Or handle the error like usual
    v, err := area()
    if err != nil {
        if err.name == "" {
            println(err.message)
        }
        return err
    }
    println(v)
    return nil
}

```

# Specification

# Primitives types

```go
bool
string
int // thought redo these 3 types to decimal??
float
decimal
byte
string
nil
error
```

# Reference Types

```go
array
map
enum
struct
func
```

# Strings

A string is a read-only array of bytes. String data is encoded using UTF-8. Strings are immutable. This means that the substring function is very efficient: no copying is performed, no extra allocations required.

TODO: use double quotes or single quotes (pref for double since its standard)

```golang
name := "Bob"
println("Hello, $name!") // `$` is used for string interpolation
println(len(name))

bobby := name + "by" // + is used to concatenate strings
println(bobby) // ==> "Bobby"

We have to either convert age to a string:
age := 1
println("age = $age")
// primitives, structs have String() method builtin println works

multiline = """
"""
```

# Arrays

```golang
nums := []int[1, 2, 3]
println(nums) // ==> [1, 2, 3]
println(nums[1]) // ==> 2
println(nums[10]) // ==> nil

names := []string["Krabs"]
names = append(names, "Spongebob") // which format to use
names << "Patrick"
println(len(names)) // ==> 3
println(array.Contains(names, "Krabs")) // ==> true
println(array.Contains(names, "Squidward")) // ==> false
```

# Maps

Keys will always be strings. Keys are ordered by default.

```golang
nums := map[int]{
    "one": 1,
    "two": 2
}
nums["three"] = 4
println(nums["one"]) // ==> "1"
println(nums["infinity"]) // ==> nil
```

# Enums

```golang
type Direction =
    | Left
    | Right
    | Up
    | Down
```

Structs

```golang
struct User {
    id string
    name string
    email string
}
```

# Funcs

### Support named args

```golang
func add(a int, b int) {
}

add(3, 5)
add(b: 4, a: 1)
```

## Types

```golang
type Operator func(a int, b int) bool
```

# Conditional Statements

## if/else if/ else

```golang
a := 10
b := 20
c := a > b ? 10 : a;
if a < b {
    println("$a < $b")
} else if a > b {
    println("$a > $b")
} else {
    println("$a == $b")
}
```

## switch

```golang
match c {
| 1 => _resize(y, x)
| 2 => _processInput(c, x, y)
| 3 => _processInput(c, x, y)
| 4 => _processInput(c, x, y)
| 255 => _exit(); break
}

switch c {
  case 1:
    x = 1
  case 2:
    x = 2
  case 3:
    x = 3
  case 4: x = 4
  case 255: x = 5
  default: x = 0
}
```

# Loops

## Normal loops

```golang
for i := 0; i < 10; i++ {
    println(i, v)
}
```

## Infinite Loop

```golang
for {
    println(i, v)
}
```

## Loop with Condition

```golang
for x > 5 {
    println(i, v)
}
```

## Array Iteration

```golang
nums := [1, 2, 3]
for i, v  := range nums {
    println(i, v)
}
```

## Map Iteration

```golang
data := map[string]{"name": "krab"}
for k, v  := range nums {
    println(k, v)
}
```

# Keywords

```golang
break
const
continue
else
enum
func
for
if
import
in
switch
package
return
struct
type
fallthrough
```

# Operators

## Arithmetic Operators

- \+ (addition)
- \- (subtraction)
- \* (multiplication)
- / (division)
- % (modulo)

## Logical Operators

- && (and)
- || (or)
- ! (not)

## Comparison Operators

- \> (gt)
- <
- />=
- <=
- ==
- !=
- ^

## Assignment Operators

- +=
- -=
- \*=
- /=
- %=

## Operator Precedence

<!-- 5 \* / % << >>

4 + - | ^
3 == != < <= > >=
2 &&
1 || -->

<!--
Thoughts builtin
enum Boolean {
    True
    False

    and(other) {
        this && other
    }

    or(other) {
        this || other
    }s
}

flag := Boolean.true; -->

# Docs

```golang
// Takes a, b
func add(a int, b int) {

}

"""
Takes a, b
"""
func add(a int, b int) {

}

/*
    Takes a, b
*/
func add(a int, b int) {

}
```

# OLD STUFF

## Features

1. No Package Management (KISS, DRY), only one way to do things. You do something do it well don't reproduce stuff. Its Highly opionated. ex: cli, no-term, http. In go we mostly only need to use the std package. All new packages will need to request to be added. All std packages will be in a separate repo in the same organization.
   Third party packages will also have the same convention and the organization will act as the registry and optimize it using some generated json file or something. This is so that it is modular and that each package can be updated separately.
2. Fully Object Oriented no primitive types only objects to work on. Everything is an object (like pony or scala will use comp)
   Whether to have functions as first class citizens.No functions only classes and objects. Maybe static.
   iler hacks to make literals behave like objects)
3. Maybe integrated into vim completely.
4. No Globals(maybe)
5. Proper Error Handling. Stream/Railway oriented design (Maybe)

## Parsing

PEG Parser with integrated lexing and maybe incremental compiling.
The compiler will format all code and also on lint errors fail. They may be separate processes though. The standard case is camelCase and no other. 2 Spaces for indentation.

For faster compilation maybe the compiler will only check the current package
source and does not check whether the function exists in other packages.
This type checking should be done by the gocode like daemon within the editor
this will speed up compilation time (maybe). Inc
remental Compilation (yeah as long as I can prouduce the same machine code every single time)
Only vim support.

The linter, formatter, compiler all will be within itself. The language should be very strict and highly opionated to and must be like a single person use case. It cannot be diverged from the de-facto standard layout.

## Type System

Try to implement much of this as classes like java, pony makes it easier to understand

0. None
1. Number -> (dec64) includes all math function no need to have a separate math package and has functions for type convertion to u16, u32 (maybe as this might not be needed for FFI we can and should be able to automatically find its value)
1. String -> includes all string functions including regex, strings
1. Array(Maybe streams or like golang io.Reader/Writer)
1. Map
1. Class -> Includes enums/FSM
1. Primitives ->
1. Unions
1. Attributes/tags (maybe)
1. Arrow functions
1. JSON as first class data storage format (maybe)
1. Byte
1. YOB/TOB (similar to gob)

## Traits

1. Composition over direct inheritance
2. Currying
3. Partials

## Enums

1. FSM

## Classes (Declared Pascal Case)

1. Single Inheritance Level
2. Static Methods (Maybe no primitives seem better)
3. Decorators (Maybe since we wont have relection it will be good for running before functions)
4. Annotations (Maybe)
5. Should be used Declaratively like golang
6. Unit Tests within methods/or next to methods (maybe) or classes like pony
7. Contracts
8. Strong typing method signatures
9. Loose typing local variables
10. No @ symbol for this, works like java

## Variables (Declared Snake Case or Camel Case decide (pony looks good))

1. Mutable/Immutable
2. let or var
3. local/ global (maybe no globals at all only const)

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

enum Test
  a = 13
  b
  c
  d

enum Result
  Success
  Failure(err)

  unwrap: ->
    os.exit(1)
  unwrapElse: ->
  string: String -> if Failure(err) "#{err}" else base.string()


trait IsAlive
  alive(): Bool

stream A String
Array[String]

class Test
  buffer: Buffer
  state: Test

  init: ()->
    buffer.write("")
    buffer.write("")
    buffer.write("")

  del: ()->
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

# The Yum Language Reference

Literals
1.Number Literal
2.String Literal
3.Boolean Literal
3.Byte Literal
4.Character Literal

Safe Functions
Unsafe Functions

Types
1.Type Aliases
2.Struct
3.Enum
4.Class
5.Trait

Expresssions
Operators

```
1.and
2.or
3.xor
4.not
5.==
6.!=
7.>=
8.>
9.<
10.<=
11.>>
12.<<
13.+
14.-
15./
16.*
```

Some/Most of these operators are implemented as operator overloading so it makes it easier for the compiler to call these functions or inline them or something. (So we wont need to make any change for our number class)

Loops

```
1. Infinite Loop
for
end
2. Loop with Condition
for cond
end
3. Loop within an range
for i in range
end
4. Loop over a 1D array
for i in [1, 2, 3, 4, 5]
end
5. Loop over a 2D array
for i, j in [[1, 2, 3], [4, 5, 6]]
end
5. Loop over a map
for k, v in {a: "b", c: "d"}
end
```

lambdas

```
class MyDispatcher
  work: func(a: Number)

Dispatcher{
  work: fun(a: Number) -> a + 5
}
```

match

```
match code
| 1 -> log('1')
| 4, 5 -> log('test)
| 7 -> {
|   logs
| }
else
  0
end
```

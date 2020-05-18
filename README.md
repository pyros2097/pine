<img src="https://github.com/pyros2097/pine/raw/master/assets/pine.png" width="100" height="100">

# Pine Programming Language

![Tests](https://github.com/pyros2097/yum/workflows/Test/badge.svg)

A programming language with a syntax largely inspired
by nim but with the simplicity of go. It has jsx support in built by
default and targets webassembly.

## Aims

- first class functions
  <!-- - first class dependent types -->
- targets webassembly
- clean syntax
- **basics**: int, float, bool, byte, enum
- **references**: string, array, map, nil, error
- **functions**: extern, proc, method, test
- **conditions**: if, elif, else, match, break, continue
- **loops**: for
- **inbuilt**: echo, assert

## Specification

#### 1. Basic Types

```golang
const nerd: bool = false
const template: string = "hello world"
const age: int = 21
const pi: float = 3.14
const token: byte = 0x05
const data: string = nil
const validation: error = error("123")
```

#### 2. Reference Types

#### A. Enums

```pony
enum Direction
  | Left
  | Right
  | Up
  | Down
```

#### B. Types

```golang
type User
  name: string
  age: string
  gender: string
  address: string
  createdAt: string
  updatedAt: string
  deletedAt: string
```

#### C. Strings

```golang
const name := "Bob"
echo "Hello, $name!" // `$` is used for string interpolation

multiline = """
"""
```

#### D. Arrays

```golang
const nums := []int[1, 2, 3]
echo nums // ==> [1, 2, 3]
echo nums[1] // ==> 2
echo nums[10] // ==> nil

const names := []string["Krabs"]
names.add("Spongebob")
echo names.length // ==> 3
echo names.contains("Krabs") // ==> true
echo names.contains("Squidward") // ==> false
```

#### E. Maps

Keys will always be strings. Keys are ordered by default.

```golang
nums := map[int]{
    "one": 1,
    "two": 2
}
nums["three"] = 4
echo nums["one"]  // ==> "1"
echo nums["infinity"] // ==> nil
```

#### 3. Functions

```golang
type Operator proc(a int, b int) bool

proc add(a int, b int) =
  a + b

add(3, 5)
add(b: 4, a: 1)
```

#### 4. Conditions

#### A. if / elif / else

```python
a := 10
b := 20
c := a > b ? 10 : a;
if a < b:
  echo "$a < $b"
elif a > b:
  echo "$a > $b"
else:
  echo "$a == $b"
```

#### B. match

```pony
c := 4
match c:
| 1 => _resize(y, x)
| 2 => _processInput(c, x, y)
| 3 => _processInput(c, x, y)
| 4 => _processInput(c, x, y)
| 255 => _exit()
```

#### 5. Loops

#### A. Normal loops

```golang
for i := 0; i < 10; i++:
  echo i, v
```

#### B. Infinite Loop

```golang
for:
  echo "infinite"
```

#### C. Loop with Condition

```golang
for x > 5:
  echo x
```

## Array Iteration

```golang
const nums := [1, 2, 3]
for i, v  := range nums:
   echo i, v
```

## Map Iteration

```golang
const data := map[string]{"name": "krab"}
for k, v  := range data:
  echo k, v
```

#### 6. Sugar

```groovy
echo 1 + 2 + 3
// Output:
6

assert x == 2

// Output:
Assertion failed:
assert x == 2
      | |
      1 false
```

#### 7. Operators

#### A. Arithmetic Operators

- \+ (addition)
- \- (subtraction)
- \* (multiplication)
- / (division)
- % (modulo)

#### B. Logical Operators

- && (and)
- || (or)
- ! (not)

#### C. Comparison Operators

- \> (gt)
- <
- />=
- <=
- ==
- !=
- ^

#### D. Assignment Operators

- :=
- =
- +=
- -=
- \*=
- /=
- %=

#### 8. Doc Strings

```golang
"""
Takes a, b
"""
func add(a int, b int) {
}
```

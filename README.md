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

```lisp
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

	(var user1 "hello")
	(var user2 123.12)
	(var user3 144)

	(defn show/dir [dir]
		(html/div dir))

  (defn appun [f & ms]
    (when (every? some? ms)
      (apply f ms)))


```

#### 2. Reference Types

#### A. Enums

```pony
type Direction enum (
  Left
  | Right
)
```

#### B. Types

```golang
type User struct (
  name      string
  age       string
  gender    string
  address   string
  createdAt string
  updatedAt string
  deletedAt string
  onCreate  (a string)
  onUpdate  (b string)
  onDelete  (c string)
}
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
nums := {
    "one": 1,
    "two": 2
}
nums["three"] = 4
echo nums["one"]  // ==> "1"
echo nums["infinity"] // ==> nil
```

#### 3. Functions

```golang
type Operator (a int, b int) bool

const add = (a: int, b: int) => {
  return a + b
}
add(3, 5)
add(b: 4, a: 1)
```

#### 4. Conditions

#### A. if / elif / else

```python
a := 10
b := 20
c := a > b ? 10 : a;
if a < b {
  echo "$a < $b"
} elif a > b {
  echo "$a > $b"
} else {
  echo "$a == $b"
}
```

#### B. switch

#### 5. Loops

#### A. Normal loops

```golang
for i := 0; i < 10; i++ {
  echo i, v
}
```

#### B. Infinite Loop

```golang
for {
  echo "infinite"
}
```

#### C. Loop with Condition

```golang
for x > 5 {
  echo x
}
```

## Array Iteration

```golang
const nums := [1, 2, 3]
for i, v  := range nums {
   echo i, v
}
```

## Map Iteration

```golang
data := map[string]{"name": "krab"}
for k, v  := range data {
  echo k, v
}
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
const add = (a int, b int) => {
}
```

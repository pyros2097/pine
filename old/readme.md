# The Krab Programming Language

```go
const name = "Bob"
age := 20
const large_number = 9999999999
println(name)
println(age)
println(large_number)
```

2. Language Features

## Power Assert

```groovy
assert x == 2


// Output:
//
// Assertion failed:
// assert x == 2
//        | |
//        1 false
```

```go
func TestBlake(t *testing.T) {
    assert(x, 2)
}
```

## Power Mocks

## Power Stubs

## Null checking

All values are default nullable/nilable you can make them non-nullable by using an exclamation mark after the type;

```go
func area(a int!, b float) {
}
```

## Multiple returns and errors

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

### Krabby Web Framework

This component will work both in frontend and backend
Components can only be served from a backend using ssr or frontend
we will have c/go http server running which will convert the tree to
string of web components with the js code generated for each
SPA/SSR fully supported
GraphQL support

```pony
class Counter
    var count: I32
    return Row(
        flex: 0.5,
        chilren: Array(
            Text("123")
        ),
        Text("123")
    )
end
```

```golang
func useState<T>(T v) (T, func(T v)) {
}


func Counter(name: string) {
    count, setCount := krab.useState<int>(0)
    return Row{
        flex: 0.5,
        children: [
            Text{
                name: "${name} Count: ${count}"
            },
        ],
        Text{
            name: "${name} Count: ${count}"
        },
    }
}

func main() {
    render(Counter{
        name: "test"
    })
}
```

```golang
func Counter(name: string) {
    count, setCount := useState<int>(0)
    return Row(flex: 0.5,
        Text(name: "${name} Count: ${count}")
    )
}

func main() {
    render(Counter(name: "test"))
}
```

```golang
func Counter(name: string) {
    count, setCount := useState<int>(0)
    return (
        <Row flex=0.5>
            <Text name={"${name} Count: ${count}"} />
        </Row>
    )
}

func main() {
    render(<Counter name="test" />)
}
```

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

## Strings

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

## Arrays

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

## Maps

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

## Enums

TODO

## Structs

```golang
struct User {
    id string
    name string
    email string
}
```

## Funcs

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

## Loop Statements

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

```

```

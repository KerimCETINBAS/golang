#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_3) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_5)

# Printing & Formatting Strings

# Introduction to the fmt Package

The fmt package in Go provides functions for formatted I/O operations. It allows you to print or format strings, numbers, and other data types in a specific way.

## Printing Output

## `Print()`

The `Print()` function is used to print output to the standard output (usually the console) without any formatting. It takes any number of arguments and concatenates them into a single string before printing.

Example:

```go
package main

import "fmt"

func main() {
    fmt.Print("Hello, ")
    fmt.Print("World!")
}
```

output:

```
Hello, World!
```

## `Println()`

The `Println()` function works similar to Print(), but it adds a newline character (\n) after printing the arguments. This results in each call to Println() printing on a new line.

Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello,")
    fmt.Println("World!")
}
```

output:

```
Hello,
World!
```

## `Printf()`

The `Printf()` function is used for formatted printing. It allows you to specify a format string with placeholders (% verbs) that will be replaced by the corresponding values. The placeholders start with a % sign and are followed by a verb that represents the type of the value to be printed.

Example:

```go
package main

import "fmt"

func main() {
    name := "John"
    age := 25
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

Output:

```
Name: John, Age: 25
```

# Format Specifiers

Format specifiers define the expected type and format of the corresponding value when using formatted printing.

Common Format Specifiers:

- `%d` or `%v` - for integers
- `%f` - for floating-point numbers
- `%s` - for strings
- `%t` - for booleans
- `%c` - for characters
- `%p` - for pointers
- `%T` - for type of the value

Example:

```go
package main

import "fmt"

func main() {
    quantity := 10
    price := 19.99
    product := "Golang Course"
    fmt.Printf("Product: %s, Quantity: %d, Price: %.2f\n", product, quantity, price)
}
```

Output:

```
Product: Golang Course, Quantity: 10, Price: 19.99
```

# String Formatting

## `Sprintf()`

The `Sprintf()` function formats and returns a formatted string without printing it. It is similar to `Printf()`, but instead of printing the result, it returns the formatted string.
Example:

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 30
    formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age)
    fmt.Println(formattedString)
}
```

Output:

```
Name: Alice, Age: 30
```

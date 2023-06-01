#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_18) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_20)

&#10;

# Lesson 19

# The `strconv` Package in Go

In Go, the `strconv` package provides functions for converting strings to various data types and vice versa. This package is particularly useful when working with user inputs, file I/O, or network communication, where data conversions between strings and other types are often required. In this lesson, we will explore the `strconv` package and its commonly used functions.

## String to Basic Types Conversions

The `strconv` package provides functions to convert strings to basic types like integers, floats, and booleans. Here are some commonly used functions:

- `Atoi`: Converts a string to an `int` value.
- `ParseFloat`: Converts a string to a `float64` value.
- `ParseBool`: Converts a string to a `bool` value.

Here's an example that demonstrates the usage of these functions:

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // String to int
    numStr := "42"
    num, _ := strconv.Atoi(numStr)
    fmt.Println(num) // Output: 42

    // String to float64
    floatStr := "3.14"
    floatVal, _ := strconv.ParseFloat(floatStr, 64)
    fmt.Println(floatVal) // Output: 3.14

    // String to bool
    boolStr := "true"
    boolVal, _ := strconv.ParseBool(boolStr)
    fmt.Println(boolVal) // Output: true
}
```

In this example, we convert strings to different basic types using the `Atoi`, `ParseFloat`, and `ParseBool` functions from the `strconv` package. The second return value from these functions is an error, which is typically handled using the blank identifier `_` for brevity.

## Basic Types to String Conversions

The `strconv` package also provides functions to convert basic types to strings. Here are some commonly used functions:

- `Itoa`: Converts an `int` value to a string.
- `FormatFloat`: Converts a `float64` value to a string.
- `FormatBool`: Converts a `bool` value to a string.

Here's an example:

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // int to string
    num := 42
    numStr := strconv.Itoa(num)
    fmt.Println(numStr) // Output: "42"

    // float64 to string
    floatVal := 3.14
    floatStr := strconv.FormatFloat(floatVal, 'f', 2, 64)
    fmt.Println(floatStr) // Output: "3.14"

    // bool to string
    boolVal := true
    boolStr := strconv.FormatBool(boolVal)
    fmt.Println(boolStr) // Output: "true"
}
```

In this example, we use the `Itoa`, `FormatFloat`, and `FormatBool` functions to convert different basic types to strings.

## Handling Errors

When using the `strconv` package, it's important to handle potential errors that may occur during conversions. All conversion functions in `strconv` return an error as the second return value. It's good practice to check the error and handle it appropriately to avoid unexpected behavior or crashes in your program.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "123abc"
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(num)
}
```

In this

example, the `Atoi` function is called with a string that cannot be converted to an integer. The error is captured in the `err` variable, and if it's not `nil`, an error message is printed.

## Conclusion

The `strconv` package in Go provides convenient functions for converting strings to basic types and vice versa. By using functions like `Atoi`, `ParseFloat`, and `ParseBool`, you can convert strings to integers, floats, and booleans. Similarly, functions like `Itoa`, `FormatFloat`, and `FormatBool` allow you to convert basic types to strings. Remember to handle errors appropriately when working with the `strconv` package to ensure the reliability of your code.

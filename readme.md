#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_9) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_11)

&#10;

# Lesson 10

# Functions in Go

Functions are an essential building block of any programming language. They allow you to encapsulate reusable code and perform specific tasks. In Go, functions are defined using the `func` keyword. Let's dive into the details of creating and using functions in Go.

## Function Declaration and Syntax

A function declaration in Go consists of the function name, a list of parameters (if any), an optional return type, and the function body enclosed in curly braces. The basic syntax of a function declaration in Go is as follows:

```go
func functionName(parameter1 type, parameter2 type) returnType {
    // Function body
    // Code to be executed
    // Return statement (if required)
}
```

- `func`: The keyword used to declare a function.
- `functionName`: The name of the function.
- `parameter1, parameter2`: The input parameters of the function, each with its corresponding type.
- `returnType`: The data type that the function returns (if any).
- `return` statement: Used to return a value from the function (if required).

## Function Examples

Let's take a look at some examples to understand the concept of functions in Go.

### Example 1: Simple Function

```go
func greet() {
    fmt.Println("Hello, World!")
}
```

In this example, we have defined a function named `greet` that does not take any parameters and does not return any value. It simply prints "Hello, World!" to the console.

### Example 2: Function with Parameters

```go
func add(a, b int) int {
    return a + b
}
```

This example demonstrates a function named `add` that takes two integer parameters (`a` and `b`) and returns their sum as an integer.

### Example 3: Function with Multiple Return Values

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

Here, the `divide` function takes two integer parameters (`a` and `b`) and returns their division as the first return value. In case of division by zero, it returns an error as the second return value using the `fmt.Errorf` function.

## Calling Functions

To execute a function, you need to call it by using its name followed by parentheses. If the function has parameters, you pass the corresponding values within the parentheses. If the function has a return value, you can store it in a variable.

```go
greet() // Call the greet function

sum := add(3, 4) // Call the add function and store the return value in the sum variable

result, err := divide(10, 2) // Call the divide function and store the return values in result and err variables
```

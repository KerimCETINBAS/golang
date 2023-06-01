#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_7) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_9)

&#10;

# Lesson 8

# Booleans, Conditionals, and Logical Operators

## Booleans in Go

In Go, the `bool` type represents boolean values, which can be either `true` or `false`. Booleans are commonly used in conditional statements and logical operations.

```go
var isTrue bool = true
var isFalse bool = false
```

## Conditional Statements

Conditional statements allow you to control the flow of your program based on certain conditions. In Go, the primary conditional statement is the `if` statement.

### If Statement

The `if` statement executes a block of code if a given condition is `true`. The general syntax of the `if` statement is as follows:

```go
if condition {
    // code to be executed if the condition is true
}
```

Example:

```go
var age int = 20

if age >= 18 {
    fmt.Println("You are an adult.")
}
```

### If-Else Statement

The `if-else` statement allows you to specify alternative code to be executed if the condition in the `if` statement is `false`. The general syntax of the `if-else` statement is as follows:

```go
if condition {
    // code to be executed if the condition is true
} else {
    // code to be executed if the condition is false
}
```

Example:

```go
var age int = 15

if age >= 18 {
    fmt.Println("You are an adult.")
} else {
    fmt.Println("You are not yet an adult.")
}
```

### If-Else If-Else Statement

The `if-else if-else` statement allows you to test multiple conditions and execute different code blocks accordingly. The general syntax of the `if-else if-else` statement is as follows:

```go
if condition1 {
    // code to be executed if condition1 is true
} else if condition2 {
    // code to be executed if condition2 is true
} else {
    // code to be executed if all conditions are false
}
```

Example:

```go
var num int = 5

if num > 0 {
    fmt.Println("The number is positive.")
} else if num < 0 {
    fmt.Println("The number is negative.")
} else {
    fmt.Println("The number is zero.")
}
```

## Logical Operators

Logical operators allow you to combine multiple conditions and perform logical operations. In Go, the logical operators are `&&` (AND), `||` (OR), and `!` (NOT).

### AND Operator (&&)

The `&&` operator returns `true` if both conditions are `true`, otherwise it returns `false`.

Example:

```go
var num int = 10

if num > 0 && num < 100 {
    fmt.Println("The number is between 0 and 100.")
}
```

### OR Operator (||)

The `||` operator returns `true` if at least one of the conditions is `true`, otherwise it returns `false`.

Example:

```go
var age int = 15

if age < 18 || age >= 65 {
    fmt.Println("You are either a minor or a senior citizen.")
}
```

### NOT Operator (!)

The `!` operator negates the value of a condition. If the condition is `true`, it becomes `false`, and vice versa.

Example:

```go
var isSunny bool = false

if !isSunny {
    fmt.Println("It's not sunny today.")
}
```

#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_8) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_10)

&#10;

# Lesson 9

# Arithmetic Operators in Go

Arithmetic operators in Go allow you to perform basic mathematical operations on numeric values. Go supports the following arithmetic operators:

- Addition (`+`)
- Subtraction (`-`)
- Multiplication (`*`)
- Division (`/`)
- Remainder or Modulo (`%`)
- Increment (`++`)
- Decrement (`--`)

## Addition (+)

The addition operator (`+`) is used to add two numeric values together.

```go
var a int = 5
var b int = 3
var sum int = a + b
```

## Subtraction (-)

The subtraction operator (`-`) is used to subtract one numeric value from another.

```go
var a int = 10
var b int = 7
var difference int = a - b
```

## Multiplication (\*)

The multiplication operator (`*`) is used to multiply two numeric values.

```go
var a int = 4
var b int = 6
var product int = a * b
```

## Division (/)

The division operator (`/`) is used to divide one numeric value by another.

```go
var a float64 = 15
var b float64 = 4
var quotient float64 = a / b
```

## Remainder or Modulo (%)

The remainder or modulo operator (`%`) is used to find the remainder after dividing one numeric value by another.

```go
var a int = 10
var b int = 3
var remainder int = a % b
```

## Increment (++) and Decrement (--)

The increment (`++`) and decrement (`--`) operators are used to increase or decrease the value of a variable by 1, respectively.

```go
var count int = 5
count++ // Increment count by 1
count-- // Decrement count by 1
```

It's important to note that increment and decrement operators can be used in both pre-increment/decrement and post-increment/decrement forms.

```go
var count int = 5
var result int

result = ++count // Pre-increment - increment count and assign the incremented value to result
result = count++ // Post-increment - assign the current value of count to result and then increment count
```

These are the basic arithmetic operators in Go. You can use them to perform mathematical calculations in your programs.

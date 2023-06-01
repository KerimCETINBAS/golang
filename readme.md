#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_17) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_19)

&#10;

# Lesson 18

# Switch Statements in Go

In Go, the switch statement provides a convenient way to execute different code blocks based on the value of an expression. Switch statements are a powerful tool for handling multiple cases and can be used with various data types. In this lesson, we will explore the syntax and usage of switch statements in Go.

## Basic Switch Statement

The basic syntax of a switch statement in Go is as follows:

```go
switch expression {
case value1:
    // code block executed when expression matches value1
case value2:
    // code block executed when expression matches value2
default:
    // code block executed when expression doesn't match any case
}
```

Here's an example of a basic switch statement:

```go
package main

import "fmt"

func main() {
    day := "Monday"

    switch day {
    case "Monday":
        fmt.Println("It's Monday!")
    case "Tuesday":
        fmt.Println("It's Tuesday!")
    case "Wednesday":
        fmt.Println("It's Wednesday!")
    default:
        fmt.Println("It's another day!")
    }
}
```

In this example, the value of the `day` variable is matched against the cases in the switch statement. If a match is found, the corresponding code block is executed. If no match is found, the code block under the `default` case is executed.

## Switch with Multiple Expressions

Go allows you to use multiple expressions in a switch statement. This can be useful when you want to compare multiple values in a single switch statement. Here's an example:

```go
package main

import "fmt"

func main() {
    language := "Go"
    version := 1

    switch language {
    case "Go", "golang":
        fmt.Println("You are using Go!")
    default:
        fmt.Println("You are using a different language.")
    }

    switch {
    case version < 1:
        fmt.Println("You are using an old version.")
    case version == 1:
        fmt.Println("You are using the latest version.")
    default:
        fmt.Println("You are using a future version.")
    }
}
```

In the first switch statement, we compare the `language` variable against multiple values, "Go" and "golang". If the value matches any of the cases, the corresponding code block is executed.

In the second switch statement, we omit the expression. Instead, we use Boolean expressions in the cases. This allows us to check different conditions and execute the corresponding code block when a condition is true.

## Fallthrough

By default, Go's switch statements automatically break after a case is executed. However, you can use the `fallthrough` keyword to execute the next case's code block as well. Here's an example:

```go
package main

import "fmt"

func main() {
    num := 2

    switch num {
    case 1:
        fmt.Println("The number is 1.")
        fallthrough
    case 2:
        fmt.Println("The number is 2.")
        fallthrough
    case 3:
        fmt.Println("The number is 3.")
    }
}
```

In this example, when `num` is 2, the code block for case 2 is executed, and the `fallthrough` statement transfers control to the next case, which is case 3. As a result, both the code blocks for case 2 and case 3 are executed.

## Conclusion

Switch statements in Go provide a concise and efficient way to handle multiple cases based on the value of an expression. By utilizing switch statements, you can simplify code that would otherwise require multiple if-

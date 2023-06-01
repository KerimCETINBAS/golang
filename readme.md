#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_16) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_18)

&#10;

# Lesson 17

# Reading User Inputs in Go

Reading user inputs is a common task in many programs. In Go, you can read user inputs from the command line using the `bufio` package and the `os` package. These packages provide convenient methods to interact with standard input and retrieve user inputs. In this lesson, we will explore how to read user inputs in Go using `bufio` and `os`.

## Reading User Inputs with bufio

The `bufio` package provides a `Scanner` type that allows you to read input from various sources, including standard input (`os.Stdin`). Here's an example of how to read a line of text entered by the user:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Enter your name: ")
    scanner.Scan() // Read the user input

    name := scanner.Text() // Get the text entered by the user

    fmt.Println("Hello,", name)
}
```

In this example, we create a new `Scanner` using `bufio.NewScanner()` and pass `os.Stdin` as the input source. We then prompt the user to enter their name using `fmt.Print()` and read their input using `scanner.Scan()`. The user's input is stored in the `Text()` method of the `scanner`.

## Reading User Inputs with os

Alternatively, you can use the `os` package to read user inputs directly. Here's an example:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Print("Enter your name: ")

    var name string
    fmt.Scanln(&name)

    fmt.Println("Hello,", name)
}
```

In this example, we use `fmt.Scanln()` to read a line of text entered by the user. The `&` symbol before `name` is used to pass the memory address of the `name` variable to the `Scanln()` function, allowing it to store the user's input directly into `name`.

## Handling Errors

When reading user inputs, it's important to handle potential errors. Both the `bufio.Scanner` and `fmt.Scanln()` functions return an error that indicates if the reading operation was successful or not. It's good practice to check for and handle errors appropriately in your code.

## Conclusion

Reading user inputs in Go is straightforward using the `bufio` package and the `os` package. By utilizing the `Scanner` type from `bufio` or the `Scanln()` function from `fmt`, you can prompt users for inputs and retrieve their responses. Remember to handle errors appropriately to ensure robustness in your code.

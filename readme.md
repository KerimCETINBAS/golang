#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_10) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_12)

&#10;

# Lesson 11

# Package Scopes in Go

In Go, the concept of package scope refers to the visibility and accessibility of variables, constants, functions, and types within a package. Understanding package scopes is important for organizing and accessing code elements effectively. Let's explore the different aspects of package scopes in Go.

## Default Package Scope

By default, identifiers declared at the package level have **package scope**, which means they are accessible throughout the package.

```go
package main

import "fmt"

// Variables with package scope
var (
    name   string = "John"
    age    int    = 30
    height float64
)

// Function with package scope
func sayHello() {
    fmt.Printf("Hello, %s!\n", name)
}

func main() {
    fmt.Println(name)     // Accessing a variable with package scope
    fmt.Println(age)      // Accessing another variable with package scope
    sayHello()             // Calling a function with package scope
}
```

In the example above, the variables `name`, `age`, and `height` have package scope, which means they can be accessed and modified within any function in the package. Similarly, the `sayHello` function has package scope and can be called from the `main` function.

## Block Scope

Go also supports block-level scope, which defines the visibility and accessibility of identifiers within a specific block of code (e.g., inside a function or an if statement).

```go
package main

import "fmt"

func main() {
    // Block scope
    if true {
        x := 10
        fmt.Println(x) // Accessible within the block
    }

    // fmt.Println(x) // Error: x is not accessible outside the block
}
```

In the example above, the variable `x` has block scope. It is declared and assigned a value within the if statement block and is accessible only within that block. Attempting to access `x` outside the block will result in a compilation error.

## Package-level Naming Conventions

In Go, the naming convention for identifiers with package scope follows the principle of **exported** and **unexported** names. An identifier is exported if its name starts with an uppercase letter. Exported names are accessible from other packages, while unexported names are only accessible within the package.

```go
package mypackage

var unexportedName string = "Unexported"   // Unexported variable
var ExportedName string = "Exported"       // Exported variable

func unexportedFunc() {
    // Unexported function
}

func ExportedFunc() {
    // Exported function
}
```

In the example above, the variable `unexportedName` and the function `unexportedFunc` have unexported names and can only be accessed within the `mypackage` package. On the other hand, the variable `ExportedName` and the function `ExportedFunc` have exported names and can be accessed from other packages as well.

## Conclusion

By utilizing package scope and following naming conventions, you can control the visibility and accessibility of variables, constants, functions, and types within your package. This helps maintain code clarity and encapsulation, and promotes reusability and modularity in your Go programs.

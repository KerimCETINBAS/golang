#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_20) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_22)

&#10;

# Lesson 21

# Interfaces in Go

In Go, interfaces provide a way to define behavior or functionality that a type must implement. They enable you to write more flexible and modular code by defining a set of methods that a type should have. In this lesson, we will explore interfaces in Go and learn how to define, implement, and use them effectively.

## Interface Basics

An interface in Go is a collection of method signatures. It specifies a contract that a type must adhere to. Any type that implements all the methods defined in an interface is said to satisfy or implement that interface.

The basic syntax to define an interface in Go is as follows:

```go
type MyInterface interface {
    Method1() // Method signature without implementation
    Method2() // Method signature without implementation
    // ...
}
```

Here's an example of an interface named `Animal` with a single method called `MakeSound`:

```go
type Animal interface {
    MakeSound()
}
```

Any type that implements the `MakeSound` method will implicitly satisfy the `Animal` interface.

## Implementing an Interface

To implement an interface in Go, a type must define all the methods specified by the interface. Let's see an example:

```go
type Dog struct {
    Name string
}

func (d Dog) MakeSound() {
    fmt.Println(d.Name, "says woof!")
}
```

In this example, we define a type `Dog` with a `MakeSound` method that satisfies the `Animal` interface. The method implementation specifies how a `Dog` makes a sound.

## Using Interfaces

Interfaces are used to achieve polymorphism in Go. By defining functions that accept interfaces as parameters, you can write code that operates on any type implementing that interface. Here's an example:

```go
func PlaySound(a Animal) {
    a.MakeSound()
}

func main() {
    dog := Dog{Name: "Buddy"}
    PlaySound(dog)
}
```

In this example, we define a `PlaySound` function that takes an `Animal` interface as a parameter. We can pass any type implementing the `Animal` interface to this function. In the `main` function, we create a `Dog` instance and pass it to `PlaySound`. Since `Dog` satisfies the `Animal` interface, the `MakeSound` method of the `Dog` type is called.

## Type Assertion and Type Switch

Interfaces in Go support type assertion and type switch, which allow you to work with underlying concrete types. Here's an example:

```go
func PrintType(v interface{}) {
    switch val := v.(type) {
    case string:
        fmt.Println("Type: string, Value:", val)
    case int:
        fmt.Println("Type: int, Value:", val)
    default:
        fmt.Println("Unknown type!")
    }
}

func main() {
    PrintType("Hello") // Type: string, Value: Hello
    PrintType(42)     // Type: int, Value: 42
    PrintType(true)   // Unknown type!
}
```

In this example, the `PrintType` function takes an empty interface as a parameter. Inside the function, we use a type switch to check the underlying type of the value passed. Depending on the type, we perform different actions.

## Empty Interface

The empty interface `interface{}` is a special type in Go that can hold values of any type. It is used when you want to work with values of unknown types or when you need maximum flexibility. Here's an example:

```go
func PrintValue(v interface{}) {
    fmt.Println("Value:", v)
}

func main() {
    PrintValue("
```

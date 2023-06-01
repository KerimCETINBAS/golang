#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_14) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_16)

&#10;

# Lesson 15

# Structs & Custom Types in Go

In Go, structs and custom types allow you to define and work with your own data types. Structs provide a way to group related data together, while custom types enable you to create new type names based on existing types. In this lesson, we will explore how to define and use structs, as well as create custom types in Go.

## Structs

A struct is a composite data type that allows you to group together values with different types into a single entity. Each value within a struct is called a field. Structs provide a convenient way to represent real-world objects or data structures.

```go
type Person struct {
    name string
    age  int
}
```

In the example above, we define a struct type named `Person` with two fields: `name` of type `string` and `age` of type `int`. You can create instances of the struct by declaring variables of type `Person`:

```go
var p Person
p.name = "John Doe"
p.age = 30
```

To access the fields of a struct, you use the dot (`.`) operator:

```go
fmt.Println(p.name) // Output: John Doe
fmt.Println(p.age)  // Output: 30
```

## Custom Types

Go allows you to create custom types by defining a new name for an existing type. This can be useful for providing more descriptive names or adding additional behavior to existing types.

```go
type ID string
```

In the example above, we define a custom type `ID` based on the underlying type `string`. With this custom type, we can create variables of type `ID`:

```go
var id ID = "abc123"
```

Custom types provide type safety and allow you to define specific behaviors for the new type.

## Conclusion

Structs and custom types in Go provide a powerful way to define and work with your own data types. Structs allow you to group related data together

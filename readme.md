#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_13) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_15)

&#10;

# Lesson 14

# Pointers in Go

Pointers in Go that allow you to work with the memory address of variables. By using pointers, you can directly access and modify the value stored in memory, which can be useful in scenarios such as passing large data structures efficiently or modifying variables within functions. In this lesson, we will explore the concept of pointers in Go and how to work with them effectively.

## Understanding Pointers

In Go, a pointer is a variable that holds the memory address of another variable. Pointers are represented using the `*` (asterisk) symbol.

```go
var ptr *int
```

In the example above, `ptr` is a pointer to an integer. It can store the memory address of an `int` variable.

To obtain the memory address of a variable, you can use the `&` (ampersand) operator.

```go
var value int = 42
ptr = &value
```

In this example, the `&value` expression gives us the memory address of the `value` variable, which is then assigned to the `ptr` pointer.

## Dereferencing Pointers

Dereferencing a pointer means accessing the value stored at the memory address it points to. It is done using the `*` (asterisk) operator.

```go
var value int = 42
var ptr *int = &value
fmt.Println(*ptr) // Output: 42
```

In the example above, `*ptr` gives us access to the value stored at the memory address pointed to by `ptr`, which is `42`.

## Modifying Values through Pointers

One of the main benefits of pointers is the ability to modify values directly in memory. By dereferencing a pointer, you can change the value stored at the memory address.

```go
var value int = 42
var ptr *int = &value
*ptr = 100
fmt.Println(value) // Output: 100
```

In this example, `*ptr = 100` modifies the value stored at the memory address pointed to by `ptr`. As a result, the value of `value` is updated to `100`.

## Pointer Parameters in Functions

Passing pointers as function parameters allows you to modify variables in the calling code.

```go
func modifyValue(ptr *int) {
    *ptr = 100
}

func main() {
    var value int = 42
    modifyValue(&value)
    fmt.Println(value) // Output: 100
}
```

In this example, the `modifyValue` function takes a pointer to an integer as a parameter. By dereferencing the pointer inside the function and assigning a new value to `*ptr`, we modify the original `value` variable in the `main` function.

## Returning Pointers from Functions

Functions can also return pointers, allowing you to allocate memory dynamically and return the memory address.

```go
func createValue() *int {
    value := 42
    return &value
}

func main() {
    ptr := createValue()
    fmt.Println(*ptr) // Output: 42
}
```

In this example, the `createValue` function allocates memory for an `int` variable and returns its memory address using `&value`. The `main` function receives the pointer and dereferences it to access the value stored in memory.

## Nil Pointers

In Go, a pointer can have a special value called `nil`, which represents the absence of a valid memory address. A `nil` pointer does not point to any memory location.

```go
var ptr *int = nil
```

When

working with pointers, it's important to check for `nil` before dereferencing the pointer to avoid runtime errors.

## Conclusion

Pointers are a powerful feature in Go that allow you to work with the memory address of variables. By understanding how to declare, assign, dereference, and modify values through pointers, you can leverage their capabilities to efficiently manage memory and modify variables directly. Proper usage of pointers can lead to more efficient and flexible code in your Go programs.

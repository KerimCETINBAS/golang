#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_12) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_14)

&#10;

# Lesson 13

# Pass by Value, Non-Pointer Values, and Pointer Wrapper Values in Go

In Go, function arguments are passed by value by default, meaning that a copy of the argument value is passed to the function. However, there are cases where you might want to modify the original value or avoid unnecessary copying. In this lesson, we will explore the concepts of pass by value, non-pointer values, and pointer wrapper values in Go.

## Pass by Value

In Go, when you pass a variable as an argument to a function, the function receives a copy of the variable's value. Any modifications made to the parameter inside the function will not affect the original variable.

```go
package main

import "fmt"

func modifyValue(x int) {
    x = 42 // Modify the parameter
}

func main() {
    value := 10
    modifyValue(value)
    fmt.Println(value) // Output: 10 (unchanged)
}
```

In the example above, the `modifyValue` function receives a copy of the `value` variable. Although the function modifies the parameter `x`, the original `value` variable remains unchanged.

## Non-Pointer Values

To modify the original value inside a function, you can pass the argument as a pointer type. However, Go also provides a convenient alternative for non-pointer values using the `&` operator and the `*` operator in function parameter and argument declarations.

```go
package main

import "fmt"

func modifyValue(ptr *int) {
    *ptr = 42 // Modify the value pointed to by ptr
}

func main() {
    value := 10
    modifyValue(&value) // Pass the address of value
    fmt.Println(value) // Output: 42 (modified)
}
```

In the example above, we pass the address of the `value` variable (`&value`) as an argument to the `modifyValue` function. Inside the function, the parameter `ptr` is a pointer to an integer. By dereferencing `ptr` using the `*` operator, we can modify the value stored at the address.

## Pointer Wrapper Values

Go also provides a convenient way to work with non-pointer values using pointer wrapper types, such as `new`, `make`, and `append`. These functions return pointer values to the allocated memory, allowing modifications to be made to the original value.

```go
package main

import "fmt"

func modifySlice(slice []int) {
    slice[0] = 42 // Modify the element in the slice
}

func main() {
    slice := make([]int, 3)
    modifySlice(slice) // Pass the slice by value (pointer wrapper)
    fmt.Println(slice) // Output: [42 0 0] (modified)
}
```

In the example above, the `make` function is used to create a slice. The `modifySlice` function receives the slice by value (pointer wrapper). Since the underlying data of the slice is stored in memory and accessed through a pointer, modifications made to the elements inside the function are reflected in the original slice.

## Conclusion

Understanding pass by value, non-pointer values, and pointer wrapper values is crucial when working with functions and modifying values in Go. By default, Go passes arguments by value, meaning modifications made inside the function won't affect the original value. However, by using pointers or pointer wrapper types, you can modify the original value directly or avoid unnecessary copying. Being mindful of these concepts will help you write more efficient and effective Go code.

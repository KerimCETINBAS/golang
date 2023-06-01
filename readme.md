#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_15) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_17)

&#10;

# Lesson 16

# Receiver Functions with Pointers in Go

In Go, receiver functions allow you to associate methods with a particular type. By using pointers as receivers, you can modify the state of the receiver directly. This can be particularly useful when working with large data structures or when you want to modify the original value. In this lesson, we will explore the concept of receiver functions with pointers in Go.

## Pointer Receivers

In Go, you can define a method on a type by specifying the type as the receiver of the function. When using a pointer as the receiver, the method can modify the underlying value directly. To use a pointer receiver, you specify the type with an asterisk (`*`).

```go
type Rectangle struct {
    width  float64
    height float64
}

func (r *Rectangle) Scale(factor float64) {
    r.width *= factor
    r.height *= factor
}
```

In the example above, we define a method `Scale` on the `Rectangle` type with a pointer receiver (`*Rectangle`). By dereferencing the pointer (`r.width` and `r.height`), we can directly modify the underlying values of the `Rectangle` struct.

## Calling Pointer Receiver Functions

To call a pointer receiver function, you can use either a pointer to the type or a value of the type. Go automatically takes care of converting the value to a pointer when calling the method.

```go
r := Rectangle{width: 10, height: 5}
r.Scale(2)
fmt.Println(r.width, r.height) // Output: 20 10

ptr := &r
ptr.Scale(0.5)
fmt.Println(r.width, r.height) // Output: 10 5
```

In the example above, we create a `Rectangle` instance `r` and call the `Scale` method using both `r` and `ptr`, which is a pointer to `r`. In both cases, the method modifies the underlying values of `r`.

## Benefits of Pointer Receivers

Using pointer receivers in receiver functions has a few advantages:

- **Efficiency**: When working with large data structures, using a pointer receiver avoids copying the entire value to the method, making it more efficient in terms of memory and performance.

- **Mutability**: Pointer receivers allow you to modify the receiver directly, which can be useful when you want to change the original value.

- **Consistency**: Using pointer receivers consistently across methods of a type helps maintain consistency and avoids potential bugs related to inconsistent behavior.

## When to Use Pointer Receivers

Use pointer receivers when:

- You want to modify the underlying value of a type directly within the method.
- You're working with large data structures to avoid unnecessary copying.
- You want to ensure consistency when modifying the state of a type across different methods.

However, if modifying the underlying value is not necessary or desired, you can use value receivers instead.

## Conclusion

Receiver functions with pointers in Go provide a way to modify the underlying value directly within a method. By using pointer receivers, you can efficiently work with large data structures and ensure consistency when modifying the state of a type. Understanding when and how to use pointer receivers will help you write more efficient and effective Go code.

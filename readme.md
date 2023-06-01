#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_6) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_8)

&#10;

# Lesson 6

# Go Loop Constructs

In Go, there are several loop constructs available to iterate over collections or repeat a block of code. The commonly used loop constructs are `for`, `for-in`, `while`, and `do-while`. Let's explore each of them:

## 1. For Loop

The `for` loop is the most common and versatile loop construct in Go. It allows you to execute a block of code repeatedly until a certain condition is met. The general syntax of the `for` loop in Go is as follows:

```go
for initialization; condition; post {
    // code to be executed
}
```

- `initialization` is an optional statement that initializes variables or assigns values before the loop starts.
- `condition` is a Boolean expression that is evaluated before each iteration. If the condition is `true`, the loop continues; otherwise, it terminates.
- `post` is an optional statement that is executed at the end of each iteration.

Example:

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Output:

```
0
1
2
3
4
```

## 2. For-in Loop

The `for-in` loop is used to iterate over collections such as arrays, slices, or maps. It automatically iterates through the elements of the collection without the need for an index variable. The general syntax of the `for-in` loop in Go is as follows:

```go
for _, element := range collection {
    // code to be executed
}
```

- `_` is used to ignore the index if you don't need it.
- `element` represents the current element of the collection.
- `range collection` specifies the collection to iterate over.

Example:

```go
numbers := []int{1, 2, 3, 4, 5}
for _, num := range numbers {
    fmt.Println(num)
}
```

Output:

```
1
2
3
4
5
```

## 3. While Loop

Go does not have a built-in `while` loop construct, but you can simulate it using the `for` loop. The `while` loop continues executing the block of code as long as the specified condition is true. The general syntax of the `while` loop in Go is as follows:

```go
for condition {
    // code to be executed
}
```

Example:

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

Output:

```
0
1
2
3
4
```

## 4. Do-While Loop

Go also does not have a built-in `do-while` loop construct, but you can simulate it using the `for` loop. The `do-while` loop executes the block of code at least once and continues as long as the specified condition is true. The general syntax of the `do-while` loop in Go is as follows:

```go
for {
    // code to be executed
    if !condition {
        break
    }
}
```

Example:

```go
i := 0
for {
    fmt.Println(i)
    i++
    if i >= 5 {
        break
    }
}
```

Output:

```
0
1
2
3
4
```

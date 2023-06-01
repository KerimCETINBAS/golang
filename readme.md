#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_11) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_13)

&#10;

# Lesson 12

# Maps in Go

Maps are a built-in data structure in Go that allow you to store key-value pairs. They provide an efficient way to look up values based on a specific key. In this lesson, we will explore the concept of maps and how to work with them in Go.

## Creating a Map

To create a map in Go, you use the `make` function with the `map` keyword, followed by the key and value types in square brackets.

```go
// Creating an empty map
var m map[keyType]valueType
```

For example, to create a map that stores string keys and int values:

```go
var m map[string]int
```

Alternatively, you can use the shorthand syntax to create and initialize a map:

```go
m := make(map[keyType]valueType)
```

## Adding and Accessing Elements

To add elements to a map, you assign a value to a specific key using the square bracket notation.

```go
m[key] = value
```

For example:

```go
m := make(map[string]int)
m["apple"] = 5
m["banana"] = 3
```

To access a value from the map, you use the square bracket notation with the key.

```go
value := m[key]
```

For example:

```go
fmt.Println(m["apple"]) // Output: 5
```

## Checking if a Key Exists

You can check if a key exists in a map using the following syntax:

```go
value, ok := m[key]
```

The `ok` variable will be `true` if the key exists in the map, and `false` otherwise. You can use this information to perform conditional logic based on key existence.

```go
value, ok := m["apple"]
if ok {
    // Key exists
    fmt.Println(value)
} else {
    // Key does not exist
    fmt.Println("Key not found")
}
```

## Deleting an Element

To delete an element from a map, you use the `delete` function and pass in the map and the key to be deleted.

```go
delete(m, key)
```

For example:

```go
delete(m, "apple")
```

## Iterating Over a Map

You can iterate over the elements of a map using a `for range` loop. It returns both the key and the value of each element.

```go
for key, value := range m {
    // Do something with key and value
    fmt.Println(key, value)
}
```

## Map Length

To get the number of elements in a map, you use the `len` function.

```go
length := len(m)
```

For example:

```go
fmt.Println(len(m))
```

## Conclusion

Maps are a powerful data structure in Go that allow you to store key-value pairs. They provide efficient look-up capabilities based on keys.

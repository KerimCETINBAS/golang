#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_4) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_6)

&#10;
# lesson 5
# Arrays, Slices, `len()`, and `append()` in Go

## Introduction

In Go, arrays and slices are fundamental data structures used to store collections of elements. Arrays have a fixed length, while slices are more flexible and can grow or shrink dynamically. The built-in functions `len()` and `append()` are often used with arrays and slices to obtain their length and modify them, respectively. This lesson will provide an overview of arrays, slices, and how to work with `len()` and `append()`.

## I. Arrays

1. Definition:

   - An array is a fixed-size sequence of elements of the same type.
   - Declaring an array: `var arr [size]Type`

2. Accessing Elements:

   - Elements in an array are accessed using zero-based indexing: `arr[index]`.
   - Example: `arr[0]` gives the first element of the array.

3. Array Initialization:
   - Initializing an array with values: `arr := [size]Type{val1, val2, ...}`
   - Example: `arr := [3]int{1, 2, 3}` initializes an integer array with three elements.

## II. Slices

1. Definition:

   - A slice is a dynamic, variable-length sequence built on top of an array.
   - Slices are more commonly used than arrays in Go due to their flexibility.

2. Slice Initialization:

   - Initializing a slice from an array: `slice := arr[start:end]`
   - Example: `slice := arr[1:4]` creates a slice from the array `arr` starting from index 1 and ending at index 3 (exclusive).

3. Modifying Slices:

   - Slices are reference types, so modifying a slice also modifies the underlying array.
   - Example: `slice[0] = 42` changes the first element of the slice and modifies the original array.

4. Dynamic Slicing:
   - Omitting the start or end index when creating a slice defaults to the beginning or end of the array, respectively.
   - Example: `slice := arr[:3]` creates a slice containing the first three elements of the array.

## III. `len()` Function

1. Definition:
   - The `len()` function returns the length of an array, slice, or a string.
   - Example: `length := len(slice)` assigns the length of the `slice` to the variable `length`.

## IV. `append()` Function

1. Definition:

   - The `append()` function is used to add elements to a slice.
   - It returns a new slice with the appended elements.
   - Example: `newSlice := append(slice, element1, element2, ...)` appends elements to `slice` and assigns the result to `newSlice`.

2. Modifying Slices with `append()`:
   - If the capacity of the underlying array is exceeded, `append()` creates a new larger array and copies the elements over.
   - Example: `slice = append(slice, newElement)` appends `newElement` to `slice`.

Conclusion:
Arrays have a fixed size, while slices are dynamic and allow for more flexible operations. The `len()` function is used to determine the length of arrays, slices, or strings, while `append()` is used to add elements to slices.

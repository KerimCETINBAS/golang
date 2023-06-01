#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_5) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_7)

&#10;

# Lesson 6

# Standard Library, Strings, and Sort Package in Go

## Introduction

The standard library in Go provides a rich set of packages that offer various functionalities for common programming tasks. Two important packages in the standard library are `strings` and `sort`. The `strings` package provides functions to manipulate strings, while the `sort` package offers sorting algorithms for different data types. This lesson will provide an overview of the standard library, focusing on the `strings` and `sort` packages.

## I. Standard Library

1. Definition

   - The standard library in Go is a collection of packages that are included with the Go programming language.
   - These packages offer a wide range of functionalities, making it easier to develop applications in Go.

2. Importing Packages

   - To use the functionalities of a package, it must be imported using the `import` keyword.
   - Example: `import "fmt"` imports the `fmt` package for formatted input/output.

3. Commonly Used Packages
   - The standard library includes packages for file I/O, networking, concurrency, cryptography, and more.
   - Examples: `fmt`, `os`, `net`, `sync`, `crypto`, etc.

## II. The `strings` Package

1. Introduction

   - The `strings` package provides functions for manipulating and working with strings.

2. Commonly Used Functions
   - `strings.Contains(s, substr)` checks if a string `s` contains a substring `substr`.
   - `strings.HasPrefix(s, prefix)` checks if a string `s` starts with a given `prefix`.
   - `strings.Join(strs, sep)` concatenates a slice of strings `strs` into a single string, separated by `sep`.
   - `strings.Split(s, sep)` splits a string `s` into a slice of substrings using a delimiter `sep`.

## III. The `sort` Package

1. Introduction

   - The `sort` package provides sorting algorithms for different data types.

2. Sorting Slices

   - `sort.Ints(slice)` sorts an integer slice in ascending order.
   - `sort.Strings(slice)` sorts a string slice in lexicographical order.

3. Custom Sorting
   - Implementing the `sort.Interface` interface allows custom sorting based on specific criteria.
   - Example: Implementing `Less`, `Len`, and `Swap` methods to sort a slice of custom structs.

## Conclusion

The standard library in Go offers a wide range of packages that provide functionalities for common programming tasks. The `strings` package is useful for manipulating and working with strings, while the `sort

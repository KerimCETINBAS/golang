#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_26) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_28)

&#10;

# Lesson 27

# Type Conversion and Encoding/JSON Package in Go

1. Type Conversion in Go:

   - Go supports explicit type conversions between compatible types.
   - The general syntax for type conversion is: `Type(expression)`.
   - Type conversions are useful when you need to convert a value from one type to another.
   - In the given code, the `TypeConverter` function performs type conversion.

2. Encoding/JSON Package in Go:

   - Go provides the `encoding/json` package for encoding and decoding JSON data.
   - It allows you to marshal Go objects into JSON representation and unmarshal JSON data into Go objects.
   - The `json.Marshal` function converts a Go object into its JSON representation.
   - The `json.Unmarshal` function converts JSON data into a Go object.
   - In the given code, the `User` struct is marshaled and unmarshaled using the `json.Marshal` and `json.Unmarshal` functions.

3. Understanding the Code:

   - The code defines two structs: `User` and `UserResponse`.
   - The `User` struct represents user data with various fields.
   - The `UserResponse` struct has a single field `Role` with a tag `json:"role"`.
   - The `TypeConverter` function is a generic function that accepts any type `R` and any data `data`.
   - It converts the `data` to its JSON representation and then unmarshals it into a new object of type `R`.
   - In the `main` function, `TypeConverter` is called with `UserResponse` as the type and an empty `User` object.
   - The returned `data` object is then modified by setting its `Role` field.
   - Finally, the modified `data` object is marshaled back into JSON and printed.

4. Running the Code:
   - After understanding the code, you can run it to observe the output.
   - The code demonstrates how to use type conversion and the `encoding/json` package to marshal and unmarshal JSON data in Go.
   - It showcases the conversion of a `User` object into a `UserResponse` object and back.
   - The modified `UserResponse` object is then printed as a JSON string.

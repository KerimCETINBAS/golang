# Lesson 3

## variables, Strings, Numbers

&nbsp;

## 1. String

No need to talk about :D
&nbsp;

## 2. declaring variables

#### 2.1 explictly typed

variables declared with var keyword following variable name and type

```go
var variableName string = "value"
```

&nbsp;

#### 2.2 implicitly typed

variables infer types from value

```go
// still string
var variableName = "value"
```

&nbsp;

#### 2.3 shorthad declaring without var keyword

only need colon (:) on initial declaration.

```go

variableName := "value"
// this same as
var variableName = "value"
```

Shorthand declerations can not be used on outside of a function

```go
variableName := "value" // error

func main() {
    variableName := "value" // good
}

```

## 3. Numbers

### 3.1 Int & UInt ( unsigned integer)

```go

 int // allias for 32bit integer
 int8 // 2^8 | -256 ~ +256
 int16 // 2^16 | -65536 ~ +65536
 int32 // 2^32 | 32 bit integer | -4294967296 ~ +4294967296
 int64 // 2^64
 // uints are same range but can not be negative
 uint // allias for 32bit uint
 uint8 // 2^8 * 2 | 0 ~ 512
 uint16 // 2^16 * 2
 uint32 // 2^32 * 2
 uint64 //  2^64 * 2
```

### 3.2 Float

used for decimal pointed numbers

```go

float32 // 32 bit float
float64 // 64 bit float
```

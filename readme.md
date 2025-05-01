# Abu

Abu — A Simple Toolkit for Casting, Reflection, and Everyday Go Utilities.

[![Go Reference](https://pkg.go.dev/badge/github.com/amberpixels/abu.svg)](https://pkg.go.dev/github.com/amberpixels/abu)
[![Go Report Card](https://goreportcard.com/badge/github.com/amberpixels/abu)](https://goreportcard.com/report/github.com/amberpixels/abu)

## Overview

Abu is a Go utility library that provides helper functions for type casting, reflection, and other common programming tasks. It's designed to make your code more concise and readable, especially in testing scenarios.

## Installation

```bash
go get github.com/amberpixels/abu
```

## Features

### Type Casting

The `cast` package provides functions for converting between different types:

- `AsString`: Convert a value to a string
- `AsBytes`: Convert a value to a byte slice
- `AsBool`: Convert a value to a boolean
- `AsInt`: Convert a value to an integer
- `AsFloat`: Convert a value to a float
- `AsKind`: Convert a value to a reflect.Kind
- `AsSliceOfAny`: Convert a value to a slice of any
- `AsStrings`: Convert a value to a slice of strings
- `AsTime`: Convert a value to a time.Time

### Type Checking

The `cast` package also provides functions for checking if a value is of a certain type:

- `IsString`: Check if a value is a string (with configurable options)
- `IsStringish`: Check if a value is string-like
- `IsNil`: Check if a value is nil
- `IsInt`: Check if a value is an integer
- `IsStrings`: Check if a value is a slice of strings
- `IsTime`: Check if a value is a time.Time

### Reflection Helpers

The `reflectish` package provides helper functions for working with reflection:

- `IndirectDeep`: Recursively dereference pointers
- `LengthOf`: Get the length of various types (maps, arrays, strings, channels, slices)

## Usage Examples

### Type Casting

```go
import "github.com/amberpixels/abu/cast"

// Convert a byte slice to a string
str = cast.AsString([]byte("byte_data")) // "byte_data"

// Convert a custom string type
type CustomString string
str = cast.AsString(CustomString("example")) // "example"

// Convert an integer
num := cast.AsInt(42) // "42"

// Convert a float
f := cast.AsFloat(3.14) // "3.14"
```

### Type Checking

```go
import "github.com/amberpixels/abu/cast"

// Check if a value is a string
if cast.IsString("example") {
    // It's a string
}

// Check if a value is string-like (string, []byte, etc.)
if cast.IsStringish([]byte("example")) {
    // It's string-like
}

// Check if a value is nil
if cast.IsNil(someValue) {
    // It's nil
}
```

### Reflection Helpers

```go
import (
    "reflect"
    "github.com/amberpixels/abu/reflect"
)

// Recursively dereference pointers
v := reflect.ValueOf(&someValue)
v = reflectish.IndirectDeep(v)

// Get the length of a value
length, ok := reflectish.LengthOf(someValue)
if ok {
    // Length is available
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
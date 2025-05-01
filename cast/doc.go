// Package cast provides utility functions for type casting and type checking.
//
// This package is designed to simplify type conversions and type checking in Go code,
// particularly useful in testing scenarios. It offers a set of functions for converting
// between different types (As* functions) and for checking if a value is of a certain type
// (Is* functions).
//
// The As* functions (AsString, AsBytes, AsBool, AsInt, AsFloat, etc.) attempt to convert
// a value to the specified type, panicking if the conversion is not possible. This approach
// is suitable for testing code where panics are acceptable.
//
// The Is* functions (IsString, IsStringish, IsNil, IsInt, etc.) check if a value is of a
// certain type or can be converted to that type, returning a boolean result.
//
// The package also provides configuration options for customizing type checking behavior,
// particularly for string types.
package cast

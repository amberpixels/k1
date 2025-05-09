// Package maybe provides a generic, type-safe container for values that may or may not be present.
//
// Maybe[T] mirrors the “Maybe”/“Optional” pattern found in other languages, offering a safer
// alternative to pointers when you need to represent “some value” vs. “no value.” It supports:
//
//   - Construction via Some(v) and None[T]().
//   - Querying presence with Some() and None() methods.
//   - Unwrapping with Unwrap(), which panics on None.
//   - JSON marshalling: encodes None as null, Some(v) as v.
//   - TOML marshalling: encodes None as the special TomlNone hack.
//   - Text unmarshalling: treats empty, “null”, or TomlNone (case-insensitive) as None,
//     otherwise attempts to parse into T (using encoding.TextUnmarshaler if available,
//     or a JSON fallback for scalars).
//
// Common helpers include True(), False(), and NoneBool() for boolean Optionals.
//
// Usage:
//
//	import "github.com/amberpixels/k1/maybe"
//
//	opt := maybe.Some(42)
//	if opt.Some() {
//	    fmt.Println("We have:", opt.Unwrap())
//	}
package maybe

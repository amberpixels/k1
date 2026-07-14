<p align="center">
  <img src="logo.svg" alt="k1" width="230">
</p>

<div align="center">

### Every gopher needs a k1t.

Type casting, reflection helpers, and everyday utilities for Go.

[![Go Reference](https://pkg.go.dev/badge/github.com/amberpixels/k1.svg)](https://pkg.go.dev/github.com/amberpixels/k1)
[![Go Version](https://img.shields.io/github/go-mod/go-version/amberpixels/k1)](go.mod)
[![License: MIT](https://img.shields.io/badge/license-MIT-yellow.svg)](LICENSE)

</div>

---

`k1` (read: "k1t") is a small toolkit of the helpers you keep rewriting between projects: type casting that survives custom types and deep pointers, an Option type, safe pointer dereferencing, and set lookups.

The core idea: try a direct type switch first, fall back to reflection. So `cast` functions accept anything shaped right, not just exact types:

```go
type UserID string
id := UserID("u-42")
p := &id

cast.AsString(id) // "u-42"
cast.AsString(&p) // "u-42" - pointers are dereferenced deeply
```

> [!NOTE]
> `As*` functions panic on impossible conversions instead of returning errors. That is by design: k1 is testing-oriented, and in tests a panic is a failure you want loud.

## Install

```bash
go get github.com/amberpixels/k1
```

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/amberpixels/k1/cast"
	"github.com/amberpixels/k1/maybe"
	"github.com/amberpixels/k1/ptr"
	"github.com/amberpixels/k1/set"
)

type UserID string

func main() {
	// cast: conversions that survive custom types and pointers
	id := UserID("u-42")
	fmt.Println(cast.AsString(&id)) // u-42

	// maybe: Option[T] instead of *T
	port := maybe.Some(8080)
	if port.Some() {
		fmt.Println(port.Unwrap()) // 8080
	}

	// ptr: dereference with a zero-value fallback
	var name *string
	fmt.Printf("%q\n", ptr.Deref(name)) // ""

	// set: map[T]struct{} without the ceremony
	admins := set.NewLookup("alice", "bob")
	fmt.Println(admins.Has("mallory")) // false
}
```

## Casting

The `cast` package converts (`As*`) and checks (`Is*`):

```go
cast.AsString([]byte("data")) // "data"
cast.AsBytes("data")          // []byte("data")
cast.AsInt(42.0)              // 42 - integral floats convert; 42.5 panics
cast.AsFloat(42)              // 42.0
cast.AsTime(&customTime)      // time.Time, also from custom time types
```

Full set: `AsString`, `AsBytes`, `AsBool`, `AsInt`, `AsFloat`, `AsTime`, `AsKind`, `AsSliceOfAny`, `AsStrings` - plus `IsString`, `IsStringish`, `IsNil`, `IsInt`, `IsStrings`, `IsTime` for checks.

`IsString` is strict by default (true only for an actual `string`); loosen it per call or globally:

```go
cast.IsString(UserID("u-42"))                          // false - strict by default
cast.IsString(UserID("u-42"), cast.AllowCustomTypes()) // true
cast.IsString([]byte("hi"), cast.AllowAll())           // true - most permissive

cast.ConfigureIsStringConfig(cast.AllowAll()) // change the default globally
```

## Optionals

The `maybe` package is an `Option[T]` for comparable types, with marshaling that behaves well in configs and APIs:

```go
port := maybe.Some(8080)
port.Some()   // true
port.Unwrap() // 8080; panics on None

none := maybe.None[int]()
json.Marshal(port) // 8080
json.Marshal(none) // null
```

None marshals as `null` in JSON and as the `"None"` sentinel in TOML; text unmarshalling treats empty, `"null"`, and `"None"` as None. Shorthands: `maybe.True()`, `maybe.False()`, `maybe.NoneBool()`, `maybe.NoneInt()`.

## Everyday Helpers

- **`ptr`** - `ptr.Deref(p)` dereferences with a zero-value fallback for nil; `ptr.Clone(p)` copies a pointee.
- **`set`** - `set.Lookup[T]` is `map[T]struct{}` with `Has`/`Add`/`Delete`/`Clear`; build one with `set.NewLookup("a", "b")`.
- **`quick`** - `quick.Append(a, b...)` appends only elements not already present; trades extra memory (and GC pressure) for speed on large slices.
- **`errs`** - `errs.UnwrapDeep(err)` walks a wrapped error chain to the root cause.
- **`reflectish`** - `IndirectDeep` for deep pointer dereferencing, `LengthOf` for the length of anything length-y, panic-safe `Interface`.
- **`k1`** (root) - `k1.JoinStringers(vals, ", ")` joins any slice of `fmt.Stringer`s.

## Feedback

k1 is a solo, opinionated project - but if you stumbled upon it and have
ideas, questions, or bug reports, an [issue](https://github.com/amberpixels/k1/issues) is always welcome :)

## License

[MIT](LICENSE) © [amberpixels](https://amberpixels.io)

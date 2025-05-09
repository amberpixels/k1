// Package reflectish provides common reflection utilities that extend Go's standard "reflect" library.
//
// It offers:
//   - IndirectDeep: recursively dereferences pointers until a non-pointer value is reached.
//   - LengthOf: returns the length of supported types (arrays, slices, maps, strings, channels),
//     along with a boolean indicating support.
//
// Usage:
//
//	import "github.com/amberpixels/k1/reflectish"
//
//	// Deeply dereference pointers
//	val := reflect.ValueOf(&&myStruct)
//	root := reflectish.IndirectDeep(val)
//
//	// Get length of a slice, map, etc.
//	length, ok := reflectish.LengthOf([]int{1,2,3})
//	if ok {
//	    fmt.Println("Length:", length)
//	}
//
// Package reflectish is intended as a lightweight helper for reflection-based operations.
package reflectish

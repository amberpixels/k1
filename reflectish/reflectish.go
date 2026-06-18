// Package reflectish contains helpers that extends standard reflect library
package reflectish

import "reflect"

// IndirectDeep does reflect.Indirect deeply.
func IndirectDeep(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	return v
}

// Interface is the panic-safe equivalent of reflect.Value.Interface.
//
// reflect.Value.Interface panics on the zero (invalid) Value — which is exactly
// what IndirectDeep yields when it dereferences through a nil pointer. Interface
// returns nil for an invalid Value instead, so the two compose cleanly:
//
//	v := reflectish.Interface(reflectish.IndirectDeep(rv)) // nil for nil pointers
func Interface(v reflect.Value) any {
	if !v.IsValid() {
		return nil
	}
	return v.Interface()
}

// IndirectInterface deeply dereferences pointers and returns the underlying
// value as any, yielding nil when it dereferences through a nil pointer. It is
// shorthand for Interface(IndirectDeep(v)).
func IndirectInterface(v reflect.Value) any {
	return Interface(IndirectDeep(v))
}

// LengthOf returns length of a given type.
func LengthOf(a any) (int, bool) {
	if a == nil {
		return 0, false
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Map, reflect.Array, reflect.String, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(a).Len(), true
	default:
		return 0, false
	}
}

package gotils

import (
	"path"
	"reflect"
)


// IsEmpty determines whether the given object should be considered empty.
// It handles various types differently:
// - nil values are always considered empty
// - Collections (slices, maps, channels) are empty when they have no elements
// - Pointers are empty if nil or if they point to an empty value
// - For other types (including arrays), it compares against the zero value
// Returns true if the object is considered empty, false otherwise.
func IsEmpty(object any) bool {
	// get nil case out of the way
	if object == nil {
		return true
	}

	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {
	// collection types are empty when they have no element
	case reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
	// pointers are empty if nil or if the value they point to is empty
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		deref := objValue.Elem().Interface()
		return IsEmpty(deref)
	// for all other types, compare against the zero value
	// array types are empty when they match their zero-initialized state
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(object, zero.Interface())
	}
}

// TypeName returns the fully qualified name of the type of the given value.
// The returned string includes the package path (base name only) and the type name,
// separated by a dot. If the type is not from a package (e.g., built-in types),
// only the type name is returned.
func TypeName(value any) string {
	ty := reflect.TypeOf(value)
	pkg := path.Base(ty.PkgPath())
	if pkg == "." {
		return ty.Name()
	}
	return pkg + "." + ty.Name()
}

package utils

import "reflect"

func CloneInterface[T any](src T) (T, bool) {
	orig := reflect.ValueOf(src)

	if orig.Kind() == reflect.Pointer {
		orig = orig.Elem()
	}

	copy := reflect.New(orig.Type()).Elem()

	copy.Set(orig)

	t, ok := copy.Addr().Interface().(T)
	return t, ok
}

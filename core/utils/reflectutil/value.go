package reflectutil

import (
	"bytes"
	"reflect"
)

// IsNil
// prevent panic from go/src/reflect/value.go:1551
func IsNil(v any) bool {
	if v == nil {
		return true
	}

	switch val := reflect.ValueOf(v); val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map,
		reflect.Pointer, reflect.UnsafePointer,
		reflect.Interface, reflect.Slice:
		return val.IsNil()
	}

	return false
}

func Equal(v1, v2 any) bool {
	if v1 == nil || v2 == nil {
		return v1 == v2
	}
	bs1, ok := v1.([]byte)
	if !ok {
		return reflect.DeepEqual(v1, v2)
	}
	bs2, ok := v2.([]byte)
	if !ok {
		return false
	}
	if bs1 == nil || bs2 == nil {
		return bs1 == nil && bs2 == nil
	}
	return bytes.Equal(bs1, bs2)
}

func IsFunc(v any) bool {
	if v == nil {
		return false
	}
	return reflect.TypeOf(v).Kind() == reflect.Func
}

// ToPtr
// if v is ptr return itself, else return ptr of v
// second return value is true if v is ptr
func ToPtr(v any) (any, bool) {
	rv, rt := reflect.ValueOf(v), reflect.TypeOf(v)
	if rv.Kind() == reflect.Ptr {
		if !rv.IsNil() {
			return v, true
		}
		return reflect.New(rt.Elem()).Interface(), true
	}
	return reflect.New(rt).Interface(), false
}

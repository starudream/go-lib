package http

import (
	"errors"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
)

func filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

func structFields(obj any, tagName string, cb func(val reflect.Value, tag string) error) error {
	rv := reflect.ValueOf(obj)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		panic("not struct")
	}
	vt := rv.Type()
	for i := 0; i < vt.NumField(); i++ {
		t := vt.Field(i).Tag.Get(tagName)
		if t == "" || t == "-" {
			continue
		}
		if e := cb(rv.Field(i), t); e != nil {
			return e
		}
	}
	return nil
}

func setAny(obj any, tagName string, values interface{ Get(string) string }) error {
	return structFields(obj, tagName, func(val reflect.Value, tag string) error {
		s := values.Get(tag)
		if s == "" {
			return nil
		}
		v, e := castAny(val.Kind(), s)
		if e != nil {
			return e
		}
		val.Set(reflect.ValueOf(v))
		return nil
	})
}

var (
	typeMultiFileHeader = reflect.TypeOf((*multipart.FileHeader)(nil))
	typeMultiFile       = reflect.TypeOf((*multipart.File)(nil)).Elem()
)

func setMulti(obj any, tagName string, form *multipart.Form) error {
	return structFields(obj, tagName, func(val reflect.Value, tag string) error {
		v, e := castMulti(val, form, tag)
		if e != nil {
			if errors.Is(e, errEmptyValue) {
				e = nil
			}
			return e
		}
		val.Set(reflect.ValueOf(v))
		return nil
	})
}

var errEmptyValue = errors.New("empty value")

func castMulti(fv reflect.Value, form *multipart.Form, t string) (any, error) {
	switch fvt := fv.Type(); {
	case fvt == typeMultiFileHeader:
		fhs := form.File[t]
		if len(fhs) == 0 {
			return nil, http.ErrMissingFile
		}
		return fhs[0], nil
	case fvt.Implements(typeMultiFile):
		fhs := form.File[t]
		if len(fhs) == 0 {
			return nil, http.ErrMissingFile
		}
		return fhs[0].Open()
	default:
		vs := form.Value[t]
		if len(vs) == 0 {
			return nil, errEmptyValue
		}
		return castAny(fv.Kind(), vs[0])
	}
}

func castAny(k reflect.Kind, s string) (any, error) {
	switch k {
	case reflect.Bool:
		return strconv.ParseBool(s)

	case reflect.Int:
		t, e := strconv.ParseInt(s, 10, 0)
		return int(t), e
	case reflect.Int8:
		t, e := strconv.ParseInt(s, 10, 8)
		return int8(t), e
	case reflect.Int16:
		t, e := strconv.ParseInt(s, 10, 16)
		return int16(t), e
	case reflect.Int32:
		t, e := strconv.ParseInt(s, 10, 32)
		return int32(t), e
	case reflect.Int64:
		return strconv.ParseInt(s, 10, 64)

	case reflect.Uint:
		t, e := strconv.ParseUint(s, 10, 0)
		return uint(t), e
	case reflect.Uint8:
		t, e := strconv.ParseUint(s, 10, 8)
		return uint8(t), e
	case reflect.Uint16:
		t, e := strconv.ParseUint(s, 10, 16)
		return uint16(t), e
	case reflect.Uint32:
		t, e := strconv.ParseUint(s, 10, 32)
		return uint32(t), e
	case reflect.Uint64:
		return strconv.ParseUint(s, 10, 64)

	case reflect.Float32:
		t, e := strconv.ParseFloat(s, 32)
		return float32(t), e
	case reflect.Float64:
		return strconv.ParseFloat(s, 64)

	case reflect.String:
		return s, nil

	default:
		panic("unsupported type: " + k.String())
	}
}

package config

import (
	"time"

	"github.com/spf13/cast"
)

type Value interface {
	Any() any
	Int64(def ...int64) int64
	Int(def ...int) int
	Float64(def ...float64) float64
	Duration(def ...time.Duration) time.Duration
	Time(def ...time.Time) time.Time
	String(def ...string) string
	Bool(def ...bool) bool
}

type value struct {
	v any
}

func NewValue(v any) Value {
	return &value{v}
}

var _ Value = (*value)(nil)

func (v *value) Any() any {
	return v.v
}

func (v *value) Int64(def ...int64) int64 {
	return iCast(v.v, cast.ToInt64E, def...)
}

func (v *value) Int(def ...int) int {
	return iCast(v.v, cast.ToIntE, def...)
}

func (v *value) Float64(def ...float64) float64 {
	return iCast(v.v, cast.ToFloat64E, def...)
}

func (v *value) Duration(def ...time.Duration) time.Duration {
	return iCast(v.v, cast.ToDurationE, def...)
}

func (v *value) Time(def ...time.Time) time.Time {
	return iCast(v.v, cast.ToTimeE, def...)
}

func (v *value) String(def ...string) string {
	return iCast(v.v, cast.ToStringE, def...)
}

func (v *value) Bool(def ...bool) bool {
	return iCast(v.v, cast.ToBoolE, def...)
}

func iCast[T any](v any, fn func(i any) (T, error), def ...T) (t T) {
	if v != nil {
		x, e := fn(v)
		if e == nil {
			return x
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return
}

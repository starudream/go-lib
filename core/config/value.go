package config

import (
	"time"

	"github.com/spf13/cast"
)

type Value interface {
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

func (v *value) Int64(def ...int64) int64 {
	return iArr[int64](def).Cast(cast.ToInt64E(v.v))
}

func (v *value) Int(def ...int) int {
	return iArr[int](def).Cast(cast.ToIntE(v.v))
}

func (v *value) Float64(def ...float64) float64 {
	return iArr[float64](def).Cast(cast.ToFloat64E(v.v))
}

func (v *value) Duration(def ...time.Duration) time.Duration {
	return iArr[time.Duration](def).Cast(cast.ToDurationE(v.v))
}

func (v *value) Time(def ...time.Time) time.Time {
	return iArr[time.Time](def).Cast(cast.ToTimeE(v.v))
}

func (v *value) String(def ...string) string {
	return iArr[string](def).Cast(cast.ToStringE(v.v))
}

func (v *value) Bool(def ...bool) bool {
	return iArr[bool](def).Cast(cast.ToBoolE(v.v))
}

type iArr[T any] []T

func (arr iArr[T]) Cast(v T, err error) (t T) {
	if err == nil {
		return v
	}
	if len(arr) > 0 {
		return arr[0]
	}
	return
}

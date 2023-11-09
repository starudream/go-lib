package sliceutil

func GetValue[T any](vs []T, idx int, def ...T) (v T, exists bool) {
	if idx >= 0 && idx < len(vs) {
		return vs[idx], true
	}
	exists = false
	if len(def) > 0 {
		v = def[0]
	}
	return
}

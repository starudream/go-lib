package codec

import (
	"fmt"

	"github.com/starudream/go-lib/core/v2/utils/reflectutil"
)

func UnmarshalTo[T any](o Unmarshaler, a any) (t T, err error) {
	v, ptr := reflectutil.ToPtr(t)
	switch x := a.(type) {
	case []byte:
		err = o.Unmarshal(x, v)
	case string:
		err = o.UnmarshalString(x, v)
	default:
		err = fmt.Errorf("unsupported type: %T", a)
	}
	if ptr {
		return v.(T), err
	}
	return *(v.(*T)), err
}

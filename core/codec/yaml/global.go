package yaml

import (
	"github.com/starudream/go-lib/core/v2/codec"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

var G = Options{
	MarshalOptions: MarshalOptions{
		IndentSpace:             2,
		IndentSequence:          true,
		SingleQuote:             false,
		FlowStyle:               false,
		LiteralStyleIfMultiline: true,
	},
	UnmarshalOptions: UnmarshalOptions{},
}

func Marshal(v any) ([]byte, error) {
	return G.Marshal(v)
}

func MarshalString(v any) (string, error) {
	return G.MarshalString(v)
}

func MustMarshal(v any) []byte {
	return G.MustMarshal(v)
}

func MustMarshalString(v any) string {
	return G.MustMarshalString(v)
}

func Unmarshal(bs []byte, v any) error {
	return G.Unmarshal(bs, v)
}

func UnmarshalString(s string, v any) error {
	return G.Unmarshal([]byte(s), v)
}

func MustUnmarshal(bs []byte, v any) {
	G.MustUnmarshal(bs, v)
}

func MustUnmarshalString(s string, v any) {
	G.MustUnmarshal([]byte(s), v)
}

func UnmarshalTo[T any](a any) (T, error) {
	return codec.UnmarshalTo[T](G, a)
}

func MustUnmarshalTo[T any](a any) T {
	return osutil.Must1(codec.UnmarshalTo[T](G, a))
}

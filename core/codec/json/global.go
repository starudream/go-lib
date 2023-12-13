package json

import (
	"github.com/starudream/go-lib/core/v2/codec"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

var G = Options{
	MarshalOptions: MarshalOptions{
		PrefixStr: "",
		IndentStr: "  ",
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

func MarshalIndent(v any) ([]byte, error) {
	return G.MarshalIndent(v)
}

func MarshalIndentString(v any) (string, error) {
	return G.MarshalIndentString(v)
}

func MustMarshalIndent(v any) []byte {
	return G.MustMarshalIndent(v)
}

func MustMarshalIndentString(v any) string {
	return G.MustMarshalIndentString(v)
}

func Compact(src []byte) ([]byte, error) {
	return G.Compact(src)
}

func CompactString(src string) (string, error) {
	return G.CompactString(src)
}

func MustCompact(src []byte) []byte {
	return G.MustCompact(src)
}

func MustCompactString(src string) string {
	return G.MustCompactString(src)
}

func Indent(src []byte) ([]byte, error) {
	return G.Indent(src)
}

func IndentString(src string) (string, error) {
	return G.IndentString(src)
}

func MustIndent(src []byte) []byte {
	return G.MustIndent(src)
}

func MustIndentString(src string) string {
	return G.MustIndentString(src)
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

package codec

type Marshaler interface {
	Marshal(v any) ([]byte, error)
	MarshalString(v any) (string, error)

	MustMarshal(v any) []byte
	MustMarshalString(v any) string
}

type IndentMarshaler interface {
	MarshalIndent(v any) ([]byte, error)
	MarshalIndentString(v any) (string, error)

	MustMarshalIndent(v any) []byte
	MustMarshalIndentString(v any) string
}

type Unmarshaler interface {
	Unmarshal(bs []byte, v any) error
	UnmarshalString(s string, v any) error

	MustUnmarshal(bs []byte, v any)
	MustUnmarshalString(s string, v any)
}

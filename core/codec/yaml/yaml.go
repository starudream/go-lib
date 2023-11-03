package yaml

import (
	"github.com/goccy/go-yaml"

	"github.com/starudream/go-lib/core/v2/codec"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

type Options struct {
	MarshalOptions
	UnmarshalOptions
}

type MarshalOptions struct {
	IndentSpace             int
	IndentSequence          bool
	SingleQuote             bool
	FlowStyle               bool
	LiteralStyleIfMultiline bool
}

var _ codec.Marshaler = MarshalOptions{}

func (o MarshalOptions) Marshal(v any) ([]byte, error) {
	opts := []yaml.EncodeOption{
		yaml.Flow(o.FlowStyle),
		yaml.Indent(o.IndentSpace),
		yaml.IndentSequence(o.IndentSequence),
		yaml.UseSingleQuote(o.SingleQuote),
		yaml.UseLiteralStyleIfMultiline(o.LiteralStyleIfMultiline),
	}
	return yaml.MarshalWithOptions(v, opts...)
}

func (o MarshalOptions) MarshalString(v any) (string, error) {
	bs, err := o.Marshal(v)
	return string(bs), err
}

func (o MarshalOptions) MustMarshal(v any) []byte {
	bs, err := o.Marshal(v)
	osutil.PanicErr(err)
	return bs
}

func (o MarshalOptions) MustMarshalString(v any) string {
	return string(o.MustMarshal(v))
}

type UnmarshalOptions struct {
}

var _ codec.Unmarshaler = UnmarshalOptions{}

func (o UnmarshalOptions) Unmarshal(bs []byte, v any) error {
	return yaml.Unmarshal(bs, v)
}

func (o UnmarshalOptions) UnmarshalString(s string, v any) error {
	return o.Unmarshal([]byte(s), v)
}

func (o UnmarshalOptions) MustUnmarshal(bs []byte, v any) {
	osutil.PanicErr(o.Unmarshal(bs, v))
}

func (o UnmarshalOptions) MustUnmarshalString(s string, v any) {
	o.MustUnmarshal([]byte(s), v)
}

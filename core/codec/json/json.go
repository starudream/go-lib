package json

import (
	"github.com/goccy/go-json"

	"github.com/starudream/go-lib/core/v2/codec"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

type Options struct {
	MarshalOptions
	UnmarshalOptions
}

type MarshalOptions struct {
	DisableHTMLEscape    bool
	DisableNormalizeUTF8 bool

	Prefix string
	Indent string
}

func (o MarshalOptions) opts() (opts []json.EncodeOptionFunc) {
	if o.DisableHTMLEscape {
		opts = append(opts, json.DisableHTMLEscape())
	}
	if o.DisableNormalizeUTF8 {
		opts = append(opts, json.DisableNormalizeUTF8())
	}
	return
}

var _ codec.Marshaler = MarshalOptions{}

func (o MarshalOptions) Marshal(v any) ([]byte, error) {
	return json.MarshalWithOption(v, o.opts()...)
}

func (o MarshalOptions) MarshalString(v any) (string, error) {
	bs, err := o.Marshal(v)
	return string(bs), err
}

func (o MarshalOptions) MustMarshal(v any) []byte {
	return osutil.Must1(o.Marshal(v))
}

func (o MarshalOptions) MustMarshalString(v any) string {
	return string(o.MustMarshal(v))
}

var _ codec.IndentMarshaler = MarshalOptions{}

func (o MarshalOptions) MarshalIndent(v any) ([]byte, error) {
	return json.MarshalIndentWithOption(v, o.Prefix, o.Indent, o.opts()...)
}

func (o MarshalOptions) MarshalIndentString(v any) (string, error) {
	bs, err := o.MarshalIndent(v)
	return string(bs), err
}

func (o MarshalOptions) MustMarshalIndent(v any) []byte {
	return osutil.Must1(o.MarshalIndent(v))
}

func (o MarshalOptions) MustMarshalIndentString(v any) string {
	return string(o.MustMarshalIndent(v))
}

type UnmarshalOptions struct {
}

var _ codec.Unmarshaler = UnmarshalOptions{}

func (o UnmarshalOptions) Unmarshal(bs []byte, v any) error {
	return json.UnmarshalWithOption(bs, v)
}

func (o UnmarshalOptions) UnmarshalString(s string, v any) error {
	return o.Unmarshal([]byte(s), v)
}

func (o UnmarshalOptions) MustUnmarshal(bs []byte, v any) {
	osutil.Must0(o.Unmarshal(bs, v))
}

func (o UnmarshalOptions) MustUnmarshalString(s string, v any) {
	o.MustUnmarshal([]byte(s), v)
}

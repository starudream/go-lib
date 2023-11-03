package testdata

type V struct {
	A string         `json:"a,omitempty" yaml:"a,omitempty"`
	B map[string]any `json:"b" yaml:"b"`

	h string
}

var (
	V1 = &V{
		A: "hello world",
		B: map[string]any{
			"int":   16,
			"bool":  true,
			"float": 3.14,
		},
	}
	V2 = &V{
		B: V1.B,
		h: "hidden",
	}
	V3 = &(*V1)
)

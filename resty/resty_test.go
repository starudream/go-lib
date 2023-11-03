package resty

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestR(t *testing.T) {
	type Res struct {
		Count       int     `json:"count"`
		Name        string  `json:"name"`
		Gender      string  `json:"gender"`
		Probability float64 `json:"probability"`
	}
	res := &Res{}

	resp, err := R(WithUserAgent(UAWindowsChrome)).SetResult(res).Get("https://api.genderize.io/?name=peter")
	testutil.LogNoErr(t, err, resp.Status(), res)
}

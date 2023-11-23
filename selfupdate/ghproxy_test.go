package selfupdate

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestGithubProxy(t *testing.T) {
	p := (&GithubProxy{}).Check()
	items := p.Items()
	for i := 0; i < len(items); i++ {
		testutil.Log(t, items[i])
	}
	testutil.Log(t, p.Fast())
}

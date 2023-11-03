package level

import (
	"testing"
)

func Test(t *testing.T) {
	t.Log(Debug, Debug.ShortString())
	t.Log(Info, Info.ShortString())
	t.Log(Warn, Warn.ShortString())
	t.Log(Error, Error.ShortString())
	t.Log(Fatal, Fatal.ShortString())
}

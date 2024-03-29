package cron

import (
	"fmt"
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/signalutil"
	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func Test(t *testing.T) {
	t.Run("second", func(t *testing.T) {
		RunJob(t, "* * * * * *", "t", func() { fmt.Println("ok") }, 1)
	})

	t.Run("panic", func(t *testing.T) {
		RunJob(t, "* * * * * *", "t", func() { panic("?") }, 1)
	})
}

func RunJob(t *testing.T, spec, name string, job func(), count int) {
	err := AddJob(spec, name, func() {
		if count <= 0 {
			return
		}
		defer func() {
			count--
			if count <= 0 {
				Remove(name)
				signalutil.Cancel()
			}
		}()
		job()
	})
	testutil.Nil(t, err)
	Run()
}

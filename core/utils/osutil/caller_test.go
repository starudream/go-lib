package osutil

import (
	"runtime"
	"testing"
)

func TestStack(t *testing.T) {
	t.Log("\n" + Stack())
}

func TestCallerPC(t *testing.T) {
	pc := CallerPC()
	f := runtime.FuncForPC(pc)
	t.Log(f.Name())
	t.Log(f.FileLine(pc))        // caller line
	t.Log(f.FileLine(f.Entry())) // caller function
}

// func TestCallerFn(t *testing.T) {
// 	skipNames := []string{
// 		"caller_test.go",
// 	}
//
// 	fn := func(frame runtime.Frame) bool {
// 		for _, name := range skipNames {
// 			if strings.Contains(frame.File, name) {
// 				return true
// 			}
// 		}
// 		return false
// 	}
//
// 	root := &cobra.Command{Use: "mod", Run: func(cmd *cobra.Command, args []string) {
// 		pc := CallerFn(fn)
// 		t.Log(pc)
// 		f := runtime.FuncForPC(pc)
// 		t.Log(f.FileLine(pc))
// 	}}
// 	_ = root.Execute()
// }

func TestCallerString(t *testing.T) {
	t.Run("std", func(t *testing.T) {
		t.Log(CallerString()) // osutil/caller_test.go:48
	})

	// t.Run("mod", func(t *testing.T) {
	// 	root := &cobra.Command{Use: "mod", Run: func(cmd *cobra.Command, args []string) {
	// 		t.Log(CallerString())  // osutil/caller_test.go:53
	// 		t.Log(CallerString(1)) // cobra@v1.7.0/command.go:944
	// 	}}
	// 	_ = root.Execute()
	// })
}

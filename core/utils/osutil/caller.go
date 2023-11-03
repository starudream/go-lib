package osutil

import (
	"bufio"
	"bytes"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

func Stack(skip ...int) string {
	t := 2
	if len(skip) > 0 {
		t += skip[0]
	}
	bs, ss := debug.Stack(), make([]string, 0, 20)
	for i, sc := 0, bufio.NewScanner(bytes.NewReader(bs)); sc.Scan(); i++ {
		if i == 0 {
			// skip first line: goroutine 20 [running]:
			continue
		}
		if i%2 != 0 {
			t--
			// skip line: package with function name
			continue
		}
		if t >= 0 {
			continue
		}
		s := sc.Text()
		idx1 := strings.LastIndexByte(s, ':')
		if idx1 == -1 {
			return string(bs)
		}
		file := strings.TrimPrefix(s[:idx1], "\t")
		line := s[idx1+1:]
		idx2 := strings.IndexByte(s[idx1+1:], ' ')
		if idx2 > -1 {
			line = s[idx1+1 : idx1+1+idx2]
		}
		ss = append(ss, file+":"+line)
	}
	if ArgTest() && len(ss) > 2 {
		// skip testing Run wrapper
		ss = ss[:len(ss)-2]
	}
	return strings.Join(ss, "\n")
}

func CallerPC(skip ...int) uintptr {
	t := 2
	if len(skip) > 0 {
		t += skip[0]
	}
	var pcs [1]uintptr
	runtime.Callers(t, pcs[:])
	return pcs[0]
}

type CallerFrame = runtime.Frame

func CallerFn(fn func(frame CallerFrame) bool) uintptr {
	var pcs [20]uintptr
	runtime.Callers(2, pcs[:])
	frames := runtime.CallersFrames(pcs[:])
	for {
		frame, more := frames.Next()
		if !fn(frame) {
			return frame.PC
		}
		if !more {
			break
		}
	}
	return 0
}

func CallerString(skip ...int) string {
	t := 1
	if len(skip) > 0 {
		t += skip[0]
	}
	_, file, line, ok := runtime.Caller(t)
	if !ok {
		return "???"
	}
	return fmt.Sprintf("%s:%d", CallerFormatPath(file), line)
}

func CallerFormatPath(path string, max ...int) string {
	t := 2
	if len(max) > 0 {
		t = max[0]
	}
	bs := []byte(path)
	for i, cnt := len(bs)-1, 0; i >= 0; i-- {
		if bs[i] == '@' {
			cnt = 1
		} else if bs[i] == '/' {
			cnt++
			if cnt >= t {
				path = path[i+1:]
				break
			}
		}
	}
	return path
}

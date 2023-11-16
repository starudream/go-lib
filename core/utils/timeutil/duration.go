package timeutil

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

func JitterDuration(min, max time.Duration, attempt int) time.Duration {
	temp := math.Min(float64(max), float64(min)*math.Exp2(float64(attempt)))
	center := time.Duration(temp / 2)
	if center == 0 {
		center = time.Nanosecond
	}
	dur := randDuration(center)
	if dur < min {
		dur = min
	}
	return dur
}

var (
	rnd   = rand.New(rand.NewSource(time.Now().UnixNano()))
	rndMu sync.Mutex
)

func randDuration(center time.Duration) time.Duration {
	rndMu.Lock()
	defer rndMu.Unlock()
	return time.Duration(math.Abs(float64(int64(center) + rnd.Int63n(int64(center)))))
}

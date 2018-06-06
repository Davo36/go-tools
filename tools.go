package tools

import "time"

// Some handy tools for go projects.

// StopWatch is just a handy way to measure time taken for various tasks.
type StopWatch struct {
	start time.Time
}

// StartTimer starts our StopWatch
func (sw *StopWatch) StartTimer() {
	sw.start = time.Now()
}

// EndTimer returns the amount of time since the StopWatch was started - as a time.Duration type.
// Use time.Duration.Seconds() to get elapsed seconds for instance.
func (sw *StopWatch) EndTimer() time.Duration {
	return (time.Since(sw.start))
}

// Go doesn't have integer min and max functions.  Only float ones. So creating my own here.

// Min : A wee min function.
func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b

}

// Max : A massive max function.
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b

}

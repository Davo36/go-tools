package tools

import "time"

// Some handy tools.

// StopWatch is just a handy way to measure time taken for various tasks.
type StopWatch struct {
	start time.Time
}

// StartTimer starts our stopwatch
func (sw *StopWatch) StartTimer() {
	sw.start = time.Now()
}

// EndTimer stops our stopwatch
func (sw *StopWatch) EndTimer() time.Duration {
	return (time.Since(sw.start))
}

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

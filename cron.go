package thenovadiary

import "time"

// TimeDeltaDays will calculate the absolute
// value of the difference (delta) in days
// between two given points in time.
func TimeDeltaDays(t1, t2 time.Time) int {
	// T0 - - T1 - T2 - - Tn ->	true
	if t2.After(t1) {
		return int(t2.Sub(t1).Hours() / 24)
	}
	return int(t1.Sub(t2).Hours() / 24)
}

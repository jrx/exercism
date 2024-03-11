// Package gigasecond contains a routine for manipulating time
package gigasecond

import "time"

// AddGigasecond returns the time after a gigasecond has passed on a given date
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}

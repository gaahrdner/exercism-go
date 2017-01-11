// Package gigasecond provides the AddGigasecond function to determine when 10^9 seconds have occurred following a given date
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4

// AddGigasecond will return the time, given a date that is exactly 10^9 seconds later.
// https://golang.org/pkg/time/#Time.Add
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1000000000) * time.Second)
}

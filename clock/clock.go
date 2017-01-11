// Package clock represents a 24-hour time notation
package clock

import "fmt"

const testVersion = 4

// Clock struct represents a 24-hour clock
type Clock struct {
	Hour   int
	Minute int
}

// New creates a new Clock type
func New(hour, minute int) Clock {
	c := Clock{}
	c.Hour = hour
	c = c.Add(minute)
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

// Add will take an arbitrary integer of minutes and update a Clock struct accordingly
func (c Clock) Add(minutes int) Clock {
	c.Minute += minutes

	// increment the hour if minutes are greater than 60, the modulus is the new minute
	if c.Minute >= 60 {
		c.Hour += c.Minute / 60
		c.Minute = c.Minute % 60
	}
	// if we have negative minutes, decrement the hour for every successful 60 division, and an extra to cover the < 60 case, then minutes are the modulus
	if c.Minute < 0 {
		c.Hour += (c.Minute / 60) + (-1)
		c.Minute = c.Minute%60 + 60
	}

	// negative hours are the modulus of 24 subtracted from 24
	if c.Hour < 0 {
		c.Hour = c.Hour%24 + 24
	}

	// 25th hour is just the moduls of 24
	if c.Hour >= 24 {
		c.Hour = c.Hour % 24
	}

	return c
}

// Leap stub file

// Package leap provides a function to determine if an integer is a leap year
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// IsLeapYear returns a boolean if a given integer is a leap year.
func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)

}

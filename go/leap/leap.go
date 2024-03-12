// Package leaps contains a routine for leap year checking
package leap

// IsLeapYear returns a bool on whether a given year is a leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

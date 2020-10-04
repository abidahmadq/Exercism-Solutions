// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"
import "math"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Duration(math.Pow(10, 9)) * time.Second

	return t.Add(gigasecond)
}

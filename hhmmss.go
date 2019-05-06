// Package hhmmss manages converting HH:MM:SS time strings to time.Duration values.
package hhmmss

import (
	"strconv"
	"time"
)

const (
	second = iota
	minute
	hour
)

// Parse converts a given stopwatch time string to a time.Duration value.
// The argument should be in the format HH:MM:SS or MM:SS. If the argument
// does not contain any colons (i.e. 1200), the argument is parsed as seconds.
func Parse(value string) (time.Duration, error) {
	var total time.Duration

	// convert argument to rune slice for parsing
	arr := []rune(value)

	// we are going to build our time duration in reverse, so we need to keep track
	// of what part we are currently adding up... 0 = second, 1 = minute, 2 = hour
	part := second

	// collect numbers between each ':' in value
	var collect []rune
	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		if v != ':' {
			collect = append(collect, v)
		}

		if v == ':' || i == 0 {
			// we built the duration in reverse, so we need to restore it to the proper order
			reverse(collect)

			c, err := strconv.Atoi(string(collect))
			if err != nil {
				return 0, err
			}

			d := time.Duration(c)

			switch part {
			case second:
				total += d * time.Second
				break
			case minute:
				total += d * time.Minute
			case hour:
				total += d * time.Hour
			}

			part++
			collect = nil
		}
	}

	return total, nil
}

// reverse reverses a slice.
func reverse(s []rune) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
}

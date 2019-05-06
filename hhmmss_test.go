package hhmmss

import (
	"testing"
	"time"
)

func TestInputs(t *testing.T) {
	type Test struct {
		arg      string
		expected time.Duration
	}

	tt := []Test{
		Test{ // should be parsed in seconds if no ':' set
			arg:      "1200",
			expected: time.Second * time.Duration(1200),
		},
		Test{
			arg:      "01:00:00", // HH:MM:SS
			expected: time.Hour * 1,
		},
		Test{ // anything after hours defined should be ignored
			arg:      "10:01:00:00", // HH:MM:SS
			expected: time.Hour * 1,
		},
		Test{
			arg:      "1:00:00",
			expected: time.Hour * 1,
		},
		Test{
			arg:      "03:30:26",
			expected: (time.Hour * 3) + (time.Minute * 30) + (time.Second * 26),
		},
		Test{
			arg:      "30:26", // MM:SS
			expected: (time.Minute * 30) + (time.Second * 26),
		},
		Test{
			arg:      "3:26", // MM:SS
			expected: (time.Minute * 3) + (time.Second * 26),
		},
		Test{
			arg:      "3:64", // MM:SS (time should roll over)
			expected: (time.Minute * 4) + (time.Second * 4),
		},
	}

	for i, test := range tt {
		dur, err := Parse(test.arg)
		if err != nil {
			t.Fatal(err)
		}

		if dur != test.expected {
			t.Errorf("#%v did not receive expected duration", i)
		}
	}
}

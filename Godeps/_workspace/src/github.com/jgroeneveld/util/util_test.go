package util

import "testing"
import "time"

func TestEndOfMonth(t *testing.T) {
	tests := []struct {
		in       time.Time
		expected time.Time
	}{
		{someDay(2014, 06, 01), EndOfDay(someDay(2014, 06, 30))},
		{someDay(2014, 12, 31), EndOfDay(someDay(2014, 12, 31))},
		{someDay(2014, 12, 15), EndOfDay(someDay(2014, 12, 31))},
		{someDay(2014, 02, 15), EndOfDay(someDay(2014, 02, 28))},
		{someDay(2016, 02, 15), EndOfDay(someDay(2016, 02, 29))},
	}

	for _, test := range tests {
		actual := EndOfMonth(test.in)
		if actual != test.expected {
			t.Errorf("EndOfMonth for %v does not match %v but was %v",
				test.in, test.expected, actual)
		}
	}
}

func someDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 12, 30, 59, 0, time.UTC)
	// return time.Date(year, month, day, 23, 59, 59, 0, nil)
}

package util

import "time"

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func BeginningOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 01, 0, 0, 0, 0, date.Location())
}

func EndOfMonth(date time.Time) time.Time {
	month := date.Month()
	year := date.Year()

	if month == time.December {
		year += 1
		month = time.January
	} else {
		month += 1
	}

	t := time.Date(year, month, 1, 0, 0, 0, 0, date.Location())
	prev := t.Add(-time.Hour)

	return EndOfDay(prev)
}

func EndOfDay(day time.Time) time.Time {
	return time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 0, day.Location())
}

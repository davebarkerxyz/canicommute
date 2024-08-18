package commute

import (
	"fmt"
	"os"
	"time"
)

func Die(errText string, args ...any) {
	fmt.Fprintf(os.Stderr, errText+"\n", args...)
	os.Exit(1)
}

// Get the next (or current) working day at the specified time
func getNextWorkingDay(hour int, min int) time.Time {
	now := time.Now()
	year, month, day := now.Date()
	nwd := time.Date(year, month, day, hour, min, 0, 0, time.Local)
	if now.Weekday() == time.Saturday {
		nwd = nwd.AddDate(0, 0, 2)
	} else if now.Weekday() == time.Sunday {
		nwd = nwd.AddDate(0, 0, 1)
	}

	return nwd
}

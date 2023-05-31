package util

import (
	"fmt"
	"time"
)

var secondsPerMin = int64(60)
var secondsPerHour = int64(60 * secondsPerMin)
var secondsPerDay = int64(24 * secondsPerHour)

// GetDuration returns long to string format
func GetDuration(seconds int64) string {
	day := seconds / secondsPerDay
	hour := (seconds - (day * secondsPerDay)) / secondsPerHour
	min := (seconds - (day * secondsPerDay) - (hour * secondsPerHour)) / secondsPerMin
	secs := (seconds - (day * secondsPerDay) - (hour * secondsPerHour) - (min * secondsPerMin))
	return fmt.Sprintf("%dD %dH %dM %dS", day, hour, min, secs)
}

func MillisBetween(from time.Time, to time.Time) int32 {
	return int32(to.Sub(from).Milliseconds())
}

func MillisToNow(from time.Time) int32 {
	return int32(time.Now().Sub(from).Milliseconds())
}

func TimeIsZero(t time.Time) bool {
	return t.Year() == 0001 && t.Hour() == 0 && t.Second() == 0 && t.Nanosecond() == 0
}

func CurrentYYMMDD() string {
	t := time.Now()
	return fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())

}

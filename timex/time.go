package timex

import "time"

const (
	DateTimeLayout = "2006-01-02 15:04:05" // time.DateTime
	DateLayout     = "2006-01-02"          // time.DateOnly
	TimeLayout     = "15:04:05"            // time.TimeOnly
)

func DateTimeString(t time.Time) string {
	return t.Format(DateTimeLayout)
}

func DateString(t time.Time) string {
	return t.Format(DateLayout)
}

func TimeString(t time.Time) string {
	return t.Format(TimeLayout)
}

func Current() string {
	return DateTimeString(time.Now())
}

func CurrentDate() string {
	return DateString(time.Now())
}

func CurrentTime() string {
	return TimeString(time.Now())
}

func CurrentUnix() int64 {
	return time.Now().Unix()
}

// CurrentMicros returns the current microseconds.
func CurrentMicros() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// CurrentMillis returns the current milliseconds.
func CurrentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// StrParseToTime return time of string
func StrParseToTime(t string, location ...time.Location) (time.Time, error) {
	var local *time.Location
	if len(location) > 0 {
		local = &location[0]
	} else {
		local = time.Local
	}

	return time.ParseInLocation(DateTimeLayout, t, local)
}

package date_utils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

//GetNow : returns the time in standard format
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString : returns the time in normal layout
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

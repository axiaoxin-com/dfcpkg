// Package dfcpkg is a datetime format converter.
// Convert any time string to a timestamp or vice versa.
package dfcpkg

import (
	"fmt"
	"time"

	"github.com/axiaoxin-com/dateparse"
)

// FormatLayouts return all formats supported when converting timestamps to datetime strings
var FormatLayouts []string = []string{
	"2006-01-02 15:04:05",
	"2006-01-02 15:04:05.999",
	"2006-01-02 15:04:05.999999",
	"2006-01-02 15:04:05.999999999",
	"2006/01/02 15:04:05.999999999",
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
}

// ResultOfLayout returns the datetime strings of the time parameter according to the FormatLayouts list
func ResultOfLayout(t time.Time) []string {
	result := []string{}
	for _, layout := range FormatLayouts {
		result = append(result, t.Format(layout))
	}
	return result
}

// ResultOfTimestamp Returns the timestamp of string type,
// including second timestamp, millisecond timestamp, microsecond timestamp and nanosecond timestamp
func ResultOfTimestamp(t time.Time) []string {
	nano := t.UnixNano()
	return []string{
		fmt.Sprint(t.Unix()),                       // 秒时间戳
		fmt.Sprint(nano / int64(time.Millisecond)), // 毫秒时间戳
		fmt.Sprint(nano / int64(time.Microsecond)), // 微秒时间戳
		fmt.Sprint(nano),                           // 纳秒时间戳
	}
}

// Convert converts any incoming string into timestamps or datetime strings in the specified location
// If no location is specified, the location of the current machine's local is used
func Convert(s string, location ...*time.Location) ([]string, error) {
	loc := time.Local
	if len(location) > 0 {
		loc = location[0]
	}

	t, dateState, err := dateparse.ParseIn(s, loc)
	if err != nil {
		return nil, err
	}

	result := []string{}
	switch dateState {
	// The datetime string of semantic type outputs both formated datetime strings and timestamps
	case dateparse.StateNow, dateparse.StateHowLongAgo:
		result = append(result, ResultOfLayout(t)...)
		result = append(result, ResultOfTimestamp(t)...)
	// Passing in a timestamp returns datetime strings
	case dateparse.StateTimestamp:
		result = ResultOfLayout(t)
	// Passing in a datetime string returns the timestamps
	default:
		result = ResultOfTimestamp(t)
	}
	return result, nil
}

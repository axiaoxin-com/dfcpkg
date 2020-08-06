package dfc

import (
	"fmt"
	"testing"
	"time"
)

func TestResultOfLayout(t *testing.T) {
	var testDatetimeString = "2020-08-06 17:29:29.292929"
	x, _ := time.Parse("2006-01-02 15:04:05.999999", testDatetimeString)
	result := ResultOfLayout(x)
	for i, s := range result {
		if s != x.Format(FormatLayouts[i]) {
			t.Error(s, " != ", x.Format(FormatLayouts[i]))
		}
	}
}

func TestResultOfTimestamp(t *testing.T) {
	var testDatetimeString = "2020-08-06 17:29:29.292929"
	x, _ := time.Parse("2006-01-02 15:04:05.999999", testDatetimeString)
	nano := x.UnixNano()
	result := ResultOfTimestamp(x)
	exp := []string{
		fmt.Sprint(x.Unix()),                       // 秒时间戳
		fmt.Sprint(nano / int64(time.Millisecond)), // 毫秒时间戳
		fmt.Sprint(nano / int64(time.Microsecond)), // 微秒时间戳
		fmt.Sprint(nano),                           // 纳秒时间戳
	}
	for i, s := range result {
		if s != exp[i] {
			t.Error(s, " != ", exp[i])
		}
	}
}

func TestConvertDatestring(t *testing.T) {
	var testDatetimeString = "2020-08-06 17:29:29.292929"
	x, _ := time.ParseInLocation("2006-01-02 15:04:05.999999", testDatetimeString, time.Local)
	result, err := Convert(testDatetimeString)
	if err != nil {
		t.Fatal(err)
	}

	nano := x.UnixNano()
	exp := []string{
		fmt.Sprint(x.Unix()),                       // 秒时间戳
		fmt.Sprint(nano / int64(time.Millisecond)), // 毫秒时间戳
		fmt.Sprint(nano / int64(time.Microsecond)), // 微秒时间戳
		fmt.Sprint(nano),                           // 纳秒时间戳
	}
	for i := 0; i < len(exp); i++ {
		if result[i] != exp[i] {
			t.Fatal(result[i], " != ", exp[i])
		}
	}
}

func TestConvertTimestamp(t *testing.T) {
	var testTimestamp int64 = 1596706169
	x := time.Unix(testTimestamp, 0)
	result, err := Convert(fmt.Sprint(testTimestamp))
	if err != nil {
		t.Fatal(err)
	}
	for i, s := range result {
		if s != x.Format(FormatLayouts[i]) {
			t.Fatal(s, " != ", x.Format(FormatLayouts[i]))
		}
	}
}

func TestConvertNow(t *testing.T) {
	now := time.Now()
	testDatetimeString := "now"
	result, err := Convert(testDatetimeString)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(FormatLayouts); i++ {
		if result[i] != now.Format(FormatLayouts[i]) {
			t.Fatal(result[i], " != ", now.Format(FormatLayouts[i]))
		}
	}
	nano := now.UnixNano()
	exp := []string{
		fmt.Sprint(now.Unix()),                     // 秒时间戳
		fmt.Sprint(nano / int64(time.Millisecond)), // 毫秒时间戳
		fmt.Sprint(nano / int64(time.Microsecond)), // 微秒时间戳
		fmt.Sprint(nano),                           // 纳秒时间戳
	}
	for i := len(FormatLayouts); i < len(FormatLayouts)+len(exp); i++ {
		if result[i] != exp[i-len(FormatLayouts)] {
			t.Fatal(result[i], " != ", now.Format(FormatLayouts[i]))
		}
	}
}

func TestConvertAgo(t *testing.T) {
	testDatetimeString := "1 hours ago"
	result, err := Convert(testDatetimeString)
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now().Add(-time.Hour * 1)
	for i := 0; i < len(FormatLayouts); i++ {
		if result[i] != now.Format(FormatLayouts[i]) {
			t.Fatal(result[i], " != ", now.Format(FormatLayouts[i]))
		}
	}
	nano := now.UnixNano()
	exp := []string{
		fmt.Sprint(now.Unix()),                     // 秒时间戳
		fmt.Sprint(nano / int64(time.Millisecond)), // 毫秒时间戳
		fmt.Sprint(nano / int64(time.Microsecond)), // 微秒时间戳
		fmt.Sprint(nano),                           // 纳秒时间戳
	}
	for i := len(FormatLayouts); i < len(FormatLayouts)+len(exp); i++ {
		if result[i] != exp[i-len(FormatLayouts)] {
			t.Fatal(result[i], " != ", now.Format(FormatLayouts[i]))
		}
	}
}

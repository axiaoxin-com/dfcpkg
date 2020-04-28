package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/axiaoxin-com/dateparse"
)

// GetTimeFromCmd 解析命令行参数中的时间，返回golang的时间对象
func GetTimeFromCmd() (time.Time, dateparse.DateState) {
	argsLen := len(os.Args)
	if argsLen < 2 {
		log.Fatal("Invalid Args. Please Run as dfc <time>")
	}
	timeArg := strings.Join(os.Args[1:argsLen], " ")
	// 解析参数为时间对象
	t, dateState, err := dateparse.ParseLocal(timeArg)
	if err != nil {
		log.Fatal("Parse time error:", err)
	}
	return t, dateState
}

// GetFormatLayouts 时间戳转字符串时展示的所有格式
func GetFormatLayouts() (layouts []string) {
	layouts = []string{
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
	return
}

// LayoutPrint 按layout列表打印时间字符串
func LayoutPrint(t time.Time) {
	for _, layout := range GetFormatLayouts() {
		fmt.Println(t.Format(layout))
	}
}

// TimestampPrint 打印秒时间戳、毫秒时间戳、微秒时间戳、纳秒时间戳
func TimestampPrint(t time.Time) {
	nano := t.UnixNano()
	fmt.Println(t.Unix())                       // 秒时间戳
	fmt.Println(nano / int64(time.Millisecond)) // 毫秒时间戳
	fmt.Println(nano / int64(time.Microsecond)) // 微秒时间戳
	fmt.Println(nano)                           // 纳秒时间戳
}

func main() {
	t, dateState := GetTimeFromCmd()
	switch dateState {
	case dateparse.StateNow, dateparse.StateHowLongAgo:
		LayoutPrint(t)
		TimestampPrint(t)
	case dateparse.StateTimestamp:
		LayoutPrint(t)
	default:
		TimestampPrint(t)
	}
}

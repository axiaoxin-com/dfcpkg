// package main
// a commandline tool of datetime format convertor
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/axiaoxin-com/dfc"
)

func main() {
	argsLen := len(os.Args)
	if argsLen < 2 {
		log.Fatal("Invalid Args. Please Run as dfc <time>")
	}
	timeArg := strings.Join(os.Args[1:argsLen], " ")
	// 解析参数为时间对象
	result, err := dfc.Convert(timeArg)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range result {
		fmt.Println(s)
	}
}

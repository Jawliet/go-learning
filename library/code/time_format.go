package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// "2006-01-02 15:04:05"->1月2号下午3时4分5秒 06年 7时区(1234567)
	// Jan 2 15:04:05 2006 MST
	// 1   2  3  4  5    6  -7
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

}

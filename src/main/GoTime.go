package main

import "time"
import "fmt"

func main() {
	var t time.Time // 定义 time.Time 类型变量
	t = time.Now()  // 获取当前时间
	fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t, t.Location(), t)
	// 时间: 2017-02-22 09:06:05.816187261 +0800 CST, 时区:  Local,  时间类型: time.Time

	// time.UTC() time 返回UTC 时区的时间
	fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t.UTC(), t.UTC().Location(), t)
	// 时间: 2017-02-22 01:07:15.179280004 +0000 UTC, 时区:  UTC,  时间类型: time.Time
}

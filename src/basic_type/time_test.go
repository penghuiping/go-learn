package basic_type

import (
	"fmt"
	"time"
)

//格式化日期必须是下面这种格式，2006-01-02 15:04:02时间是go语言诞生的日子
const (
	date        = "2006-01-02"
	shortdate   = "06-01-02"
	times       = "15:04:05"
	shorttime   = "15:04"
	datetime    = "2006-01-02 15:04:05"
	newdatetime = "2006/01/02 15~04~05"
	newtime     = "15~04~02"
)

func DoActionTime() {
	//格式化时间
	t := time.Now()
	fmt.Println(t.Format(date))
	fmt.Println(t.Format(datetime))

	//计算三天后的时间
	tt := t.AddDate(0, 0, -3)
	fmt.Println(tt.Format(datetime))
}

package main

import(
	"fmt"
	"time"
)

func main() {
	var t time.Time
	t = time.Now()

	fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t, t.Location(), t)

	var dur time.Duration
	dur = 120000000000
	fmt.Printf(dur.String())
}

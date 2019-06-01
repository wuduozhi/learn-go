package main

import (
	"fmt"
	"time"
)

func main()  {
    now := time.Now()

    fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Clock())

	// hour, min, sec := now.Clock()

	

}
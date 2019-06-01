package main

import (
	"fmt"
	"errors"
)

func main(){
	test_panic()
}

func test_panic(){
	defer func(){
		fmt.Println("Finish")
		if err := recover();err != nil{
			fmt.Println(err)
		}
	}()
	fmt.Println("I love you")
	panic(errors.New("I miss you"))
}
package main

import (
	"bytes"
	"fmt"
	"encoding/base64"
	"runtime"
)

func main() {
	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))

	b, _ := base64.StdEncoding.DecodeString("V3VkdW96aGlAcXEuY29t")
	fmt.Println(b)

	call()
	call()
}

func call() {
    var calldepth = 1;
    fmt.Println(runtime.Caller(calldepth))
}
package main

import (
	"io"
	"log"
	"os"
	"strings"
	"github.com/tdewolff/parse/buffer"
)

func main() {
	r := strings.NewReader("some io.Reader stram to be read\n")
	lr := io.MultiReader(
		io.LimitReader(r, 4),
		buffer.NewReader([]byte("wuduozhi.\r\n")),
	)

	if _,err := io.Copy(os.Stdout,lr);err != nil{
		log.Fatal(err)
	}

}
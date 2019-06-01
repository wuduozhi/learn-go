package main

import (
	"fmt"
	"path"
	"strconv"
)

func main() {
	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
	}

	for _,p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}

	fmt.Println(path.Base("/a/b"))
	fmt.Println(path.Dir("/a/b/c"))
	fmt.Println(path.Ext("/a/b/c/bar.css"))
	fmt.Println(path.IsAbs("/dev/null"))

	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))
	fmt.Println(path.Join("a/b", "/c"))

	fmt.Println(path.Split("static/myfile.css"))

	rst := make([]byte, 0)
	rst = strconv.AppendBool(rst, 0 < 1)
	fmt.Printf("%s\n", rst) // true
	rst = strconv.AppendBool(rst, 0 > 1)
	fmt.Printf("%s\n", rst) // truefalse

	fmt.Println(strconv.ParseInt("123", 10, 8))
	// 123
	fmt.Println(strconv.ParseInt("12345", 10, 8))
	// 127 strconv.ParseInt: parsing "12345": value out of range
	fmt.Println(strconv.ParseInt("2147483647", 10, 0))
	// 2147483647
	fmt.Println(strconv.ParseInt("0xFF", 16, 0))
	// 0 strconv.ParseInt: parsing "0xFF": invalid syntax
	fmt.Println(strconv.ParseInt("FF", 16, 0))
	// 255
	fmt.Println(strconv.ParseInt("0xFF", 0, 0))
	// 255

	s := "0.12345678901234567890"
	f, err := strconv.ParseFloat(s, 32)
	fmt.Println(f, err)          // 0.12345679104328156
	fmt.Println(float32(f), err) // 0.12345679
	f, err = strconv.ParseFloat(s, 64)
	fmt.Println(f, err) // 0.12345678901234568
}
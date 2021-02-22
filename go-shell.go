package main

import "os"

func main()  {
	str, _ := os.Getwd()	// _ 是空白值，可以不使用

	println( str)
}

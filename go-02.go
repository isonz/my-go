package main

import (
	"fmt"
	"time"
)

const ABC = "abc"		//可以不使用

func main()  {
	var age int		//变量的声明必须使用空格隔开
	name := "Ken"	//使用冒号定义字符串变量并赋值
	var b, c = 1, 2
	var a *int		// 类型为 nil：

	//这种不带声明格式的只能在函数体中出现
	g, h := 123, "hello"

	println(ABC)
	println(age)
	println(name)
	println(b, c)
	println(a)
	println(g, h)


	// %d 表示整型数字，%s 表示字符串
	var stockCode=123
	var endDate=time.Now()
	var targetUrl=fmt.Sprintf("Code=%d&endDate=%s", stockCode, endDate)
	println(targetUrl)

}

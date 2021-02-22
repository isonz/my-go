/**
条件语句，循环语句
 */
package main

import (
	"fmt"
	"time"
)

func main()  {
	findSuShu()
	switchFunc()
	selectFunc()
	forFunc()
	gotoFunc()
	gotoTag(); print9x()
}

// 寻找到 100 以内的所有的素数
func findSuShu()  {
	var count int
	var flag bool
	count=1
	//while(count<100) {}    //go没有while
	for count < 100 {
		count++
		flag = true
		//注意tmp变量  :=
		for tmp:=2;tmp<count;tmp++ {
			if count%tmp==0{
				flag = false
			}
		}

		// 每一个 if else 都需要加入括号 同时 else 位置不能在新一行
		if flag == true {
			println(count,"素数")
		}else{
			continue
		}
	}
}

// switch 使用, 不同的 case 之间不使用 break 分隔，默认只会执行一个 case。
// 如果想要执行多个 case，需要使用 fallthrough 关键字，也可用 break 终止。
func switchFunc()  {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50,60,70 : grade = "C"
	default: grade = "D"
	}

	switch {
	case grade == "A" :
		fmt.Printf("优秀!\n" )
	case grade == "B", grade == "C" :
		fmt.Printf("良好\n" )
	case grade == "D" :
		fmt.Printf("及格\n" )
	case grade == "F":
		fmt.Printf("不及格\n" )
	default:
		fmt.Printf("差\n" )
	}
	fmt.Printf("你的等级是 %s\n", grade )


	// switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T",i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型" )
	default:
		fmt.Printf("未知型")
	}


	// 使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}


	// 支持多条件匹配
	a := 1
	switch a{
		case 1,2,3,4:
			println("1,2,3,4")
		default:
			println("default")
	}
}

// select 会循环检测条件，如果有满足则执行并退出，否则一直循环检测。
func selectFunc()  {
	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go Chann(ch, stopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("Receive", c)
			fmt.Println("channel")
		case s := <-ch:
			fmt.Println("Receive", s)
		case _ = <-stopCh:
			goto end
		}
	}
end:
}

func Chann(ch chan int, stopCh chan bool) {
	var i int
	i = 10
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}


func forFunc()  {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	/* 定义局部变量 */
	var a = 10
	/* for 循环 */
	for a < 20 {
		if a == 15 {
			/* 跳过此次循环 */
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
}

func gotoFunc()  {
	/* 定义局部变量 */
	var a int = 10
	/* 循环 */
	LOOP: for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a ---- 的值为 : %d\n", a)
		a++
	}


}

//for循环配合goto打印九九乘法表
func gotoTag() {
	for m := 1; m < 10; m++ {
		n := 1
		LOOP: if n <= m {
			fmt.Printf("%dx%d=%d ",n,m,m*n)
			n++
			goto LOOP
		} else {
			fmt.Println("")
		}
		n++
	}
}

//嵌套for循环打印九九乘法表
func print9x() {
	for m := 1; m < 10; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%dx%d=%d ",n,m,m*n)
		}
		fmt.Println("")
	}
}




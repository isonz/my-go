/**
递归函数, 类型转换, 接口

递归，就是在运行的过程中调用自己。
接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
*/

package main

import "fmt"

func main()  {

	var i = 5
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	var j int
	for j = 0; j < 10; j++ {
		fmt.Printf("%d\t", fibonacci(j))
	}

	println()

	// 类型转换用于将一种数据类型的变量转换为另外一种类型的变量。
	var sum = 17
	var count = 5
	var mean float32
	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)


	println()

	// go 不支持隐式转换类型，比如 :
	/*
	var a int64 = 3
	var b int32
	b = a				// 此时会报错, cannot use a (type int64) as type int32 in assignment, cannot use b (type int32) as type string in argument to fmt.Printf
	fmt.Printf("b 为 : %d", b)
	 */

	// 但是如果改成 b = int32(a) 就不会报错了:
	var a int64 = 3
	var b int32
	b = int32(a)
	fmt.Printf("b 为 : %d", b)


	println()
	println()

	// 接口
	/**
	定义了一个接口Phone，接口里面有一个方法call()。
	然后我们在main函数里面定义了一个Phone类型变量，并分别为之赋值为NokiaPhone和IPhone。
	 */
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()

}


// 阶乘
func Factorial(n uint64)(result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}


// 斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}



// 接口
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}


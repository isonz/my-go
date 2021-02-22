/**
函数, 变量作用域
*/
package main

import (
	"fmt"
	"math"
)

/* 声明全局变量 */
var g = 20

func main()  {

	/* 声明局部变量 */
	var g = 10
	fmt.Printf ("结果： g = %d\n",  g)

	/* 定义局部变量 */
	var a = 100
	var b = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)
	println(fmt.Sprintf("最大值是 : %d", ret ))

	a, b = b, a
	println(fmt.Sprintf("a : %d, b : %d", a, b ))

	swap2(&a, &b)
	fmt.Printf("交换后，a 的值 : %d\n", a )
	fmt.Printf("交换后，b 的值 : %d\n", b )

	c, d := swap1("Google", "Apple")
	println(c, d)


	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))


	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调2，x：%d\n", x)
		return x
	})


	// 闭包
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()
	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	println(nextNumber(), nextNumber(), nextNumber())


	// 方法
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())

}


/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// Go 函数可以返回多个值
func swap1(x, y string) (string, string) {
	return y, x
}


/**
 * &a 指向 a 指针，a 变量的地址
 * &b 指向 b 指针，b 变量的地址
 */
func swap2(x *int, y *int) {
	var temp int
	temp = *x    /* 保存 x 地址上的值 */
	*x = *y      /* 将 y 值赋给 x */
	*y = temp    /* 将 temp 值赋给 y */
}


// 声明一个函数类型
type cb func(int) int
func testCallBack(x int, f cb) {
	f(x)
}
func callBack(x int) int {
	fmt.Printf("我是回调1，x：%d\n", x)
	return x
}


//闭包
func getSequence() func() int {		// 与return同步返回类型
	i:=0
	return func() int {			// 返回类型为 func() int
		i+=1
		return i
	}
}


// 方法
// Java, C++ 等语言中，实现类的方法做法都是编译器隐式的给函数加一个 this 指针，而在 Go 里，这个 this 指针需要明确的申明出来，其实和其它 OO 语言并没有很大的区别。
/** C++：
class Circle {
  public:
    float getArea() {
       return 3.14 * radius * radius;
    }
  private:
    float radius;
}
// 其中 getArea 经过编译器处理大致变为
float getArea(Circle *const c) {
  ...
}
 */
/* 定义结构体 */
type Circle struct {
	radius float64
}
//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}



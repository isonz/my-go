/**
指针 —— 存储地址的东西
&a == ip, *ip == a
*/
package main

import "fmt"

const MAX int = 3

func main()  {
	var a int= 20   /* 声明实际变量 */
	var ip *int     /* 声明指针变量 */
	ip = &a  		/* 指针变量的存储地址 */

	// &a == ip, *ip == a
	fmt.Printf("a 变量的地址是: %x\n", &a)
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)


	// 当一个指针被定义后没有分配到任何变量时，它的值为 nil。
	// if(ptr != nil) ptr 不是空指针    if(ptr == nil)  ptr 是空指针
	var  ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr  )
	ptr = &a
	fmt.Printf("ptr 的值为 : %x\n", ptr  )

	// 指针数组
	aa := []int{10,100,200}
	var i int
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, aa[i] )
	}
	var ptrArr [MAX]*int
	for  i = 0; i < MAX; i++ {
		ptrArr[i] = &aa[i] 		// 整数地址赋值给指针数组
	}
	for  i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptrArr[i] )
	}

	// 指向指针的指针, 可以是多重指针，2重，3重，多重，几重加几个 *
	// 如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
	//当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：
	var b int
	var ptrb *int
	var pptrb **int
	b = 3000
	/* 指针 ptrb 地址 */
	ptrb = &b
	/* 指向指针 ptrb 地址 */
	pptrb = &ptrb
	/* 获取 pptrb 的值 */
	fmt.Printf("变量 b = %d\n", b )
	fmt.Printf("指针变量 *ptrb = %d, ptrb = %x \n", *ptrb, ptrb)
	fmt.Printf("指向指针的指针变量 **pptrb = %d, pptrb = %x \n", **pptrb, pptrb)

	// 向函数传递指针
	/* 定义局部变量 */
	var e int = 100
	var f int= 200
	fmt.Printf("交换前 e 的值: %d, f 的值 : %d\n", e, f )
	/* 调用函数用于交换值
	 * &e 指向 e 变量的地址
	 * &f 指向 f 变量的地址
	 */
	swap3(&e, &f)
	fmt.Printf("交换后 e 的值: %d, f 的值 : %d\n", e, f )
	swap4(&e, &f)
	fmt.Printf("交换后 e 的值: %d, f 的值 : %d\n", e, f )
}

// 向函数传递指针, 同一个 main 包，go-04.go 中已经有了 swap1， swap2，这里不可再定义
func swap3(x *int, y *int) {		// x,y 分别是地址，指向整数型变量
	var temp int
	temp = *x    /* 保存 x 地址的值 */
	*x = *y      /* 将 y 赋值给 x */
	*y = temp    /* 将 temp 赋值给 y */
}

// 向函数传递指针
func swap4(x *int, y *int) {		// x,y 分别是地址，指向整数型变量
	*x, *y = *y, *x
}
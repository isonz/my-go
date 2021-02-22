/**
数组, 多维数组，向函数传递数组
*/
package main

import "fmt"

func main()  {
	// 初始化数组
	balance1 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}		//固定长度
	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}	//不确定长度
	balance3 := []float32{1000.0, 2.0, 3.4, 7.0, 50.0}		//不确定长度
	balance4 := [6]float32{1:2.0,3:7.0}		//  固定长度下将索引为 1 和 3 的元素初始化,非固定长度下不可这样设置

	println(len(balance1), len(balance2), len(balance3), len(balance4))
	println(balance1[0], balance2[1], balance3[2], balance4[3])

	/* 输出数组元素 */
	for i := 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance1[i] )
	}

	// 二维数组中的元素可通过 a[ i ][ j ] 来访问。
	// Step 1: 创建数组
	var values [][]int
	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)
	// Step 3: 显示两行数据
	fmt.Println("Row 1", values[0])
	fmt.Println("Row 2", values[1])
	// Step 4: 访问第一个元素
	fmt.Println("二维数组的第一行第一个元素为：", values[0][0])
	fmt.Println("二维数组的第二行第一个元素为：", values[1][0])


	aa := [3][4]int{
		{0, 1, 2, 3} ,   /*  第一行索引为 0 */
		{4, 5, 6, 7} ,   /*  第二行索引为 1 */
		{8, 9, 10, 11},   /* 第三行索引为 2 */
	}
	// 注意：以上代码中倒数第二行的 } 必须要有逗号，因为最后一行的 } 不能单独一行，也可以写成这样：
	bb := [3][4]int{
		{0, 1, 2, 3} ,   /*  第一行索引为 0 */
		{4, 5, 6, 7} ,   /*  第二行索引为 1 */
		{8, 9, 10, 11}}   /* 第三行索引为 2 */
	fmt.Println(aa)		// println(aa)会报错
	fmt.Println(bb)		// println(bb)会报错


	// 创建二维数组
	sites := [2][2]string{}
	// 向二维数组添加元素
	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"
	// 显示结果
	fmt.Println(sites)


	// 创建各个维度元素数量不一致的多维数组：
	// 创建空的二维数组
	var animals [][]string

	// 创建三个一维数组，各数组长度不同
	row3 := []string{"fish", "shark", "eel"}
	row4 := []string{"bird"}
	row5 := []string{"lizard", "salamander"}
	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, row3)
	animals = append(animals, row4)
	animals = append(animals, row5)
	// 循环输出
	for i := range animals {
		fmt.Printf("Row %v:%s  ", i, animals[i])
	}
	println()

	// 向函数传递数组
	/* 数组长度为 5 */
	var  balance = [5]int {1000, 2, 3, 17, 50}
	var avg float32
	/* 数组作为参数传递给函数 */
	avg = getAverage( balance, 5)
	/* 输出返回的平均值 */
	fmt.Printf( "平均值为: %f \n", avg)
}



func getAverage(arr [5]int, size int) float32 {
	var i,sum int
	var avg float32

	for i = 0; i < size;i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
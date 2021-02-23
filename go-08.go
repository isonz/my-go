/**
切片(Slice) - 数组的长度不可改变，切片的长度是不固定的，不需要说明长度, 可以追加元素，在追加时可能使切片的容量增大。
*/
package main

import "fmt"

func main()  {
	var numbers = make([]int,3,5)  // sliceName []T = make([]T, length, capacity)
	printSlice(numbers)


	var numbers1 []int
	printSlice(numbers1)
	if(numbers1 == nil){
		fmt.Printf("切片是空的 \n")
	}


	/* 创建切片 */
	numbers2 := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers2)

	/* 打印原始切片 */
	fmt.Println("numbers2 ==", numbers2)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers2[1:4] ==", numbers2[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers2[:3] ==", numbers2[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers2[4:] ==", numbers2[4:])

	numbers3 := make([]int,0,5)
	printSlice(numbers3)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number4 := numbers2[:2]
	printSlice(number4)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number5 := numbers2[2:5]
	printSlice(number5)


	println()

	var numbers6 []int
	printSlice(numbers6)

	/* 允许追加空切片 */
	numbers6 = append(numbers6, 0)
	printSlice(numbers6)

	/* 向切片添加一个元素 */
	numbers6 = append(numbers6, 1)
	printSlice(numbers6)

	/* 同时添加多个元素 */
	numbers6 = append(numbers6, 2,3,4)
	printSlice(numbers6)

	/* 创建切片 numbers7 是之前切片的两倍容量*/
	numbers7 := make([]int, len(numbers6), (cap(numbers6))*2)

	/* 拷贝 numbers6 的内容到 numbers7 */
	copy(numbers7,numbers6)
	printSlice(numbers7)


}

// len() 方法获取长度， cap() 可以测量切片最长可以达到多少。
func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
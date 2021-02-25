/**
并发, 通道（channel）

Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。
同一个程序中的所有 goroutine 共享同一个地址空间。


通道（channel）是用来传递数据的一个数据结构。
通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。

通道缓冲区-- 通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小： ch := make(chan int, 100)
带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。

如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。
如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；
如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。

*/

package main

import (
	"fmt"
	"time"
)

func main()  {

	// 并发
	go say("world")
	say("hello")


	// 使用通道
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x, y, x+y)


	// 通道缓冲区, 通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小： ch := make(chan int, 100)
	// 这里我们定义了一个可以存储整数类型的带缓冲通道 缓冲区大小为2
	ch := make(chan int, 2)
	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据, 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2
	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)



	// Go 遍历通道与关闭通道， 通过 range 关键字来实现遍历读取到的数据，类似于与数组或切片。
	d := make(chan int, 10)		//缓冲区大小为 10
	go fibonacci1(cap(d), d)	// cap(d) == 10
	// range 函数遍历每个从通道接收到的数据，因为 d 在发送完 10 个数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据之后就结束了。
	// 如果上面的 d 通道不关闭，那么 range 函数就不会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range d {
		fmt.Println(i)
	}

}


func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 d
}


func fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
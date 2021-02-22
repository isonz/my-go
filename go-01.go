//必须要有 main 包，否则会报错误：cannot run non-main package
package main
// package my_go	// 看底部的解释

import (
	"runtime"
	"unicode/utf8"
) // 导入外包

/* 这是程序入口函数 */
func main()  {		// { 不能单独放在一行
	println(runtime.Version())				// go1.16
	println("Hello Golang...");		// 可以使用分号，但不推荐使用。 如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中我们并不鼓励这种做法。
	println("Hello" + " World...")	// 与 Print("hello, world\n") 可以得到相同的结果。

	// 计算UTF8的字符串长度
	str := "hello 世界"		// 理论上是 5+1+2 = 8
	println("len(str):", len(str))	// 实际上是 5+1+3+3 , 汉字占3个字节
	println("RuneCountInString:", utf8.RuneCountInString(str))
	println("rune:", len([]rune(str)))	// 将字符串转换为rune类型的数组再计算长度
	println("A")
	println('A')	//  打印出65，ASCII 码值
	// println('AB')	//  报错，不能打印2个 rune
}

/**
当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，
那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），
这被称为导出（像面向对象语言中的 public）；
标识符如果以小写字母开头，则对包外是不可见的，
但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
 */

/**
单引号，双引号，反引号：
Go语言的字符串类型string在本质上就与其他语言的字符串类型不同：
Java的String、C++的std::string以及Python3的str类型都只是定宽字符序列
Go语言的字符串是一个用UTF-8编码的变宽字符序列，它的每一个字符都用一个或多个字节表示
即：一个Go语言字符串是一个任意字节的常量序列。
Golang的双引号和反引号都可用于表示一个常量字符串，不同在于：
双引号用来创建可解析的字符串字面量(支持转义，但不能用来引用多行)，
	其实质是一个byte类型的数组，str:="hello world" fmt.Println(str[3:5]),这时的输出结果为“lo”.
	但如果将打印的内容改为“str[3]”，则会输出108。
反引号用来创建原生的字符串字面量，这些字符串可能由多行组成(不支持任何转义序列)，原生的字符串字面量多用于书写多行消息、HTML以及正则表达式
而单引号则用于表示Golang的一个特殊类型：rune，类似其他语言的byte但又不完全一样，是指：码点字面量（Unicode code point），不做任何转义的原始内容。
	rune： int32的别名，几乎在所有方面等同于int32，它用来区分字符值和整数值
 */

/**
关于包，根据本地测试得出以下几点：
  1、文件名与包名没有直接关系，不一定要将文件名与包名定成同一个。
  2、文件夹名与包名没有直接关系，并非需要一致。
  3、同一个文件夹下的文件只能有一个包名，否则编译报错。
文件结构:

Test
--helloworld.go

myMath
--myMath1.go
--myMath2.go

测试代码:
// helloworld.go
package main
import (
"fmt"
"./myMath"
)

func main(){
    fmt.Println("Hello World!")
    fmt.Println(mathClass.Add(1,1))
    fmt.Println(mathClass.Sub(1,1))
}


// myMath1.go
package mathClass
func Add(x,y int) int {
    return x + y
}


// myMath2.go
package mathClass
func Sub(x,y int) int {
    return x - y
}


 */
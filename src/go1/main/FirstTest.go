package main

import "fmt" // 导入功能包，实际是导入该目录名下面的.go文件，在使用包中的方法时，需要通过go文件package定义的包名
// import . "math"  // 使用 . 表明在使用该目录下的方法时可以省略包名，直接使用方法名调用
// import f "fmt"  // 使用发f作为fmt包的别名，这样在调用该包下的方法时可以使用f来调用
// import _ "fmt"  // _ 操作是引入某个包，但不直接使用包里的函数，而是调用该包里面的init函数

const ( // 使用const块代码声明常量
	NUM int = 100             // 显式指定常量的数据类型
	STR     = "Hello GoLand!" // 隐式通过赋值的常量值字面量确定常量的数据类型
)

func main() { // 必须在main包中的main方法作为程序入口
	fmt.Println("Hello World!")
	fmt.Println(STR)
	fmt.Println(NUM)
	a := 1        // 声明简式变量，简式变量只能在函数中声明使用
	var b int = 2 // 声明变量并显式指定变量数据类型
	var c = 4     // 声明变量，隐式通过赋值的类型确定变量的数据类型
	fmt.Println(a, b, c)
	if a := 6; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		fmt.Println(a, b, c)
	}
}

func init() {
	fmt.Println("init func....")
}

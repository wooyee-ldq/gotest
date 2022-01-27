package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

const PI = 3.1415926

func main() {
	fmt.Println("运算符....")
	fmt.Println("内置运算符：")
	a := 1 + 1
	b := a + 2
	c := b - 1
	fmt.Println("c:", c)
	d := 0xffff
	e := 0x0000
	f1 := d & e  // 按位与
	f2 := d | e  // 按位或
	f3 := d ^ e  // 按位异或
	f4 := d << 2 // 左移
	f5 := ^b     // 按位取反
	fmt.Println(f1, f2, f3, f4, f5)
	// 定义一个int类型指针
	var p *int = &a // 和c中一样取变量a的地址，也就是返回一个指针
	fmt.Println("a=", a)
	a++
	fmt.Println("a=", a)
	*p = 4
	fmt.Println("a=", a)
	fmt.Println("a变量的地址是：", p)
	// 定义一个string类型指针
	var ps *string
	str := "hello goland!"
	ps = &str
	fmt.Println(*ps)
	fmt.Println("-------------------")
	type myByte uint8 // 适用type定义myByte作为uint8类型别名，也可以这样写：type myByte = uint8
	// 内置的byte,rune类型实际也是别名：
	// type byte = uint8
	// type rune = int32
	var bt myByte = 10
	fmt.Println(bt)
	// 适用[]索引获取字符串中的字节
	fmt.Println(str[0])
	*ps = "你好，世界！"
	fmt.Printf("%c\n", str[0]) // 字符串在使用下标索引获取内容时，获取的是该下标索引位置的字节对应数值(0-255)
	fmt.Println(len(str))      // len(str)方法获取str字符串的字节长度，因为go使用的utf-8编码，中文一个字符占用3个字节
	var arr [6]rune = [6]rune{}
	arr[0] = 104
	fmt.Println(arr)
	var bl bool = true
	if bl {
		fmt.Printf("%c", arr[0])
	}
	fmt.Println("--------------------")
	var buf bytes.Buffer
	buf.WriteString("测")
	buf.WriteString("试")
	fmt.Println(buf.String())
	var bd strings.Builder
	bd.WriteString("abc")
	bd.WriteString("de")
	fmt.Println(bd.String())
	fmt.Println("--------------------------")
	fmt.Println("数组:")
	var arrAge = [5]int{18, 20, 15, 22, 16}
	var arrName = [5]string{3: "Chris", 4: "Ron"} //指定索引位置初始化
	// {"","","","Chris","Ron"}
	fmt.Println(arrAge, arrName)
	var arrLazy = [...]int{1, 2, 3, 4} // 根据初始化元素个数确定数组大小
	var arrRoom [4]int
	var arrBed = new([2]int) // 使用new创建的数组，返回的数组名是对应数组的指针类型
	var arrp *[2]int = arrBed
	arrp[0] = 10
	arrp[1] = 20
	arrRoom[0] = 4
	fmt.Println(arrLazy)
	fmt.Println("cap:", cap(arrLazy)) // cap()方法计算数组容量
	fmt.Println(arrRoom, *arrp)
	arrHandle(arrLazy[0:len(arrLazy)])
	fmt.Println("\n----------------------")
	type char int32
	var ch char = 'd'
	var ch1 int8 = '\u0041' // go语言中使用utf-8的unicode编码，一个字符占至少用2个字节
	// \u开头只能接4位16进制数，也就是16位2进制数，即是16bit大小
	// \U开头的只能接8位16进制数，也就是32位2进制数，即是32bit大小
	// 所以go中的字符可以使用int16或者int表示
	fmt.Printf("%d:%c\t%d:%c", ch, ch, ch1, ch1)
	fmt.Println("\n-------------------")
	fmt.Println("切片：")
	// 切片（slice） 是对底层数组一个连续片段的引用（该数组我们称之为相关数组，通常是匿名的），所以切片是一个引用类型（和数组不一样）。
	// 切片没初始化前是nil
	// 切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中切片比数组更常用。
	var arrslice1 []int = arrLazy[0:2] // 定义并初始化一个切片
	var arrslice2 = []int{4, 3, 2, 1}  // 也可以这样定义并初始化一个切片
	fmt.Println(arrslice1, arrslice2)
	// 使用 make() 函数来创建一个切片，同时创建好相关数组：var slice1 []type = make([]type, len,cap)
	// type是切片关联数组的类型，len是切片的长度，cap是切片的容量也就是数组的长度/容量
	var arrslice []int = make([]int, 2, 4) // 适用make()创建数组并定义初始化数组
	fmt.Println(arrslice)
	// 通过改变切片长度得到新切片的过程称之为切片重组 reslicing，做法如下：slice1 = slice1[0:end]，其中 end 是新的末尾索引（即长度）。
	fmt.Println(arrslice2, len(arrslice2), cap(arrslice2))
	arrslice2 = arrslice2[1:4] // 切片重组
	fmt.Println(arrslice2, len(arrslice2), cap(arrslice2))
	// 切片重组后新旧切片还是关联原来的同一个数组，需要注意防止因为新切片占用小部分数组数据导致数组无法被回收GC
	// 使用内置函数copy()，拷贝数据（而不是重新划分slice）到新切片。
	copy(arrslice2, arrslice1[0:2]) // 拷贝复制数据到新的切片，两个切片关联的数组不同
	// append()函数将 0 个或多个具有相同类型 S 的元素追加到切片s后面并且返回新的切片；
	//追加的元素必须和原切片的元素同类型。
	//如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。
	// 这样会导致返回的新切片可能已经关联到一个新的数组，也就是返回的新切片和原切片关联的数组不一样
	// 在append添加数据的时候，如果切片的容量不足，会创建一个原切片容量2倍大的数组，把原数据复制到数组后，把添加的数据添加到新数组后面
	newslice := append(arrslice2, 10) // 创建了新的数组，新数组大小是原来切片容量的2倍大小
	fmt.Println(newslice, len(newslice), cap(newslice))
	fmt.Println("---------------------------")
	fmt.Println("字典：")
	// map是一种元素对的无序集合，一组称为元素value，另一组为唯一键索引key。 未初始化map的值为nil。
	// map 是引用类型，可以使用如下声明：var map1 map[keytype]valuetype
	// map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据。
	// 在一个 nil 的slice中添加元素是没问题的，但对一个map做同样的事将会生成一个运行时的panic。
	// map 是引用类型的，内存用 make 方法来分配。map 的初始化：var map1 = make(map[keytype]valuetype)
	var map1 = make(map[string]interface{})
	map1["a1"] = 10
	map1["a2"] = "20"
	var map2 = map[string]int{"a": 1, "b": 2}
	fmt.Println(map1, map2)
	// 从 map1 中删除 key1，直接 delete(map1, key1) 就可以。如果 key1 不存在，该操作不会产生错误。
	delete(map1, "a2")
	delete(map1, "a3")
	fmt.Println(map1)
	// 标明 map 的初始容量 capacity，就像这样：make(map[keytype]valuetype，cap)。
	var map3 = make(map[string]int, 4)
	map3["b"] = 2
	fmt.Println(map3, len(map3))
	// map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序。
	//如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序（使用 sort 包）。
	fmt.Println("---------range----------")
	// 在"range"语句中生成的数据的值是真实集合元素的拷贝，它们不是原有元素的引用。
	//这意味着更新这些值将不会修改原来的数据。同时也意味着使用这些值的地址将不会得到原有数据的指针。
	data := []int{1, 2, 3, 4}
	for _, v := range data {
		v += 1
	}
	fmt.Println(data)
	// 如果你需要更新原有集合中的数据，使用索引操作符来获得数据。
	for i, _ := range data {
		data[i] += 1
	}
	fmt.Println(data)
	fmt.Println("---------------------------------")
	fmt.Println("流程控制：")
	// switch:
	sit := 4
	switch sit { // 常规的switch用法
	case 1:
		fmt.Println("sit's value is 1.")
	case 2:
		fmt.Println("sit's value is 2.")
	default:
		fmt.Println("sit's value is default.")
	}

	switch sit := 1; {
	case sit > 1:
		fmt.Println("sit > 1")
	case sit > 0:
		fmt.Println("sit > 0")
	default:
		fmt.Println("default....")
	}

	switch { // 在case使用bool表达式判断是否执行
	case sit > 2:
		fmt.Println("sit > 2")
		fallthrough // fallthrough强制执行下一条case代码，fallthrough不会判断下一条case的expr结果是否为true。
	case sit > 0:
		fmt.Println("sit > 0")
	default:
		fmt.Println("default....")
	}

	// select控制
	// select是Go中的一个控制结构，类似于switch语句，用于处理异步IO操作。
	// select会监听case语句中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。
	var c1, c2, c3 chan int
	var r1, r2 int
	select {
	case r1 = <-c1:
		fmt.Println("received ", r1, "from c1.")
	case c2 <- r2:
		fmt.Println("send ", r2, "to c2.")
	case i3, ok := (<-c3):
		if ok {
			fmt.Println("received ", i3, "from c3.")
		} else {
			fmt.Println("c3 was closed.")
		}
	case <-time.After(time.Second * 2): // 超时退出
		fmt.Println("request time out.")
	}
	// for循环
	for i := 0; i < len(arrLazy); i++ {
		fmt.Println(arrLazy[i])
	}
	var ss string = "你好，世界！"
	for i := 0; i < len(ss); i++ {
		fmt.Printf("%d:%c\t", i, ss[i])
	}
	fmt.Println()
	for i, v := range ss {
		fmt.Printf("%d:%c\t", i, v)
	}
	// Go语言中没有while，直接使用for来实现while
	fmt.Println()
	for {
		fmt.Println("use the 'for' to replace 'while'.")
		break
	}
	fmt.Println("---------------------------")
}

func arrHandle(arr []int) {
	// 使用range遍历数组
	for i, v := range arr {
		fmt.Println(i, ":", v)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
}

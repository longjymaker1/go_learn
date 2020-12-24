package main


import (
	"fmt"
)


func swap(a, b * int){
	// 取a指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}

func swap2(a, b *int){
	b, a = a, b
}

func main(){
	/**
	每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
	Go语言中使用在变量名前面添加&操作符（前缀）来获取变量的内存地址（取地址操作）
	**/
	// ptr := &v    // v 的类型为 T
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p\r\n \r\n", &cat, &str) // 使用 fmt.Printf 的动词%p打印 cat 和 str 变量的内存地址，指针的值是带有0x十六进制前缀的一组数据

	/** 从指针获取指针指向的值 **/
	fmt.Println("------------------------ 从指针获取指针指向的值 ----------------------")
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后,值的类型
	fmt.Printf("value type: %T,\nvalue= %s\n", value, value)

	/**
	取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

	变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
		对变量进行取地址操作使用&操作符，可以获得这个变量的指针变量
		指针变量的值是指针地址
		对指针变量进行取值操作使用*操作符，可以获得指针变量指向的原变量的值
	**/

	/** 使用指针修改值 **/
	fmt.Println("------------------------ 使用指针修改值 ----------------------")
	x, y := 1, 2
	fmt.Printf("swap 交换前 x = %d, y = %d\n", x, y)
	// 交换变量值
	swap(&x, &y)
	fmt.Printf("swap 交换后 x = %d, y = %d\n", x, y)

	c, d := 1, 2
	fmt.Printf("swap 交换前 c = %d, d = %d\n", c, d)
	swap2(&c, &d)
	fmt.Printf("swap2 交换后 c = %d, d = %d\n", c, d)

	/** 创建指针的另一种方法 new()函数 **/
	fmt.Println("------------------------ new()函数 ----------------------")
	sss := new(string)
	*sss = "我去"
	fmt.Println(*sss)
	//new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值
}
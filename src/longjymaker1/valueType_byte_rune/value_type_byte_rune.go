package main

import(
	"fmt"
	"math"
)

func main() {
	// 输出各数值范围
	fmt.Println("int 8 range: ", math.MinInt8, math.MaxInt8)
	fmt.Println("int 16 range: ", math.MinInt16, math.MaxInt16)
	fmt.Println("int 32 range: ", math.MinInt32, math.MaxInt32)
	fmt.Println("int 64 range: ", math.MinInt64, math.MaxInt64)

	// 初始化一个32位整型值
	var a int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("a int32: 0x%x %d\n", a, a)

	b := int16(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("b int16: 0x%x %d\n", b, b)

	// 将常量保存为float32类型
	var c float32 = math.Pi
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(c))
}
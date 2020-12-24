package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
	"bytes"
)

func main() {
	var str = "我嘞个去\n艹艹艹艹艹"
	fmt.Println(str)                    // ASCII 字符串长度使用 len() 函数

	str1 := "aaaaaaaa" + "dddddddd"
	fmt.Println(str1)

	const str2 = `第一行
	第二行
	第三行
	\r\n
	`
	fmt.Println(str2)

	fmt.Println("======================= 计算长度 ============================")
	tip1 := "ggsdfer  is sdg a dff"
	fmt.Println(len(tip1))

	tip2 := "我嘞个去"
	fmt.Println(len(tip2))

	fmt.Println(utf8.RuneCountInString(tip2))  // Unicode 字符串长度使用 utf8.RuneCountInString() 函数

	tip3 := "我嘞个去sdff"
	fmt.Println(utf8.RuneCountInString(tip3))

	fmt.Println("======================= 遍历字符串 ============================")
	// 遍历每一个ASCII字符
	theme := "我嘞个去你大爷的"
	for i := 0; i < len(theme); i++{
		fmt.Printf("ascii: c=%c  d=%d\n", theme[i], theme[i])
	}
	fmt.Println("--------------------------------------------------------------")
	for _, s := range theme {
		fmt.Printf("Unicode: c=%c  d=%d \n", s, s)
	}
	fmt.Println("======================= 字符串截取 ============================")
	theme1 := "我嘞个去, 你大爷的"
	comma := strings.Index(theme1, ",")

	fmt.Println(comma, theme1[comma])
	fmt.Println("======================= 字符串修改 ============================")
	msg := "Heros never die"
	msgBytes := []byte(msg)  // []byte 和 string 可以通过强制类型转换互转
	for i := 5; i <= 10; i++{
		msgBytes[i] = ' '
	}
	fmt.Println(string(msgBytes))
	fmt.Println("======================= 字符串拼接 ============================")
	stra := "我嘞个去"
	strb := "你大爷"

	// 声明字节缓冲
	var stringBuilder bytes.Buffer

	// 把字符串写入缓冲
	stringBuilder.WriteString(stra)
	stringBuilder.WriteString(strb)

	// 将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String())

	fmt.Println("======================= 格式化输出 ============================")
	//  fmt.Sprintf(格式化样式, 参数列表…)
	// 格式化样式：字符串形式，格式化动词以%开头
	// 参数列表：多个参数以逗号分隔，个数必须与格式化样式中的个数一一对应，否则运行时会报错
	var progress = 2
	var target = 8

	title := fmt.Sprintf("aaaaa%daaa, bbbb%ddddd", progress, target)
	fmt.Println(title)

	pi := 3.1415926
	variant := fmt.Sprintf("%v, %v, %v","fffffff", pi, true)
	fmt.Println(variant)

	profile := &struct {
		Name string
		HP int
	}{
		Name: "rat",
		HP: 150,
	}

	fmt.Printf("使用'%%+v' %+v\n", profile)
	fmt.Printf("使用'%%#v' %#v\n", profile)
	fmt.Printf("使用'%%T' %T\n", profile)
}
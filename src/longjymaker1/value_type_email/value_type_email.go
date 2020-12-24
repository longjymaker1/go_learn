package main


import (
	"encoding/base64"
	"fmt"
)


func main(){

	// 需要处理的字符串
	msg := "Away from keyboard. https://golang.org/"  // 行为需要编码的消息，消息可以是字符串，也可以是二进制数据。

	// 编码消息
	encodedMsg := base64.StdEncoding.EncodeToString([]byte (msg))  // base64 包有多种编码方法，这里使用 base64.StdEnoding 
														           // 的标准编码方法进行编码。传入的字符串需要转换为字节数组才能供这个函数使用

	// 输出编码完成的消息
	fmt.Println(encodedMsg)  // 编码完成后一定会输出字符串类型，打印输出

	//解码消息
	data, err := base64.StdEncoding.DecodeString(encodedMsg)  // 解码时可能会发生错误，使用 err 变量接收错误

	// 出错处理
	if err != nil {
		fmt.Println(err)  // 出错时，打印错误
	} else {
		//打印解码完成的数据
		fmt.Println(string(data))  // 正确时，将返回的字节数组（[]byte）转换为字符串
	}
}

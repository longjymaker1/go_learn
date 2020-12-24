package main

import "fmt"

func GetData() (int, int) {
	return 100, 200
}

func main(){
	var age int
	fmt.Println("我的年龄是", age)
	age = 29 // 赋值
	fmt.Println("my age is", age)
	age = 54 // 赋值
	fmt.Println("my new age is", age)
	var age1 = 30
	fmt.Println("年龄 = ", age1)
	var width, height int = 100, 50
	fmt.Println("width = ", width, "height = ", height)
	
	a, _ := GetData()
	_, b := GetData()
	fmt.Println("a = ", a, "b = ",b)
}

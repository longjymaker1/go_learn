# Hello World
## go工作区
在编写代码之前，我们首先应该建立 Go 的工作区（Workspace）。
在 Mac 或 Linux 操作系统下，Go 工作区应该设置在 `$HOME/go`。所以要在 `$HOME` 目录下创建 `go` 目录。
而在 Windows 下，工作区应该设置在 `C:\Users\YourName\go`。所以将 go 目录放置在 `C:\Users\YourName`。
其实也可以通过设置 `GOPATH` 环境变量，用其他目录来作为工作区
当前GOPATH
```shell script
C:\Users\longjy>go env
set GOCACHE=C:\Users\longjy\AppData\Local\go-build
set GOENV=C:\Users\longjy\AppData\Roaming\go\env
set GOPATH=D:\pro\go_work
```
* 所有 Go 源文件都应该放置在工作区里的 src 目录下
* 所有 Go 项目都应该依次在 src 里面设置自己的子目录

在 src 里面创建一个目录 hello 来放置整个 hello world 项目
目录结构为
```shell script
# GOPATH = D:\pro\go_work
GOPATH
    src
        hello
            hello.go
```
编写hello.go代码
```go
package main  // 每一个 Go 文件都应该在开头进行 package name 的声明（译注：只有可执行程序的包名应当为 main）。
			  // 包（Packages）用于代码的封装与重用，这里的包名称是main

import "fmt"  // 我们引入了 fmt 包，用于在 main 函数里面打印文本到标准输出

func main() {  // main 是一个特殊的函数。整个程序就是从 main 函数开始运行的。
	           // main 函数必须放置在 main 包中。{ 和 } 分别表示 main 函数的开始和结束部分
	fmt.Println("Hello World!")  // fmt 包中的 Println 函数用于把文本写入标准输出
}

```
1. `run`命令，在命令行中输入 `go run workspacepath/src/hello/hello.go`,workspacepath为实际路径
```shell script
PS C:\Windows\System32> go run D:\pro\go_work\learn\src\hello.go
Hello World!
```
2. `install`命令, 命令行输入 `go install workspacepath/src/hello/hello.go`, 然后`workspacepath/bin/hello`运行程序
当你输入 go install hello 时，go 工具会在工作区中搜索 hello 包（hello 称之为包，后面会更加详细地讨论包）。
接下来它会在工作区的 bin 目录下，创建一个名为 hello（Windows 下名为 hello.exe）的二进制文件。运行 go install hello 后，
其目录结构如下
```shell script
# GOPATH = D:\pro\go_work
GOPATH
    src
        hello
            hello.go
    bin
        hello
```
```shell script
PS C:\Windows\System32> go install hello
PS C:\Windows\System32> D:\pro\go_work\bin\hello.exe
Hello World!
```

# variable 变量
变量指定了某存储单元（Memory Location）的名称，该存储单元会存储特定类型的值
## 声明变量
### 单个变量
var name type 是声明单个变量的语法
```go
package main

import "fmt"

func main() {
    var age int // 变量声明
    fmt.Println("my age is", age) // 未初始化会被自动赋值为零值
}
```
### 声明变量并初始化
声明变量的同时可以给定初始值。 `var name type = initialvalue` 的语法用于声明变量并初始化
```go
package main

import "fmt"

func main() {
    var age int = 29 // 声明变量并初始化

    fmt.Println("my age is", age)
}
```
### 类型推断（Type Inference）
如果变量有初始值，那么 Go 能够自动推断具有初始值的变量的类型。因此，如果变量有初始值，就可以在变量声明中省略 type。

如果变量声明的语法是 `var name = initialvalue`，Go 能够根据初始值自动推断变量的类型
```go
package main

import "fmt"

func main() {
    var age = 29 // 声明变量并初始化

    fmt.Println("my age is", age)
}
```
### 声明多个变量
声明多个变量的语法是 `var name1, name2 type = initialvalue1, initialvalue2`
```go
package main

import "fmt"

func main() {
    var width, height int = 100, 50 // 声明多个变量

    fmt.Println("width is", width, "height is", height)
}
```
在有些情况下，我们可能会想要在一个语句中声明不同类型的变量
```go
var (  
    name1 = initialvalue1,
    name2 = initialvalue2
)
```

```go
package main

import "fmt"

func main() {
    var (
        name   = "naveen"
        age    = 29
        height int
    )
    fmt.Println("my name is", name, ", age is", age, "and height is", height)
}
```
### 简短声明
Go 也支持一种声明变量的简洁形式，称为简短声明（Short Hand Declaration），该声明使用了 := 操作符
声明变量的简短语法是 `name := initialvalue`
```go
package main

import "fmt"

func main() {  
    name, age := "naveen", 29 // 简短声明

    fmt.Println("my name is", name, "age is", age)
}
```
短声明的语法要求 := 操作符的左边至少有一个变量是尚未声明的
```go
package main

import "fmt"

func main() {
    a, b := 20, 30 // 声明变量a和b
    fmt.Println("a is", a, "b is", b)
    b, c := 40, 50 // b已经声明，但c尚未声明
    fmt.Println("b is", b, "c is", c)
    b, c = 80, 90 // 给已经声明的变量b和c赋新值
    fmt.Println("changed b is", b, "c is", c)
}

/**
a is 20 b is 30
b is 40 c is 50
changed b is 80 c is 90
*/
```
```go
package main

import "fmt"

func main() {  
    a, b := 20, 30 // 声明a和b
    fmt.Println("a is", a, "b is", b)
    a, b := 40, 50 // 错误，没有尚未声明的变量
}
/**
no new variables on left side of :=
*/
```
**由于 Go 是强类型（Strongly Typed）语言，因此不允许某一类型的变量赋值为其他类型的值**
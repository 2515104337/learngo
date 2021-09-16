// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// package main 是一个特殊的包
// 它允许 Go 创建一个可执行文件
package main

/*
这是一个多行注释。

import 关键字使另一个包可用于此 .go “文件”。

import "fmt" 允许您在此文件中访问 fmt 包的功能。
*/
import "fmt"

// “func main” 是一个特殊的函数。
//
// Go 必须知道从哪里开始执行
//
// func main 为 Go 创建一个起点
//
// 编译代码后，Go 运行时会先运行这个函数
func main() {
	// import “fmt” 之后
	// “fmt” 包的 Println 功能可用

	// 通过在控制台中键入下面命令来查看它的样子：
	//   go doc -src fmt Println

	// Println 只是一个输出函数
	//   "fmt" package

	// 输出 = 第一个字母是大写
	fmt.Println("Hello Gopher!")

	// Go 不能自己调用​​ Println 函数.
	// 这就是为什么你需要在这里调用它。
	// 它只会自动调用 `func main`.

	// -----

	// Go 支持字符串文字中的 Unicode 字符
	// 而且在源代码中: KÖSTEBEK!
	//
	// 因为: 字面 ~= 源代码

	// 练习: 删除下面的注释 --> //
	fmt.Println("Merhaba Köstebek!")

	// 不需要注释:
	// "Merhaba Köstebek" 等同于 "Hello Gopher"
	// 在土耳其语言中
}

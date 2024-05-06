package main

import (
	"fmt"
	"strings"
)

func main() {

	// 声明字符串变量
	var question string

	// 无限循环
	for {
		fmt.Println("用户： ")
		_, err := fmt.Scanln(&question)

		/*
			go传参模式：
			1. 普通传参：位置传参，可读性不佳，破坏兼容性，解决方法：引入一个参数结构体（sturct）类型
			2. 参数结构体
			3. 函数式选项
		*/

		// 错误处理，如果不等nil会直接返回
		if err != nil {
			return
		}

		// string包中Replace函数式字符串替代函数
		question = strings.Replace(question, "?", "", -1)
		question = strings.Replace(question, "吗", "", -1)

		// 机器人回复
		fmt.Printf("机器人：%s%s\r\n", question, "!")
	}
}

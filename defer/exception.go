package main

import "fmt"

func testPanic() {
	panic("Panic A")
}

func main() {
	// 用defer配合recover来实现异常捕获，这样做的好处在于可以保证程序继续往下走，而不至于崩溃
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
	}()
	fmt.Println("Starting completes here!")
	testPanic()
	fmt.Println("Main block completes here!") // 注意这里的不会被执行
}

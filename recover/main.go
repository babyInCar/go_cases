package main

import "fmt"

func a() {
	defer b()
	panic("panic a ")
}

func b() {
	defer catch("b")
	defer fb()
	panic("panic b")
}

// 如果把recover放在调用链中间或者最后面一环，那么不会往回回溯，只会调用当前以及当前代码调用的函数发生的Panic
func fb() {
	panic("panic fb")
}

// recover 只会捕获最近的一次Panic
func catch(funcname string) {
	if r := recover(); r != nil {
		fmt.Println(funcname, "recover", r)
	}
}

func main() {
	defer catch("main")
	a()
}

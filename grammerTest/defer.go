package main

import "fmt"
/*
执行方式类似其它语言中的析构函数
即使函数发生严重错误也会执行
常用于资源清理、文件关闭、解锁以及记录时间等操作
Go 没有异常机制，但有 panic/recover 模式来处理错误
Panic 可以在任何地方引发，但recover只有在defer调用的函数中有效
 */

func main() {
	for i:=0;i<3;i++  {
		defer fmt.Println(i)
	}

	////输出3 3 3
	//for i := 0; i < 3; i++ {
	//	defer func() {
	//		fmt.Println(i)
	//	}()
	//}

	A()
	B()
	C()
}

func A()  {
	fmt.Println("Func A")
}

func B(){

	defer func() {
		if err := recover();err!=nil{
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic B")
}

func C()  {
	fmt.Println("Func C")
}

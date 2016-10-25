package main

import "fmt"
/*
小写是包私有的
大写是公有的导出的
 */
type A struct {
	Name string
}
type B struct {
	Name string
}

//类型别名
type TZ int

func main() {

	a:=A{}
	a.Print()

	b:=B{}
	b.Print()

	var z TZ
	//通过结构体类模式调用
	z.Print()
	//通过类型调用
	(*TZ).Print(&z)
}

//方法,接受参数放前，函数的第一个接收者，作为其方法
func (a A) Print() {
	fmt.Println("A")
}

//方法是和结构体绑定的
func (b B) Print() {
	fmt.Println("B")
}
//int 的也可以绑定一个方法
func (tz *TZ)Print(){
	fmt.Println("TZ")
}
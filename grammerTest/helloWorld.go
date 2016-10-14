package main

import "fmt"

var (
	g1 int
)

func main() {

	fmt.Println("hello world! Welcome to Go!");
	//变量声明与赋值 var identifier type | v_name := value (这种赋值方式不适合声明全局变量)
	//变量声明后没有使用是会编译报错的
	var a int = 10
	var b float32
	c := 20

	fmt.Println(a,b,c)

	//多变量的并行赋值
	var s1,s2,s3 = "s1","s2","s3"
	fmt.Println(s1+s2+s3)

	//抛弃不要的值 _
	p1,_,p2 := 1,2,3;
	fmt.Println("p1=",p1,"p2=",p2);

	//因式分解写法，用于全局变量声明,此处并不是
	var (
		globle_1 int = 100
		globle_2 string = "全局变量"
	)

	fmt.Println(globle_1,globle_2);


	/*
	值类型：int float bool string
	引用类型：
	 */
	var a1 int = 7;
	var a2 = a1;

	fmt.Println(&a1,&a2)

	/*
	常量的类型：布尔型，数字型，字符串型
	const indentifier [type] = value
	 */
	const C1 string = "abc"
	const C2 = "abc"
	const C3,C4,C5 = 1,false,"str"
	const (
		UNKNOW = 0
		FEMALE = 1
		MALE = 2
	)

	//快速交换值
	var s11,s21 int = 1,2
	fmt.Println("交换前：s11 = ",s11," s21 = ",s21)
	s11,s21 = s21,s11
	fmt.Println("交换后：s11 = ",s11," s21 = ",s21)

	grade := "B"
	test := "1qaz"
	fmt.Println("字符串输出测试: %s\n",test);
	fmt.Printf("你的等级是 %s\n", grade );









}

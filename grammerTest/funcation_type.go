package main

import(
	"fmt"
	"math"
)

func main() {
	/*
	func function_name( [parameter list] ) [return_types] {
   		函数体
	}
	 */
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf( "最大值是 : %d\n", ret )

	name1,name2 := swap("hzqiuxm","handanrujiaren")
	fmt.Println(name1,name2)
	fmt.Println("--------------------------------------")
	/*
	函数作为返回值
 	*/
	getquereRoot := func(x float64) float64{
		return math.Sqrt(x);
	}

	fmt.Println(getquereRoot(9))
	fmt.Println("--------------------------------------")

	/*
	函数闭包
	 */
	/*
	 nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	fmt.Println("--------------------------------------")
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())



}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}
/*
返回多个值
 */
func swap(x, y string) (string, string) {
	return y, x
}

/*
函数闭包
*/
func getSequence() func() int {
	i:=0
	return func() int {
		i+=1
		return i
	}
}

/*
Go 语言中同时有函数和方法。
一个方法就是一个包含了接受者的函数，
接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
所有给定类型的方法属于该类型的方法集。语法格式如下：
func (variable_name variable_data_type) function_name() [return_type]{
   	函数体
}
 */
/* 定义函数 */
type Circle struct {
	radius float64
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

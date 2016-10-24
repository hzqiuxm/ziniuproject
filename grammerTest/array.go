package main

import "fmt"
/*
数组在go中是数值类型
数组是固定长度的，不同长度数组属于不同类型
你可以省略用...代替，但是不能为空，为空就编程slice了
 */
func main() {
	//多维数组测试
	//array_test();

	var a int = 20;
	var ip *int;
	ip = &a
	fmt.Println(ip,"  ",*ip);
	fmt.Println("----------------------------------");
	var b [2]int
	 c := [...]int{19:1}
	var p *[20]int = &c
	fmt.Println(b,c,p)
	var p1 = new ([10]int) //指向数组的指针
	fmt.Println(p1)

	fmt.Println("----------------冒泡排序------------------");
	a1 := [...]int {13,5,78,11,54,21,52,98,31}
	fmt.Println("排序前：",a1)
	//var i,j int;
	num := len(a1)
	for i:=0;i<num;i++{
		for j:= i+1;j<num;j++{
			if(a1[i]>a1[j]){
				a1[i],a1[j] = a1[j],a1[i]
			}
		}
		fmt.Println("i=",i,"序列：",a1)
	}

	fmt.Println("排序后：",a1)



}

func array_test(){

	/* 数组 - 5 行 2 列*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}


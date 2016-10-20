package main

import "fmt"

func main() {

	var n int = 15
	fmt.Printf("%d 的阶乘是 %d\n", n, Factorial(n))

	var i int
	for  i = 0; i < 10; i++ {
		fmt.Printf("斐波那契数列 %d\t", fibonaci(i))
	}


}
/*/
阶乘
 */
func Factorial(x int) (result int) {
	if x == 0 {
		result = 1;
	} else {
		result = x * Factorial(x - 1);
	}
	return;
}

/*
斐波那契数列
 */
func fibonaci(n int) int {
	if n < 2 {
		return n
	}
	return fibonaci(n-2) + fibonaci(n-1)
}
package main

import "fmt"
/*
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，
功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
 */
func main() {
	//make([]T, length, capacity)
	var numbers = make([]int,3,6)
	printSlice(numbers)

	numbers2 := []int {0,1,2,3,4,5,6,7,8}
	printSlice(numbers2)

	var numbers3 [] int
	printSlice(numbers3)
	if(nil == numbers3){
		println("numbers3 is empty!")
	}

	//切片使用
	fmt.Println("numbers2[1:4] ==", numbers2[1:4])
	fmt.Println("numbers2[:4] ==", numbers2[:4])
	fmt.Println("numbers2[1:] ==", numbers2[1:])
	fmt.Println("numbers2[:] ==", numbers2[:])

	//切片的追加和复制
	numbers = append(numbers,1,2,3)
	printSlice(numbers)
	numbers = append(numbers,4)//超过预设的capacity大小后成倍增长
	printSlice(numbers)
	var numbers_new = make([]int,len(numbers),cap(numbers)*2)
	copy(numbers_new,numbers)
	printSlice(numbers_new)
}
func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
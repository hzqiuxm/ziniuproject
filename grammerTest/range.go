package main

import "fmt"
/*
Go 语言中 range 关键字用于for循环中迭代数组(array)、切片(slice)、链表(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引值，在集合中返回 key-value 对的 key 值。
 */
func main() {
	numbers := []int{0,1,2,3,4}
	sum :=0
	for _,num := range numbers{
		sum += num
	}

	fmt.Println("sum=",sum)

	for i,num:=range numbers{
		if(3 == num){
			println("index is ",i)
		}
		fmt.Println("num = ",num)
	}

	//range在map中的使用
	kvs := map[string]string{"a":"apple","b":"banana"}

	for k,v:= range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

package main

import (
	"fmt"
	"runtime"
)
/*
超级线程池
主要是每个实例通过4-5K栈内存，大幅减少的创建和销毁，是制造Go号称高并发根本原因
并发不等于并行
goroutine简单易用，也在语言层面上给予了极大方便
通过通信来共享内存，而不是通过内存在实现通讯
桥梁Channel  引用类型  可以设置单向或双向通道  可以设置缓存
注意无缓存的时候，是阻塞的，读写顺序不重要；有缓存的时候是异步的，
否则写完会存在爱读不读情况

 */
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go1(c,i)
	}
	<-c

}

//多个CPU执行，执行顺序不定,有些不会执行
//如果9先执行，其他就不执行了
func Go1(c chan bool,index int){
	a := 1
	for i:=0;i<100000000;i++ {
		a +=i
	}
	fmt.Println(index,a)

	if index == 9{
		c <- true
	}
}
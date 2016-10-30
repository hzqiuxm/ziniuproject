package main

import (
	"fmt"
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
Select :处理多个channel接收和发送
有多个可执行的select是随机执行的
可设置超时
可阻塞main函数
 */
func main() {
	c1,c2 := make(chan int),make(chan string)
	//避免进程提前结束通道
	o := make(chan bool)

	go func() {
		for{
			select {
			case v,ok := <-c1 :
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1:",v)
			case v,ok := <-c2 :
				if !ok{
					o <- true
					break
				}
				fmt.Println("c2:",v)
			}
		}
	}()

	c1<-1
	c2<- "hi"
	c1<-3
	c2<- "hello"

	close(c1)
	close(c2)

	//读出o通道的值
	<-o

}

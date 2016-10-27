package main

import "fmt"
/*
接口只有方法声明，没有实现，没有数据字段
接口可以匿名嵌入其它接口，或嵌入到结构中
接口调用不会做receiver的自动转换
接口同样支持匿名字段方法
接口也可实现类似OOP中的多态
空接口可以作为任何类型数据的容器
将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个
复制品的指针，既无法修改复制品的状态，也无法获取指针

 */

type USB interface {
	Name() string
	Connector
}

type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func (phoneC PhoneConnector) Name() string{
	return phoneC.name
}
func (phoneC PhoneConnector) Connect() {
	fmt.Println("Connector with ",phoneC.name)
}

//为了让Disconnect 方法更加使用，将传入形参改成interface{},就可以接受任何参数了，以为
//默认所有的方法都是实现了空接口的
func Disconnect(usb interface{}){
	if phoneC,ok := usb.(PhoneConnector);ok{
		fmt.Println("Disconnect from ",phoneC.name)
		return
	}
	fmt.Println("Unknow device !")

}
func main() {
	//var pc USB
	pc := PhoneConnector{"PhoneConnector"}
	pc.Connect()
	Disconnect(pc)
	//b拿到的是copy,对b属性修改不起作用
	var b Connector
	b = Connector(pc)
	b.Connect()
	pc.name = "hello"
	b.Connect()

}


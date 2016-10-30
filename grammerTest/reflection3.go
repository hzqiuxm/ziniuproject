package main

import (
	"fmt"
	"reflect"
)
/*
通过反射调用方法
 */

type User struct {
	Id int
	Name string
	Age int
}

func (u User) SayHello(name string){
	fmt.Println("Helo",name ,", my name is",u.Name)
}

func main() {
	u := User{1,"hzqiuxm",30}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("SayHello")

	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}

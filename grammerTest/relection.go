package main
/*
主要是二个函数： Typeof()，Valueof() 获取接口中目标对象的信息
反射会将匿名字段作为独立字段
想利用反射修改对象状态，前提是interface.data 是settable即pointer-interface
通过反射可以动态调用方法
 */
import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello(){
	fmt.Println("I'm hello method ! ")
}

func main() {

	u := User{1,"qiuxm",30}
	Info(u)

}

func Info(o interface{}){
	t := reflect.TypeOf(o)
	fmt.Println("Type: ",t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields: ")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()

		fmt.Println(" 字段名称：",f.Name," 字段类型：",f.Type,"字段值：",val)

	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println("方法名称：",m.Name,"方法类型：",m.Type)
	}


}

package main
/*
反射嵌套结构和通过反射修改结构属性值
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

type Manager struct {
	User
	title string

}

func Set(o interface{}){
	v:= reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet(){
		fmt.Println("Cant modified")
		return
	}else {
		v = v.Elem()
	}

	//如果找到就修改否则提示没有找到需要修改的字段
	f:= v.FieldByName("Name")

	if !f.IsValid(){
		fmt.Println("修改的属性不存在")
		return
	}

	if f.Kind()==reflect.String{
		f.SetString("weclome")
	}

}
func main() {

	m := Manager{User:User{1,"hzqiuxm",30},title:"qxm"}
	t := reflect.TypeOf(m)

	fmt.Println("Manager的User的name: ",t.FieldByIndex([]int{0,1}))
	fmt.Println((t.FieldByName("User")))

	x := 123
	v := reflect.ValueOf(&x)

	v.Elem().SetInt(999)

	fmt.Println("x= ",x)

	//修改一个结构对象
	u2 := User{1,"dingjiahong",25}
	Set(&u2)
	fmt.Println(u2)

}

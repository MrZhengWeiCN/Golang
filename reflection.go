package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Print("hello ", name, ",my name is ", u.Name)
}

func main() {
	u := User{1, "OJBk", 12}
	v := reflect.ValueOf(u)                           //取得反射对象
	mv := v.MethodByName("Hello")                     //获得Hello方法
	args := []reflect.Value{reflect.ValueOf("Wayne")} //参数获取
	mv.Call(args)                                     //执行方法
}

package object_function

import (
	"fmt"
	"reflect"
	"testing"
)

//结构体定义
type student struct {
	age  int
	name string
}

//带标签的结构体
type teacher struct {
	age  int    "年龄"
	name string "名字"
}

//两个属性项都是首字符大写，代表是可以public访问
type Teacher struct {
	Age  int    "年龄"
	Name string "名字"
}

func TestStructExample(in *testing.T) {

	//结构体初始化new方式
	var stu = new(student)

	stu.name = "jack"
	stu.age = 12

	fmt.Println("name:", stu.name, ";age:", stu.age)

	//简写方式，底层依旧调用了new,&可以省略
	var stu0 = &student{12, "jack"}
	var stu1 = student{12, "jack"}

	fmt.Println("name:", stu0.name, ";age:", stu0.age)
	fmt.Println("name:", stu1.name, ";age:", stu1.age)

	//通过反射获取结构体标签
	var tea = teacher{30, "小张"}
	teaType := reflect.TypeOf(tea)
	teaAgeField := teaType.Field(0)
	fmt.Println(teaAgeField.Tag)
}

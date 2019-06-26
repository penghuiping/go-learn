package object_function

import (
	"fmt"
	"reflect"
	"testing"
)

type student1 struct {
	name string "名字"
	age  int    "年龄"
}

func TestReflectExample(in *testing.T) {
	//使用反射分析结构体
	var stu = student1{"jack", 12}
	var stuRef = reflect.ValueOf(stu)

	var i = 0
	for i = 0; i < stuRef.NumField(); i++ {
		fmt.Println(stuRef.Type().Field(i).Name, ":", stuRef.Type().Field(i).Tag, ":", stuRef.Field(i))
	}

}

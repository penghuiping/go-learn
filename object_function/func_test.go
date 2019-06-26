package object_function

import (
	"fmt"
	"time"
	"testing"
)

//多个参数，一个返回值
func add(a int, b int) int {
	return a + b
}

//方法名第一个字符大写代表public方法,否则为私有方法
func Add(a int, b int) int {
	return add(a, b)
}

//多个返回值
func addAndMinus(a int, b int) (int, int) {
	return a + b, a - b
}

//入参数不确定
func addMulti(a ...int) int {
	tmp := 0
	for _, value := range a {
		tmp = tmp + value
	}
	return tmp
}

//函数作为函数的参数
func phpAdd(a1 int, b1 int, f func(a int, b int) int) int {
	return f(a1, b1)
}

//用函数最为函数对的返回值
func phpAdd1(a1 int, b1 int) int {
	return func(a int, b int) int {
		return a + b
	}(a1, b1)
}

func TestFuncExample(in *testing.T) {
	i := add(1, 2)
	fmt.Println(i)

	a, b := addAndMinus(1, 2)
	fmt.Println(a, b)
	result := addMulti(1, 2, 3, 4, 5)
	fmt.Println(result)

	result1 := phpAdd(1, 2, add)
	fmt.Println(result1)

	//匿名函数,函数闭包
	func() {
		fmt.Println("hello world")
	}()

	fmt.Println(phpAdd1(1, 2))

	//计算函数的执行时间
	start := time.Now()
	phpAdd1(1, 2)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

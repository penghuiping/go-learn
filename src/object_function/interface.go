package object_function

import "fmt"

//接口定义
type shape interface {
	area() float32
}

type square struct {
	side float32
}


//square实现接口
func (sq *square) area() float32 {
	return sq.side * sq.side
}

type rectangle struct {
	width  float32
	height float32
}

//rectangle实现接口
func (rc *rectangle) area() float32 {
	return rc.width * rc.height
}

func cacArea(a shape) {
	fmt.Println((a).area())
}

func InterfaceExample() {
	//多态,注意var r [shape]后缀
	var r shape = &rectangle{12.0, 13.0}
	var s shape = &square{12.0}

	//多态
	cacArea(r)
	cacArea(s)

	//判断r是否是rectangle类型,类型判断
	if _, ok := r.(*rectangle); ok {
		fmt.Println(ok)
	}
}

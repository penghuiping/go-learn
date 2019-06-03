package basic_type

import (
	"fmt"
	"strconv"
	"testing"
)

func TestDoActionBase(in *testing.T) {

	//打印
	fmt.Println("Hello, World!")

	//for循环
	var arr = [3]int{1, 2, 3}
	i := 0
	for i = 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//for range循环
	for pos, value := range arr {
		fmt.Println(strconv.Itoa(pos) + "," + strconv.Itoa(value))
	}

	//分之语句
	var tmp = 2
	if tmp > 2 {
		fmt.Println("tmp is bigger than two")
	} else if tmp == 2 {
		fmt.Println("tmp is equal two")
	} else {
		fmt.Println("tmp is smaller than two")
	}

	//switch语句
	var a = "0"
	switch a {
	case "1":
		fmt.Println("a is 1")
		break
	case "2":
		fmt.Println("a is 2")
		break
	default:
		fmt.Println("a is default value 1")
		break
	}
}

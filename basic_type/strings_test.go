package basic_type

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestDoActionStrings(in *testing.T) {

	var b1 = "hello" + "world";
	fmt.Println(b1);

	var a = "hello world"
	fmt.Println(a)
	if strings.Contains(a, "hello") {
		fmt.Println("a constans hello")
	}

	var tmp = strings.Split(a, " ")
	var i = 0
	for i = 0; i < len(tmp); i++ {
		fmt.Println(tmp[i])
	}

	c := 1
	fmt.Println(c)

	var b, _ = strconv.Atoi("1")
	fmt.Println(b + 1)
}

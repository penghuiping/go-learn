package _io

import (
	"bufio"
	"fmt"
	"os"
)

func StdioExample() {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入一些数据")
	var value, err = reader.ReadString('\n')
	if err == nil {
		fmt.Println(value)
	}
	var firstname, lastname string

	fmt.Scan(&firstname, &lastname)

	fmt.Println(firstname, lastname)
}

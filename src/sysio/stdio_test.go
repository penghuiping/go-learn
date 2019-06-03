package sysio

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestStdioExample(in *testing.T) {
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

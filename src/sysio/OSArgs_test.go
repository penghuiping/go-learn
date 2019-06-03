package sysio

import (
	"fmt"
	"os"
)

func OSArgsExample() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1], os.Args[2])
	}
}

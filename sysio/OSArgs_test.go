package sysio

import (
	"fmt"
	"os"
	"testing"
)

func TestOSArgsExample(in *testing.T) {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1], os.Args[2])
	}
}

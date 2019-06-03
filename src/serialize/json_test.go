package serialize

import (
	"encoding/json"
	"fmt"
	"testing"
)

type teacher struct {
	name string
	age  int
}

type school struct {
	name     string
	teachers []*teacher
}

func TestJsonExample(in *testing.T) {
	var teacher0 = &teacher{"jack", 22}
	var teacher1 = &teacher{"mary", 22}

	var teachers = []*teacher{teacher0, teacher1}
	var school1 = &school{"小学", teachers}

	bArr, _ := json.Marshal(teacher0)
	fmt.Println(string(bArr))

	barr1, _ := json.Marshal(school1)
	fmt.Println(string(barr1))
}

package serialize

import (
	"encoding/json"
	"fmt"
)

type teacher struct {
	name string
	age  int
}

type school struct {
	name     string
	teachers []*teacher
}

func JsonExample() {
	var teacher0 = &teacher{"jack", 22}
	var teacher1 = &teacher{"mary", 22}

	var teachers = []*teacher{teacher0, teacher1}
	fmt.Println(teachers)
	var school1 = school{"小学", teachers}
	fmt.Println(school1)
	bArr, _ := json.Marshal(teacher0)
	fmt.Printf(string(bArr))
}

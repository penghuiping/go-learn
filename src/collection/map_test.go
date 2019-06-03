package collection

import "fmt"

func printMap(a *map[string]int) {
	for p, v := range *a {
		fmt.Println(p, v)
	}
}

func MapExample() {
	fmt.Println("===========>华丽的分割线<===========")
	//初始化
	var mapLit = map[string]int{"key0": 1, "key1": 2}

	var mapLit1 = make(map[string]int)

	mapLit1["key0"] = 1
	mapLit1["key1"] = 2

	//遍历
	for p, v := range mapLit {
		fmt.Println(p, v)
	}

	fmt.Println("===========>华丽的分割线<===========")
	//新增
	mapLit["key2"] = 3

	//遍历
	printMap(&mapLit)

	fmt.Println("===========>华丽的分割线<===========")
	//删除
	delete(mapLit, "key2")

	//遍历
	printMap(&mapLit)

}

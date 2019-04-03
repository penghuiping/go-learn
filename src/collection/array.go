package collection

import "fmt"

func ArrayExample() {

	fmt.Println("==================>初始化数组默认值为0<==================")
	var arr1 [5]int
	for p, v := range arr1 {
		fmt.Println(p, v)
	}

	fmt.Println("==================>初始化数组并且赋值<==================")
	var arr2 = [5]int{1, 2, 3, 4, 5}
	for p, v := range arr2 {
		fmt.Println(p, v)
	}

	fmt.Println("==================>数组初始化，通过声明值来确定数组大小<==================")
	var arr3 = [...]int{1, 2, 3, 4, 5}
	for p, v := range arr3 {
		fmt.Println(p, v)
	}

	fmt.Println("==================>数组初始化，初始化部分索引参数<==================")
	var arr4 = [5]string{1: "hello", 3: "world"}
	for p, v := range arr4 {
		fmt.Println(p, v)
	}

	fmt.Println("==================>多维数组<==================")

	var arr5 = [3][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {5, 3, 2, 1, 1}}

	var i = 0
	var j = 0
	for i = 0; i < len(arr5); i++ {
		for j = 0; j < len(arr5[i]); j++ {
			fmt.Print(arr5[i][j], " ")
		}
		fmt.Println("")
	}
}

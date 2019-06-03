package collection

import "fmt"

//用切片作为函数的入参
func add(a []int) int {
	i := 0
	for _, v := range a {
		i += v
	}
	return i
}

/**
* slice是数组的引用
 */
func SliceExample() {
	var arr1 = [3]int{1, 2, 3}

	//声明切片
	var slice1 = arr1[1:]

	for _, v := range slice1 {
		fmt.Print(v, " ")
	}

	fmt.Println()
	fmt.Println("===============>华丽的分割线<===============")

	fmt.Println("slice1的len为", len(slice1))
	fmt.Println("slice1的capacity为", cap(slice1))

	fmt.Println("===============>用切片作为函数的入参<===============")
	fmt.Println(add(slice1))

	fmt.Println("===============>使用make函数创建一个切片<===============")
	//使用make创建一个切片
	var slice2 = make([]int, 5, 10)

	slice2[1] = 1
	slice2[2] = 2
	slice2[3] = 3

	for _, v := range slice2 {
		fmt.Print(v, " ")
	}

	fmt.Println()
	fmt.Println(len(slice2), "", cap(slice2))

	fmt.Println("===============>切片扩容<===============")
	slice2 = slice2[0 : len(slice2)+2]
	for _, v := range slice2 {
		fmt.Print(v, " ")
	}
	fmt.Println("\n===============>切片复制<===============")
	var slice3 = make([]int, 6)
	copy(slice3, slice2)
	for _, v := range slice3 {
		fmt.Print(v, " ")
	}

	fmt.Println("\n===============>切片追加<===============")
	slice2 = append(slice2, 1, 2, 3)
	for _, v := range slice2 {
		fmt.Print(v, " ")
	}

}

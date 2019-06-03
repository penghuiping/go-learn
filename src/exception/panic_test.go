package exception

import "fmt"

func panicFunc() {
	defer fmt.Println("defer才是最后执行的,就算遇到panic错误也会执行")
	fmt.Println("这种错误会中断程序")
	panic("我要中断程序")
	fmt.Println("这是最后一行代码")
}

func catchPanicFunc() {
	//捕捉panic异常，需要defer方式，并写在函数第一行
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕捉panic异常")
			fmt.Println("panic异常是:" + fmt.Sprint(err))
		}
	}()
	panicFunc()
}

func PanicExample() {
	catchPanicFunc()
}

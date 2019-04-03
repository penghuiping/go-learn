package thread

import (
	"fmt"
	"sync"
	"runtime"
	"time"
)

func GoRoutineExample() {
	//设置线程数量，类似定义一个线程池
	runtime.GOMAXPROCS(2)

	//类似java中的 countDownLatch
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			fmt.Println("")
			fmt.Println("1111111")
			fmt.Println("3333333")
			fmt.Println("")
			time.Sleep(1 * time.Millisecond)
		}

	}()

	go func() {
		defer wg.Done()
		for {
			fmt.Println("")
			fmt.Println("2222222")
			fmt.Println("4444444")
			fmt.Println("")
			time.Sleep(1 * time.Millisecond)
		}
	}()
	wg.Wait()
}

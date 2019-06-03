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
		var count = 0
		for {
			fmt.Println("")
			fmt.Println("1111111")
			fmt.Println("3333333")
			fmt.Println("")
			time.Sleep(1000 * time.Millisecond)
			count++
			if 5 == count {
				break;
			}
		}

	}()

	go func() {
		defer wg.Done()
		var count = 0
		for {
			fmt.Println("")
			fmt.Println("2222222")
			fmt.Println("4444444")
			fmt.Println("")
			time.Sleep(1000 * time.Millisecond)
			count++
			if 5 == count {
				break;
			}
		}
	}()
	wg.Wait()
}

func GoRoutingExample2() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "2s"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}
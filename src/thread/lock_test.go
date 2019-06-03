package thread

import (
	"sync"
	"runtime"
	"fmt"
)

var count = 0

var mutex sync.Mutex

func countPlus() {
	mutex.Lock()
	count = count + 1;
	defer mutex.Unlock()
}

func countMinus() {
	mutex.Lock()
	count = count - 1
	defer mutex.Unlock()
}

/*
使用互斥锁保证同步
 */
func LockExample() {
	//设置线程数量，类似定义一个线程池
	runtime.GOMAXPROCS(8)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10000; i++ {
			countPlus()
		}
		defer wg.Done()
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			countMinus()
		}
		defer wg.Done()
	}()
	wg.Wait()
	fmt.Println(count)
}

package thread

import (
	"context"
	"time"
	"fmt"
	"testing"
)

func TestContext(in *testing.T) {
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(),1*time.Second)
	go func() {
		ctx1 := context.WithValue(ctx,"name","gorouting1")
		for {
			select {
			case <-ctx1.Done():
				fmt.Println("gorouting1 finished")
				return
			default:
				fmt.Println("gorouting1 runing:","value:",ctx1.Value("name"))
			}
			time.Sleep(time.Second)
		}

	}()

	go func() {
		ctx2 := context.WithValue(ctx,"name","gorouting2")
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("gorouting2 finished")
				return
			default:
				fmt.Println("gorouting2 runing:","value:",ctx2.Value("name"))
			}
			time.Sleep(time.Second)
		}

	}()

	go func() {
		time.Sleep(10 * time.Second)
		cancel()
	}()

	time.Sleep(10000 * time.Second)

	fmt.Println("cancel...")

}

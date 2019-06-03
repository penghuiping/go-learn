package thread

import (
	"runtime"
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
	"math/rand"
	"strconv"
	"testing"
)

type Message struct {
	header string
	body   string
}

func (msg *Message) print() {
	fmt.Println(msg.header)
	fmt.Println(msg.body)
}

func send(ch chan *Message, msg *Message) {
	ch <- msg;
}

func receive(ch chan *Message, fn func(msg *Message)) {
	word, ok := <-ch
	if ok {
		fn(word);
	}
}

func TestChannelExample(in *testing.T) {
	//设置线程数量，类似定义一个线程池
	runtime.GOMAXPROCS(2)

	ch := make(chan *Message)
	defer close(ch);
	go func() {
		for {
			msg := new(Message);
			msg.header = "head_" + strconv.FormatInt(rand.Int63(), 10);
			fmt.Println(time.Now().Format("2006-01-02 15:04:05")+"  send a message")
			send(ch, msg)
			time.Sleep(1 * time.Second);
		}

	}()

	go func() {
		for ; ; {
			receive(ch, func(msg *Message) {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05")+"  receive a message")
				fmt.Println(msg.header)
				fmt.Println()
			})
			time.Sleep(10 * time.Second);
		}
	}()

	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入exit关闭")
	for ; ; {
		var value, err = reader.ReadString('\n')
		if err == nil {
			fmt.Println(value)
			if strings.Trim(value, "\n") == "exit" {
				fmt.Println("将要关闭")
				break
			}
		}
	}
}

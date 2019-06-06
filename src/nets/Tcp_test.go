package nets

import (
	"net"
	"fmt"
	"testing"
	"strings"
)

var conns chan net.Conn

func TestTcpServerExample(in *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("出错啦!", err)
		return
	}
	conns = make(chan net.Conn)

	for {
		con, err := lis.Accept()
		if err != nil {
			fmt.Println("出错啦!", err)
			continue
		}
		go conn(con)
	}
}

func TestTcpClientExample(in *testing.T) {
	var i int
	for i = 0; i < 1000; i++ {
		go func() {
			conn, err := net.Dial("tcp", "127.0.0.1:8080")
			if nil != err {
				fmt.Println("连接失败，错误为:", err)
				return
			}
			for {
				buf := make([]byte, 32)
				count, err := conn.Read(buf)
				if err != nil {
					continue
				} else {
					fmt.Println(string(buf[0:count]))
				}

			}
		}()
	}
	select {}
}

func conn(conn net.Conn) {
	conns <- conn
	fmt.Println("连接成功,当前数量", len(conns))
}

func TestNet(in *testing.T) {
	server()
}

func server() {
	li, _ := net.Listen("tcp", "127.0.0.1:8081")
	for {
		conn, _ := li.Accept()
		go func(c *net.Conn) {
			defer (*c).Close()
			buf := make([]byte, 1024)

			for {
				start := 0;
				index,_:= (*c).Read(buf)
				content := string(buf[start:index])
				fmt.Println(content)
				if strings.Contains(content,"quit") {
					break
				}

			}
		}(&conn)
	}

}

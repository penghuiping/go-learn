package nets

import (
	"net"
	"fmt"
)

var conns chan net.Conn

func TcpServerExample() {
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

func TcpClientExample() {
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

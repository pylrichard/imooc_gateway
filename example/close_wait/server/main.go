package main

import (
	"fmt"
	"net"
)

const BufSize int = 128

func main() {
	//监听9090端口
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}
	for {
		//为客户端创建连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}
		go func(conn net.Conn) {
			defer conn.Close()
			for {
				var buf [BufSize]byte
				n, err := conn.Read(buf[:])
				if err != nil {
					fmt.Printf("read from connect failed, err: %v\n", err)
					break
				}
				str := string(buf[:n])
				fmt.Printf("receive from client, data: %v\n", str)
			}
		}(conn)
	}
}


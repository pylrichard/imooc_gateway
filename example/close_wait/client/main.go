package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main()  {
	doSend()
	fmt.Print("doSend over")
}


func doSend() {
	//连接服务端
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	//读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)
	for {
		//一直读取直到读到\n
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err: %v\n", err)
			break
		}
		//读取到Q或者q时终止客户端
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" || trimmedInput == "q" {
			break
		}
		//将输入发送给服务端
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	/*
		创建连接池
	 */
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			//连接超时
			Timeout:   30 * time.Second,
			//探活时间
			KeepAlive: 30 * time.Second,
		}).DialContext,
		//最大空闲连接
		MaxIdleConns:          100,
		//空闲超时时间
		IdleConnTimeout:       90 * time.Second,
		//tls握手超时时间
		TLSHandshakeTimeout:   10 * time.Second,
		//continue状态码超时时间
		ExpectContinueTimeout: 1 * time.Second,
	}
	/*
		创建客户端
	 */
	client := &http.Client{
		//请求超时时间
		Timeout:   30 * time.Second,
		Transport: transport,
	}
	//请求数据
	resp, err := client.Get("http://127.0.0.1:1210/bye")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//读取内容
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
package main

import (
	"fmt"
	"net"
	"testing"
	"time"
	"zinxLearn/znet"
)

/**
模拟客户端
*/

func ClientTest() {

	fmt.Println("Client Test ... start")
	// 3秒之后发起测试请求
	time.Sleep(3 * time.Second)

	dial, err := net.Dial("tcp", "127.0.0.1:8081")

	if err != nil {
		fmt.Println("clint start err, exit!")
		return
	}

	for {
		_, err := dial.Write([]byte("hello zinx"))
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		bytes := make([]byte, 512)

		read, err := dial.Read(bytes)
		if err != nil {
			fmt.Println(" read buf error: ", bytes, " ", read)
			return
		}

		fmt.Println(" server call back ", string(bytes[:read]), read)
		time.Sleep(time.Second)

	}

}

func TestServer(t *testing.T) {
	/**
	服务器测试
	*/

	server := znet.NewServer("myZinx")

	go ClientTest()

	server.Server()

}

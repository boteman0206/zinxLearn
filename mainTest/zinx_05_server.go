package main

import (
	"fmt"
	"zinxLearn/ziface"
	"zinxLearn/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgId, ", data=", string(request.GetData()))

	//回写数据
	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	// 创建服务器的句柄
	server := znet.NewServer("zinx_05")

	// 添加路由
	router := PingRouter{}
	server.AddRouter(&router)

	server.Server()
}

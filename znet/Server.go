package znet

import (
	"fmt"
	"net"
)

//定义实现IServer实现的结构体
type Server struct {

	//服务名称
	Name string
	//tcp4或者其他服务协议
	IPVersion string

	// ip
	IP string

	// 端口
	Port int
}

//实现IServer接口的所有方法
//启动
func (s Server) Start() {

}

//停止
func (s Server) Stop() {

}

//开启服务
func (s Server) Server() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	//开启一个go去做服务端的Lister业务
	go func() {
		//1： 获取一个tcp的连接addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprint("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		//2：监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp err: ", err)
			return
		}

		// 已经监听成功
		fmt.Println("start zinx server: ", s.Name, " succ, now listing...")

		// 3：启动server网络连接业务
		for {
			// 3.1: 阻塞等待客户端连接
			accept, err := listenner.Accept()

			if err != nil {
				fmt.Println("accept err : ", err)
				continue
			}

		}

	}()

}

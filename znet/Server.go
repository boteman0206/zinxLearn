package znet

import (
	"fmt"
	"net"
	"time"
	"zinxLearn/ziface"
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
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	//开启一个go去做服务端的Lister业务
	go func() {
		//1： 获取一个tcp的连接addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
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

			go func() {
				for {
					buf := make([]byte, 512)
					read, err2 := accept.Read(buf)

					if err2 != nil {
						fmt.Println(" recv buf err: ", err2)
						continue
					}

					// 回显
					if _, err3 := accept.Write(buf[:read]); err3 != nil {
						fmt.Println(" write back err: ", err3)
						continue
					}
				}
			}()
		}

	}()

}

//停止
func (s Server) Stop() {
	fmt.Println(" [stop] zinx server, name ", s.Name)
	// todo Server.stop()
}

//开启服务
func (s Server) Server() {
	s.Start()

	// todo 启动服务，可以处理其他的事情

	// 阻塞
	for {
		time.Sleep(10 * time.Second)
	}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		name,
		"tcp4",
		"0.0.0.0",
		8081,
	}
	return s
}

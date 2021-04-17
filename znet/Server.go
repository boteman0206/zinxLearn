package znet

import (
	"fmt"
	"github.com/gofrs/uuid"
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

	// 04，15 当前Server由用户绑定的会掉router，也就是Server注册的连接对应的业务
	Router ziface.IRouter
}

// --------- 定义当前的客户端连接的handle api-----------------
//func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
//	fmt.Println("[Conn Handle] CallBackToClient....")
//	if _, err := conn.Write(data[:cnt]); err != nil {
//		fmt.Println("write back buf err : ", err)
//		return errors.New("Call back err")
//	}
//	return nil
//}

//-----------------实现 ziface.Iserver里面的全部接口的方法
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
			accept, err := listenner.AcceptTCP()

			if err != nil {
				fmt.Println("accept err : ", err)
				continue
			}

			// 3.2 todo Server.Start() 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接

			// 3.3 处理改新连接请求的业务方法 此时应该有 handler 和 conn是绑定的
			v1, _ := uuid.NewV1()
			conntion := NewConntion(accept, v1.String(), s.Router)

			go conntion.Start()
			// 版本一注释掉
			//		go func() {
			//			for {
			//				buf := make([]byte, 512)
			//				read, err2 := accept.Read(buf)
			//
			//				if err2 != nil {
			//					fmt.Println(" recv buf err: ", err2)
			//					continue
			//				}
			//
			//				// 回显
			//				if _, err3 := accept.Write(buf[:read]); err3 != nil {
			//					fmt.Println(" write back err: ", err3)
			//					continue
			//				}
			//			}
			//		}()
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
		nil,
	}
	return s
}

// todo 这里必须要用指针传递 否则router会报nil空指针， 因为你并没有真正的赋值, 是值拷贝 3.0 添加路由方法
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}

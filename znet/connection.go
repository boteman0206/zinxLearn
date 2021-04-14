package znet

import (
	"fmt"
	"net"
	"zinxLearn/ziface"
)

// 实现 iconnection接口
type Connection struct {

	// 当前的socket tcp连接
	Conn *net.TCPConn

	//当前的ID 也可以成为SessionID 全局唯一
	ConnID string

	// 当前连接的关闭状态
	isClosed bool

	//处理连接的api方法
	handleAPI ziface.HandFunc

	// 告知该连接已经退出/停止的channel
	ExitBuffChan chan bool
}

// 创建连接的方法
func NewConntion(conn *net.TCPConn, connID string, callback_api ziface.HandFunc) *Connection {

	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		handleAPI:    callback_api,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}

//处理当前数据的Groutine
func (c *Connection) StartReader() {
	fmt.Println(" 开始执行 reader Groutine 方法 。。。。")
	defer fmt.Println(c.Conn.RemoteAddr().String(), " 连接 reader exit!")

	defer c.Stop()

	for {
		//读取我们最大的数据
		buf := make([]byte, 512)
		read, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read data error ...")
			continue
		}

		// 调用当前连接业务（执行当前的conn绑定的handle方法）
		err = c.handleAPI(c.Conn, buf, read)
		if err != nil {
			fmt.Println("handler func error...")
			c.ExitBuffChan <- true
			return
		}
	}

}

// 启动连接，让连接开始工作
func (c *Connection) Start() {

	// 开始处理该链接读取客户端数据之后的请求业务
	go c.StartReader()

	for {
		select {
		case <-c.ExitBuffChan:
			// 得到消息退出，不用阻塞
			return
		}
	}

}

// 停止连接，结束当前连接
func (c *Connection) Stop() {
	// 1. 如果当前连接已经关闭
	if c.isClosed == true {
		return
	}

	c.isClosed = true
	//TODO Connection Stop() 如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用

	c.Conn.Close()

	// 通知从缓冲队列读取数据的业务， 该链接已经关闭
	c.ExitBuffChan <- true

	// 关闭连接全部通道
	close(c.ExitBuffChan)

}

// 从当前连接获取原始连接的socket TCPConn
func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

// 从当前的获取连接的ID
func (c *Connection) GetConnID() string {
	return c.ConnID

}

// 获取远程的客户端的地址
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

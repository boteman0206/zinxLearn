package znet

import (
	"fmt"
	"strconv"
	"zinxLearn/ziface"
)

type MsgHandle struct {

	// 存放每一个MsgId锁对应的处理方法的map属性
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

// 马上以非阻塞的方式
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {

	router, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api msgId = ", request.GetMsgId(), "is not found")
	}

	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)

}

// 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	// 1 判断当前的msg绑定的api处理方法是否已经存在
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeated api , msgId = " + strconv.Itoa(int(msgId)))
	}

	// 添加msg与api的绑定关系
	mh.Apis[msgId] = router
	fmt.Println("add api msgIfd = ", msgId)

}

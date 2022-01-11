package api

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
)

//创建一个玩家对象
func AddRouter(s ziface.IServer) {
	fmt.Println("------------ AddRouter ------------")

	s.AddRouter(2, &WorldChatApi{})
	s.AddRouter(3, &MoveApi{})
}

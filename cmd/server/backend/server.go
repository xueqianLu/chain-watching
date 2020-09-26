package backend

import (
	"github.com/hpb-project/chain-watching/types"
	"github.com/labstack/echo"
)

func HttpStart() {
	e := echo.New()
	// 注册路由
	e.GET("/", GetIndex)
	e.GET("/node/:id", GetNodeInfo)
	e.GET("/nodes", GetPeersInfo)

	e.POST(types.NODESTATUS, ReceiveNodeStatus)
	e.POST(types.PEERLIST, ReceivePeerList)
	// 开启 HTTP Server
	e.Logger.Fatal(e.Start(":2020"))
}

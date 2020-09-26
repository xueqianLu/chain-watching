package backend

import (
	"encoding/json"
	"github.com/hpb-project/chain-watching/cmd/server/database"
	"github.com/hpb-project/chain-watching/types"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

// e.GET("/node/:id", getNodeInfo)
func GetNodeInfo(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	return c.String(http.StatusOK, id)
}

func GetIndex(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func GetPeersInfo(c echo.Context) error {
	return c.String(http.StatusOK, "get peer info")
}

func ReceivePeerList(c echo.Context) error {
	defer c.Request().Body.Close()

	data, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusOK, "")
	} else {
		var peerList = make([]types.PeerInfo, 0)
		if err = json.Unmarshal(data, &peerList); err == nil {
			for _, info := range peerList {
				database.SavePeerInfo(&info)
			}
		}
	}
	return c.String(http.StatusOK, "")
}

func ReceiveNodeStatus(c echo.Context) error {
	defer c.Request().Body.Close()

	data, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusOK, "")
	} else {
		var nodeStatus types.Status
		if err = json.Unmarshal(data, &nodeStatus); err == nil {
			database.SaveNodeStatus(&nodeStatus)
		}
	}
	return c.String(http.StatusOK, "")
}

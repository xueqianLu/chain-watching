package hpbclient

import (
	"encoding/json"
	"github.com/hpb-project/chain-watching/types"
)

type PeerNetwork struct {
	Local  string `json:"local"`
	Remote string `json:"remote"`
}

type RPCPeerInfo struct {
	PeerID      string      `json:"id"`
	Coinbase    string      `json:"coinbase"`
	Network     PeerNetwork `json:"network"`
	NodeType    string      `json:"remote"`
	Version     string      `json:"version"`
	Mining      string      `json:"mining"`
	ConnectTime string      `json:"start"`
}

func (p RPCPeerInfo) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}

func (p RPCPeerInfo) ToPeerInfo() *types.PeerInfo {
	pinfo := &types.PeerInfo{}
	pinfo.PeerID = p.PeerID
	pinfo.Coinbase = p.Coinbase
	pinfo.ConnectTime = ParseRPCTime(p.ConnectTime)
	pinfo.NodeType = p.NodeType
	pinfo.RemoteIp = p.Network.Remote
	pinfo.GHPBVersion, pinfo.BOEVersion = ParseRPCVersion(p.Version)

	return pinfo
}

type RCPQuest interface {
	body() string
}

type AdminPeersResponse struct {
	Jsonrpc   string        `json:"jsonrpc"`
	Id        int           `json:"id"`
	PeersInfo []RPCPeerInfo `json:"result"`
}

type AdminPeers struct {
	Params  []string `json:"params"`
	Id      int      `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
}

func NewAdminPeers(id int) *AdminPeers {
	return &AdminPeers{
		Method: "admin_peers",
		Id:     id,
	}
}

func (p *AdminPeers) body() string {
	datas := make([]interface{}, 0)
	datas = append(datas, *p)

	body, _ := json.Marshal(datas)
	return string(body)
}

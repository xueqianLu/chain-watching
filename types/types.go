package types

import "encoding/json"

type Status struct {
	PeerID        string `json:"peerid"`
	CurrentHeader string `json:"currentHeader"`
	SyncStatus    string `json:"syncStatus"`
	PeerNumber    string `json:"peerNumber"`
	BOEVersion    string `json:"boeVersion"`
	GHPBVersion   string `json:"ghpbVersion"`
	Mining        string `json:"mining"`
	Starttime     string `json:"starttime"`
}

type PeerInfo struct {
	PeerID      string `json:"peerid"`
	Coinbase    string `json:"coinbase"`
	RemoteIp    string `json:"remoteip"`
	BOEVersion  string `json:"boeVersion"`
	GHPBVersion string `json:"ghpbVersion"`
	ConnectTime string `json:"connectTime"`
	NodeType    string `json:"nodetype"`
}

func (p PeerInfo) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}

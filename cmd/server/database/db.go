package database

import (
	"github.com/hpb-project/chain-watching/database"
	"github.com/hpb-project/chain-watching/types"
)

const (
	nodeStatusKeyPrefix = "ns"
	peerInfoKeyPrefix   = "pi"
	currentPeersListKey = "peer"
)

func SavePeerInfo(info *types.PeerInfo) {
	pid := info.PeerID
	key := peerInfoKeyPrefix + pid
	db := database.GetConnect()
	defer db.Close()
	db.Do("SET", key, info.String())
}

func SaveNodeStatus(status *types.Status) {
	pid := status.PeerID
	key := nodeStatusKeyPrefix + pid
	db := database.GetConnect()
	defer db.Close()
	db.Do("SET", key, status.String())
}

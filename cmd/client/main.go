package main

import (
	"encoding/json"
	"github.com/hpb-project/chain-watching/cmd/client/hpbclient"
	"sync"
	"time"
)

func watchingPeers(client *hpbclient.Client) {
	var id = 10
	questPeers := hpbclient.NewAdminPeers(id)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			response, err := client.Do(questPeers)
			if err != nil {
				println("request failed ", err)
				continue
			}
			var res = make([]hpbclient.AdminPeersResponse, 0)
			err = json.Unmarshal(response, &res)
			if err != nil {
				println("unmarshal failed ", err)
				continue
			}

			for _, r := range res {
				for _, p := range r.PeersInfo {
					println("peer:", p.String())
					tp := p.ToPeerInfo()
					println("transfer to type.PeerInfo:", tp.String())
					// Todo: for test.
					break
				}
			}
		}
	}
}

func watchStatus(client *hpbclient.Client) {
	var id = 20
	for {
		request := hpbclient.NewAdminPeers(id)
		response, err := client.Do(request)
		if err != nil {
			println("request failed ", err)
			return
		}
		//println("got res :", string(response))
		var res = make([]hpbclient.AdminPeersResponse, 0)
		err = json.Unmarshal(response, &res)
		if err != nil {
			println("unmarshal failed ", err)
			return
		}

		for _, r := range res {
			for _, p := range r.PeersInfo {
				println("peer:", p.String())
			}
		}

	}
}

func main() {
	client := hpbclient.NewClient("http://127.0.0.1:8545")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go watchingPeers(client)
	wg.Wait()
	//go watchStatus(client)
}

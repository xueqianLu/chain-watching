package main

import (
	"encoding/json"
	"github.com/hpb-project/chain-watching/cmd/client/hpbclient"
	"log"
	"sync"
	"time"
)

func watchingPeers(client *hpbclient.Client) {
	var id = 10
	questPeers := hpbclient.NewAdminPeers(id)

	peersticker := time.NewTicker(time.Second)
	statusticker := time.NewTicker(time.Second)
	defer peersticker.Stop()
	defer statusticker.Stop()
	for {
		select {
		case <-statusticker.C:
			//Todo: query peer status.

		case <-peersticker.C:
			response, err := client.Do(questPeers)
			if err != nil {
				log.Println("request failed ", err)
				continue
			}
			var res = make([]hpbclient.AdminPeersResponse, 0)
			err = json.Unmarshal(response, &res)
			if err != nil {
				log.Println("unmarshal failed ", err)
				continue
			}

			datas := make([]interface{}, 0)

			for _, r := range res {
				for _, p := range r.PeersInfo {
					tp := p.ToPeerInfo()
					datas = append(datas, tp)
					//log.Println("transfer to type.PeerInfo:", tp.String())
				}
			}
			if body, e := json.Marshal(datas); e == nil && len(datas) > 0 {
				e = client.PostPeerInfo(string(body))
				if e != nil {
					log.Println("Post to Server failed:", e.Error())
				} else {
					log.Println("Post to Server finished.")
				}
			}
		}
	}
}

func main() {
	client := hpbclient.NewClient("http://127.0.0.1:8545", "http://127.0.0.1:2020")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go watchingPeers(client)
	wg.Wait()
	//go watchStatus(client)
}

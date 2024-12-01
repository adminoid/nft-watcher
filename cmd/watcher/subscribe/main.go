package main

import (
	"github.com/adminoid/nft-watcher/internal/config"
	"github.com/adminoid/nft-watcher/internal/watcher/subscribe"
)

func main() {
	cfg, _ := config.NewConfig()
	//fmt.Println(cfg.TONCONSOLE_KEY)
	go subscribe.ToTransactions(cfg.TONCONSOLE_KEY)

	//go SubscribeToTraces(token)
	//go SubscribeToMempool(token)
	//go SubscribeToBlocks(token)
	select {}
}

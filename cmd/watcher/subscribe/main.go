package main

import "github.com/adminoid/nft-watcher/internal/watcher/subscribe"

func main() {
	// When working with tonapi.io, you should consider getting an API key at https://tonconsole.com/
	// because tonapi.io has per-ip limits for sse and websocket connections.
	//
	// You can configure it with:
	//         streamingAPI = tonapi.NewStreamingAPI(tonapi.WithStreamingToken("<private-key>"))
	//
	// To work with a local version of tonapi.io (opentonapi) use:
	//         streamingAPI = tonapi.NewStreamingAPI(tonapi.WithStreamingEndpoint("http://127.0.0.1:8081"))
	//
	token := "AHQGYRXC2JAB2TQAAAAG37FZFQA37NPCMJ65A6UDDTLZHPHJGWO7BZZYOX7I4OXE5MPT3DA"

	//go SubscribeToTraces(token)
	//go SubscribeToMempool(token)
	go subscribe.ToTransactions(token)
	//go SubscribeToBlocks(token)
	select {}
}

package subscribe

import (
	"context"
	"fmt"
	"github.com/tonkeeper/tonapi-go"
)

var addresses = []string{
	"0:06d811f426598591b32b2c49f29f66c821368e4acb1de16762b04e0174532465",
	"0:80d78a35f955a14b679faa887ff4cd5bfc0f43b4a4eea2a7e6927f3701b273c2",
	"0:0e41dc1dc3c9067ed24248580e12b3359818d83dee0304fabcf80845eafafdb2",
}

func ToTransactions(token string) {
	streamingAPI := tonapi.NewStreamingAPI(tonapi.WithStreamingToken(token))
	for {
		err := streamingAPI.SubscribeToTransactions(context.Background(),
			addresses,
			// this "operations" is optional,
			// if not set, you will receive all transactions.
			// if defined, you will receive only transactions with these operations.
			nil,
			func(data tonapi.TransactionEventData) {
				fmt.Printf("New tx with hash: %v\n", data.TxHash)
				//fmt.Printf("%+v\n", data)
				fmt.Printf("%#v\n", data)
				s1 := data.AccountID.Workchain
				fmt.Println(s1)
				//s2 := string(bytes(data.AccountID.Address))
				//s2 := fmt.Sprintf("%s", data.AccountID.Address)
				//fmt.Println(s2)
				fmt.Printf("%x\n", data.AccountID.Address)
				fmt.Println(data.TxHash)
				fmt.Println("----------")
			})
		if err != nil {
			fmt.Printf("tx error: %v, reconnecting...\n", err)
		}
	}
}

//func SubscribeToMempool(token string) {
//	streamingAPI := tonapi.NewStreamingAPI(tonapi.WithStreamingToken(token))
//	for {
//		err := streamingAPI.SubscribeToMempool(context.Background(),
//			// this "accounts" parameter is optional,
//			// if not set, you will receive all mempool events.
//			// if defined, you will receive only mempool events that involve these accounts.
//			addresses,
//			func(data tonapi.MempoolEventData) {
//				value, _ := json.Marshal(data)
//				fmt.Printf("mempool event: %#v\n", value)
//			})
//		if err != nil {
//			fmt.Printf("mempool error: %v, reconnecting...\n", err)
//		}
//	}
//}

//func SubscribeToTraces(token string) {
//	streamingAPI := tonapi.NewStreamingAPI(tonapi.WithStreamingToken(token))
//	for {
//		err := streamingAPI.SubscribeToTraces(context.Background(), addresses,
//			func(data tonapi.TraceEventData) {
//				fmt.Printf("New trace with hash: %v\n", data.Hash)
//				//fmt.Printf("%+v\n", data)
//				fmt.Printf("%#v\n", data)
//			})
//		if err != nil {
//			fmt.Printf("trace error: %v, reconnecting...\n", err)
//		}
//	}
//}

//func intPointer(x int) *int {
//	return &x
//}
//func SubscribeToBlocks(token string) {
//	streamingAPI := tonapi.NewStreamingAPI(tonapi.WithStreamingToken(token))
//	for {
//		err := streamingAPI.SubscribeToBlocks(context.Background(), intPointer(-1),
//			func(data tonapi.BlockEventData) {
//				fmt.Printf("New block: (%v,%v,%v)\n", data.Workchain, data.Shard, data.Seqno)
//			})
//		if err != nil {
//			fmt.Printf("block error: %v, reconnecting...\n", err)
//		}
//	}
//}

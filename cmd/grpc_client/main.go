package main

import (
	"log"

	"benchprotos/grpc/proto/pb"
	"benchprotos/grpc_client/client"
)

func main()  {

	quotes := []*pb.Quote{{Ask: "ETH", Bid: "BTC", Price: 123}}
	if err := client.QuotesUpdate(quotes); err != nil {
		log.Println("quotes update failed", err)
	}
}

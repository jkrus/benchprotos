package main

import (
	"benchprotos/grpc/proto/pb"
	"benchprotos/grpc_client/client"
)

func main()  {

	quotes := []*pb.Quote{{Ask: "ETH", Bid: "BTC", Price: 123}}
	client.QuotesUpdate(quotes)

}

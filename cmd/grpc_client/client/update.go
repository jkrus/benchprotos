// Copyright © 2020-2021 The EVEN Solutions Developers Team

package client

import (
	"fmt"
	"log"

	"benchprotos/grpc/proto/pb"
)

// QuotesUpdate sends a list of quotes.
func QuotesUpdate(quotes []*pb.Quote) error {
	defer terminate()
	if err := connect(); err != nil {
		return fmt.Errorf("rpc connect: %w", err)
	}

	request := &pb.QuotesUpdateRequest{Quote: quotes}
	response, err := pb.NewQuotesServiceClient(con).QuotesUpdate(ctx, request)
	if err != nil {
		return fmt.Errorf("send list of quotes: %w", err)
	}

	log.Println(response.Response)

	return nil
}

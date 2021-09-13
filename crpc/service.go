package crpc

import (
	"log"

	"benchprotos/crpc/cpb"
)

// QuotesUpdate implements service api interface.
func (s *capnservice) QuotesUpdate(update cpb.QuotesUpdate_quotesUpdate) error {
	quote, _ := update.Params.List()
	log.Println(quote)
	if err := update.Results.SetResult("Update success"); err != nil {
		log.Println("quotes update set result failed", err)
	}

	return nil
}

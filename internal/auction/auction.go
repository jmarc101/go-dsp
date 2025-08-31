package auction

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jmarc101/go-dsp/pkg/openrtb"
	"github.com/prebid/openrtb/v20/openrtb2"
)

var (
	ErrInvalidRequest = errors.New("invalid request") // malformed or missing required fields
	ErrNoBid          = errors.New("no bid")          // nothing eligible
	ErrTimeout        = errors.New("auction time out")
	ErrNoCreatives    = errors.New("no eligible creatives")
)

func Run(ctx context.Context, br *openrtb.BidRequest) (openrtb.BidRepsonse, error) {
	log.Printf("Starting auction for bid: %s", br.ID)

	return openrtb2.BidResponse{
		ID:    br.ID,
		BidID: uuid.New().String(),
	}, nil
}

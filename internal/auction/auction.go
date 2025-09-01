package auction

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
	ErrTimeout        = errors.New("auction time out")
	ErrNoCreatives    = errors.New("no eligible creatives")
)

func Run(ctx context.Context, br *BidRequest) (BidResponse, error) {
	log.Printf("Starting auction for bid: %s", br.ID)

	return BidResponse{
		ID:    br.ID,
		BidID: uuid.New().String(),
	}, nil
}

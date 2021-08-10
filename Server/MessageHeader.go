package Server

import (
	"context"

	protos "/home/scot/Desktop/grpc/protos/MessageHeader"

	"github.com/hashicorp/go-hclog"
)

type MessageHeader struct {
	log hclog.Logger
}

// NewCurrency creates a new Currency server
func NewMessageHeader(l hclog.Logger) *MessageHeader {
	return &MessageHeader{l}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *MessageHeader) Send(ctx context.Context, rr *protos.EdgeRequest) (*protos.EdgeResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.Gethdr(), "dest", rr.Gettext())
	return &protos.EdgeResponse{text: 1}, nil
}

package binance

import (
	"context"
	"github.com/aiviaio/go-binance/v2"
)

type Client struct {
	client *binance.Client
}

func NewClient() *Client {
	return &Client{
		client: binance.NewClient("", ""),
	}
}

func (c *Client) GetSymbols() ([]string, error) {
	info, err := c.client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	var symbols []string
	for i, info := range info.Symbols {
		if i >= 5 {
			break
		}
		symbols = append(symbols, info.Symbol)
	}
	return symbols, nil
}

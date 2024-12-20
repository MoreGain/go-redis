package redis

import "context"

type Client struct {
	ctx context.Context
	opt *Options
}

func NewClient(opt *Options) *Client {
	opt.init()
	return &Client{
		ctx: context.Background(),
		opt: opt,
	}
}

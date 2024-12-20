package redis

import (
	"context"
	"crypto/tls"
	"net"
	"strings"
	"time"
)

type Options struct {
	Network  string
	Addr     string
	Dialer   func(ctx context.Context, network, addr string) (net.Conn, error)
	Password string
	DB       int

	DialTimeout time.Duration
	TLSConfig   *tls.Config
}

func (opt *Options) init() {
	if opt.Addr == "" {
		opt.Addr = "127.0.0.1:6379"
	}
	if opt.Network == "" {
		if strings.HasPrefix(opt.Addr, "/") {
			opt.Network = "unix"
		} else {
			opt.Network = "tcp"
		}
	}
	if opt.Dialer == nil {
		opt.Dialer = NewDialer(opt)
	}
}

func NewDialer(opt *Options) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		netDialer := net.Dialer{
			Timeout:   opt.DialTimeout,
			KeepAlive: 5 * time.Minute,
		}
		if opt.TLSConfig == nil {
			return netDialer.DialContext(ctx, network, addr)
		}
		return tls.DialWithDialer(&netDialer, network, addr, opt.TLSConfig)
	}
}

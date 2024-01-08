package tcp

import (
	"fmt"
	"net"
)

type MustListenOptions struct {
	Addr *string
	Port *string
}

type MustListenOption func(*MustListenOptions)

// WithAddr sets the address to listen on.
func WithAddr(addr string) MustListenOption {
	return func(o *MustListenOptions) {
		o.Addr = &addr
	}
}

// WithPort sets the port to listen on.
func WithPort(port string) MustListenOption {
	return func(o *MustListenOptions) {
		o.Port = &port
	}
}

// MustListen returns a net.Listener.
//
// If the listener cannot be created, it panics.
func MustListen(opts ...MustListenOption) net.Listener {
	var options MustListenOptions
	for _, opt := range opts {
		opt(&options)
	}

	addr := ""
	if options.Addr != nil {
		addr = *options.Addr
	}
	port := "8080"
	if options.Port != nil {
		port = *options.Port
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", addr, port))
	if err != nil {
		panic(err)
	}
	return listener
}

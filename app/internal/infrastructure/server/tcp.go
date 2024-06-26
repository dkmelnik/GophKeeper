// Package server provides functionality for a TCP server.
package server

import (
	"net"
)

// TCPServer represents a TCP server instance.
type TCPServer struct {
	listener *net.TCPListener
}

// NewTCPServer creates a new TCPServer instance bound to the specified address.
// It resolves the address and initializes a TCP listener.
// Returns a pointer to TCPServer instance and nil error on success, or nil and error on failure.
func NewTCPServer(address string) (*TCPServer, error) {
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &TCPServer{listener}, nil
}

// Listener returns the TCPListener associated with the TCPServer instance.
func (t *TCPServer) Listener() *net.TCPListener {
	return t.listener
}

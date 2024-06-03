package server

import "time"

type Server struct {
	host    string
	port    int
	timeout time.Duration
	// New requirement
	maxConn int
}

func WithPort(port int) func(server *Server) {
	return func(server *Server) {
		server.port = port
	}
}

func WithHost(host string) func(server *Server) {
	return func(server *Server) {
		server.host = host
	}
}

func WithTimeout(timeout time.Duration) func(server *Server) {
	return func(server *Server) {
		server.timeout = timeout
	}
}

func WithMaxConn(maxConn int) func(server *Server) {
	return func(server *Server) {
		server.maxConn = maxConn
	}
}
func NewServer(options ...func(server *Server)) *Server {
	server := &Server{}
	for _, opt := range options {
		opt(server)
	}
	return server
}

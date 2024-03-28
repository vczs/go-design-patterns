package main

import "fmt"

type Server struct {
	Id      string
	MaxConn int
	Tls     bool
}

type Option func(*Server)

func WithId(id string) Option {
	return func(s *Server) {
		if id != "" {
			s.Id = id
		}
	}
}

func WithMaxConn(maxConn int) Option {
	return func(s *Server) {
		if maxConn > 0 {
			s.MaxConn = maxConn
		}
	}
}

func WithTls(tls bool) Option {
	return func(s *Server) {
		s.Tls = tls
	}
}

func NewServer(opt ...Option) *Server {
	s := &Server{Id: "000", MaxConn: 100, Tls: false}
	for _, o := range opt {
		o(s)
	}
	return s
}

func main() {
	fmt.Printf("%+v", NewServer(WithId("666"), WithMaxConn(200), WithTls(true)))
}

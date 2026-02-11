package main

import (
	"github.com/Quavke/market/internal/network"
)

func main() {
	cfg := network.NewServerConfig()
	s := network.NewServer(cfg)
	if err := s.Start(); err != nil {
		panic(err)
	}
}

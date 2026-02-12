package main

import (
	"fmt"
	"os"

	"github.com/Quavke/market/internal/network"
)

func main() {
	cfg := network.NewServerConfig()
	s, err := network.NewServer(cfg)
	if err != nil {
		fmt.Printf("Failed to create server: %v\n", err)
		os.Exit(1)
	}
	if err := s.Start(); err != nil {
		panic(err)
	}
}

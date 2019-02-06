package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/anjmao/wic/server"
	"github.com/anjmao/wic/services/winter"
	game "github.com/anjmao/wic/game"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "Port to listen on")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()

	s, err := server.New()
	if err != nil {
		logrus.Fatalf("could not create server: %v", err)
	}

	s.Register(func(gs *grpc.Server) {
		winter.RegisterService(gs, game.NewController())
	})

	if err := s.Run(*port); err != nil {
		logrus.Fatalf("could not run server: %v", err)
	}
}

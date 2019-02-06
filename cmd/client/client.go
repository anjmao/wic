package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/anjmao/wic/api/winter"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	target     = flag.String("target", "127.0.0.1:8080", "Game server target address")
	serverName = flag.String("serverName", "server", "Game server name")
	crt        = flag.String("crt", "certs/client.crt", "Client TLS public key path")
	key        = flag.String("key", "certs/client.key", "Client TLS private key path")
	ca         = flag.String("ca", "certs/ca.crt", "Client TLS ca path")
)

var (
	errInvalidArgs    = errors.New("invalid args")
	errServerNotReady = errors.New("server is not ready")
	errGameNotStarted = errors.New("game is not started")
)

var help = `
START <player_name> - Start a new game.
LIST - List running games.
JOIN <game_id> <player_name> - Join existing game.
SHOOT <x> <y> - Hit a zombie.
EXIT - Close game client.
	`

func main() {
	flag.Parse()
	client, err := newClient(*target, *serverName, *crt, *key, *ca)
	if err != nil {
		logrus.Fatalf("could not create secure client: %v", err)
	}

	fmt.Print(help)
	fmt.Println("")

	handler := &cmdHandler{client: client}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		handler.handleCmd(scanner.Text())
	}

	if scanner.Err() != nil {
		os.Exit(0)
	}
}

type cmdHandler struct {
	client     winter.WinterGameClient
	playerName string
	stream     winter.WinterGame_PlayClient
}

func (h *cmdHandler) handleCmd(input string) {
	s := strings.Split(input, " ")
	cmd, args := s[0], s[1:]

	var err error
	switch strings.ToUpper(cmd) {
	case "START":
		err = h.handleStartCmd(args)
	case "JOIN":
		err = h.handleJoinCmd(args)
	case "LIST":
		err = h.handleListCmd()
	case "SHOOT":
		err = h.handleShootCmd(args)
	case "EXIT":
		os.Exit(0)
	default:
		logrus.Errorf("unknown command %s", input)
	}

	if err != nil {
		logrus.Error(err)
	}
}

func (h *cmdHandler) handleStartCmd(args []string) error {
	if len(args) != 1 {
		return errInvalidArgs
	}

	playerName := args[0]
	return h.startGame(playerName, "")
}

func (h *cmdHandler) handleJoinCmd(args []string) error {
	if len(args) != 2 {
		return errInvalidArgs
	}

	gameID := args[0]
	playerName := args[1]
	return h.startGame(playerName, gameID)
}

func (h *cmdHandler) startGame(playerName, gameID string) error {
	ctx := context.Background()

	// Create new game stream.
	stream, err := h.client.Play(ctx)
	if err != nil {
		return err
	}
	h.stream = stream
	h.playerName = playerName

	// Make initial request to start a game.
	initialReq := &winter.PlayRequest{
		Start: &winter.Start{
			PlayerName: playerName,
			GameID:     gameID,
		},
	}
	if err := h.stream.Send(initialReq); err != nil {
		return err
	}

	// Receive server ready message which indicates that
	// we can start playing.
	res, err := h.stream.Recv()
	if err != nil {
		return err
	}
	if res.Ready == nil {
		return errServerNotReady
	}
	fmt.Printf("Game %s is ready.\n", res.Ready.GameID)

	// Game is ready to play, start listening to incoming server events.
	go h.listenGameEvents()

	return nil
}

func (h *cmdHandler) handleShootCmd(args []string) error {
	if h.stream == nil {
		return errGameNotStarted
	}

	if len(args) != 2 {
		return errInvalidArgs
	}
	x, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid x: %v", err)
	}
	y, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid y: %v", err)
	}

	req := &winter.PlayRequest{
		ShootAt: &winter.ShootAt{
			X: int32(x),
			Y: int32(y),
		},
	}
	if err := h.stream.Send(req); err != nil {
		return err
	}
	return nil
}

func (h *cmdHandler) handleListCmd() error {
	ctx := context.Background()
	res, err := h.client.ListGames(ctx, &winter.ListGamesRequest{})
	if err != nil {
		return fmt.Errorf("could not list games: %v", err)
	}

	if len(res.Games) == 0 {
		fmt.Println("No games found.")
		return nil
	}

	for _, g := range res.Games {
		fmt.Println(g)
	}
	fmt.Println("")
	return nil
}

func (h *cmdHandler) listenGameEvents()  {
	for {
		res, err := h.stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Errorf("could not handle stream data: %v", err)
			break
		}

		shootRes := res.ShootResult
		if shootRes != nil {
			fmt.Printf("BOOM %s %d %s\n", shootRes.PlayerName, shootRes.Points, shootRes.ZombieName)
			if shootRes.Points > 0 {
				fmt.Println("Game ended.")
			}
		}

		zombie := res.Zombie
		if zombie != nil {
			if zombie.ReachedWall {
				fmt.Printf("Game over. Zombie %s reached the wall\n", zombie.Name)
			}
		}
	}
}

// newClient creates new game GRPC client with Mutual TLS encryption.
func newClient(target, serverName, crt, key, ca string) (winter.WinterGameClient, error) {
	// Load the client certificates from disk.
	certificate, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		return nil, fmt.Errorf("could not load client key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority.
	certPool := x509.NewCertPool()
	caFile, err := ioutil.ReadFile(ca)
	if err != nil {
		return nil, fmt.Errorf("could not read ca certificate: %s", err)
	}

	// Append the certificates from the CA.
	if ok := certPool.AppendCertsFromPEM(caFile); !ok {
		return nil, errors.New("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   serverName,
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS12,
	})

	// Create a connection with the TLS credentials.
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("could not connect to game server on addr %s: %v", target, err)
	}

	client := winter.NewWinterGameClient(conn)
	return client, nil
}

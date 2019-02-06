package winter

import (
	"context"
	api "github.com/anjmao/wic/api/winter"
	"github.com/anjmao/wic/game"
	"log"
	"net"
	"reflect"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

type bufDialer func(string, time.Duration) (net.Conn, error)

func startServer(gameCtrl game.Controller) bufDialer {
	lis := bufconn.Listen(bufSize)
	s := grpc.NewServer()
	RegisterService(s, gameCtrl)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	return func(string, time.Duration) (net.Conn, error) {
		return lis.Dial()
	}
}

func TestGamePlay(t *testing.T) {
	ctx := context.Background()
	dialer := startServer(newMockGameController())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := api.NewWinterGameClient(conn)
	stream, err := client.Play(ctx)

	playerName := "trump"
	zombieName := "night-king"

	// Test initial game start request.
	req := &api.PlayRequest{Start:&api.Start{PlayerName:playerName}}
	err = stream.Send(req)
	if err != nil {
		t.Fatalf("could not send play request: %v", err)
	}
	reply, err := stream.Recv()
	if err != nil {
		t.Fatalf("could not start game: %v", err)
	}
	if reply.Ready == nil || reply.Ready.GameID == "" {
		t.Fatalf("expected game ready status, got %#v", reply.Ready)
	}

	// Test shoot miss.
	shootRes := shoot(t, stream, 0 , 1)
	expectedShootRes := &api.ShootResult{
		ZombieName:zombieName,
		PlayerName:playerName,
		Points:0,
	}
	if ok := reflect.DeepEqual(expectedShootRes, shootRes); !ok {
		t.Fatalf("expected shoot %v, got %v", expectedShootRes, shootRes)
	}

	// Test shoot hit.
	shootRes = shoot(t, stream, 0 , 0)
	expectedShootRes = &api.ShootResult{
		ZombieName:zombieName,
		PlayerName:playerName,
		Points:1,
	}
	if ok := reflect.DeepEqual(expectedShootRes, shootRes); !ok {
		t.Fatalf("expected shoot %v, got %v", expectedShootRes, shootRes)
	}
}

func shoot(t *testing.T, stream api.WinterGame_PlayClient, x, y int32) *api.ShootResult {
	req := &api.PlayRequest{ShootAt:&api.ShootAt{X:x, Y:y}}
	err := stream.Send(req)
	if err != nil {
		t.Fatalf("could not send shoot request: %v", err)
	}

	reply, err := stream.Recv()
	if err != nil {
		t.Fatalf("could not get shoot result: %v", err)
	}
	return reply.ShootResult
}

type mockGameController struct {
	g *game.Game
}

func newMockGameController() game.Controller {
	return &mockGameController{}
}

func (m *mockGameController) StartNewGame(config *game.Config) *game.Game {
	m.g = game.NewGame("g1", config)
	go m.g.Run()
	return m.g
}

func (m *mockGameController) GetGame(id string) *game.Game {
	return m.g
}

func (mockGameController) GetGameIDs() []string {
	return []string{"g1"}
}
func (mockGameController) DeleteEndedGame(id string) {}



package winter

import (
	"context"
	api "github.com/anjmao/wic/api/winter"
	"github.com/anjmao/wic/game"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errGameNotFound               = status.Error(codes.NotFound, "Game not found")
	errMissingPlayerName          = status.Error(codes.InvalidArgument, "PlayerName is required")
	errMissingInitialStartRequest = status.Error(codes.InvalidArgument, "Missing game start request")
)

const shootsChanBuffer = 100

// RegisterService registers winter service implementation.
func RegisterService(s *grpc.Server, controller game.Controller) {
	srv := &service{
		controller: controller,
	}
	api.RegisterWinterGameServer(s, srv)
}

type service struct {
	controller game.Controller
}

func (s *service) ListGames(ctx context.Context, in *api.ListGamesRequest) (*api.ListGamesReply, error) {
	ids := s.controller.GetGameIDs()
	var games []string
	for _, id := range ids {
		games = append(games, id)
	}
	return &api.ListGamesReply{Games: games}, nil
}

func (s *service) Play(stream api.WinterGame_PlayServer) error {
	ctx := stream.Context()

	// Receive initial client request with start data.
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	if err := s.validateStart(req.Start); err != nil {
		return err
	}

	// Start new game or join existing one.
	g := s.gameFrom(req.Start)
	if g == nil {
		return errGameNotFound
	}
	player := g.JoinPlayer(req.Start.PlayerName)

	// Notify client that game is ready.
	ready := &api.PlayReply{Ready: &api.Ready{GameID: g.ID()}}
	if err := stream.Send(ready); err != nil {
		return err
	}

	// Start listening to client shootsCh.
	shootsCh := make(chan *api.ShootAt, shootsChanBuffer)
	go func() {
		for {
			recv, err := stream.Recv()
			if err != nil {
				// Client died.
				return
			}
			if recv.ShootAt != nil {
				shootsCh <- recv.ShootAt
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			g.RemovePlayer(player.Name())
			s.controller.DeleteEndedGame(g.ID())
			return ctx.Err()
		case shoot := <-shootsCh:
			g.ShootZombie(player, game.Coord(shoot.X), game.Coord(shoot.Y))
		case event := <-player.Notify():
			reply := toProtoPlayReply(event)
			if err := stream.Send(reply); err != nil {
				return err
			}
		}
	}
}

func (s *service) validateStart(start *api.Start) error {
	if start == nil {
		return errMissingInitialStartRequest
	}
	if start.PlayerName == "" {
		return errMissingPlayerName
	}
	return nil
}

func (s *service) gameFrom(req *api.Start) *game.Game {
	if req.GameID == "" {
		return s.controller.StartNewGame(game.DefaultGameConfig)
	}
	return s.controller.GetGame(req.GameID)
}

func toProtoPlayReply(event game.Event) *api.PlayReply {
	switch e := event.(type) {
	case *game.ZombieMoveEvent:
		return &api.PlayReply{
			Zombie: &api.Zombie{
				Name: e.Name(),
				X: int32(e.X()),
				Y: int32(e.Y()),
			},
		}
	case *game.ZombieReachedWallEvent:
		return &api.PlayReply{
			Zombie: &api.Zombie{
				Name: e.Name(),
				X: int32(e.X()),
				Y: int32(e.Y()),
				ReachedWall:true,
			},
		}
	case *game.ShootResultEvent:
		return &api.PlayReply{
			ShootResult: &api.ShootResult{
				PlayerName:e.Player.Name(),
				Points:e.Points,
				ZombieName:e.Zombie.Name(),
			},
		}
	default:
		return &api.PlayReply{}
	}
}


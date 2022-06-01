package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/murat/codenames/internal/game"
	pb "github.com/murat/codenames/protos"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

// holds states of the games
type db struct {
	games []*pb.Game
	mu    sync.RWMutex
}

type codenamesServer struct {
	pb.UnimplementedCodenamesServer
	db db
}

func (c *codenamesServer) CreateGame(ctx context.Context, request *pb.GameRequest) (*pb.GameResponse, error) {
	g, err := game.NewGame(request.File)
	if err != nil {
		return nil, fmt.Errorf("could not create game, %w", err)
	}

	pbg := &pb.Game{
		ID:       uuid.NewV4().String(),
		Board:    g.ConvertToRequest(),
		RedTeam:  &pb.Team{},
		BlueTeam: &pb.Team{},
	}

	c.db.games = append(c.db.games, pbg)

	return &pb.GameResponse{
		Game: pbg,
	}, nil
}

func (c *codenamesServer) JoinGame(ctx context.Context, request *pb.JoinGameRequest) (*pb.Game, error) {
	var found *pb.Game
	for _, g := range c.db.games {
		if g.ID == request.GameID {
			found = g
		}
	}

	var team *pb.Team
	switch request.Team {
	case pb.TEAM_RED:
		team = found.RedTeam
	case pb.TEAM_BLUE:
		team = found.BlueTeam
	default:
		return nil, fmt.Errorf("unknown team")
	}

	team.Players = append(team.Players, &pb.Player{Name: request.Name})

	return found, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("could not create listener, %v", err)
	}

	log.Println("listening on :8080")

	gs := grpc.NewServer()
	pb.RegisterCodenamesServer(gs, &codenamesServer{})
	if err := gs.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

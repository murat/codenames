package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/murat/codenames/internal/game"
	pb "github.com/murat/codenames/protos"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

type codenamesServer struct {
	pb.UnimplementedCodenamesServer

	// holds states of the games
	db []*pb.Game
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

	c.db = append(c.db, pbg)

	return &pb.GameResponse{
		Game: pbg,
	}, nil
}

func (c *codenamesServer) JoinGame(ctx context.Context, request *pb.JoinGameRequest) (*empty.Empty, error) {
	return nil, fmt.Errorf("not implemented yet")
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

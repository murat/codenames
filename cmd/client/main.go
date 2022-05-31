package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/murat/codenames/protos"
	"google.golang.org/grpc"
)

var (
	port, command, file, gameID, player, team string
)

func main() {
	flag.StringVar(&port, "port", "8080", "-port 8080")
	flag.StringVar(&command, "command", "join", "-command join (available commands: create, join...)")
	flag.StringVar(&file, "file", "file", "-file words/en.txt")
	flag.StringVar(&gameID, "game", "game", "-game XYZ (get by create command)")
	flag.StringVar(&player, "player", "player", "-player murat")
	flag.StringVar(&team, "team", "team", "-team red (available teams: red, blue)")
	flag.Parse()

	dial, err := grpc.Dial("localhost:"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not dial, %v", err)
	}
	defer func(dial *grpc.ClientConn) {
		err := dial.Close()
		if err != nil {
			log.Fatalf("could not close, %v", err)
		}
	}(dial)

	client := pb.NewCodenamesClient(dial)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch command {
	case "create":
		r, err := client.CreateGame(ctx, &pb.GameRequest{File: file})
		if err != nil {
			log.Fatalf("could not create game, %v", err)
		}

		fmt.Println(r.Game.ID)

		for _, card := range r.Game.Board.Cards {
			fmt.Printf("%s ", card.Text)
		}
	case "join":
		_, err := client.JoinGame(ctx, &pb.JoinGameRequest{
			GameID: gameID,
			Name:   player,
			Team:   0,
		})
		if err != nil {
			log.Fatalf("could not join the game, %v", err)
		}
	}
}

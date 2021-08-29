package grpc

import (
	"context"
	"net"

	"github.com/rs/zerolog/log"
	pbHighscore "github.com/rwbailey/m-apis/highscore/v1"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

var HighScore = 999999.0

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) SetHighScore(ctx context.Context, input *pbHighscore.SetHighScoreRequest) (*pbHighscore.SetHighScoreResponse, error) {
	log.Info.Msg("SetHighScore in m-highscore was called")
	HighScore = input.GetHighScore()

	return &pbHighscore.SetHighScoreResponse{
		Set: true,
	}, nil
}

func (g *Grpc) GetHighScore(ctx context.Context, input *pbHighscore.GetHighScoreRequest) (*pbHighscore.GetHighScoreResponse, error) {
	log.Info.Msg("GetHighScore in m-highscore was called")
	return &pbHighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil
}

func (g *Grpc) ListenAndServ() error {
	lst, err := net.Listen("tcp", g.address)
	if err != nil {
		return err
	}

	serverOpts := []grpc.ServerOptions{}

	g.srv = grpc.NewServer(serverOpts)

	pbHighscore.RegisterGameServer(g.srv, g)

	log.Info.Msg("Starting gRPC server for highscore at address:", g.address)

	err = g.srv.Serve(lst)
	if err != nil {
		return err
	}

	return nil
}

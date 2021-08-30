package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	pbHighscore "github.com/rwbailey/m-apis/highscore/v1"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:50051", "address to connect")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to address: " + *addressPtr)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close connection")
		}
	}()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := pbHighscore.NewGameClient(conn)
	if c == nil {
		log.Info().Msg("client nil")
	}

	resp, err := c.GetHighScore(timeoutCtx, &pbHighscore.GetHighScoreRequest{})
	if err != nil {
		log.Fatal().Err(err).Msg("error calling GetHighScore")
	}

	log.Info().Msg(fmt.Sprintf("%f", resp.GetHighScore()))
}

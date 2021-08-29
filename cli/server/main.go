package main

import (
	"flag"

	"github.com/rs/zerolog/log"
	grpcSetup "github.com/rwbailey/m-highscore/internal/server/grpc"
)

func main() {
	var addressPtr = flag.String("address", ":50051", "address to connect")
	flag.Parse()

	s := grpcSetup.NewServer(*addressPtr)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start highscore server")
	}
}

package protos

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	UnimplementedMusicServer
}

func (s *server) RequestMusic(ctx context.Context, in *MusicRequest) (*MusicResponse, error) {
	return &MusicResponse{
		Url: fmt.Sprintf("http://127.0.0.1:9001%s", in.GetUuid()),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	RegisterMusicServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

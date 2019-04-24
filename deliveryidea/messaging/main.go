package main

import (
	"context"
	"log"
	"net"

	"../pb"
	"./flags"
	"./twilio"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const ()

type server struct{}

var twilioClient *twilio.Client

func main() {
	lis, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Starting Cart service")
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func setup() {
	twilioClient = twilio.CreateClient(
		&twilio.Config{
			AccountID: flags.GetString(flags.TwilioAccountID),
			AuthToken: flags.GetString(flags.TwilioAuthToken),
			From:      flags.GetString(flags.TwilioFromNumber),
		})
}

func (s *server) Send(ctx context.Context, msg *pb.Message) (*empty.Empty, error) {
	for _, val := range msg.Recipients {
		if err := twilioClient.SendMessage(&twilio.MessageRequest{To: val, Body: msg.Message}); err != nil {
			return nil, err
		}
	}
	return new(empty.Empty), nil
}

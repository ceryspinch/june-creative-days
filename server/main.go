package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	pb "ticket_service/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	jsonDBFile = "testdata/tickets_db.json"
	port       = flag.Int("port", 50051, "The server port")
)

type ticketManagerServer struct {
	pb.UnimplementedTicketManagerServer
	purchasedTickets map[string]*pb.Ticket
}

func newTicketManagerServer() *ticketManagerServer {
	s := &ticketManagerServer{}
	s.loadTickets(jsonDBFile)
	return s
}

// BuyTicket creates a new ticket.
func (s *ticketManagerServer) BuyTicket(ctx context.Context, req *pb.BuyTicketRequest) (*pb.BuyTicketResponse, error) {
	// Create ticket with new information
	newTicket := &pb.Ticket{
		Id:                uuid.New().String(),
		Purchaser:         req.Purchaser,
		IsBringingGuest:   req.IsBringingGuest,
		HasReceivedTicket: false,
	}

	// Update the map
	s.purchasedTickets[newTicket.Id] = newTicket

	// Update database
	err := s.updateDatabase()
	if err != nil {
		return nil, err
	}

	return &pb.BuyTicketResponse{Ticket: newTicket}, nil
}

// ListTicket returns a single ticket as found by the ID provided in the request.
func (s *ticketManagerServer) ListTicket(ctx context.Context, req *pb.ListTicketRequest) (*pb.ListTicketResponse, error) {
	ticket, ok := s.purchasedTickets[req.Id]
	if !ok {
		return nil, errors.New("ticket not found")
	}

	return &pb.ListTicketResponse{
		Ticket: &pb.Ticket{
			Id:                ticket.Id,
			Purchaser:         ticket.Purchaser,
			IsBringingGuest:   ticket.IsBringingGuest,
			HasReceivedTicket: ticket.HasReceivedTicket,
		},
	}, nil
}

// ListAllTickets returns a list of all the purchased tickets in the database.
func (s *ticketManagerServer) ListAllTickets(ctx context.Context, req *pb.ListAllTicketsRequest) (*pb.ListAllTicketsResponse, error) {
	// Create a slice of tickets from the map
	var tickets []*pb.Ticket
	for _, ticket := range s.purchasedTickets {
		tickets = append(tickets, ticket)
	}

	return &pb.ListAllTicketsResponse{
		Tickets: tickets,
	}, nil
}

// UpdateTicketInformation updates the isBringingGuest and hasReceivedTicket fields on a specific ticket.
func (s *ticketManagerServer) UpdateTicketInformation(ctx context.Context, req *pb.UpdateTicketInformationRequest) (*pb.UpdateTicketInformationResponse, error) {
	// Find the ticket in the map
	ticket, ok := s.purchasedTickets[req.Id]
	if !ok {
		return nil, errors.New("ticket not found")
	}

	// Create a new ticket with the updated information.
	updatedTicket := &pb.Ticket{
		Id:                ticket.Id,
		Purchaser:         ticket.Purchaser,
		IsBringingGuest:   req.IsBringingGuest,
		HasReceivedTicket: req.HasReceivedTicket,
	}

	// Update the map with the updated ticket
	s.purchasedTickets[req.Id] = updatedTicket

	// Update database
	err := s.updateDatabase()
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTicketInformationResponse{
		Ticket: updatedTicket,
	}, nil

}

// DeleteTicket deletes a ticket with a specific ID.
func (s *ticketManagerServer) DeleteTicket(ctx context.Context, req *pb.DeleteTicketRequest) (*pb.DeleteTicketResponse, error) {
	// Find the ticket to delete
	delete(s.purchasedTickets, req.Id)

	// Update database
	err := s.updateDatabase()
	if err != nil {
		return nil, err
	}

	return &pb.DeleteTicketResponse{}, nil
}

// loadTickets loads tickets from a JSON file.
func (s *ticketManagerServer) loadTickets(filePath string) error {
	var data []byte
	var err error

	// Check if filePath is empty
	if filePath == "" {
		return errors.New("file path not specified, could not load data")
	}

	data, err = os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to load tickets from file: %v", err)
	}

	if err := json.Unmarshal(data, &s.purchasedTickets); err != nil {
		return fmt.Errorf("failed to unmarshal tickets: %v", err)
	}
	return nil
}

func (s *ticketManagerServer) updateDatabase() error {
	updatedDatabase, err := json.MarshalIndent(s.purchasedTickets, "", "    ")
	if err != nil {
		return err
	}
	ioutil.WriteFile(jsonDBFile, updatedDatabase, 0644)
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterTicketManagerServer(grpcServer, newTicketManagerServer())

	grpcServer.Serve(lis)
}

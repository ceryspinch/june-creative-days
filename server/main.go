package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "ticket_service/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
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

// Create
func (s *ticketManagerServer) BuyTicket(ctx context.Context, req *pb.BuyTicketRequest) (*pb.BuyTicketResponse, error) {
	newTicket := &pb.Ticket{
		Id:                uuid.New().String(),
		Purchaser:         req.Purchaser,
		IsBringingGuest:   req.IsBringingGuest,
		HasReceivedTicket: false,
	}

	s.purchasedTickets[newTicket.Id] = newTicket

	return &pb.BuyTicketResponse{Ticket: newTicket}, nil
}

// ListTicket returns a single ticket as found by the ID provided in the request.
func (s *ticketManagerServer) ListTicket(ctx context.Context, req *pb.ListTicketRequest) (*pb.ListTicketResponse, error) {
	ticket, ok := s.purchasedTickets[req.Id]
	if !ok {
		return nil, errors.New("Ticket not found")
	}
	// TODO: Get by name/email instead of id?
	return &pb.ListTicketResponse{
		Ticket: &pb.Ticket{
			Id:                ticket.Id,
			Purchaser:         ticket.Purchaser,
			IsBringingGuest:   ticket.IsBringingGuest,
			HasReceivedTicket: ticket.HasReceivedTicket,
		},
	}, nil
}

// ListAllTickets returns a list of all the purchased tickets in the database
func (s *ticketManagerServer) ListAllTickets(ctx context.Context, req *pb.ListAllTicketsRequest) (*pb.ListAllTicketsResponse, error) {

	var tickets []*pb.Ticket
	for _, ticket := range s.purchasedTickets {
		tickets = append(tickets, ticket)
	}

	return &pb.ListAllTicketsResponse{
		Tickets: tickets,
	}, nil
}

func (s *ticketManagerServer) UpdateTicketInformation(ctx context.Context, req *pb.UpdateTicketInformationRequest) (*pb.UpdateTicketInformationResponse, error) {
	ticket, ok := s.purchasedTickets[req.Id]
	if !ok {
		return nil, errors.New("Ticket not found")
	}

	updatedTicket := &pb.Ticket{
		Id:                ticket.Id,
		Purchaser:         ticket.Purchaser,
		IsBringingGuest:   req.IsBringingGuest,
		HasReceivedTicket: req.HasReceivedTicket,
	}

	s.purchasedTickets[req.Id] = updatedTicket

	// TODO: Update db with new ticket
	return &pb.UpdateTicketInformationResponse{
		Ticket: updatedTicket,
	}, nil

}

func (s *ticketManagerServer) DeleteTicket(ctx context.Context, req *pb.DeleteTicketRequest) (*pb.DeleteTicketResponse, error) {
	delete(s.purchasedTickets, req.Id)

	return &pb.DeleteTicketResponse{}, nil
}

func (s *ticketManagerServer) findTicket(id string) (*pb.Ticket, error) {
	for _, ticket := range s.purchasedTickets {
		if ticket.Id == id {
			return ticket, nil
		}
	}
	return nil, fmt.Errorf("Could not find ticket %s", id)
}

// loadTickets loads tickets from a JSON file.
func (s *ticketManagerServer) loadTickets(filePath string) error {
	var data []byte
	var err error

	// Check if filePath is empty
	if filePath == "" {
		return errors.New("File path not specified, could not load data")
	}

	data, err = os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Failed to load tickets from file: %v", err)
	}

	if err := json.Unmarshal(data, &s.purchasedTickets); err != nil {
		return fmt.Errorf("Failed to unmarshal tickets: %v", err)
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer() //opts...)
	pb.RegisterTicketManagerServer(grpcServer, newTicketManagerServer())
	grpcServer.Serve(lis)
}

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "ticket_service/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	DatabaseConnection()
}

var DB *gorm.DB
var err error

type Ticket struct {
	ID                string `gorm:"primarykey"`
	Purchaser         string
	IsBringingGuest   bool
	HasReceivedTicket bool
	CreatedAt         time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime:false"`
}

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "postgres"
	dbUser := "postgres"
	password := "temp"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Ticket{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

type server struct {
	pb.UnimplementedTicketManagerServer
}

// Create
func (*server) BuyTicket(ctx context.Context, req *pb.BuyTicketRequest) (*pb.BuyTicketResponse, error) {
	fmt.Println("Buy Ticket")

	// Create a new ID for the ticket
	newTicketID := uuid.New().String()

	// Create the ticket data
	data := Ticket{
		ID:                newTicketID,
		Purchaser:         req.Purchaser,
		IsBringingGuest:   req.IsBringingGuest,
		HasReceivedTicket: false,
	}

	// Add new ticket to the database
	res := DB.Create(&data)
	if res.RowsAffected == 0 {
		return nil, errors.New("ticket creation unsuccessful")
	}

	return &pb.BuyTicketResponse{
		Ticket: &pb.Ticket{
			Id:                newTicketID,
			Purchaser:         req.Purchaser,
			IsBringingGuest:   req.IsBringingGuest,
			HasReceivedTicket: false,
		},
	}, nil
}

// Get
func (*server) ListTicket(ctx context.Context, req *pb.ListTicketRequest) (*pb.ListTicketResponse, error) {
	fmt.Println("List Ticket", req.GetId())

	var ticket Ticket

	// Find the ticket with the same ID as the request in the database
	res := DB.Find(&ticket, "id = ?", req.GetId())
	if res.RowsAffected == 0 {
		return nil, errors.New("ticket not found")
	}

	return &pb.ListTicketResponse{
		Ticket: &pb.Ticket{
			Id:                ticket.ID,
			Purchaser:         ticket.Purchaser,
			IsBringingGuest:   ticket.IsBringingGuest,
			HasReceivedTicket: ticket.HasReceivedTicket,
		},
	}, nil
}

// Get
func (*server) ListAllTickets(ctx context.Context, req *pb.ListAllTicketsRequest) (*pb.ListAllTicketsResponse, error) {
	fmt.Println("List All Tickets")

	tickets := []*pb.Ticket{}

	// Find all tickets in the database
	res := DB.Find(&tickets)
	if res.RowsAffected == 0 {
		return nil, errors.New("ticket not found")
	}

	return &pb.ListAllTicketsResponse{
		Tickets: tickets,
	}, nil
}

func (*server) UpdateTicketInformation(ctx context.Context, req *pb.UpdateTicketInformationRequest) (*pb.UpdateTicketInformationResponse, error) {
	fmt.Println("Update Ticket")
	var ticket Ticket

	res := DB.Model(&ticket).Where("id=?", req.Id).Updates(Ticket{ID: req.Id, Purchaser: ticket.Purchaser, IsBringingGuest: req.IsBringingGuest, HasReceivedTicket: req.HasReceivedTicket})
	if res.RowsAffected == 0 {
		return nil, errors.New("ticket not found")
	}

	return &pb.UpdateTicketInformationResponse{
		Ticket: &pb.Ticket{
			Id:                ticket.ID,
			Purchaser:         ticket.Purchaser,
			IsBringingGuest:   ticket.IsBringingGuest,
			HasReceivedTicket: ticket.HasReceivedTicket,
		},
	}, nil
}

func (*server) DeleteTicket(ctx context.Context, req *pb.DeleteTicketRequest) (*pb.DeleteTicketResponse, error) {
	fmt.Println("Delete Ticket")

	var ticket Ticket
	res := DB.Where("id = ?", req.GetId()).Delete(&ticket)
	if res.RowsAffected == 0 {
		return nil, errors.New("ticket not found")
	}

	return &pb.DeleteTicketResponse{}, nil
}

func main() {
	fmt.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTicketManagerServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}

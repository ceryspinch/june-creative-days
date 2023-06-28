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

	"ticket_service/data"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// func init() {
// 	DatabaseConnection()
// }

// var DB *gorm.DB
// var err error

// type Ticket struct {
// 	ID                string `gorm:"primarykey"`
// 	Purchaser         string
// 	IsBringingGuest   bool
// 	HasReceivedTicket bool
// 	CreatedAt         time.Time `gorm:"autoCreateTime:false"`
// 	UpdatedAt         time.Time `gorm:"autoUpdateTime:false"`
// }

// func DatabaseConnection() {
// 	host := "localhost"
// 	port := "5432"
// 	dbName := "postgres"
// 	dbUser := "postgres"
// 	password := "temp"
// 	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
// 		host,
// 		port,
// 		dbUser,
// 		dbName,
// 		password,
// 	)
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	DB.AutoMigrate(Ticket{})
// 	if err != nil {
// 		log.Fatal("Error connecting to the database...", err)
// 	}
// 	fmt.Println("Database connection successful...")
// }

// var (
// 	port = flag.Int("port", 50051, "gRPC server port")
// )

type ticketManagerServer struct {
	pb.UnimplementedTicketManagerServer
	purchasedTickets []*pb.Ticket
}

func newTicketManagerServer() *ticketManagerServer {
	s := &ticketManagerServer{}
	s.loadTickets(*jsonDBFile)
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

	s.purchasedTickets = append(s.purchasedTickets, newTicket)

	return &pb.BuyTicketResponse{Ticket: newTicket}, nil

	// fmt.Println("Buy Ticket")

	// // Create a new ID for the ticket
	// newTicketID := uuid.New().String()

	// // Create the ticket data
	// data := Ticket{
	// 	ID:                newTicketID,
	// 	Purchaser:         req.Purchaser,
	// 	IsBringingGuest:   req.IsBringingGuest,
	// 	HasReceivedTicket: false,
	// }

	// // Add new ticket to the database
	// res := DB.Create(&data)
	// if res.RowsAffected == 0 {
	// 	return nil, errors.New("ticket creation unsuccessful")
	// }

	// return &pb.BuyTicketResponse{
	// 	Ticket: &pb.Ticket{
	// 		Id:                newTicketID,
	// 		Purchaser:         req.Purchaser,
	// 		IsBringingGuest:   req.IsBringingGuest,
	// 		HasReceivedTicket: false,
	// 	},
	// }, nil
}

// ListTicket returns a single ticket as found by the ID provided in the request.
func (s *ticketManagerServer) ListTicket(ctx context.Context, req *pb.ListTicketRequest) (*pb.ListTicketResponse, error) {
	ticket, err := s.findTicket(req.Id)
	if err != nil {
		return nil, err
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

// fmt.Println("List Ticket", req.GetId())

// var ticket Ticket

// // Find the ticket with the same ID as the request in the database
// res := DB.Find(&ticket, "id = ?", req.GetId())
// if res.RowsAffected == 0 {
// 	return nil, errors.New("ticket not found")
// }

// return &pb.ListTicketResponse{
// 	Ticket: &pb.Ticket{
// 		Id:                ticket.ID,
// 		Purchaser:         ticket.Purchaser,
// 		IsBringingGuest:   ticket.IsBringingGuest,
// 		HasReceivedTicket: ticket.HasReceivedTicket,
// 	},
// }, nil

// ListAllTickets returns a list of all the purchased tickets in the database
func (s *ticketManagerServer) ListAllTickets(ctx context.Context, req *pb.ListAllTicketsRequest) (*pb.ListAllTicketsResponse, error) {
	// fmt.Println("List All Tickets")

	// tickets := []*pb.Ticket{}

	// // Find all tickets in the database
	// res := DB.Find(&tickets)
	// if res.RowsAffected == 0 {
	// 	return nil, errors.New("ticket not found")
	// }

	return &pb.ListAllTicketsResponse{
		Tickets: s.purchasedTickets,
	}, nil
}

func (s *ticketManagerServer) UpdateTicketInformation(ctx context.Context, req *pb.UpdateTicketInformationRequest) (*pb.UpdateTicketInformationResponse, error) {
	ticket, err := s.findTicket(req.Id)
	if err != nil {
		return nil, err
	}

	updatedTicket := &pb.Ticket{
		Id:                ticket.Id,
		Purchaser:         ticket.Purchaser,
		IsBringingGuest:   req.IsBringingGuest,
		HasReceivedTicket: req.HasReceivedTicket,
	}

	// TODO: Update db with new ticket
	return &pb.UpdateTicketInformationResponse{
		Ticket: updatedTicket,
	}, nil

	// fmt.Println("Update Ticket")
	// var ticket Ticket

	// res := DB.Model(&ticket).Where("id=?", req.Id).Updates(Ticket{ID: req.Id, Purchaser: ticket.Purchaser, IsBringingGuest: req.IsBringingGuest, HasReceivedTicket: req.HasReceivedTicket})
	// if res.RowsAffected == 0 {
	// 	return nil, errors.New("ticket not found")
	// }

	// return &pb.UpdateTicketInformationResponse{
	// 	Ticket: &pb.Ticket{
	// 		Id:                ticket.ID,
	// 		Purchaser:         ticket.Purchaser,
	// 		IsBringingGuest:   ticket.IsBringingGuest,
	// 		HasReceivedTicket: ticket.HasReceivedTicket,
	// 	},
	// }, nil
}

func (s *ticketManagerServer) DeleteTicket(ctx context.Context, req *pb.DeleteTicketRequest) (*pb.DeleteTicketResponse, error) {
	for i, ticket := range s.purchasedTickets {
		if ticket.Id == req.Id {
			s.purchasedTickets = append(s.purchasedTickets[:i], s.purchasedTickets[i+1])
		}
	}

	// fmt.Println("Delete Ticket")

	// var ticket Ticket
	// res := DB.Where("id = ?", req.GetId()).Delete(&ticket)
	// if res.RowsAffected == 0 {
	// 	return nil, errors.New("ticket not found")
	// }

	// return &pb.DeleteTicketResponse{}, nil
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

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = data.Path("x509/server_cert.pem")
		}
		if *keyFile == "" {
			*keyFile = data.Path("x509/server_key.pem")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTicketManagerServer(grpcServer, newTicketManagerServer())
	grpcServer.Serve(lis)
}

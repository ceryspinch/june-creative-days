package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "ticket_service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

type Ticket struct {
	ID                string `json:"id"`
	Purchaser         string `json:"Purchaser"`
	IsBringingGuest   bool   `json:"IsBringingGuest"`
	HasReceivedTicket bool   `json:"HasReceivedTicket"`
}

func buyTicket(ctx context.Context, client pb.TicketManagerClient, purchaser string, isBringingGuest bool) error {
	log.Printf("Purchasing ticket for %s", purchaser)
	ticket, err := client.BuyTicket(ctx, &pb.BuyTicketRequest{Purchaser: purchaser, IsBringingGuest: isBringingGuest})
	if err != nil {
		return err
	}
	log.Println(ticket)
	return nil
}

func listTicket(ctx context.Context, client pb.TicketManagerClient, id string) error {
	log.Printf("Getting ticket information for ticket %s", id)

	tickets, err := client.ListTicket(ctx, &pb.ListTicketRequest{Id: id})
	if err != nil {
		return err
	}
	log.Println(tickets)
	return nil
}

func listAllTickets(ctx context.Context, client pb.TicketManagerClient) (*pb.ListAllTicketsResponse, error) {
	log.Println("Getting all ticket information")

	tickets, err := client.ListAllTickets(ctx, &pb.ListAllTicketsRequest{})
	if err != nil {
		//pb.ListAllTicketsResponse{}, err
	}
	//log.Println(tickets)
	return tickets, nil
}

func updateTicketInformation(ctx context.Context, client pb.TicketManagerClient, id string, isBringingGuest, hasReceivedTicket bool) error {
	log.Printf("Updating ticket information for ticket %s", id)
	ticket, err := client.UpdateTicketInformation(ctx, &pb.UpdateTicketInformationRequest{Id: id, IsBringingGuest: isBringingGuest, HasReceivedTicket: hasReceivedTicket})
	if err != nil {
		return err
	}
	log.Println(ticket)
	return nil
}

func deleteTicket(ctx context.Context, client pb.TicketManagerClient, id string) error {
	log.Printf("Deleting ticket %s", id)
	_, err := client.DeleteTicket(ctx, &pb.DeleteTicketRequest{Id: id})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	// Use insecure credentials for now as only running on bench
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create client
	client := pb.NewTicketManagerClient(conn)

	fmt.Println()
	buyTicket(ctx, client, "Cerys Pinch", false)
	fmt.Println()
	ticketsList, _ := listAllTickets(ctx, client)
	cerysId := ticketsList.Tickets[0].Id
	fmt.Println(ticketsList)
	fmt.Println()
	listTicket(ctx, client, cerysId)
	fmt.Println()
	updateTicketInformation(ctx, client, cerysId, false, true)
	fmt.Println()
	deleteTicket(ctx, client, cerysId)

}

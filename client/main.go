package main

import (
	"context"
	"flag"
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

func listTicket(ctx context.Context, client pb.TicketManagerClient, id string) {
	log.Printf("Getting ticket information for ticket %d, id")

	tickets, err := client.ListTicket(ctx, &pb.ListTicketRequest{Id: id})
	if err != nil {
		log.Fatalf("client.ListTicket failed: %v", err)
	}
	log.Println(tickets)
}

func listAllTickets(ctx context.Context, client pb.TicketManagerClient) {
	log.Println("Getting all ticket information")

	tickets, err := client.ListAllTickets(ctx, &pb.ListAllTicketsRequest{})
	if err != nil {
		log.Fatalf("client.ListAllTickets failed: %v", err)
	}
	log.Println(tickets)
}

func buyTicket(ctx context.Context, client pb.TicketManagerClient, purchaser string, isBringingGuest bool) {
	log.Printf("Purchasing ticket for %s, purchaser")
	ticket, err := client.BuyTicket(ctx, &pb.BuyTicketRequest{Purchaser: purchaser, IsBringingGuest: isBringingGuest})
	if err != nil {
		log.Fatalf("client.updateTicketInformation failed: %v", err)
	}
	log.Println(ticket)
}

func updateTicketInformation(ctx context.Context, client pb.TicketManagerClient, id string, isBringingGuest, hasReceivedTicket bool) {
	log.Printf("Updating ticket information for ticket %d, id")
	ticket, err := client.UpdateTicketInformation(ctx, &pb.UpdateTicketInformationRequest{Id: id, IsBringingGuest: isBringingGuest, HasReceivedTicket: hasReceivedTicket})
	if err != nil {
		log.Fatalf("client.updateTicketInformation failed: %v", err)
	}
	log.Println(ticket)
}

func deleteTicket(ctx context.Context, client pb.TicketManagerClient, id string) {
	log.Printf("Deleting ticket %d, id")
	client.DeleteTicket(ctx, &pb.DeleteTicketRequest{Id: id})
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

	buyTicket(ctx, client, "Cerys Pinch", false)
	listTicket(ctx, client, "TODO")
	updateTicketInformation(ctx, client, "TODO", false, true)
	deleteTicket(ctx, client, "TODO")
	listAllTickets(ctx, client)
}

// func main() {
// 	flag.Parse()
// 	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}

// 	defer conn.Close()
// 	client := pb.NewTicketManagerClient(conn)

// 	//generic request
// 	r := gin.Default()

// 	//request for all tickets
// 	r.GET("/tickets", func(ctx *gin.Context) {
// 		res, err := client.ListAllTickets(ctx, &pb.ListAllTicketsRequest{})
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err,
// 			})
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"tickets": res.Tickets,
// 		})
// 	})

// 	//request for singular ticket
// 	r.GET("/tickets/:id", func(ctx *gin.Context) {
// 		id := ctx.Param("id")
// 		res, err := client.ListTicket(ctx, &pb.ListTicketRequest{Id: id})
// 		if err != nil {
// 			ctx.JSON(http.StatusNotFound, gin.H{
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"ticket": res.Ticket,
// 		})
// 	})

// 	//buy ticket
// 	r.POST("/tickets", func(ctx *gin.Context) {
// 		var ticket Ticket

// 		err := ctx.ShouldBind(&ticket)
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err,
// 			})
// 			return
// 		}
// 		res, err := client.BuyTicket(ctx, &pb.BuyTicketRequest{Purchaser: ticket.Purchaser, IsBringingGuest: ticket.IsBringingGuest})
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err,
// 			})
// 			return
// 		}
// 		ctx.JSON(http.StatusCreated, gin.H{
// 			"Ticket": res.Ticket,
// 		})
// 	})

// 	//update ticket information
// 	r.PUT("/tickets/:id", func(ctx *gin.Context) {
// 		var ticket Ticket
// 		err := ctx.ShouldBind(&ticket)
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		res, err := client.UpdateTicketInformation(ctx, &pb.UpdateTicketInformationRequest{Id: ticket.ID, IsBringingGuest: ticket.IsBringingGuest, HasReceivedTicket: ticket.HasReceivedTicket})
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"ticket": res.Ticket,
// 		})
// 		return
// 	})

// 	//delete ticket
// 	r.DELETE("/ticket/:id", func(ctx *gin.Context) {
// 		id := ctx.Param("id")
// 		_, err := client.DeleteTicket(ctx, &pb.DeleteTicketRequest{Id: id})
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})

// 	})
// 	r.Run(":5000")
// }

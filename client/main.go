package main

import (
	"flag"
	"log"
	"net/http"

	pb "ticket_service/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type Ticket struct {
	ID                string `json:"id"`
	Purchaser         string `json:"Purchaser"`
	IsBringingGuest   bool   `json:"IsBringingGuest"`
	HasReceivedTicket bool   `json:"HasReceivedTicket"`
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewTicketManagerClient(conn)

	//generic request
	r := gin.Default()

	//request for all tickets
	r.GET("/tickets", func(ctx *gin.Context) {
		res, err := client.ListAllTickets(ctx, &pb.ListAllTicketsRequest{})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"tickets": res.Tickets,
		})
	})

	//request for singular ticket
	r.GET("/tickets/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		res, err := client.ListTicket(ctx, &pb.ListTicketRequest{Id: id})
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"ticket": res.Ticket,
		})
	})

	//buy ticket
	r.POST("/tickets", func(ctx *gin.Context) {
		var ticket Ticket

		err := ctx.ShouldBind(&ticket)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		res, err := client.BuyTicket(ctx, &pb.BuyTicketRequest{Purchaser: ticket.Purchaser, IsBringingGuest: ticket.IsBringingGuest})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"Ticket": res.Ticket,
		})
	})

	//update ticket information
	r.PUT("/tickets/:id", func(ctx *gin.Context) {
		var ticket Ticket
		err := ctx.ShouldBind(&ticket)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		res, err := client.UpdateTicketInformation(ctx, &pb.UpdateTicketInformationRequest{ID: ticket.ID, IsBringingGuest: ticket.IsBringingGuest, HasReceivedTicket: ticket.HasReceivedTicket})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"ticket": res.Ticket,
		})
		return
	})

	//delete ticket
	r.DELETE("/ticket/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		_, err := client.DeleteTicket(ctx, &pb.DeleteTicketRequest{Id: id})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})

	})
	r.Run(":5000")
}

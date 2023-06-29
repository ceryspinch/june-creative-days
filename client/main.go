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

func buyTicket(ctx context.Context, client pb.TicketManagerClient, purchaser string, isBringingGuest bool) (*pb.BuyTicketResponse, error) {
	//log.Printf("Purchasing ticket for %s", purchaser)

	ticket, err := client.BuyTicket(ctx, &pb.BuyTicketRequest{Purchaser: purchaser, IsBringingGuest: isBringingGuest})
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func listTicket(ctx context.Context, client pb.TicketManagerClient, id string) (*pb.ListTicketResponse, error) {
	//log.Printf("Getting ticket information for ticket %s", id)

	ticket, err := client.ListTicket(ctx, &pb.ListTicketRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func listAllTickets(ctx context.Context, client pb.TicketManagerClient) (*pb.ListAllTicketsResponse, error) {
	//log.Println("Getting all ticket information")

	tickets, err := client.ListAllTickets(ctx, &pb.ListAllTicketsRequest{})
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func updateTicketInformation(ctx context.Context, client pb.TicketManagerClient, id string, isBringingGuest, hasReceivedTicket bool) (*pb.UpdateTicketInformationResponse, error) {
	//log.Printf("Updating ticket information for ticket %s", id)
	updatedTicket, err := client.UpdateTicketInformation(ctx, &pb.UpdateTicketInformationRequest{Id: id, IsBringingGuest: isBringingGuest, HasReceivedTicket: hasReceivedTicket})
	if err != nil {
		return nil, err
	}
	return updatedTicket, nil
}

func deleteTicket(ctx context.Context, client pb.TicketManagerClient, id string) (*pb.DeleteTicketResponse, error) {
	//log.Printf("Deleting ticket %s", id)
	resp, err := client.DeleteTicket(ctx, &pb.DeleteTicketRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func formatting(ticket *pb.Ticket) {
	fmt.Println("Printing ticket info")
	fmt.Println("Name: ", ticket.Purchaser)
	fmt.Println("Id: ", ticket.Id)
	if ticket.IsBringingGuest {
		fmt.Println("Plus one? Yes")
	}
	if ticket.HasReceivedTicket {
		fmt.Println("Has received ticket? Yes")
	} else {
		fmt.Println("Has received ticket? No")
	}
	fmt.Println()

}

func testing(ctx context.Context, client pb.TicketManagerClient) {
	// Testing purposes
	ticketsList, err := listAllTickets(ctx, client)
	if err != nil {
		fmt.Println(err)
	}
	for _, ticket := range ticketsList.Tickets {
		formatting(ticket)
	}
	//fmt.Println(ticketsList)

	// ticket, err := listTicket(ctx, client, "124")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// formatting(ticket.Ticket)
	// // fmt.Println(ticket)

	// newTicket, err := buyTicket(ctx, client, "Mark", false)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// ticketsList, err = listAllTickets(ctx, client)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // fmt.Println(ticketsList)

	// updatedTicket, err := updateTicketInformation(ctx, client, newTicket.Ticket.Id, true, true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// //formatting(updatedTicket.Ticket)
	// // fmt.Println(updatedTicket)

	// ticket, err = listTicket(ctx, client, newTicket.Ticket.Id)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // fmt.Println(ticket)

	// _, err = deleteTicket(ctx, client, newTicket.Ticket.Id)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// ticketsList, err = listAllTickets(ctx, client)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // fmt.Println(ticketsList)
}

func menu(ctx context.Context, client pb.TicketManagerClient) {
	fmt.Println("****************************")
	fmt.Println("      Ticketing Client")
	fmt.Println("****************************\n\n")

	fmt.Println("            MENU")
	fmt.Println("Enter for the following options")
	fmt.Println("1. To view all tickets")
	fmt.Println("2. To buy a tickets")
	fmt.Println("3. To view a specific ticket")
	fmt.Println("4. To update a ticket")
	fmt.Println("5. To delete a ticket")
	fmt.Println("6. To quit")
	fmt.Print("Enter your choice: ")
	var option int
	fmt.Scanln(&option)
	fmt.Println()

	for option != 6 {
		switch option {
		case 1:
			ticketsList, err := listAllTickets(ctx, client)
			if err != nil {
				fmt.Println(err)
			}
			for _, ticket := range ticketsList.Tickets {
				formatting(ticket)
			}
		case 2:
			var name string
			var isBringingGuest bool
			fmt.Print("Enter the name of the purchaser: ")
			fmt.Scanln(&name)
			fmt.Print("Is the purchaser bringing a +1? Enter true or false: ")
			fmt.Scanln(&isBringingGuest)
			newTicket, err := buyTicket(ctx, client, name, isBringingGuest)
			if err != nil {
				fmt.Println(err)
			}
			formatting(newTicket.Ticket)
		case 3:
			var id string
			fmt.Print("Enter the id of the ticket to list: ")
			fmt.Scanln(&id)
			ticket, err := listTicket(ctx, client, id)
			if err != nil {
				fmt.Println(err)
			}
			formatting(ticket.Ticket)
		case 4:
			var id string
			var isBringingGuest bool
			var hasReceivedTicket bool
			fmt.Print("Enter the id of the ticket to update: ")
			fmt.Scanln(&id)
			fmt.Print("Is the purchaser bringing a +1? Enter true or false: ")
			fmt.Scanln(&isBringingGuest)
			fmt.Print("Has the purchaser recieved a ticket? Enter the id of the ticket to update: ")
			fmt.Scanln(&hasReceivedTicket)
			updatedTicket, err := updateTicketInformation(ctx, client, id, isBringingGuest, hasReceivedTicket)
			if err != nil {
				fmt.Println(err)
			}
			formatting(updatedTicket.Ticket)
		case 5:
			var id string
			fmt.Print("Enter the id of the ticket to delete: ")
			fmt.Scanln(&id)
			_, err := deleteTicket(ctx, client, id)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Ticket deleted")
		default:
			fmt.Println("That was not a valid option")
		}
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&option)
		fmt.Println()
	}
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

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	// Create client
	client := pb.NewTicketManagerClient(conn)

	menu(ctx, client)
	//testing(ctx, client)

}

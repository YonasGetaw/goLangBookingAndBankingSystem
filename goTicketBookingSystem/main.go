package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Event struct {
	ID    int
	Name  string
	Seats int
}

type Booking struct {
	Name      string
	EventName string
	Tickets   int
}

var events = []Event{
	{ID: 1, Name: "Music Concert", Seats: 100},
	{ID: 2, Name: "Tech Conference", Seats: 50},
	{ID: 3, Name: "Art Festival", Seats: 30},
}

var bookings []Booking

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n🎟️  Welcome to Go Ticket Booking System")
		fmt.Println("1. Show Events")
		fmt.Println("2. Book Ticket")
		fmt.Println("3. Show All Bookings")
		fmt.Println("4. Exit")
		fmt.Print("👉 Enter your choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)

		if err != nil {
			fmt.Println("❌ Invalid input, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			showEvents()
		case 2:
			bookTicket(reader)
		case 3:
			showBookings()
		case 4:
			fmt.Println("👋 Thank you! Exiting the system.")
			return
		default:
			fmt.Println("❌ Invalid option. Try again.")
		}
	}
}

func showEvents() {
	fmt.Println("\n📅 Available Events:")
	for _, event := range events {
		fmt.Printf("ID: %d | Event: %s | Available Seats: %d\n", event.ID, event.Name, event.Seats)
	}
}

func bookTicket(reader *bufio.Reader) {
	showEvents()

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter Event ID to book: ")
	eventIDStr, _ := reader.ReadString('\n')
	eventIDStr = strings.TrimSpace(eventIDStr)
	eventID, err := strconv.Atoi(eventIDStr)

	if err != nil || eventID < 1 || eventID > len(events) {
		fmt.Println("❌ Invalid Event ID.")
		return
	}

	selectedEvent := &events[eventID-1]

	fmt.Printf("Enter number of tickets for %s: ", selectedEvent.Name)
	ticketStr, _ := reader.ReadString('\n')
	ticketStr = strings.TrimSpace(ticketStr)
	tickets, err := strconv.Atoi(ticketStr)

	if err != nil || tickets < 1 {
		fmt.Println("❌ Invalid number of tickets.")
		return
	}

	if selectedEvent.Seats < tickets {
		fmt.Printf("❌ Only %d seats available for %s.\n", selectedEvent.Seats, selectedEvent.Name)
		return
	}

	selectedEvent.Seats -= tickets
	bookings = append(bookings, Booking{
		Name:      name,
		EventName: selectedEvent.Name,
		Tickets:   tickets,
	})

	fmt.Printf("✅ Booking confirmed for %s (%d ticket(s)) to %s!\n", name, tickets, selectedEvent.Name)
}

func showBookings() {
	if len(bookings) == 0 {
		fmt.Println("📭 No bookings yet.")
		return
	}

	fmt.Println("\n📖 All Bookings:")
	for i, b := range bookings {
		fmt.Printf("%d. %s booked %d ticket(s) for %s\n", i+1, b.Name, b.Tickets, b.EventName)
	}
}

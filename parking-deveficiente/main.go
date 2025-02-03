package main

import (
	"log"
	"time"
)

func main() {
	ticketer := Ticketer{}

	carType := 1
	ticket, err := ticketer.IssueTicket(carType, time.Now())

	if err != nil {
		log.Printf("error issuing ticket: %s", err)
	}

	log.Print(ticket)
	// send ticket to the driver
}

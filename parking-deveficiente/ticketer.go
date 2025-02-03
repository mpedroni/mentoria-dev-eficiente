package main

import (
	"errors"
	"time"
)

type Ticketer struct{}

func (t *Ticketer) IssueTicket(vehicleType int, time time.Time) (string, error) {
	if !t.HasAvailableSpots(vehicleType) {
		return "", errors.New("no spots available")
	}

	return "ticket with id {id} issued at {time}", nil
}

func (t *Ticketer) HasAvailableSpots(vehicleType int) bool {
	// check available spots for the vehicle type
	// create a stub tot he
	return false
}

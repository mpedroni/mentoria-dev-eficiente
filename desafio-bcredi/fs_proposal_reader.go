package main

import (
	"fmt"
	"strings"
)

type ProposalReader interface {
	Read() ([]Proposal, error)
}

type fsProposalReader struct {
	f string
}

func NewFileSystemProposalReader(f string) ProposalReader {
	return &fsProposalReader{f}
}

func (r *fsProposalReader) Read() ([]Proposal, error) {
	var proposalID string
	w := make([]Warranty, 0)

	for _, line := range strings.Split(r.f, "\n") {
		fields := strings.Split(line, ",")

		eventType := fields[1] + "," + fields[2]

		switch eventType {
		case "proposal,created":
			proposalID = fields[4]
			w = make([]Warranty, 0)
		case "warranty,added":
			id := fields[5]
			w = append(w, NewWarranty(id, proposalID))
		case "proponent,added":
		default:
			return nil, fmt.Errorf("unknown event type: %s", eventType)
		}
	}

	return []Proposal{
		NewProposal(proposalID, w),
	}, nil
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileField int

const (
	EventID FileField = iota
	EventType
	Action
	Timestamp
	ProposalID

	RequiredValue    = 5
	DeadlineInMonths = 6

	WarrantyID       = 5
	WarrantyPrice    = 6
	WarrantyProvince = 7

	ProponentID     = 5
	ProponentName   = 6
	ProponentAge    = 7
	ProponentIncome = 8
	ProponentIsMain = 9
)

type ProposalReader interface {
	Read() ([]ProposalReaderResponse, error)
}

type fsProposalReader struct {
	content string
}

type ProposalReaderResponse struct {
	Proposal
	Err error
}

type proposalData struct {
	proposalID       string
	requiredValue    float64
	deadlineInMonths int
	warranties       []Warranty
	proponents       []Proponent
}

func NewFileSystemProposalReader(filename string) (ProposalReader, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return &fsProposalReader{}, err
	}

	return &fsProposalReader{content: string(content)}, nil
}

func (r *fsProposalReader) Read() ([]ProposalReaderResponse, error) {
	proposals := make([]ProposalReaderResponse, 0)
	pdata := make([]*proposalData, 0)

	for n, line := range strings.Split(r.content, "\n") {
		fields := strings.Split(line, ",")

		eventType := fields[EventType] + "," + fields[Action]

		switch eventType {
		case "proposal,created":
			pd := &proposalData{}
			pd.proposalID = fields[ProposalID]

			valueAsString := fields[RequiredValue]
			value, err := strconv.ParseFloat(valueAsString, 64)
			if err != nil {
				return nil, err
			}
			pd.requiredValue = value

			deadlineAsString := fields[DeadlineInMonths]
			deadline, err := strconv.Atoi(deadlineAsString)
			if err != nil {
				return nil, err
			}
			pd.deadlineInMonths = deadline

			pdata = append(pdata, pd)
		case "warranty,added":
			if len(pdata) == 0 {
				return nil, fmt.Errorf("warranty added before proposal created at line: %v", n+1)
			}
			id := fields[WarrantyID]
			province := fields[WarrantyProvince]
			priceAsString := fields[WarrantyPrice]

			price, err := strconv.ParseFloat(priceAsString, 64)
			if err != nil {
				return nil, err
			}

			pd := pdata[len(pdata)-1]
			pd.warranties = append(pd.warranties, NewWarranty(id, pd.proposalID, price, province))
		case "proponent,added":
			if len(pdata) == 0 {
				return nil, fmt.Errorf("proponent added before proposal created at line: %v", n+1)
			}

			id := fields[ProponentID]
			name := fields[ProponentName]
			ageAsString := fields[ProponentAge]
			age, err := strconv.Atoi(ageAsString)
			if err != nil {
				return nil, err
			}

			incomeAsString := fields[ProponentIncome]
			income, err := strconv.ParseFloat(incomeAsString, 64)
			if err != nil {
				return nil, err
			}

			isMainAsString := fields[ProponentIsMain]
			isMain, err := strconv.ParseBool(isMainAsString)
			if err != nil {
				return nil, err
			}

			pd := pdata[len(pdata)-1]
			pd.proponents = append(pd.proponents, NewProponent(id, pd.proposalID, name, age, income, isMain))
		default:
			return nil, fmt.Errorf("unknown event type: %s", eventType)
		}
	}

	for _, pd := range pdata {
		var res ProposalReaderResponse
		res.Proposal, res.Err = NewProposal(pd.proposalID, pd.requiredValue, pd.deadlineInMonths, pd.warranties, pd.proponents)
		proposals = append(proposals, res)
	}

	return proposals, nil
}

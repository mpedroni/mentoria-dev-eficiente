package proposalreader

import (
	"desafio_bcredi/internal/loan"
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

type fsProposalReader struct {
	content string
}

func NewFileSystemProposalReader(filename string) (ProposalReader, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return &fsProposalReader{}, err
	}

	return &fsProposalReader{content: string(content)}, nil
}

func (r *fsProposalReader) Read() ([]loan.Proposal, error) {
	proposals := make([]*loan.Proposal, 0)

	for n, line := range strings.Split(r.content, "\n") {
		fields := strings.Split(line, ",")

		eventType := fields[EventType] + "," + fields[Action]

		switch eventType {
		case "proposal,created":
			proposalID := fields[ProposalID]

			valueAsString := fields[RequiredValue]
			value, err := strconv.ParseFloat(valueAsString, 64)
			if err != nil {
				return nil, err
			}
			requiredValue := value

			deadlineAsString := fields[DeadlineInMonths]
			deadline, err := strconv.Atoi(deadlineAsString)
			if err != nil {
				return nil, err
			}
			deadlineInMonths := deadline

			proposal := loan.NewProposal(proposalID, requiredValue, deadlineInMonths)
			proposals = append(
				proposals,
				&proposal,
			)

		case "warranty,added":
			if len(proposals) == 0 {
				return nil, fmt.Errorf("warranty added before proposal created at line: %v", n+1)
			}
			id := fields[WarrantyID]
			province := fields[WarrantyProvince]
			priceAsString := fields[WarrantyPrice]

			price, err := strconv.ParseFloat(priceAsString, 64)
			if err != nil {
				return nil, err
			}

			proposal := proposals[len(proposals)-1]
			proposal.AddWarranty(loan.NewWarranty(id, proposal.ID, price, province))
		case "proponent,added":
			if len(proposals) == 0 {
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

			proposal := proposals[len(proposals)-1]
			proposal.AddProponent(loan.NewProponent(id, proposal.ID, name, age, income, isMain))
		default:
			return nil, fmt.Errorf("unknown event type: %s", eventType)
		}
	}

	pp := make([]loan.Proposal, len(proposals))
	for i, p := range proposals {
		pp[i] = *p
	}

	return pp, nil
}

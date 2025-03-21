package loan

type Loan struct {
	proposal Proposal
}

type LoanRule func(Proposal) error

func NewCustomLoan(proposal Proposal, rules ...LoanRule) (Loan, error) {
	l := Loan{proposal}

	for _, rule := range rules {
		if err := rule(proposal); err != nil {
			return Loan{}, err
		}
	}

	return l, nil
}

// TODO: idk if these proxy methods makes sense or if it indicates a design problem
func (l Loan) Proponents() []Proponent {
	return l.proposal.Proponents()
}

func (l Loan) Warranties() []Warranty {
	return l.proposal.Warranties()
}

func (l Loan) DeadlineInMonths() int {
	return l.proposal.DeadlineInMonths()
}

func (l Loan) RequiredValue() float64 {
	return l.proposal.RequiredValue()
}

func (l Loan) ProposalID() string {
	return l.proposal.ID
}

package main

type Loan struct {
	proposal Proposal
}

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

func NewLoan(proposal Proposal) (Loan, error) {
	l := Loan{proposal}

	if err := l.selfValidate(); err != nil {
		return Loan{}, err
	}

	return l, nil
}

func (l *Loan) selfValidate() error {
	return l.proposal.SelfValidate()
}

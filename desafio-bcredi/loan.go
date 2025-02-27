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
	mainProponentsCount := 0

	for _, prop := range l.proposal.Proponents() {
		if prop.IsMain() {
			mainProponentsCount++
		}
	}

	if mainProponentsCount == 0 {
		return ErrMainProponentNotFound
	}

	if mainProponentsCount > 1 {
		return ErrInvalidNumberOfMainProponents
	}

	if l.proposal.MainProponent().Age() < 18 {
		return ErrMainProponentUnderage
	}

	if l.proposal.MainProponent().MonthlyIncome() < l.proposal.Installment() {
		return ErrMainProponentIncomeNotEnough
	}

	if len(l.proposal.Proponents()) < 2 {
		return ErrNotEnoughProponents
	}

	if l.proposal.WarrantiesValue() < l.proposal.RequiredValue()*2 {
		return ErrWarrantiesValueNotEnough
	}

	return nil
}

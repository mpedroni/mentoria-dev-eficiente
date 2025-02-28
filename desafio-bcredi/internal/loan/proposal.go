package loan

import (
	"errors"
)

var (
	// TODO: maybe now makes sense keep these validation errors closer to the rule definitions
	ErrMainProponentNotFound        = errors.New("a proposal must have a main proponent")
	ErrTooMuchMainProponents        = errors.New("a proposal must have only one main proponent")
	ErrMainProponentUnderage        = errors.New("the main proponent must be at least 18 years old")
	ErrMainProponentIncomeNotEnough = errors.New("the main proponent income must be enough to pay the proposal installment")
	ErrNotEnoughProponents          = errors.New("a proposal must have at least two proponents")
	ErrWarrantiesValueNotEnough     = errors.New("the warranties value must be at least twice the proposal value")
)

type Proposal struct {
	ID               string
	requiredValue    float64
	deadlineInMonths int
	warranties       []Warranty
	proponents       []Proponent
}

func NewProposal(id string, requiredValue float64, deadlineInMonths int) Proposal {
	return Proposal{
		ID:               id,
		requiredValue:    requiredValue,
		deadlineInMonths: deadlineInMonths,
	}
}

func (p *Proposal) AddWarranty(warranty Warranty) {
	p.warranties = append(p.warranties, warranty)
}

func (p *Proposal) AddProponent(proponent Proponent) {
	p.proponents = append(p.proponents, proponent)
}

// TODO: a proposal itself can have many or none main proponents, so maybe this method makes sense in Loan only
// probably it will imply receiving the Loan directly instead of the proposal in loan rules
func (p *Proposal) MainProponent() Proponent {
	for _, prop := range p.proponents {
		if prop.IsMain() {
			return prop
		}
	}

	return Proponent{}
}

func (p *Proposal) Installment() float64 {
	return p.requiredValue / float64(p.deadlineInMonths)
}

func (p *Proposal) WarrantiesValue() float64 {
	var total float64

	for _, w := range p.Warranties() {
		total += w.Price()
	}

	return total
}

func (p *Proposal) RequiredValue() float64 {
	return p.requiredValue
}

func (p *Proposal) DeadlineInMonths() int {
	return p.deadlineInMonths
}

func (p *Proposal) Warranties() []Warranty {
	return p.warranties
}

func (p *Proposal) Proponents() []Proponent {
	return p.proponents
}

func (p *Proposal) MainProponents() []Proponent {
	var mainProponents []Proponent

	for _, prop := range p.Proponents() {
		if prop.IsMain() {
			mainProponents = append(mainProponents, prop)
		}
	}

	return mainProponents
}

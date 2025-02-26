package main

import (
	"errors"
)

var (
	ErrMainProponentNotFound         = errors.New("a proposal must have a main proponent")
	ErrInvalidNumberOfMainProponents = errors.New("a proposal must have only one main proponent")
	ErrMainProponentUnderage         = errors.New("the main proponent must be at least 18 years old")
	ErrMainProponentIncomeNotEnough  = errors.New("the main proponent income must be enough to pay the proposal installment")
	ErrNotEnoughProponents           = errors.New("a proposal must have at least two proponents")
	ErrWarrantiesValueNotEnough      = errors.New("the warranties value must be at least twice the proposal value")
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

func (p *Proposal) SelfValidate() error {
	mainProponentsCount := 0

	for _, p := range p.Proponents() {
		if p.IsMain() {
			mainProponentsCount++
		}
	}

	if mainProponentsCount == 0 {
		return ErrMainProponentNotFound
	}

	if mainProponentsCount > 1 {
		return ErrInvalidNumberOfMainProponents
	}

	if p.MainProponent().Age() < 18 {
		return ErrMainProponentUnderage
	}

	if p.MainProponent().MonthlyIncome() < p.Installment() {
		return ErrMainProponentIncomeNotEnough
	}

	if len(p.Proponents()) < 2 {
		return ErrNotEnoughProponents
	}

	if p.WarrantiesValue() < p.RequiredValue()*2 {
		return ErrWarrantiesValueNotEnough
	}

	return nil
}

func (p *Proposal) MainProponent() Proponent {
	for _, p := range p.Proponents() {
		if p.IsMain() {
			return p
		}
	}

	panic("main proponent not found")
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

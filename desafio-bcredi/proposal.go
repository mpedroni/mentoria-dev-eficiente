package main

type Warranty struct {
	ID         string
	ProposalID string
	price      float64
}

func NewWarranty(ID string, proposalID string) Warranty {
	return Warranty{
		ID:         ID,
		ProposalID: proposalID,
	}
}

func (w *Warranty) Price() float64 {
	return w.price
}

type Proposal struct {
	ID         string
	warranties []Warranty
}

func NewProposal(id string, warranties []Warranty) Proposal {
	return Proposal{
		ID:         id,
		warranties: warranties,
	}
}

func (p *Proposal) Warranties() []Warranty {
	return p.warranties
}

package loan

type Warranty struct {
	ID         string
	ProposalID string
	price      float64
	province   string
}

func NewWarranty(ID string, proposalID string, price float64, province string) Warranty {
	return Warranty{
		ID:         ID,
		ProposalID: proposalID,
		price:      price,
		province:   province,
	}
}

func (w *Warranty) Price() float64 {
	return w.price
}

func (w *Warranty) Province() string {
	return w.province
}

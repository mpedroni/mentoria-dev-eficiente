package loan

type Proponent struct {
	ID         string
	ProposalID string
	name       string
	age        int
	income     float64
	isMain     bool
}

func NewProponent(ID string, proposalID string, name string, age int, income float64, isMain bool) Proponent {
	return Proponent{
		ID:         ID,
		ProposalID: proposalID,
		name:       name,
		age:        age,
		income:     income,
		isMain:     isMain,
	}
}

func (p Proponent) Name() string {
	return p.name
}

func (p Proponent) Age() int {
	return p.age
}

func (p Proponent) MonthlyIncome() float64 {
	return p.income
}

func (p Proponent) IsMain() bool {
	return p.isMain
}

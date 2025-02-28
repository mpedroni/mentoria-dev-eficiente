package loan

func NewRegularLoan(p Proposal) (Loan, error) {
	return NewCustomLoan(p,
		HasMainProponent(),
		MainProponentsMax(1),
		MainProponentIsLegalAge(),
		MainProponentIncomeEnough(),
		ProponentsMin(2),

		WarrantiesValueEnough(),
	)
}

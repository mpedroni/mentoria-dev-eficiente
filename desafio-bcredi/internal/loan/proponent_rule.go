package loan

func HasMainProponent() LoanRule {
	return func(p Proposal) error {
		if len(p.MainProponents()) == 0 {
			return ErrMainProponentNotFound
		}

		return nil
	}
}

func MainProponentsMax(max int) LoanRule {
	return func(p Proposal) error {
		if len(p.MainProponents()) > max {
			return ErrTooMuchMainProponents
		}

		return nil
	}
}

func ProponentsMin(min int) LoanRule {
	return func(p Proposal) error {
		if len(p.Proponents()) < min {
			return ErrNotEnoughProponents
		}

		return nil
	}
}

func MainProponentIsLegalAge() LoanRule {
	return func(p Proposal) error {
		if p.MainProponent().Age() < 18 {
			return ErrMainProponentUnderage
		}

		return nil
	}
}

func MainProponentIncomeEnough() LoanRule {
	return func(p Proposal) error {
		if p.MainProponent().MonthlyIncome() < p.Installment() {
			return ErrMainProponentIncomeNotEnough
		}

		return nil
	}
}

package main

func HasMainProponent() LoanRule {
	return func(p Proposal) error {
		if p.MainProponent().ID == "" {
			return ErrMainProponentNotFound
		}

		return nil
	}
}

func MainProponentsMax(max int) LoanRule {
	return func(p Proposal) error {
		mainProponentsCount := 0

		for _, prop := range p.Proponents() {
			if prop.IsMain() {
				mainProponentsCount++
			}
		}

		if mainProponentsCount > max {
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

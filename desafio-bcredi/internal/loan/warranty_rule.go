package loan

func WarrantiesValueEnough() LoanRule {
	return func(p Proposal) error {
		if p.WarrantiesValue() < p.RequiredValue()*2 {
			return ErrWarrantiesValueNotEnough
		}

		return nil
	}
}

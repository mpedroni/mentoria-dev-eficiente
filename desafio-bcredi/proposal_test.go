package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProposal_Proponent(t *testing.T) {
	warranties := []Warranty{
		NewWarranty("warranty_1", "prop_1", 3245356.0, "SC"),
		NewWarranty("warranty_2", "prop_1", 5000000.0, "PR"),
	}

	t.Run("should create a proposal when there are only one main proponent", func(t *testing.T) {
		proponents := []Proponent{
			NewProponent("proponent_1", "prop_1", "Ismael Streich Jr.", 42, 62615.64, true),
			NewProponent("proponent_2", "prop_1", "Arlene Kassulke", 36, 48925.52, false),
		}

		sut := NewProposal("prop_1", 1141424., 240)
		sut.AddWarranty(warranties[0])
		sut.AddWarranty(warranties[1])
		sut.AddProponent(proponents[0])
		sut.AddProponent(proponents[1])

		assert.Nil(t, sut.SelfValidate())
		assert.Equal(t, "prop_1", sut.ID)
		assert.Equal(t, 1141424., sut.RequiredValue())
		assert.Equal(t, 240, sut.DeadlineInMonths())
		assert.Len(t, sut.Warranties(), 2)
		assert.Len(t, sut.Proponents(), 2)
	})

	t.Run("should return error when there aren't a main proponent", func(t *testing.T) {
		proponents := []Proponent{
			NewProponent("proponent_1", "prop_1", "Ismael Streich Jr.", 42, 62615.64, false),
			NewProponent("proponent_2", "prop_1", "Arlene Kassulke", 36, 48925.52, false),
		}

		sut := NewProposal("prop_1", 1141424., 240)
		sut.AddWarranty(warranties[0])
		sut.AddWarranty(warranties[1])
		sut.AddProponent(proponents[0])
		sut.AddProponent(proponents[1])

		assert.ErrorIs(t, sut.SelfValidate(), ErrMainProponentNotFound)
	})

	t.Run("should return error when there are more than one main proponent", func(t *testing.T) {
		proponents := []Proponent{
			NewProponent("proponent_1", "prop_1", "Ismael Streich Jr.", 42, 62615.64, true),
			NewProponent("proponent_2", "prop_1", "Arlene Kassulke", 36, 48925.52, true),
		}

		sut := NewProposal("prop_1", 1141424., 240)
		sut.AddWarranty(warranties[0])
		sut.AddWarranty(warranties[1])
		sut.AddProponent(proponents[0])
		sut.AddProponent(proponents[1])

		assert.ErrorIs(t, sut.SelfValidate(), ErrInvalidNumberOfMainProponents)
	})

	t.Run("should return error when the main proponent is underage", func(t *testing.T) {
		proponents := []Proponent{
			NewProponent("proponent_1", "prop_1", "Ismael Streich Jr.", 17, 62615.64, true),
			NewProponent("proponent_2", "prop_1", "Arlene Kassulke", 36, 48925.52, false),
		}

		sut := NewProposal("prop_1", 1141424., 240)
		sut.AddWarranty(warranties[0])
		sut.AddWarranty(warranties[1])
		sut.AddProponent(proponents[0])
		sut.AddProponent(proponents[1])

		assert.ErrorIs(t, sut.SelfValidate(), ErrMainProponentUnderage)
	})

	t.Run("should return error when the main proponent income is less than the proposal installment", func(t *testing.T) {
		proposalPrice := 12000.0
		monthlyIncome := proposalPrice/12 - 1
		proponents := []Proponent{
			NewProponent("proponent_1", "prop_1", "Ismael Streich Jr.", 42, monthlyIncome, true),
			NewProponent("proponent_2", "prop_1", "Arlene Kassulke", 36, 48925.52, false),
		}

		sut := NewProposal("prop_1", proposalPrice, 12)
		sut.AddWarranty(warranties[0])
		sut.AddWarranty(warranties[1])
		sut.AddProponent(proponents[0])
		sut.AddProponent(proponents[1])

		assert.ErrorIs(t, sut.SelfValidate(), ErrMainProponentIncomeNotEnough)
	})

	t.Run("should return error when there are less than two proponents", func(t *testing.T) {
		proponents := []Proponent{
			NewProponent("proponent_1", "prop_1", "Ismael Streich Jr.", 42, 62615.64, true),
		}

		sut := NewProposal("prop_1", 1141424., 240)
		sut.AddWarranty(warranties[0])
		sut.AddWarranty(warranties[1])
		sut.AddProponent(proponents[0])

		assert.ErrorIs(t, sut.SelfValidate(), ErrNotEnoughProponents)
	})
}

func TestProposal_Warranties(t *testing.T) {
	proponents := []Proponent{
		NewProponent("proponent_1", "prop_1", "Ismael Streich Jr.", 42, 62615.64, true),
		NewProponent("proponent_2", "prop_1", "Arlene Kassulke", 36, 48925.52, false),
	}

	t.Run("should return error when the sum of the warranties values ins't at least twice the proposal value", func(t *testing.T) {
		warranties := []Warranty{
			NewWarranty("warranty_1", "prop_1", 10.0, "SC"),
			NewWarranty("warranty_2", "prop_1", 20.0, "PR"),
		}

		sut := NewProposal("prop_1", 60.1, 240)

		sut.AddWarranty(warranties[0])
		sut.AddWarranty(warranties[1])
		sut.AddProponent(proponents[0])
		sut.AddProponent(proponents[1])

		assert.ErrorIs(t, sut.SelfValidate(), ErrWarrantiesValueNotEnough)
	})
}

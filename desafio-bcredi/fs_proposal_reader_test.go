package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFsProposalReaderRead(t *testing.T) {
	sut, err := NewFileSystemProposalReader("./proposals/example001.txt")
	assert.Nil(t, err)

	t.Run("should return a list of proposals", func(t *testing.T) {
		proposals, err := sut.Read()
		assert.Nil(t, err)

		assert.Equal(t, 1, len(proposals))
	})

	t.Run("should identify the proposal", func(t *testing.T) {
		proposals, err := sut.Read()
		assert.Nil(t, err)

		p := proposals[0]
		assert.Equal(t, "prop_1", p.ID)
		assert.Equal(t, 1141424., p.RequiredValue())
		assert.Equal(t, 240, p.DeadlineInMonths())
	})

	t.Run("should identify warranties on the proposal", func(t *testing.T) {
		proposals, err := sut.Read()
		assert.Nil(t, err)

		proposal := proposals[0]
		warranties := proposal.Warranties()
		assert.Equal(t, 2, len(warranties))

		w := warranties[0]
		assert.Equal(t, "warranty_1", w.ID)
		assert.Equal(t, "prop_1", w.ProposalID)
		assert.Equal(t, 3245356.0, w.Price())
	})

	t.Run("should identify the proponents on the proposal", func(t *testing.T) {
		proposals, err := sut.Read()
		assert.Nil(t, err)

		proposal := proposals[0]
		proponents := proposal.Proponents()
		assert.Equal(t, 2, len(proponents))

		p := proponents[0]
		assert.Equal(t, "proponent_1", p.ID)
		assert.Equal(t, "prop_1", p.ProposalID)
		assert.Equal(t, "Ismael Streich Jr.", p.Name())
		assert.Equal(t, 42, p.Age())
		assert.Equal(t, 62615.64, p.MonthlyIncome())
		assert.Equal(t, true, p.IsMain())
	})

	t.Run("should handle files with more than one proposal", func(t *testing.T) {
		sut, err := NewFileSystemProposalReader("./proposals/example002.txt")
		assert.Nil(t, err)

		proposals, err := sut.Read()
		assert.Nil(t, err)

		assert.Equal(t, 2, len(proposals))

		p1 := proposals[0]
		assert.Equal(t, "80921e5f-4307-4623-9ddb-5bf826a31dd7", p1.ID)
		assert.Equal(t, 1141424., p1.RequiredValue())
		assert.Equal(t, 240, p1.DeadlineInMonths())
		assert.Equal(t, 2, len(p1.Warranties()))
		assert.Equal(t, 2, len(p1.Proponents()))

		p2 := proposals[1]
		assert.Equal(t, "52f0b3f2-f838-4ce2-96ee-9876dd2c0cf6", p2.ID)
		assert.Equal(t, 2689584.0, p2.RequiredValue())
		assert.Equal(t, 72, p2.DeadlineInMonths())
		assert.Equal(t, 2, len(p2.Warranties()))
		assert.Equal(t, 3, len(p2.Proponents()))
	})
}

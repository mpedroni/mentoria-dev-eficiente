package main

import (
	"os"
	"testing"
)

func TestFsProposalReaderRead(t *testing.T) {
	f, err := os.ReadFile("./proposals/example001.txt")
	if err != nil {
		t.Fatal(err)
	}

	sut := NewFileSystemProposalReader(string(f))

	t.Run("should return a list of proposals", func(t *testing.T) {
		proposals, err := sut.Read()
		if err != nil {
			t.Fatal(err)
		}

		if len(proposals) != 1 {
			t.Fatalf("expected 1 proposal, got %d", len(proposals))
		}
	})

	t.Run("should identify the proposal id", func(t *testing.T) {
		proposals, err := sut.Read()
		if err != nil {
			t.Fatal(err)
		}

		proposal := proposals[0]
		if proposal.ID != "prop_1" {
			t.Fatalf("expected proposal id %s, got %s", "prop_1", proposal.ID)
		}
	})

	t.Run("should identify warranties on the proposal", func(t *testing.T) {
		proposals, err := sut.Read()
		if err != nil {
			t.Fatal(err)
		}

		proposal := proposals[0]
		warranties := proposal.Warranties()
		if len(warranties) != 2 {
			t.Fatalf("expected 2 warranties, got %d", len(warranties))
		}

		w := warranties[0]
		if w.ProposalID != "prop_1" {
			t.Fatalf("expected warranty proposal id %s, got %s", "prop_1", w.ProposalID)
		}

		if w.ID != "warranty_1" {
			t.Fatalf("expected warranty id %s, got %s", "warr_1", w.ID)
		}

		if w.Price() != 3245356.0 {
			t.Fatalf("expected warranty price %f, got %f", 3245356.0, w.Price())
		}
	})

}

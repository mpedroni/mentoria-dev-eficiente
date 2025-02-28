package main

import (
	"desafio_bcredi/internal/loan"
	"desafio_bcredi/internal/proposalreader"
	"fmt"
)

type CLI struct {
	r proposalreader.ProposalReader
}

func NewCLI(r proposalreader.ProposalReader) *CLI {
	return &CLI{r}
}

func (c *CLI) Run() error {
	proposals, err := c.r.Read()

	if err != nil {
		return err
	}

	for _, p := range proposals {
		_, err := loan.NewRegularLoan(p)
		if err != nil {
			fmt.Println(p.ID, "is invalid:", err)
		} else {
			fmt.Println(p.ID, "is valid")
		}
	}

	return nil
}

package main

import "fmt"

type CLI struct {
	r ProposalReader
}

func NewCLI(r ProposalReader) *CLI {
	return &CLI{r}
}

func (c *CLI) Run() error {
	proposals, err := c.r.Read()

	if err != nil {
		return err
	}

	for _, p := range proposals {
		if p.Err != nil {
			fmt.Println(p.ID, "is invalid:", p.Err)
		} else {
			fmt.Println(p.ID, "is valid")
		}
	}

	return nil
}

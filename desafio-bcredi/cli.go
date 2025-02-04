package main

type CLI struct {
	r ProposalReader
}

func NewCLI(r ProposalReader) *CLI {
	return &CLI{r}
}

func (c *CLI) Run() error {
	return nil
}

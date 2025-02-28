package main

import "desafio_bcredi/internal/proposalreader"

func main() {
	fsProposalReader, err := proposalreader.NewFileSystemProposalReader("./proposals/example002.txt")
	if err != nil {
		panic(err)
	}

	cli := NewCLI(fsProposalReader)

	if err := cli.Run(); err != nil {
		panic(err)
	}
}

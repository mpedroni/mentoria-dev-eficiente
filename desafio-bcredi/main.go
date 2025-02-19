package main

func main() {
	fsProposalReader, err := NewFileSystemProposalReader("./proposals/example002.txt")
	if err != nil {
		panic(err)
	}

	cli := NewCLI(fsProposalReader)

	if err := cli.Run(); err != nil {
		panic(err)
	}
}

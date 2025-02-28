package proposalreader

import "desafio_bcredi/internal/loan"

type ProposalReader interface {
	Read() ([]loan.Proposal, error)
}

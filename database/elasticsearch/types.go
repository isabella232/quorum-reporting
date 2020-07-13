package elasticsearch

import (
	"quorumengineering/quorum-report/types"
)

type Contract struct {
	Address             types.Address `json:"address"`
	TemplateName        string        `json:"templateName"`
	CreationTransaction types.Hash    `json:"creationTx"`
	LastFiltered        uint64        `json:"lastFiltered"`
}

type Template struct {
	TemplateName string `json:"templateName"`
	ABI          string `json:"abi"`
	StorageABI   string `json:"storageAbi"`
}

type State struct {
	Address     types.Address `json:"address"`
	BlockNumber uint64        `json:"blockNumber"`
	StorageRoot types.Hash    `json:"storageRoot"`
}

type Storage struct {
	StorageRoot types.Hash        `json:"storageRoot"`
	StorageMap  map[string]string `json:"storageMap"`
}

//

type ContractQueryResult struct {
	Source Contract `json:"_source"`
}

type TemplateQueryResult struct {
	Source Template `json:"_source"`
}

type TransactionQueryResult struct {
	Source *types.Transaction `json:"_source"`
}

type BlockQueryResult struct {
	Source *types.Block `json:"_source"`
}

type StateQueryResult struct {
	Source State `json:"_source"`
}

type StorageQueryResult struct {
	Source Storage `json:"_source"`
}

type LastPersistedResult struct {
	Source struct {
		LastPersisted uint64 `json:"lastPersisted"`
	} `json:"_source"`
}

type SearchQueryResult struct {
	Hits struct {
		Hits []IndividualResult `json:"hits"`
	} `json:"hits"`
}

type CountQueryResult struct {
	Count uint64 `json:"count"`
}

type IndividualResult struct {
	Id     string                 `json:"_id"`
	Source map[string]interface{} `json:"_source"`
}

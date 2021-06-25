package types

const (
	// ModuleName is the name of the module
	ModuleName = "blockchaintest"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

const (
	WhoisPrefix      = "whois-value-"
	WhoisCountPrefix = "whois-count-"
)

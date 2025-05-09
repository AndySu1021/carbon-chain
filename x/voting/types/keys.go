package types

const (
	// ModuleName defines the module name
	ModuleName = "voting"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_voting"
)

var (
	ParamsKey = []byte("p_voting")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

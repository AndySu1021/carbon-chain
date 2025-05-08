package types

const (
	// ModuleName defines the module name
	ModuleName = "carbonchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_carbonchain"
)

var (
	ParamsKey = []byte("p_carbonchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

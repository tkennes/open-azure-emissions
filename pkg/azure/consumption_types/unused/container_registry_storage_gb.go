package azure_unit_of_measure

var (
	// [registryType: string]: Storage in GB
	CONTAINER_REGISTRY_STORAGE_GB = map[string]int{
		"Basic":    10,
		"Standard": 100,
		"Premium":  500,
	}
)

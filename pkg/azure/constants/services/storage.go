package azure_services

var (
	STORAGE_REDUNDANCY_TO_REPLICAS = map[string]int{
		"":        1,
		"LRS":     3,
		"ZRS":     3,
		"GRS":     6,
		"RA-GRS":  6,
		"GZRS":    6,
		"RA-GZRS": 6,
	}
)

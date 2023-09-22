package azure_unit_of_measure

var (
	NETWORKING_USAGE_UNITS = map[string]string{
		"GB_1":   "1 GB",
		"TB_1":   "1 TB",
		"GB_10":  "10 GB",
		"GB_100": "100 GB",
		"GB_200": "200 GB",
	}

	NETWORKING_USAGE_TYPES = []string{
		"Geo-Replication Data transfer",
		"Data Transfer Out",
		"Egress",
		"Geo-Replication v2 Data transfer",
		"Data Processed - Egress",
	}
)

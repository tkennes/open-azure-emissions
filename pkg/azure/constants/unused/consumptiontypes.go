package azure_constants

var (
	UNKNOWN_USAGE_TYPES = [...]string{
		"Server - Free",
		"Requests",
		"Custom Domain",
	}

	UNSUPPORTED_USAGE_TYPES = [...]string{
		"Rulesets",
		"Rules",
		"Policies",
		"Kafka Surcharge",
		"License",
	}

	COMPUTE_USAGE_UNITS = [...]string{
		"1 Hour",
		"10 Hour",
		"10 Hours",
		"100 Hour",
		"100 Hours",
		"1000 Hour",
		"1000 Hours",
	}

	STORAGE_USAGE_UNITS = [...]string{
		"1 /Month",
		"100 /Month",
		"1 GB/Month",
		"10 GB/Month",
		"100 GB/Month",
		"10 /Day",
		"30 /Day",
		"1 TB/Month",
	}

	NETWORKING_USAGE_UNITS = [...]string{
		"1 GB",
		"1 TB",
		"10 GB",
		"100 GB",
		"200 GB",
	}

	MEMORY_USAGE_UNITS = [...]string{
		"50000 GB Seconds",
		"1000 GB Hours",
	}

	STORAGE_USAGE_TYPES = [...]string{
		"Data Stored",
		"Metadata",
		"Registry Unit",
		"ZRS",
		"LRS",
		"GRS",
		"GZRS",
		"Data Retention",
		"Pay-as-you-go Data at Rest",
		"Standard Instances",
		"Node",
		"10 DTUs",
		"S0 DTUs",
		"B DTUs",
		"B DTU",
		"eDTUs",
		"On Premises Server Protected Instances",
		"Standard Trial Nodes",
		"Azure VM Protected Instances",
		"Standard User",
		"Multi-step Web Test",
		"S0 Secondary Active DTUs",
		"Resource Monitored at 5 Minute Frequency",
		"VM Replicated to Azure",
		"Basic User",
		"Standard Nodes",
		"Microsoft-hosted CI",
		"Data Processing Unit",
		"Pay as you go Warm Storage",
		"Standard Unit",
	}

	NETWORKING_USAGE_TYPES = [...]string{
		"Geo-Replication Data transfer",
		"Data Transfer Out",
		"Egress",
		"Geo-Replication v2 Data transfer",
		"Data Processed - Egress",
	}

	MEMORY_USAGE_TYPES = [...]string{
		"Execution Time",
		"Memory Duration",
		"C1 Cache Instance",
	}

	// Registry Type
	CONTAINER_REGISTRY_STORAGE_GB = map[string]int{
		"Basic":    10,
		"Standard": 100,
		"Premium":  500,
	}

	// Azure GroupBys (Query Date Types)
	AZURE_QUERY_GROUP_BY = map[string]string{
		"day":     "day",
		"week":    "isoWeek",
		"month":   "month",
		"quarter": "quarter",
		"year":    "year",
	}
)

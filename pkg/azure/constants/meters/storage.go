package azure_constants_meters

var (
	ESTIMABLE_STORAGE_SUBCATEGORIES = []string{
		"Standard SSD Managed Disks",
		"Premium SSD Managed Disks",
		"Premium Page Blob",
		"Standard HDD Managed Disks",
	}

	UNESTIMABLE_STORAGE_METER_NAMES = []string{
		"Disk Delete Operations",
		"Disk Delete Operations - Free",

		"Disk Operations",
		"Disk Operations - Free",

		"Disk Read Operations",
		"Disk Read Operations - Free",

		"Disk Write Operations",
		"Disk Write Operations - Free",

		// Redundancy operations: adding all of them even though
		// they are not all seen as meter names in the data yet
		"LRS Disk Write Operations",
		"LRS Disk Write Operations - Free",

		"ZRS Disk Write Operations",
		"ZRS Disk Write Operations - Free",

		"GRS Disk Write Operations",
		"GRS Disk Write Operations - Free",

		"RA-GRS Disk Write Operations",
		"RA-GRS Disk Write Operations - Free",

		"GZRS Disk Write Operations",
		"GZRS Disk Write Operations - Free",

		// Snapshots
		"LRS Snapshots",
		"ZRS Snapshots",
		"GRS Snapshots",
		"RA-GRS Snapshots",
		"GZRS Snapshots",
	}
)

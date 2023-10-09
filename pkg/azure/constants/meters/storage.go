package azure_constants_meters

type AzureEstimableStorage struct {
	MeterSubCategory string
	StorageType      string
	MeterName        string
	MeterNameHas     string
}

type AzureEstimableStorages map[string]AzureEstimableStorage

func (azureEstimableStorages AzureEstimableStorages) meterSubCategories() []string {
	var list []string
	for _, azureEstimableStorage := range azureEstimableStorages {
		list = append(list, azureEstimableStorage.MeterSubCategory)
	}
	return list
}

var (
	ESTIMABLE_STORAGE_USAGE = map[string]AzureEstimableStorage{
		// Blob/Object Storage
		"Azure Data Lake Storage Gen2 Flat Namespace": {
			MeterSubCategory: "Azure Data Lake Storage Gen2 Flat Namespace",
			StorageType:      "Blob",
			MeterNameHas:     "Data Stored",
		},
		"Azure Data Lake Storage Gen2 Hierarchical Namespace": {
			MeterSubCategory: "Azure Data Lake Storage Gen2 Hierarchical Namespace",
			StorageType:      "Blob",
			MeterNameHas:     "Data Stored",
		},
		"Blob Storage": {
			MeterSubCategory: "Blob Storage",
			StorageType:      "Blob",
			MeterNameHas:     "Data Stored",
		},
		"Files": {
			MeterSubCategory: "Files",
			StorageType:      "Blob",
			MeterNameHas:     "Data Stored",
		},
		"General Block Blob": {
			MeterSubCategory: "General Block Blob",
			StorageType:      "Blob",
			MeterNameHas:     "Data Stored",
		},
		"General Block Blob v2": {
			MeterSubCategory: "General Block Blob v2",
			StorageType:      "Blob",
			MeterNameHas:     "Data Stored",
		},
		"General Block Blob v2 Hierarchical Namespace": {
			MeterSubCategory: "General Block Blob v2 Hierarchical Namespace",
			StorageType:      "Blob",
			MeterNameHas:     "Data Stored",
		},
		"Premium Block Blob": {
			MeterSubCategory: "Premium Block Blob",
			StorageType:      "Blob",
			MeterNameHas:     "Data Stored",
		},
		"Queues": {
			MeterSubCategory: "Queues",
			StorageType:      "Blob",
			MeterNameHas:     "Data Stored",
		},
		"Standard Page Blob": {
			MeterSubCategory: "Standard Page Blob",
			StorageType:      "Blob",
			MeterNameHas:     "Data Stored",
		},
		"Standard Page Blob v2": {
			MeterSubCategory: "Standard Page Blob v2",
			StorageType:      "Blob",
			MeterNameHas:     "Data Stored",
		},
		"Tables": {
			MeterSubCategory: "Tables",
			StorageType:      "Blob",
			MeterNameHas:     "Data Stored",
		},

		// Disk Storage
		"Standard SSD Managed Disks": {
			MeterSubCategory: "Standard SSD Managed Disks",
			StorageType:      "Disk",
		},
		"Premium SSD Managed Disks": {
			MeterSubCategory: "Premium SSD Managed Disks",
			StorageType:      "Disk",
		},
		"Premium Page Blob": {
			MeterSubCategory: "Premium Page Blob",
			StorageType:      "Disk",
		},
		"Standard HDD Managed Disks": {
			MeterSubCategory: "Standard HDD Managed Disks",
			StorageType:      "Disk",
		},
	}

	ESTIMABLE_STORAGE_SUBCATEGORIES = ESTIMABLE_STORAGE_USAGE.meterSubCategories()

	UNESTIMABLE_DISK_STORAGE_METER_NAMES = []string{
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

		// Snapshots -> Microsoft.Compute
		"LRS Snapshots",
		"ZRS Snapshots",
		"GRS Snapshots",
		"RA-GRS Snapshots",
		"GZRS Snapshots",
	}
)

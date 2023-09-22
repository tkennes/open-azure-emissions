package azure_services

import (
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

var (
	MANAGED_DISKS = map[string]azure_models.AzureDisk{
		// SSD
		"P1": azure_models.AzureDisk{
			Name: "P1",
			Size: 4,
			Type: "SSD",
		},
		"P2": azure_models.AzureDisk{
			Name: "P2",
			Size: 8,
			Type: "SSD",
		},
		"P3": azure_models.AzureDisk{
			Name: "P3",
			Size: 16,
			Type: "SSD",
		},
		"P4": azure_models.AzureDisk{
			Name: "P4",
			Size: 32,
			Type: "SSD",
		},
		"P6": azure_models.AzureDisk{
			Name: "P6",
			Size: 64,
			Type: "SSD",
		},
		"P10": azure_models.AzureDisk{
			Name: "P10",
			Size: 128,
			Type: "SSD",
		},
		"P15": azure_models.AzureDisk{
			Name: "P15",
			Size: 256,
			Type: "SSD",
		},
		"P20": azure_models.AzureDisk{
			Name: "P20",
			Size: 512,
			Type: "SSD",
		},
		"P30": azure_models.AzureDisk{
			Name: "P30",
			Size: 1024,
			Type: "SSD",
		},
		"P40": azure_models.AzureDisk{
			Name: "P40",
			Size: 2048,
			Type: "SSD",
		},
		"P50": azure_models.AzureDisk{
			Name: "P50",
			Size: 4096,
			Type: "SSD",
		},
		"P60": azure_models.AzureDisk{
			Name: "P60",
			Size: 8192,
			Type: "SSD",
		},
		"P70": azure_models.AzureDisk{
			Name: "P70",
			Size: 16384,
			Type: "SSD",
		},
		"P80": azure_models.AzureDisk{
			Name: "P80",
			Size: 32767,
			Type: "SSD",
		},
		"E1": azure_models.AzureDisk{
			Name: "E1",
			Size: 4,
			Type: "SSD",
		},
		"E2": azure_models.AzureDisk{
			Name: "E2",
			Size: 8,
			Type: "SSD",
		},
		"E3": azure_models.AzureDisk{
			Name: "E3",
			Size: 16,
			Type: "SSD",
		},
		"E4": azure_models.AzureDisk{
			Name: "E4",
			Size: 32,
			Type: "SSD",
		},
		"E6": azure_models.AzureDisk{
			Name: "E6",
			Size: 64,
			Type: "SSD",
		},
		"E10": azure_models.AzureDisk{
			Name: "E10",
			Size: 128,
			Type: "SSD",
		},
		"E15": azure_models.AzureDisk{
			Name: "E15",
			Size: 256,
			Type: "SSD",
		},
		"E20": azure_models.AzureDisk{
			Name: "E20",
			Size: 512,
			Type: "SSD",
		},
		"E30": azure_models.AzureDisk{
			Name: "E30",
			Size: 1024,
			Type: "SSD",
		},
		"E40": azure_models.AzureDisk{
			Name: "E40",
			Size: 2048,
			Type: "SSD",
		},
		"E50": azure_models.AzureDisk{
			Name: "E50",
			Size: 4096,
			Type: "SSD",
		},
		"E60": azure_models.AzureDisk{
			Name: "E60",
			Size: 8192,
			Type: "SSD",
		},
		"E70": azure_models.AzureDisk{
			Name: "E70",
			Size: 16384,
			Type: "SSD",
		},
		"E80": azure_models.AzureDisk{
			Name: "E80",
			Size: 32767,
			Type: "SSD",
		},

		// HDDs
		"S4": azure_models.AzureDisk{
			Name: "S4",
			Size: 32,
			Type: "HDD",
		},
		"S6": azure_models.AzureDisk{
			Name: "S6",
			Size: 64,
			Type: "HDD",
		},
		"S10": azure_models.AzureDisk{
			Name: "S10",
			Size: 128,
			Type: "HDD",
		},
		"S15": azure_models.AzureDisk{
			Name: "S15",
			Size: 256,
			Type: "HDD",
		},
		"S20": azure_models.AzureDisk{
			Name: "S20",
			Size: 512,
			Type: "HDD",
		},
		"S30": azure_models.AzureDisk{
			Name: "S30",
			Size: 1024,
			Type: "HDD",
		},
		"S40": azure_models.AzureDisk{
			Name: "S40",
			Size: 2048,
			Type: "HDD",
		},
		"S50": azure_models.AzureDisk{
			Name: "S50",
			Size: 4096,
			Type: "HDD",
		},
		"S60": azure_models.AzureDisk{
			Name: "S60",
			Size: 8192,
			Type: "HDD",
		},
		"S70": azure_models.AzureDisk{
			Name: "S70",
			Size: 16384,
			Type: "HDD",
		},
		"S80": azure_models.AzureDisk{
			Name: "S80",
			Size: 767,
			Type: "HDD",
		},
	}

	MANAGED_DISK_REPLICAS = map[string]int{
		"":        1,
		"LRS":     3,
		"ZRS":     3,
		"GRS":     6,
		"RA-GRS":  6,
		"GZRS":    6,
		"RA-GZRS": 6,
	}
)

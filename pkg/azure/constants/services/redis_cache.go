package azure_services

import (
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

var (
	REDIS_CACHES = map[string]azure_models.RedisCache{
		"C0": azure_models.RedisCache{
			Name: "C0",
			Size: 0.25,
		},
		"C1": azure_models.RedisCache{
			Name: "C1",
			Size: 1,
		},
		"C2": azure_models.RedisCache{
			Name: "C2",
			Size: 2.5,
		},
		"C3": azure_models.RedisCache{
			Name: "C3",
			Size: 6,
		},
		"C4": azure_models.RedisCache{
			Name: "C4",
			Size: 13,
		},
		"C5": azure_models.RedisCache{
			Name: "C5",
			Size: 26,
		},
		"C6": azure_models.RedisCache{
			Name: "C6",
			Size: 53,
		},
		"P1": azure_models.RedisCache{
			Name:   "P1",
			Size:   6,
			Series: "Premium",
		},
		"P2": azure_models.RedisCache{
			Name:   "P2",
			Size:   13,
			Series: "Premium",
		},
		"P3": azure_models.RedisCache{
			Name:   "P3",
			Size:   26,
			Series: "Premium",
		},
		"P4": azure_models.RedisCache{
			Name:   "P4",
			Size:   53,
			Series: "Premium",
		},
		"P5": azure_models.RedisCache{
			Name:   "P5",
			Size:   120,
			Series: "Premium",
		},
		"E10": azure_models.RedisCache{
			Name:   "E10",
			Size:   12,
			Series: "Enterprise",
		},
		"E20": azure_models.RedisCache{
			Name:   "E20",
			Size:   25,
			Series: "Enterprise",
		},
		"E50": azure_models.RedisCache{
			Name:   "E50",
			Size:   50,
			Series: "Enterprise",
		},
		"E100": azure_models.RedisCache{
			Name:   "E100",
			Size:   100,
			Series: "Enterprise",
		},
		"F300": azure_models.RedisCache{
			Name:   "F300",
			Size:   384,
			Series: "EnterpriseFlash",
		},
		"F700": azure_models.RedisCache{
			Name:   "F700",
			Size:   715,
			Series: "EnterpriseFlash",
		},
		"F1500": azure_models.RedisCache{
			Name:   "F1500",
			Size:   1455,
			Series: "EnterpriseFlash",
		},
	}

	REDIS_CACHE_REPLICAS = map[string]int{
		"Basic":    1,
		"Standard": 2,
		// Not Sure about these numbers
		"Premium":         3,
		"Enterprise":      3,
		"EnterpriseFlash": 3,
	}
)

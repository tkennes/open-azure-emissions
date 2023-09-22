package azure_services

import (
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

var (
	// [VCPUs , Memory , embodied emissions]
	APP_SERVICE_PLANS = map[string]azure_models.AppServicePlan{
		// None
		"vCPU Duration": azure_models.AppServicePlan{
			Name:   "vCPU Duration",
			Series: "None",
			VCPUs:  1,
		},

		"vCore": azure_models.AppServicePlan{
			Name:   "vCore",
			Series: "None",
			VCPUs:  1,
		},

		"2 vCore": azure_models.AppServicePlan{
			Name:   "2 vCore",
			Series: "None",
			VCPUs:  2,
		},

		// Free Plan
		"F1": azure_models.AppServicePlan{
			Name:   "F1",
			Series: "Free Service Plan",
			VCPUs:  1,
			Memory: 3.5,
		},

		// Shared Service Plan
		"D1": azure_models.AppServicePlan{
			Name:   "D1",
			Series: "Shared Service Plan",
			VCPUs:  1,
			Memory: 3.5,
		},
		"D2": azure_models.AppServicePlan{
			Name:   "D2",
			Series: "Shared Service Plan",
			VCPUs:  2,
			Memory: 7,
		},
		"D3": azure_models.AppServicePlan{
			Name:   "D3",
			Series: "Shared Service Plan",
			VCPUs:  4,
			Memory: 14,
		},
		"D4": azure_models.AppServicePlan{
			Name:   "D4",
			Series: "Shared Service Plan",
			VCPUs:  8,
			Memory: 28,
		},

		// Basic Service Plan
		"B1": azure_models.AppServicePlan{
			Name:   "B1",
			Series: "Basic Service Plan",
			VCPUs:  1,
			Memory: 1.75,
		},
		"B2": azure_models.AppServicePlan{
			Name:   "B2",
			Series: "Basic Service Plan",
			VCPUs:  2,
			Memory: 3.5,
		},
		"B3": azure_models.AppServicePlan{
			Name:   "B3",
			Series: "Basic Service Plan",
			VCPUs:  4,
			Memory: 7,
		},

		// Standard Service Plan
		"S1": azure_models.AppServicePlan{
			Name:   "S1",
			Series: "Standard Service Plan",
			VCPUs:  1,
			Memory: 1.75,
		},
		"S2": azure_models.AppServicePlan{
			Name:   "S2",
			Series: "Standard Service Plan",
			VCPUs:  2,
			Memory: 3.5,
		},
		"S3": azure_models.AppServicePlan{
			Name:   "S3",
			Series: "Standard Service Plan",
			VCPUs:  4,
			Memory: 7,
		},

		// Premium v2 Service Plan
		"P1 v2": azure_models.AppServicePlan{
			Name:   "P1 v2",
			Series: "Premium v2 Service Plan",
			VCPUs:  1,
			Memory: 3.5,
		},
		"P2 v2": azure_models.AppServicePlan{
			Name:   "P2 v2",
			Series: "Premium v2 Service Plan",
			VCPUs:  2,
			Memory: 7,
		},
		"P3 v2": azure_models.AppServicePlan{
			Name:   "P3 v2",
			Series: "Premium v2 Service Plan",
			VCPUs:  4,
			Memory: 14,
		},

		// Premium v3 Service Plan
		"P1 v3": azure_models.AppServicePlan{
			Name:   "P1 v3",
			Series: "Premium v3 Service Plan",
			VCPUs:  2,
			Memory: 8,
		},
		"P2 v3": azure_models.AppServicePlan{
			Name:   "P2 v3",
			Series: "Premium v3 Service Plan",
			VCPUs:  4,
			Memory: 16,
		},
		"P3 v3": azure_models.AppServicePlan{
			Name:   "P3 v3",
			Series: "Premium v3 Service Plan",
			VCPUs:  8,
			Memory: 32,
		},

		// Isolated Service Plan
		"I1": azure_models.AppServicePlan{
			Name:   "I1",
			Series: "Isolated Service Plan",
			VCPUs:  1,
			Memory: 3.5,
		},
		"I2": azure_models.AppServicePlan{
			Name:   "I2",
			Series: "Isolated Service Plan",
			VCPUs:  2,
			Memory: 7,
		},
		"I3": azure_models.AppServicePlan{
			Name:   "I3",
			Series: "Isolated Service Plan",
			VCPUs:  4,
			Memory: 14,
		},

		// Isolated v2 Plan
		"I1 v2": azure_models.AppServicePlan{
			Name:   "I1 v2",
			Series: "Isolated v2 Plan",
			VCPUs:  2,
			Memory: 8,
		},
		"I2 v2": azure_models.AppServicePlan{
			Name:   "I2 v2",
			Series: "Isolated v2 Plan",
			VCPUs:  4,
			Memory: 16,
		},
		"I3 v2": azure_models.AppServicePlan{
			Name:   "I3 v2",
			Series: "Isolated v2 Plan",
			VCPUs:  8,
			Memory: 32,
		},
	}
)

package azure_services

import (
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

var (
	// [registryType: string]: Storage in GB
	CONTAINER_REGISTRIES = map[string]azure_models.AzureContainerRegistry{
		"Basic": azure_models.AzureContainerRegistry{
			Name: "Basic",
			Size: 10,
		},
		"Standard": azure_models.AzureContainerRegistry{
			Name: "Standard",
			Size: 100,
		},
		"Premium": azure_models.AzureContainerRegistry{
			Name: "Premium",
			Size: 500,
		},
	}
)

package azure_footprint

import (
	"errors"
	"fmt"

	"github.com/tkennes/open-azure-emissions/pkg/util/parsers"

	azure_constants_services "github.com/tkennes/open-azure-emissions/pkg/azure/constants/services"
	azure_footprint_core "github.com/tkennes/open-azure-emissions/pkg/azure/footprint/core"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func EstimateContainerRegistryEnergyConsumption(data azure_models.AzureCostDetails) (float64, error) {
	sku := parsers.ParseContainerRegistrySku(data.product)

	if containerRegistry, found := azure_constants_services.CONTAINER_REGISTRIES[sku]; found {
		replicas := 1
		terabytes := containerRegistry.Size / 1000
		// We're assuming SSD
		diskType := "SSD"

		return azure_footprint_core.EstimateStorageEnergy(
			data,
			diskType,
			terabytes,
			replicas,
		)
	} else {
		return 0, errors.New(fmt.Sprintf("Container Registry Product \"%s\" (sku: %s): could not find container registry", data.Product, sku))
	}
}

package azure_footprint

import (
	"github.com/tkennes/open-azure-emissions/pkg/util/parsers"

	azure_constants_services "github.com/tkennes/open-azure-emissions/pkg/azure/constants/services"
	azure_footprint_core "github.com/tkennes/open-azure-emissions/pkg/azure/footprint/core"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func EstimateBlobEnergyConsumption(data azure_models.AzureCostDetails) (float64, error) {

	redundancy := parsers.ParseStorageRedundancy(data.MeterName)
	replicas := azure_constants_services.STORAGE_REDUNDANCY_TO_REPLICAS[redundancy]
	diskType := "SSD"
	terabytes := 1 / 1000

	return azure_footprint_core.EstimateStorageEnergy(
		data,
		diskType,
		terabytes,
		replicas,
	)
}

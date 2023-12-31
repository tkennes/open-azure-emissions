package azure_footprint

import (
	"errors"
	"fmt"

	"github.com/tkennes/open-azure-emissions/pkg/util/parsers"

	azure_constants_services "github.com/tkennes/open-azure-emissions/pkg/azure/constants/services"
	azure_footprint_core "github.com/tkennes/open-azure-emissions/pkg/azure/footprint/core"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func EstimateManagedDiskEnergyConsumption(data azure_models.AzureCostDetails) (float64, error) {
	diskName := parsers.ParseManagedDiskType(data.MeterName)
	if managedDisk, found := azure_constants_services.MANAGED_DISKS[diskName]; found {
		redundancy := parsers.ParseStorageRedundancy(data.MeterName)
		replicas := azure_constants_services.STORAGE_REDUNDANCY_TO_REPLICAS[redundancy]

		terabytes := managedDisk.Size / 1000

		return azure_footprint_core.EstimateStorageEnergy(
			data,
			managedDisk.Type,
			terabytes,
			replicas,
		)

	} else {
		return 0, errors.New(fmt.Sprintf("Managed Disk Meter Name \"%s\" (disk name: %s): could not parse/find machine", data.MeterName, diskName))
	}
}

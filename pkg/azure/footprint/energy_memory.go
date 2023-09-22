package azure_footprint

import (
	"github.com/tkennes/open-azure-emissions/pkg/util/parsers"

	azure_constants "github.com/tkennes/open-azure-emissions/pkg/azure/constants"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func MemoryEnergyCalculation(
	gigaByteHours float64,
	memoryCoefficient float64,
	replicas int,
) float64 {
	return gigaByteHours * memoryCoefficient * float64(replicas)
}

func EstimateMemoryEnergy(data azure_models.AzureCostDetails, memorySize float64, replicas int) (float64, error) {
	hours, err := parsers.ParseHours(data)
	if err != nil {
		return 0, err
	}
	return MemoryEnergyCalculation(
		calculateGigaByteHours(memorySize, data.Quantity, hours),
		azure_constants.MEMORY_ENERGY_COEFFICIENT,
		replicas,
	), nil
}

func calculateGigaByteHours(gigabytes float64, quantity float64, hours float64) float64 {
	return gigabytes * quantity * hours
}

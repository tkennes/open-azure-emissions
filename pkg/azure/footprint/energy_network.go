package azure_footprint

import (
	"github.com/tkennes/open-azure-emissions/pkg/util/parsers"

	azure_constants "github.com/tkennes/open-azure-emissions/pkg/azure/constants"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func NetworkEnergyCalculation(gigabytes float64, energyCoefficient float64) (float64, error) {
	return gigabytes * energyCoefficient, nil
}

func EstimateNetworkingEnergyConsumption(data azure_models.AzureCostDetails) (float64, error) {
	gigabyteMeasure, err := parsers.ParseGigabyteMeasure(data.UnitOfMeasure)
	if err != nil {
		return 0, err
	}
	return NetworkEnergyCalculation(
		data.Quantity*gigabyteMeasure,
		azure_constants.NETWORKING_ENERGY_COEFFICIENT,
	)
}

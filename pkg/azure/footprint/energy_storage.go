package azure_footprint

import (
	"errors"
	"fmt"

	"github.com/tkennes/open-azure-emissions/pkg/util/parsers"

	azure_constants "github.com/tkennes/open-azure-emissions/pkg/azure/constants"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func StorageEnergyCalculation(
	teraByteHours float64,
	StorageCoefficient float64,
	replicationFactor int,
) float64 {
	return teraByteHours * StorageCoefficient * float64(replicationFactor) / 1000
}

func EstimateStorageEnergy(data azure_models.AzureCostDetails, diskType string, terabytes float64, replicas int) (float64, error) {
	hours, err := parsers.ParseHours(data)
	if err != nil {
		return 0, err
	}

	var storageCoefficient float64
	if diskType == "SSD" {
		storageCoefficient = azure_constants.SSDCOEFFICIENT
	} else if diskType == "HDD" {
		storageCoefficient = azure_constants.HDDCOEFFICIENT
	} else {
		return 0, errors.New(fmt.Sprintf("Disk Type \"%s\" not found (expected SSD or HDD)", diskType))
	}

	return StorageEnergyCalculation(
		calculateTerabyteHours(terabytes, data.Quantity, hours),
		storageCoefficient,
		replicas,
	), nil
}

func calculateTerabyteHours(terabytes float64, quantity float64, hours float64) float64 {
	return terabytes * quantity * hours
}

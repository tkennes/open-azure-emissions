package azure_footprint_core

import (
	"errors"
	"fmt"

	azure_constants "github.com/tkennes/open-azure-emissions/pkg/azure/constants"
)

func EstimateFootprint(
	kilowattHours float64,
	region string,
) (float64, error) {
	if emissions_factor, found := azure_constants.AZURE_REGION_EMISSIONS_FACTOR_TON_PER_KWH[region]; found {
		return kilowattHours * emissions_factor, nil
	} else {
		return 0, errors.New(fmt.Sprintf("Region \"%s\" not found", region))
	}
}

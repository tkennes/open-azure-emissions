package azure_footprint

import (
	"errors"
	"fmt"

	"github.com/tkennes/open-azure-emissions/pkg/util/parsers"

	azure_constants_compute_processors "github.com/tkennes/open-azure-emissions/pkg/azure/constants/compute_processors"
	azure_constants_services "github.com/tkennes/open-azure-emissions/pkg/azure/constants/services"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func EstimateAppServiceComputeEnergyConsumption(data azure_models.AzureCostDetails) (float64, error) {
	machine := parsers.ParseAppServiceMachine(data.MeterName)
	if appServicePlan, found := azure_constants_services.APP_SERVICE_PLANS[machine]; found {
		if instanceTypes, found := azure_constants_compute_processors.INSTANCE_TYPE_COMPUTE_PROCESSOR_MAPPING[appServicePlan.Name]; found {
			return EstimateComputeEnergy(
				instanceTypes,
				data,
				appServicePlan.VCPUs,
			)
		} else {
			return 0, errors.New(fmt.Sprintf("App Service Plan Product \"%s\" (machine: %s): could not find instance type", appServicePlan.Name, machine))
		}
	} else {
		return 0, errors.New(fmt.Sprintf("App Service Plan Product \"%s\" (machine: %s): could not parse/find machine", data.Product, machine))
	}
}

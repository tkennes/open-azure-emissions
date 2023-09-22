package azure_footprint_core

import (
	"errors"
	"fmt"

	"github.com/tkennes/open-azure-emissions/pkg/util/parsers"

	azure_constants "github.com/tkennes/open-azure-emissions/pkg/azure/constants"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func ComputeEnergyCalculation(
	averageCPUUtilization float64,
	virtualCPUHours float64,
	minWatts float64,
	maxWatts float64,
) (float64, error) {
	if averageCPUUtilization > 1 {
		fmt.Errorf("averageCPUUtilization must be between 0 and 1, since it reflects a percentage")
		return 0, errors.New("averageCPUUtilization must be between 0 and 1, since it reflects a percentage")
	}
	averageWatts := minWatts + (averageCPUUtilization/100)*(maxWatts-minWatts)
	kilowattHours := averageWatts * virtualCPUHours / 1000
	return kilowattHours, nil
}

func EstimateComputeEnergy(instanceTypes []string, data azure_models.AzureCostDetails, vcpus int) (float64, error) {
	min_watts, err := GetMinWatts(instanceTypes)
	if err != nil {
		return 0, err
	}
	max_watts, err := GetMaxWatts(instanceTypes)
	if err != nil {
		return 0, err
	}
	hours, err := parsers.ParseHours(data)
	if err != nil {
		return 0, err
	}
	return ComputeEnergyCalculation(
		float64(azure_constants.AVG_CPU_UTILIZATION_2020),
		calculateVCPUHours(vcpus, data.Quantity, hours),
		min_watts,
		max_watts,
	)
}

func calculateVCPUHours(vcpus int, quantity float64, hours float64) float64 {
	return float64(vcpus) * quantity * hours
}

func GetMinWatts(instanceTypes []string) (float64, error) {
	var totalMinWatts float64 = 0
	for _, instance_type := range instanceTypes {
		if _, found := azure_constants.MIN_WATTS_BY_COMPUTE_PROCESSOR[instance_type]; found {
			totalMinWatts = totalMinWatts + azure_constants.MIN_WATTS_BY_COMPUTE_PROCESSOR[instance_type]
		} else {
			return 0, errors.New(fmt.Sprintf("Instance type \"%s\" not found in MIN_WATTS_BY_COMPUTE_PROCESSOR", instance_type))
		}
	}
	return totalMinWatts / float64(len(instanceTypes)), nil
}

func GetMaxWatts(instanceTypes []string) (float64, error) {
	var totalMaxWatts float64 = 0
	for _, instance_type := range instanceTypes {
		if _, found := azure_constants.MAX_WATTS_BY_COMPUTE_PROCESSOR[instance_type]; found {
			totalMaxWatts = totalMaxWatts + azure_constants.MAX_WATTS_BY_COMPUTE_PROCESSOR[instance_type]
		} else {
			return 0, errors.New(fmt.Sprintf("Instance type \"%s\" not found in MAX_WATTS_BY_COMPUTE_PROCESSOR", instance_type))
		}
	}
	return totalMaxWatts / float64(len(instanceTypes)), nil
}

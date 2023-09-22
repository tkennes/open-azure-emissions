package azure_footprint

import (
	"errors"
	"fmt"

	"github.com/tkennes/open-azure-emissions/pkg/util/parsers"

	azure_constant_compute_processors "github.com/tkennes/open-azure-emissions/pkg/azure/constants/compute_processors"
	azure_constant_services "github.com/tkennes/open-azure-emissions/pkg/azure/constants/services"
	azure_footprint_core "github.com/tkennes/open-azure-emissions/pkg/azure/footprint/core"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func EstimateVirtualMachineEnergyConsumption(data azure_models.AzureCostDetails) (float64, error) {
	machine := parsers.ParseVirtualmachineMachine(data.MeterName)
	if virtualMachine, found := azure_constant_services.VIRTUAL_MACHINE_TYPES[machine]; found {
		if instanceTypes, found := azure_constant_compute_processors.INSTANCE_TYPE_COMPUTE_PROCESSOR_MAPPING[virtualMachine.Name]; found {
			computeEnergy, err := azure_footprint_core.EstimateComputeEnergy(
				instanceTypes,
				data,
				virtualMachine.VCPUs,
			)
			if err != nil {
				return 0, err
			}
			replicas := 1
			memoryEnergy, err := azure_footprint_core.EstimateMemoryEnergy(
				data,
				virtualMachine.Memory,
				replicas,
			)
			if err != nil {
				return 0, err
			}
			return computeEnergy + memoryEnergy, nil
		} else {
			return 0, errors.New(fmt.Sprintf("Virtual Machine: could not find instance type \"%s\"", machine))
		}
	} else {
		return 0, errors.New(fmt.Sprintf("Virtual Machine: could not find machine \"%s\"", machine))
	}
}

package azure_footprint

import (
	"testing"

	azure_footprint_core "github.com/tkennes/open-azure-emissions/pkg/azure/footprint/core"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
	test_helpers "github.com/tkennes/open-azure-emissions/pkg/util/test_helpers"
)

func TestEstimateVirtualMachineEnergyConsumption(t *testing.T) {

	var testcases = []azure_models.AzureCostDetailsTestCase{
		azure_models.AzureCostDetailsTestCase{
			Data: azure_models.AzureCostDetails{
				MeterName:        "D4ds v4",
				MeterCategory:    "Virtual Machines",
				MeterSubCategory: "Virtual Machines Ddsv4 Series",
				Product:          "Virtual Machines Ddsv4 Series - D4ds v4",
				Quantity:         24,
				UnitOfMeasure:    "1 Hour",
			},
			ExpectedFloat: 0.211968,
		},
		azure_models.AzureCostDetailsTestCase{
			Data: azure_models.AzureCostDetails{
				MeterName:        "F8",
				MeterCategory:    "Virtual Machines",
				MeterSubCategory: "Virtual Machines F Series",
				Product:          "Virtual Machines Ddsv4 Series - A4 v2",
				Quantity:         21,
				UnitOfMeasure:    "1 Hour",
			},
			ExpectedFloat: 0.257712,
		},
	}

	for _, testcase := range testcases {
		result, err := azure_footprint_core.EstimateVirtualMachineEnergyConsumption(testcase.Data)
		if err != nil {
			t.Errorf("TestEstimateVirtualMachineEnergyConsumption Error: %s", err)
		} else if !test_helpers.FloatsAlmostEqual(result, testcase.ExpectedFloat) {
			t.Errorf("TestEstimateVirtualMachineEnergyConsumption: %f, got %f", testcase.ExpectedFloat, result)
		}
	}
}

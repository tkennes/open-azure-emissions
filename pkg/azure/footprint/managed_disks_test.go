package azure_footprint

import (
	"testing"

	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
	test_helpers "github.com/tkennes/open-azure-emissions/pkg/util/test_helpers"
)

func TestEstimateManagedDiskEnergyConsumption(t *testing.T) {
	var testcases = []azure_models.AzureCostDetailsTestCase{
		azure_models.AzureCostDetailsTestCase{
			Data: azure_models.AzureCostDetails{
				Date:             "2023-01-31",
				MeterName:        "P15 LRS Disk",
				MeterCategory:    "Storage",
				MeterSubCategory: "Premium SSD Managed Disks",
				Quantity:         0.002688,
				UnitOfMeasure:    "1/Month",
			},
			ExpectedFloat: 0.0018430820351999996,
		},
		azure_models.AzureCostDetailsTestCase{
			Data: azure_models.AzureCostDetails{
				Date:             "2023-02-01",
				MeterName:        "E6 Disks",
				MeterCategory:    "Storage",
				MeterSubCategory: "Standard SSD Managed Disks",
				Quantity:         0.035712,
				UnitOfMeasure:    "1/Month",
			},
			ExpectedFloat: 0.0018430820351999998,
		},
		azure_models.AzureCostDetailsTestCase{
			Data: azure_models.AzureCostDetails{
				Date:             "2023-05-01",
				MeterName:        "S10 Disks",
				MeterCategory:    "Storage",
				MeterSubCategory: "Standard HDD Managed Disks",
				Quantity:         0.032256,
				UnitOfMeasure:    "1/Month",
			},
			ExpectedFloat: 0.0019966722048,
		},
	}

	for _, testcase := range testcases {
		result, err := EstimateManagedDiskEnergyConsumption(testcase.Data)
		if err != nil {
			t.Errorf("Error: %s", err)
		} else if !test_helpers.FloatsAlmostEqual(result, testcase.ExpectedFloat) {
			t.Errorf("Expected: %f, got %f", testcase.ExpectedFloat, result)
		}
	}
}

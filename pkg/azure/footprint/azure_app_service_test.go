package azure_footprint

import (
	"testing"

	azure_footprint_core "github.com/tkennes/open-azure-emissions/pkg/azure/footprint/core"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
	test_helpers "github.com/tkennes/open-azure-emissions/pkg/util/test_helpers"
)

func TestEstimateAppServiceComputeEnergyConsumption(t *testing.T) {
	var testcases = []azure_models.AzureCostDetailsTestCase{
		azure_models.AzureCostDetailsTestCase{
			Data: azure_models.AzureCostDetails{
				MeterName:        "F1 App",
				MeterCategory:    "Azure App Service",
				MeterSubCategory: "Azure App Service Free Plan - Linux",
				Product:          "Azure App Service Free Plan - Linux - F1",
				UnitOfMeasure:    "1 Hour",
				Quantity:         0.032256000000000014,
			},
			ExpectedFloat: 2.419200000000001e-05,
		},
		azure_models.AzureCostDetailsTestCase{
			Data: azure_models.AzureCostDetails{
				MeterName:        "F1 App",
				MeterCategory:    "Azure App Service",
				MeterSubCategory: "Azure App Service Free Plan",
				Product:          "Azure App Service Free Plan - F1",
				UnitOfMeasure:    "1 Hour",
				Quantity:         0.032256000000000014,
			},
			ExpectedFloat: 2.419200000000001e-05,
		},
	}

	for _, testcase := range testcases {
		result, err := azure_footprint_core.EstimateAppServiceComputeEnergyConsumption(testcase.Data)
		if err != nil {
			t.Errorf("Error: %s", err)
		} else if !test_helpers.FloatsAlmostEqual(result, testcase.ExpectedFloat) {
			t.Errorf("Expected: %f, got %f", testcase.ExpectedFloat, result)
		}
	}
}

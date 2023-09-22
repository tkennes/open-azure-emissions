package azure_footprint

import (
	"testing"

	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
	test_helpers "github.com/tkennes/open-azure-emissions/pkg/util/test_helpers"
)

func TestEstimateRedisCacheEnergyConsumption(t *testing.T) {

	var testcases = []azure_models.AzureCostDetailsTestCase{
		azure_models.AzureCostDetailsTestCase{
			Data: azure_models.AzureCostDetails{
				MeterName:        "C1 Cache Instance",
				MeterCategory:    "Redis Cache",
				MeterSubCategory: "Azure Redis Cache Standard",
				Product:          "Azure Redis Cache Standard - C1 - EU West",
				UnitOfMeasure:    "1 Hour",
				Quantity:         48,
				Location:         "EU West",
			},
			ExpectedFloat: 0.037632,
		},
		azure_models.AzureCostDetailsTestCase{
			Data: azure_models.AzureCostDetails{
				MeterName:        "P1 Cache Instance",
				MeterCategory:    "Redis Cache",
				MeterSubCategory: "Azure Redis Cache Premium",
				Product:          "Azure Redis Cache Premium - P1 - EU West",
				UnitOfMeasure:    "1 Hour",
				Quantity:         48,
				Location:         "EU West",
			},
			ExpectedFloat: 0.338688,
		},
	}

	for _, testcase := range testcases {
		result, err := EstimateRedisCacheEnergyConsumption(testcase.Data)
		if err != nil {
			t.Errorf("TestEstimateRedisCacheEnergyConsumption Error: %s", err)
		} else if !test_helpers.FloatsAlmostEqual(result, testcase.ExpectedFloat) {
			t.Errorf("TestEstimateRedisCacheEnergyConsumption: %f, got %f", testcase.ExpectedFloat, result)
		}
	}
}

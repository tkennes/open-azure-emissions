package azure

import (
	"errors"
	"fmt"

	"golang.org/x/exp/slices"

	"github.com/tkennes/open-azure-emissions/pkg/log"
	"github.com/tkennes/open-azure-emissions/pkg/util/stringutil"

	azure_constants_meters "github.com/tkennes/open-azure-emissions/pkg/azure/constants/meters"
	azure_footprint "github.com/tkennes/open-azure-emissions/pkg/azure/footprint"
	azure_footprint_core "github.com/tkennes/open-azure-emissions/pkg/azure/footprint/core"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func AssessFootprint(data azure_models.AzureCostDetails) (float64, error) {
	if IsMeterCategoryFootprintEstimable(data.MeterCategory) {
		switch {
		case data.MeterCategory == "Azure App Service":
			if !slices.Contains(azure_constants_meters.AZURE_APP_SERVICE_METER_SUBCATEGORIES_WITHOUT_COMPUTE, data.MeterSubCategory) {
				// Ignoring the unknown Azure App Services
				return azure_footprint.EstimateAppServiceComputeEnergyConsumption(data)
			} else {
				// If unknown, return 0 and an error
				message := fmt.Sprintf("Meter category \"%s\": subcategory  \"%s\" not estimable", data.MeterCategory, data.MeterSubCategory)
				log.Infof(message)
				return 0, errors.New(message)
			}
		case data.MeterCategory == "Virtual Machines":
			if data.MeterName != "" {
				return azure_footprint.EstimateVirtualMachineEnergyConsumption(data)
			} else {
				message := fmt.Sprintf("Meter category \"%s\": Meter name is empty \"%s\", probably because \"%s\" contains reservations", data.MeterCategory, data.MeterName, data.Product)
				log.Infof(message)
				return 0, errors.New(message)
			}
		case data.MeterCategory == "Redis Cache":
			return azure_footprint.EstimateRedisCacheEnergyConsumption(data)
		case data.MeterCategory == "Storage":
			if slices.Contains(azure_constants_meters.ESTIMABLE_STORAGE_SUBCATEGORIES, data.MeterSubCategory) {
				if !slices.Contains(azure_constants_meters.UNESTIMABLE_STORAGE_METER_NAMES, data.MeterName) {
					return azure_footprint.EstimateManagedDiskEnergyConsumption(data)
				} else {
					message := fmt.Sprintf("Meter category \"%s\": Meter name \"%s\" not estimable", data.MeterCategory, data.MeterName)
					log.Infof(message)
					return 0, errors.New(message)
				}
			} else {
				message := fmt.Sprintf("Meter category \"%s\": Meter subcategory \"%s\" not estimable", data.MeterCategory, data.MeterSubCategory)
				log.Infof(message)
				return 0, errors.New(message)
			}
		case stringutil.Contains(data.MeterCategory, azure_constants_meters.ESTIMABLE_NETWORKING_CATEGORIES):
			if data.UnitOfMeasure == "1 GB" {
				return azure_footprint_core.EstimateNetworkingEnergyConsumption(data)
			} else {
				message := fmt.Sprintf("Meter category \"%s\": Meter UnitOfMeasure \"%s\" not estimable, product: \"%s\"", data.MeterCategory, data.UnitOfMeasure, data.Product)
				log.Infof(message)
				return 0, errors.New(message)
			}
		default:
			// If unknown, return 0 and an error
			message := fmt.Sprintf("Meter category \"%s\" not estimable", data.MeterCategory)
			log.Infof(message)
			return 0, errors.New(message)
		}
	} else {
		// Ignoring the unknown Meter Categories
		return 0, errors.New(fmt.Sprintf("Meter category \"%s\" not estimable", data.MeterCategory))
	}
}

func IsMeterCategoryFootprintEstimable(meterCategory string) bool {
	return slices.Contains(azure_constants_meters.ESTIMABLE_METER_CATEGORIES, meterCategory)
}

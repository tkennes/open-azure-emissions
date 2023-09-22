package azure

import (
	"errors"
	"fmt"

	"golang.org/x/exp/slices"

	"github.com/tkennes/open-azure-emissions/pkg/log"
	"github.com/tkennes/open-azure-emissions/pkg/util/stringutil"

	azure_constants_meters "github.com/tkennes/open-azure-emissions/pkg/azure/constants/meters"
	azure_footprint "github.com/tkennes/open-azure-emissions/pkg/azure/footprint"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func AssessFootprint(data azure_models.AzureCostDetails) (float64, error) {
	if IsMeterCategoryFootprintEstimable(data.MeterCategory) {
		switch meterCategory := data.MeterCategory; meterCategory {
		case "Azure App Service":
			if !slices.Contains(azure_constants_meters.AZURE_APP_SERVICE_METER_SUBCATEGORIES_WITHOUT_COMPUTE, data.MeterSubCategory) {
				// Ignoring the unknown Azure App Services
				return azure_footprint.EstimateAppServiceComputeEnergyConsumption(data)
			} else {
				// If unknown, return 0 and an error
				message := fmt.Sprintf("Meter category \"%s\": subcategory  \"%s\" not estimable", data.MeterCategory, data.MeterSubCategory)
				log.Infof(message)
				return 0, errors.New(message)
			}
		case "Virtual Machines":
			if data.MeterName != "" {
				return azure_footprint.EstimateVirtualMachineEnergyConsumption(data)
			} else {
				message := fmt.Sprintf("Meter category \"%s\": Meter name is empty \"%s\", probably because \"%s\" contains reservations", data.MeterCategory, data.Product)
				log.Infof(message)
				return 0, errors.New(message)
			}
		case "Redis Cache":
			return azure_footprint.EstimateRedisCacheEnergyConsumption(data)
		case "Storage":
			if slices.Contains(azure_constants_meters.ESTIMABLE_STORAGE_SUBCATEGORIES, data.MeterSubCategory) {
				if !slices.Contains(azure_constants_meters.UNESTIMABLE_STORAGE_METER_NAMES, data.MeterName) {
					return azure_footprint.EstimateManagedDiskEnergyConsumption(data)
				} else {
					message := fmt.Sprintf("Meter category \"%s\": Meter name \"%s\" not estimable", data.MeterName)
					log.Infof(message)
					return 0, errors.New(message)
				}
			} else {
				message := fmt.Sprintf("Meter category \"%s\": Meter subcategory \"%s\" not estimable", data.MeterSubCategory)
				log.Infof(message)
				return 0, errors.New(message)
			}
		case stringutil.Contains(meterCategory, azure_constants_meters.ESTIMABLE_NETWORKING_CATEGORIES):
			if data.UnitOfMeasure == "1 GB" {
				return azure_footprint.EstimateNetworkingEnergyConsumption(data)
			} else {
				message := fmt.Sprintf("Meter category \"%s\": Meter UnitOfMeasure \"%s\" not estimable, product: \"%s\"", data.meterCategory, data.UnifOfMeasure, data.Product)
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

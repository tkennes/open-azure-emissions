package azure

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/tkennes/open-azure-emissions/pkg/log"
	"github.com/tkennes/open-azure-emissions/pkg/util/stringutil"

	azure_constants_meters "github.com/tkennes/open-azure-emissions/pkg/azure/constants/meters"
	azure_footprint "github.com/tkennes/open-azure-emissions/pkg/azure/footprint"
	azure_footprint_core "github.com/tkennes/open-azure-emissions/pkg/azure/footprint/core"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func AssessFootprint(data azure_models.AzureCostDetails) (float64, error) {
	switch {
	// App Service
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

	// Container Registry
	case data.MeterCategory == "Container Registry":
		if slices.Contains(azure_constants_meters.CONTAINER_REGISTRY_STORAGE_METER_NAMES, data.MeterName) {
			return azure_footprint.EstimateContainerRegistryEnergyConsumption(data)
		} else {
			message := fmt.Sprintf("Meter category \"%s\": Meter name \"%s\" not estimable", data.MeterCategory, data.MeterName)
			log.Infof(message)
			return 0, errors.New(message)
		}

	// Networking
	case stringutil.Contains(data.MeterCategory, azure_constants_meters.ESTIMABLE_NETWORKING_CATEGORIES):
		if data.UnitOfMeasure == "1 GB" {
			return azure_footprint_core.EstimateNetworkingEnergyConsumption(data)
		} else {
			message := fmt.Sprintf("Meter category \"%s\": Meter UnitOfMeasure \"%s\" not estimable, product: \"%s\"", data.MeterCategory, data.UnitOfMeasure, data.Product)
			log.Infof(message)
			return 0, errors.New(message)
		}

	// Redis Cache
	case data.MeterCategory == "Redis Cache":
		return azure_footprint.EstimateRedisCacheEnergyConsumption(data)

	// Storage
	case data.MeterCategory == "Storage":
		if storage_usage, found := azure_constants_meters.ESTIMABLE_STORAGE_USAGE[data.MeterSubCategory]; found {
			// Handle Disk Types
			if storage_usage.StorageType == "Disk" {
				if !slices.Contains(azure_constants_meters.UNESTIMABLE_DISK_STORAGE_METER_NAMES, data.MeterName) {
					return azure_footprint.EstimateManagedDiskEnergyConsumption(data)
				} else {
					message := fmt.Sprintf("Meter category-subcategory \"%s\"-\"%s\": Meter name \"%s\" not estimable", data.MeterCategory, data.MeterSubCategory, data.MeterName)
					log.Infof(message)
					return 0, errors.New(message)
				}

				// Handle Blob Types
			} else if storage_usage.StorageType == "Blob" {
				if strings.Contains(data.MeterName, storage_usage.MeterNameHas) {
					return azure_footprint_core.EstimateBlobEnergyConsumption(data)
				} else {
					message := fmt.Sprintf("Meter category-subcategory \"%s\"-\"%s\": Meter name \"%s\" not estimable", data.MeterCategory, data.MeterSubCategory, data.MeterName)
					log.Infof(message)
					return 0, errors.New(message)
				}
			} else {
				message := fmt.Sprintf("Meter category \"%s\"-\"%s\": Custom Storage type \"%s\" not estimable", data.MeterCategory, data.MeterSubCategory, storage_usage.StorageType)
				log.Infof(message)
				return 0, errors.New(message)
			}
		} else {
			message := fmt.Sprintf("Meter category \"%s\": Meter subcategory \"%s\" not estimable", data.MeterCategory, data.MeterSubCategory)
			log.Infof(message)
			return 0, errors.New(message)
		}

	// Virtual Machines
	case data.MeterCategory == "Virtual Machines":
		if data.MeterName != "" {
			return azure_footprint.EstimateVirtualMachineEnergyConsumption(data)
		} else {
			message := fmt.Sprintf("Meter category \"%s\": Meter name is empty \"%s\", probably because \"%s\" contains reservations", data.MeterCategory, data.MeterName, data.Product)
			log.Infof(message)
			return 0, errors.New(message)
		}
	default:
		// If unknown, return 0 and an error
		message := fmt.Sprintf("Meter category \"%s\" not estimable", data.MeterCategory)
		log.Infof(message)
		return 0, errors.New(message)
	}
}

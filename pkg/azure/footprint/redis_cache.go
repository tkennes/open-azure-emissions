package azure_footprint

import (
	"errors"
	"fmt"

	azure_constant_services "github.com/tkennes/open-azure-emissions/pkg/azure/constants/services"
	azure_footprint_core "github.com/tkennes/open-azure-emissions/pkg/azure/footprint/core"
	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
	"github.com/tkennes/open-azure-emissions/pkg/util/parsers"
)

func EstimateRedisCacheEnergyConsumption(data azure_models.AzureCostDetails) (float64, error) {
	instance := parsers.ParseRedisCacheInstance(data.MeterName)
	if redisCache, found := azure_constant_services.REDIS_CACHES[instance]; found {
		if redisCache.Series == "" {
			series, err := parsers.ParseRedisCacheSeries(data.MeterSubCategory)
			redisCache.Series = series
			if err != nil {
				return 0, err
			}
		}
		return azure_footprint_core.EstimateMemoryEnergy(
			data,
			redisCache.Size,
			azure_constant_services.REDIS_CACHE_REPLICAS[redisCache.Series],
		)
	} else {
		return 0, errors.New(fmt.Sprintf("Redis Cache (meterName: \"%s\", instance: %s): was not found", data.MeterName, instance))
	}
}

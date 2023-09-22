package azure

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/tkennes/open-azure-emissions/pkg/log"

	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func GetFootprint(raw_data []byte) (interface{}, error) {
	var azureCostDetails azure_models.AzureCostDetails
	_ = json.Unmarshal(raw_data, &azureCostDetails)

	validate := validator.New()
	err := validate.Struct(azureCostDetails)
	if err != nil {
		log.Infof("Error validating AzureCostDetails: %s", err)
		return nil, err
	}
	footprint, err := AssessFootprint(azureCostDetails)
	if err != nil {
		log.Infof("Error assessing footprint: %s", err)
		return nil, err
	}
	return footprint, nil
}

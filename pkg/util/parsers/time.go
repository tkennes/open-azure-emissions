package parsers

import (
	"errors"
	"fmt"
	"time"

	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func ParseHours(data azure_models.AzureCostDetails) (float64, error) {
	if data.UnitOfMeasure == "1 Hour" {
		return 1, nil
	} else if data.UnitOfMeasure == "1/Month" || data.UnitOfMeasure == "1 GB/Month" {
		date, err := time.Parse("2006-01-02", data.Date)
		if err != nil {
			return 0, err
		}
		return float64(getHoursInMonth(date.Month(), date.Year())), nil
	} else {
		return 0, errors.New(fmt.Sprintf("Could not parse hours: %s", data.UnitOfMeasure))
	}
}

func getHoursInMonth(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day() * 24
}

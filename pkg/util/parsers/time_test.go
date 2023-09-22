package parsers

import (
	"strconv"
	"testing"

	azure_models "github.com/tkennes/open-azure-emissions/pkg/azure/models"
)

func TestParseHours(t *testing.T) {
	data := azure_models.AzureCostDetails{
		UnitOfMeasure: "",
		Date:          "2020-01-01",
	}

	result, _ := ParseHours(data)
	if result != 0 {
		t.Errorf("Result 1 was incorrect. Expected 0, got %s", strconv.FormatFloat(result, 'f', -1, 64))
	}

	data.UnitOfMeasure = "1 Hour"
	result_2, err := ParseHours(data)
	if result_2 != 1 {
		t.Errorf("Result 2 was incorrect. Expected 1, got %s: error %s", strconv.FormatFloat(result, 'f', -1, 64), err)
	}

	data.UnitOfMeasure = "1/Month"
	result_3, err := ParseHours(data)
	if result_3 != 744 {
		t.Errorf("Result 3 was incorrect. Expected 1, got %s: error: %s", strconv.FormatFloat(result, 'f', -1, 64), err)
	}
}

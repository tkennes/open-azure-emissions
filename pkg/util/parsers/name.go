package parsers

import (
	"errors"
	"fmt"
	"strings"
)

func ParseAppServiceMachine(meterName string) string {
	return strings.Replace(meterName, " App", "", 1)
}

func ParseRedisCacheInstance(meterName string) string {
	name := strings.Replace(meterName, " Instance", "", 1)
	return strings.Replace(name, " Cache", "", 1)
}

func ParseRedisCacheSeries(meterSubCategory string) (string, error) {
	if strings.Contains(meterSubCategory, "Standard") {
		return "Standard", nil
	} else if strings.Contains(meterSubCategory, "Basic") {
		return "Basic", nil
	} else {
		return "", errors.New(fmt.Sprintf("Redis Cache (meterSubCategory: \"%s\"): was not detected as Basic or Standard", meterSubCategory))
	}
}

func ParseGigabyteMeasure(unitOfMeasure string) (float64, error) {
	if unitOfMeasure == "1 GB" {
		return 1, nil
	} else {
		return 0, errors.New(fmt.Sprintf("Unit of measure \"%s\" not estimable", unitOfMeasure))
	}
}

func ParseManagedDiskType(meterName string) string {
	// There are several options we are tackling here:
	// Mostly one of the following three formats
	// E1 LRS Disks - Free
	// S10 GRS Disks
	// D10 Disks

	name := strings.Replace(meterName, " - Free", "", 1)
	name = strings.Replace(name, " Disks", "", 1)
	name = strings.Replace(name, " Disk", "", 1)

	name = strings.Replace(name, " LRS", "", 1)
	name = strings.Replace(name, " ZRS", "", 1)
	name = strings.Replace(name, " RA-GRS", "", 1)
	name = strings.Replace(name, " RA-GZRS", "", 1)
	name = strings.Replace(name, " GRS", "", 1)
	name = strings.Replace(name, " GZRS", "", 1)

	return name
}

func ParseManagedDiskRedundancy(meterName string) string {
	meterNameLower := strings.ToLower(meterName)
	// LRS
	if strings.Contains(meterNameLower, "lrs") {
		return "LRS"

		// ZRS
	} else if strings.Contains(meterNameLower, "ra-gzrs") {
		return "RA-GZRS"
	} else if strings.Contains(meterNameLower, "gzrs") {
		return "GZRS"
	} else if strings.Contains(meterNameLower, "zrs") {
		return "ZRS"

		// GRS
	} else if strings.Contains(meterNameLower, "ra-grs") {
		return "RA-GRS"
	} else if strings.Contains(meterNameLower, "grs") {
		return "GRS"
	} else {
		return ""
	}

}

func ParseVirtualmachineMachine(meterName string) string {
	// Virtual Machine Meter Names can have the structure: D2 v3/D2s v3, or D2 v3
	// In both cases, we only want the first part.
	// And we remove " Spot" from the meter name.
	return strings.Replace(strings.SplitN(meterName, "/", 2)[0], " Spot", "", 1)
}

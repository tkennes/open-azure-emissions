package parsers

import (
	"testing"
)

func TestParseAppServiceMachine(t *testing.T) {
	result := ParseAppServiceMachine("S1 App")
	if result != "S1" {
		t.Errorf("Result was incorrect. Expected S1, got %s", result)
	}
}

func TestParseRedisCacheInstance(t *testing.T) {
	result := ParseRedisCacheInstance("C0 Instance")
	if result != "C0" {
		t.Errorf("Result 1 was incorrect. Expected C0, got %s", result)
	}
	result_2 := ParseRedisCacheInstance("C0 Cache")
	if result_2 != "C0" {
		t.Errorf("Result 2 was incorrect. Expected C0, got %s", result_2)
	}
	result_3 := ParseRedisCacheInstance("C0 Instance Cache")
	if result_3 != "C0" {
		t.Errorf("Result 3 was incorrect. Expected C0, got %s", result_3)
	}
}

func TestParseRedisCacheSeries(t *testing.T) {
	result, _ := ParseRedisCacheSeries("Basic C0 Cache")
	if result != "Basic" {
		t.Errorf("Result 1 was incorrect. Expected Basic, got %s", result)
	}
	result_2, _ := ParseRedisCacheSeries("Standard C0 Cache")
	if result_2 != "Standard" {
		t.Errorf("Result 2 was incorrect. Expected Standard, got %s", result_2)
	}
	result_3, _ := ParseRedisCacheSeries("C0 Cache")
	if result_3 != "" {
		t.Errorf("Result 3 was incorrect. Expected \"\", got %s", result_3)
	}
}

func TestParseManagedDiskType(t *testing.T) {
	result := ParseManagedDiskType("E1 LRS Disks - Free")
	if result != "E1" {
		t.Errorf("Result 1 was incorrect. Expected E1, got %s", result)
	}
	result_2 := ParseManagedDiskType("S10 GRS Disks")
	if result_2 != "S10" {
		t.Errorf("Result 2 was incorrect. Expected S10, got %s", result_2)
	}
	result_3 := ParseManagedDiskType("D10 Disks")
	if result_3 != "D10" {
		t.Errorf("Result 3 was incorrect. Expected D10, got %s", result_3)
	}
	result_4 := ParseManagedDiskType("P15 LRS Disk")
	if result_4 != "P15" {
		t.Errorf("Result 4 was incorrect. Expected P15, got %s", result_4)
	}
}

func TestParseManagedDiskRedundancy(t *testing.T) {
	result := ParseManagedDiskRedundancy("E1 LRS Disks - Free")
	if result != "LRS" {
		t.Errorf("Result 1 was incorrect. Expected LRS, got %s", result)
	}
	result_2 := ParseManagedDiskRedundancy("S10 GRS Disks")
	if result_2 != "GRS" {
		t.Errorf("Result 2 was incorrect. Expected GRS, got %s", result_2)
	}
	result_3 := ParseManagedDiskRedundancy("D10 Disks")
	if result_3 != "" {
		t.Errorf("Result 3 was incorrect. Expected \"\", got %s", result_3)
	}
	result_4 := ParseManagedDiskRedundancy("D10 RA-GZRS Disks")
	if result_4 != "RA-GZRS" {
		t.Errorf("Result 4 was incorrect. Expected RA-GZRS, got %s", result_4)
	}
}

func TestParseVirtualmachineMachine(t *testing.T) {
	result := ParseVirtualmachineMachine("Standard_D2s_v3")
	if result != "Standard_D2s_v3" {
		t.Errorf("Result 1 was incorrect. Expected Standard_D2s_v3, got %s", result)
	}
	result_2 := ParseVirtualmachineMachine("D2 v3/D2s v3")
	if result_2 != "D2 v3" {
		t.Errorf("Result 2 was incorrect. Expected D2 v3, got %s", result_2)
	}
	result_3 := ParseVirtualmachineMachine("D2 v3 Spot")
	if result_3 != "D2 v3" {
		t.Errorf("Result 3 was incorrect. Expected D2 v3, got %s", result_3)
	}
}

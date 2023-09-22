package azure_unit_of_measure

var (
	MEMORY_METER_CATEGORIES = []string{
		"Redis Cache",
	}

	// [cacheName: string]: number
	CACHE_MEMORY_GB = map[string]float64{
		"C0":    0.25,
		"C1":    1,
		"C2":    2.5,
		"C3":    6,
		"C4":    13,
		"C5":    26,
		"C6":    53,
		"P1":    6,
		"P2":    13,
		"P3":    26,
		"P4":    53,
		"P5":    120,
		"E10":   12,
		"E20":   25,
		"E50":   50,
		"E100":  100,
		"F300":  384,
		"F700":  715,
		"F1500": 1455,
	}
)

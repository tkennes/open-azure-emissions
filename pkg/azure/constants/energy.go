package azure_constants

var (
	// STORAGE
	SSDCOEFFICIENT = 1.2  // watt hours / terabyte hour
	HDDCOEFFICIENT = 0.65 // watt hours / terabyte hour

	// MEMORY
	MEMORY_AVG = 80.47

	// Gigabytes / physical Chip
	MEMORY_BY_COMPUTE_PROCESSOR = map[string]float64{
		"CASCADE_LAKE":     98.12,
		"SKYLAKE":          81.32,
		"BROADWELL":        69.65,
		"HASWELL":          27.71,
		"COFFEE_LAKE":      19.56,
		"SANDY_BRIDGE":     16.7,
		"IVY_BRIDGE":       9.67,
		"AMD_EPYC_1ST_GEN": 89.6,
		"AMD_EPYC_2ND_GEN": 129.78,
		"AMD_EPYC_3RD_GEN": 128.0,
	}

	// COMPUTE
	MIN_WATTS_AVG                  = 0.74
	MIN_WATTS_BY_COMPUTE_PROCESSOR = map[string]float64{
		// CPUs
		"CASCADE_LAKE":     0.64,
		"SKYLAKE":          0.65,
		"BROADWELL":        0.71,
		"HASWELL":          1,
		"COFFEE_LAKE":      1.14,
		"SANDY_BRIDGE":     2.17,
		"IVY_BRIDGE":       3.04,
		"AMD_EPYC_1ST_GEN": 0.82,
		"AMD_EPYC_2ND_GEN": 0.47,
		"AMD_EPYC_3RD_GEN": 0.45,
		// GPUs
		"NVIDIA_T4":         8,
		"NVIDIA_TESLA_K80":  35,
		"NVIDIA_TESLA_P100": 36,
		"NVIDIA_TESLA_V100": 35,
		"NVIDIA_TESLA_M60":  35,
		"NVIDIA_TESLA_P40":  30,
		"NVIDIA_TESLA_A100": 46,
		"XILINX_ALVEO_U250": 27,
	}

	MAX_WATTS_AVG                  = 3.54
	MAX_WATTS_BY_COMPUTE_PROCESSOR = map[string]float64{
		// CPUs
		"CASCADE_LAKE":     3.97,
		"SKYLAKE":          4.26,
		"BROADWELL":        3.69,
		"HASWELL":          4.74,
		"COFFEE_LAKE":      5.42,
		"SANDY_BRIDGE":     8.58,
		"IVY_BRIDGE":       8.25,
		"AMD_EPYC_1ST_GEN": 2.55,
		"AMD_EPYC_2ND_GEN": 1.69,
		"AMD_EPYC_3RD_GEN": 2.02,
		// GPUs
		"NVIDIA_T4":         71,
		"NVIDIA_TESLA_K80":  306,
		"NVIDIA_TESLA_P100": 306,
		"NVIDIA_TESLA_V100": 306,
		"NVIDIA_TESLA_M60":  306,
		"NVIDIA_TESLA_P40":  255,
		"NVIDIA_TESLA_A100": 407,
		"XILINX_ALVEO_U250": 229.5,
	}

	// NETWORKING
	NETWORKING_ENERGY_COEFFICIENT = 0.001 // kWh / Gb

	// MEMORY
	MEMORY_ENERGY_COEFFICIENT = 0.000392 // kWh / Gb

	// PUE
	PUE_AVG = 1.185

	// UTILIZATION
	AVG_CPU_UTILIZATION_2020 = 50 / 100 // 50%

	// REPLICATION
	REPLICATION_FACTORS = map[string]float64{
		"STORAGE_LRS":    3,
		"STORAGE_ZRS":    3,
		"STORAGE_GRS":    6,
		"STORAGE_GZRS":   6,
		"STORAGE_DISKS":  3,
		"DATABASE_MYSQL": 3,
		"COSMOS_DB":      4,
		"SQL_DB":         3,
		"REDIS_CACHE":    2,
		"DEFAULT":        1,
	}

	// Lifespan
	SERVER_EXPECTED_LIFESPAN = 365 * 24 * 4 // 4 years in hours
)

func getMemory(computeProcessor string) float64 {
	return MEMORY_BY_COMPUTE_PROCESSOR[computeProcessor]
}

func getMinWatts(computeProcessor string) float64 {
	return MIN_WATTS_BY_COMPUTE_PROCESSOR[computeProcessor]
}

func getMaxWatts(computeProcessor string) float64 {
	return MAX_WATTS_BY_COMPUTE_PROCESSOR[computeProcessor]
}

func getPUE() float64 {
	return PUE_AVG
}

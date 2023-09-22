package azure_models

type VirtualMachine struct {
	Name              string
	Series            string
	VCPUs             int
	Memory            float64
	EmbodiedEmissions float64
}

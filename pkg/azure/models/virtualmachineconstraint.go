package azure_models

type VirtualMachineVCPUConstraint struct {
	Name                string
	VCPUs               int
	Memory              float64
	HighestVCPUPossible int
	EmbodiedEmissions   float64
}

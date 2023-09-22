package env

const (
	AzureOfferIDEnvVar        = "AZURE_OFFER_ID"
	AzureBillingAccountEnvVar = "AZURE_BILLING_ACCOUNT"

	OpenAzureEmissionsPortEnvVar = "OPEN_AZURE_EMISSIONS_PORT"
)


func GetOpenAzureEmissionsPort() int {
	return GetInt(OpenAzureEmissionsPortEnvVar, 9005)
}


// GetAzureOfferID returns the environment variable value for AzureOfferIDEnvVar which represents
// the Azure offer ID for determining prices.
func GetAzureOfferID() string {
	return Get(AzureOfferIDEnvVar, "")
}

// GetAzureBillingAccount returns the environment variable value for
// AzureBillingAccountEnvVar which represents the Azure billing
// account for determining prices. If this is specified
// customer-specific prices will be downloaded from the consumption
// price sheet API.
func GetAzureBillingAccount() string {
	return Get(AzureBillingAccountEnvVar, "")
}

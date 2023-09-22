package azure_models

type AzureCostDetails struct {
	AccountName                  string  `json:"accountname"`                       // 	EA, pay-as-you-go 	Display name of the EA enrollment account or pay-as-you-go billing account.
	AccountOwnerId               string  `json:"accountownerid"`                    // 	EA, pay-as-you-go 	Unique identifier for the EA enrollment account or pay-as-you-go billing account.
	AdditionalInfo               string  `json:"additionalinfo"`                    // 	All 				Service-specific metadata. For example, an image type for a virtual machine.
	BenefitId                    string  `json:"benefitid"`                         // 	EA, MCA 			Unique identifier for the purchased savings plan instance.
	BenefitName                  string  `json:"benefitname"`                       // 	EA, MCA 			Unique identifier for the purchased savings plan instance.
	BillingAccountId             string  `json:"billingaccountid"`                  // 	All 				Unique identifier for the root billing account.
	BillingAccountName           string  `json:"billingaccountname"`                // 	All 				Name of the billing account.
	BillingCurrency              string  `json:"billingcurrency"`                   // 	All 				Currency associated with the billing account.
	BillingCurrencyCode          string  `json:"billingcurrencycode"`               // 	All 				See BillingCurrency.
	BillingPeriod                string  `json:"billingperiod"`                     // 	EA, pay-as-you-go 	The billing period of the charge.
	BillingPeriodEndDate         string  `json:"billingperiodenddate"`              // 	All 				The end date of the billing period.
	BillingPeriodStartDate       string  `json:"billingperiodstartdate"`            // 	All 				The start date of the billing period.
	BillingProfileId             string  `json:"billingprofileid"`                  // 	All 				Unique identifier of the EA enrollment, pay-as-you-go subscription, MCA billing profile, or AWS consolidated account.
	BillingProfileName           string  `json:"billingprofilename"`                // 	All 				Name of the EA enrollment, pay-as-you-go subscription, MCA billing profile, or AWS consolidated account.
	ChargeType                   string  `json:"chargetype"`                        // 	All 				Indicates whether the charge represents usage (Usage), a purchase (Purchase), or a refund (Refund).
	ConsumedService              string  `json:"consumedservice"`                   // 	All 				Name of the service the charge is associated with.
	CostCenter                   string  `json:"costcenter"`                        // 	EA, MCA 			The cost center defined for the subscription for tracking costs (only available in open billing periods for MCA accounts).
	Cost                         string  `json:"cost"`                              // 	EA, pay-as-you-go 	See CostInBillingCurrency.
	CostAllocationRuleName       string  `json:"costallocationrulename"`            // 	EA, MCA 			Name of the Cost Allocation rule that's applicable to the record.
	CostInBillingCurrency        string  `json:"costinbillingcurrency"`             // 	EA, MCA 			Cost of the charge in the billing currency before credits or taxes.
	CostInPricingCurrency        string  `json:"costinpricingcurrency"`             // 	MCA 				Cost of the charge in the pricing currency before credits or taxes.
	Currency                     string  `json:"currency"`                          // 	EA, pay-as-you-go 	See BillingCurrency.
	CustomerName                 string  `json:"customername"`                      // 	MPA 				Name of the Azure Active Directory tenant for the customer's subscription.
	CustomerTenantId             string  `json:"customertenantid"`                  // 	MPA 				Identifier of the Azure Active Directory tenant of the customer's subscription.
	Date                         string  `json:"date"`                              // 	All 				The usage or purchase date of the charge.
	EffectivePrice               string  `json:"effectiveprice"`                    // 	All 				Blended unit price for the period. Blended prices average out any fluctuations in the unit price, like graduated tiering, which lowers the price as quantity increases over time.
	ExchangeRateDate             string  `json:"exchangeratedate"`                  // 	MCA 				Date the exchange rate was established.
	ExchangeRatePricingToBilling string  `json:"exchangeratepricingtobilling"`      // 	MCA 				Exchange rate used to convert the cost in the pricing currency to the billing currency.
	Frequency                    string  `json:"frequency"`                         // 	All 				Indicates whether a charge is expected to repeat. Charges can either happen once (OneTime), repeat on a monthly or yearly basis (Recurring), or be based on usage (UsageBased).
	InvoiceId                    string  `json:"invoiceid"`                         // 	pay-as-you-go, MCA 	The unique document ID listed on the invoice PDF.
	InvoiceSection               string  `json:"invoicesection"`                    // 	MCA 				See InvoiceSectionName.
	InvoiceSectionId             string  `json:"invoicesectionid"`                  // 	EA, MCA 			Unique identifier for the EA department or MCA invoice section.
	InvoiceSectionName           string  `json:"invoicesectionname"`                // 	EA, MCA 			Name of the EA department or MCA invoice section.
	IsAzureCreditEligible        string  `json:"isazurecrediteligible"`             // 	All 				Indicates if the charge is eligible to be paid for using Azure credits (Values: True or False).
	Location                     string  `json:"location" validate:"required"`      // 	MCA 				Normalized location of the resource, if different resource locations are configured for the same regions. Purchases and Marketplace usage may be shown as blank or unassigned.
	MeterCategory                string  `json:"metercategory" validate:"required"` // 	All 				Name of the classification category for the meter. For example, Cloud services and Networking.
	MeterId                      string  `json:"meterid"`                           // 	All 				The unique identifier for the meter.
	MeterName                    string  `json:"metername" validate:"required"`     // 	All 				The name of the meter. Purchases and Marketplace usage may be shown as blank or unassigned.
	MeterRegion                  string  `json:"meterregion"`                       // 	All 				Name of the datacenter location for services priced based on location. See Location.
	MeterSubCategory             string  `json:"metersubcategory"`                  // 	All 				Name of the meter subclassification category. Purchases and Marketplace usage may be shown as blank or unassigned.
	OfferId                      string  `json:"offerid"`                           // 	All 				Name of the offer purchased.
	pay_as_you_goPrice           string  `json:"pay-as-you-goPrice"`                // 	All 				Retail price for the resource.
	PartnerEarnedCreditApplied   string  `json:"partnerearnedcreditapplied"`        // 	MPA 				Indicates whether the partner earned credit has been applied.
	PartnerEarnedCreditRate      string  `json:"partnerearnedcreditrate"`           // 	MPA 				Rate of discount applied if there's a partner earned credit (PEC), based on partner admin link access.
	PartnerName                  string  `json:"partnername"`                       // 	MPA 				Name of the partner Azure Active Directory tenant.
	PartnerTenantId              string  `json:"partnertenantid"`                   // 	MPA 				Identifier for the partner's Azure Active Directory tenant.
	PartNumber                   string  `json:"partnumber"`                        // 	EA, pay-as-you-go 	Identifier used to get specific meter pricing.
	PlanName                     string  `json:"planname"`                          // 	EA, pay-as-you-go 	Marketplace plan name.
	PreviousInvoiceId            string  `json:"previousinvoiceid"`                 // 	MCA 				Reference to an original invoice if the line item is a refund.
	PricingCurrency              string  `json:"pricingcurrency"`                   // 	MCA 				Currency used when rating based on negotiated prices.
	PricingModel                 string  `json:"pricingmodel"`                      // 	All 				Identifier that indicates how the meter is priced. (Values: On Demand, Reservation, and Spot)
	Product                      string  `json:"product"`                           // 	All 				Name of the product.
	ProductId                    string  `json:"productid"`                         // 	MCA 				Unique identifier for the product.
	ProductOrderId               string  `json:"productorderid"`                    // 	All 				Unique identifier for the product order.
	ProductOrderName             string  `json:"productordername"`                  // 	All 				Unique name for the product order.
	Provider                     string  `json:"provider"`                          // 	All 				Identifier for product category or Line of Business. For example, Azure, Microsoft 365, and AWS.
	PublisherId                  string  `json:"publisherid"`                       // 	MCA 				The ID of the publisher. It's only available after the invoice is generated.
	PublisherName                string  `json:"publishername"`                     // 	All 				Publisher for Marketplace services.
	PublisherType                string  `json:"publishertype"`                     // 	All 				Supported values: Microsoft, Azure, AWS, Marketplace. Values are Microsoft for MCA accounts and Azure for EA and pay-as-you-go accounts.
	Quantity                     float64 `json:"quantity" validate:"required"`      // 	All 				The number of units purchased or consumed.
	ResellerName                 string  `json:"resellername"`                      // 	MPA 				The name of the reseller associated with the subscription.
	ResellerMpnId                string  `json:"resellermpnid"`                     // 	MPA 				ID for the reseller associated with the subscription.
	ReservationId                string  `json:"reservationid"`                     // 	EA, MCA 			Unique identifier for the purchased reservation instance.
	ReservationName              string  `json:"reservationname"`                   // 	EA, MCA 			Name of the purchased reservation instance.
	ResourceGroup                string  `json:"resourcegroup"`                     // 	All 				Name of the resource group the resource is in. Not all charges come from resources deployed to resource groups. Charges that don't have a resource group will be shown as null or empty, Others, or Not applicable.
	ResourceId                   string  `json:"resourceid"`                        // 	All 				Unique identifier of the Azure Resource Manager resource.
	ResourceLocation             string  `json:"resourcelocation"`                  // 	All 				Datacenter location where the resource is running. See Location.
	ResourceName                 string  `json:"resourcename"`                      // 	EA, pay-as-you-go 	Name of the resource. Not all charges come from deployed resources. Charges that don't have a resource type will be shown as null/empty, Others , or Not applicable.
	ResourceType                 string  `json:"resourcetype"`                      // 	MCA 				Type of resource instance. Not all charges come from deployed resources. Charges that don't have a resource type will be shown as null/empty, Others , or Not applicable.
	RoundingAdjustment           string  `json:"roundingadjustment"`                // 	EA, MCA 			Rounding adjustment represents the quantization that occurs during cost calculation. When the calculated costs are converted to the invoiced total, small rounding errors can occur. The rounding errors are represented as rounding adjustment to ensure that the costs shown in Cost Management align to the invoice.
	ServiceFamily                string  `json:"servicefamily"`                     // 	MCA 				Service family that the service belongs to.
	ServiceInfo                  string  `json:"serviceinfo"`                       // 	All 				Service-specific metadata.
	ServiceInfo2                 string  `json:"serviceinfo2"`                      // 	All 				Legacy field with optional service-specific metadata.
	ServicePeriodEndDate         string  `json:"serviceperiodenddate"`              // 	MCA 				The end date of the rating period that defined and locked pricing for the consumed or purchased service.
	ServicePeriodStartDate       string  `json:"serviceperiodstartdate"`            // 	MCA 				The start date of the rating period that defined and locked pricing for the consumed or purchased service.
	SubscriptionId               string  `json:"subscriptionid"`                    // 	All 				Unique identifier for the Azure subscription.
	SubscriptionName             string  `json:"subscriptionname"`                  // 	All 				Name of the Azure subscription.
	Tags                         string  `json:"tags"`                              // 	All 				Tags assigned to the resource. Doesn't include resource group tags. Can be used to group or distribute costs for internal chargeback. For more information, see Organize your Azure resources with tags.
	Term                         string  `json:"term"`                              // 	All 				Displays the term for the validity of the offer. For example: In case of reserved instances, it displays 12 months as the Term. For one-time purchases or recurring purchases, Term is one month (SaaS, Marketplace Support). Not applicable for Azure consumption.
	UnitOfMeasure                string  `json:"unitofmeasure" validate:"required"` // 	All 				The unit of measure for billing for the service. For example, compute services are billed per hour.
	UnitPrice                    string  `json:"unitprice"`                         // 	EA, pay-as-you-go 	The price per unit for the charge.
}

type AzureCostDetailsTestCase struct {
	Data           AzureCostDetails
	ExpectedInt    int
	ExpectedFloat  float64
	ExpectedString string
}

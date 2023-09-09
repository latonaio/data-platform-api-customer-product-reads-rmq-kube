package requests

type CustomerProduct struct {
	BusinessPartner              int     `json:"BusinessPartner"`
	Customer                     int     `json:"Customer"`
	SalesOrganization            string  `json:"SalesOrganization"`
	DistributionChannel          string  `json:"DistributionChannel"`
	Product                      string  `json:"Product"`
	ProductByCustomer            *string `json:"ProductByCustomer"`
	ProductDescriptionByCustomer *string `json:"ProductDescriptionByCustomer"`
	BaseUnit                     *string `json:"BaseUnit"`
	SalesUnit                    *string `json:"SalesUnit"`
}

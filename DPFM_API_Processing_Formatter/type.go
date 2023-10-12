package dpfm_api_processing_formatter

type PostalCodeUpdates struct {
	PostalCode          string `json:"PostalCode"`
	Country             string `json:"Country"`
	LocalSubRegion      string `json:"LocalSubRegion"`
	LocalRegion         string `json:"LocalRegion"`
	GlobalRegion        string `json:"GlobalRegion"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type PostalCodeAddressUpdates struct {
	PostalCode                  string  `json:"PostalCode"`
	Country                     string  `json:"Country"`
	PostalCodeAddressDetailText *string `json:"PostalCodeAddressDetailText"`
	CityName                    *string `json:"CityName"`
	Building                    *string `json:"Building"`
	Floor                       *int    `json:"Floor"`
	PostalCodeAddressTotalText  *string `json:"PostalCodeAddressTotalText"`
}

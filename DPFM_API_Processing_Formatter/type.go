package dpfm_api_processing_formatter

type PostalCodeUpdates struct {
	PostalCode string `json:"PostalCode"`
}

type PostalCodeAddressUpdates struct {
	PostalCode                  string  `json:"PostalCode"`
	Country                     string  `json:"Country"`
	PostalCodeAddressDetailText *string `json:"PostalCodeAddressDetailText"`
}

package requests

type PostalCodeAddress struct {
	PostalCode                  string  `json:"PostalCode"`
	Country                     string  `json:"Country"`
	PostalCodeAddressDetailText *string `json:"PostalCodeAddressDetailText"`
	CityName                    *string `json:"CityName"`
	Building                    *string `json:"Building"`
	Floor                       *int    `json:"Floor"`
	PostalCodeAddressTotalText  *string `json:"PostalCodeAddressTotalText"`
	IsMarkedForDeletion         *bool   `json:"IsMarkedForDeletion"`
}

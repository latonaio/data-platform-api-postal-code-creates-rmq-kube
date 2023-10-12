package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-postal-code-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToPostalCodeUpdates(postalCode dpfm_api_input_reader.PostalCode) *PostalCodeUpdates {
	data := postalCode

	return &PostalCodeUpdates{
		PostalCode          string `json:"PostalCode"`
		Country             string `json:"Country"`
		LocalSubRegion      string `json:"LocalSubRegion"`
		LocalRegion         string `json:"LocalRegion"`
		GlobalRegion        string `json:"GlobalRegion"`
		IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
			}
}

func ConvertToStorageLoationUpdates(postalCode dpfm_api_input_reader.PostalCode, accounting dpfm_api_input_reader.PostalCodeAddress) *PostalCodeAddressUpdates {
	dataPostalCode := postalCode
	data := postalCodeAddress

	return &PostalCodeAddressUpdates{
		PostalCode                  string  `json:"PostalCode"`
		Country                     string  `json:"Country"`
		PostalCodeAddressDetailText *string `json:"PostalCodeAddressDetailText"`
		CityName                    *string `json:"CityName"`
		Building                    *string `json:"Building"`
		Floor                       *int    `json:"Floor"`
		PostalCodeAddressTotalText  *string `json:"PostalCodeAddressTotalText"`
		IsMarkedForDeletion         *bool   `json:"IsMarkedForDeletion"`
		}
}

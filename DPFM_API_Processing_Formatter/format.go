package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-postal-code-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToPostalCodeUpdates(postalCode dpfm_api_input_reader.PostalCode) *PostalCodeUpdates {
	data := postalCode

	return &PostalCodeUpdates{
		PostalCode: *data.PostalCode,
	}
}

func ConvertToPostalCodeAddressUpdates(postalCode dpfm_api_input_reader.PostalCode, postalCodeAddress dpfm_api_input_reader.PostalCodeAddress) *PostalCodeAddressUpdates {
	//	dataPostalCode := postalCode
	data := postalCodeAddress

	return &PostalCodeAddressUpdates{
		PostalCode:                  *data.PostalCode,
		Country:                     *data.Country,
		PostalCodeAddressDetailText: data.PostalCodeAddressDetailText,
	}
}

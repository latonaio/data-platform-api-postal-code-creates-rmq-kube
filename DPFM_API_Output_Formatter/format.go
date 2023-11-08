package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-postal-code-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-postal-code-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToPostalCodeCreates(sdc *dpfm_api_input_reader.SDC) (*PostalCode, error) {
	data := sdc.PostalCode

	postalCode, err := TypeConverter[*PostalCode](data)
	if err != nil {
		return nil, err
	}

	return postalCode, nil
}

func ConvertToPostalCodeAddressCreates(sdc *dpfm_api_input_reader.SDC) (*[]PostalCodeAddress, error) {
	items := make([]PostalCodeAddress, 0)

	for _, data := range sdc.PostalCode.PostalCodeAddress {
		PostalCodeAddress, err := TypeConverter[*PostalCodeAddress](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *PostalCodeAddress)
	}

	return &items, nil
}

func ConvertToPostalCodeUpdates(postalCodeData dpfm_api_input_reader.PostalCode) (*PostalCode, error) {
	data := postalCodeData

	postalCode, err := TypeConverter[*PostalCode](data)
	if err != nil {
		return nil, err
	}

	return postalCode, nil
}

func ConvertToPostalCodeAddressUpdates(postalCodeAddressUpdates *[]dpfm_api_processing_formatter.PostalCodeAddressUpdates) (*[]PostalCodeAddress, error) {
	items := make([]PostalCodeAddress, 0)

	for _, data := range *postalCodeAddressUpdates {
		postalCodeAddress, err := TypeConverter[*PostalCodeAddress](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *postalCodeAddress)
	}

	return &items, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}

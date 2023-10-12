package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-postal-code-creates-rmq-kube/DPFM_API_Input_Reader"
	"encoding/json"
	"time"

	"golang.org/x/xerrors"
)

func ConvertToPostalCodeCreates(sdc *dpfm_api_input_reader.SDC) (*PostalCode, error) {
	data := sdc.PostalCode

	postalCode, err := TypeConverter[*PostalCode](data)
	if err != nil {
		return nil, err
	}
	// postalCode.CreationDate = *getSystemDatePtr()
	// postalCode.CreationTime = *getSystemTimePtr()
	// postalCode.LastChangeDate = getSystemDatePtr()
	// postalCode.LastChangeTime = getSystemTimePtr()

	return postalCode, nil
}

func ConvertToPostalCodeAddressCreates(sdc *dpfm_api_input_reader.SDC) (*PostalCodeAddress, error) {
	data := sdc.PostalCodeAddress

	postalCodeAddress, err := TypeConverter[*PostalCodeAddress](data)
	if err != nil {
		return nil, err
	}

	return postalCodeAddress, nil
}

func ConvertToPostalCodeUpdates(postalCodeData dpfm_api_input_reader.PostalCode) (*PostalCode, error) {
	data := postalCodeData

	postalCode, err := TypeConverter[*PostalCode](data)
	if err != nil {
		return nil, err
	}

	return postalCode, nil
}

func ConvertToPostalCodeAddressUpdates(postalCodeData dpfm_api_input_reader.PostalCodeAddress) (*PostalCodeAddress, error) {
	data := postalCodeAddressData

	postalCodeAddress, err := TypeConverter[*PostalCodeAddress](data)
	if err != nil {
		return nil, err
	}

	return postalCodeAddress, nil
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

func getSystemDatePtr() *string {
	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// day := time.Now().In(jst)

	day := time.Now()
	res := day.Format("2006-01-02")
	return &res
}

func getSystemTimePtr() *string {
	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// day := time.Now().In(jst)

	day := time.Now()
	res := day.Format("15:04:05")
	return &res
}

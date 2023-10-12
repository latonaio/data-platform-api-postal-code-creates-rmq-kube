package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-postal-code-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-postal-code-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-postal-code-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var postalCode *dpfm_api_output_formatter.PostalCode
	var postalCodeAddress *[]dpfm_api_output_formatter.PostalCodeAddress
	for _, fn := range accepter {
		switch fn {
		case "PostalCode":
			postalCode = c.postalCodeCreateSql(nil, mtx, input, output, errs, log)
		case "PostalCodeAddress":
			postalCodeAddress = c.postalCodeAddressCreateSql(nil, mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		PostalCode:        postalCode,
		PostalCodeAddress: postalCodeAddress,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var postalCode *dpfm_api_output_formatter.PostalCode
	var postalCodeAddress *[]dpfm_api_output_formatter.PostalCodeAddress
	for _, fn := range accepter {
		switch fn {
		case "PostalCode":
			postalCode = c.postalCodeUpdateSql(mtx, input, output, errs, log)
		case "PostalCodeAddress":
			postalCodeAddress = c.postalCodeAddressUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		PostalCode:        postalCode,
		PostalCodeAddress: postalCodeAddress,
	}

	return data
}

func (c *DPFMAPICaller) postalCodeCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.PostalCode {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	postalCodeData := input.PostalCode
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": postalCodeData, "function": "PostalCodePostalCode", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "PostalCode Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPostalCodeCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) postalCodeAddressCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PostalCodeAddress {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.PostalCode.PostalCodeAddress {
		input.PostalCode.PostalCodeAddress[i].PostalCode = input.PostalCode.PostalCode
		postalCodeAddressData := input.PostalCode.PostalCodeAddress[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": accountingData, "function": "PostalCodePostalCodeAddress", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "PostalCodeAddress Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPostalCodeAddressCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) postalCodeUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.PostalCode {
	postalCode := input.PostalCode
	postalCodeData := dpfm_api_processing_formatter.ConvertToPostalCodeUpdates(postalCode)

	sessionID := input.RuntimeSessionID
	if postalCodeIsUpdate(postalCodeData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": postalCodeData, "function": "PostalCodePostalCode", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "PostalCode Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPostalCodeUpdates(postalCode)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) postalCodeAddressUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PostalCodeAddress {
	req := make([]dpfm_api_processing_formatter.PostalCodeAddressUpdates, 0)
	sessionID := input.RuntimeSessionID

	postalCode := input.PostalCode
	for _, postalCodeAddress := range postalCode.PostalCodeAddress {
		postalCodeAddressData := *dpfm_api_processing_formatter.ConvertToPostalCodeAddressUpdates(postalCode, postalCodeAddress)

		if postalCodeAddressIsUpdate(&postalCodeAddressData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": postalCodeAddressData, "function": "PostalCodePostalCodeAddress", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "PostalCodeAddress Data cannot update"
				return nil
			}
		}
		req = append(req, postalCodeAddressData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPostalCodeAddressUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func postalCodeIsUpdate(postalCode *dpfm_api_processing_formatter.PostalCodeUpdates) bool {
	postalCode := postalCode.PostalCode
	country := postalCode.Country

	return !(postalCode == 0 || postalCode == "")
}

func postalCodeAddressIsUpdate(postalCodeAddress *dpfm_api_processing_formatter.PostalCodeAddressUpdates) bool {
	postalCode := postalCodeAddress.PostalCode
	country := postalCodeAddress.Country

	return !(postalCode == 0 || postalCode == "")
}

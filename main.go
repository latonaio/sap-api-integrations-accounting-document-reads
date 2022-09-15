package main

import (
	sap_api_caller "sap-api-integrations-accounting-document-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-accounting-document-reads/SAP_API_Input_Reader"
	"sap-api-integrations-accounting-document-reads/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	sap_api_time_value_converter "github.com/latonaio/sap-api-time-value-converter"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	fr := sap_api_input_reader.NewFileReader()
	gc := sap_api_request_client_header_setup.NewSAPRequestClientWithOption(conf.SAP)
	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		"100",
		gc,
		l,
	)
	inputSDC := fr.ReadSDC("./Inputs/SDC_Accounting_Document_Item_sample.json")
	sap_api_time_value_converter.ChangeTimeFormatToSAPFormatStruct(&inputSDC)
	accepter := inputSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Item",
		}
	}

	caller.AsyncGetAccountingDocument(
		inputSDC.AccountingDocument.CompanyCode,
		inputSDC.AccountingDocument.FiscalYear,
		inputSDC.AccountingDocument.AccountingDocument,
		accepter,
	)
}

package main

import (
	sap_api_caller "sap-api-integrations-product-master-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-product-master-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inputSDC := fr.ReadSDC("./Inputs/SDC_Product_Master_General_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inputSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"General", "Plant", "MRPArea", "Procurement",
			"WorkScheduling", "SalesPlant",
			"Accounting", "SalesOrganization", "ProductDescByProduct", "ProductDescByDesc",
			"Quality", "SalesTax",
		}
	}

	caller.AsyncGetProductMaster(
		inputSDC.Product.Product,
		inputSDC.Product.Plant.Plant,
		inputSDC.Product.Plant.MRPArea.MRPArea,
		inputSDC.Product.Accounting.ValuationArea,
		inputSDC.Product.SalesOrganization.ProductSalesOrg,
		inputSDC.Product.SalesOrganization.ProductDistributionChnl,
		inputSDC.Product.ProductDescription.Language,
		inputSDC.Product.ProductDescription.ProductDescription,
		inputSDC.Product.SalesTax.Country,
		inputSDC.Product.SalesTax.TaxCategory,
		accepter,
	)
}

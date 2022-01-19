package main

import (
	sap_api_caller "sap-api-integrations-product-master-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-product-master-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Product_Master_Sales_Tax_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"General", "Plant", "MRPArea", "Procurement",
			"WorkScheduling", "SalesPlant",
			"Accounting", "SalesOrganization", "ProductDescByProduct", "ProductDescByDesc",
			"Quality", "SalesTax",
		}
	}

	caller.AsyncGetProductMaster(
		inoutSDC.Product.Product,
		inoutSDC.Product.Plant.Plant,
		inoutSDC.Product.Plant.MRPArea.MRPArea,
		inoutSDC.Product.Accounting.ValuationArea,
		inoutSDC.Product.SalesOrganization.ProductSalesOrg,
		inoutSDC.Product.SalesOrganization.ProductDistributionChnl,
		inoutSDC.Product.ProductDescription.Language,
		inoutSDC.Product.ProductDescription.ProductDescription,
		inoutSDC.Product.SalesTax.Country,
		inoutSDC.Product.SalesTax.TaxCategory,
		accepter,
	)
}

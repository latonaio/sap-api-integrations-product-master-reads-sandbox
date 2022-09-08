package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-product-master-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_get_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_get_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_get_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetProductMaster(product, plant, mrpArea, valuationArea, productSalesOrg, productDistributionChnl, language, productDescription, country, taxCategory string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(product)
				wg.Done()
			}()
		case "Plant":
			func() {
				c.Plant(product, plant)
				wg.Done()
			}()
		case "MRPArea":
			func() {
				c.MRPArea(product, plant, mrpArea)
				wg.Done()
			}()
		case "Procurement":
			func() {
				c.Procurement(product, plant)
				wg.Done()
			}()
		case "WorkScheduling":
			func() {
				c.WorkScheduling(product, plant)
				wg.Done()
			}()
		case "SalesPlant":
			func() {
				c.SalesPlant(product, plant)
				wg.Done()
			}()
		case "Accounting":
			func() {
				c.Accounting(product, valuationArea)
				wg.Done()
			}()
		case "SalesOrganization":
			func() {
				c.SalesOrganization(product, productSalesOrg, productDistributionChnl)
				wg.Done()
			}()
		case "ProductDescByProduct":
			func() {
				c.ProductDescByProduct(product, language)
				wg.Done()
			}()
		case "ProductDescByDesc":
			func() {
				c.ProductDescByDesc(language, productDescription)
				wg.Done()
			}()
		case "Quality":
			func() {
				c.Quality(product, plant)
				wg.Done()
			}()
		case "SalesTax":
			func() {
				c.SalesTax(product, country, taxCategory)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) General(product string) {
	generalData, err := c.callProductSrvAPIRequirementGeneral("A_Product", product)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(generalData)

	productDescData, err := c.callToProductDesc(generalData[0].ToProductDesc)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(productDescData)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementGeneral(api, product string) ([]sap_api_output_formatter.General, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	param := c.getQueryWithGeneral(map[string]string{}, product)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToProductDesc(url string) ([]sap_api_output_formatter.ToProductDesc, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToProductDesc(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Plant(product, plant string) {
	data, err := c.callProductSrvAPIRequirementPlant("A_ProductPlant", product, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementPlant(api, product, plant string) ([]sap_api_output_formatter.Plant, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")

	param := c.getQueryWithPlant(map[string]string{}, product, plant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPlant(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MRPArea(product, plant, mrpArea string) {
	data, err := c.callProductSrvAPIRequirementMRPArea("A_ProductPlantMRPArea", product, plant, mrpArea)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementMRPArea(api, product, plant, mrpArea string) ([]sap_api_output_formatter.MRPArea, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")

	param := c.getQueryWithMRPArea(map[string]string{}, product, plant, mrpArea)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToMRPArea(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Procurement(product, plant string) {
	data, err := c.callProductSrvAPIRequirementProcurement("A_ProductPlantProcurement", product, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementProcurement(api, product, plant string) ([]sap_api_output_formatter.Procurement, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")

	param := c.getQueryWithProcurement(map[string]string{}, product, plant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToProcurement(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) WorkScheduling(product, plant string) {
	data, err := c.callProductSrvAPIRequirementWorkScheduling("A_ProductWorkScheduling", product, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementWorkScheduling(api, product, plant string) ([]sap_api_output_formatter.WorkScheduling, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")

	param := c.getQueryWithWorkScheduling(map[string]string{}, product, plant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToWorkScheduling(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SalesPlant(product, plant string) {
	data, err := c.callProductSrvAPIRequirementSalesPlant("A_ProductPlantSales", product, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementSalesPlant(api, product, plant string) ([]sap_api_output_formatter.SalesPlant, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")

	param := c.getQueryWithSalesPlant(map[string]string{}, product, plant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSalesPlant(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Accounting(product, valuationArea string) {
	data, err := c.callProductSrvAPIRequirementAccounting("A_ProductValuation", product, valuationArea)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementAccounting(api, product, valuationArea string) ([]sap_api_output_formatter.Accounting, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")

	param := c.getQueryWithAccounting(map[string]string{}, product, valuationArea)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToAccounting(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SalesOrganization(product, productSalesOrg, productDistributionChnl string) {
	data, err := c.callProductSrvAPIRequirementSalesOrganization("A_ProductSalesDelivery", product, productSalesOrg, productDistributionChnl)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementSalesOrganization(api, product, productSalesOrg, productDistributionChnl string) ([]sap_api_output_formatter.SalesOrganization, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")

	param := c.getQueryWithSalesOrganization(map[string]string{}, product, productSalesOrg, productDistributionChnl)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSalesOrganization(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ProductDescByProduct(product, language string) {
	data, err := c.callProductSrvAPIRequirementProductDescByProduct("A_ProductDescription", product, language)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementProductDescByProduct(api, product, language string) ([]sap_api_output_formatter.ProductDesc, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")

	param := c.getQueryWithProductDescByProduct(map[string]string{}, product, language)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToProductDesc(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ProductDescByDesc(language, productDescription string) {
	data, err := c.callProductSrvAPIRequirementProductDescByDesc("A_ProductDescription", language, productDescription)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementProductDescByDesc(api, language, productDescription string) ([]sap_api_output_formatter.ProductDesc, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")

	param := c.getQueryWithProductDescByDesc(map[string]string{}, language, productDescription)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToProductDesc(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Quality(product, plant string) {
	data, err := c.callProductSrvAPIRequirementQuality("A_ProductPlantQualityMgmt", product, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementQuality(api, product, plant string) ([]sap_api_output_formatter.Quality, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")

	param := c.getQueryWithQuality(map[string]string{}, product, plant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToQuality(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SalesTax(product, country, taxCategory string) {
	data, err := c.callProductSrvAPIRequirementSalesTax("A_ProductSalesTax", product, country, taxCategory)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementSalesTax(api, product, country, taxCategory string) ([]sap_api_output_formatter.SalesTax, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")

	param := c.getQueryWithSalesTax(map[string]string{}, product, country, taxCategory)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSalesTax(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

// func (c *SAPAPICaller) setHeaderAPIKeyAccept(params map[string]string) map[string]string {
// 	if len(params) == 0 {
// 		params = make(map[string]string, 1)
// 	}

// 	params["APIKey"] = c.apiKey
// 	params["Accept"] = "application/json"
// 	return params
// }

func (c *SAPAPICaller) getQueryWithGeneral(params map[string]string, product string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Product eq '%s'", product)
	return params
}

func (c *SAPAPICaller) getQueryWithPlant(params map[string]string, product, plant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Product eq '%s' and Plant eq '%s'", product, plant)
	return params
}

func (c *SAPAPICaller) getQueryWithMRPArea(params map[string]string, product, plant, mrpArea string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Product eq '%s' and Plant eq '%s' and MRPArea eq '%s'", product, plant, mrpArea)
	return params
}

func (c *SAPAPICaller) getQueryWithProcurement(params map[string]string, product, plant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Product eq '%s' and Plant eq '%s'", product, plant)
	return params
}

func (c *SAPAPICaller) getQueryWithWorkScheduling(params map[string]string, product, plant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Product eq '%s' and Plant eq '%s'", product, plant)
	return params
}

func (c *SAPAPICaller) getQueryWithSalesPlant(params map[string]string, product, plant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Product eq '%s' and Plant eq '%s'", product, plant)
	return params
}

func (c *SAPAPICaller) getQueryWithAccounting(params map[string]string, product, valuationArea string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Product eq '%s' and ValuationArea eq '%s'", product, valuationArea)
	return params
}

func (c *SAPAPICaller) getQueryWithSalesOrganization(params map[string]string, product, productSalesOrg, productDistributionChnl string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Product eq '%s' and ProductSalesOrg eq '%s' and ProductDistributionChnl eq '%s'", product, productSalesOrg, productDistributionChnl)
	return params
}

func (c *SAPAPICaller) getQueryWithProductDescByProduct(params map[string]string, product, language string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Product eq '%s' and Language eq '%s'", product, language)
	return params
}

func (c *SAPAPICaller) getQueryWithProductDescByDesc(params map[string]string, language, productDescription string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Language eq '%s' and substringof('%s', ProductDescription)", language, productDescription)
	return params
}

func (c *SAPAPICaller) getQueryWithQuality(params map[string]string, product, plant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Product eq '%s' and Plant eq '%s'", product, plant)
	return params
}

func (c *SAPAPICaller) getQueryWithSalesTax(params map[string]string, product, country, taxCategory string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Product eq '%s' and Country eq '%s' and TaxCategory eq '%s'", product, country, taxCategory)
	return params
}

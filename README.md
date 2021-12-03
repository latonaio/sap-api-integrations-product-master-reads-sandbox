# sap-api-integrations-product-master-reads  
sap-api-integrations-product-master-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で品目マスタデータを取得するマイクロサービスです。  
sap-api-integrations-product-master-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-product-master-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_PRODUCT_SRV_0001/overview  

## 動作環境  
sap-api-integrations-product-master-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）   
・ AION のリソース （推奨)   
・ OS: LinuxOS （必須）   
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用
sap-api-integrations-product-master-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 品目マスタ の 一般データ が取得された結果の JSON の例です。  
以下の項目のうち、"BaseUnit" ～ "WeightUnit" は、/SAP_API_Caller/formatter.go 内 の type Material struct {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"BaseUnit": "AU",
	"Division": "00",
	"GrossWeight": "0.000",
	"Material": "A001",
	"ProductGroup": "A001",
	"ProductStandardID": "",
	"Product_desc": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCT_SRV/A_Product('A001')/to_Description",
	"SizeOrDimensionText": "",
	"ValidityStartDate": "",
	"WeightUnit": "KG",
	"cursor": "/Users/kyotatashiro/go/src/sap-api-integrations-product-master-reads/SAP_API_Caller.(*SAPAPICaller).Product",
	"level": "INFO",
	"time": "2021-11-26T15:50:30.156715+09:00"
}

```

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAllの値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Product", "Plant", "Accounting" が指定されています。    
  
```
  "api_schema": "sap.s4.beh.product.v1.Product.Created.v1",
  "accepter": ["Product", "Plant", "Accounting"],
  "material_code": "21",
  "deleted": false
```
  
* 全データを取得する際ののsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
  "api_schema": "sap.s4.beh.product.v1.Product.Created.v1",
  "accepter": ["All"],
  "material_code": "21",
  "deleted": false
```

## 指定されたデータ種別のコール

accepter におけるデータ種別の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetProductMaster(product, plant, mrpArea, valuationArea, productSalesOrg, productDistributionChnl string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Product":
			func() {
				c.Product(product)
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
		case "ProductDesc":
			func() {
				c.ProductDesc(product)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

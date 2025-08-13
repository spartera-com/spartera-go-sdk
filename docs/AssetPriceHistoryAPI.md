# \AssetPriceHistoryAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompaniesCompanyIdAssetsAssetIdPricesActiveGet**](AssetPriceHistoryAPI.md#CompaniesCompanyIdAssetsAssetIdPricesActiveGet) | **Get** /companies/{company_id}/assets/{asset_id}/prices/active | Get the currently active price for an asset
[**CompaniesCompanyIdAssetsAssetIdPricesAphIdDelete**](AssetPriceHistoryAPI.md#CompaniesCompanyIdAssetsAssetIdPricesAphIdDelete) | **Delete** /companies/{company_id}/assets/{asset_id}/prices/{aph_id} | Delete single price history record by ID
[**CompaniesCompanyIdAssetsAssetIdPricesAphIdGet**](AssetPriceHistoryAPI.md#CompaniesCompanyIdAssetsAssetIdPricesAphIdGet) | **Get** /companies/{company_id}/assets/{asset_id}/prices/{aph_id} | Get single price history record by ID
[**CompaniesCompanyIdAssetsAssetIdPricesAphIdPatch**](AssetPriceHistoryAPI.md#CompaniesCompanyIdAssetsAssetIdPricesAphIdPatch) | **Patch** /companies/{company_id}/assets/{asset_id}/prices/{aph_id} | Update an existing price history record by ID
[**CompaniesCompanyIdAssetsAssetIdPricesCalculateCreditsPost**](AssetPriceHistoryAPI.md#CompaniesCompanyIdAssetsAssetIdPricesCalculateCreditsPost) | **Post** /companies/{company_id}/assets/{asset_id}/prices/calculate_credits | Calculate the credit equivalent for a given USD price without saving
[**CompaniesCompanyIdAssetsAssetIdPricesDiscountPost**](AssetPriceHistoryAPI.md#CompaniesCompanyIdAssetsAssetIdPricesDiscountPost) | **Post** /companies/{company_id}/assets/{asset_id}/prices/discount | POST /companies/{company_id}/assets/{asset_id}/prices/discount
[**CompaniesCompanyIdAssetsAssetIdPricesGet**](AssetPriceHistoryAPI.md#CompaniesCompanyIdAssetsAssetIdPricesGet) | **Get** /companies/{company_id}/assets/{asset_id}/prices | Get all price history records for a specific asset
[**CompaniesCompanyIdAssetsAssetIdPricesPost**](AssetPriceHistoryAPI.md#CompaniesCompanyIdAssetsAssetIdPricesPost) | **Post** /companies/{company_id}/assets/{asset_id}/prices | Create a new price history record for an asset



## CompaniesCompanyIdAssetsAssetIdPricesActiveGet

> CompaniesCompanyIdAssetsAssetIdPricesGet200Response CompaniesCompanyIdAssetsAssetIdPricesActiveGet(ctx, companyId, assetId).Execute()

Get the currently active price for an asset

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spartera-com/spartera-go-sdk"
)

func main() {
	companyId := "companyId_example" // string | 
	assetId := "assetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesActiveGet(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesActiveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdPricesActiveGet`: CompaniesCompanyIdAssetsAssetIdPricesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesActiveGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdPricesActiveGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdAssetsAssetIdPricesGet200Response**](CompaniesCompanyIdAssetsAssetIdPricesGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdPricesAphIdDelete

> CompaniesCompanyIdAssetsAssetIdPricesAphIdDelete200Response CompaniesCompanyIdAssetsAssetIdPricesAphIdDelete(ctx, companyId, assetId, aphId).Execute()

Delete single price history record by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spartera-com/spartera-go-sdk"
)

func main() {
	companyId := "companyId_example" // string | 
	assetId := "assetId_example" // string | 
	aphId := "aphId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesAphIdDelete(context.Background(), companyId, assetId, aphId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesAphIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdPricesAphIdDelete`: CompaniesCompanyIdAssetsAssetIdPricesAphIdDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesAphIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 
**aphId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdPricesAphIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CompaniesCompanyIdAssetsAssetIdPricesAphIdDelete200Response**](CompaniesCompanyIdAssetsAssetIdPricesAphIdDelete200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdPricesAphIdGet

> CompaniesCompanyIdAssetsAssetIdPricesAphIdGet200Response CompaniesCompanyIdAssetsAssetIdPricesAphIdGet(ctx, companyId, assetId, aphId).Execute()

Get single price history record by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spartera-com/spartera-go-sdk"
)

func main() {
	companyId := "companyId_example" // string | 
	assetId := "assetId_example" // string | 
	aphId := "aphId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesAphIdGet(context.Background(), companyId, assetId, aphId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesAphIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdPricesAphIdGet`: CompaniesCompanyIdAssetsAssetIdPricesAphIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesAphIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 
**aphId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdPricesAphIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CompaniesCompanyIdAssetsAssetIdPricesAphIdGet200Response**](CompaniesCompanyIdAssetsAssetIdPricesAphIdGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdPricesAphIdPatch

> CompaniesCompanyIdAssetsAssetIdPricesAphIdPatch200Response CompaniesCompanyIdAssetsAssetIdPricesAphIdPatch(ctx, companyId, assetId, aphId).AssetPriceHistoryUpdate(assetPriceHistoryUpdate).Execute()

Update an existing price history record by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spartera-com/spartera-go-sdk"
)

func main() {
	companyId := "companyId_example" // string | 
	assetId := "assetId_example" // string | 
	aphId := "aphId_example" // string | 
	assetPriceHistoryUpdate := *openapiclient.NewAssetPriceHistoryUpdate() // AssetPriceHistoryUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesAphIdPatch(context.Background(), companyId, assetId, aphId).AssetPriceHistoryUpdate(assetPriceHistoryUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesAphIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdPricesAphIdPatch`: CompaniesCompanyIdAssetsAssetIdPricesAphIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesAphIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 
**aphId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdPricesAphIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **assetPriceHistoryUpdate** | [**AssetPriceHistoryUpdate**](AssetPriceHistoryUpdate.md) |  | 

### Return type

[**CompaniesCompanyIdAssetsAssetIdPricesAphIdPatch200Response**](CompaniesCompanyIdAssetsAssetIdPricesAphIdPatch200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdPricesCalculateCreditsPost

> CompaniesCompanyIdAssetsAssetIdPricesPost200Response CompaniesCompanyIdAssetsAssetIdPricesCalculateCreditsPost(ctx, companyId, assetId).AssetPriceHistoryInput(assetPriceHistoryInput).Execute()

Calculate the credit equivalent for a given USD price without saving

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spartera-com/spartera-go-sdk"
)

func main() {
	companyId := "companyId_example" // string | 
	assetId := "assetId_example" // string | 
	assetPriceHistoryInput := *openapiclient.NewAssetPriceHistoryInput("AssetId_example") // AssetPriceHistoryInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesCalculateCreditsPost(context.Background(), companyId, assetId).AssetPriceHistoryInput(assetPriceHistoryInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesCalculateCreditsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdPricesCalculateCreditsPost`: CompaniesCompanyIdAssetsAssetIdPricesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesCalculateCreditsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdPricesCalculateCreditsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assetPriceHistoryInput** | [**AssetPriceHistoryInput**](AssetPriceHistoryInput.md) |  | 

### Return type

[**CompaniesCompanyIdAssetsAssetIdPricesPost200Response**](CompaniesCompanyIdAssetsAssetIdPricesPost200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdPricesDiscountPost

> CompaniesCompanyIdAssetsAssetIdPricesPost200Response CompaniesCompanyIdAssetsAssetIdPricesDiscountPost(ctx, companyId, assetId).AssetPriceHistoryInput(assetPriceHistoryInput).Execute()

POST /companies/{company_id}/assets/{asset_id}/prices/discount

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spartera-com/spartera-go-sdk"
)

func main() {
	companyId := "companyId_example" // string | 
	assetId := "assetId_example" // string | 
	assetPriceHistoryInput := *openapiclient.NewAssetPriceHistoryInput("AssetId_example") // AssetPriceHistoryInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesDiscountPost(context.Background(), companyId, assetId).AssetPriceHistoryInput(assetPriceHistoryInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesDiscountPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdPricesDiscountPost`: CompaniesCompanyIdAssetsAssetIdPricesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesDiscountPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdPricesDiscountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assetPriceHistoryInput** | [**AssetPriceHistoryInput**](AssetPriceHistoryInput.md) |  | 

### Return type

[**CompaniesCompanyIdAssetsAssetIdPricesPost200Response**](CompaniesCompanyIdAssetsAssetIdPricesPost200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdPricesGet

> CompaniesCompanyIdAssetsAssetIdPricesGet200Response CompaniesCompanyIdAssetsAssetIdPricesGet(ctx, companyId, assetId).Execute()

Get all price history records for a specific asset

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spartera-com/spartera-go-sdk"
)

func main() {
	companyId := "companyId_example" // string | 
	assetId := "assetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesGet(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdPricesGet`: CompaniesCompanyIdAssetsAssetIdPricesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdPricesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdAssetsAssetIdPricesGet200Response**](CompaniesCompanyIdAssetsAssetIdPricesGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdPricesPost

> CompaniesCompanyIdAssetsAssetIdPricesPost200Response CompaniesCompanyIdAssetsAssetIdPricesPost(ctx, companyId, assetId).AssetPriceHistoryInput(assetPriceHistoryInput).Execute()

Create a new price history record for an asset

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spartera-com/spartera-go-sdk"
)

func main() {
	companyId := "companyId_example" // string | 
	assetId := "assetId_example" // string | 
	assetPriceHistoryInput := *openapiclient.NewAssetPriceHistoryInput("AssetId_example") // AssetPriceHistoryInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesPost(context.Background(), companyId, assetId).AssetPriceHistoryInput(assetPriceHistoryInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdPricesPost`: CompaniesCompanyIdAssetsAssetIdPricesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.CompaniesCompanyIdAssetsAssetIdPricesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdPricesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assetPriceHistoryInput** | [**AssetPriceHistoryInput**](AssetPriceHistoryInput.md) |  | 

### Return type

[**CompaniesCompanyIdAssetsAssetIdPricesPost200Response**](CompaniesCompanyIdAssetsAssetIdPricesPost200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


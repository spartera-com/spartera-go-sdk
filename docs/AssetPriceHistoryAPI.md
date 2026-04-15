# \AssetPriceHistoryAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssetPriceHistory**](AssetPriceHistoryAPI.md#CreateAssetPriceHistory) | **Post** /companies/{company_id}/assets/{asset_id}/prices | Create a new price history record for an asset
[**CreateAssetPriceHistoryPricesCalculateCredits**](AssetPriceHistoryAPI.md#CreateAssetPriceHistoryPricesCalculateCredits) | **Post** /companies/{company_id}/assets/{asset_id}/prices/calculate_credits | Calculate the credit equivalent for a given USD price without saving
[**CreateAssetPriceHistoryPricesDiscount**](AssetPriceHistoryAPI.md#CreateAssetPriceHistoryPricesDiscount) | **Post** /companies/{company_id}/assets/{asset_id}/prices/discount | POST /companies/{company_id}/assets/{asset_id}/prices/discount
[**DeleteAssetPriceHistory**](AssetPriceHistoryAPI.md#DeleteAssetPriceHistory) | **Delete** /companies/{company_id}/assets/{asset_id}/prices/{aph_id} | Delete single price history record by ID
[**GetAssetPriceHistoryById**](AssetPriceHistoryAPI.md#GetAssetPriceHistoryById) | **Get** /companies/{company_id}/assets/{asset_id}/prices | Get all price history records for a specific asset
[**GetAssetPriceHistoryByIdAssetsPrices**](AssetPriceHistoryAPI.md#GetAssetPriceHistoryByIdAssetsPrices) | **Get** /companies/{company_id}/assets/{asset_id}/prices/{aph_id} | Get single price history record by ID
[**GetAssetPriceHistoryByIdPricesActive**](AssetPriceHistoryAPI.md#GetAssetPriceHistoryByIdPricesActive) | **Get** /companies/{company_id}/assets/{asset_id}/prices/active | Get the currently active price for an asset
[**UpdateAssetPriceHistory**](AssetPriceHistoryAPI.md#UpdateAssetPriceHistory) | **Patch** /companies/{company_id}/assets/{asset_id}/prices/{aph_id} | Update an existing price history record by ID



## CreateAssetPriceHistory

> CreateAssetPriceHistory200Response CreateAssetPriceHistory(ctx, companyId, assetId).AssetPriceHistoryInput(assetPriceHistoryInput).Execute()

Create a new price history record for an asset

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	companyId := "companyId_example" // string | Unique identifier for the Company
	assetId := "assetId_example" // string | Unique identifier for the Asset
	assetPriceHistoryInput := *openapiclient.NewAssetPriceHistoryInput() // AssetPriceHistoryInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.CreateAssetPriceHistory(context.Background(), companyId, assetId).AssetPriceHistoryInput(assetPriceHistoryInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.CreateAssetPriceHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetPriceHistory`: CreateAssetPriceHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.CreateAssetPriceHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetPriceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assetPriceHistoryInput** | [**AssetPriceHistoryInput**](AssetPriceHistoryInput.md) |  | 

### Return type

[**CreateAssetPriceHistory200Response**](CreateAssetPriceHistory200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetPriceHistoryPricesCalculateCredits

> CreateAssetPriceHistory200Response CreateAssetPriceHistoryPricesCalculateCredits(ctx, companyId, assetId).AssetPriceHistoryInput(assetPriceHistoryInput).Execute()

Calculate the credit equivalent for a given USD price without saving

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	companyId := "companyId_example" // string | Unique identifier for the Company
	assetId := "assetId_example" // string | Unique identifier for the Asset
	assetPriceHistoryInput := *openapiclient.NewAssetPriceHistoryInput() // AssetPriceHistoryInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.CreateAssetPriceHistoryPricesCalculateCredits(context.Background(), companyId, assetId).AssetPriceHistoryInput(assetPriceHistoryInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.CreateAssetPriceHistoryPricesCalculateCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetPriceHistoryPricesCalculateCredits`: CreateAssetPriceHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.CreateAssetPriceHistoryPricesCalculateCredits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetPriceHistoryPricesCalculateCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assetPriceHistoryInput** | [**AssetPriceHistoryInput**](AssetPriceHistoryInput.md) |  | 

### Return type

[**CreateAssetPriceHistory200Response**](CreateAssetPriceHistory200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetPriceHistoryPricesDiscount

> CreateAssetPriceHistory200Response CreateAssetPriceHistoryPricesDiscount(ctx, companyId, assetId).AssetPriceHistoryInput(assetPriceHistoryInput).Execute()

POST /companies/{company_id}/assets/{asset_id}/prices/discount

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	companyId := "companyId_example" // string | Unique identifier for the Company
	assetId := "assetId_example" // string | Unique identifier for the Asset
	assetPriceHistoryInput := *openapiclient.NewAssetPriceHistoryInput() // AssetPriceHistoryInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.CreateAssetPriceHistoryPricesDiscount(context.Background(), companyId, assetId).AssetPriceHistoryInput(assetPriceHistoryInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.CreateAssetPriceHistoryPricesDiscount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetPriceHistoryPricesDiscount`: CreateAssetPriceHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.CreateAssetPriceHistoryPricesDiscount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetPriceHistoryPricesDiscountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assetPriceHistoryInput** | [**AssetPriceHistoryInput**](AssetPriceHistoryInput.md) |  | 

### Return type

[**CreateAssetPriceHistory200Response**](CreateAssetPriceHistory200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssetPriceHistory

> DeleteAssetPriceHistory200Response DeleteAssetPriceHistory(ctx, companyId, assetId, aphId).Execute()

Delete single price history record by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	companyId := "companyId_example" // string | Unique identifier for the Company
	assetId := "assetId_example" // string | Unique identifier for the Asset
	aphId := "aphId_example" // string | Unique identifier for the Aph

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.DeleteAssetPriceHistory(context.Background(), companyId, assetId, aphId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.DeleteAssetPriceHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAssetPriceHistory`: DeleteAssetPriceHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.DeleteAssetPriceHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 
**aphId** | **string** | Unique identifier for the Aph | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetPriceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeleteAssetPriceHistory200Response**](DeleteAssetPriceHistory200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetPriceHistoryById

> GetAssetPriceHistoryById200Response GetAssetPriceHistoryById(ctx, companyId, assetId).Execute()

Get all price history records for a specific asset

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	companyId := "companyId_example" // string | Unique identifier for the Company
	assetId := "assetId_example" // string | Unique identifier for the Asset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.GetAssetPriceHistoryById(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.GetAssetPriceHistoryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetPriceHistoryById`: GetAssetPriceHistoryById200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.GetAssetPriceHistoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetPriceHistoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAssetPriceHistoryById200Response**](GetAssetPriceHistoryById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetPriceHistoryByIdAssetsPrices

> GetAssetPriceHistoryById200Response GetAssetPriceHistoryByIdAssetsPrices(ctx, companyId, assetId, aphId).Execute()

Get single price history record by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	companyId := "companyId_example" // string | Unique identifier for the Company
	assetId := "assetId_example" // string | Unique identifier for the Asset
	aphId := "aphId_example" // string | Unique identifier for the Aph

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.GetAssetPriceHistoryByIdAssetsPrices(context.Background(), companyId, assetId, aphId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.GetAssetPriceHistoryByIdAssetsPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetPriceHistoryByIdAssetsPrices`: GetAssetPriceHistoryById200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.GetAssetPriceHistoryByIdAssetsPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 
**aphId** | **string** | Unique identifier for the Aph | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetPriceHistoryByIdAssetsPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetAssetPriceHistoryById200Response**](GetAssetPriceHistoryById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetPriceHistoryByIdPricesActive

> GetAssetPriceHistoryById200Response GetAssetPriceHistoryByIdPricesActive(ctx, companyId, assetId).Execute()

Get the currently active price for an asset

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	companyId := "companyId_example" // string | Unique identifier for the Company
	assetId := "assetId_example" // string | Unique identifier for the Asset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.GetAssetPriceHistoryByIdPricesActive(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.GetAssetPriceHistoryByIdPricesActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetPriceHistoryByIdPricesActive`: GetAssetPriceHistoryById200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.GetAssetPriceHistoryByIdPricesActive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetPriceHistoryByIdPricesActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAssetPriceHistoryById200Response**](GetAssetPriceHistoryById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetPriceHistory

> UpdateAssetPriceHistory200Response UpdateAssetPriceHistory(ctx, companyId, assetId, aphId).AssetPriceHistoryUpdate(assetPriceHistoryUpdate).Execute()

Update an existing price history record by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	companyId := "companyId_example" // string | Unique identifier for the Company
	assetId := "assetId_example" // string | Unique identifier for the Asset
	aphId := "aphId_example" // string | Unique identifier for the Aph
	assetPriceHistoryUpdate := *openapiclient.NewAssetPriceHistoryUpdate() // AssetPriceHistoryUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetPriceHistoryAPI.UpdateAssetPriceHistory(context.Background(), companyId, assetId, aphId).AssetPriceHistoryUpdate(assetPriceHistoryUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetPriceHistoryAPI.UpdateAssetPriceHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAssetPriceHistory`: UpdateAssetPriceHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetPriceHistoryAPI.UpdateAssetPriceHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 
**aphId** | **string** | Unique identifier for the Aph | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetPriceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **assetPriceHistoryUpdate** | [**AssetPriceHistoryUpdate**](AssetPriceHistoryUpdate.md) |  | 

### Return type

[**UpdateAssetPriceHistory200Response**](UpdateAssetPriceHistory200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


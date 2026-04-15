# \AssetsAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssets**](AssetsAPI.md#CreateAssets) | **Post** /companies/{company_id}/assets | Create a new asset
[**CreateAssetsAnalyze**](AssetsAPI.md#CreateAssetsAnalyze) | **Post** /analyze/{company_handle}/assets/{asset_slug} | Process (analyze) an asset with dynamic rate limiting applied via decorator.
[**CreateAssetsScanColumn**](AssetsAPI.md#CreateAssetsScanColumn) | **Post** /companies/{company_id}/assets/{asset_id}/scan_column | Scan a column in the asset&#39;s table to retrieve distinct values      Request Body:         column (str): Column name to scan         limit (int, optional): Maximum distinct values to return (default 1000, max 5000)      Returns:         Flask Response with scan results
[**CreateAssetsTest**](AssetsAPI.md#CreateAssetsTest) | **Post** /companies/{company_id}/assets/{asset_id}/test | POST /companies/{company_id}/assets/{asset_id}/test
[**DeleteAssets**](AssetsAPI.md#DeleteAssets) | **Delete** /companies/{company_id}/assets/{asset_id} | Delete single asset by ID
[**GetAssetsById**](AssetsAPI.md#GetAssetsById) | **Get** /companies/{company_id}/assets/{asset_id} | Get single asset by ID
[**GetAssetsById2**](AssetsAPI.md#GetAssetsById2) | **Get** /companies/{company_id}/assets/{asset_id}/statistics | Get statistics for a specific asset (public endpoint)
[**GetAssetsByIdAnalyze**](AssetsAPI.md#GetAssetsByIdAnalyze) | **Get** /analyze/{company_handle}/assets/{asset_slug} | Process (analyze) an asset with dynamic rate limiting applied via decorator.
[**GetAssetsByIdInfoschema**](AssetsAPI.md#GetAssetsByIdInfoschema) | **Get** /companies/{company_id}/assets/{asset_id}/infoschema | Get the information schema for a specific asset&#39;s table
[**GetAssetsByIdInfoschemaSave**](AssetsAPI.md#GetAssetsByIdInfoschemaSave) | **Get** /companies/{company_id}/assets/{asset_id}/infoschema/save | Retrieve and save an asset&#39;s information schema
[**GetAssetsByIdPredictedPrice**](AssetsAPI.md#GetAssetsByIdPredictedPrice) | **Get** /companies/{company_id}/assets/{asset_id}/predicted_price | Get AI-predicted pricing for a specific asset
[**GetAssetsByIdStatistics**](AssetsAPI.md#GetAssetsByIdStatistics) | **Get** /companies/{company_id}/assets/statistics | Get statistics for all assets the user has access to
[**GetAssetsByIdTest**](AssetsAPI.md#GetAssetsByIdTest) | **Get** /companies/{company_id}/assets/{asset_id}/test | GET /companies/{company_id}/assets/{asset_id}/test
[**ListAssets**](AssetsAPI.md#ListAssets) | **Get** /companies/{company_id}/assets | Get all assets for a specific company
[**ListAssetsSearch**](AssetsAPI.md#ListAssetsSearch) | **Get** /companies/{company_id}/assets/search | Search and filter assets with advanced options      Query Parameters:         q: Search query string         category: Filter by category         sport: Filter by sport tag         sort_by: Sort field (name|recent|popular|trending)         limit: Number of results (default 20, max 100)         offset: Offset for pagination         include_recommended: Include recommendations (true/false)         include_schema: Include asset_schema in response (true/false, default false)
[**UpdateAssets**](AssetsAPI.md#UpdateAssets) | **Patch** /companies/{company_id}/assets/{asset_id} | Update an existing asset by ID



## CreateAssets

> CreateAssetsAnalyze200Response CreateAssets(ctx, companyId).AssetsInput(assetsInput).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Create a new asset

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
	assetsInput := *openapiclient.NewAssetsInput("company_id_abc123", "Example Name", "MANUAL") // AssetsInput | 
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	limit := int32(56) // int32 | Number of items per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	sortOrder := "sortOrder_example" // string | Sort order (ascending or descending) (optional) (default to "desc")
	search := "search_example" // string | Search term to filter results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CreateAssets(context.Background(), companyId).AssetsInput(assetsInput).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CreateAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssets`: CreateAssetsAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CreateAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetsInput** | [**AssetsInput**](AssetsInput.md) |  | 
 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**CreateAssetsAnalyze200Response**](CreateAssetsAnalyze200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetsAnalyze

> CreateAssetsAnalyze200Response CreateAssetsAnalyze(ctx, assetSlug, companyHandle).AssetsInput(assetsInput).Execute()

Process (analyze) an asset with dynamic rate limiting applied via decorator.

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
	assetSlug := "assetSlug_example" // string | URL-friendly slug for the Asset
	companyHandle := "companyHandle_example" // string | Human-readable handle for the Company
	assetsInput := *openapiclient.NewAssetsInput("company_id_abc123", "Example Name", "MANUAL") // AssetsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CreateAssetsAnalyze(context.Background(), assetSlug, companyHandle).AssetsInput(assetsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CreateAssetsAnalyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetsAnalyze`: CreateAssetsAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CreateAssetsAnalyze`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetSlug** | **string** | URL-friendly slug for the Asset | 
**companyHandle** | **string** | Human-readable handle for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetsAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assetsInput** | [**AssetsInput**](AssetsInput.md) |  | 

### Return type

[**CreateAssetsAnalyze200Response**](CreateAssetsAnalyze200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetsScanColumn

> CreateAssetsAnalyze200Response CreateAssetsScanColumn(ctx, companyId, assetId).AssetsInput(assetsInput).Execute()

Scan a column in the asset's table to retrieve distinct values      Request Body:         column (str): Column name to scan         limit (int, optional): Maximum distinct values to return (default 1000, max 5000)      Returns:         Flask Response with scan results

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
	assetsInput := *openapiclient.NewAssetsInput("company_id_abc123", "Example Name", "MANUAL") // AssetsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CreateAssetsScanColumn(context.Background(), companyId, assetId).AssetsInput(assetsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CreateAssetsScanColumn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetsScanColumn`: CreateAssetsAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CreateAssetsScanColumn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetsScanColumnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assetsInput** | [**AssetsInput**](AssetsInput.md) |  | 

### Return type

[**CreateAssetsAnalyze200Response**](CreateAssetsAnalyze200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetsTest

> CreateAssetsAnalyze200Response CreateAssetsTest(ctx, companyId, assetId).AssetsInput(assetsInput).Execute()

POST /companies/{company_id}/assets/{asset_id}/test

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
	assetsInput := *openapiclient.NewAssetsInput("company_id_abc123", "Example Name", "MANUAL") // AssetsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CreateAssetsTest(context.Background(), companyId, assetId).AssetsInput(assetsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CreateAssetsTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetsTest`: CreateAssetsAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CreateAssetsTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetsTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assetsInput** | [**AssetsInput**](AssetsInput.md) |  | 

### Return type

[**CreateAssetsAnalyze200Response**](CreateAssetsAnalyze200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssets

> DeleteAssets200Response DeleteAssets(ctx, companyId, assetId).Execute()

Delete single asset by ID

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
	resp, r, err := apiClient.AssetsAPI.DeleteAssets(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.DeleteAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAssets`: DeleteAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.DeleteAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteAssets200Response**](DeleteAssets200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetsById

> GetAssetsByIdAnalyze200Response GetAssetsById(ctx, companyId, assetId).Execute()

Get single asset by ID

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
	resp, r, err := apiClient.AssetsAPI.GetAssetsById(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAssetsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetsById`: GetAssetsByIdAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetAssetsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAssetsByIdAnalyze200Response**](GetAssetsByIdAnalyze200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetsById2

> GetAssetsByIdAnalyze200Response GetAssetsById2(ctx, companyId, assetId).Execute()

Get statistics for a specific asset (public endpoint)

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
	resp, r, err := apiClient.AssetsAPI.GetAssetsById2(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAssetsById2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetsById2`: GetAssetsByIdAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetAssetsById2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsById2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAssetsByIdAnalyze200Response**](GetAssetsByIdAnalyze200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetsByIdAnalyze

> GetAssetsByIdAnalyze200Response GetAssetsByIdAnalyze(ctx, assetSlug, companyHandle).Execute()

Process (analyze) an asset with dynamic rate limiting applied via decorator.

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
	assetSlug := "assetSlug_example" // string | URL-friendly slug for the Asset
	companyHandle := "companyHandle_example" // string | Human-readable handle for the Company

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.GetAssetsByIdAnalyze(context.Background(), assetSlug, companyHandle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAssetsByIdAnalyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetsByIdAnalyze`: GetAssetsByIdAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetAssetsByIdAnalyze`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetSlug** | **string** | URL-friendly slug for the Asset | 
**companyHandle** | **string** | Human-readable handle for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsByIdAnalyzeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAssetsByIdAnalyze200Response**](GetAssetsByIdAnalyze200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetsByIdInfoschema

> GetAssetsByIdAnalyze200Response GetAssetsByIdInfoschema(ctx, companyId, assetId).Execute()

Get the information schema for a specific asset's table

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
	resp, r, err := apiClient.AssetsAPI.GetAssetsByIdInfoschema(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAssetsByIdInfoschema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetsByIdInfoschema`: GetAssetsByIdAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetAssetsByIdInfoschema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsByIdInfoschemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAssetsByIdAnalyze200Response**](GetAssetsByIdAnalyze200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetsByIdInfoschemaSave

> GetAssetsByIdAnalyze200Response GetAssetsByIdInfoschemaSave(ctx, companyId, assetId).Execute()

Retrieve and save an asset's information schema

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
	resp, r, err := apiClient.AssetsAPI.GetAssetsByIdInfoschemaSave(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAssetsByIdInfoschemaSave``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetsByIdInfoschemaSave`: GetAssetsByIdAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetAssetsByIdInfoschemaSave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsByIdInfoschemaSaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAssetsByIdAnalyze200Response**](GetAssetsByIdAnalyze200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetsByIdPredictedPrice

> GetAssetsByIdAnalyze200Response GetAssetsByIdPredictedPrice(ctx, companyId, assetId).Execute()

Get AI-predicted pricing for a specific asset

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
	resp, r, err := apiClient.AssetsAPI.GetAssetsByIdPredictedPrice(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAssetsByIdPredictedPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetsByIdPredictedPrice`: GetAssetsByIdAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetAssetsByIdPredictedPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsByIdPredictedPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAssetsByIdAnalyze200Response**](GetAssetsByIdAnalyze200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetsByIdStatistics

> GetAssetsByIdAnalyze200Response GetAssetsByIdStatistics(ctx, companyId).Execute()

Get statistics for all assets the user has access to

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.GetAssetsByIdStatistics(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAssetsByIdStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetsByIdStatistics`: GetAssetsByIdAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetAssetsByIdStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsByIdStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAssetsByIdAnalyze200Response**](GetAssetsByIdAnalyze200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetsByIdTest

> GetAssetsByIdAnalyze200Response GetAssetsByIdTest(ctx, companyId, assetId).Execute()

GET /companies/{company_id}/assets/{asset_id}/test

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
	resp, r, err := apiClient.AssetsAPI.GetAssetsByIdTest(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetAssetsByIdTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetsByIdTest`: GetAssetsByIdAnalyze200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetAssetsByIdTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsByIdTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAssetsByIdAnalyze200Response**](GetAssetsByIdAnalyze200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssets

> ListAssets200Response ListAssets(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get all assets for a specific company

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
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	limit := int32(56) // int32 | Number of items per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	sortOrder := "sortOrder_example" // string | Sort order (ascending or descending) (optional) (default to "desc")
	search := "search_example" // string | Search term to filter results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.ListAssets(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.ListAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssets`: ListAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.ListAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListAssets200Response**](ListAssets200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetsSearch

> ListAssets200Response ListAssetsSearch(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Search and filter assets with advanced options      Query Parameters:         q: Search query string         category: Filter by category         sport: Filter by sport tag         sort_by: Sort field (name|recent|popular|trending)         limit: Number of results (default 20, max 100)         offset: Offset for pagination         include_recommended: Include recommendations (true/false)         include_schema: Include asset_schema in response (true/false, default false)

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
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	limit := int32(56) // int32 | Number of items per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	sortOrder := "sortOrder_example" // string | Sort order (ascending or descending) (optional) (default to "desc")
	search := "search_example" // string | Search term to filter results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.ListAssetsSearch(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.ListAssetsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssetsSearch`: ListAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.ListAssetsSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListAssets200Response**](ListAssets200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssets

> UpdateAssets200Response UpdateAssets(ctx, companyId, assetId).AssetsUpdate(assetsUpdate).Execute()

Update an existing asset by ID

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
	assetsUpdate := *openapiclient.NewAssetsUpdate() // AssetsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.UpdateAssets(context.Background(), companyId, assetId).AssetsUpdate(assetsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.UpdateAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAssets`: UpdateAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.UpdateAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assetsUpdate** | [**AssetsUpdate**](AssetsUpdate.md) |  | 

### Return type

[**UpdateAssets200Response**](UpdateAssets200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


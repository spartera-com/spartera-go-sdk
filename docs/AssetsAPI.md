# \AssetsAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeCompanyHandleAssetsAssetSlugGet**](AssetsAPI.md#AnalyzeCompanyHandleAssetsAssetSlugGet) | **Get** /analyze/{company_handle}/assets/{asset_slug} | Process assets route that handles both owned and purchased assets.             Minimal route function that passes all logic to crudder.process_asset              Args:                 asset_path: The path after /analyze/ containing asset information                 company_handle: The subdomain from Flask&#39;s routing (if available)
[**CompaniesCompanyIdAssetsAssetIdDelete**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdDelete) | **Delete** /companies/{company_id}/assets/{asset_id} | Delete single asset by ID
[**CompaniesCompanyIdAssetsAssetIdGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdGet) | **Get** /companies/{company_id}/assets/{asset_id} | Get single asset by ID
[**CompaniesCompanyIdAssetsAssetIdInfoschemaGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdInfoschemaGet) | **Get** /companies/{company_id}/assets/{asset_id}/infoschema | Get the information schema for a specific asset&#39;s table
[**CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet) | **Get** /companies/{company_id}/assets/{asset_id}/infoschema/save | Get the information schema for a specific asset and save it to the asset&#39;s asset_schema field
[**CompaniesCompanyIdAssetsAssetIdPatch**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdPatch) | **Patch** /companies/{company_id}/assets/{asset_id} | Update an existing asset by ID
[**CompaniesCompanyIdAssetsAssetIdPredictedPriceGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdPredictedPriceGet) | **Get** /companies/{company_id}/assets/{asset_id}/predicted_price | Get AI-predicted pricing for a specific asset
[**CompaniesCompanyIdAssetsAssetIdRecommendationsExplainGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdRecommendationsExplainGet) | **Get** /companies/{company_id}/assets/{asset_id}/recommendations/explain | Get detailed explanation of how asset recommendations are calculated for debugging purposes.
[**CompaniesCompanyIdAssetsAssetIdRecommendationsGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdRecommendationsGet) | **Get** /companies/{company_id}/assets/{asset_id}/recommendations | Get asset recommendations for a specific asset using heuristic waterfall matching     Returns list of similar assets based on industry, company, connection, tags, etc.      Query Parameters:     - limit: Number of recommendations to return (default: 12, max: 50)     - min_score: Minimum similarity score threshold (default: 0.1)     - include_details: Include component similarity scores (default: false)
[**CompaniesCompanyIdAssetsAssetIdStatisticsGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdStatisticsGet) | **Get** /companies/{company_id}/assets/{asset_id}/statistics | Get statistics for a specific asset (public endpoint)
[**CompaniesCompanyIdAssetsAssetIdTestGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdTestGet) | **Get** /companies/{company_id}/assets/{asset_id}/test | Test out an Asset (on a subset of data)
[**CompaniesCompanyIdAssetsGet**](AssetsAPI.md#CompaniesCompanyIdAssetsGet) | **Get** /companies/{company_id}/assets | Get all assets for a specific company
[**CompaniesCompanyIdAssetsPost**](AssetsAPI.md#CompaniesCompanyIdAssetsPost) | **Post** /companies/{company_id}/assets | Create a new asset
[**CompaniesCompanyIdAssetsRecommendationsBulkPost**](AssetsAPI.md#CompaniesCompanyIdAssetsRecommendationsBulkPost) | **Post** /companies/{company_id}/assets/recommendations/bulk | Get recommendations for multiple assets in a single request. Useful for pre-loading recommendations.
[**CompaniesCompanyIdAssetsRecommendationsHealthGet**](AssetsAPI.md#CompaniesCompanyIdAssetsRecommendationsHealthGet) | **Get** /companies/{company_id}/assets/recommendations/health | Health check for the recommendations system with sample data and performance metrics.
[**CompaniesCompanyIdAssetsStatisticsGet**](AssetsAPI.md#CompaniesCompanyIdAssetsStatisticsGet) | **Get** /companies/{company_id}/assets/statistics | Get statistics for all assets the user has access to



## AnalyzeCompanyHandleAssetsAssetSlugGet

> map[string]interface{} AnalyzeCompanyHandleAssetsAssetSlugGet(ctx, companyHandle, assetSlug).Execute()

Process assets route that handles both owned and purchased assets.             Minimal route function that passes all logic to crudder.process_asset              Args:                 asset_path: The path after /analyze/ containing asset information                 company_handle: The subdomain from Flask's routing (if available)

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
	companyHandle := "companyHandle_example" // string | 
	assetSlug := "assetSlug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.AnalyzeCompanyHandleAssetsAssetSlugGet(context.Background(), companyHandle, assetSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AnalyzeCompanyHandleAssetsAssetSlugGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyzeCompanyHandleAssetsAssetSlugGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AnalyzeCompanyHandleAssetsAssetSlugGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyHandle** | **string** |  | 
**assetSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyzeCompanyHandleAssetsAssetSlugGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdDelete

> map[string]interface{} CompaniesCompanyIdAssetsAssetIdDelete(ctx, companyId, assetId).Execute()

Delete single asset by ID

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
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsAssetIdDelete(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsAssetIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsAssetIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdGet

> map[string]interface{} CompaniesCompanyIdAssetsAssetIdGet(ctx, companyId, assetId).Execute()

Get single asset by ID

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
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsAssetIdGet(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsAssetIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsAssetIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdInfoschemaGet

> map[string]interface{} CompaniesCompanyIdAssetsAssetIdInfoschemaGet(ctx, companyId, assetId).Execute()

Get the information schema for a specific asset's table

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
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsAssetIdInfoschemaGet(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsAssetIdInfoschemaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdInfoschemaGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsAssetIdInfoschemaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdInfoschemaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet

> map[string]interface{} CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet(ctx, companyId, assetId).Execute()

Get the information schema for a specific asset and save it to the asset's asset_schema field

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
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdInfoschemaSaveGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdPatch

> map[string]interface{} CompaniesCompanyIdAssetsAssetIdPatch(ctx, companyId, assetId).Asset(asset).Execute()

Update an existing asset by ID

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
	asset := *openapiclient.NewAsset("CompanyId_example", "Name_example", "Source_example") // Asset | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsAssetIdPatch(context.Background(), companyId, assetId).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsAssetIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdPatch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsAssetIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asset** | [**Asset**](Asset.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdPredictedPriceGet

> map[string]interface{} CompaniesCompanyIdAssetsAssetIdPredictedPriceGet(ctx, companyId, assetId).Execute()

Get AI-predicted pricing for a specific asset

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
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsAssetIdPredictedPriceGet(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsAssetIdPredictedPriceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdPredictedPriceGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsAssetIdPredictedPriceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdPredictedPriceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdRecommendationsExplainGet

> map[string]interface{} CompaniesCompanyIdAssetsAssetIdRecommendationsExplainGet(ctx, companyId, assetId).Execute()

Get detailed explanation of how asset recommendations are calculated for debugging purposes.

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
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsAssetIdRecommendationsExplainGet(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsAssetIdRecommendationsExplainGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdRecommendationsExplainGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsAssetIdRecommendationsExplainGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdRecommendationsExplainGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdRecommendationsGet

> map[string]interface{} CompaniesCompanyIdAssetsAssetIdRecommendationsGet(ctx, companyId, assetId).Limit(limit).MinScore(minScore).IncludeDetails(includeDetails).Execute()

Get asset recommendations for a specific asset using heuristic waterfall matching     Returns list of similar assets based on industry, company, connection, tags, etc.      Query Parameters:     - limit: Number of recommendations to return (default: 12, max: 50)     - min_score: Minimum similarity score threshold (default: 0.1)     - include_details: Include component similarity scores (default: false)

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
	limit := "limit_example" // string |  (optional)
	minScore := "minScore_example" // string |  (optional)
	includeDetails := "includeDetails_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsAssetIdRecommendationsGet(context.Background(), companyId, assetId).Limit(limit).MinScore(minScore).IncludeDetails(includeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsAssetIdRecommendationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdRecommendationsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsAssetIdRecommendationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdRecommendationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **string** |  | 
 **minScore** | **string** |  | 
 **includeDetails** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdStatisticsGet

> map[string]interface{} CompaniesCompanyIdAssetsAssetIdStatisticsGet(ctx, companyId, assetId).Execute()

Get statistics for a specific asset (public endpoint)

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
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsAssetIdStatisticsGet(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsAssetIdStatisticsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdStatisticsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsAssetIdStatisticsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdStatisticsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdTestGet

> map[string]interface{} CompaniesCompanyIdAssetsAssetIdTestGet(ctx, companyId, assetId).Execute()

Test out an Asset (on a subset of data)

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
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsAssetIdTestGet(context.Background(), companyId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsAssetIdTestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdTestGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsAssetIdTestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsAssetIdTestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsGet

> map[string]interface{} CompaniesCompanyIdAssetsGet(ctx, companyId).Execute()

Get all assets for a specific company

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsPost

> map[string]interface{} CompaniesCompanyIdAssetsPost(ctx, companyId).Asset(asset).Execute()

Create a new asset

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
	asset := *openapiclient.NewAsset("CompanyId_example", "Name_example", "Source_example") // Asset | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsPost(context.Background(), companyId).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asset** | [**Asset**](Asset.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsRecommendationsBulkPost

> map[string]interface{} CompaniesCompanyIdAssetsRecommendationsBulkPost(ctx, companyId).Asset(asset).Execute()

Get recommendations for multiple assets in a single request. Useful for pre-loading recommendations.

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
	asset := *openapiclient.NewAsset("CompanyId_example", "Name_example", "Source_example") // Asset | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsRecommendationsBulkPost(context.Background(), companyId).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsRecommendationsBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsRecommendationsBulkPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsRecommendationsBulkPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsRecommendationsBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asset** | [**Asset**](Asset.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsRecommendationsHealthGet

> map[string]interface{} CompaniesCompanyIdAssetsRecommendationsHealthGet(ctx, companyId).Execute()

Health check for the recommendations system with sample data and performance metrics.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsRecommendationsHealthGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsRecommendationsHealthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsRecommendationsHealthGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsRecommendationsHealthGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsRecommendationsHealthGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsStatisticsGet

> map[string]interface{} CompaniesCompanyIdAssetsStatisticsGet(ctx, companyId).Execute()

Get statistics for all assets the user has access to

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsStatisticsGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsStatisticsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsStatisticsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CompaniesCompanyIdAssetsStatisticsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAssetsStatisticsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \AssetsAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeCompanyHandleAssetsAssetSlugGet**](AssetsAPI.md#AnalyzeCompanyHandleAssetsAssetSlugGet) | **Get** /analyze/{company_handle}/assets/{asset_slug} | Process (analyze) an asset.
[**CompaniesCompanyIdAssetsAssetIdDelete**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdDelete) | **Delete** /companies/{company_id}/assets/{asset_id} | Delete single asset by ID
[**CompaniesCompanyIdAssetsAssetIdGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdGet) | **Get** /companies/{company_id}/assets/{asset_id} | Get single asset by ID
[**CompaniesCompanyIdAssetsAssetIdInfoschemaGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdInfoschemaGet) | **Get** /companies/{company_id}/assets/{asset_id}/infoschema | Get the information schema for a specific asset&#39;s table
[**CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet) | **Get** /companies/{company_id}/assets/{asset_id}/infoschema/save | Retrieve and save an asset&#39;s information schema
[**CompaniesCompanyIdAssetsAssetIdPatch**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdPatch) | **Patch** /companies/{company_id}/assets/{asset_id} | Update an existing asset by ID
[**CompaniesCompanyIdAssetsAssetIdPredictedPriceGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdPredictedPriceGet) | **Get** /companies/{company_id}/assets/{asset_id}/predicted_price | Get AI-predicted pricing for a specific asset
[**CompaniesCompanyIdAssetsAssetIdStatisticsGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdStatisticsGet) | **Get** /companies/{company_id}/assets/{asset_id}/statistics | Get statistics for a specific asset (public endpoint)
[**CompaniesCompanyIdAssetsAssetIdTestGet**](AssetsAPI.md#CompaniesCompanyIdAssetsAssetIdTestGet) | **Get** /companies/{company_id}/assets/{asset_id}/test | Test out an Asset (on a subset of data)
[**CompaniesCompanyIdAssetsGet**](AssetsAPI.md#CompaniesCompanyIdAssetsGet) | **Get** /companies/{company_id}/assets | Get all assets for a specific company
[**CompaniesCompanyIdAssetsPost**](AssetsAPI.md#CompaniesCompanyIdAssetsPost) | **Post** /companies/{company_id}/assets | Create a new asset
[**CompaniesCompanyIdAssetsStatisticsGet**](AssetsAPI.md#CompaniesCompanyIdAssetsStatisticsGet) | **Get** /companies/{company_id}/assets/statistics | Get statistics for all assets the user has access to



## AnalyzeCompanyHandleAssetsAssetSlugGet

> AnalyzeCompanyHandleAssetsAssetSlugGet200Response AnalyzeCompanyHandleAssetsAssetSlugGet(ctx, assetSlug, companyHandle).Execute()

Process (analyze) an asset.

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
	assetSlug := "assetSlug_example" // string | 
	companyHandle := "companyHandle_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.AnalyzeCompanyHandleAssetsAssetSlugGet(context.Background(), assetSlug, companyHandle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.AnalyzeCompanyHandleAssetsAssetSlugGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyzeCompanyHandleAssetsAssetSlugGet`: AnalyzeCompanyHandleAssetsAssetSlugGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.AnalyzeCompanyHandleAssetsAssetSlugGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetSlug** | **string** |  | 
**companyHandle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyzeCompanyHandleAssetsAssetSlugGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AnalyzeCompanyHandleAssetsAssetSlugGet200Response**](AnalyzeCompanyHandleAssetsAssetSlugGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdDelete

> CompaniesCompanyIdAssetsAssetIdDelete200Response CompaniesCompanyIdAssetsAssetIdDelete(ctx, companyId, assetId).Execute()

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
	// response from `CompaniesCompanyIdAssetsAssetIdDelete`: CompaniesCompanyIdAssetsAssetIdDelete200Response
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

[**CompaniesCompanyIdAssetsAssetIdDelete200Response**](CompaniesCompanyIdAssetsAssetIdDelete200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdGet

> AnalyzeCompanyHandleAssetsAssetSlugGet200Response CompaniesCompanyIdAssetsAssetIdGet(ctx, companyId, assetId).Execute()

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
	// response from `CompaniesCompanyIdAssetsAssetIdGet`: AnalyzeCompanyHandleAssetsAssetSlugGet200Response
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

[**AnalyzeCompanyHandleAssetsAssetSlugGet200Response**](AnalyzeCompanyHandleAssetsAssetSlugGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdInfoschemaGet

> CompaniesCompanyIdAssetsGet200Response CompaniesCompanyIdAssetsAssetIdInfoschemaGet(ctx, companyId, assetId).Execute()

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
	// response from `CompaniesCompanyIdAssetsAssetIdInfoschemaGet`: CompaniesCompanyIdAssetsGet200Response
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

[**CompaniesCompanyIdAssetsGet200Response**](CompaniesCompanyIdAssetsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet

> CompaniesCompanyIdAssetsGet200Response CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet(ctx, companyId, assetId).Execute()

Retrieve and save an asset's information schema

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
	// response from `CompaniesCompanyIdAssetsAssetIdInfoschemaSaveGet`: CompaniesCompanyIdAssetsGet200Response
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

[**CompaniesCompanyIdAssetsGet200Response**](CompaniesCompanyIdAssetsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdPatch

> CompaniesCompanyIdAssetsAssetIdPatch200Response CompaniesCompanyIdAssetsAssetIdPatch(ctx, companyId, assetId).AssetsUpdate(assetsUpdate).Execute()

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
	assetsUpdate := *openapiclient.NewAssetsUpdate() // AssetsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsAssetIdPatch(context.Background(), companyId, assetId).AssetsUpdate(assetsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsAssetIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsAssetIdPatch`: CompaniesCompanyIdAssetsAssetIdPatch200Response
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


 **assetsUpdate** | [**AssetsUpdate**](AssetsUpdate.md) |  | 

### Return type

[**CompaniesCompanyIdAssetsAssetIdPatch200Response**](CompaniesCompanyIdAssetsAssetIdPatch200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdPredictedPriceGet

> CompaniesCompanyIdAssetsGet200Response CompaniesCompanyIdAssetsAssetIdPredictedPriceGet(ctx, companyId, assetId).Execute()

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
	// response from `CompaniesCompanyIdAssetsAssetIdPredictedPriceGet`: CompaniesCompanyIdAssetsGet200Response
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

[**CompaniesCompanyIdAssetsGet200Response**](CompaniesCompanyIdAssetsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdStatisticsGet

> CompaniesCompanyIdAssetsGet200Response CompaniesCompanyIdAssetsAssetIdStatisticsGet(ctx, companyId, assetId).Execute()

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
	// response from `CompaniesCompanyIdAssetsAssetIdStatisticsGet`: CompaniesCompanyIdAssetsGet200Response
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

[**CompaniesCompanyIdAssetsGet200Response**](CompaniesCompanyIdAssetsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsAssetIdTestGet

> CompaniesCompanyIdAssetsGet200Response CompaniesCompanyIdAssetsAssetIdTestGet(ctx, companyId, assetId).Execute()

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
	// response from `CompaniesCompanyIdAssetsAssetIdTestGet`: CompaniesCompanyIdAssetsGet200Response
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

[**CompaniesCompanyIdAssetsGet200Response**](CompaniesCompanyIdAssetsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsGet

> CompaniesCompanyIdAssetsGet200Response CompaniesCompanyIdAssetsGet(ctx, companyId).Execute()

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
	// response from `CompaniesCompanyIdAssetsGet`: CompaniesCompanyIdAssetsGet200Response
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

[**CompaniesCompanyIdAssetsGet200Response**](CompaniesCompanyIdAssetsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsPost

> CompaniesCompanyIdAssetsPost200Response CompaniesCompanyIdAssetsPost(ctx, companyId).AssetsInput(assetsInput).Execute()

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
	assetsInput := *openapiclient.NewAssetsInput("CompanyId_example", "Name_example", "Source_example") // AssetsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CompaniesCompanyIdAssetsPost(context.Background(), companyId).AssetsInput(assetsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CompaniesCompanyIdAssetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAssetsPost`: CompaniesCompanyIdAssetsPost200Response
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

 **assetsInput** | [**AssetsInput**](AssetsInput.md) |  | 

### Return type

[**CompaniesCompanyIdAssetsPost200Response**](CompaniesCompanyIdAssetsPost200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdAssetsStatisticsGet

> CompaniesCompanyIdAssetsGet200Response CompaniesCompanyIdAssetsStatisticsGet(ctx, companyId).Execute()

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
	// response from `CompaniesCompanyIdAssetsStatisticsGet`: CompaniesCompanyIdAssetsGet200Response
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

[**CompaniesCompanyIdAssetsGet200Response**](CompaniesCompanyIdAssetsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


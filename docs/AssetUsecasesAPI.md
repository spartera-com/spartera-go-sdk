# \AssetUsecasesAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssetUsecasesById**](AssetUsecasesAPI.md#GetAssetUsecasesById) | **Get** /asset_usecases/{auc_id} | Get single asset use case by ID
[**ListAssetUsecases**](AssetUsecasesAPI.md#ListAssetUsecases) | **Get** /asset_usecases | Get a list of all asset use cases



## GetAssetUsecasesById

> GetAssetUsecasesById200Response GetAssetUsecasesById(ctx, aucId).Execute()

Get single asset use case by ID

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
	aucId := "aucId_example" // string | Unique identifier for the Auc

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetUsecasesAPI.GetAssetUsecasesById(context.Background(), aucId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetUsecasesAPI.GetAssetUsecasesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetUsecasesById`: GetAssetUsecasesById200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetUsecasesAPI.GetAssetUsecasesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aucId** | **string** | Unique identifier for the Auc | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetUsecasesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAssetUsecasesById200Response**](GetAssetUsecasesById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetUsecases

> ListAssetUsecases200Response ListAssetUsecases(ctx).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get a list of all asset use cases

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
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	limit := int32(56) // int32 | Number of items per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	sortOrder := "sortOrder_example" // string | Sort order (ascending or descending) (optional) (default to "desc")
	search := "search_example" // string | Search term to filter results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetUsecasesAPI.ListAssetUsecases(context.Background()).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetUsecasesAPI.ListAssetUsecases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssetUsecases`: ListAssetUsecases200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetUsecasesAPI.ListAssetUsecases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetUsecasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListAssetUsecases200Response**](ListAssetUsecases200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \StorageEnginesAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStorageEnginesById**](StorageEnginesAPI.md#GetStorageEnginesById) | **Get** /cloud-providers/{provider_id}/storage-engines/{engine_id} | Get single storage engine by ID
[**ListStorageEngines**](StorageEnginesAPI.md#ListStorageEngines) | **Get** /cloud-providers/{provider_id}/storage-engines | Get a list of all storage engines



## GetStorageEnginesById

> GetStorageEnginesById200Response GetStorageEnginesById(ctx, providerId, engineId).Execute()

Get single storage engine by ID

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
	providerId := "providerId_example" // string | Unique identifier for the Provider
	engineId := "engineId_example" // string | Unique identifier for the Engine

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageEnginesAPI.GetStorageEnginesById(context.Background(), providerId, engineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageEnginesAPI.GetStorageEnginesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageEnginesById`: GetStorageEnginesById200Response
	fmt.Fprintf(os.Stdout, "Response from `StorageEnginesAPI.GetStorageEnginesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | Unique identifier for the Provider | 
**engineId** | **string** | Unique identifier for the Engine | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageEnginesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetStorageEnginesById200Response**](GetStorageEnginesById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStorageEngines

> ListStorageEngines200Response ListStorageEngines(ctx, providerId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get a list of all storage engines

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
	providerId := "providerId_example" // string | Unique identifier for the Provider
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	limit := int32(56) // int32 | Number of items per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	sortOrder := "sortOrder_example" // string | Sort order (ascending or descending) (optional) (default to "desc")
	search := "search_example" // string | Search term to filter results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageEnginesAPI.ListStorageEngines(context.Background(), providerId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageEnginesAPI.ListStorageEngines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStorageEngines`: ListStorageEngines200Response
	fmt.Fprintf(os.Stdout, "Response from `StorageEnginesAPI.ListStorageEngines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | Unique identifier for the Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStorageEnginesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListStorageEngines200Response**](ListStorageEngines200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \JobFunctionsAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJobFunctionsById**](JobFunctionsAPI.md#GetJobFunctionsById) | **Get** /job-functions/{function_id} | Get single job function by ID
[**ListJobFunctions**](JobFunctionsAPI.md#ListJobFunctions) | **Get** /job-functions | Get a list of all job functions



## GetJobFunctionsById

> GetJobFunctionsById200Response GetJobFunctionsById(ctx, functionId).Execute()

Get single job function by ID

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
	functionId := "functionId_example" // string | Unique identifier for the Function

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobFunctionsAPI.GetJobFunctionsById(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobFunctionsAPI.GetJobFunctionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobFunctionsById`: GetJobFunctionsById200Response
	fmt.Fprintf(os.Stdout, "Response from `JobFunctionsAPI.GetJobFunctionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** | Unique identifier for the Function | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobFunctionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetJobFunctionsById200Response**](GetJobFunctionsById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJobFunctions

> ListJobFunctions200Response ListJobFunctions(ctx).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get a list of all job functions

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
	resp, r, err := apiClient.JobFunctionsAPI.ListJobFunctions(context.Background()).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobFunctionsAPI.ListJobFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListJobFunctions`: ListJobFunctions200Response
	fmt.Fprintf(os.Stdout, "Response from `JobFunctionsAPI.ListJobFunctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListJobFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListJobFunctions200Response**](ListJobFunctions200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


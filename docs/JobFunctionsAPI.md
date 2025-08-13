# \JobFunctionsAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobFunctionsFunctionIdGet**](JobFunctionsAPI.md#JobFunctionsFunctionIdGet) | **Get** /job-functions/{function_id} | Get single job function by ID
[**JobFunctionsGet**](JobFunctionsAPI.md#JobFunctionsGet) | **Get** /job-functions | Get a list of all job functions



## JobFunctionsFunctionIdGet

> JobFunctionsFunctionIdGet200Response JobFunctionsFunctionIdGet(ctx, functionId).Execute()

Get single job function by ID

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
	functionId := "functionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobFunctionsAPI.JobFunctionsFunctionIdGet(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobFunctionsAPI.JobFunctionsFunctionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobFunctionsFunctionIdGet`: JobFunctionsFunctionIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `JobFunctionsAPI.JobFunctionsFunctionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobFunctionsFunctionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobFunctionsFunctionIdGet200Response**](JobFunctionsFunctionIdGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobFunctionsGet

> JobFunctionsGet200Response JobFunctionsGet(ctx).Execute()

Get a list of all job functions

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobFunctionsAPI.JobFunctionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobFunctionsAPI.JobFunctionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobFunctionsGet`: JobFunctionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `JobFunctionsAPI.JobFunctionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiJobFunctionsGetRequest struct via the builder pattern


### Return type

[**JobFunctionsGet200Response**](JobFunctionsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


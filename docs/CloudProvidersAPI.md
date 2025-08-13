# \CloudProvidersAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudProvidersGet**](CloudProvidersAPI.md#CloudProvidersGet) | **Get** /cloud-providers | Get a list of all cloud providers
[**CloudProvidersProviderIdGet**](CloudProvidersAPI.md#CloudProvidersProviderIdGet) | **Get** /cloud-providers/{provider_id} | Get single cloud provider by ID



## CloudProvidersGet

> CloudProvidersGet200Response CloudProvidersGet(ctx).Execute()

Get a list of all cloud providers

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
	resp, r, err := apiClient.CloudProvidersAPI.CloudProvidersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.CloudProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProvidersGet`: CloudProvidersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.CloudProvidersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProvidersGetRequest struct via the builder pattern


### Return type

[**CloudProvidersGet200Response**](CloudProvidersGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudProvidersProviderIdGet

> CloudProvidersProviderIdGet200Response CloudProvidersProviderIdGet(ctx, providerId).Execute()

Get single cloud provider by ID

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
	providerId := "providerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProvidersAPI.CloudProvidersProviderIdGet(context.Background(), providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.CloudProvidersProviderIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProvidersProviderIdGet`: CloudProvidersProviderIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.CloudProvidersProviderIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProvidersProviderIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudProvidersProviderIdGet200Response**](CloudProvidersProviderIdGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


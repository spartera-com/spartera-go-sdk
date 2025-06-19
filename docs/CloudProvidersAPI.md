# \CloudProvidersAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudProvidersGet**](CloudProvidersAPI.md#CloudProvidersGet) | **Get** /cloud-providers | Get a list of all cloud providers
[**CloudProvidersPost**](CloudProvidersAPI.md#CloudProvidersPost) | **Post** /cloud-providers | Create single cloud provider
[**CloudProvidersProviderIdDelete**](CloudProvidersAPI.md#CloudProvidersProviderIdDelete) | **Delete** /cloud-providers/{provider_id} | Delete single cloud provider by ID
[**CloudProvidersProviderIdGet**](CloudProvidersAPI.md#CloudProvidersProviderIdGet) | **Get** /cloud-providers/{provider_id} | Get single cloud provider by ID
[**CloudProvidersProviderIdPatch**](CloudProvidersAPI.md#CloudProvidersProviderIdPatch) | **Patch** /cloud-providers/{provider_id} | Update an existing cloud provider by ID



## CloudProvidersGet

> map[string]interface{} CloudProvidersGet(ctx).Execute()

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
	// response from `CloudProvidersGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.CloudProvidersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProvidersGetRequest struct via the builder pattern


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


## CloudProvidersPost

> map[string]interface{} CloudProvidersPost(ctx).Execute()

Create single cloud provider

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
	resp, r, err := apiClient.CloudProvidersAPI.CloudProvidersPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.CloudProvidersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProvidersPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.CloudProvidersPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProvidersPostRequest struct via the builder pattern


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


## CloudProvidersProviderIdDelete

> map[string]interface{} CloudProvidersProviderIdDelete(ctx, providerId).Execute()

Delete single cloud provider by ID

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
	resp, r, err := apiClient.CloudProvidersAPI.CloudProvidersProviderIdDelete(context.Background(), providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.CloudProvidersProviderIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProvidersProviderIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.CloudProvidersProviderIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProvidersProviderIdDeleteRequest struct via the builder pattern


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


## CloudProvidersProviderIdGet

> map[string]interface{} CloudProvidersProviderIdGet(ctx, providerId).Execute()

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
	// response from `CloudProvidersProviderIdGet`: map[string]interface{}
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

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudProvidersProviderIdPatch

> map[string]interface{} CloudProvidersProviderIdPatch(ctx, providerId).Execute()

Update an existing cloud provider by ID

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
	resp, r, err := apiClient.CloudProvidersAPI.CloudProvidersProviderIdPatch(context.Background(), providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.CloudProvidersProviderIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProvidersProviderIdPatch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.CloudProvidersProviderIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProvidersProviderIdPatchRequest struct via the builder pattern


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


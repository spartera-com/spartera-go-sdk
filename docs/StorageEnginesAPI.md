# \StorageEnginesAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudProvidersProviderIdStorageEnginesEngineIdGet**](StorageEnginesAPI.md#CloudProvidersProviderIdStorageEnginesEngineIdGet) | **Get** /cloud-providers/{provider_id}/storage-engines/{engine_id} | Get single storage engine by ID
[**CloudProvidersProviderIdStorageEnginesGet**](StorageEnginesAPI.md#CloudProvidersProviderIdStorageEnginesGet) | **Get** /cloud-providers/{provider_id}/storage-engines | Get a list of all storage engines



## CloudProvidersProviderIdStorageEnginesEngineIdGet

> CompaniesCompanyIdApiKeysGet200Response CloudProvidersProviderIdStorageEnginesEngineIdGet(ctx, providerId, engineId).Execute()

Get single storage engine by ID

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
	engineId := "engineId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageEnginesAPI.CloudProvidersProviderIdStorageEnginesEngineIdGet(context.Background(), providerId, engineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageEnginesAPI.CloudProvidersProviderIdStorageEnginesEngineIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProvidersProviderIdStorageEnginesEngineIdGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `StorageEnginesAPI.CloudProvidersProviderIdStorageEnginesEngineIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**engineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProvidersProviderIdStorageEnginesEngineIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdApiKeysGet200Response**](CompaniesCompanyIdApiKeysGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudProvidersProviderIdStorageEnginesGet

> CompaniesCompanyIdApiKeysGet200Response CloudProvidersProviderIdStorageEnginesGet(ctx, providerId).Execute()

Get a list of all storage engines

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
	resp, r, err := apiClient.StorageEnginesAPI.CloudProvidersProviderIdStorageEnginesGet(context.Background(), providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageEnginesAPI.CloudProvidersProviderIdStorageEnginesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudProvidersProviderIdStorageEnginesGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `StorageEnginesAPI.CloudProvidersProviderIdStorageEnginesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudProvidersProviderIdStorageEnginesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompaniesCompanyIdApiKeysGet200Response**](CompaniesCompanyIdApiKeysGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


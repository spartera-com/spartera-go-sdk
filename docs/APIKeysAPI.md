# \APIKeysAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompaniesCompanyIdApiKeysApiKeyIdGet**](APIKeysAPI.md#CompaniesCompanyIdApiKeysApiKeyIdGet) | **Get** /companies/{company_id}/api-keys/{api_key_id} | Get single API key by ID
[**CompaniesCompanyIdApiKeysApiKeyIdPatch**](APIKeysAPI.md#CompaniesCompanyIdApiKeysApiKeyIdPatch) | **Patch** /companies/{company_id}/api-keys/{api_key_id} | Update an existing API key by ID
[**CompaniesCompanyIdApiKeysGet**](APIKeysAPI.md#CompaniesCompanyIdApiKeysGet) | **Get** /companies/{company_id}/api-keys | Get all API keys
[**CompaniesCompanyIdApiKeysPost**](APIKeysAPI.md#CompaniesCompanyIdApiKeysPost) | **Post** /companies/{company_id}/api-keys | Create single API key
[**CompaniesCompanyIdApiKeysTokenDelete**](APIKeysAPI.md#CompaniesCompanyIdApiKeysTokenDelete) | **Delete** /companies/{company_id}/api-keys/{token} | Delete single API key by token



## CompaniesCompanyIdApiKeysApiKeyIdGet

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdApiKeysApiKeyIdGet(ctx, companyId, apiKeyId).Execute()

Get single API key by ID

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
	apiKeyId := "apiKeyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdGet(context.Background(), companyId, apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdApiKeysApiKeyIdGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdApiKeysApiKeyIdGetRequest struct via the builder pattern


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


## CompaniesCompanyIdApiKeysApiKeyIdPatch

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdApiKeysApiKeyIdPatch(ctx, companyId, apiKeyId).Execute()

Update an existing API key by ID

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
	apiKeyId := "apiKeyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdPatch(context.Background(), companyId, apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdApiKeysApiKeyIdPatch`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdApiKeysApiKeyIdPatchRequest struct via the builder pattern


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


## CompaniesCompanyIdApiKeysGet

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdApiKeysGet(ctx, companyId).Execute()

Get all API keys

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
	resp, r, err := apiClient.APIKeysAPI.CompaniesCompanyIdApiKeysGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CompaniesCompanyIdApiKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdApiKeysGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.CompaniesCompanyIdApiKeysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdApiKeysGetRequest struct via the builder pattern


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


## CompaniesCompanyIdApiKeysPost

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdApiKeysPost(ctx, companyId).Execute()

Create single API key

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
	resp, r, err := apiClient.APIKeysAPI.CompaniesCompanyIdApiKeysPost(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CompaniesCompanyIdApiKeysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdApiKeysPost`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.CompaniesCompanyIdApiKeysPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdApiKeysPostRequest struct via the builder pattern


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


## CompaniesCompanyIdApiKeysTokenDelete

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdApiKeysTokenDelete(ctx, companyId, token).Execute()

Delete single API key by token

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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.CompaniesCompanyIdApiKeysTokenDelete(context.Background(), companyId, token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CompaniesCompanyIdApiKeysTokenDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdApiKeysTokenDelete`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.CompaniesCompanyIdApiKeysTokenDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdApiKeysTokenDeleteRequest struct via the builder pattern


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


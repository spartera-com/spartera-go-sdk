# \APIKeysAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompaniesCompanyIdApiKeysApiKeyIdDelete**](APIKeysAPI.md#CompaniesCompanyIdApiKeysApiKeyIdDelete) | **Delete** /companies/{company_id}/api-keys/{api_key_id} | Delete (deactivate) single API key by api_key_id.     Uses the api_key_id returned during creation for clean identification.     Fixed to use correct primary key lookup.
[**CompaniesCompanyIdApiKeysApiKeyIdGet**](APIKeysAPI.md#CompaniesCompanyIdApiKeysApiKeyIdGet) | **Get** /companies/{company_id}/api-keys/{api_key_id} | Get single API key by ID.     Returns masked API key for security (sk_****1234).
[**CompaniesCompanyIdApiKeysApiKeyIdPatch**](APIKeysAPI.md#CompaniesCompanyIdApiKeysApiKeyIdPatch) | **Patch** /companies/{company_id}/api-keys/{api_key_id} | Update an existing API key by ID.     Can update metadata like name, expiration, rate limits, etc.     Cannot update the actual key value (for security).
[**CompaniesCompanyIdApiKeysApiKeyIdRevokePost**](APIKeysAPI.md#CompaniesCompanyIdApiKeysApiKeyIdRevokePost) | **Post** /companies/{company_id}/api-keys/{api_key_id}/revoke | Explicitly revoke an API key with reason tracking.     This is different from delete as it includes revocation metadata.
[**CompaniesCompanyIdApiKeysApiKeyIdStatsGet**](APIKeysAPI.md#CompaniesCompanyIdApiKeysApiKeyIdStatsGet) | **Get** /companies/{company_id}/api-keys/{api_key_id}/stats | Get usage statistics for a specific API key.     Returns usage count, last used date, failed attempts, etc.
[**CompaniesCompanyIdApiKeysGet**](APIKeysAPI.md#CompaniesCompanyIdApiKeysGet) | **Get** /companies/{company_id}/api-keys | Get all API keys for a company.     Returns masked API keys for security (sk_****1234).
[**CompaniesCompanyIdApiKeysPost**](APIKeysAPI.md#CompaniesCompanyIdApiKeysPost) | **Post** /companies/{company_id}/api-keys | Create single API key.     Returns the actual sk_ key (only time it&#39;s exposed) and api_key_id for future operations.



## CompaniesCompanyIdApiKeysApiKeyIdDelete

> CompaniesCompanyIdApiKeysApiKeyIdDelete200Response CompaniesCompanyIdApiKeysApiKeyIdDelete(ctx, companyId, apiKeyId).Execute()

Delete (deactivate) single API key by api_key_id.     Uses the api_key_id returned during creation for clean identification.     Fixed to use correct primary key lookup.

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
	resp, r, err := apiClient.APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdDelete(context.Background(), companyId, apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdApiKeysApiKeyIdDelete`: CompaniesCompanyIdApiKeysApiKeyIdDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdApiKeysApiKeyIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdApiKeysApiKeyIdDelete200Response**](CompaniesCompanyIdApiKeysApiKeyIdDelete200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdApiKeysApiKeyIdGet

> CompaniesCompanyIdApiKeysApiKeyIdGet200Response CompaniesCompanyIdApiKeysApiKeyIdGet(ctx, companyId, apiKeyId).Execute()

Get single API key by ID.     Returns masked API key for security (sk_****1234).

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
	// response from `CompaniesCompanyIdApiKeysApiKeyIdGet`: CompaniesCompanyIdApiKeysApiKeyIdGet200Response
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

[**CompaniesCompanyIdApiKeysApiKeyIdGet200Response**](CompaniesCompanyIdApiKeysApiKeyIdGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdApiKeysApiKeyIdPatch

> CompaniesCompanyIdApiKeysApiKeyIdPatch200Response CompaniesCompanyIdApiKeysApiKeyIdPatch(ctx, companyId, apiKeyId).ApiKeysUpdate(apiKeysUpdate).Execute()

Update an existing API key by ID.     Can update metadata like name, expiration, rate limits, etc.     Cannot update the actual key value (for security).

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
	apiKeysUpdate := *openapiclient.NewApiKeysUpdate() // ApiKeysUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdPatch(context.Background(), companyId, apiKeyId).ApiKeysUpdate(apiKeysUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdApiKeysApiKeyIdPatch`: CompaniesCompanyIdApiKeysApiKeyIdPatch200Response
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


 **apiKeysUpdate** | [**ApiKeysUpdate**](ApiKeysUpdate.md) |  | 

### Return type

[**CompaniesCompanyIdApiKeysApiKeyIdPatch200Response**](CompaniesCompanyIdApiKeysApiKeyIdPatch200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdApiKeysApiKeyIdRevokePost

> CompaniesCompanyIdApiKeysPost200Response CompaniesCompanyIdApiKeysApiKeyIdRevokePost(ctx, companyId, apiKeyId).ApiKeysInput(apiKeysInput).Execute()

Explicitly revoke an API key with reason tracking.     This is different from delete as it includes revocation metadata.

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
	apiKeysInput := *openapiclient.NewApiKeysInput("CompanyId_example", int64(123)) // ApiKeysInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdRevokePost(context.Background(), companyId, apiKeyId).ApiKeysInput(apiKeysInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdRevokePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdApiKeysApiKeyIdRevokePost`: CompaniesCompanyIdApiKeysPost200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdRevokePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdApiKeysApiKeyIdRevokePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiKeysInput** | [**ApiKeysInput**](ApiKeysInput.md) |  | 

### Return type

[**CompaniesCompanyIdApiKeysPost200Response**](CompaniesCompanyIdApiKeysPost200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdApiKeysApiKeyIdStatsGet

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdApiKeysApiKeyIdStatsGet(ctx, companyId, apiKeyId).Execute()

Get usage statistics for a specific API key.     Returns usage count, last used date, failed attempts, etc.

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
	resp, r, err := apiClient.APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdStatsGet(context.Background(), companyId, apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdStatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdApiKeysApiKeyIdStatsGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.CompaniesCompanyIdApiKeysApiKeyIdStatsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**apiKeyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdApiKeysApiKeyIdStatsGetRequest struct via the builder pattern


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

Get all API keys for a company.     Returns masked API keys for security (sk_****1234).

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

> CompaniesCompanyIdApiKeysPost200Response CompaniesCompanyIdApiKeysPost(ctx, companyId).ApiKeysInput(apiKeysInput).Execute()

Create single API key.     Returns the actual sk_ key (only time it's exposed) and api_key_id for future operations.

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
	apiKeysInput := *openapiclient.NewApiKeysInput("CompanyId_example", int64(123)) // ApiKeysInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.CompaniesCompanyIdApiKeysPost(context.Background(), companyId).ApiKeysInput(apiKeysInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CompaniesCompanyIdApiKeysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdApiKeysPost`: CompaniesCompanyIdApiKeysPost200Response
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

 **apiKeysInput** | [**ApiKeysInput**](ApiKeysInput.md) |  | 

### Return type

[**CompaniesCompanyIdApiKeysPost200Response**](CompaniesCompanyIdApiKeysPost200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


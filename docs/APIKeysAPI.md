# \APIKeysAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiKeys**](APIKeysAPI.md#CreateApiKeys) | **Post** /companies/{company_id}/api-keys | Create single API key.     Returns the actual sk_ key (only time it&#39;s exposed) and api_key_id for future operations.
[**CreateApiKeysApiKeysRevoke**](APIKeysAPI.md#CreateApiKeysApiKeysRevoke) | **Post** /companies/{company_id}/api-keys/{api_key_id}/revoke | Explicitly revoke an API key with reason tracking.     This is different from delete as it includes revocation metadata.
[**DeleteApiKeys**](APIKeysAPI.md#DeleteApiKeys) | **Delete** /companies/{company_id}/api-keys/{api_key_id} | Delete (deactivate) single API key by api_key_id.     Uses the api_key_id returned during creation for clean identification.     Fixed to use correct primary key lookup.
[**GetApiKeysById**](APIKeysAPI.md#GetApiKeysById) | **Get** /companies/{company_id}/api-keys/{api_key_id} | Get single API key by ID.     Returns masked API key for security (sk_****1234).
[**GetApiKeysByIdApiKeysStats**](APIKeysAPI.md#GetApiKeysByIdApiKeysStats) | **Get** /companies/{company_id}/api-keys/{api_key_id}/stats | Get usage statistics for a specific API key.     Returns usage count, last used date, failed attempts, etc.
[**ListApiKeys**](APIKeysAPI.md#ListApiKeys) | **Get** /companies/{company_id}/api-keys | Get all API keys for a company.     Returns masked API keys for security (sk_****1234).
[**UpdateApiKeys**](APIKeysAPI.md#UpdateApiKeys) | **Patch** /companies/{company_id}/api-keys/{api_key_id} | Update an existing API key by ID.     Can update metadata like name, expiration, rate limits, etc.     Cannot update the actual key value (for security).



## CreateApiKeys

> CreateApiKeys200Response CreateApiKeys(ctx, companyId).ApiKeysInput(apiKeysInput).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Create single API key.     Returns the actual sk_ key (only time it's exposed) and api_key_id for future operations.

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
	companyId := "companyId_example" // string | Unique identifier for the Company
	apiKeysInput := *openapiclient.NewApiKeysInput("company_id_abc123", int64(1)) // ApiKeysInput | 
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	limit := int32(56) // int32 | Number of items per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	sortOrder := "sortOrder_example" // string | Sort order (ascending or descending) (optional) (default to "desc")
	search := "search_example" // string | Search term to filter results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.CreateApiKeys(context.Background(), companyId).ApiKeysInput(apiKeysInput).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CreateApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiKeys`: CreateApiKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.CreateApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKeysInput** | [**ApiKeysInput**](ApiKeysInput.md) |  | 
 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**CreateApiKeys200Response**](CreateApiKeys200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiKeysApiKeysRevoke

> CreateApiKeys200Response CreateApiKeysApiKeysRevoke(ctx, companyId, apiKeyId).ApiKeysInput(apiKeysInput).Execute()

Explicitly revoke an API key with reason tracking.     This is different from delete as it includes revocation metadata.

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
	companyId := "companyId_example" // string | Unique identifier for the Company
	apiKeyId := "apiKeyId_example" // string | Unique identifier for the Api Key
	apiKeysInput := *openapiclient.NewApiKeysInput("company_id_abc123", int64(1)) // ApiKeysInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.CreateApiKeysApiKeysRevoke(context.Background(), companyId, apiKeyId).ApiKeysInput(apiKeysInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.CreateApiKeysApiKeysRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiKeysApiKeysRevoke`: CreateApiKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.CreateApiKeysApiKeysRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**apiKeyId** | **string** | Unique identifier for the Api Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKeysApiKeysRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiKeysInput** | [**ApiKeysInput**](ApiKeysInput.md) |  | 

### Return type

[**CreateApiKeys200Response**](CreateApiKeys200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiKeys

> DeleteApiKeys200Response DeleteApiKeys(ctx, companyId, apiKeyId).Execute()

Delete (deactivate) single API key by api_key_id.     Uses the api_key_id returned during creation for clean identification.     Fixed to use correct primary key lookup.

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
	companyId := "companyId_example" // string | Unique identifier for the Company
	apiKeyId := "apiKeyId_example" // string | Unique identifier for the Api Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.DeleteApiKeys(context.Background(), companyId, apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.DeleteApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiKeys`: DeleteApiKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.DeleteApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**apiKeyId** | **string** | Unique identifier for the Api Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteApiKeys200Response**](DeleteApiKeys200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeysById

> GetApiKeysById200Response GetApiKeysById(ctx, companyId, apiKeyId).Execute()

Get single API key by ID.     Returns masked API key for security (sk_****1234).

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
	companyId := "companyId_example" // string | Unique identifier for the Company
	apiKeyId := "apiKeyId_example" // string | Unique identifier for the Api Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.GetApiKeysById(context.Background(), companyId, apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.GetApiKeysById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKeysById`: GetApiKeysById200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.GetApiKeysById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**apiKeyId** | **string** | Unique identifier for the Api Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeysByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetApiKeysById200Response**](GetApiKeysById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeysByIdApiKeysStats

> GetApiKeysById200Response GetApiKeysByIdApiKeysStats(ctx, companyId, apiKeyId).Execute()

Get usage statistics for a specific API key.     Returns usage count, last used date, failed attempts, etc.

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
	companyId := "companyId_example" // string | Unique identifier for the Company
	apiKeyId := "apiKeyId_example" // string | Unique identifier for the Api Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.GetApiKeysByIdApiKeysStats(context.Background(), companyId, apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.GetApiKeysByIdApiKeysStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKeysByIdApiKeysStats`: GetApiKeysById200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.GetApiKeysByIdApiKeysStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**apiKeyId** | **string** | Unique identifier for the Api Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeysByIdApiKeysStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetApiKeysById200Response**](GetApiKeysById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiKeys

> ListApiKeys200Response ListApiKeys(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get all API keys for a company.     Returns masked API keys for security (sk_****1234).

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
	companyId := "companyId_example" // string | Unique identifier for the Company
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	limit := int32(56) // int32 | Number of items per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	sortOrder := "sortOrder_example" // string | Sort order (ascending or descending) (optional) (default to "desc")
	search := "search_example" // string | Search term to filter results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.ListApiKeys(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.ListApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApiKeys`: ListApiKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.ListApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListApiKeys200Response**](ListApiKeys200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiKeys

> UpdateApiKeys200Response UpdateApiKeys(ctx, companyId, apiKeyId).ApiKeysUpdate(apiKeysUpdate).Execute()

Update an existing API key by ID.     Can update metadata like name, expiration, rate limits, etc.     Cannot update the actual key value (for security).

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
	companyId := "companyId_example" // string | Unique identifier for the Company
	apiKeyId := "apiKeyId_example" // string | Unique identifier for the Api Key
	apiKeysUpdate := *openapiclient.NewApiKeysUpdate() // ApiKeysUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIKeysAPI.UpdateApiKeys(context.Background(), companyId, apiKeyId).ApiKeysUpdate(apiKeysUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIKeysAPI.UpdateApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiKeys`: UpdateApiKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `APIKeysAPI.UpdateApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**apiKeyId** | **string** | Unique identifier for the Api Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiKeysUpdate** | [**ApiKeysUpdate**](ApiKeysUpdate.md) |  | 

### Return type

[**UpdateApiKeys200Response**](UpdateApiKeys200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


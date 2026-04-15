# \PostgenIntegrationsAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePostgenIntegrations**](PostgenIntegrationsAPI.md#CreatePostgenIntegrations) | **Post** /companies/{company_id}/postgen_integrations | POST /companies/{company_id}/postgen_integrations
[**CreatePostgenIntegrationsTest**](PostgenIntegrationsAPI.md#CreatePostgenIntegrationsTest) | **Post** /companies/{company_id}/postgen_integrations/test | POST /companies/{company_id}/postgen_integrations/test
[**DeletePostgenIntegrations**](PostgenIntegrationsAPI.md#DeletePostgenIntegrations) | **Delete** /companies/{company_id}/postgen_integrations/{integration_id} | Delete single integration by ID.     Also deletes credentials from GCP Secret Manager.
[**GetPostgenIntegrationsById**](PostgenIntegrationsAPI.md#GetPostgenIntegrationsById) | **Get** /companies/{company_id}/postgen_integrations/{integration_id} | Get single integration by ID.
[**ListPostgenIntegrations**](PostgenIntegrationsAPI.md#ListPostgenIntegrations) | **Get** /companies/{company_id}/postgen_integrations | Get a list of all postgen integrations for the company.     All company users can view integrations.
[**UpdatePostgenIntegrations**](PostgenIntegrationsAPI.md#UpdatePostgenIntegrations) | **Patch** /companies/{company_id}/postgen_integrations/{integration_id} | Update an existing integration by ID.      Can update credentials, is_active status, etc.



## CreatePostgenIntegrations

> CreatePostgenIntegrations200Response CreatePostgenIntegrations(ctx, companyId).PostgenIntegrationsInput(postgenIntegrationsInput).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

POST /companies/{company_id}/postgen_integrations

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
	postgenIntegrationsInput := *openapiclient.NewPostgenIntegrationsInput("company_id_abc123", "user_id_abc123", "example_value", "Example Name", "Example Name") // PostgenIntegrationsInput | 
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	limit := int32(56) // int32 | Number of items per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	sortOrder := "sortOrder_example" // string | Sort order (ascending or descending) (optional) (default to "desc")
	search := "search_example" // string | Search term to filter results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostgenIntegrationsAPI.CreatePostgenIntegrations(context.Background(), companyId).PostgenIntegrationsInput(postgenIntegrationsInput).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgenIntegrationsAPI.CreatePostgenIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePostgenIntegrations`: CreatePostgenIntegrations200Response
	fmt.Fprintf(os.Stdout, "Response from `PostgenIntegrationsAPI.CreatePostgenIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostgenIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postgenIntegrationsInput** | [**PostgenIntegrationsInput**](PostgenIntegrationsInput.md) |  | 
 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**CreatePostgenIntegrations200Response**](CreatePostgenIntegrations200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePostgenIntegrationsTest

> CreatePostgenIntegrations200Response CreatePostgenIntegrationsTest(ctx, companyId).PostgenIntegrationsInput(postgenIntegrationsInput).Execute()

POST /companies/{company_id}/postgen_integrations/test

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
	postgenIntegrationsInput := *openapiclient.NewPostgenIntegrationsInput("company_id_abc123", "user_id_abc123", "example_value", "Example Name", "Example Name") // PostgenIntegrationsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostgenIntegrationsAPI.CreatePostgenIntegrationsTest(context.Background(), companyId).PostgenIntegrationsInput(postgenIntegrationsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgenIntegrationsAPI.CreatePostgenIntegrationsTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePostgenIntegrationsTest`: CreatePostgenIntegrations200Response
	fmt.Fprintf(os.Stdout, "Response from `PostgenIntegrationsAPI.CreatePostgenIntegrationsTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostgenIntegrationsTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postgenIntegrationsInput** | [**PostgenIntegrationsInput**](PostgenIntegrationsInput.md) |  | 

### Return type

[**CreatePostgenIntegrations200Response**](CreatePostgenIntegrations200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePostgenIntegrations

> DeletePostgenIntegrations200Response DeletePostgenIntegrations(ctx, companyId, integrationId).Execute()

Delete single integration by ID.     Also deletes credentials from GCP Secret Manager.

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
	integrationId := "integrationId_example" // string | Unique identifier for the Integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostgenIntegrationsAPI.DeletePostgenIntegrations(context.Background(), companyId, integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgenIntegrationsAPI.DeletePostgenIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePostgenIntegrations`: DeletePostgenIntegrations200Response
	fmt.Fprintf(os.Stdout, "Response from `PostgenIntegrationsAPI.DeletePostgenIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**integrationId** | **string** | Unique identifier for the Integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostgenIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeletePostgenIntegrations200Response**](DeletePostgenIntegrations200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostgenIntegrationsById

> GetPostgenIntegrationsById200Response GetPostgenIntegrationsById(ctx, companyId, integrationId).Execute()

Get single integration by ID.

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
	integrationId := "integrationId_example" // string | Unique identifier for the Integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostgenIntegrationsAPI.GetPostgenIntegrationsById(context.Background(), companyId, integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgenIntegrationsAPI.GetPostgenIntegrationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostgenIntegrationsById`: GetPostgenIntegrationsById200Response
	fmt.Fprintf(os.Stdout, "Response from `PostgenIntegrationsAPI.GetPostgenIntegrationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**integrationId** | **string** | Unique identifier for the Integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostgenIntegrationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetPostgenIntegrationsById200Response**](GetPostgenIntegrationsById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPostgenIntegrations

> ListPostgenIntegrations200Response ListPostgenIntegrations(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get a list of all postgen integrations for the company.     All company users can view integrations.

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
	resp, r, err := apiClient.PostgenIntegrationsAPI.ListPostgenIntegrations(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgenIntegrationsAPI.ListPostgenIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPostgenIntegrations`: ListPostgenIntegrations200Response
	fmt.Fprintf(os.Stdout, "Response from `PostgenIntegrationsAPI.ListPostgenIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPostgenIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListPostgenIntegrations200Response**](ListPostgenIntegrations200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePostgenIntegrations

> UpdatePostgenIntegrations200Response UpdatePostgenIntegrations(ctx, companyId, integrationId).PostgenIntegrationsUpdate(postgenIntegrationsUpdate).Execute()

Update an existing integration by ID.      Can update credentials, is_active status, etc.

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
	integrationId := "integrationId_example" // string | Unique identifier for the Integration
	postgenIntegrationsUpdate := *openapiclient.NewPostgenIntegrationsUpdate() // PostgenIntegrationsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostgenIntegrationsAPI.UpdatePostgenIntegrations(context.Background(), companyId, integrationId).PostgenIntegrationsUpdate(postgenIntegrationsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgenIntegrationsAPI.UpdatePostgenIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePostgenIntegrations`: UpdatePostgenIntegrations200Response
	fmt.Fprintf(os.Stdout, "Response from `PostgenIntegrationsAPI.UpdatePostgenIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**integrationId** | **string** | Unique identifier for the Integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostgenIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postgenIntegrationsUpdate** | [**PostgenIntegrationsUpdate**](PostgenIntegrationsUpdate.md) |  | 

### Return type

[**UpdatePostgenIntegrations200Response**](UpdatePostgenIntegrations200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


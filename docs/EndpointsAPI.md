# \EndpointsAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndpoints**](EndpointsAPI.md#CreateEndpoints) | **Post** /companies/{company_id}/endpoints | Create a new endpoint
[**CreateEndpointsKeys**](EndpointsAPI.md#CreateEndpointsKeys) | **Post** /companies/{company_id}/endpoints/{endpoint_id}/keys | POST /companies/{company_id}/endpoints/{endpoint_id}/keys
[**DeleteEndpoints**](EndpointsAPI.md#DeleteEndpoints) | **Delete** /companies/{company_id}/endpoints/{endpoint_id} | Delete single endpoint by ID
[**DeleteEndpointsKeys**](EndpointsAPI.md#DeleteEndpointsKeys) | **Delete** /companies/{company_id}/endpoints/{endpoint_id}/keys/{api_key_id} | DELETE /companies/{company_id}/endpoints/{endpoint_id}/keys/{api_key_id}
[**GetEndpointsById**](EndpointsAPI.md#GetEndpointsById) | **Get** /companies/{company_id}/endpoints/{endpoint_id} | Get single endpoint by ID
[**GetEndpointsByIdAvailableEndpoints**](EndpointsAPI.md#GetEndpointsByIdAvailableEndpoints) | **Get** /companies/{company_id}/endpoints/{endpoint_id}/available-endpoints | GET /companies/{company_id}/endpoints/{endpoint_id}/available-endpoints
[**GetEndpointsByIdConnectionsDescribe**](EndpointsAPI.md#GetEndpointsByIdConnectionsDescribe) | **Get** /companies/{company_id}/endpoints/../connections/{connection_id}/describe | Get schema information for a connection      Query parameters:         include_fields: Whether to include field information (default: true)         schemas: Optional comma-separated schemas to include         tables: Optional comma-separated tables to include
[**GetEndpointsByIdExecute**](EndpointsAPI.md#GetEndpointsByIdExecute) | **Get** /companies/{company_id}/endpoints/{endpoint_id}/execute | Execute an endpoint with pagination support.      Customer-facing route that returns paginated data.     Query params: ?start&#x3D;0&amp;limit&#x3D;100
[**GetEndpointsByIdKeys**](EndpointsAPI.md#GetEndpointsByIdKeys) | **Get** /companies/{company_id}/endpoints/{endpoint_id}/keys | GET /companies/{company_id}/endpoints/{endpoint_id}/keys
[**GetEndpointsByIdStats**](EndpointsAPI.md#GetEndpointsByIdStats) | **Get** /companies/{company_id}/endpoints/{endpoint_id}/stats | Get usage statistics for an endpoint      Query parameters:         days: Number of days to analyze (default: 30)
[**GetEndpointsByIdTest**](EndpointsAPI.md#GetEndpointsByIdTest) | **Get** /companies/{company_id}/endpoints/{endpoint_id}/test | Test an endpoint with sample data      Request body (optional):         limit: Number of sample rows to return (default: 10)
[**GetEndpointsByIdUrl**](EndpointsAPI.md#GetEndpointsByIdUrl) | **Get** /companies/{company_id}/endpoints/{endpoint_id}/url | GET /companies/{company_id}/endpoints/{endpoint_id}/url
[**ListEndpoints**](EndpointsAPI.md#ListEndpoints) | **Get** /companies/{company_id}/endpoints | Get all endpoints for a specific company with pagination, filtering, and sorting
[**UpdateEndpoints**](EndpointsAPI.md#UpdateEndpoints) | **Patch** /companies/{company_id}/endpoints/{endpoint_id} | Update an existing endpoint by ID



## CreateEndpoints

> CreateEndpoints200Response CreateEndpoints(ctx, companyId).EndpointsInput(endpointsInput).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Create a new endpoint

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
	endpointsInput := *openapiclient.NewEndpointsInput("company_id_abc123", "connection_id_abc123", "Example Name") // EndpointsInput | 
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	limit := int32(56) // int32 | Number of items per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	sortOrder := "sortOrder_example" // string | Sort order (ascending or descending) (optional) (default to "desc")
	search := "search_example" // string | Search term to filter results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.CreateEndpoints(context.Background(), companyId).EndpointsInput(endpointsInput).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.CreateEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpoints`: CreateEndpoints200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.CreateEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointsInput** | [**EndpointsInput**](EndpointsInput.md) |  | 
 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**CreateEndpoints200Response**](CreateEndpoints200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEndpointsKeys

> CreateEndpoints200Response CreateEndpointsKeys(ctx, companyId, endpointId).EndpointsInput(endpointsInput).Execute()

POST /companies/{company_id}/endpoints/{endpoint_id}/keys

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
	endpointId := "endpointId_example" // string | Unique identifier for the Endpoint
	endpointsInput := *openapiclient.NewEndpointsInput("company_id_abc123", "connection_id_abc123", "Example Name") // EndpointsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.CreateEndpointsKeys(context.Background(), companyId, endpointId).EndpointsInput(endpointsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.CreateEndpointsKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpointsKeys`: CreateEndpoints200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.CreateEndpointsKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**endpointId** | **string** | Unique identifier for the Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointsKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endpointsInput** | [**EndpointsInput**](EndpointsInput.md) |  | 

### Return type

[**CreateEndpoints200Response**](CreateEndpoints200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEndpoints

> DeleteEndpoints200Response DeleteEndpoints(ctx, companyId, endpointId).Execute()

Delete single endpoint by ID

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
	endpointId := "endpointId_example" // string | Unique identifier for the Endpoint

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.DeleteEndpoints(context.Background(), companyId, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DeleteEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEndpoints`: DeleteEndpoints200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.DeleteEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**endpointId** | **string** | Unique identifier for the Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteEndpoints200Response**](DeleteEndpoints200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEndpointsKeys

> DeleteEndpoints200Response DeleteEndpointsKeys(ctx, companyId, endpointId, apiKeyId).Execute()

DELETE /companies/{company_id}/endpoints/{endpoint_id}/keys/{api_key_id}

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
	endpointId := "endpointId_example" // string | Unique identifier for the Endpoint
	apiKeyId := "apiKeyId_example" // string | Unique identifier for the Api Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.DeleteEndpointsKeys(context.Background(), companyId, endpointId, apiKeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.DeleteEndpointsKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEndpointsKeys`: DeleteEndpoints200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.DeleteEndpointsKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**endpointId** | **string** | Unique identifier for the Endpoint | 
**apiKeyId** | **string** | Unique identifier for the Api Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointsKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeleteEndpoints200Response**](DeleteEndpoints200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointsById

> GetEndpointsByIdConnectionsDescribe200Response GetEndpointsById(ctx, companyId, endpointId).Execute()

Get single endpoint by ID

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
	endpointId := "endpointId_example" // string | Unique identifier for the Endpoint

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.GetEndpointsById(context.Background(), companyId, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointsById`: GetEndpointsByIdConnectionsDescribe200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**endpointId** | **string** | Unique identifier for the Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEndpointsByIdConnectionsDescribe200Response**](GetEndpointsByIdConnectionsDescribe200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointsByIdAvailableEndpoints

> GetEndpointsByIdConnectionsDescribe200Response GetEndpointsByIdAvailableEndpoints(ctx, companyId, endpointId).Execute()

GET /companies/{company_id}/endpoints/{endpoint_id}/available-endpoints

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
	endpointId := "endpointId_example" // string | Unique identifier for the Endpoint

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.GetEndpointsByIdAvailableEndpoints(context.Background(), companyId, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointsByIdAvailableEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointsByIdAvailableEndpoints`: GetEndpointsByIdConnectionsDescribe200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointsByIdAvailableEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**endpointId** | **string** | Unique identifier for the Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointsByIdAvailableEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEndpointsByIdConnectionsDescribe200Response**](GetEndpointsByIdConnectionsDescribe200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointsByIdConnectionsDescribe

> GetEndpointsByIdConnectionsDescribe200Response GetEndpointsByIdConnectionsDescribe(ctx, companyId, connectionId).Execute()

Get schema information for a connection      Query parameters:         include_fields: Whether to include field information (default: true)         schemas: Optional comma-separated schemas to include         tables: Optional comma-separated tables to include

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
	connectionId := "connectionId_example" // string | Unique identifier for the Connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.GetEndpointsByIdConnectionsDescribe(context.Background(), companyId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointsByIdConnectionsDescribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointsByIdConnectionsDescribe`: GetEndpointsByIdConnectionsDescribe200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointsByIdConnectionsDescribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**connectionId** | **string** | Unique identifier for the Connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointsByIdConnectionsDescribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEndpointsByIdConnectionsDescribe200Response**](GetEndpointsByIdConnectionsDescribe200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointsByIdExecute

> GetEndpointsByIdConnectionsDescribe200Response GetEndpointsByIdExecute(ctx, companyId, endpointId).Execute()

Execute an endpoint with pagination support.      Customer-facing route that returns paginated data.     Query params: ?start=0&limit=100

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
	endpointId := "endpointId_example" // string | Unique identifier for the Endpoint

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.GetEndpointsByIdExecute(context.Background(), companyId, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointsByIdExecute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointsByIdExecute`: GetEndpointsByIdConnectionsDescribe200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointsByIdExecute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**endpointId** | **string** | Unique identifier for the Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointsByIdExecuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEndpointsByIdConnectionsDescribe200Response**](GetEndpointsByIdConnectionsDescribe200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointsByIdKeys

> GetEndpointsByIdConnectionsDescribe200Response GetEndpointsByIdKeys(ctx, companyId, endpointId).Execute()

GET /companies/{company_id}/endpoints/{endpoint_id}/keys

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
	endpointId := "endpointId_example" // string | Unique identifier for the Endpoint

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.GetEndpointsByIdKeys(context.Background(), companyId, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointsByIdKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointsByIdKeys`: GetEndpointsByIdConnectionsDescribe200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointsByIdKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**endpointId** | **string** | Unique identifier for the Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointsByIdKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEndpointsByIdConnectionsDescribe200Response**](GetEndpointsByIdConnectionsDescribe200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointsByIdStats

> GetEndpointsByIdConnectionsDescribe200Response GetEndpointsByIdStats(ctx, companyId, endpointId).Execute()

Get usage statistics for an endpoint      Query parameters:         days: Number of days to analyze (default: 30)

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
	endpointId := "endpointId_example" // string | Unique identifier for the Endpoint

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.GetEndpointsByIdStats(context.Background(), companyId, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointsByIdStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointsByIdStats`: GetEndpointsByIdConnectionsDescribe200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointsByIdStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**endpointId** | **string** | Unique identifier for the Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointsByIdStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEndpointsByIdConnectionsDescribe200Response**](GetEndpointsByIdConnectionsDescribe200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointsByIdTest

> GetEndpointsByIdConnectionsDescribe200Response GetEndpointsByIdTest(ctx, companyId, endpointId).Execute()

Test an endpoint with sample data      Request body (optional):         limit: Number of sample rows to return (default: 10)

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
	endpointId := "endpointId_example" // string | Unique identifier for the Endpoint

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.GetEndpointsByIdTest(context.Background(), companyId, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointsByIdTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointsByIdTest`: GetEndpointsByIdConnectionsDescribe200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointsByIdTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**endpointId** | **string** | Unique identifier for the Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointsByIdTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEndpointsByIdConnectionsDescribe200Response**](GetEndpointsByIdConnectionsDescribe200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointsByIdUrl

> GetEndpointsByIdConnectionsDescribe200Response GetEndpointsByIdUrl(ctx, companyId, endpointId).Execute()

GET /companies/{company_id}/endpoints/{endpoint_id}/url

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
	endpointId := "endpointId_example" // string | Unique identifier for the Endpoint

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.GetEndpointsByIdUrl(context.Background(), companyId, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpointsByIdUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpointsByIdUrl`: GetEndpointsByIdConnectionsDescribe200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpointsByIdUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**endpointId** | **string** | Unique identifier for the Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointsByIdUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEndpointsByIdConnectionsDescribe200Response**](GetEndpointsByIdConnectionsDescribe200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndpoints

> ListEndpoints200Response ListEndpoints(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get all endpoints for a specific company with pagination, filtering, and sorting

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
	resp, r, err := apiClient.EndpointsAPI.ListEndpoints(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.ListEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEndpoints`: ListEndpoints200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.ListEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListEndpoints200Response**](ListEndpoints200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpoints

> UpdateEndpoints200Response UpdateEndpoints(ctx, companyId, endpointId).EndpointsUpdate(endpointsUpdate).Execute()

Update an existing endpoint by ID

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
	endpointId := "endpointId_example" // string | Unique identifier for the Endpoint
	endpointsUpdate := *openapiclient.NewEndpointsUpdate() // EndpointsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.UpdateEndpoints(context.Background(), companyId, endpointId).EndpointsUpdate(endpointsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.UpdateEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpoints`: UpdateEndpoints200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.UpdateEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**endpointId** | **string** | Unique identifier for the Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endpointsUpdate** | [**EndpointsUpdate**](EndpointsUpdate.md) |  | 

### Return type

[**UpdateEndpoints200Response**](UpdateEndpoints200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


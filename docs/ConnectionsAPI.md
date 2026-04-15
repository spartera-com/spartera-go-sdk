# \ConnectionsAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnections**](ConnectionsAPI.md#CreateConnections) | **Post** /companies/{company_id}/connections | Create a new connection by ID
[**DeleteConnections**](ConnectionsAPI.md#DeleteConnections) | **Delete** /companies/{company_id}/connections/{connection_id} | Delete single connection by ID
[**GetConnectionsById**](ConnectionsAPI.md#GetConnectionsById) | **Get** /companies/{company_id}/connections/{connection_id} | Get single connection by ID
[**GetConnectionsById2**](ConnectionsAPI.md#GetConnectionsById2) | **Get** /companies/{company_id}/connections/{connection_id}/test | Test the specified connection
[**GetConnectionsByIdInfoschema**](ConnectionsAPI.md#GetConnectionsByIdInfoschema) | **Get** /companies/{company_id}/connections/{connection_id}/infoschema | Retrieve the information schema for the specified connection
[**ListConnections**](ConnectionsAPI.md#ListConnections) | **Get** /companies/{company_id}/connections | Get all connections for a specific company
[**UpdateConnections**](ConnectionsAPI.md#UpdateConnections) | **Patch** /companies/{company_id}/connections/{connection_id} | Update an existing connection by ID



## CreateConnections

> CreateConnections200Response CreateConnections(ctx, companyId).ConnectionsInput(connectionsInput).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Create a new connection by ID

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
	connectionsInput := *openapiclient.NewConnectionsInput(int64(1), "company_id_abc123") // ConnectionsInput | 
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	limit := int32(56) // int32 | Number of items per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	sortOrder := "sortOrder_example" // string | Sort order (ascending or descending) (optional) (default to "desc")
	search := "search_example" // string | Search term to filter results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.CreateConnections(context.Background(), companyId).ConnectionsInput(connectionsInput).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CreateConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnections`: CreateConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CreateConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionsInput** | [**ConnectionsInput**](ConnectionsInput.md) |  | 
 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**CreateConnections200Response**](CreateConnections200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnections

> DeleteConnections200Response DeleteConnections(ctx, companyId, connectionId).Execute()

Delete single connection by ID

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
	resp, r, err := apiClient.ConnectionsAPI.DeleteConnections(context.Background(), companyId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.DeleteConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteConnections`: DeleteConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.DeleteConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**connectionId** | **string** | Unique identifier for the Connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteConnections200Response**](DeleteConnections200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionsById

> GetConnectionsById200Response GetConnectionsById(ctx, companyId, connectionId).Execute()

Get single connection by ID

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
	resp, r, err := apiClient.ConnectionsAPI.GetConnectionsById(context.Background(), companyId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.GetConnectionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionsById`: GetConnectionsById200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.GetConnectionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**connectionId** | **string** | Unique identifier for the Connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetConnectionsById200Response**](GetConnectionsById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionsById2

> GetConnectionsById200Response GetConnectionsById2(ctx, companyId, connectionId).Execute()

Test the specified connection

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
	resp, r, err := apiClient.ConnectionsAPI.GetConnectionsById2(context.Background(), companyId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.GetConnectionsById2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionsById2`: GetConnectionsById200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.GetConnectionsById2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**connectionId** | **string** | Unique identifier for the Connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsById2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetConnectionsById200Response**](GetConnectionsById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionsByIdInfoschema

> GetConnectionsById200Response GetConnectionsByIdInfoschema(ctx, companyId, connectionId).Execute()

Retrieve the information schema for the specified connection

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
	resp, r, err := apiClient.ConnectionsAPI.GetConnectionsByIdInfoschema(context.Background(), companyId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.GetConnectionsByIdInfoschema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionsByIdInfoschema`: GetConnectionsById200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.GetConnectionsByIdInfoschema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**connectionId** | **string** | Unique identifier for the Connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsByIdInfoschemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetConnectionsById200Response**](GetConnectionsById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnections

> ListConnections200Response ListConnections(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get all connections for a specific company

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
	resp, r, err := apiClient.ConnectionsAPI.ListConnections(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ListConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnections`: ListConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ListConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListConnections200Response**](ListConnections200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnections

> UpdateConnections200Response UpdateConnections(ctx, companyId, connectionId).ConnectionsUpdate(connectionsUpdate).Execute()

Update an existing connection by ID

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
	connectionsUpdate := *openapiclient.NewConnectionsUpdate() // ConnectionsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.UpdateConnections(context.Background(), companyId, connectionId).ConnectionsUpdate(connectionsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.UpdateConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConnections`: UpdateConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.UpdateConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**connectionId** | **string** | Unique identifier for the Connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectionsUpdate** | [**ConnectionsUpdate**](ConnectionsUpdate.md) |  | 

### Return type

[**UpdateConnections200Response**](UpdateConnections200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


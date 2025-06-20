# \ConnectionsAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompaniesCompanyIdConnectionsConnectionIdDelete**](ConnectionsAPI.md#CompaniesCompanyIdConnectionsConnectionIdDelete) | **Delete** /companies/{company_id}/connections/{connection_id} | Delete single connection by ID
[**CompaniesCompanyIdConnectionsConnectionIdGet**](ConnectionsAPI.md#CompaniesCompanyIdConnectionsConnectionIdGet) | **Get** /companies/{company_id}/connections/{connection_id} | Get single connection by ID
[**CompaniesCompanyIdConnectionsConnectionIdInfoschemaGet**](ConnectionsAPI.md#CompaniesCompanyIdConnectionsConnectionIdInfoschemaGet) | **Get** /companies/{company_id}/connections/{connection_id}/infoschema | Retrieve the information schema for the specified connection
[**CompaniesCompanyIdConnectionsConnectionIdPatch**](ConnectionsAPI.md#CompaniesCompanyIdConnectionsConnectionIdPatch) | **Patch** /companies/{company_id}/connections/{connection_id} | Update an existing connection by ID
[**CompaniesCompanyIdConnectionsConnectionIdTestGet**](ConnectionsAPI.md#CompaniesCompanyIdConnectionsConnectionIdTestGet) | **Get** /companies/{company_id}/connections/{connection_id}/test | Test the specified connection
[**CompaniesCompanyIdConnectionsGet**](ConnectionsAPI.md#CompaniesCompanyIdConnectionsGet) | **Get** /companies/{company_id}/connections | Get all connections for a specific company
[**CompaniesCompanyIdConnectionsPost**](ConnectionsAPI.md#CompaniesCompanyIdConnectionsPost) | **Post** /companies/{company_id}/connections | Create a new connection by ID



## CompaniesCompanyIdConnectionsConnectionIdDelete

> CompaniesCompanyIdConnectionsConnectionIdDelete200Response CompaniesCompanyIdConnectionsConnectionIdDelete(ctx, companyId, connectionId).Execute()

Delete single connection by ID

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
	connectionId := "connectionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdDelete(context.Background(), companyId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdConnectionsConnectionIdDelete`: CompaniesCompanyIdConnectionsConnectionIdDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdConnectionsConnectionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdConnectionsConnectionIdDelete200Response**](CompaniesCompanyIdConnectionsConnectionIdDelete200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdConnectionsConnectionIdGet

> CompaniesCompanyIdConnectionsConnectionIdGet200Response CompaniesCompanyIdConnectionsConnectionIdGet(ctx, companyId, connectionId).Execute()

Get single connection by ID

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
	connectionId := "connectionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdGet(context.Background(), companyId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdConnectionsConnectionIdGet`: CompaniesCompanyIdConnectionsConnectionIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdConnectionsConnectionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdConnectionsConnectionIdGet200Response**](CompaniesCompanyIdConnectionsConnectionIdGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdConnectionsConnectionIdInfoschemaGet

> CompaniesCompanyIdConnectionsGet200Response CompaniesCompanyIdConnectionsConnectionIdInfoschemaGet(ctx, companyId, connectionId).Execute()

Retrieve the information schema for the specified connection

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
	connectionId := "connectionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdInfoschemaGet(context.Background(), companyId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdInfoschemaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdConnectionsConnectionIdInfoschemaGet`: CompaniesCompanyIdConnectionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdInfoschemaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdConnectionsConnectionIdInfoschemaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdConnectionsGet200Response**](CompaniesCompanyIdConnectionsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdConnectionsConnectionIdPatch

> CompaniesCompanyIdConnectionsConnectionIdPatch200Response CompaniesCompanyIdConnectionsConnectionIdPatch(ctx, companyId, connectionId).Connection(connection).Execute()

Update an existing connection by ID

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
	connectionId := "connectionId_example" // string | 
	connection := *openapiclient.NewConnection("EngineId_example", "CompanyId_example", "Visibility_example") // Connection | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdPatch(context.Background(), companyId, connectionId).Connection(connection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdConnectionsConnectionIdPatch`: CompaniesCompanyIdConnectionsConnectionIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdConnectionsConnectionIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connection** | [**Connection**](Connection.md) |  | 

### Return type

[**CompaniesCompanyIdConnectionsConnectionIdPatch200Response**](CompaniesCompanyIdConnectionsConnectionIdPatch200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdConnectionsConnectionIdTestGet

> CompaniesCompanyIdConnectionsGet200Response CompaniesCompanyIdConnectionsConnectionIdTestGet(ctx, companyId, connectionId).Execute()

Test the specified connection

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
	connectionId := "connectionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdTestGet(context.Background(), companyId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdTestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdConnectionsConnectionIdTestGet`: CompaniesCompanyIdConnectionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CompaniesCompanyIdConnectionsConnectionIdTestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdConnectionsConnectionIdTestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdConnectionsGet200Response**](CompaniesCompanyIdConnectionsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdConnectionsGet

> CompaniesCompanyIdConnectionsGet200Response CompaniesCompanyIdConnectionsGet(ctx, companyId).Execute()

Get all connections for a specific company

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
	resp, r, err := apiClient.ConnectionsAPI.CompaniesCompanyIdConnectionsGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CompaniesCompanyIdConnectionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdConnectionsGet`: CompaniesCompanyIdConnectionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CompaniesCompanyIdConnectionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdConnectionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompaniesCompanyIdConnectionsGet200Response**](CompaniesCompanyIdConnectionsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdConnectionsPost

> CompaniesCompanyIdConnectionsPost200Response CompaniesCompanyIdConnectionsPost(ctx, companyId).Connection(connection).Execute()

Create a new connection by ID

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
	connection := *openapiclient.NewConnection("EngineId_example", "CompanyId_example", "Visibility_example") // Connection | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.CompaniesCompanyIdConnectionsPost(context.Background(), companyId).Connection(connection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CompaniesCompanyIdConnectionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdConnectionsPost`: CompaniesCompanyIdConnectionsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CompaniesCompanyIdConnectionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdConnectionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connection** | [**Connection**](Connection.md) |  | 

### Return type

[**CompaniesCompanyIdConnectionsPost200Response**](CompaniesCompanyIdConnectionsPost200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


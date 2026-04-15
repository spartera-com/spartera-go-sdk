# \AlertsAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlerts**](AlertsAPI.md#CreateAlerts) | **Post** /companies/{company_id}/users/{user_id}/alerts | POST /companies/{company_id}/users/{user_id}/alerts
[**DeleteAlerts**](AlertsAPI.md#DeleteAlerts) | **Delete** /companies/{company_id}/users/{user_id}/alerts/{alert_id} | Delete single alert by ID
[**GetAlertsById**](AlertsAPI.md#GetAlertsById) | **Get** /companies/{company_id}/users/{user_id}/alerts | Get a list of all alerts for a specific user
[**GetAlertsByIdAssetAll**](AlertsAPI.md#GetAlertsByIdAssetAll) | **Get** /companies/{company_id}/users/{user_id}/alerts/asset/{asset_id}/all | Get all alerts for a specific asset
[**GetAlertsByIdUsers**](AlertsAPI.md#GetAlertsByIdUsers) | **Get** /companies/{company_id}/users/{user_id}/alerts/{alert_id} | Get single alert by ID
[**GetAlertsByIdUsersAsset**](AlertsAPI.md#GetAlertsByIdUsersAsset) | **Get** /companies/{company_id}/users/{user_id}/alerts/asset/{asset_id} | Get all alerts for a specific asset (by user)
[**UpdateAlerts**](AlertsAPI.md#UpdateAlerts) | **Patch** /companies/{company_id}/users/{user_id}/alerts/{alert_id} | Update an existing alert by ID



## CreateAlerts

> CreateAlerts200Response CreateAlerts(ctx, companyId, userId).AlertsInput(alertsInput).Execute()

POST /companies/{company_id}/users/{user_id}/alerts

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
	userId := "userId_example" // string | Unique identifier for the User
	alertsInput := *openapiclient.NewAlertsInput("asset_id_abc123", "company_id_abc123") // AlertsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CreateAlerts(context.Background(), companyId, userId).AlertsInput(alertsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CreateAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlerts`: CreateAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CreateAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alertsInput** | [**AlertsInput**](AlertsInput.md) |  | 

### Return type

[**CreateAlerts200Response**](CreateAlerts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlerts

> DeleteAlerts200Response DeleteAlerts(ctx, companyId, userId, alertId).Execute()

Delete single alert by ID

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
	userId := "userId_example" // string | Unique identifier for the User
	alertId := "alertId_example" // string | Unique identifier for the Alert

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DeleteAlerts(context.Background(), companyId, userId, alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlerts`: DeleteAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 
**alertId** | **string** | Unique identifier for the Alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeleteAlerts200Response**](DeleteAlerts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsById

> GetAlertsById200Response GetAlertsById(ctx, companyId, userId).Execute()

Get a list of all alerts for a specific user

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
	userId := "userId_example" // string | Unique identifier for the User

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertsById(context.Background(), companyId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertsById`: GetAlertsById200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAlertsById200Response**](GetAlertsById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsByIdAssetAll

> GetAlertsById200Response GetAlertsByIdAssetAll(ctx, companyId, userId, assetId).Execute()

Get all alerts for a specific asset

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
	userId := "userId_example" // string | Unique identifier for the User
	assetId := "assetId_example" // string | Unique identifier for the Asset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertsByIdAssetAll(context.Background(), companyId, userId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertsByIdAssetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertsByIdAssetAll`: GetAlertsById200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertsByIdAssetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsByIdAssetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetAlertsById200Response**](GetAlertsById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsByIdUsers

> GetAlertsById200Response GetAlertsByIdUsers(ctx, companyId, userId, alertId).Execute()

Get single alert by ID

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
	userId := "userId_example" // string | Unique identifier for the User
	alertId := "alertId_example" // string | Unique identifier for the Alert

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertsByIdUsers(context.Background(), companyId, userId, alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertsByIdUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertsByIdUsers`: GetAlertsById200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertsByIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 
**alertId** | **string** | Unique identifier for the Alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsByIdUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetAlertsById200Response**](GetAlertsById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsByIdUsersAsset

> GetAlertsById200Response GetAlertsByIdUsersAsset(ctx, companyId, userId, assetId).Execute()

Get all alerts for a specific asset (by user)

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
	userId := "userId_example" // string | Unique identifier for the User
	assetId := "assetId_example" // string | Unique identifier for the Asset

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertsByIdUsersAsset(context.Background(), companyId, userId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertsByIdUsersAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertsByIdUsersAsset`: GetAlertsById200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertsByIdUsersAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 
**assetId** | **string** | Unique identifier for the Asset | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsByIdUsersAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetAlertsById200Response**](GetAlertsById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlerts

> UpdateAlerts200Response UpdateAlerts(ctx, companyId, userId, alertId).AlertsUpdate(alertsUpdate).Execute()

Update an existing alert by ID

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
	userId := "userId_example" // string | Unique identifier for the User
	alertId := "alertId_example" // string | Unique identifier for the Alert
	alertsUpdate := *openapiclient.NewAlertsUpdate() // AlertsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateAlerts(context.Background(), companyId, userId, alertId).AlertsUpdate(alertsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlerts`: UpdateAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 
**alertId** | **string** | Unique identifier for the Alert | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **alertsUpdate** | [**AlertsUpdate**](AlertsUpdate.md) |  | 

### Return type

[**UpdateAlerts200Response**](UpdateAlerts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


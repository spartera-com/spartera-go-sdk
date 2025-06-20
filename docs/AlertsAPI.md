# \AlertsAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompaniesCompanyIdUsersUserIdAlertsAlertIdDelete**](AlertsAPI.md#CompaniesCompanyIdUsersUserIdAlertsAlertIdDelete) | **Delete** /companies/{company_id}/users/{user_id}/alerts/{alert_id} | Delete single alert by ID
[**CompaniesCompanyIdUsersUserIdAlertsAlertIdGet**](AlertsAPI.md#CompaniesCompanyIdUsersUserIdAlertsAlertIdGet) | **Get** /companies/{company_id}/users/{user_id}/alerts/{alert_id} | Get single alert by ID
[**CompaniesCompanyIdUsersUserIdAlertsAlertIdPatch**](AlertsAPI.md#CompaniesCompanyIdUsersUserIdAlertsAlertIdPatch) | **Patch** /companies/{company_id}/users/{user_id}/alerts/{alert_id} | Update an existing alert by ID
[**CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdAllGet**](AlertsAPI.md#CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdAllGet) | **Get** /companies/{company_id}/users/{user_id}/alerts/asset/{asset_id}/all | Get all alerts for a specific asset
[**CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdGet**](AlertsAPI.md#CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdGet) | **Get** /companies/{company_id}/users/{user_id}/alerts/asset/{asset_id} | Get all alerts for a specific asset (by user)
[**CompaniesCompanyIdUsersUserIdAlertsGet**](AlertsAPI.md#CompaniesCompanyIdUsersUserIdAlertsGet) | **Get** /companies/{company_id}/users/{user_id}/alerts | Get a list of all alerts for a specific user
[**CompaniesCompanyIdUsersUserIdAlertsPost**](AlertsAPI.md#CompaniesCompanyIdUsersUserIdAlertsPost) | **Post** /companies/{company_id}/users/{user_id}/alerts | POST /companies/{company_id}/users/{user_id}/alerts



## CompaniesCompanyIdUsersUserIdAlertsAlertIdDelete

> CompaniesCompanyIdUsersUserIdAlertsAlertIdDelete200Response CompaniesCompanyIdUsersUserIdAlertsAlertIdDelete(ctx, companyId, userId, alertId).Execute()

Delete single alert by ID

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
	userId := "userId_example" // string | 
	alertId := "alertId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAlertIdDelete(context.Background(), companyId, userId, alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAlertIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdAlertsAlertIdDelete`: CompaniesCompanyIdUsersUserIdAlertsAlertIdDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAlertIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 
**alertId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdAlertsAlertIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CompaniesCompanyIdUsersUserIdAlertsAlertIdDelete200Response**](CompaniesCompanyIdUsersUserIdAlertsAlertIdDelete200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdAlertsAlertIdGet

> CompaniesCompanyIdUsersUserIdAlertsAlertIdGet200Response CompaniesCompanyIdUsersUserIdAlertsAlertIdGet(ctx, companyId, userId, alertId).Execute()

Get single alert by ID

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
	userId := "userId_example" // string | 
	alertId := "alertId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAlertIdGet(context.Background(), companyId, userId, alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAlertIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdAlertsAlertIdGet`: CompaniesCompanyIdUsersUserIdAlertsAlertIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAlertIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 
**alertId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdAlertsAlertIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CompaniesCompanyIdUsersUserIdAlertsAlertIdGet200Response**](CompaniesCompanyIdUsersUserIdAlertsAlertIdGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdAlertsAlertIdPatch

> CompaniesCompanyIdUsersUserIdAlertsAlertIdPatch200Response CompaniesCompanyIdUsersUserIdAlertsAlertIdPatch(ctx, companyId, userId, alertId).Alert(alert).Execute()

Update an existing alert by ID

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
	userId := "userId_example" // string | 
	alertId := "alertId_example" // string | 
	alert := *openapiclient.NewAlert("AssetId_example", "CompanyId_example", "AlertType_example") // Alert | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAlertIdPatch(context.Background(), companyId, userId, alertId).Alert(alert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAlertIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdAlertsAlertIdPatch`: CompaniesCompanyIdUsersUserIdAlertsAlertIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAlertIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 
**alertId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdAlertsAlertIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **alert** | [**Alert**](Alert.md) |  | 

### Return type

[**CompaniesCompanyIdUsersUserIdAlertsAlertIdPatch200Response**](CompaniesCompanyIdUsersUserIdAlertsAlertIdPatch200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdAllGet

> CompaniesCompanyIdUsersUserIdAlertsGet200Response CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdAllGet(ctx, companyId, userId, assetId).Execute()

Get all alerts for a specific asset

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
	userId := "userId_example" // string | 
	assetId := "assetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdAllGet(context.Background(), companyId, userId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdAllGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdAllGet`: CompaniesCompanyIdUsersUserIdAlertsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdAllGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdAlertsAssetAssetIdAllGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CompaniesCompanyIdUsersUserIdAlertsGet200Response**](CompaniesCompanyIdUsersUserIdAlertsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdGet

> CompaniesCompanyIdUsersUserIdAlertsAlertIdGet200Response CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdGet(ctx, companyId, userId, assetId).Execute()

Get all alerts for a specific asset (by user)

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
	userId := "userId_example" // string | 
	assetId := "assetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdGet(context.Background(), companyId, userId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdGet`: CompaniesCompanyIdUsersUserIdAlertsAlertIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsAssetAssetIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdAlertsAssetAssetIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CompaniesCompanyIdUsersUserIdAlertsAlertIdGet200Response**](CompaniesCompanyIdUsersUserIdAlertsAlertIdGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdAlertsGet

> CompaniesCompanyIdUsersUserIdAlertsGet200Response CompaniesCompanyIdUsersUserIdAlertsGet(ctx, companyId, userId).Execute()

Get a list of all alerts for a specific user

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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsGet(context.Background(), companyId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdAlertsGet`: CompaniesCompanyIdUsersUserIdAlertsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdAlertsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdUsersUserIdAlertsGet200Response**](CompaniesCompanyIdUsersUserIdAlertsGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdAlertsPost

> CompaniesCompanyIdUsersUserIdAlertsPost200Response CompaniesCompanyIdUsersUserIdAlertsPost(ctx, companyId, userId).Alert(alert).Execute()

POST /companies/{company_id}/users/{user_id}/alerts

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
	userId := "userId_example" // string | 
	alert := *openapiclient.NewAlert("AssetId_example", "CompanyId_example", "AlertType_example") // Alert | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsPost(context.Background(), companyId, userId).Alert(alert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdAlertsPost`: CompaniesCompanyIdUsersUserIdAlertsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CompaniesCompanyIdUsersUserIdAlertsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdAlertsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alert** | [**Alert**](Alert.md) |  | 

### Return type

[**CompaniesCompanyIdUsersUserIdAlertsPost200Response**](CompaniesCompanyIdUsersUserIdAlertsPost200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


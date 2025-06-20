# \CompaniesAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompaniesCompanyIdAnalyticsAssetsGet**](CompaniesAPI.md#CompaniesCompanyIdAnalyticsAssetsGet) | **Get** /companies/{company_id}/analytics/assets | Get asset performance analytics     Query params: start_date, end_date, limit, sort_by, include
[**CompaniesCompanyIdAnalyticsCustomersGet**](CompaniesAPI.md#CompaniesCompanyIdAnalyticsCustomersGet) | **Get** /companies/{company_id}/analytics/customers | Get customer analytics including growth and segmentation     Query params: start_date, end_date, group_by, segment_by
[**CompaniesCompanyIdAnalyticsDashboardGet**](CompaniesAPI.md#CompaniesCompanyIdAnalyticsDashboardGet) | **Get** /companies/{company_id}/analytics/dashboard | Get comprehensive dashboard analytics for seller dashboard     Includes all metrics needed for dashboard charts in one call     Query params: start_date, end_date, period (day/week/month/quarter)
[**CompaniesCompanyIdAnalyticsSalesGet**](CompaniesAPI.md#CompaniesCompanyIdAnalyticsSalesGet) | **Get** /companies/{company_id}/analytics/sales | Get sales over time analytics     Query params: start_date, end_date, group_by (day/week/month/quarter), metrics
[**CompaniesCompanyIdGet**](CompaniesAPI.md#CompaniesCompanyIdGet) | **Get** /companies/{company_id} | Get details of the requestor&#39;s own company
[**CompaniesCompanyIdObjectsGet**](CompaniesAPI.md#CompaniesCompanyIdObjectsGet) | **Get** /companies/{company_id}/objects | Get all objects (connections, assets) of a single company
[**CompaniesCompanyIdPatch**](CompaniesAPI.md#CompaniesCompanyIdPatch) | **Patch** /companies/{company_id} | Update an existing company by ID
[**CompaniesCompanyIdRequestsPlanGet**](CompaniesAPI.md#CompaniesCompanyIdRequestsPlanGet) | **Get** /companies/{company_id}/requests/plan | Get the total number of requests allocated in the company&#39;s current subscription plan.
[**CompaniesCompanyIdRequestsUsageGet**](CompaniesAPI.md#CompaniesCompanyIdRequestsUsageGet) | **Get** /companies/{company_id}/requests/usage | Get company request usage data for a specific month. Returns JSON metrics by default or CSV logs when download parameter is included.



## CompaniesCompanyIdAnalyticsAssetsGet

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdAnalyticsAssetsGet(ctx, companyId).Execute()

Get asset performance analytics     Query params: start_date, end_date, limit, sort_by, include

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
	resp, r, err := apiClient.CompaniesAPI.CompaniesCompanyIdAnalyticsAssetsGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesCompanyIdAnalyticsAssetsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAnalyticsAssetsGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesCompanyIdAnalyticsAssetsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAnalyticsAssetsGetRequest struct via the builder pattern


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


## CompaniesCompanyIdAnalyticsCustomersGet

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdAnalyticsCustomersGet(ctx, companyId).Execute()

Get customer analytics including growth and segmentation     Query params: start_date, end_date, group_by, segment_by

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
	resp, r, err := apiClient.CompaniesAPI.CompaniesCompanyIdAnalyticsCustomersGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesCompanyIdAnalyticsCustomersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAnalyticsCustomersGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesCompanyIdAnalyticsCustomersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAnalyticsCustomersGetRequest struct via the builder pattern


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


## CompaniesCompanyIdAnalyticsDashboardGet

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdAnalyticsDashboardGet(ctx, companyId).Execute()

Get comprehensive dashboard analytics for seller dashboard     Includes all metrics needed for dashboard charts in one call     Query params: start_date, end_date, period (day/week/month/quarter)

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
	resp, r, err := apiClient.CompaniesAPI.CompaniesCompanyIdAnalyticsDashboardGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesCompanyIdAnalyticsDashboardGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAnalyticsDashboardGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesCompanyIdAnalyticsDashboardGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAnalyticsDashboardGetRequest struct via the builder pattern


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


## CompaniesCompanyIdAnalyticsSalesGet

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdAnalyticsSalesGet(ctx, companyId).Execute()

Get sales over time analytics     Query params: start_date, end_date, group_by (day/week/month/quarter), metrics

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
	resp, r, err := apiClient.CompaniesAPI.CompaniesCompanyIdAnalyticsSalesGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesCompanyIdAnalyticsSalesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdAnalyticsSalesGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesCompanyIdAnalyticsSalesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdAnalyticsSalesGetRequest struct via the builder pattern


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


## CompaniesCompanyIdGet

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdGet(ctx, companyId).Execute()

Get details of the requestor's own company

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
	resp, r, err := apiClient.CompaniesAPI.CompaniesCompanyIdGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesCompanyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesCompanyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdGetRequest struct via the builder pattern


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


## CompaniesCompanyIdObjectsGet

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdObjectsGet(ctx, companyId).Execute()

Get all objects (connections, assets) of a single company

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
	resp, r, err := apiClient.CompaniesAPI.CompaniesCompanyIdObjectsGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesCompanyIdObjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdObjectsGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesCompanyIdObjectsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdObjectsGetRequest struct via the builder pattern


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


## CompaniesCompanyIdPatch

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdPatch(ctx, companyId).Execute()

Update an existing company by ID

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
	resp, r, err := apiClient.CompaniesAPI.CompaniesCompanyIdPatch(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesCompanyIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdPatch`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesCompanyIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdPatchRequest struct via the builder pattern


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


## CompaniesCompanyIdRequestsPlanGet

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdRequestsPlanGet(ctx, companyId).Execute()

Get the total number of requests allocated in the company's current subscription plan.

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
	resp, r, err := apiClient.CompaniesAPI.CompaniesCompanyIdRequestsPlanGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesCompanyIdRequestsPlanGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdRequestsPlanGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesCompanyIdRequestsPlanGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdRequestsPlanGetRequest struct via the builder pattern


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


## CompaniesCompanyIdRequestsUsageGet

> CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdRequestsUsageGet(ctx, companyId).Execute()

Get company request usage data for a specific month. Returns JSON metrics by default or CSV logs when download parameter is included.

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
	resp, r, err := apiClient.CompaniesAPI.CompaniesCompanyIdRequestsUsageGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CompaniesCompanyIdRequestsUsageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdRequestsUsageGet`: CompaniesCompanyIdApiKeysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CompaniesCompanyIdRequestsUsageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdRequestsUsageGetRequest struct via the builder pattern


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


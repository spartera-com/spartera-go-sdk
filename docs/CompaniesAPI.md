# \CompaniesAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompaniesById**](CompaniesAPI.md#GetCompaniesById) | **Get** /companies/{company_id} | Get details of the requestor&#39;s own company
[**ListCompanies**](CompaniesAPI.md#ListCompanies) | **Get** /companies/{company_id}/requests/plan | Get the total number of requests allocated in the company&#39;s current subscription plan.
[**ListCompaniesAnalyticsAssets**](CompaniesAPI.md#ListCompaniesAnalyticsAssets) | **Get** /companies/{company_id}/analytics/assets | Get asset performance analytics.     Query params: start_date, end_date, limit, sort_by, include
[**ListCompaniesAnalyticsCustomers**](CompaniesAPI.md#ListCompaniesAnalyticsCustomers) | **Get** /companies/{company_id}/analytics/customers | Get customer analytics including growth and segmentation.     Query params: start_date, end_date, group_by, segment_by
[**ListCompaniesAnalyticsDashboard**](CompaniesAPI.md#ListCompaniesAnalyticsDashboard) | **Get** /companies/{company_id}/analytics/dashboard | Get comprehensive dashboard analytics for seller dashboard.     Query params: start_date, end_date, period (day/week/month/quarter)
[**ListCompaniesAnalyticsMetrics**](CompaniesAPI.md#ListCompaniesAnalyticsMetrics) | **Get** /companies/{company_id}/analytics/metrics | Get summary metrics (counts only). No date params needed — returns all-time totals.     Performance: &lt; 0.1s
[**ListCompaniesAnalyticsSales**](CompaniesAPI.md#ListCompaniesAnalyticsSales) | **Get** /companies/{company_id}/analytics/sales | Get sales over time analytics.     Query params: start_date, end_date, group_by (day/week/month/quarter), metrics
[**ListCompaniesObjects**](CompaniesAPI.md#ListCompaniesObjects) | **Get** /companies/{company_id}/objects | Get all objects (connections, assets) of a single company
[**ListCompaniesRequestsUsage**](CompaniesAPI.md#ListCompaniesRequestsUsage) | **Get** /companies/{company_id}/requests/usage | Get company request usage data for a specific month.
[**UpdateCompanies**](CompaniesAPI.md#UpdateCompanies) | **Patch** /companies/{company_id} | Update an existing company by ID



## GetCompaniesById

> GetCompaniesById200Response GetCompaniesById(ctx, companyId).Execute()

Get details of the requestor's own company

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.GetCompaniesById(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetCompaniesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompaniesById`: GetCompaniesById200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.GetCompaniesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCompaniesById200Response**](GetCompaniesById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanies

> ListCompaniesAnalyticsAssets200Response ListCompanies(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get the total number of requests allocated in the company's current subscription plan.

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
	resp, r, err := apiClient.CompaniesAPI.ListCompanies(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompanies`: ListCompaniesAnalyticsAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListCompaniesAnalyticsAssets200Response**](ListCompaniesAnalyticsAssets200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompaniesAnalyticsAssets

> ListCompaniesAnalyticsAssets200Response ListCompaniesAnalyticsAssets(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get asset performance analytics.     Query params: start_date, end_date, limit, sort_by, include

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
	resp, r, err := apiClient.CompaniesAPI.ListCompaniesAnalyticsAssets(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListCompaniesAnalyticsAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompaniesAnalyticsAssets`: ListCompaniesAnalyticsAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListCompaniesAnalyticsAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesAnalyticsAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListCompaniesAnalyticsAssets200Response**](ListCompaniesAnalyticsAssets200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompaniesAnalyticsCustomers

> ListCompaniesAnalyticsAssets200Response ListCompaniesAnalyticsCustomers(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get customer analytics including growth and segmentation.     Query params: start_date, end_date, group_by, segment_by

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
	resp, r, err := apiClient.CompaniesAPI.ListCompaniesAnalyticsCustomers(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListCompaniesAnalyticsCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompaniesAnalyticsCustomers`: ListCompaniesAnalyticsAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListCompaniesAnalyticsCustomers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesAnalyticsCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListCompaniesAnalyticsAssets200Response**](ListCompaniesAnalyticsAssets200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompaniesAnalyticsDashboard

> ListCompaniesAnalyticsAssets200Response ListCompaniesAnalyticsDashboard(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get comprehensive dashboard analytics for seller dashboard.     Query params: start_date, end_date, period (day/week/month/quarter)

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
	resp, r, err := apiClient.CompaniesAPI.ListCompaniesAnalyticsDashboard(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListCompaniesAnalyticsDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompaniesAnalyticsDashboard`: ListCompaniesAnalyticsAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListCompaniesAnalyticsDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesAnalyticsDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListCompaniesAnalyticsAssets200Response**](ListCompaniesAnalyticsAssets200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompaniesAnalyticsMetrics

> ListCompaniesAnalyticsAssets200Response ListCompaniesAnalyticsMetrics(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get summary metrics (counts only). No date params needed — returns all-time totals.     Performance: < 0.1s

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
	resp, r, err := apiClient.CompaniesAPI.ListCompaniesAnalyticsMetrics(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListCompaniesAnalyticsMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompaniesAnalyticsMetrics`: ListCompaniesAnalyticsAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListCompaniesAnalyticsMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesAnalyticsMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListCompaniesAnalyticsAssets200Response**](ListCompaniesAnalyticsAssets200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompaniesAnalyticsSales

> ListCompaniesAnalyticsAssets200Response ListCompaniesAnalyticsSales(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get sales over time analytics.     Query params: start_date, end_date, group_by (day/week/month/quarter), metrics

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
	resp, r, err := apiClient.CompaniesAPI.ListCompaniesAnalyticsSales(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListCompaniesAnalyticsSales``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompaniesAnalyticsSales`: ListCompaniesAnalyticsAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListCompaniesAnalyticsSales`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesAnalyticsSalesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListCompaniesAnalyticsAssets200Response**](ListCompaniesAnalyticsAssets200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompaniesObjects

> ListCompaniesAnalyticsAssets200Response ListCompaniesObjects(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get all objects (connections, assets) of a single company

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
	resp, r, err := apiClient.CompaniesAPI.ListCompaniesObjects(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListCompaniesObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompaniesObjects`: ListCompaniesAnalyticsAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListCompaniesObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListCompaniesAnalyticsAssets200Response**](ListCompaniesAnalyticsAssets200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompaniesRequestsUsage

> ListCompaniesAnalyticsAssets200Response ListCompaniesRequestsUsage(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get company request usage data for a specific month.

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
	resp, r, err := apiClient.CompaniesAPI.ListCompaniesRequestsUsage(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListCompaniesRequestsUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompaniesRequestsUsage`: ListCompaniesAnalyticsAssets200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListCompaniesRequestsUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesRequestsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListCompaniesAnalyticsAssets200Response**](ListCompaniesAnalyticsAssets200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompanies

> UpdateCompanies200Response UpdateCompanies(ctx, companyId).CompaniesUpdate(companiesUpdate).Execute()

Update an existing company by ID

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
	companiesUpdate := *openapiclient.NewCompaniesUpdate() // CompaniesUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.UpdateCompanies(context.Background(), companyId).CompaniesUpdate(companiesUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.UpdateCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCompanies`: UpdateCompanies200Response
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.UpdateCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companiesUpdate** | [**CompaniesUpdate**](CompaniesUpdate.md) |  | 

### Return type

[**UpdateCompanies200Response**](UpdateCompanies200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


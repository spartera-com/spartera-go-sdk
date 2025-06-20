# \UsersAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompaniesCompanyIdUsersGet**](UsersAPI.md#CompaniesCompanyIdUsersGet) | **Get** /companies/{company_id}/users | Get a list of all users in a company
[**CompaniesCompanyIdUsersPost**](UsersAPI.md#CompaniesCompanyIdUsersPost) | **Post** /companies/{company_id}/users | Create a new user
[**CompaniesCompanyIdUsersUserIdDelete**](UsersAPI.md#CompaniesCompanyIdUsersUserIdDelete) | **Delete** /companies/{company_id}/users/{user_id} | Delete single user by ID
[**CompaniesCompanyIdUsersUserIdGet**](UsersAPI.md#CompaniesCompanyIdUsersUserIdGet) | **Get** /companies/{company_id}/users/{user_id} | Get single user by ID
[**CompaniesCompanyIdUsersUserIdPatch**](UsersAPI.md#CompaniesCompanyIdUsersUserIdPatch) | **Patch** /companies/{company_id}/users/{user_id} | Update an existing user by ID
[**MeGet**](UsersAPI.md#MeGet) | **Get** /me | Get current authenticated user&#39;s profile information.              Returns:                 JSON response with user profile data from database



## CompaniesCompanyIdUsersGet

> CompaniesCompanyIdUsersGet200Response CompaniesCompanyIdUsersGet(ctx, companyId).Execute()

Get a list of all users in a company

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
	resp, r, err := apiClient.UsersAPI.CompaniesCompanyIdUsersGet(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CompaniesCompanyIdUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersGet`: CompaniesCompanyIdUsersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CompaniesCompanyIdUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompaniesCompanyIdUsersGet200Response**](CompaniesCompanyIdUsersGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersPost

> CompaniesCompanyIdUsersPost200Response CompaniesCompanyIdUsersPost(ctx, companyId).User(user).Execute()

Create a new user

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
	user := *openapiclient.NewUser("CompanyId_example", "Status_example") // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CompaniesCompanyIdUsersPost(context.Background(), companyId).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CompaniesCompanyIdUsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersPost`: CompaniesCompanyIdUsersPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CompaniesCompanyIdUsersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

### Return type

[**CompaniesCompanyIdUsersPost200Response**](CompaniesCompanyIdUsersPost200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdDelete

> CompaniesCompanyIdUsersUserIdDelete200Response CompaniesCompanyIdUsersUserIdDelete(ctx, companyId, userId).Execute()

Delete single user by ID

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
	resp, r, err := apiClient.UsersAPI.CompaniesCompanyIdUsersUserIdDelete(context.Background(), companyId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CompaniesCompanyIdUsersUserIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdDelete`: CompaniesCompanyIdUsersUserIdDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CompaniesCompanyIdUsersUserIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdUsersUserIdDelete200Response**](CompaniesCompanyIdUsersUserIdDelete200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdGet

> CompaniesCompanyIdUsersUserIdGet200Response CompaniesCompanyIdUsersUserIdGet(ctx, companyId, userId).Execute()

Get single user by ID

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
	resp, r, err := apiClient.UsersAPI.CompaniesCompanyIdUsersUserIdGet(context.Background(), companyId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CompaniesCompanyIdUsersUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdGet`: CompaniesCompanyIdUsersUserIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CompaniesCompanyIdUsersUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdUsersUserIdGet200Response**](CompaniesCompanyIdUsersUserIdGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdPatch

> CompaniesCompanyIdUsersUserIdPatch200Response CompaniesCompanyIdUsersUserIdPatch(ctx, companyId, userId).User(user).Execute()

Update an existing user by ID

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
	user := *openapiclient.NewUser("CompanyId_example", "Status_example") // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CompaniesCompanyIdUsersUserIdPatch(context.Background(), companyId, userId).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CompaniesCompanyIdUsersUserIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdPatch`: CompaniesCompanyIdUsersUserIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CompaniesCompanyIdUsersUserIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **user** | [**User**](User.md) |  | 

### Return type

[**CompaniesCompanyIdUsersUserIdPatch200Response**](CompaniesCompanyIdUsersUserIdPatch200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeGet

> MeGet200Response MeGet(ctx).Execute()

Get current authenticated user's profile information.              Returns:                 JSON response with user profile data from database

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.MeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.MeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MeGet`: MeGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.MeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetRequest struct via the builder pattern


### Return type

[**MeGet200Response**](MeGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \FavoritesAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFavorites**](FavoritesAPI.md#CreateFavorites) | **Post** /companies/{company_id}/users/{user_id}/favorites | POST /companies/{company_id}/users/{user_id}/favorites
[**DeleteFavorites**](FavoritesAPI.md#DeleteFavorites) | **Delete** /companies/{company_id}/users/{user_id}/favorites/{favorite_id} | Delete single favorite by ID (unfavorite an asset)
[**GetFavoritesById**](FavoritesAPI.md#GetFavoritesById) | **Get** /companies/{company_id}/users/{user_id}/favorites | Get a list of all favorites for a specific user
[**GetFavoritesByIdUsers**](FavoritesAPI.md#GetFavoritesByIdUsers) | **Get** /companies/{company_id}/users/{user_id}/favorites/{favorite_id} | Get single favorite by ID
[**GetFavoritesByIdUsersCategory**](FavoritesAPI.md#GetFavoritesByIdUsersCategory) | **Get** /companies/{company_id}/users/{user_id}/favorites/category/{category} | Get all favorites for the specified user in a specific category
[**GetFavoritesByIdUsersCheck**](FavoritesAPI.md#GetFavoritesByIdUsersCheck) | **Get** /companies/{company_id}/users/{user_id}/favorites/check/{asset_id} | Check if the specified user has favorited a specific asset
[**GetFavoritesByIdUsersUncategorized**](FavoritesAPI.md#GetFavoritesByIdUsersUncategorized) | **Get** /companies/{company_id}/users/{user_id}/favorites/uncategorized | Get all favorites for the specified user that don&#39;t have a category
[**UpdateFavorites**](FavoritesAPI.md#UpdateFavorites) | **Patch** /companies/{company_id}/users/{user_id}/favorites/{favorite_id} | Update an existing favorite by ID



## CreateFavorites

> CreateFavorites200Response CreateFavorites(ctx, companyId, userId).FavoritesInput(favoritesInput).Execute()

POST /companies/{company_id}/users/{user_id}/favorites

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
	favoritesInput := *openapiclient.NewFavoritesInput("asset_id_abc123", "company_id_abc123") // FavoritesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.CreateFavorites(context.Background(), companyId, userId).FavoritesInput(favoritesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.CreateFavorites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFavorites`: CreateFavorites200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.CreateFavorites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFavoritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **favoritesInput** | [**FavoritesInput**](FavoritesInput.md) |  | 

### Return type

[**CreateFavorites200Response**](CreateFavorites200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFavorites

> DeleteFavorites200Response DeleteFavorites(ctx, companyId, userId, favoriteId).Execute()

Delete single favorite by ID (unfavorite an asset)

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
	favoriteId := "favoriteId_example" // string | Unique identifier for the Favorite

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.DeleteFavorites(context.Background(), companyId, userId, favoriteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.DeleteFavorites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFavorites`: DeleteFavorites200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.DeleteFavorites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 
**favoriteId** | **string** | Unique identifier for the Favorite | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFavoritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeleteFavorites200Response**](DeleteFavorites200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoritesById

> GetFavoritesById200Response GetFavoritesById(ctx, companyId, userId).Execute()

Get a list of all favorites for a specific user

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
	resp, r, err := apiClient.FavoritesAPI.GetFavoritesById(context.Background(), companyId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.GetFavoritesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoritesById`: GetFavoritesById200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.GetFavoritesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoritesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetFavoritesById200Response**](GetFavoritesById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoritesByIdUsers

> GetFavoritesById200Response GetFavoritesByIdUsers(ctx, companyId, userId, favoriteId).Execute()

Get single favorite by ID

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
	favoriteId := "favoriteId_example" // string | Unique identifier for the Favorite

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.GetFavoritesByIdUsers(context.Background(), companyId, userId, favoriteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.GetFavoritesByIdUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoritesByIdUsers`: GetFavoritesById200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.GetFavoritesByIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 
**favoriteId** | **string** | Unique identifier for the Favorite | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoritesByIdUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetFavoritesById200Response**](GetFavoritesById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoritesByIdUsersCategory

> GetFavoritesById200Response GetFavoritesByIdUsersCategory(ctx, companyId, userId, category).Execute()

Get all favorites for the specified user in a specific category

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
	category := "category_example" // string | Parameter for category

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.GetFavoritesByIdUsersCategory(context.Background(), companyId, userId, category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.GetFavoritesByIdUsersCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoritesByIdUsersCategory`: GetFavoritesById200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.GetFavoritesByIdUsersCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 
**category** | **string** | Parameter for category | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoritesByIdUsersCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetFavoritesById200Response**](GetFavoritesById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoritesByIdUsersCheck

> GetFavoritesById200Response GetFavoritesByIdUsersCheck(ctx, companyId, userId, assetId).Execute()

Check if the specified user has favorited a specific asset

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
	resp, r, err := apiClient.FavoritesAPI.GetFavoritesByIdUsersCheck(context.Background(), companyId, userId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.GetFavoritesByIdUsersCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoritesByIdUsersCheck`: GetFavoritesById200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.GetFavoritesByIdUsersCheck`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetFavoritesByIdUsersCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetFavoritesById200Response**](GetFavoritesById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoritesByIdUsersUncategorized

> GetFavoritesById200Response GetFavoritesByIdUsersUncategorized(ctx, companyId, userId).Execute()

Get all favorites for the specified user that don't have a category

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
	resp, r, err := apiClient.FavoritesAPI.GetFavoritesByIdUsersUncategorized(context.Background(), companyId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.GetFavoritesByIdUsersUncategorized``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoritesByIdUsersUncategorized`: GetFavoritesById200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.GetFavoritesByIdUsersUncategorized`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoritesByIdUsersUncategorizedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetFavoritesById200Response**](GetFavoritesById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFavorites

> UpdateFavorites200Response UpdateFavorites(ctx, companyId, userId, favoriteId).FavoritesUpdate(favoritesUpdate).Execute()

Update an existing favorite by ID

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
	favoriteId := "favoriteId_example" // string | Unique identifier for the Favorite
	favoritesUpdate := *openapiclient.NewFavoritesUpdate() // FavoritesUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.UpdateFavorites(context.Background(), companyId, userId, favoriteId).FavoritesUpdate(favoritesUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.UpdateFavorites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFavorites`: UpdateFavorites200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.UpdateFavorites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**userId** | **string** | Unique identifier for the User | 
**favoriteId** | **string** | Unique identifier for the Favorite | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFavoritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **favoritesUpdate** | [**FavoritesUpdate**](FavoritesUpdate.md) |  | 

### Return type

[**UpdateFavorites200Response**](UpdateFavorites200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


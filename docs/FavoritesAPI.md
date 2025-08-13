# \FavoritesAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet) | **Get** /companies/{company_id}/users/{user_id}/favorites/category/{category} | Get all favorites for the specified user in a specific category
[**CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet) | **Get** /companies/{company_id}/users/{user_id}/favorites/check/{asset_id} | Check if the specified user has favorited a specific asset
[**CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete) | **Delete** /companies/{company_id}/users/{user_id}/favorites/{favorite_id} | Delete single favorite by ID (unfavorite an asset)
[**CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet) | **Get** /companies/{company_id}/users/{user_id}/favorites/{favorite_id} | Get single favorite by ID
[**CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch) | **Patch** /companies/{company_id}/users/{user_id}/favorites/{favorite_id} | Update an existing favorite by ID
[**CompaniesCompanyIdUsersUserIdFavoritesGet**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesGet) | **Get** /companies/{company_id}/users/{user_id}/favorites | Get a list of all favorites for a specific user
[**CompaniesCompanyIdUsersUserIdFavoritesPost**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesPost) | **Post** /companies/{company_id}/users/{user_id}/favorites | POST /companies/{company_id}/users/{user_id}/favorites
[**CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet) | **Get** /companies/{company_id}/users/{user_id}/favorites/uncategorized | Get all favorites for the specified user that don&#39;t have a category



## CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet

> CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet200Response CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet(ctx, companyId, userId, category).Execute()

Get all favorites for the specified user in a specific category

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
	category := "category_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet(context.Background(), companyId, userId, category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet`: CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 
**category** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet200Response**](CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet

> CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet200Response CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet(ctx, companyId, userId, assetId).Execute()

Check if the specified user has favorited a specific asset

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
	resp, r, err := apiClient.FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet(context.Background(), companyId, userId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet`: CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet200Response**](CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete

> CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete200Response CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete(ctx, companyId, userId, favoriteId).Execute()

Delete single favorite by ID (unfavorite an asset)

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
	favoriteId := "favoriteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete(context.Background(), companyId, userId, favoriteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete`: CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 
**favoriteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete200Response**](CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet

> CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet200Response CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet(ctx, companyId, userId, favoriteId).Execute()

Get single favorite by ID

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
	favoriteId := "favoriteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet(context.Background(), companyId, userId, favoriteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet`: CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 
**favoriteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet200Response**](CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch

> CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch200Response CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch(ctx, companyId, userId, favoriteId).FavoritesUpdate(favoritesUpdate).Execute()

Update an existing favorite by ID

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
	favoriteId := "favoriteId_example" // string | 
	favoritesUpdate := *openapiclient.NewFavoritesUpdate() // FavoritesUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch(context.Background(), companyId, userId, favoriteId).FavoritesUpdate(favoritesUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch`: CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 
**favoriteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **favoritesUpdate** | [**FavoritesUpdate**](FavoritesUpdate.md) |  | 

### Return type

[**CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch200Response**](CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesGet

> CompaniesCompanyIdUsersUserIdFavoritesGet200Response CompaniesCompanyIdUsersUserIdFavoritesGet(ctx, companyId, userId).Execute()

Get a list of all favorites for a specific user

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
	resp, r, err := apiClient.FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesGet(context.Background(), companyId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdFavoritesGet`: CompaniesCompanyIdUsersUserIdFavoritesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdFavoritesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdUsersUserIdFavoritesGet200Response**](CompaniesCompanyIdUsersUserIdFavoritesGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesPost

> CompaniesCompanyIdUsersUserIdFavoritesPost200Response CompaniesCompanyIdUsersUserIdFavoritesPost(ctx, companyId, userId).FavoritesInput(favoritesInput).Execute()

POST /companies/{company_id}/users/{user_id}/favorites

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
	favoritesInput := *openapiclient.NewFavoritesInput("AssetId_example", "CompanyId_example") // FavoritesInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesPost(context.Background(), companyId, userId).FavoritesInput(favoritesInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdFavoritesPost`: CompaniesCompanyIdUsersUserIdFavoritesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdFavoritesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **favoritesInput** | [**FavoritesInput**](FavoritesInput.md) |  | 

### Return type

[**CompaniesCompanyIdUsersUserIdFavoritesPost200Response**](CompaniesCompanyIdUsersUserIdFavoritesPost200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet

> CompaniesCompanyIdUsersUserIdFavoritesGet200Response CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet(ctx, companyId, userId).Execute()

Get all favorites for the specified user that don't have a category

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
	resp, r, err := apiClient.FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet(context.Background(), companyId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet`: CompaniesCompanyIdUsersUserIdFavoritesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompaniesCompanyIdUsersUserIdFavoritesUncategorizedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CompaniesCompanyIdUsersUserIdFavoritesGet200Response**](CompaniesCompanyIdUsersUserIdFavoritesGet200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


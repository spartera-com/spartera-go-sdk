# \FavoritesAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet) | **Get** /companies/{company_id}/users/{user_id}/favorites/category/{category} | Get all favorites for the specified user in a specific category
[**CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet) | **Get** /companies/{company_id}/users/{user_id}/favorites/check/{asset_id} | Check if the specified user has favorited a specific asset     Returns the favorite record if it exists, or empty result if not     Useful for UI to show/hide favorite button states
[**CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete) | **Delete** /companies/{company_id}/users/{user_id}/favorites/{favorite_id} | Delete single favorite by ID (unfavorite an asset)
[**CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet) | **Get** /companies/{company_id}/users/{user_id}/favorites/{favorite_id} | Get single favorite by ID
[**CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch) | **Patch** /companies/{company_id}/users/{user_id}/favorites/{favorite_id} | Update an existing favorite by ID     Can update notes, category, priority
[**CompaniesCompanyIdUsersUserIdFavoritesGet**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesGet) | **Get** /companies/{company_id}/users/{user_id}/favorites | Get a list of all favorites for a specific user     Query params:     - category: filter by category     - sort: sort field (priority, date_created, etc.)     - order: sort direction (asc, desc)
[**CompaniesCompanyIdUsersUserIdFavoritesPost**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesPost) | **Post** /companies/{company_id}/users/{user_id}/favorites | POST /companies/{company_id}/users/{user_id}/favorites
[**CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet**](FavoritesAPI.md#CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet) | **Get** /companies/{company_id}/users/{user_id}/favorites/uncategorized | Get all favorites for the specified user that don&#39;t have a category



## CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet

> map[string]interface{} CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet(ctx, companyId, userId, category).Execute()

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
	// response from `CompaniesCompanyIdUsersUserIdFavoritesCategoryCategoryGet`: map[string]interface{}
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

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet

> map[string]interface{} CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet(ctx, companyId, userId, assetId).Execute()

Check if the specified user has favorited a specific asset     Returns the favorite record if it exists, or empty result if not     Useful for UI to show/hide favorite button states

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
	// response from `CompaniesCompanyIdUsersUserIdFavoritesCheckAssetIdGet`: map[string]interface{}
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

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete

> map[string]interface{} CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete(ctx, companyId, userId, favoriteId).Execute()

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
	// response from `CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdDelete`: map[string]interface{}
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

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet

> map[string]interface{} CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet(ctx, companyId, userId, favoriteId).Execute()

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
	// response from `CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdGet`: map[string]interface{}
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

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch

> map[string]interface{} CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch(ctx, companyId, userId, favoriteId).Favorite(favorite).Execute()

Update an existing favorite by ID     Can update notes, category, priority

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
	favorite := *openapiclient.NewFavorite("AssetId_example", "CompanyId_example", "Priority_example") // Favorite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch(context.Background(), companyId, userId, favoriteId).Favorite(favorite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdFavoritesFavoriteIdPatch`: map[string]interface{}
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



 **favorite** | [**Favorite**](Favorite.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesGet

> map[string]interface{} CompaniesCompanyIdUsersUserIdFavoritesGet(ctx, companyId, userId).Execute()

Get a list of all favorites for a specific user     Query params:     - category: filter by category     - sort: sort field (priority, date_created, etc.)     - order: sort direction (asc, desc)

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
	// response from `CompaniesCompanyIdUsersUserIdFavoritesGet`: map[string]interface{}
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

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesPost

> map[string]interface{} CompaniesCompanyIdUsersUserIdFavoritesPost(ctx, companyId, userId).Favorite(favorite).Execute()

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
	favorite := *openapiclient.NewFavorite("AssetId_example", "CompanyId_example", "Priority_example") // Favorite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesPost(context.Background(), companyId, userId).Favorite(favorite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.CompaniesCompanyIdUsersUserIdFavoritesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompaniesCompanyIdUsersUserIdFavoritesPost`: map[string]interface{}
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


 **favorite** | [**Favorite**](Favorite.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet

> map[string]interface{} CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet(ctx, companyId, userId).Execute()

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
	// response from `CompaniesCompanyIdUsersUserIdFavoritesUncategorizedGet`: map[string]interface{}
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

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


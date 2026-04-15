# \PostsAPI

All URIs are relative to *https://api.spartera.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePosts**](PostsAPI.md#CreatePosts) | **Post** /companies/{company_id}/posts | POST /companies/{company_id}/posts
[**CreatePosts2**](PostsAPI.md#CreatePosts2) | **Post** /companies/{company_id}/posts/{post_id}/publish/{integration_id} | Publish a post to an external platform via an integration.      Args:         post_id: ID of the post to publish         integration_id: ID of the integration to use (from postgen_integrations)      Returns:         Publication record with external_url and external_post_id
[**CreatePostsPublish**](PostsAPI.md#CreatePostsPublish) | **Post** /companies/{company_id}/posts/{post_id}/publish | Publish a post (make it publicly visible).
[**CreatePostsUnpublish**](PostsAPI.md#CreatePostsUnpublish) | **Post** /companies/{company_id}/posts/{post_id}/unpublish | Unpublish a post (make it private/draft again).
[**CreatePostsView**](PostsAPI.md#CreatePostsView) | **Post** /companies/{company_id}/posts/{post_id}/view | Increment view count for a post.     Public endpoint (no authentication required).
[**DeletePosts**](PostsAPI.md#DeletePosts) | **Delete** /companies/{company_id}/posts/{post_id} | Delete single post by ID.
[**GetPostsById**](PostsAPI.md#GetPostsById) | **Get** /companies/{company_id}/posts/{post_id} | Get single post by ID.
[**GetPostsByIdPublications**](PostsAPI.md#GetPostsByIdPublications) | **Get** /companies/{company_id}/posts/{post_id}/publications | Get all publications for a post.     Shows where this post has been published to external platforms.      Returns:         Array of publication records with platform, URL, status
[**ListPosts**](PostsAPI.md#ListPosts) | **Get** /companies/{company_id}/posts | Get a list of all posts for the user&#39;s company.     Supports filtering, sorting, pagination.
[**ListPostsSummary**](PostsAPI.md#ListPostsSummary) | **Get** /companies/{company_id}/posts/summary | GET /companies/{company_id}/posts/summary
[**UpdatePosts**](PostsAPI.md#UpdatePosts) | **Patch** /companies/{company_id}/posts/{post_id} | Update an existing post by ID.      Note: last_edited_at is automatically updated.



## CreatePosts

> CreatePosts200Response CreatePosts(ctx, companyId).PostsInput(postsInput).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

POST /companies/{company_id}/posts

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
	postsInput := *openapiclient.NewPostsInput("user_id_abc123", "company_id_abc123", "example_value", "example_value") // PostsInput | 
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	limit := int32(56) // int32 | Number of items per page (optional) (default to 20)
	sortBy := "sortBy_example" // string | Field to sort by (optional)
	sortOrder := "sortOrder_example" // string | Sort order (ascending or descending) (optional) (default to "desc")
	search := "search_example" // string | Search term to filter results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.CreatePosts(context.Background(), companyId).PostsInput(postsInput).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.CreatePosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePosts`: CreatePosts200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.CreatePosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postsInput** | [**PostsInput**](PostsInput.md) |  | 
 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**CreatePosts200Response**](CreatePosts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePosts2

> CreatePosts200Response CreatePosts2(ctx, companyId, postId, integrationId).PostsInput(postsInput).Execute()

Publish a post to an external platform via an integration.      Args:         post_id: ID of the post to publish         integration_id: ID of the integration to use (from postgen_integrations)      Returns:         Publication record with external_url and external_post_id

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
	postId := "postId_example" // string | Unique identifier for the Post
	integrationId := "integrationId_example" // string | Unique identifier for the Integration
	postsInput := *openapiclient.NewPostsInput("user_id_abc123", "company_id_abc123", "example_value", "example_value") // PostsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.CreatePosts2(context.Background(), companyId, postId, integrationId).PostsInput(postsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.CreatePosts2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePosts2`: CreatePosts200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.CreatePosts2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**postId** | **string** | Unique identifier for the Post | 
**integrationId** | **string** | Unique identifier for the Integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePosts2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postsInput** | [**PostsInput**](PostsInput.md) |  | 

### Return type

[**CreatePosts200Response**](CreatePosts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePostsPublish

> CreatePosts200Response CreatePostsPublish(ctx, companyId, postId).PostsInput(postsInput).Execute()

Publish a post (make it publicly visible).

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
	postId := "postId_example" // string | Unique identifier for the Post
	postsInput := *openapiclient.NewPostsInput("user_id_abc123", "company_id_abc123", "example_value", "example_value") // PostsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.CreatePostsPublish(context.Background(), companyId, postId).PostsInput(postsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.CreatePostsPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePostsPublish`: CreatePosts200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.CreatePostsPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**postId** | **string** | Unique identifier for the Post | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostsPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postsInput** | [**PostsInput**](PostsInput.md) |  | 

### Return type

[**CreatePosts200Response**](CreatePosts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePostsUnpublish

> CreatePosts200Response CreatePostsUnpublish(ctx, companyId, postId).PostsInput(postsInput).Execute()

Unpublish a post (make it private/draft again).

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
	postId := "postId_example" // string | Unique identifier for the Post
	postsInput := *openapiclient.NewPostsInput("user_id_abc123", "company_id_abc123", "example_value", "example_value") // PostsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.CreatePostsUnpublish(context.Background(), companyId, postId).PostsInput(postsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.CreatePostsUnpublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePostsUnpublish`: CreatePosts200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.CreatePostsUnpublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**postId** | **string** | Unique identifier for the Post | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostsUnpublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postsInput** | [**PostsInput**](PostsInput.md) |  | 

### Return type

[**CreatePosts200Response**](CreatePosts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePostsView

> CreatePosts200Response CreatePostsView(ctx, companyId, postId).PostsInput(postsInput).Execute()

Increment view count for a post.     Public endpoint (no authentication required).

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
	postId := "postId_example" // string | Unique identifier for the Post
	postsInput := *openapiclient.NewPostsInput("user_id_abc123", "company_id_abc123", "example_value", "example_value") // PostsInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.CreatePostsView(context.Background(), companyId, postId).PostsInput(postsInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.CreatePostsView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePostsView`: CreatePosts200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.CreatePostsView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**postId** | **string** | Unique identifier for the Post | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostsViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postsInput** | [**PostsInput**](PostsInput.md) |  | 

### Return type

[**CreatePosts200Response**](CreatePosts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePosts

> DeletePosts200Response DeletePosts(ctx, companyId, postId).Execute()

Delete single post by ID.

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
	postId := "postId_example" // string | Unique identifier for the Post

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.DeletePosts(context.Background(), companyId, postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.DeletePosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePosts`: DeletePosts200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.DeletePosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**postId** | **string** | Unique identifier for the Post | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeletePosts200Response**](DeletePosts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostsById

> GetPostsById200Response GetPostsById(ctx, companyId, postId).Execute()

Get single post by ID.

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
	postId := "postId_example" // string | Unique identifier for the Post

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.GetPostsById(context.Background(), companyId, postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.GetPostsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostsById`: GetPostsById200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.GetPostsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**postId** | **string** | Unique identifier for the Post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetPostsById200Response**](GetPostsById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostsByIdPublications

> GetPostsById200Response GetPostsByIdPublications(ctx, companyId, postId).Execute()

Get all publications for a post.     Shows where this post has been published to external platforms.      Returns:         Array of publication records with platform, URL, status

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
	postId := "postId_example" // string | Unique identifier for the Post

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.GetPostsByIdPublications(context.Background(), companyId, postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.GetPostsByIdPublications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostsByIdPublications`: GetPostsById200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.GetPostsByIdPublications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**postId** | **string** | Unique identifier for the Post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostsByIdPublicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetPostsById200Response**](GetPostsById200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPosts

> ListPosts200Response ListPosts(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

Get a list of all posts for the user's company.     Supports filtering, sorting, pagination.

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
	resp, r, err := apiClient.PostsAPI.ListPosts(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.ListPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPosts`: ListPosts200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.ListPosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListPosts200Response**](ListPosts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPostsSummary

> ListPosts200Response ListPostsSummary(ctx, companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()

GET /companies/{company_id}/posts/summary

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
	resp, r, err := apiClient.PostsAPI.ListPostsSummary(context.Background(), companyId).Page(page).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.ListPostsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPostsSummary`: ListPosts200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.ListPostsSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPostsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **limit** | **int32** | Number of items per page | [default to 20]
 **sortBy** | **string** | Field to sort by | 
 **sortOrder** | **string** | Sort order (ascending or descending) | [default to &quot;desc&quot;]
 **search** | **string** | Search term to filter results | 

### Return type

[**ListPosts200Response**](ListPosts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePosts

> UpdatePosts200Response UpdatePosts(ctx, companyId, postId).PostsUpdate(postsUpdate).Execute()

Update an existing post by ID.      Note: last_edited_at is automatically updated.

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
	postId := "postId_example" // string | Unique identifier for the Post
	postsUpdate := *openapiclient.NewPostsUpdate() // PostsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.UpdatePosts(context.Background(), companyId, postId).PostsUpdate(postsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.UpdatePosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePosts`: UpdatePosts200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.UpdatePosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | Unique identifier for the Company | 
**postId** | **string** | Unique identifier for the Post | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postsUpdate** | [**PostsUpdate**](PostsUpdate.md) |  | 

### Return type

[**UpdatePosts200Response**](UpdatePosts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


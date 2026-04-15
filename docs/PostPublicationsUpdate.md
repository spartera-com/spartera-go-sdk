# PostPublicationsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostId** | Pointer to **string** | Post that was published | [optional] 
**CompanyId** | Pointer to **string** | Company this publication belongs to | [optional] 
**IntegrationId** | Pointer to **string** | Integration used for publishing | [optional] 
**Platform** | Pointer to **string** | Platform identifier (beehiiv, wordpress, etc) | [optional] 
**ExternalPostId** | Pointer to **string** | ID of the post on the external platform | [optional] 
**ExternalPostUrl** | Pointer to **string** | URL to view the post on the external platform | [optional] 
**PublishedAt** | Pointer to **time.Time** | When the post was published to this platform | [optional] 
**Status** | Pointer to **string** | Status: published, failed, deleted | [optional] 
**ErrorMessage** | Pointer to **string** | Error message if publication failed | [optional] 

## Methods

### NewPostPublicationsUpdate

`func NewPostPublicationsUpdate() *PostPublicationsUpdate`

NewPostPublicationsUpdate instantiates a new PostPublicationsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPublicationsUpdateWithDefaults

`func NewPostPublicationsUpdateWithDefaults() *PostPublicationsUpdate`

NewPostPublicationsUpdateWithDefaults instantiates a new PostPublicationsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostId

`func (o *PostPublicationsUpdate) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *PostPublicationsUpdate) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *PostPublicationsUpdate) SetPostId(v string)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *PostPublicationsUpdate) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetCompanyId

`func (o *PostPublicationsUpdate) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PostPublicationsUpdate) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PostPublicationsUpdate) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *PostPublicationsUpdate) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *PostPublicationsUpdate) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *PostPublicationsUpdate) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *PostPublicationsUpdate) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *PostPublicationsUpdate) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetPlatform

`func (o *PostPublicationsUpdate) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PostPublicationsUpdate) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PostPublicationsUpdate) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PostPublicationsUpdate) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetExternalPostId

`func (o *PostPublicationsUpdate) GetExternalPostId() string`

GetExternalPostId returns the ExternalPostId field if non-nil, zero value otherwise.

### GetExternalPostIdOk

`func (o *PostPublicationsUpdate) GetExternalPostIdOk() (*string, bool)`

GetExternalPostIdOk returns a tuple with the ExternalPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPostId

`func (o *PostPublicationsUpdate) SetExternalPostId(v string)`

SetExternalPostId sets ExternalPostId field to given value.

### HasExternalPostId

`func (o *PostPublicationsUpdate) HasExternalPostId() bool`

HasExternalPostId returns a boolean if a field has been set.

### GetExternalPostUrl

`func (o *PostPublicationsUpdate) GetExternalPostUrl() string`

GetExternalPostUrl returns the ExternalPostUrl field if non-nil, zero value otherwise.

### GetExternalPostUrlOk

`func (o *PostPublicationsUpdate) GetExternalPostUrlOk() (*string, bool)`

GetExternalPostUrlOk returns a tuple with the ExternalPostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPostUrl

`func (o *PostPublicationsUpdate) SetExternalPostUrl(v string)`

SetExternalPostUrl sets ExternalPostUrl field to given value.

### HasExternalPostUrl

`func (o *PostPublicationsUpdate) HasExternalPostUrl() bool`

HasExternalPostUrl returns a boolean if a field has been set.

### GetPublishedAt

`func (o *PostPublicationsUpdate) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PostPublicationsUpdate) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PostPublicationsUpdate) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *PostPublicationsUpdate) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetStatus

`func (o *PostPublicationsUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostPublicationsUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostPublicationsUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostPublicationsUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PostPublicationsUpdate) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PostPublicationsUpdate) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PostPublicationsUpdate) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PostPublicationsUpdate) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



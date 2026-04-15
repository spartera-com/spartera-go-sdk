# PostPublicationsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostId** | **string** | Post that was published | 
**CompanyId** | **string** | Company this publication belongs to | 
**IntegrationId** | Pointer to **string** | Integration used for publishing | [optional] 
**Platform** | **string** | Platform identifier (beehiiv, wordpress, etc) | 
**ExternalPostId** | Pointer to **string** | ID of the post on the external platform | [optional] 
**ExternalPostUrl** | Pointer to **string** | URL to view the post on the external platform | [optional] 
**PublishedAt** | Pointer to **time.Time** | When the post was published to this platform | [optional] 
**Status** | Pointer to **string** | Status: published, failed, deleted | [optional] 
**ErrorMessage** | Pointer to **string** | Error message if publication failed | [optional] 

## Methods

### NewPostPublicationsInput

`func NewPostPublicationsInput(postId string, companyId string, platform string, ) *PostPublicationsInput`

NewPostPublicationsInput instantiates a new PostPublicationsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPublicationsInputWithDefaults

`func NewPostPublicationsInputWithDefaults() *PostPublicationsInput`

NewPostPublicationsInputWithDefaults instantiates a new PostPublicationsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostId

`func (o *PostPublicationsInput) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *PostPublicationsInput) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *PostPublicationsInput) SetPostId(v string)`

SetPostId sets PostId field to given value.


### GetCompanyId

`func (o *PostPublicationsInput) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PostPublicationsInput) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PostPublicationsInput) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetIntegrationId

`func (o *PostPublicationsInput) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *PostPublicationsInput) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *PostPublicationsInput) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *PostPublicationsInput) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetPlatform

`func (o *PostPublicationsInput) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PostPublicationsInput) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PostPublicationsInput) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetExternalPostId

`func (o *PostPublicationsInput) GetExternalPostId() string`

GetExternalPostId returns the ExternalPostId field if non-nil, zero value otherwise.

### GetExternalPostIdOk

`func (o *PostPublicationsInput) GetExternalPostIdOk() (*string, bool)`

GetExternalPostIdOk returns a tuple with the ExternalPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPostId

`func (o *PostPublicationsInput) SetExternalPostId(v string)`

SetExternalPostId sets ExternalPostId field to given value.

### HasExternalPostId

`func (o *PostPublicationsInput) HasExternalPostId() bool`

HasExternalPostId returns a boolean if a field has been set.

### GetExternalPostUrl

`func (o *PostPublicationsInput) GetExternalPostUrl() string`

GetExternalPostUrl returns the ExternalPostUrl field if non-nil, zero value otherwise.

### GetExternalPostUrlOk

`func (o *PostPublicationsInput) GetExternalPostUrlOk() (*string, bool)`

GetExternalPostUrlOk returns a tuple with the ExternalPostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPostUrl

`func (o *PostPublicationsInput) SetExternalPostUrl(v string)`

SetExternalPostUrl sets ExternalPostUrl field to given value.

### HasExternalPostUrl

`func (o *PostPublicationsInput) HasExternalPostUrl() bool`

HasExternalPostUrl returns a boolean if a field has been set.

### GetPublishedAt

`func (o *PostPublicationsInput) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PostPublicationsInput) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PostPublicationsInput) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *PostPublicationsInput) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetStatus

`func (o *PostPublicationsInput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostPublicationsInput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostPublicationsInput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostPublicationsInput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PostPublicationsInput) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PostPublicationsInput) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PostPublicationsInput) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PostPublicationsInput) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



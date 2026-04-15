# PostPublications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** | Optional. | [optional] 
**LastUpdated** | Pointer to **time.Time** | Optional. | [optional] 
**PublicationId** | Pointer to **string** | Unique identifier. | [optional] 
**PostId** | **string** | Post that was published | 
**CompanyId** | **string** | Company this publication belongs to | 
**IntegrationId** | Pointer to **string** | Integration used for publishing | [optional] 
**Platform** | **string** | Platform identifier (beehiiv, wordpress, etc) | 
**ExternalPostId** | Pointer to **string** | ID of the post on the external platform | [optional] 
**ExternalPostUrl** | Pointer to **string** | URL to view the post on the external platform | [optional] 
**PublishedAt** | Pointer to **time.Time** | When the post was published to this platform | [optional] 
**Status** | **string** | Status: published, failed, deleted | 
**ErrorMessage** | Pointer to **string** | Error message if publication failed | [optional] 

## Methods

### NewPostPublications

`func NewPostPublications(postId string, companyId string, platform string, status string, ) *PostPublications`

NewPostPublications instantiates a new PostPublications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPublicationsWithDefaults

`func NewPostPublicationsWithDefaults() *PostPublications`

NewPostPublicationsWithDefaults instantiates a new PostPublications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *PostPublications) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *PostPublications) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *PostPublications) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *PostPublications) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PostPublications) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PostPublications) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PostPublications) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PostPublications) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPublicationId

`func (o *PostPublications) GetPublicationId() string`

GetPublicationId returns the PublicationId field if non-nil, zero value otherwise.

### GetPublicationIdOk

`func (o *PostPublications) GetPublicationIdOk() (*string, bool)`

GetPublicationIdOk returns a tuple with the PublicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationId

`func (o *PostPublications) SetPublicationId(v string)`

SetPublicationId sets PublicationId field to given value.

### HasPublicationId

`func (o *PostPublications) HasPublicationId() bool`

HasPublicationId returns a boolean if a field has been set.

### GetPostId

`func (o *PostPublications) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *PostPublications) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *PostPublications) SetPostId(v string)`

SetPostId sets PostId field to given value.


### GetCompanyId

`func (o *PostPublications) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PostPublications) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PostPublications) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetIntegrationId

`func (o *PostPublications) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *PostPublications) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *PostPublications) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *PostPublications) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetPlatform

`func (o *PostPublications) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PostPublications) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PostPublications) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetExternalPostId

`func (o *PostPublications) GetExternalPostId() string`

GetExternalPostId returns the ExternalPostId field if non-nil, zero value otherwise.

### GetExternalPostIdOk

`func (o *PostPublications) GetExternalPostIdOk() (*string, bool)`

GetExternalPostIdOk returns a tuple with the ExternalPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPostId

`func (o *PostPublications) SetExternalPostId(v string)`

SetExternalPostId sets ExternalPostId field to given value.

### HasExternalPostId

`func (o *PostPublications) HasExternalPostId() bool`

HasExternalPostId returns a boolean if a field has been set.

### GetExternalPostUrl

`func (o *PostPublications) GetExternalPostUrl() string`

GetExternalPostUrl returns the ExternalPostUrl field if non-nil, zero value otherwise.

### GetExternalPostUrlOk

`func (o *PostPublications) GetExternalPostUrlOk() (*string, bool)`

GetExternalPostUrlOk returns a tuple with the ExternalPostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPostUrl

`func (o *PostPublications) SetExternalPostUrl(v string)`

SetExternalPostUrl sets ExternalPostUrl field to given value.

### HasExternalPostUrl

`func (o *PostPublications) HasExternalPostUrl() bool`

HasExternalPostUrl returns a boolean if a field has been set.

### GetPublishedAt

`func (o *PostPublications) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PostPublications) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PostPublications) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *PostPublications) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetStatus

`func (o *PostPublications) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostPublications) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostPublications) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *PostPublications) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PostPublications) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PostPublications) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PostPublications) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PostgenIntegrations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** | Optional. | [optional] 
**LastUpdated** | Pointer to **time.Time** | Optional. | [optional] 
**IntegrationId** | Pointer to **string** | Unique identifier. | [optional] 
**CompanyId** | **string** | Company this integration belongs to | 
**UserId** | **string** | User who created this integration | 
**Platform** | **string** | Platform identifier (beehiiv, wordpress, medium, etc) | 
**PlatformName** | **string** | Display name of the platform | 
**CredentialsSecretName** | **string** | GCP Secret Manager secret name containing encrypted credentials | 
**IsActive** | **bool** | Whether this integration is currently active | 
**LastUsedAt** | Pointer to **time.Time** | Last time this integration was used for publishing | [optional] 

## Methods

### NewPostgenIntegrations

`func NewPostgenIntegrations(companyId string, userId string, platform string, platformName string, credentialsSecretName string, isActive bool, ) *PostgenIntegrations`

NewPostgenIntegrations instantiates a new PostgenIntegrations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgenIntegrationsWithDefaults

`func NewPostgenIntegrationsWithDefaults() *PostgenIntegrations`

NewPostgenIntegrationsWithDefaults instantiates a new PostgenIntegrations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *PostgenIntegrations) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *PostgenIntegrations) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *PostgenIntegrations) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *PostgenIntegrations) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PostgenIntegrations) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PostgenIntegrations) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PostgenIntegrations) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PostgenIntegrations) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetIntegrationId

`func (o *PostgenIntegrations) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *PostgenIntegrations) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *PostgenIntegrations) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *PostgenIntegrations) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetCompanyId

`func (o *PostgenIntegrations) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PostgenIntegrations) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PostgenIntegrations) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetUserId

`func (o *PostgenIntegrations) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostgenIntegrations) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostgenIntegrations) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPlatform

`func (o *PostgenIntegrations) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PostgenIntegrations) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PostgenIntegrations) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPlatformName

`func (o *PostgenIntegrations) GetPlatformName() string`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *PostgenIntegrations) GetPlatformNameOk() (*string, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *PostgenIntegrations) SetPlatformName(v string)`

SetPlatformName sets PlatformName field to given value.


### GetCredentialsSecretName

`func (o *PostgenIntegrations) GetCredentialsSecretName() string`

GetCredentialsSecretName returns the CredentialsSecretName field if non-nil, zero value otherwise.

### GetCredentialsSecretNameOk

`func (o *PostgenIntegrations) GetCredentialsSecretNameOk() (*string, bool)`

GetCredentialsSecretNameOk returns a tuple with the CredentialsSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsSecretName

`func (o *PostgenIntegrations) SetCredentialsSecretName(v string)`

SetCredentialsSecretName sets CredentialsSecretName field to given value.


### GetIsActive

`func (o *PostgenIntegrations) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PostgenIntegrations) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PostgenIntegrations) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLastUsedAt

`func (o *PostgenIntegrations) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *PostgenIntegrations) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *PostgenIntegrations) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *PostgenIntegrations) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



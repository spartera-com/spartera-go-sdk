# PostgenIntegrationsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** | Company this integration belongs to | [optional] 
**UserId** | Pointer to **string** | User who created this integration | [optional] 
**Platform** | Pointer to **string** | Platform identifier (beehiiv, wordpress, medium, etc) | [optional] 
**PlatformName** | Pointer to **string** | Display name of the platform | [optional] 
**CredentialsSecretName** | Pointer to **string** | GCP Secret Manager secret name containing encrypted credentials | [optional] 
**IsActive** | Pointer to **bool** | Whether this integration is currently active | [optional] 
**LastUsedAt** | Pointer to **time.Time** | Last time this integration was used for publishing | [optional] 

## Methods

### NewPostgenIntegrationsUpdate

`func NewPostgenIntegrationsUpdate() *PostgenIntegrationsUpdate`

NewPostgenIntegrationsUpdate instantiates a new PostgenIntegrationsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgenIntegrationsUpdateWithDefaults

`func NewPostgenIntegrationsUpdateWithDefaults() *PostgenIntegrationsUpdate`

NewPostgenIntegrationsUpdateWithDefaults instantiates a new PostgenIntegrationsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *PostgenIntegrationsUpdate) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PostgenIntegrationsUpdate) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PostgenIntegrationsUpdate) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *PostgenIntegrationsUpdate) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetUserId

`func (o *PostgenIntegrationsUpdate) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostgenIntegrationsUpdate) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostgenIntegrationsUpdate) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PostgenIntegrationsUpdate) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPlatform

`func (o *PostgenIntegrationsUpdate) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PostgenIntegrationsUpdate) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PostgenIntegrationsUpdate) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PostgenIntegrationsUpdate) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformName

`func (o *PostgenIntegrationsUpdate) GetPlatformName() string`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *PostgenIntegrationsUpdate) GetPlatformNameOk() (*string, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *PostgenIntegrationsUpdate) SetPlatformName(v string)`

SetPlatformName sets PlatformName field to given value.

### HasPlatformName

`func (o *PostgenIntegrationsUpdate) HasPlatformName() bool`

HasPlatformName returns a boolean if a field has been set.

### GetCredentialsSecretName

`func (o *PostgenIntegrationsUpdate) GetCredentialsSecretName() string`

GetCredentialsSecretName returns the CredentialsSecretName field if non-nil, zero value otherwise.

### GetCredentialsSecretNameOk

`func (o *PostgenIntegrationsUpdate) GetCredentialsSecretNameOk() (*string, bool)`

GetCredentialsSecretNameOk returns a tuple with the CredentialsSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsSecretName

`func (o *PostgenIntegrationsUpdate) SetCredentialsSecretName(v string)`

SetCredentialsSecretName sets CredentialsSecretName field to given value.

### HasCredentialsSecretName

`func (o *PostgenIntegrationsUpdate) HasCredentialsSecretName() bool`

HasCredentialsSecretName returns a boolean if a field has been set.

### GetIsActive

`func (o *PostgenIntegrationsUpdate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PostgenIntegrationsUpdate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PostgenIntegrationsUpdate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PostgenIntegrationsUpdate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *PostgenIntegrationsUpdate) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *PostgenIntegrationsUpdate) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *PostgenIntegrationsUpdate) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *PostgenIntegrationsUpdate) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



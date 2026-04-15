# PostgenIntegrationsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **string** | Company this integration belongs to | 
**UserId** | **string** | User who created this integration | 
**Platform** | **string** | Platform identifier (beehiiv, wordpress, medium, etc) | 
**PlatformName** | **string** | Display name of the platform | 
**CredentialsSecretName** | **string** | GCP Secret Manager secret name containing encrypted credentials | 
**IsActive** | Pointer to **bool** | Whether this integration is currently active | [optional] 
**LastUsedAt** | Pointer to **time.Time** | Last time this integration was used for publishing | [optional] 

## Methods

### NewPostgenIntegrationsInput

`func NewPostgenIntegrationsInput(companyId string, userId string, platform string, platformName string, credentialsSecretName string, ) *PostgenIntegrationsInput`

NewPostgenIntegrationsInput instantiates a new PostgenIntegrationsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgenIntegrationsInputWithDefaults

`func NewPostgenIntegrationsInputWithDefaults() *PostgenIntegrationsInput`

NewPostgenIntegrationsInputWithDefaults instantiates a new PostgenIntegrationsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *PostgenIntegrationsInput) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PostgenIntegrationsInput) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PostgenIntegrationsInput) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetUserId

`func (o *PostgenIntegrationsInput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostgenIntegrationsInput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostgenIntegrationsInput) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPlatform

`func (o *PostgenIntegrationsInput) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PostgenIntegrationsInput) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PostgenIntegrationsInput) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPlatformName

`func (o *PostgenIntegrationsInput) GetPlatformName() string`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *PostgenIntegrationsInput) GetPlatformNameOk() (*string, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *PostgenIntegrationsInput) SetPlatformName(v string)`

SetPlatformName sets PlatformName field to given value.


### GetCredentialsSecretName

`func (o *PostgenIntegrationsInput) GetCredentialsSecretName() string`

GetCredentialsSecretName returns the CredentialsSecretName field if non-nil, zero value otherwise.

### GetCredentialsSecretNameOk

`func (o *PostgenIntegrationsInput) GetCredentialsSecretNameOk() (*string, bool)`

GetCredentialsSecretNameOk returns a tuple with the CredentialsSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsSecretName

`func (o *PostgenIntegrationsInput) SetCredentialsSecretName(v string)`

SetCredentialsSecretName sets CredentialsSecretName field to given value.


### GetIsActive

`func (o *PostgenIntegrationsInput) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PostgenIntegrationsInput) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PostgenIntegrationsInput) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PostgenIntegrationsInput) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *PostgenIntegrationsInput) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *PostgenIntegrationsInput) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *PostgenIntegrationsInput) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *PostgenIntegrationsInput) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



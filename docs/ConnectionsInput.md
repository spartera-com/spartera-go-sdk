# ConnectionsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**EngineId** | **int64** |  | 
**CompanyId** | **string** |  | 
**CredentialType** | Pointer to **string** | Enum type: CredentialType | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ProviderDomain** | Pointer to **string** |  | [optional] 
**VerifiedUsageAbility** | Pointer to **bool** |  | [optional] 

## Methods

### NewConnectionsInput

`func NewConnectionsInput(engineId int64, companyId string, ) *ConnectionsInput`

NewConnectionsInput instantiates a new ConnectionsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionsInputWithDefaults

`func NewConnectionsInputWithDefaults() *ConnectionsInput`

NewConnectionsInputWithDefaults instantiates a new ConnectionsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ConnectionsInput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ConnectionsInput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ConnectionsInput) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ConnectionsInput) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEngineId

`func (o *ConnectionsInput) GetEngineId() int64`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *ConnectionsInput) GetEngineIdOk() (*int64, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *ConnectionsInput) SetEngineId(v int64)`

SetEngineId sets EngineId field to given value.


### GetCompanyId

`func (o *ConnectionsInput) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ConnectionsInput) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ConnectionsInput) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetCredentialType

`func (o *ConnectionsInput) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *ConnectionsInput) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *ConnectionsInput) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *ConnectionsInput) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetName

`func (o *ConnectionsInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionsInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionsInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionsInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectionsInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectionsInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectionsInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectionsInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProviderDomain

`func (o *ConnectionsInput) GetProviderDomain() string`

GetProviderDomain returns the ProviderDomain field if non-nil, zero value otherwise.

### GetProviderDomainOk

`func (o *ConnectionsInput) GetProviderDomainOk() (*string, bool)`

GetProviderDomainOk returns a tuple with the ProviderDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDomain

`func (o *ConnectionsInput) SetProviderDomain(v string)`

SetProviderDomain sets ProviderDomain field to given value.

### HasProviderDomain

`func (o *ConnectionsInput) HasProviderDomain() bool`

HasProviderDomain returns a boolean if a field has been set.

### GetVerifiedUsageAbility

`func (o *ConnectionsInput) GetVerifiedUsageAbility() bool`

GetVerifiedUsageAbility returns the VerifiedUsageAbility field if non-nil, zero value otherwise.

### GetVerifiedUsageAbilityOk

`func (o *ConnectionsInput) GetVerifiedUsageAbilityOk() (*bool, bool)`

GetVerifiedUsageAbilityOk returns a tuple with the VerifiedUsageAbility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedUsageAbility

`func (o *ConnectionsInput) SetVerifiedUsageAbility(v bool)`

SetVerifiedUsageAbility sets VerifiedUsageAbility field to given value.

### HasVerifiedUsageAbility

`func (o *ConnectionsInput) HasVerifiedUsageAbility() bool`

HasVerifiedUsageAbility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



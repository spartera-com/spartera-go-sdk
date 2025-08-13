# Connections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**EngineId** | **int64** |  | 
**CompanyId** | **string** |  | 
**CredentialType** | Pointer to **string** | Enum type: CredentialType | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ProviderDomain** | Pointer to **string** |  | [optional] 
**VerifiedUsageAbility** | Pointer to **bool** |  | [optional] 

## Methods

### NewConnections

`func NewConnections(engineId int64, companyId string, ) *Connections`

NewConnections instantiates a new Connections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionsWithDefaults

`func NewConnectionsWithDefaults() *Connections`

NewConnectionsWithDefaults instantiates a new Connections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *Connections) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Connections) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Connections) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Connections) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Connections) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Connections) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Connections) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Connections) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetConnectionId

`func (o *Connections) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Connections) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Connections) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Connections) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetUserId

`func (o *Connections) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Connections) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Connections) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Connections) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEngineId

`func (o *Connections) GetEngineId() int64`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *Connections) GetEngineIdOk() (*int64, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *Connections) SetEngineId(v int64)`

SetEngineId sets EngineId field to given value.


### GetCompanyId

`func (o *Connections) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Connections) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Connections) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetCredentialType

`func (o *Connections) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *Connections) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *Connections) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *Connections) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetName

`func (o *Connections) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Connections) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Connections) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Connections) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Connections) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Connections) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Connections) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Connections) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProviderDomain

`func (o *Connections) GetProviderDomain() string`

GetProviderDomain returns the ProviderDomain field if non-nil, zero value otherwise.

### GetProviderDomainOk

`func (o *Connections) GetProviderDomainOk() (*string, bool)`

GetProviderDomainOk returns a tuple with the ProviderDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDomain

`func (o *Connections) SetProviderDomain(v string)`

SetProviderDomain sets ProviderDomain field to given value.

### HasProviderDomain

`func (o *Connections) HasProviderDomain() bool`

HasProviderDomain returns a boolean if a field has been set.

### GetVerifiedUsageAbility

`func (o *Connections) GetVerifiedUsageAbility() bool`

GetVerifiedUsageAbility returns the VerifiedUsageAbility field if non-nil, zero value otherwise.

### GetVerifiedUsageAbilityOk

`func (o *Connections) GetVerifiedUsageAbilityOk() (*bool, bool)`

GetVerifiedUsageAbilityOk returns a tuple with the VerifiedUsageAbility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedUsageAbility

`func (o *Connections) SetVerifiedUsageAbility(v bool)`

SetVerifiedUsageAbility sets VerifiedUsageAbility field to given value.

### HasVerifiedUsageAbility

`func (o *Connections) HasVerifiedUsageAbility() bool`

HasVerifiedUsageAbility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**EngineId** | **string** |  | 
**CompanyId** | **string** |  | 
**CredentialType** | Pointer to **string** |  | [optional] 
**ApiProvider** | Pointer to **string** |  | [optional] 
**ApiEndpoint** | Pointer to **string** |  | [optional] 
**ApiKeyLocation** | Pointer to **string** |  | [optional] 
**ApiKeyParam** | Pointer to **string** |  | [optional] 
**ApiKeyValue** | Pointer to **string** |  | [optional] 
**Visibility** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GcpSecretId** | Pointer to **string** |  | [optional] 
**ProviderDomain** | Pointer to **string** |  | [optional] 
**VerifiedUsageAbility** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **string** |  | [optional] 

## Methods

### NewConnection

`func NewConnection(engineId string, companyId string, visibility string, ) *Connection`

NewConnection instantiates a new Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionWithDefaults

`func NewConnectionWithDefaults() *Connection`

NewConnectionWithDefaults instantiates a new Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *Connection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Connection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Connection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Connection) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetUserId

`func (o *Connection) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Connection) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Connection) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Connection) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEngineId

`func (o *Connection) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *Connection) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *Connection) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.


### GetCompanyId

`func (o *Connection) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Connection) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Connection) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetCredentialType

`func (o *Connection) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *Connection) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *Connection) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *Connection) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetApiProvider

`func (o *Connection) GetApiProvider() string`

GetApiProvider returns the ApiProvider field if non-nil, zero value otherwise.

### GetApiProviderOk

`func (o *Connection) GetApiProviderOk() (*string, bool)`

GetApiProviderOk returns a tuple with the ApiProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvider

`func (o *Connection) SetApiProvider(v string)`

SetApiProvider sets ApiProvider field to given value.

### HasApiProvider

`func (o *Connection) HasApiProvider() bool`

HasApiProvider returns a boolean if a field has been set.

### GetApiEndpoint

`func (o *Connection) GetApiEndpoint() string`

GetApiEndpoint returns the ApiEndpoint field if non-nil, zero value otherwise.

### GetApiEndpointOk

`func (o *Connection) GetApiEndpointOk() (*string, bool)`

GetApiEndpointOk returns a tuple with the ApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEndpoint

`func (o *Connection) SetApiEndpoint(v string)`

SetApiEndpoint sets ApiEndpoint field to given value.

### HasApiEndpoint

`func (o *Connection) HasApiEndpoint() bool`

HasApiEndpoint returns a boolean if a field has been set.

### GetApiKeyLocation

`func (o *Connection) GetApiKeyLocation() string`

GetApiKeyLocation returns the ApiKeyLocation field if non-nil, zero value otherwise.

### GetApiKeyLocationOk

`func (o *Connection) GetApiKeyLocationOk() (*string, bool)`

GetApiKeyLocationOk returns a tuple with the ApiKeyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyLocation

`func (o *Connection) SetApiKeyLocation(v string)`

SetApiKeyLocation sets ApiKeyLocation field to given value.

### HasApiKeyLocation

`func (o *Connection) HasApiKeyLocation() bool`

HasApiKeyLocation returns a boolean if a field has been set.

### GetApiKeyParam

`func (o *Connection) GetApiKeyParam() string`

GetApiKeyParam returns the ApiKeyParam field if non-nil, zero value otherwise.

### GetApiKeyParamOk

`func (o *Connection) GetApiKeyParamOk() (*string, bool)`

GetApiKeyParamOk returns a tuple with the ApiKeyParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyParam

`func (o *Connection) SetApiKeyParam(v string)`

SetApiKeyParam sets ApiKeyParam field to given value.

### HasApiKeyParam

`func (o *Connection) HasApiKeyParam() bool`

HasApiKeyParam returns a boolean if a field has been set.

### GetApiKeyValue

`func (o *Connection) GetApiKeyValue() string`

GetApiKeyValue returns the ApiKeyValue field if non-nil, zero value otherwise.

### GetApiKeyValueOk

`func (o *Connection) GetApiKeyValueOk() (*string, bool)`

GetApiKeyValueOk returns a tuple with the ApiKeyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyValue

`func (o *Connection) SetApiKeyValue(v string)`

SetApiKeyValue sets ApiKeyValue field to given value.

### HasApiKeyValue

`func (o *Connection) HasApiKeyValue() bool`

HasApiKeyValue returns a boolean if a field has been set.

### GetVisibility

`func (o *Connection) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Connection) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Connection) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetName

`func (o *Connection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Connection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Connection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Connection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Connection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Connection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Connection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Connection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGcpSecretId

`func (o *Connection) GetGcpSecretId() string`

GetGcpSecretId returns the GcpSecretId field if non-nil, zero value otherwise.

### GetGcpSecretIdOk

`func (o *Connection) GetGcpSecretIdOk() (*string, bool)`

GetGcpSecretIdOk returns a tuple with the GcpSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpSecretId

`func (o *Connection) SetGcpSecretId(v string)`

SetGcpSecretId sets GcpSecretId field to given value.

### HasGcpSecretId

`func (o *Connection) HasGcpSecretId() bool`

HasGcpSecretId returns a boolean if a field has been set.

### GetProviderDomain

`func (o *Connection) GetProviderDomain() string`

GetProviderDomain returns the ProviderDomain field if non-nil, zero value otherwise.

### GetProviderDomainOk

`func (o *Connection) GetProviderDomainOk() (*string, bool)`

GetProviderDomainOk returns a tuple with the ProviderDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDomain

`func (o *Connection) SetProviderDomain(v string)`

SetProviderDomain sets ProviderDomain field to given value.

### HasProviderDomain

`func (o *Connection) HasProviderDomain() bool`

HasProviderDomain returns a boolean if a field has been set.

### GetVerifiedUsageAbility

`func (o *Connection) GetVerifiedUsageAbility() string`

GetVerifiedUsageAbility returns the VerifiedUsageAbility field if non-nil, zero value otherwise.

### GetVerifiedUsageAbilityOk

`func (o *Connection) GetVerifiedUsageAbilityOk() (*string, bool)`

GetVerifiedUsageAbilityOk returns a tuple with the VerifiedUsageAbility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedUsageAbility

`func (o *Connection) SetVerifiedUsageAbility(v string)`

SetVerifiedUsageAbility sets VerifiedUsageAbility field to given value.

### HasVerifiedUsageAbility

`func (o *Connection) HasVerifiedUsageAbility() bool`

HasVerifiedUsageAbility returns a boolean if a field has been set.

### GetDateCreated

`func (o *Connection) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Connection) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Connection) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Connection) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Connection) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Connection) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Connection) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Connection) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *Connection) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Connection) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Connection) SetActive(v string)`

SetActive sets Active field to given value.

### HasActive

`func (o *Connection) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



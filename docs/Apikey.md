# Apikey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | **string** |  | [readonly] 
**UserId** | Pointer to **string** |  | [optional] 
**CompanyId** | **string** |  | 
**RoleId** | **string** |  | 
**Token** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**ExpirationDateUtc** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **string** |  | [optional] 

## Methods

### NewApikey

`func NewApikey(apiKeyId string, companyId string, roleId string, token string, ) *Apikey`

NewApikey instantiates a new Apikey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApikeyWithDefaults

`func NewApikeyWithDefaults() *Apikey`

NewApikeyWithDefaults instantiates a new Apikey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *Apikey) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *Apikey) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *Apikey) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.


### GetUserId

`func (o *Apikey) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Apikey) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Apikey) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Apikey) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *Apikey) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Apikey) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Apikey) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetRoleId

`func (o *Apikey) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *Apikey) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *Apikey) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetToken

`func (o *Apikey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Apikey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Apikey) SetToken(v string)`

SetToken sets Token field to given value.


### GetName

`func (o *Apikey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Apikey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Apikey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Apikey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpirationDateUtc

`func (o *Apikey) GetExpirationDateUtc() string`

GetExpirationDateUtc returns the ExpirationDateUtc field if non-nil, zero value otherwise.

### GetExpirationDateUtcOk

`func (o *Apikey) GetExpirationDateUtcOk() (*string, bool)`

GetExpirationDateUtcOk returns a tuple with the ExpirationDateUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateUtc

`func (o *Apikey) SetExpirationDateUtc(v string)`

SetExpirationDateUtc sets ExpirationDateUtc field to given value.

### HasExpirationDateUtc

`func (o *Apikey) HasExpirationDateUtc() bool`

HasExpirationDateUtc returns a boolean if a field has been set.

### GetDateCreated

`func (o *Apikey) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Apikey) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Apikey) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Apikey) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Apikey) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Apikey) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Apikey) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Apikey) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *Apikey) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Apikey) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Apikey) SetActive(v string)`

SetActive sets Active field to given value.

### HasActive

`func (o *Apikey) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



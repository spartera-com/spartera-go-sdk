# ApiKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**ApiKeyId** | **int64** |  | 
**UserId** | Pointer to **string** | User who owns this API key | [optional] 
**CompanyId** | **string** | Company this API key belongs to | 
**RoleId** | **int64** | Role/permission level for this API key | 
**Name** | Pointer to **string** | Human-readable name for this API key | [optional] 
**ExpirationDateUtc** | Pointer to **time.Time** | When this API key expires (NULL &#x3D; never expires) | [optional] 

## Methods

### NewApiKeys

`func NewApiKeys(apiKeyId int64, companyId string, roleId int64, ) *ApiKeys`

NewApiKeys instantiates a new ApiKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeysWithDefaults

`func NewApiKeysWithDefaults() *ApiKeys`

NewApiKeysWithDefaults instantiates a new ApiKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *ApiKeys) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiKeys) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiKeys) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiKeys) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ApiKeys) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ApiKeys) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ApiKeys) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ApiKeys) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetApiKeyId

`func (o *ApiKeys) GetApiKeyId() int64`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *ApiKeys) GetApiKeyIdOk() (*int64, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *ApiKeys) SetApiKeyId(v int64)`

SetApiKeyId sets ApiKeyId field to given value.


### GetUserId

`func (o *ApiKeys) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiKeys) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiKeys) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiKeys) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *ApiKeys) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApiKeys) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApiKeys) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetRoleId

`func (o *ApiKeys) GetRoleId() int64`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ApiKeys) GetRoleIdOk() (*int64, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ApiKeys) SetRoleId(v int64)`

SetRoleId sets RoleId field to given value.


### GetName

`func (o *ApiKeys) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeys) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeys) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKeys) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpirationDateUtc

`func (o *ApiKeys) GetExpirationDateUtc() time.Time`

GetExpirationDateUtc returns the ExpirationDateUtc field if non-nil, zero value otherwise.

### GetExpirationDateUtcOk

`func (o *ApiKeys) GetExpirationDateUtcOk() (*time.Time, bool)`

GetExpirationDateUtcOk returns a tuple with the ExpirationDateUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateUtc

`func (o *ApiKeys) SetExpirationDateUtc(v time.Time)`

SetExpirationDateUtc sets ExpirationDateUtc field to given value.

### HasExpirationDateUtc

`func (o *ApiKeys) HasExpirationDateUtc() bool`

HasExpirationDateUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



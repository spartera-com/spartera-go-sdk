# ApiKeysInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | User who owns this API key | [optional] 
**CompanyId** | **string** | Company this API key belongs to | 
**RoleId** | **int64** | Role/permission level for this API key | 
**Name** | Pointer to **string** | Human-readable name for this API key | [optional] 
**ExpirationDateUtc** | Pointer to **time.Time** | When this API key expires (NULL &#x3D; never expires) | [optional] 

## Methods

### NewApiKeysInput

`func NewApiKeysInput(companyId string, roleId int64, ) *ApiKeysInput`

NewApiKeysInput instantiates a new ApiKeysInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeysInputWithDefaults

`func NewApiKeysInputWithDefaults() *ApiKeysInput`

NewApiKeysInputWithDefaults instantiates a new ApiKeysInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ApiKeysInput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiKeysInput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiKeysInput) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiKeysInput) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *ApiKeysInput) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApiKeysInput) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApiKeysInput) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetRoleId

`func (o *ApiKeysInput) GetRoleId() int64`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ApiKeysInput) GetRoleIdOk() (*int64, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ApiKeysInput) SetRoleId(v int64)`

SetRoleId sets RoleId field to given value.


### GetName

`func (o *ApiKeysInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeysInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeysInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKeysInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpirationDateUtc

`func (o *ApiKeysInput) GetExpirationDateUtc() time.Time`

GetExpirationDateUtc returns the ExpirationDateUtc field if non-nil, zero value otherwise.

### GetExpirationDateUtcOk

`func (o *ApiKeysInput) GetExpirationDateUtcOk() (*time.Time, bool)`

GetExpirationDateUtcOk returns a tuple with the ExpirationDateUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateUtc

`func (o *ApiKeysInput) SetExpirationDateUtc(v time.Time)`

SetExpirationDateUtc sets ExpirationDateUtc field to given value.

### HasExpirationDateUtc

`func (o *ApiKeysInput) HasExpirationDateUtc() bool`

HasExpirationDateUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



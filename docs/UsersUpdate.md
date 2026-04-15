# UsersUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** | References companies.company_id — A Spartera seller or buyer company account. See GET /companies for valid values. Required. | [optional] 
**RoleId** | Pointer to **int64** | User&#39;s role for RBAC - single source of truth | [optional] 
**FunctionId** | Pointer to **int64** | User&#39;s job function/title | [optional] 
**Status** | Pointer to **string** | Required. One of: ACTIVE, PENDING, INACTIVE, BANNED. | [optional] 
**EmailAddress** | Pointer to **string** | Optional. Must be unique. | [optional] 
**Timezone** | Pointer to **string** | Optional. | [optional] 
**MarketingOptOut** | Pointer to **bool** | Whether user has opted out of marketing communications. Default false &#x3D; opted in per ToS. | [optional] 

## Methods

### NewUsersUpdate

`func NewUsersUpdate() *UsersUpdate`

NewUsersUpdate instantiates a new UsersUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersUpdateWithDefaults

`func NewUsersUpdateWithDefaults() *UsersUpdate`

NewUsersUpdateWithDefaults instantiates a new UsersUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *UsersUpdate) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *UsersUpdate) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *UsersUpdate) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *UsersUpdate) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetRoleId

`func (o *UsersUpdate) GetRoleId() int64`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UsersUpdate) GetRoleIdOk() (*int64, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UsersUpdate) SetRoleId(v int64)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UsersUpdate) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetFunctionId

`func (o *UsersUpdate) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *UsersUpdate) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *UsersUpdate) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *UsersUpdate) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### GetStatus

`func (o *UsersUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsersUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsersUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UsersUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEmailAddress

`func (o *UsersUpdate) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UsersUpdate) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UsersUpdate) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *UsersUpdate) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetTimezone

`func (o *UsersUpdate) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UsersUpdate) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UsersUpdate) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UsersUpdate) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetMarketingOptOut

`func (o *UsersUpdate) GetMarketingOptOut() bool`

GetMarketingOptOut returns the MarketingOptOut field if non-nil, zero value otherwise.

### GetMarketingOptOutOk

`func (o *UsersUpdate) GetMarketingOptOutOk() (*bool, bool)`

GetMarketingOptOutOk returns a tuple with the MarketingOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptOut

`func (o *UsersUpdate) SetMarketingOptOut(v bool)`

SetMarketingOptOut sets MarketingOptOut field to given value.

### HasMarketingOptOut

`func (o *UsersUpdate) HasMarketingOptOut() bool`

HasMarketingOptOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



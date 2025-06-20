# MeGet200ResponseProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | User ID (primary key) | [optional] 
**Type** | Pointer to **string** | Profile type | [optional] 
**CompanyId** | Pointer to **string** | Company ID | [optional] 
**FunctionId** | Pointer to **int32** | User function/role ID | [optional] 
**Status** | Pointer to **string** | User status | [optional] 
**EmailAddress** | Pointer to **string** | User email address | [optional] 
**Timezone** | Pointer to **string** | User timezone | [optional] 
**DateCreated** | Pointer to **time.Time** | Account creation date | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last profile update | [optional] 
**Active** | Pointer to **bool** | Account active status | [optional] 

## Methods

### NewMeGet200ResponseProfile

`func NewMeGet200ResponseProfile() *MeGet200ResponseProfile`

NewMeGet200ResponseProfile instantiates a new MeGet200ResponseProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeGet200ResponseProfileWithDefaults

`func NewMeGet200ResponseProfileWithDefaults() *MeGet200ResponseProfile`

NewMeGet200ResponseProfileWithDefaults instantiates a new MeGet200ResponseProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MeGet200ResponseProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeGet200ResponseProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeGet200ResponseProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MeGet200ResponseProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MeGet200ResponseProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MeGet200ResponseProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MeGet200ResponseProfile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MeGet200ResponseProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCompanyId

`func (o *MeGet200ResponseProfile) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *MeGet200ResponseProfile) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *MeGet200ResponseProfile) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *MeGet200ResponseProfile) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetFunctionId

`func (o *MeGet200ResponseProfile) GetFunctionId() int32`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *MeGet200ResponseProfile) GetFunctionIdOk() (*int32, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *MeGet200ResponseProfile) SetFunctionId(v int32)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *MeGet200ResponseProfile) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### GetStatus

`func (o *MeGet200ResponseProfile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MeGet200ResponseProfile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MeGet200ResponseProfile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MeGet200ResponseProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEmailAddress

`func (o *MeGet200ResponseProfile) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MeGet200ResponseProfile) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *MeGet200ResponseProfile) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *MeGet200ResponseProfile) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetTimezone

`func (o *MeGet200ResponseProfile) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MeGet200ResponseProfile) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MeGet200ResponseProfile) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *MeGet200ResponseProfile) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetDateCreated

`func (o *MeGet200ResponseProfile) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *MeGet200ResponseProfile) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *MeGet200ResponseProfile) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *MeGet200ResponseProfile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *MeGet200ResponseProfile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *MeGet200ResponseProfile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *MeGet200ResponseProfile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *MeGet200ResponseProfile) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *MeGet200ResponseProfile) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MeGet200ResponseProfile) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MeGet200ResponseProfile) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MeGet200ResponseProfile) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



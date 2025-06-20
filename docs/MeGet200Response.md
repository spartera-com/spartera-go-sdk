# MeGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | Firebase user ID | 
**Email** | Pointer to **string** | User email address | [optional] 
**AuthMethod** | **string** | Authentication method used | 
**Platform** | Pointer to **string** | Platform origin | [optional] 
**OriginDomain** | Pointer to **string** | Request origin domain | [optional] 
**Profile** | [**MeGet200ResponseProfile**](MeGet200ResponseProfile.md) |  | 
**CompanyId** | **string** | Company ID from authentication claims | 
**RoleId** | **int32** | Role ID from authentication claims | 
**TokenMetadata** | Pointer to [**MeGet200ResponseTokenMetadata**](MeGet200ResponseTokenMetadata.md) |  | [optional] 
**ApiKeyInfo** | Pointer to [**MeGet200ResponseApiKeyInfo**](MeGet200ResponseApiKeyInfo.md) |  | [optional] 

## Methods

### NewMeGet200Response

`func NewMeGet200Response(userId string, authMethod string, profile MeGet200ResponseProfile, companyId string, roleId int32, ) *MeGet200Response`

NewMeGet200Response instantiates a new MeGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeGet200ResponseWithDefaults

`func NewMeGet200ResponseWithDefaults() *MeGet200Response`

NewMeGet200ResponseWithDefaults instantiates a new MeGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *MeGet200Response) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MeGet200Response) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MeGet200Response) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetEmail

`func (o *MeGet200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MeGet200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MeGet200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MeGet200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAuthMethod

`func (o *MeGet200Response) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *MeGet200Response) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *MeGet200Response) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetPlatform

`func (o *MeGet200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *MeGet200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *MeGet200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *MeGet200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetOriginDomain

`func (o *MeGet200Response) GetOriginDomain() string`

GetOriginDomain returns the OriginDomain field if non-nil, zero value otherwise.

### GetOriginDomainOk

`func (o *MeGet200Response) GetOriginDomainOk() (*string, bool)`

GetOriginDomainOk returns a tuple with the OriginDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDomain

`func (o *MeGet200Response) SetOriginDomain(v string)`

SetOriginDomain sets OriginDomain field to given value.

### HasOriginDomain

`func (o *MeGet200Response) HasOriginDomain() bool`

HasOriginDomain returns a boolean if a field has been set.

### GetProfile

`func (o *MeGet200Response) GetProfile() MeGet200ResponseProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MeGet200Response) GetProfileOk() (*MeGet200ResponseProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MeGet200Response) SetProfile(v MeGet200ResponseProfile)`

SetProfile sets Profile field to given value.


### GetCompanyId

`func (o *MeGet200Response) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *MeGet200Response) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *MeGet200Response) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetRoleId

`func (o *MeGet200Response) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *MeGet200Response) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *MeGet200Response) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.


### GetTokenMetadata

`func (o *MeGet200Response) GetTokenMetadata() MeGet200ResponseTokenMetadata`

GetTokenMetadata returns the TokenMetadata field if non-nil, zero value otherwise.

### GetTokenMetadataOk

`func (o *MeGet200Response) GetTokenMetadataOk() (*MeGet200ResponseTokenMetadata, bool)`

GetTokenMetadataOk returns a tuple with the TokenMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMetadata

`func (o *MeGet200Response) SetTokenMetadata(v MeGet200ResponseTokenMetadata)`

SetTokenMetadata sets TokenMetadata field to given value.

### HasTokenMetadata

`func (o *MeGet200Response) HasTokenMetadata() bool`

HasTokenMetadata returns a boolean if a field has been set.

### GetApiKeyInfo

`func (o *MeGet200Response) GetApiKeyInfo() MeGet200ResponseApiKeyInfo`

GetApiKeyInfo returns the ApiKeyInfo field if non-nil, zero value otherwise.

### GetApiKeyInfoOk

`func (o *MeGet200Response) GetApiKeyInfoOk() (*MeGet200ResponseApiKeyInfo, bool)`

GetApiKeyInfoOk returns a tuple with the ApiKeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyInfo

`func (o *MeGet200Response) SetApiKeyInfo(v MeGet200ResponseApiKeyInfo)`

SetApiKeyInfo sets ApiKeyInfo field to given value.

### HasApiKeyInfo

`func (o *MeGet200Response) HasApiKeyInfo() bool`

HasApiKeyInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



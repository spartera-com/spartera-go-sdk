# FavoritesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**CompanyId** | **string** |  | 
**Notes** | Pointer to **string** | Optional user notes about this favorite | [optional] 
**Category** | Pointer to **string** | Optional category for organizing favorites (e.g., &#39;Work&#39;, &#39;Research&#39;) | [optional] 
**Priority** | Pointer to **int64** | User-defined priority for sorting (higher &#x3D; more important) | [optional] 

## Methods

### NewFavoritesInput

`func NewFavoritesInput(assetId string, companyId string, ) *FavoritesInput`

NewFavoritesInput instantiates a new FavoritesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoritesInputWithDefaults

`func NewFavoritesInputWithDefaults() *FavoritesInput`

NewFavoritesInputWithDefaults instantiates a new FavoritesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *FavoritesInput) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *FavoritesInput) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *FavoritesInput) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetUserId

`func (o *FavoritesInput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FavoritesInput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FavoritesInput) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FavoritesInput) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *FavoritesInput) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *FavoritesInput) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *FavoritesInput) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetNotes

`func (o *FavoritesInput) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *FavoritesInput) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *FavoritesInput) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *FavoritesInput) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCategory

`func (o *FavoritesInput) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FavoritesInput) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FavoritesInput) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FavoritesInput) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPriority

`func (o *FavoritesInput) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *FavoritesInput) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *FavoritesInput) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *FavoritesInput) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



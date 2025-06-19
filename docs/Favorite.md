# Favorite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FavoriteId** | Pointer to **string** |  | [optional] [readonly] 
**AssetId** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**CompanyId** | **string** |  | 
**Notes** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Priority** | **string** |  | 
**DateCreated** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **string** |  | [optional] [readonly] 
**Active** | Pointer to **string** |  | [optional] 

## Methods

### NewFavorite

`func NewFavorite(assetId string, companyId string, priority string, ) *Favorite`

NewFavorite instantiates a new Favorite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoriteWithDefaults

`func NewFavoriteWithDefaults() *Favorite`

NewFavoriteWithDefaults instantiates a new Favorite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFavoriteId

`func (o *Favorite) GetFavoriteId() string`

GetFavoriteId returns the FavoriteId field if non-nil, zero value otherwise.

### GetFavoriteIdOk

`func (o *Favorite) GetFavoriteIdOk() (*string, bool)`

GetFavoriteIdOk returns a tuple with the FavoriteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteId

`func (o *Favorite) SetFavoriteId(v string)`

SetFavoriteId sets FavoriteId field to given value.

### HasFavoriteId

`func (o *Favorite) HasFavoriteId() bool`

HasFavoriteId returns a boolean if a field has been set.

### GetAssetId

`func (o *Favorite) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Favorite) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Favorite) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetUserId

`func (o *Favorite) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Favorite) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Favorite) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Favorite) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *Favorite) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Favorite) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Favorite) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetNotes

`func (o *Favorite) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Favorite) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Favorite) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Favorite) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCategory

`func (o *Favorite) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Favorite) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Favorite) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Favorite) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPriority

`func (o *Favorite) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Favorite) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Favorite) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetDateCreated

`func (o *Favorite) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Favorite) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Favorite) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Favorite) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Favorite) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Favorite) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Favorite) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Favorite) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *Favorite) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Favorite) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Favorite) SetActive(v string)`

SetActive sets Active field to given value.

### HasActive

`func (o *Favorite) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



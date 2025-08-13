# Favorites

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**FavoriteId** | Pointer to **int64** |  | [optional] 
**AssetId** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**CompanyId** | **string** |  | 
**Notes** | Pointer to **string** | Optional user notes about this favorite | [optional] 
**Category** | Pointer to **string** | Optional category for organizing favorites (e.g., &#39;Work&#39;, &#39;Research&#39;) | [optional] 
**Priority** | **int64** | User-defined priority for sorting (higher &#x3D; more important) | 

## Methods

### NewFavorites

`func NewFavorites(assetId string, companyId string, priority int64, ) *Favorites`

NewFavorites instantiates a new Favorites object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoritesWithDefaults

`func NewFavoritesWithDefaults() *Favorites`

NewFavoritesWithDefaults instantiates a new Favorites object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *Favorites) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Favorites) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Favorites) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Favorites) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Favorites) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Favorites) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Favorites) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Favorites) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetFavoriteId

`func (o *Favorites) GetFavoriteId() int64`

GetFavoriteId returns the FavoriteId field if non-nil, zero value otherwise.

### GetFavoriteIdOk

`func (o *Favorites) GetFavoriteIdOk() (*int64, bool)`

GetFavoriteIdOk returns a tuple with the FavoriteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteId

`func (o *Favorites) SetFavoriteId(v int64)`

SetFavoriteId sets FavoriteId field to given value.

### HasFavoriteId

`func (o *Favorites) HasFavoriteId() bool`

HasFavoriteId returns a boolean if a field has been set.

### GetAssetId

`func (o *Favorites) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Favorites) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Favorites) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetUserId

`func (o *Favorites) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Favorites) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Favorites) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Favorites) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *Favorites) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Favorites) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Favorites) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetNotes

`func (o *Favorites) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Favorites) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Favorites) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Favorites) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCategory

`func (o *Favorites) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Favorites) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Favorites) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Favorites) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPriority

`func (o *Favorites) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Favorites) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Favorites) SetPriority(v int64)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



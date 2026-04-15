# AssetUsecases

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** | Optional. | [optional] 
**LastUpdated** | Pointer to **time.Time** | Optional. | [optional] 
**AucId** | Pointer to **int64** | Unique identifier. | [optional] 
**AucName** | **string** | Required. Must be unique. | 
**Slug** | Pointer to **string** | URL-friendly slug derived from auc_name (e.g. &#39;competitive-benchmarking&#39;) | [optional] 
**AucDescription** | Pointer to **string** | Optional. | [optional] 

## Methods

### NewAssetUsecases

`func NewAssetUsecases(aucName string, ) *AssetUsecases`

NewAssetUsecases instantiates a new AssetUsecases object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetUsecasesWithDefaults

`func NewAssetUsecasesWithDefaults() *AssetUsecases`

NewAssetUsecasesWithDefaults instantiates a new AssetUsecases object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *AssetUsecases) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AssetUsecases) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AssetUsecases) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AssetUsecases) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AssetUsecases) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AssetUsecases) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AssetUsecases) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AssetUsecases) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAucId

`func (o *AssetUsecases) GetAucId() int64`

GetAucId returns the AucId field if non-nil, zero value otherwise.

### GetAucIdOk

`func (o *AssetUsecases) GetAucIdOk() (*int64, bool)`

GetAucIdOk returns a tuple with the AucId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAucId

`func (o *AssetUsecases) SetAucId(v int64)`

SetAucId sets AucId field to given value.

### HasAucId

`func (o *AssetUsecases) HasAucId() bool`

HasAucId returns a boolean if a field has been set.

### GetAucName

`func (o *AssetUsecases) GetAucName() string`

GetAucName returns the AucName field if non-nil, zero value otherwise.

### GetAucNameOk

`func (o *AssetUsecases) GetAucNameOk() (*string, bool)`

GetAucNameOk returns a tuple with the AucName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAucName

`func (o *AssetUsecases) SetAucName(v string)`

SetAucName sets AucName field to given value.


### GetSlug

`func (o *AssetUsecases) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AssetUsecases) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AssetUsecases) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AssetUsecases) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetAucDescription

`func (o *AssetUsecases) GetAucDescription() string`

GetAucDescription returns the AucDescription field if non-nil, zero value otherwise.

### GetAucDescriptionOk

`func (o *AssetUsecases) GetAucDescriptionOk() (*string, bool)`

GetAucDescriptionOk returns a tuple with the AucDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAucDescription

`func (o *AssetUsecases) SetAucDescription(v string)`

SetAucDescription sets AucDescription field to given value.

### HasAucDescription

`func (o *AssetUsecases) HasAucDescription() bool`

HasAucDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



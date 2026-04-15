# Industries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** | Optional. | [optional] 
**LastUpdated** | Pointer to **time.Time** | Optional. | [optional] 
**IndustryId** | Pointer to **int64** | Auto-generated unique identifier. | [optional] 
**IndustryName** | **string** | Required. Must be unique. | 
**ShortName** | **string** | Required. | 
**Slug** | Pointer to **string** | Optional. | [optional] 
**NaiscCode** | Pointer to **int64** | Optional. | [optional] 
**Description** | Pointer to **string** | Optional. | [optional] 

## Methods

### NewIndustries

`func NewIndustries(industryName string, shortName string, ) *Industries`

NewIndustries instantiates a new Industries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndustriesWithDefaults

`func NewIndustriesWithDefaults() *Industries`

NewIndustriesWithDefaults instantiates a new Industries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *Industries) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Industries) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Industries) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Industries) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Industries) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Industries) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Industries) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Industries) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetIndustryId

`func (o *Industries) GetIndustryId() int64`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *Industries) GetIndustryIdOk() (*int64, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *Industries) SetIndustryId(v int64)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *Industries) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetIndustryName

`func (o *Industries) GetIndustryName() string`

GetIndustryName returns the IndustryName field if non-nil, zero value otherwise.

### GetIndustryNameOk

`func (o *Industries) GetIndustryNameOk() (*string, bool)`

GetIndustryNameOk returns a tuple with the IndustryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryName

`func (o *Industries) SetIndustryName(v string)`

SetIndustryName sets IndustryName field to given value.


### GetShortName

`func (o *Industries) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *Industries) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *Industries) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetSlug

`func (o *Industries) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Industries) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Industries) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Industries) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetNaiscCode

`func (o *Industries) GetNaiscCode() int64`

GetNaiscCode returns the NaiscCode field if non-nil, zero value otherwise.

### GetNaiscCodeOk

`func (o *Industries) GetNaiscCodeOk() (*int64, bool)`

GetNaiscCodeOk returns a tuple with the NaiscCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaiscCode

`func (o *Industries) SetNaiscCode(v int64)`

SetNaiscCode sets NaiscCode field to given value.

### HasNaiscCode

`func (o *Industries) HasNaiscCode() bool`

HasNaiscCode returns a boolean if a field has been set.

### GetDescription

`func (o *Industries) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Industries) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Industries) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Industries) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



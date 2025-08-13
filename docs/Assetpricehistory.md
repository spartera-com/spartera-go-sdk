# AssetPriceHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**AphId** | Pointer to **string** |  | [optional] 
**AssetId** | **string** |  | 
**PriceUsd** | Pointer to **float64** |  | [optional] 
**DateEnded** | Pointer to **time.Time** | When did the price end (Datetime) | [optional] 

## Methods

### NewAssetPriceHistory

`func NewAssetPriceHistory(assetId string, ) *AssetPriceHistory`

NewAssetPriceHistory instantiates a new AssetPriceHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetPriceHistoryWithDefaults

`func NewAssetPriceHistoryWithDefaults() *AssetPriceHistory`

NewAssetPriceHistoryWithDefaults instantiates a new AssetPriceHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *AssetPriceHistory) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AssetPriceHistory) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AssetPriceHistory) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AssetPriceHistory) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AssetPriceHistory) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AssetPriceHistory) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AssetPriceHistory) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AssetPriceHistory) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAphId

`func (o *AssetPriceHistory) GetAphId() string`

GetAphId returns the AphId field if non-nil, zero value otherwise.

### GetAphIdOk

`func (o *AssetPriceHistory) GetAphIdOk() (*string, bool)`

GetAphIdOk returns a tuple with the AphId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAphId

`func (o *AssetPriceHistory) SetAphId(v string)`

SetAphId sets AphId field to given value.

### HasAphId

`func (o *AssetPriceHistory) HasAphId() bool`

HasAphId returns a boolean if a field has been set.

### GetAssetId

`func (o *AssetPriceHistory) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetPriceHistory) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetPriceHistory) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetPriceUsd

`func (o *AssetPriceHistory) GetPriceUsd() float64`

GetPriceUsd returns the PriceUsd field if non-nil, zero value otherwise.

### GetPriceUsdOk

`func (o *AssetPriceHistory) GetPriceUsdOk() (*float64, bool)`

GetPriceUsdOk returns a tuple with the PriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUsd

`func (o *AssetPriceHistory) SetPriceUsd(v float64)`

SetPriceUsd sets PriceUsd field to given value.

### HasPriceUsd

`func (o *AssetPriceHistory) HasPriceUsd() bool`

HasPriceUsd returns a boolean if a field has been set.

### GetDateEnded

`func (o *AssetPriceHistory) GetDateEnded() time.Time`

GetDateEnded returns the DateEnded field if non-nil, zero value otherwise.

### GetDateEndedOk

`func (o *AssetPriceHistory) GetDateEndedOk() (*time.Time, bool)`

GetDateEndedOk returns a tuple with the DateEnded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnded

`func (o *AssetPriceHistory) SetDateEnded(v time.Time)`

SetDateEnded sets DateEnded field to given value.

### HasDateEnded

`func (o *AssetPriceHistory) HasDateEnded() bool`

HasDateEnded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



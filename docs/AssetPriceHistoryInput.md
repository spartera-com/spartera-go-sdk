# AssetPriceHistoryInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** |  | 
**PriceUsd** | Pointer to **float64** |  | [optional] 
**DateEnded** | Pointer to **time.Time** | When did the price end (Datetime) | [optional] 

## Methods

### NewAssetPriceHistoryInput

`func NewAssetPriceHistoryInput(assetId string, ) *AssetPriceHistoryInput`

NewAssetPriceHistoryInput instantiates a new AssetPriceHistoryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetPriceHistoryInputWithDefaults

`func NewAssetPriceHistoryInputWithDefaults() *AssetPriceHistoryInput`

NewAssetPriceHistoryInputWithDefaults instantiates a new AssetPriceHistoryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *AssetPriceHistoryInput) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetPriceHistoryInput) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetPriceHistoryInput) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetPriceUsd

`func (o *AssetPriceHistoryInput) GetPriceUsd() float64`

GetPriceUsd returns the PriceUsd field if non-nil, zero value otherwise.

### GetPriceUsdOk

`func (o *AssetPriceHistoryInput) GetPriceUsdOk() (*float64, bool)`

GetPriceUsdOk returns a tuple with the PriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUsd

`func (o *AssetPriceHistoryInput) SetPriceUsd(v float64)`

SetPriceUsd sets PriceUsd field to given value.

### HasPriceUsd

`func (o *AssetPriceHistoryInput) HasPriceUsd() bool`

HasPriceUsd returns a boolean if a field has been set.

### GetDateEnded

`func (o *AssetPriceHistoryInput) GetDateEnded() time.Time`

GetDateEnded returns the DateEnded field if non-nil, zero value otherwise.

### GetDateEndedOk

`func (o *AssetPriceHistoryInput) GetDateEndedOk() (*time.Time, bool)`

GetDateEndedOk returns a tuple with the DateEnded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnded

`func (o *AssetPriceHistoryInput) SetDateEnded(v time.Time)`

SetDateEnded sets DateEnded field to given value.

### HasDateEnded

`func (o *AssetPriceHistoryInput) HasDateEnded() bool`

HasDateEnded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



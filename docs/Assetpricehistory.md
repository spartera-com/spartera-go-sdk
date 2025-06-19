# Assetpricehistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AphId** | Pointer to **string** |  | [optional] 
**AssetId** | **string** |  | 
**PriceUsd** | Pointer to **float64** |  | [optional] 
**PriceCredits** | Pointer to **string** |  | [optional] 
**DiscountPercentage** | Pointer to **float32** |  | [optional] 
**SaleStartDate** | Pointer to **string** |  | [optional] 
**SaleEndDate** | Pointer to **string** |  | [optional] 
**DateEnded** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **string** |  | [optional] [readonly] 
**Active** | **string** |  | 

## Methods

### NewAssetpricehistory

`func NewAssetpricehistory(assetId string, active string, ) *Assetpricehistory`

NewAssetpricehistory instantiates a new Assetpricehistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetpricehistoryWithDefaults

`func NewAssetpricehistoryWithDefaults() *Assetpricehistory`

NewAssetpricehistoryWithDefaults instantiates a new Assetpricehistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAphId

`func (o *Assetpricehistory) GetAphId() string`

GetAphId returns the AphId field if non-nil, zero value otherwise.

### GetAphIdOk

`func (o *Assetpricehistory) GetAphIdOk() (*string, bool)`

GetAphIdOk returns a tuple with the AphId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAphId

`func (o *Assetpricehistory) SetAphId(v string)`

SetAphId sets AphId field to given value.

### HasAphId

`func (o *Assetpricehistory) HasAphId() bool`

HasAphId returns a boolean if a field has been set.

### GetAssetId

`func (o *Assetpricehistory) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Assetpricehistory) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Assetpricehistory) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetPriceUsd

`func (o *Assetpricehistory) GetPriceUsd() float64`

GetPriceUsd returns the PriceUsd field if non-nil, zero value otherwise.

### GetPriceUsdOk

`func (o *Assetpricehistory) GetPriceUsdOk() (*float64, bool)`

GetPriceUsdOk returns a tuple with the PriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUsd

`func (o *Assetpricehistory) SetPriceUsd(v float64)`

SetPriceUsd sets PriceUsd field to given value.

### HasPriceUsd

`func (o *Assetpricehistory) HasPriceUsd() bool`

HasPriceUsd returns a boolean if a field has been set.

### GetPriceCredits

`func (o *Assetpricehistory) GetPriceCredits() string`

GetPriceCredits returns the PriceCredits field if non-nil, zero value otherwise.

### GetPriceCreditsOk

`func (o *Assetpricehistory) GetPriceCreditsOk() (*string, bool)`

GetPriceCreditsOk returns a tuple with the PriceCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCredits

`func (o *Assetpricehistory) SetPriceCredits(v string)`

SetPriceCredits sets PriceCredits field to given value.

### HasPriceCredits

`func (o *Assetpricehistory) HasPriceCredits() bool`

HasPriceCredits returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *Assetpricehistory) GetDiscountPercentage() float32`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *Assetpricehistory) GetDiscountPercentageOk() (*float32, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *Assetpricehistory) SetDiscountPercentage(v float32)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *Assetpricehistory) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetSaleStartDate

`func (o *Assetpricehistory) GetSaleStartDate() string`

GetSaleStartDate returns the SaleStartDate field if non-nil, zero value otherwise.

### GetSaleStartDateOk

`func (o *Assetpricehistory) GetSaleStartDateOk() (*string, bool)`

GetSaleStartDateOk returns a tuple with the SaleStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleStartDate

`func (o *Assetpricehistory) SetSaleStartDate(v string)`

SetSaleStartDate sets SaleStartDate field to given value.

### HasSaleStartDate

`func (o *Assetpricehistory) HasSaleStartDate() bool`

HasSaleStartDate returns a boolean if a field has been set.

### GetSaleEndDate

`func (o *Assetpricehistory) GetSaleEndDate() string`

GetSaleEndDate returns the SaleEndDate field if non-nil, zero value otherwise.

### GetSaleEndDateOk

`func (o *Assetpricehistory) GetSaleEndDateOk() (*string, bool)`

GetSaleEndDateOk returns a tuple with the SaleEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleEndDate

`func (o *Assetpricehistory) SetSaleEndDate(v string)`

SetSaleEndDate sets SaleEndDate field to given value.

### HasSaleEndDate

`func (o *Assetpricehistory) HasSaleEndDate() bool`

HasSaleEndDate returns a boolean if a field has been set.

### GetDateEnded

`func (o *Assetpricehistory) GetDateEnded() string`

GetDateEnded returns the DateEnded field if non-nil, zero value otherwise.

### GetDateEndedOk

`func (o *Assetpricehistory) GetDateEndedOk() (*string, bool)`

GetDateEndedOk returns a tuple with the DateEnded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnded

`func (o *Assetpricehistory) SetDateEnded(v string)`

SetDateEnded sets DateEnded field to given value.

### HasDateEnded

`func (o *Assetpricehistory) HasDateEnded() bool`

HasDateEnded returns a boolean if a field has been set.

### GetDateCreated

`func (o *Assetpricehistory) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Assetpricehistory) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Assetpricehistory) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Assetpricehistory) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Assetpricehistory) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Assetpricehistory) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Assetpricehistory) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Assetpricehistory) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *Assetpricehistory) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Assetpricehistory) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Assetpricehistory) SetActive(v string)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



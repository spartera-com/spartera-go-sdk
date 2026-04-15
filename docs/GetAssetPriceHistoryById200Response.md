# GetAssetPriceHistoryById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**AssetPriceHistory**](AssetPriceHistory.md) |  | 

## Methods

### NewGetAssetPriceHistoryById200Response

`func NewGetAssetPriceHistoryById200Response(message string, data AssetPriceHistory, ) *GetAssetPriceHistoryById200Response`

NewGetAssetPriceHistoryById200Response instantiates a new GetAssetPriceHistoryById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetPriceHistoryById200ResponseWithDefaults

`func NewGetAssetPriceHistoryById200ResponseWithDefaults() *GetAssetPriceHistoryById200Response`

NewGetAssetPriceHistoryById200ResponseWithDefaults instantiates a new GetAssetPriceHistoryById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetAssetPriceHistoryById200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAssetPriceHistoryById200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAssetPriceHistoryById200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetAssetPriceHistoryById200Response) GetData() AssetPriceHistory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAssetPriceHistoryById200Response) GetDataOk() (*AssetPriceHistory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAssetPriceHistoryById200Response) SetData(v AssetPriceHistory)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



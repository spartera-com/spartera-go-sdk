# DeleteAssetPriceHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**DeleteAssetPriceHistory200ResponseData**](DeleteAssetPriceHistory200ResponseData.md) |  | 

## Methods

### NewDeleteAssetPriceHistory200Response

`func NewDeleteAssetPriceHistory200Response(message string, data DeleteAssetPriceHistory200ResponseData, ) *DeleteAssetPriceHistory200Response`

NewDeleteAssetPriceHistory200Response instantiates a new DeleteAssetPriceHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAssetPriceHistory200ResponseWithDefaults

`func NewDeleteAssetPriceHistory200ResponseWithDefaults() *DeleteAssetPriceHistory200Response`

NewDeleteAssetPriceHistory200ResponseWithDefaults instantiates a new DeleteAssetPriceHistory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DeleteAssetPriceHistory200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteAssetPriceHistory200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteAssetPriceHistory200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *DeleteAssetPriceHistory200Response) GetData() DeleteAssetPriceHistory200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteAssetPriceHistory200Response) GetDataOk() (*DeleteAssetPriceHistory200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteAssetPriceHistory200Response) SetData(v DeleteAssetPriceHistory200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



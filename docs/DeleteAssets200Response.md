# DeleteAssets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**DeleteAssets200ResponseData**](DeleteAssets200ResponseData.md) |  | 

## Methods

### NewDeleteAssets200Response

`func NewDeleteAssets200Response(message string, data DeleteAssets200ResponseData, ) *DeleteAssets200Response`

NewDeleteAssets200Response instantiates a new DeleteAssets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAssets200ResponseWithDefaults

`func NewDeleteAssets200ResponseWithDefaults() *DeleteAssets200Response`

NewDeleteAssets200ResponseWithDefaults instantiates a new DeleteAssets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DeleteAssets200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteAssets200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteAssets200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *DeleteAssets200Response) GetData() DeleteAssets200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteAssets200Response) GetDataOk() (*DeleteAssets200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteAssets200Response) SetData(v DeleteAssets200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



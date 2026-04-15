# DeleteEndpoints200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**DeleteEndpoints200ResponseData**](DeleteEndpoints200ResponseData.md) |  | 

## Methods

### NewDeleteEndpoints200Response

`func NewDeleteEndpoints200Response(message string, data DeleteEndpoints200ResponseData, ) *DeleteEndpoints200Response`

NewDeleteEndpoints200Response instantiates a new DeleteEndpoints200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteEndpoints200ResponseWithDefaults

`func NewDeleteEndpoints200ResponseWithDefaults() *DeleteEndpoints200Response`

NewDeleteEndpoints200ResponseWithDefaults instantiates a new DeleteEndpoints200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DeleteEndpoints200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteEndpoints200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteEndpoints200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *DeleteEndpoints200Response) GetData() DeleteEndpoints200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteEndpoints200Response) GetDataOk() (*DeleteEndpoints200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteEndpoints200Response) SetData(v DeleteEndpoints200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



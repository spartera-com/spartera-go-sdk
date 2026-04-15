# UpdateUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**UpdateUsers200ResponseData**](UpdateUsers200ResponseData.md) |  | 

## Methods

### NewUpdateUsers200Response

`func NewUpdateUsers200Response(message string, data UpdateUsers200ResponseData, ) *UpdateUsers200Response`

NewUpdateUsers200Response instantiates a new UpdateUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUsers200ResponseWithDefaults

`func NewUpdateUsers200ResponseWithDefaults() *UpdateUsers200Response`

NewUpdateUsers200ResponseWithDefaults instantiates a new UpdateUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateUsers200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateUsers200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateUsers200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *UpdateUsers200Response) GetData() UpdateUsers200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateUsers200Response) GetDataOk() (*UpdateUsers200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateUsers200Response) SetData(v UpdateUsers200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



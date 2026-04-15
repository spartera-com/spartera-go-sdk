# UpdateConnections200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**UpdateConnections200ResponseData**](UpdateConnections200ResponseData.md) |  | 

## Methods

### NewUpdateConnections200Response

`func NewUpdateConnections200Response(message string, data UpdateConnections200ResponseData, ) *UpdateConnections200Response`

NewUpdateConnections200Response instantiates a new UpdateConnections200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConnections200ResponseWithDefaults

`func NewUpdateConnections200ResponseWithDefaults() *UpdateConnections200Response`

NewUpdateConnections200ResponseWithDefaults instantiates a new UpdateConnections200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateConnections200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateConnections200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateConnections200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *UpdateConnections200Response) GetData() UpdateConnections200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateConnections200Response) GetDataOk() (*UpdateConnections200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateConnections200Response) SetData(v UpdateConnections200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateAlerts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**UpdateAlerts200ResponseData**](UpdateAlerts200ResponseData.md) |  | 

## Methods

### NewUpdateAlerts200Response

`func NewUpdateAlerts200Response(message string, data UpdateAlerts200ResponseData, ) *UpdateAlerts200Response`

NewUpdateAlerts200Response instantiates a new UpdateAlerts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlerts200ResponseWithDefaults

`func NewUpdateAlerts200ResponseWithDefaults() *UpdateAlerts200Response`

NewUpdateAlerts200ResponseWithDefaults instantiates a new UpdateAlerts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateAlerts200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateAlerts200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateAlerts200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *UpdateAlerts200Response) GetData() UpdateAlerts200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateAlerts200Response) GetDataOk() (*UpdateAlerts200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateAlerts200Response) SetData(v UpdateAlerts200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



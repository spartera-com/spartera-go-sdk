# UpdateApiKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**UpdateApiKeys200ResponseData**](UpdateApiKeys200ResponseData.md) |  | 

## Methods

### NewUpdateApiKeys200Response

`func NewUpdateApiKeys200Response(message string, data UpdateApiKeys200ResponseData, ) *UpdateApiKeys200Response`

NewUpdateApiKeys200Response instantiates a new UpdateApiKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApiKeys200ResponseWithDefaults

`func NewUpdateApiKeys200ResponseWithDefaults() *UpdateApiKeys200Response`

NewUpdateApiKeys200ResponseWithDefaults instantiates a new UpdateApiKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateApiKeys200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateApiKeys200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateApiKeys200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *UpdateApiKeys200Response) GetData() UpdateApiKeys200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateApiKeys200Response) GetDataOk() (*UpdateApiKeys200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateApiKeys200Response) SetData(v UpdateApiKeys200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



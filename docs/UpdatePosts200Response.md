# UpdatePosts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**UpdatePosts200ResponseData**](UpdatePosts200ResponseData.md) |  | 

## Methods

### NewUpdatePosts200Response

`func NewUpdatePosts200Response(message string, data UpdatePosts200ResponseData, ) *UpdatePosts200Response`

NewUpdatePosts200Response instantiates a new UpdatePosts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePosts200ResponseWithDefaults

`func NewUpdatePosts200ResponseWithDefaults() *UpdatePosts200Response`

NewUpdatePosts200ResponseWithDefaults instantiates a new UpdatePosts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdatePosts200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdatePosts200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdatePosts200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *UpdatePosts200Response) GetData() UpdatePosts200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdatePosts200Response) GetDataOk() (*UpdatePosts200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdatePosts200Response) SetData(v UpdatePosts200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



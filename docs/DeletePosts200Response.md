# DeletePosts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**DeletePosts200ResponseData**](DeletePosts200ResponseData.md) |  | 

## Methods

### NewDeletePosts200Response

`func NewDeletePosts200Response(message string, data DeletePosts200ResponseData, ) *DeletePosts200Response`

NewDeletePosts200Response instantiates a new DeletePosts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePosts200ResponseWithDefaults

`func NewDeletePosts200ResponseWithDefaults() *DeletePosts200Response`

NewDeletePosts200ResponseWithDefaults instantiates a new DeletePosts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DeletePosts200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeletePosts200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeletePosts200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *DeletePosts200Response) GetData() DeletePosts200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeletePosts200Response) GetDataOk() (*DeletePosts200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeletePosts200Response) SetData(v DeletePosts200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



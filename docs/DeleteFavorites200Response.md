# DeleteFavorites200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**DeleteFavorites200ResponseData**](DeleteFavorites200ResponseData.md) |  | 

## Methods

### NewDeleteFavorites200Response

`func NewDeleteFavorites200Response(message string, data DeleteFavorites200ResponseData, ) *DeleteFavorites200Response`

NewDeleteFavorites200Response instantiates a new DeleteFavorites200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteFavorites200ResponseWithDefaults

`func NewDeleteFavorites200ResponseWithDefaults() *DeleteFavorites200Response`

NewDeleteFavorites200ResponseWithDefaults instantiates a new DeleteFavorites200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DeleteFavorites200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteFavorites200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteFavorites200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *DeleteFavorites200Response) GetData() DeleteFavorites200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteFavorites200Response) GetDataOk() (*DeleteFavorites200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteFavorites200Response) SetData(v DeleteFavorites200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



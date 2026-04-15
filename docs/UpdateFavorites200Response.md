# UpdateFavorites200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**UpdateFavorites200ResponseData**](UpdateFavorites200ResponseData.md) |  | 

## Methods

### NewUpdateFavorites200Response

`func NewUpdateFavorites200Response(message string, data UpdateFavorites200ResponseData, ) *UpdateFavorites200Response`

NewUpdateFavorites200Response instantiates a new UpdateFavorites200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFavorites200ResponseWithDefaults

`func NewUpdateFavorites200ResponseWithDefaults() *UpdateFavorites200Response`

NewUpdateFavorites200ResponseWithDefaults instantiates a new UpdateFavorites200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateFavorites200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateFavorites200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateFavorites200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *UpdateFavorites200Response) GetData() UpdateFavorites200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateFavorites200Response) GetDataOk() (*UpdateFavorites200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateFavorites200Response) SetData(v UpdateFavorites200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



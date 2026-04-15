# CreateFavorites200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**CreateFavorites200ResponseData**](CreateFavorites200ResponseData.md) |  | 

## Methods

### NewCreateFavorites200Response

`func NewCreateFavorites200Response(message string, data CreateFavorites200ResponseData, ) *CreateFavorites200Response`

NewCreateFavorites200Response instantiates a new CreateFavorites200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFavorites200ResponseWithDefaults

`func NewCreateFavorites200ResponseWithDefaults() *CreateFavorites200Response`

NewCreateFavorites200ResponseWithDefaults instantiates a new CreateFavorites200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateFavorites200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateFavorites200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateFavorites200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *CreateFavorites200Response) GetData() CreateFavorites200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateFavorites200Response) GetDataOk() (*CreateFavorites200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateFavorites200Response) SetData(v CreateFavorites200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



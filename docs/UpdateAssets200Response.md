# UpdateAssets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**UpdateAssets200ResponseData**](UpdateAssets200ResponseData.md) |  | 

## Methods

### NewUpdateAssets200Response

`func NewUpdateAssets200Response(message string, data UpdateAssets200ResponseData, ) *UpdateAssets200Response`

NewUpdateAssets200Response instantiates a new UpdateAssets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAssets200ResponseWithDefaults

`func NewUpdateAssets200ResponseWithDefaults() *UpdateAssets200Response`

NewUpdateAssets200ResponseWithDefaults instantiates a new UpdateAssets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateAssets200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateAssets200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateAssets200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *UpdateAssets200Response) GetData() UpdateAssets200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateAssets200Response) GetDataOk() (*UpdateAssets200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateAssets200Response) SetData(v UpdateAssets200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



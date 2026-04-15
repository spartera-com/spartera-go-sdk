# CreateUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**CreateUsers200ResponseData**](CreateUsers200ResponseData.md) |  | 

## Methods

### NewCreateUsers200Response

`func NewCreateUsers200Response(message string, data CreateUsers200ResponseData, ) *CreateUsers200Response`

NewCreateUsers200Response instantiates a new CreateUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUsers200ResponseWithDefaults

`func NewCreateUsers200ResponseWithDefaults() *CreateUsers200Response`

NewCreateUsers200ResponseWithDefaults instantiates a new CreateUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateUsers200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateUsers200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateUsers200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *CreateUsers200Response) GetData() CreateUsers200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateUsers200Response) GetDataOk() (*CreateUsers200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateUsers200Response) SetData(v CreateUsers200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



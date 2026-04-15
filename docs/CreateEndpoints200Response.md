# CreateEndpoints200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**CreateEndpoints200ResponseData**](CreateEndpoints200ResponseData.md) |  | 

## Methods

### NewCreateEndpoints200Response

`func NewCreateEndpoints200Response(message string, data CreateEndpoints200ResponseData, ) *CreateEndpoints200Response`

NewCreateEndpoints200Response instantiates a new CreateEndpoints200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEndpoints200ResponseWithDefaults

`func NewCreateEndpoints200ResponseWithDefaults() *CreateEndpoints200Response`

NewCreateEndpoints200ResponseWithDefaults instantiates a new CreateEndpoints200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateEndpoints200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateEndpoints200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateEndpoints200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *CreateEndpoints200Response) GetData() CreateEndpoints200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateEndpoints200Response) GetDataOk() (*CreateEndpoints200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateEndpoints200Response) SetData(v CreateEndpoints200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



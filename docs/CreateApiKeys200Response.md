# CreateApiKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**CreateApiKeys200ResponseData**](CreateApiKeys200ResponseData.md) |  | 

## Methods

### NewCreateApiKeys200Response

`func NewCreateApiKeys200Response(message string, data CreateApiKeys200ResponseData, ) *CreateApiKeys200Response`

NewCreateApiKeys200Response instantiates a new CreateApiKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKeys200ResponseWithDefaults

`func NewCreateApiKeys200ResponseWithDefaults() *CreateApiKeys200Response`

NewCreateApiKeys200ResponseWithDefaults instantiates a new CreateApiKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateApiKeys200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateApiKeys200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateApiKeys200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *CreateApiKeys200Response) GetData() CreateApiKeys200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateApiKeys200Response) GetDataOk() (*CreateApiKeys200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateApiKeys200Response) SetData(v CreateApiKeys200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



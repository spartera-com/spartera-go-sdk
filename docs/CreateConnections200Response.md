# CreateConnections200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**CreateConnections200ResponseData**](CreateConnections200ResponseData.md) |  | 

## Methods

### NewCreateConnections200Response

`func NewCreateConnections200Response(message string, data CreateConnections200ResponseData, ) *CreateConnections200Response`

NewCreateConnections200Response instantiates a new CreateConnections200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnections200ResponseWithDefaults

`func NewCreateConnections200ResponseWithDefaults() *CreateConnections200Response`

NewCreateConnections200ResponseWithDefaults instantiates a new CreateConnections200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateConnections200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateConnections200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateConnections200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *CreateConnections200Response) GetData() CreateConnections200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateConnections200Response) GetDataOk() (*CreateConnections200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateConnections200Response) SetData(v CreateConnections200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



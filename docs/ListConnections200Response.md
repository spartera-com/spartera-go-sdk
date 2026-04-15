# ListConnections200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**[]Connections**](Connections.md) |  | 

## Methods

### NewListConnections200Response

`func NewListConnections200Response(message string, data []Connections, ) *ListConnections200Response`

NewListConnections200Response instantiates a new ListConnections200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnections200ResponseWithDefaults

`func NewListConnections200ResponseWithDefaults() *ListConnections200Response`

NewListConnections200ResponseWithDefaults instantiates a new ListConnections200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListConnections200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListConnections200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListConnections200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *ListConnections200Response) GetData() []Connections`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListConnections200Response) GetDataOk() (*[]Connections, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListConnections200Response) SetData(v []Connections)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



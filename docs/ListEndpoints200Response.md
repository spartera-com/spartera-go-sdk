# ListEndpoints200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**[]Endpoints**](Endpoints.md) |  | 

## Methods

### NewListEndpoints200Response

`func NewListEndpoints200Response(message string, data []Endpoints, ) *ListEndpoints200Response`

NewListEndpoints200Response instantiates a new ListEndpoints200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEndpoints200ResponseWithDefaults

`func NewListEndpoints200ResponseWithDefaults() *ListEndpoints200Response`

NewListEndpoints200ResponseWithDefaults instantiates a new ListEndpoints200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListEndpoints200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListEndpoints200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListEndpoints200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *ListEndpoints200Response) GetData() []Endpoints`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListEndpoints200Response) GetDataOk() (*[]Endpoints, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListEndpoints200Response) SetData(v []Endpoints)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



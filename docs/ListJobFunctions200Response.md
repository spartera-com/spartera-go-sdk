# ListJobFunctions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**[]JobFunctions**](JobFunctions.md) |  | 

## Methods

### NewListJobFunctions200Response

`func NewListJobFunctions200Response(message string, data []JobFunctions, ) *ListJobFunctions200Response`

NewListJobFunctions200Response instantiates a new ListJobFunctions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListJobFunctions200ResponseWithDefaults

`func NewListJobFunctions200ResponseWithDefaults() *ListJobFunctions200Response`

NewListJobFunctions200ResponseWithDefaults instantiates a new ListJobFunctions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListJobFunctions200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListJobFunctions200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListJobFunctions200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *ListJobFunctions200Response) GetData() []JobFunctions`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListJobFunctions200Response) GetDataOk() (*[]JobFunctions, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListJobFunctions200Response) SetData(v []JobFunctions)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



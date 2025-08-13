# JobFunctionsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**[]JobFunctions**](JobFunctions.md) |  | 

## Methods

### NewJobFunctionsGet200Response

`func NewJobFunctionsGet200Response(message string, data []JobFunctions, ) *JobFunctionsGet200Response`

NewJobFunctionsGet200Response instantiates a new JobFunctionsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobFunctionsGet200ResponseWithDefaults

`func NewJobFunctionsGet200ResponseWithDefaults() *JobFunctionsGet200Response`

NewJobFunctionsGet200ResponseWithDefaults instantiates a new JobFunctionsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *JobFunctionsGet200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *JobFunctionsGet200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *JobFunctionsGet200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *JobFunctionsGet200Response) GetData() []JobFunctions`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobFunctionsGet200Response) GetDataOk() (*[]JobFunctions, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobFunctionsGet200Response) SetData(v []JobFunctions)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ListIndustries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**[]Industries**](Industries.md) |  | 

## Methods

### NewListIndustries200Response

`func NewListIndustries200Response(message string, data []Industries, ) *ListIndustries200Response`

NewListIndustries200Response instantiates a new ListIndustries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIndustries200ResponseWithDefaults

`func NewListIndustries200ResponseWithDefaults() *ListIndustries200Response`

NewListIndustries200ResponseWithDefaults instantiates a new ListIndustries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListIndustries200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListIndustries200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListIndustries200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *ListIndustries200Response) GetData() []Industries`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListIndustries200Response) GetDataOk() (*[]Industries, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListIndustries200Response) SetData(v []Industries)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



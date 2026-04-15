# ListStorageEngines200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**[]StorageEngines**](StorageEngines.md) |  | 

## Methods

### NewListStorageEngines200Response

`func NewListStorageEngines200Response(message string, data []StorageEngines, ) *ListStorageEngines200Response`

NewListStorageEngines200Response instantiates a new ListStorageEngines200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageEngines200ResponseWithDefaults

`func NewListStorageEngines200ResponseWithDefaults() *ListStorageEngines200Response`

NewListStorageEngines200ResponseWithDefaults instantiates a new ListStorageEngines200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListStorageEngines200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListStorageEngines200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListStorageEngines200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *ListStorageEngines200Response) GetData() []StorageEngines`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListStorageEngines200Response) GetDataOk() (*[]StorageEngines, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListStorageEngines200Response) SetData(v []StorageEngines)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetStorageEnginesById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**StorageEngines**](StorageEngines.md) |  | 

## Methods

### NewGetStorageEnginesById200Response

`func NewGetStorageEnginesById200Response(message string, data StorageEngines, ) *GetStorageEnginesById200Response`

NewGetStorageEnginesById200Response instantiates a new GetStorageEnginesById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorageEnginesById200ResponseWithDefaults

`func NewGetStorageEnginesById200ResponseWithDefaults() *GetStorageEnginesById200Response`

NewGetStorageEnginesById200ResponseWithDefaults instantiates a new GetStorageEnginesById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetStorageEnginesById200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetStorageEnginesById200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetStorageEnginesById200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetStorageEnginesById200Response) GetData() StorageEngines`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetStorageEnginesById200Response) GetDataOk() (*StorageEngines, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetStorageEnginesById200Response) SetData(v StorageEngines)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



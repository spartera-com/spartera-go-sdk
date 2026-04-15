# CreateAlerts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**CreateAlerts200ResponseData**](CreateAlerts200ResponseData.md) |  | 

## Methods

### NewCreateAlerts200Response

`func NewCreateAlerts200Response(message string, data CreateAlerts200ResponseData, ) *CreateAlerts200Response`

NewCreateAlerts200Response instantiates a new CreateAlerts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlerts200ResponseWithDefaults

`func NewCreateAlerts200ResponseWithDefaults() *CreateAlerts200Response`

NewCreateAlerts200ResponseWithDefaults instantiates a new CreateAlerts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateAlerts200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateAlerts200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateAlerts200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *CreateAlerts200Response) GetData() CreateAlerts200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateAlerts200Response) GetDataOk() (*CreateAlerts200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateAlerts200Response) SetData(v CreateAlerts200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



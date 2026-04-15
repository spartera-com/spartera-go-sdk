# GetAlertsById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**Alerts**](Alerts.md) |  | 

## Methods

### NewGetAlertsById200Response

`func NewGetAlertsById200Response(message string, data Alerts, ) *GetAlertsById200Response`

NewGetAlertsById200Response instantiates a new GetAlertsById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlertsById200ResponseWithDefaults

`func NewGetAlertsById200ResponseWithDefaults() *GetAlertsById200Response`

NewGetAlertsById200ResponseWithDefaults instantiates a new GetAlertsById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetAlertsById200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAlertsById200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAlertsById200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetAlertsById200Response) GetData() Alerts`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAlertsById200Response) GetDataOk() (*Alerts, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAlertsById200Response) SetData(v Alerts)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



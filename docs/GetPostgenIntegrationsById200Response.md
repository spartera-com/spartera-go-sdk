# GetPostgenIntegrationsById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**PostgenIntegrations**](PostgenIntegrations.md) |  | 

## Methods

### NewGetPostgenIntegrationsById200Response

`func NewGetPostgenIntegrationsById200Response(message string, data PostgenIntegrations, ) *GetPostgenIntegrationsById200Response`

NewGetPostgenIntegrationsById200Response instantiates a new GetPostgenIntegrationsById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPostgenIntegrationsById200ResponseWithDefaults

`func NewGetPostgenIntegrationsById200ResponseWithDefaults() *GetPostgenIntegrationsById200Response`

NewGetPostgenIntegrationsById200ResponseWithDefaults instantiates a new GetPostgenIntegrationsById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetPostgenIntegrationsById200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetPostgenIntegrationsById200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetPostgenIntegrationsById200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetPostgenIntegrationsById200Response) GetData() PostgenIntegrations`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPostgenIntegrationsById200Response) GetDataOk() (*PostgenIntegrations, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPostgenIntegrationsById200Response) SetData(v PostgenIntegrations)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



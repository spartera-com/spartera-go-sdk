# ListPostgenIntegrations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**[]PostgenIntegrations**](PostgenIntegrations.md) |  | 

## Methods

### NewListPostgenIntegrations200Response

`func NewListPostgenIntegrations200Response(message string, data []PostgenIntegrations, ) *ListPostgenIntegrations200Response`

NewListPostgenIntegrations200Response instantiates a new ListPostgenIntegrations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPostgenIntegrations200ResponseWithDefaults

`func NewListPostgenIntegrations200ResponseWithDefaults() *ListPostgenIntegrations200Response`

NewListPostgenIntegrations200ResponseWithDefaults instantiates a new ListPostgenIntegrations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListPostgenIntegrations200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListPostgenIntegrations200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListPostgenIntegrations200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *ListPostgenIntegrations200Response) GetData() []PostgenIntegrations`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPostgenIntegrations200Response) GetDataOk() (*[]PostgenIntegrations, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPostgenIntegrations200Response) SetData(v []PostgenIntegrations)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



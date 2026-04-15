# GetEndpointsByIdConnectionsDescribe200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**Endpoints**](Endpoints.md) |  | 

## Methods

### NewGetEndpointsByIdConnectionsDescribe200Response

`func NewGetEndpointsByIdConnectionsDescribe200Response(message string, data Endpoints, ) *GetEndpointsByIdConnectionsDescribe200Response`

NewGetEndpointsByIdConnectionsDescribe200Response instantiates a new GetEndpointsByIdConnectionsDescribe200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEndpointsByIdConnectionsDescribe200ResponseWithDefaults

`func NewGetEndpointsByIdConnectionsDescribe200ResponseWithDefaults() *GetEndpointsByIdConnectionsDescribe200Response`

NewGetEndpointsByIdConnectionsDescribe200ResponseWithDefaults instantiates a new GetEndpointsByIdConnectionsDescribe200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetEndpointsByIdConnectionsDescribe200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetEndpointsByIdConnectionsDescribe200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetEndpointsByIdConnectionsDescribe200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetEndpointsByIdConnectionsDescribe200Response) GetData() Endpoints`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEndpointsByIdConnectionsDescribe200Response) GetDataOk() (*Endpoints, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEndpointsByIdConnectionsDescribe200Response) SetData(v Endpoints)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



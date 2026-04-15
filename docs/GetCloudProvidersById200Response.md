# GetCloudProvidersById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**CloudProviders**](CloudProviders.md) |  | 

## Methods

### NewGetCloudProvidersById200Response

`func NewGetCloudProvidersById200Response(message string, data CloudProviders, ) *GetCloudProvidersById200Response`

NewGetCloudProvidersById200Response instantiates a new GetCloudProvidersById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCloudProvidersById200ResponseWithDefaults

`func NewGetCloudProvidersById200ResponseWithDefaults() *GetCloudProvidersById200Response`

NewGetCloudProvidersById200ResponseWithDefaults instantiates a new GetCloudProvidersById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetCloudProvidersById200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetCloudProvidersById200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetCloudProvidersById200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetCloudProvidersById200Response) GetData() CloudProviders`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCloudProvidersById200Response) GetDataOk() (*CloudProviders, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCloudProvidersById200Response) SetData(v CloudProviders)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



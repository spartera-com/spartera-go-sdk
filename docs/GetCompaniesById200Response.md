# GetCompaniesById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**Companies**](Companies.md) |  | 

## Methods

### NewGetCompaniesById200Response

`func NewGetCompaniesById200Response(message string, data Companies, ) *GetCompaniesById200Response`

NewGetCompaniesById200Response instantiates a new GetCompaniesById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCompaniesById200ResponseWithDefaults

`func NewGetCompaniesById200ResponseWithDefaults() *GetCompaniesById200Response`

NewGetCompaniesById200ResponseWithDefaults instantiates a new GetCompaniesById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetCompaniesById200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetCompaniesById200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetCompaniesById200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetCompaniesById200Response) GetData() Companies`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCompaniesById200Response) GetDataOk() (*Companies, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCompaniesById200Response) SetData(v Companies)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



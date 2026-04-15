# GetAssetUsecasesById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**AssetUsecases**](AssetUsecases.md) |  | 

## Methods

### NewGetAssetUsecasesById200Response

`func NewGetAssetUsecasesById200Response(message string, data AssetUsecases, ) *GetAssetUsecasesById200Response`

NewGetAssetUsecasesById200Response instantiates a new GetAssetUsecasesById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetUsecasesById200ResponseWithDefaults

`func NewGetAssetUsecasesById200ResponseWithDefaults() *GetAssetUsecasesById200Response`

NewGetAssetUsecasesById200ResponseWithDefaults instantiates a new GetAssetUsecasesById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetAssetUsecasesById200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAssetUsecasesById200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAssetUsecasesById200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetAssetUsecasesById200Response) GetData() AssetUsecases`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAssetUsecasesById200Response) GetDataOk() (*AssetUsecases, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAssetUsecasesById200Response) SetData(v AssetUsecases)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



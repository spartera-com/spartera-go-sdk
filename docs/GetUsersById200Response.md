# GetUsersById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**Users**](Users.md) |  | 

## Methods

### NewGetUsersById200Response

`func NewGetUsersById200Response(message string, data Users, ) *GetUsersById200Response`

NewGetUsersById200Response instantiates a new GetUsersById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsersById200ResponseWithDefaults

`func NewGetUsersById200ResponseWithDefaults() *GetUsersById200Response`

NewGetUsersById200ResponseWithDefaults instantiates a new GetUsersById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetUsersById200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetUsersById200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetUsersById200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetUsersById200Response) GetData() Users`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUsersById200Response) GetDataOk() (*Users, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUsersById200Response) SetData(v Users)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



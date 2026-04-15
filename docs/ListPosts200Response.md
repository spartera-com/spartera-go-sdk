# ListPosts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**[]Posts**](Posts.md) |  | 

## Methods

### NewListPosts200Response

`func NewListPosts200Response(message string, data []Posts, ) *ListPosts200Response`

NewListPosts200Response instantiates a new ListPosts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPosts200ResponseWithDefaults

`func NewListPosts200ResponseWithDefaults() *ListPosts200Response`

NewListPosts200ResponseWithDefaults instantiates a new ListPosts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListPosts200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListPosts200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListPosts200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *ListPosts200Response) GetData() []Posts`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPosts200Response) GetDataOk() (*[]Posts, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPosts200Response) SetData(v []Posts)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



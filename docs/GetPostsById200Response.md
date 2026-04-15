# GetPostsById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Response status message | 
**Data** | [**Posts**](Posts.md) |  | 

## Methods

### NewGetPostsById200Response

`func NewGetPostsById200Response(message string, data Posts, ) *GetPostsById200Response`

NewGetPostsById200Response instantiates a new GetPostsById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPostsById200ResponseWithDefaults

`func NewGetPostsById200ResponseWithDefaults() *GetPostsById200Response`

NewGetPostsById200ResponseWithDefaults instantiates a new GetPostsById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetPostsById200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetPostsById200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetPostsById200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetPostsById200Response) GetData() Posts`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPostsById200Response) GetDataOk() (*Posts, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPostsById200Response) SetData(v Posts)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



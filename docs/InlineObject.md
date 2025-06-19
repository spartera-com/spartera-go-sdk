# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject() *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InlineObject) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineObject) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineObject) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineObject) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *InlineObject) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineObject) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineObject) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineObject) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *InlineObject) GetErrors() map[string][]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineObject) GetErrorsOk() (*map[string][]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineObject) SetErrors(v map[string][]string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineObject) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



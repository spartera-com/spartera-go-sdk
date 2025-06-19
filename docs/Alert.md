# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertId** | Pointer to **string** |  | [optional] [readonly] 
**AssetId** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**CompanyId** | **string** |  | 
**AlertType** | **string** |  | 
**IsActive** | Pointer to **string** |  | [optional] 
**LastTriggered** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **string** |  | [optional] [readonly] 
**Active** | Pointer to **string** |  | [optional] 

## Methods

### NewAlert

`func NewAlert(assetId string, companyId string, alertType string, ) *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertId

`func (o *Alert) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *Alert) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *Alert) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *Alert) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetAssetId

`func (o *Alert) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Alert) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Alert) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetUserId

`func (o *Alert) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Alert) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Alert) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Alert) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *Alert) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Alert) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Alert) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetAlertType

`func (o *Alert) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *Alert) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *Alert) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.


### GetIsActive

`func (o *Alert) GetIsActive() string`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Alert) GetIsActiveOk() (*string, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Alert) SetIsActive(v string)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Alert) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastTriggered

`func (o *Alert) GetLastTriggered() string`

GetLastTriggered returns the LastTriggered field if non-nil, zero value otherwise.

### GetLastTriggeredOk

`func (o *Alert) GetLastTriggeredOk() (*string, bool)`

GetLastTriggeredOk returns a tuple with the LastTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggered

`func (o *Alert) SetLastTriggered(v string)`

SetLastTriggered sets LastTriggered field to given value.

### HasLastTriggered

`func (o *Alert) HasLastTriggered() bool`

HasLastTriggered returns a boolean if a field has been set.

### GetDateCreated

`func (o *Alert) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Alert) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Alert) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Alert) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Alert) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Alert) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Alert) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Alert) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *Alert) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Alert) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Alert) SetActive(v string)`

SetActive sets Active field to given value.

### HasActive

`func (o *Alert) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



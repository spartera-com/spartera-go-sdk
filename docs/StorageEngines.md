# StorageEngines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** | Optional. | [optional] 
**LastUpdated** | Pointer to **time.Time** | Optional. | [optional] 
**EngineId** | Pointer to **int64** | Unique identifier. | [optional] 
**ProviderId** | **int64** | References cloud_providers.provider_id — Supported cloud platforms and database engines available for connections. See GET /cloud_providers for valid values. Required. | 
**ServiceName** | **string** | Required. | 
**EngineType** | **string** | Required. One of: EDW, LLM, RDBMS, OBJ, API_MODEL, … (6 total). | 
**IntegrationComplete** | Pointer to **bool** | Optional. | [optional] 
**TestFuncComplete** | Pointer to **bool** | Optional. | [optional] 

## Methods

### NewStorageEngines

`func NewStorageEngines(providerId int64, serviceName string, engineType string, ) *StorageEngines`

NewStorageEngines instantiates a new StorageEngines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageEnginesWithDefaults

`func NewStorageEnginesWithDefaults() *StorageEngines`

NewStorageEnginesWithDefaults instantiates a new StorageEngines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *StorageEngines) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *StorageEngines) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *StorageEngines) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *StorageEngines) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *StorageEngines) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *StorageEngines) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *StorageEngines) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *StorageEngines) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEngineId

`func (o *StorageEngines) GetEngineId() int64`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *StorageEngines) GetEngineIdOk() (*int64, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *StorageEngines) SetEngineId(v int64)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *StorageEngines) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetProviderId

`func (o *StorageEngines) GetProviderId() int64`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *StorageEngines) GetProviderIdOk() (*int64, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *StorageEngines) SetProviderId(v int64)`

SetProviderId sets ProviderId field to given value.


### GetServiceName

`func (o *StorageEngines) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *StorageEngines) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *StorageEngines) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetEngineType

`func (o *StorageEngines) GetEngineType() string`

GetEngineType returns the EngineType field if non-nil, zero value otherwise.

### GetEngineTypeOk

`func (o *StorageEngines) GetEngineTypeOk() (*string, bool)`

GetEngineTypeOk returns a tuple with the EngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineType

`func (o *StorageEngines) SetEngineType(v string)`

SetEngineType sets EngineType field to given value.


### GetIntegrationComplete

`func (o *StorageEngines) GetIntegrationComplete() bool`

GetIntegrationComplete returns the IntegrationComplete field if non-nil, zero value otherwise.

### GetIntegrationCompleteOk

`func (o *StorageEngines) GetIntegrationCompleteOk() (*bool, bool)`

GetIntegrationCompleteOk returns a tuple with the IntegrationComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationComplete

`func (o *StorageEngines) SetIntegrationComplete(v bool)`

SetIntegrationComplete sets IntegrationComplete field to given value.

### HasIntegrationComplete

`func (o *StorageEngines) HasIntegrationComplete() bool`

HasIntegrationComplete returns a boolean if a field has been set.

### GetTestFuncComplete

`func (o *StorageEngines) GetTestFuncComplete() bool`

GetTestFuncComplete returns the TestFuncComplete field if non-nil, zero value otherwise.

### GetTestFuncCompleteOk

`func (o *StorageEngines) GetTestFuncCompleteOk() (*bool, bool)`

GetTestFuncCompleteOk returns a tuple with the TestFuncComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestFuncComplete

`func (o *StorageEngines) SetTestFuncComplete(v bool)`

SetTestFuncComplete sets TestFuncComplete field to given value.

### HasTestFuncComplete

`func (o *StorageEngines) HasTestFuncComplete() bool`

HasTestFuncComplete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



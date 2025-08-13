# CloudProviders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**ProviderId** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**ParentCompany** | Pointer to **string** |  | [optional] 
**MarketingHomepageUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudProviders

`func NewCloudProviders(name string, ) *CloudProviders`

NewCloudProviders instantiates a new CloudProviders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProvidersWithDefaults

`func NewCloudProvidersWithDefaults() *CloudProviders`

NewCloudProvidersWithDefaults instantiates a new CloudProviders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *CloudProviders) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CloudProviders) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CloudProviders) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CloudProviders) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CloudProviders) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CloudProviders) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CloudProviders) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CloudProviders) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProviderId

`func (o *CloudProviders) GetProviderId() int64`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CloudProviders) GetProviderIdOk() (*int64, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CloudProviders) SetProviderId(v int64)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *CloudProviders) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetName

`func (o *CloudProviders) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudProviders) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudProviders) SetName(v string)`

SetName sets Name field to given value.


### GetParentCompany

`func (o *CloudProviders) GetParentCompany() string`

GetParentCompany returns the ParentCompany field if non-nil, zero value otherwise.

### GetParentCompanyOk

`func (o *CloudProviders) GetParentCompanyOk() (*string, bool)`

GetParentCompanyOk returns a tuple with the ParentCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCompany

`func (o *CloudProviders) SetParentCompany(v string)`

SetParentCompany sets ParentCompany field to given value.

### HasParentCompany

`func (o *CloudProviders) HasParentCompany() bool`

HasParentCompany returns a boolean if a field has been set.

### GetMarketingHomepageUrl

`func (o *CloudProviders) GetMarketingHomepageUrl() string`

GetMarketingHomepageUrl returns the MarketingHomepageUrl field if non-nil, zero value otherwise.

### GetMarketingHomepageUrlOk

`func (o *CloudProviders) GetMarketingHomepageUrlOk() (*string, bool)`

GetMarketingHomepageUrlOk returns a tuple with the MarketingHomepageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingHomepageUrl

`func (o *CloudProviders) SetMarketingHomepageUrl(v string)`

SetMarketingHomepageUrl sets MarketingHomepageUrl field to given value.

### HasMarketingHomepageUrl

`func (o *CloudProviders) HasMarketingHomepageUrl() bool`

HasMarketingHomepageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



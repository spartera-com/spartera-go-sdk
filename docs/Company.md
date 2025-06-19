# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**IndustryId** | Pointer to **string** |  | [optional] 
**CountryId** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CompanyDescription** | Pointer to **string** |  | [optional] 
**CompanyHandle** | **string** |  | 
**CompanyDomain** | **string** |  | 
**CreditsBalance** | **string** |  | 
**AcceptedEula** | Pointer to **string** |  | [optional] 
**StripeAccountId** | Pointer to **string** |  | [optional] 
**StripeAccountStatus** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **string** |  | [optional] 

## Methods

### NewCompany

`func NewCompany(companyHandle string, companyDomain string, creditsBalance string, ) *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *Company) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Company) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Company) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *Company) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetIndustryId

`func (o *Company) GetIndustryId() string`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *Company) GetIndustryIdOk() (*string, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *Company) SetIndustryId(v string)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *Company) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetCountryId

`func (o *Company) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *Company) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *Company) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *Company) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetCompanyName

`func (o *Company) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Company) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Company) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Company) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanyDescription

`func (o *Company) GetCompanyDescription() string`

GetCompanyDescription returns the CompanyDescription field if non-nil, zero value otherwise.

### GetCompanyDescriptionOk

`func (o *Company) GetCompanyDescriptionOk() (*string, bool)`

GetCompanyDescriptionOk returns a tuple with the CompanyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDescription

`func (o *Company) SetCompanyDescription(v string)`

SetCompanyDescription sets CompanyDescription field to given value.

### HasCompanyDescription

`func (o *Company) HasCompanyDescription() bool`

HasCompanyDescription returns a boolean if a field has been set.

### GetCompanyHandle

`func (o *Company) GetCompanyHandle() string`

GetCompanyHandle returns the CompanyHandle field if non-nil, zero value otherwise.

### GetCompanyHandleOk

`func (o *Company) GetCompanyHandleOk() (*string, bool)`

GetCompanyHandleOk returns a tuple with the CompanyHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyHandle

`func (o *Company) SetCompanyHandle(v string)`

SetCompanyHandle sets CompanyHandle field to given value.


### GetCompanyDomain

`func (o *Company) GetCompanyDomain() string`

GetCompanyDomain returns the CompanyDomain field if non-nil, zero value otherwise.

### GetCompanyDomainOk

`func (o *Company) GetCompanyDomainOk() (*string, bool)`

GetCompanyDomainOk returns a tuple with the CompanyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDomain

`func (o *Company) SetCompanyDomain(v string)`

SetCompanyDomain sets CompanyDomain field to given value.


### GetCreditsBalance

`func (o *Company) GetCreditsBalance() string`

GetCreditsBalance returns the CreditsBalance field if non-nil, zero value otherwise.

### GetCreditsBalanceOk

`func (o *Company) GetCreditsBalanceOk() (*string, bool)`

GetCreditsBalanceOk returns a tuple with the CreditsBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsBalance

`func (o *Company) SetCreditsBalance(v string)`

SetCreditsBalance sets CreditsBalance field to given value.


### GetAcceptedEula

`func (o *Company) GetAcceptedEula() string`

GetAcceptedEula returns the AcceptedEula field if non-nil, zero value otherwise.

### GetAcceptedEulaOk

`func (o *Company) GetAcceptedEulaOk() (*string, bool)`

GetAcceptedEulaOk returns a tuple with the AcceptedEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedEula

`func (o *Company) SetAcceptedEula(v string)`

SetAcceptedEula sets AcceptedEula field to given value.

### HasAcceptedEula

`func (o *Company) HasAcceptedEula() bool`

HasAcceptedEula returns a boolean if a field has been set.

### GetStripeAccountId

`func (o *Company) GetStripeAccountId() string`

GetStripeAccountId returns the StripeAccountId field if non-nil, zero value otherwise.

### GetStripeAccountIdOk

`func (o *Company) GetStripeAccountIdOk() (*string, bool)`

GetStripeAccountIdOk returns a tuple with the StripeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeAccountId

`func (o *Company) SetStripeAccountId(v string)`

SetStripeAccountId sets StripeAccountId field to given value.

### HasStripeAccountId

`func (o *Company) HasStripeAccountId() bool`

HasStripeAccountId returns a boolean if a field has been set.

### GetStripeAccountStatus

`func (o *Company) GetStripeAccountStatus() string`

GetStripeAccountStatus returns the StripeAccountStatus field if non-nil, zero value otherwise.

### GetStripeAccountStatusOk

`func (o *Company) GetStripeAccountStatusOk() (*string, bool)`

GetStripeAccountStatusOk returns a tuple with the StripeAccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeAccountStatus

`func (o *Company) SetStripeAccountStatus(v string)`

SetStripeAccountStatus sets StripeAccountStatus field to given value.

### HasStripeAccountStatus

`func (o *Company) HasStripeAccountStatus() bool`

HasStripeAccountStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *Company) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Company) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Company) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Company) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Company) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Company) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Company) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Company) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *Company) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Company) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Company) SetActive(v string)`

SetActive sets Active field to given value.

### HasActive

`func (o *Company) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



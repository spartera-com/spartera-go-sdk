# Companies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**IndustryId** | Pointer to **int64** |  | [optional] 
**CountryId** | Pointer to **int64** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CompanyDescription** | Pointer to **string** |  | [optional] 
**CompanyHandle** | **string** |  | 
**CompanyDomain** | **string** |  | 
**CreditsBalance** | **int64** | Current balance of credits for this company (buyer) | 
**AcceptedEula** | Pointer to **bool** |  | [optional] 
**StripeAccountId** | Pointer to **string** | Stripe Connect account ID for marketplace sellers | [optional] 
**StripeAccountStatus** | Pointer to **string** | Status of the Stripe account (verified, pending, etc.) | [optional] 

## Methods

### NewCompanies

`func NewCompanies(companyHandle string, companyDomain string, creditsBalance int64, ) *Companies`

NewCompanies instantiates a new Companies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompaniesWithDefaults

`func NewCompaniesWithDefaults() *Companies`

NewCompaniesWithDefaults instantiates a new Companies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *Companies) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Companies) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Companies) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Companies) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Companies) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Companies) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Companies) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Companies) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCompanyId

`func (o *Companies) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Companies) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Companies) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *Companies) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetIndustryId

`func (o *Companies) GetIndustryId() int64`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *Companies) GetIndustryIdOk() (*int64, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *Companies) SetIndustryId(v int64)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *Companies) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetCountryId

`func (o *Companies) GetCountryId() int64`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *Companies) GetCountryIdOk() (*int64, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *Companies) SetCountryId(v int64)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *Companies) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetCompanyName

`func (o *Companies) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Companies) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Companies) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Companies) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanyDescription

`func (o *Companies) GetCompanyDescription() string`

GetCompanyDescription returns the CompanyDescription field if non-nil, zero value otherwise.

### GetCompanyDescriptionOk

`func (o *Companies) GetCompanyDescriptionOk() (*string, bool)`

GetCompanyDescriptionOk returns a tuple with the CompanyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDescription

`func (o *Companies) SetCompanyDescription(v string)`

SetCompanyDescription sets CompanyDescription field to given value.

### HasCompanyDescription

`func (o *Companies) HasCompanyDescription() bool`

HasCompanyDescription returns a boolean if a field has been set.

### GetCompanyHandle

`func (o *Companies) GetCompanyHandle() string`

GetCompanyHandle returns the CompanyHandle field if non-nil, zero value otherwise.

### GetCompanyHandleOk

`func (o *Companies) GetCompanyHandleOk() (*string, bool)`

GetCompanyHandleOk returns a tuple with the CompanyHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyHandle

`func (o *Companies) SetCompanyHandle(v string)`

SetCompanyHandle sets CompanyHandle field to given value.


### GetCompanyDomain

`func (o *Companies) GetCompanyDomain() string`

GetCompanyDomain returns the CompanyDomain field if non-nil, zero value otherwise.

### GetCompanyDomainOk

`func (o *Companies) GetCompanyDomainOk() (*string, bool)`

GetCompanyDomainOk returns a tuple with the CompanyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDomain

`func (o *Companies) SetCompanyDomain(v string)`

SetCompanyDomain sets CompanyDomain field to given value.


### GetCreditsBalance

`func (o *Companies) GetCreditsBalance() int64`

GetCreditsBalance returns the CreditsBalance field if non-nil, zero value otherwise.

### GetCreditsBalanceOk

`func (o *Companies) GetCreditsBalanceOk() (*int64, bool)`

GetCreditsBalanceOk returns a tuple with the CreditsBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsBalance

`func (o *Companies) SetCreditsBalance(v int64)`

SetCreditsBalance sets CreditsBalance field to given value.


### GetAcceptedEula

`func (o *Companies) GetAcceptedEula() bool`

GetAcceptedEula returns the AcceptedEula field if non-nil, zero value otherwise.

### GetAcceptedEulaOk

`func (o *Companies) GetAcceptedEulaOk() (*bool, bool)`

GetAcceptedEulaOk returns a tuple with the AcceptedEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedEula

`func (o *Companies) SetAcceptedEula(v bool)`

SetAcceptedEula sets AcceptedEula field to given value.

### HasAcceptedEula

`func (o *Companies) HasAcceptedEula() bool`

HasAcceptedEula returns a boolean if a field has been set.

### GetStripeAccountId

`func (o *Companies) GetStripeAccountId() string`

GetStripeAccountId returns the StripeAccountId field if non-nil, zero value otherwise.

### GetStripeAccountIdOk

`func (o *Companies) GetStripeAccountIdOk() (*string, bool)`

GetStripeAccountIdOk returns a tuple with the StripeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeAccountId

`func (o *Companies) SetStripeAccountId(v string)`

SetStripeAccountId sets StripeAccountId field to given value.

### HasStripeAccountId

`func (o *Companies) HasStripeAccountId() bool`

HasStripeAccountId returns a boolean if a field has been set.

### GetStripeAccountStatus

`func (o *Companies) GetStripeAccountStatus() string`

GetStripeAccountStatus returns the StripeAccountStatus field if non-nil, zero value otherwise.

### GetStripeAccountStatusOk

`func (o *Companies) GetStripeAccountStatusOk() (*string, bool)`

GetStripeAccountStatusOk returns a tuple with the StripeAccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeAccountStatus

`func (o *Companies) SetStripeAccountStatus(v string)`

SetStripeAccountStatus sets StripeAccountStatus field to given value.

### HasStripeAccountStatus

`func (o *Companies) HasStripeAccountStatus() bool`

HasStripeAccountStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



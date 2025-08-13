# CompaniesUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndustryId** | Pointer to **int64** |  | [optional] 
**CountryId** | Pointer to **int64** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**CompanyDescription** | Pointer to **string** |  | [optional] 
**CompanyHandle** | Pointer to **string** |  | [optional] 
**CompanyDomain** | Pointer to **string** |  | [optional] 
**CreditsBalance** | Pointer to **int64** | Current balance of credits for this company (buyer) | [optional] 
**AcceptedEula** | Pointer to **bool** |  | [optional] 
**StripeAccountId** | Pointer to **string** | Stripe Connect account ID for marketplace sellers | [optional] 
**StripeAccountStatus** | Pointer to **string** | Status of the Stripe account (verified, pending, etc.) | [optional] 

## Methods

### NewCompaniesUpdate

`func NewCompaniesUpdate() *CompaniesUpdate`

NewCompaniesUpdate instantiates a new CompaniesUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompaniesUpdateWithDefaults

`func NewCompaniesUpdateWithDefaults() *CompaniesUpdate`

NewCompaniesUpdateWithDefaults instantiates a new CompaniesUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndustryId

`func (o *CompaniesUpdate) GetIndustryId() int64`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *CompaniesUpdate) GetIndustryIdOk() (*int64, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *CompaniesUpdate) SetIndustryId(v int64)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *CompaniesUpdate) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetCountryId

`func (o *CompaniesUpdate) GetCountryId() int64`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *CompaniesUpdate) GetCountryIdOk() (*int64, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *CompaniesUpdate) SetCountryId(v int64)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *CompaniesUpdate) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetCompanyName

`func (o *CompaniesUpdate) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompaniesUpdate) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompaniesUpdate) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompaniesUpdate) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanyDescription

`func (o *CompaniesUpdate) GetCompanyDescription() string`

GetCompanyDescription returns the CompanyDescription field if non-nil, zero value otherwise.

### GetCompanyDescriptionOk

`func (o *CompaniesUpdate) GetCompanyDescriptionOk() (*string, bool)`

GetCompanyDescriptionOk returns a tuple with the CompanyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDescription

`func (o *CompaniesUpdate) SetCompanyDescription(v string)`

SetCompanyDescription sets CompanyDescription field to given value.

### HasCompanyDescription

`func (o *CompaniesUpdate) HasCompanyDescription() bool`

HasCompanyDescription returns a boolean if a field has been set.

### GetCompanyHandle

`func (o *CompaniesUpdate) GetCompanyHandle() string`

GetCompanyHandle returns the CompanyHandle field if non-nil, zero value otherwise.

### GetCompanyHandleOk

`func (o *CompaniesUpdate) GetCompanyHandleOk() (*string, bool)`

GetCompanyHandleOk returns a tuple with the CompanyHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyHandle

`func (o *CompaniesUpdate) SetCompanyHandle(v string)`

SetCompanyHandle sets CompanyHandle field to given value.

### HasCompanyHandle

`func (o *CompaniesUpdate) HasCompanyHandle() bool`

HasCompanyHandle returns a boolean if a field has been set.

### GetCompanyDomain

`func (o *CompaniesUpdate) GetCompanyDomain() string`

GetCompanyDomain returns the CompanyDomain field if non-nil, zero value otherwise.

### GetCompanyDomainOk

`func (o *CompaniesUpdate) GetCompanyDomainOk() (*string, bool)`

GetCompanyDomainOk returns a tuple with the CompanyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDomain

`func (o *CompaniesUpdate) SetCompanyDomain(v string)`

SetCompanyDomain sets CompanyDomain field to given value.

### HasCompanyDomain

`func (o *CompaniesUpdate) HasCompanyDomain() bool`

HasCompanyDomain returns a boolean if a field has been set.

### GetCreditsBalance

`func (o *CompaniesUpdate) GetCreditsBalance() int64`

GetCreditsBalance returns the CreditsBalance field if non-nil, zero value otherwise.

### GetCreditsBalanceOk

`func (o *CompaniesUpdate) GetCreditsBalanceOk() (*int64, bool)`

GetCreditsBalanceOk returns a tuple with the CreditsBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsBalance

`func (o *CompaniesUpdate) SetCreditsBalance(v int64)`

SetCreditsBalance sets CreditsBalance field to given value.

### HasCreditsBalance

`func (o *CompaniesUpdate) HasCreditsBalance() bool`

HasCreditsBalance returns a boolean if a field has been set.

### GetAcceptedEula

`func (o *CompaniesUpdate) GetAcceptedEula() bool`

GetAcceptedEula returns the AcceptedEula field if non-nil, zero value otherwise.

### GetAcceptedEulaOk

`func (o *CompaniesUpdate) GetAcceptedEulaOk() (*bool, bool)`

GetAcceptedEulaOk returns a tuple with the AcceptedEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedEula

`func (o *CompaniesUpdate) SetAcceptedEula(v bool)`

SetAcceptedEula sets AcceptedEula field to given value.

### HasAcceptedEula

`func (o *CompaniesUpdate) HasAcceptedEula() bool`

HasAcceptedEula returns a boolean if a field has been set.

### GetStripeAccountId

`func (o *CompaniesUpdate) GetStripeAccountId() string`

GetStripeAccountId returns the StripeAccountId field if non-nil, zero value otherwise.

### GetStripeAccountIdOk

`func (o *CompaniesUpdate) GetStripeAccountIdOk() (*string, bool)`

GetStripeAccountIdOk returns a tuple with the StripeAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeAccountId

`func (o *CompaniesUpdate) SetStripeAccountId(v string)`

SetStripeAccountId sets StripeAccountId field to given value.

### HasStripeAccountId

`func (o *CompaniesUpdate) HasStripeAccountId() bool`

HasStripeAccountId returns a boolean if a field has been set.

### GetStripeAccountStatus

`func (o *CompaniesUpdate) GetStripeAccountStatus() string`

GetStripeAccountStatus returns the StripeAccountStatus field if non-nil, zero value otherwise.

### GetStripeAccountStatusOk

`func (o *CompaniesUpdate) GetStripeAccountStatusOk() (*string, bool)`

GetStripeAccountStatusOk returns a tuple with the StripeAccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeAccountStatus

`func (o *CompaniesUpdate) SetStripeAccountStatus(v string)`

SetStripeAccountStatus sets StripeAccountStatus field to given value.

### HasStripeAccountStatus

`func (o *CompaniesUpdate) HasStripeAccountStatus() bool`

HasStripeAccountStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



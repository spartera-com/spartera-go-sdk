# EndpointsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** | Required. | [optional] 
**LastUpdated** | Pointer to **time.Time** | Required. | [optional] 
**UserId** | Pointer to **string** | User who created this endpoint | [optional] 
**CompanyId** | **string** | References companies.company_id — A Spartera seller or buyer company account. See GET /companies for valid values. Required. | 
**ConnectionId** | **string** | Connection to the data source | 
**IndustryId** | Pointer to **int64** | Industry / category for marketplace discovery | [optional] 
**AucId** | Pointer to **int64** | Primary use case for marketplace discovery | [optional] 
**ApprovalStatus** | Pointer to **string** | Approval status before marketplace publication | [optional] 
**ApprovedByUserId** | Pointer to **string** | User who approved this endpoint for marketplace | [optional] 
**ApprovedAt** | Pointer to **time.Time** | Timestamp of marketplace approval | [optional] 
**SellInMarketplace** | Pointer to **bool** | Whether this endpoint appears in the public marketplace | [optional] 
**PriceCredits** | Pointer to **int64** | Credits deducted from the buyer&#39;s pool per successful (200 OK) request. Same credit pool as assets. price_usd kept for billing records / dashboards. | [optional] 
**Name** | **string** | Human-readable name for the endpoint | 
**Slug** | Pointer to **string** | Human-readable, URL-safe slug derived from name at creation time. e.g. &#39;NFL Live Moneyline &amp; Spread Odds&#39; → &#39;nfl-live-moneyline-spread-odds&#39;. Never changes after creation. Unique within company (DB constraint). Creation fails with 409 if a duplicate name exists in the same company. | [optional] 
**Description** | Pointer to **string** | Description of what this endpoint provides | [optional] 
**SourceSchemaName** | Pointer to **string** | Schema/database name where the table resides | [optional] 
**SourceTableName** | Pointer to **string** | Table name to query from | [optional] 
**CustomerName** | Pointer to **string** | Named customer for B2B deals (marketplace uses price_credits instead) | [optional] 
**PriceUsd** | Pointer to **float32** | USD reference price for billing records and seller dashboards | [optional] 
**EndpointSchema** | Pointer to **map[string]interface{}** | Column configurations including aggregations, filters, and visibility. Format: {columns: [{name, type, aggregation, filter, is_hidden, alias, ...}]}. This is the source of truth — SQL is generated at runtime from this configuration. | [optional] 
**RateLimitRequests** | Pointer to **int64** | Number of requests allowed per rate_limit_period | [optional] 
**RateLimitPeriod** | Pointer to **string** | Time period for rate limiting (HOUR, DAY, MONTH) | [optional] 
**RateLimitGranularity** | Pointer to **string** | How to group rate limits (IP, USER, COMPANY, API_KEY, GLOBAL) | [optional] 
**AccessMethod** | Pointer to **string** | Access control method (OPEN, API_KEY, IP, EMAIL, DOMAIN) | [optional] 
**AccessWhitelist** | Pointer to **map[string]interface{}** | List of allowed IPs, emails, or domains based on access_method. Format: {type: &#39;ip&#39;|&#39;email&#39;|&#39;domain&#39;, values: [&#39;192.168.1.1&#39;, ...]} | [optional] 
**Status** | Pointer to **string** | Current status of the endpoint (ACTIVE, INACTIVE, DEPRECATED) | [optional] 
**Tags** | Pointer to **string** | Comma-separated tags for organizing endpoints | [optional] 
**LastAccessed** | Pointer to **time.Time** | When this endpoint was last called | [optional] 
**MaxRecordsPerRequest** | Pointer to **int64** | Seller-enforced row cap per request. Buyers cannot exceed this. Default 1000. | [optional] 
**SampleResponse** | Pointer to **map[string]interface{}** | Last successful {spartera, request, response} envelope. Saved on each successful marketplace run. Displayed as preview on product page load. | [optional] 
**Active** | Pointer to **bool** | Required. | [optional] 

## Methods

### NewEndpointsInput

`func NewEndpointsInput(companyId string, connectionId string, name string, ) *EndpointsInput`

NewEndpointsInput instantiates a new EndpointsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointsInputWithDefaults

`func NewEndpointsInputWithDefaults() *EndpointsInput`

NewEndpointsInputWithDefaults instantiates a new EndpointsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *EndpointsInput) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *EndpointsInput) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *EndpointsInput) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *EndpointsInput) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *EndpointsInput) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EndpointsInput) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EndpointsInput) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *EndpointsInput) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetUserId

`func (o *EndpointsInput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EndpointsInput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EndpointsInput) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EndpointsInput) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *EndpointsInput) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *EndpointsInput) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *EndpointsInput) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetConnectionId

`func (o *EndpointsInput) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *EndpointsInput) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *EndpointsInput) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetIndustryId

`func (o *EndpointsInput) GetIndustryId() int64`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *EndpointsInput) GetIndustryIdOk() (*int64, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *EndpointsInput) SetIndustryId(v int64)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *EndpointsInput) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetAucId

`func (o *EndpointsInput) GetAucId() int64`

GetAucId returns the AucId field if non-nil, zero value otherwise.

### GetAucIdOk

`func (o *EndpointsInput) GetAucIdOk() (*int64, bool)`

GetAucIdOk returns a tuple with the AucId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAucId

`func (o *EndpointsInput) SetAucId(v int64)`

SetAucId sets AucId field to given value.

### HasAucId

`func (o *EndpointsInput) HasAucId() bool`

HasAucId returns a boolean if a field has been set.

### GetApprovalStatus

`func (o *EndpointsInput) GetApprovalStatus() string`

GetApprovalStatus returns the ApprovalStatus field if non-nil, zero value otherwise.

### GetApprovalStatusOk

`func (o *EndpointsInput) GetApprovalStatusOk() (*string, bool)`

GetApprovalStatusOk returns a tuple with the ApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatus

`func (o *EndpointsInput) SetApprovalStatus(v string)`

SetApprovalStatus sets ApprovalStatus field to given value.

### HasApprovalStatus

`func (o *EndpointsInput) HasApprovalStatus() bool`

HasApprovalStatus returns a boolean if a field has been set.

### GetApprovedByUserId

`func (o *EndpointsInput) GetApprovedByUserId() string`

GetApprovedByUserId returns the ApprovedByUserId field if non-nil, zero value otherwise.

### GetApprovedByUserIdOk

`func (o *EndpointsInput) GetApprovedByUserIdOk() (*string, bool)`

GetApprovedByUserIdOk returns a tuple with the ApprovedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedByUserId

`func (o *EndpointsInput) SetApprovedByUserId(v string)`

SetApprovedByUserId sets ApprovedByUserId field to given value.

### HasApprovedByUserId

`func (o *EndpointsInput) HasApprovedByUserId() bool`

HasApprovedByUserId returns a boolean if a field has been set.

### GetApprovedAt

`func (o *EndpointsInput) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *EndpointsInput) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *EndpointsInput) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *EndpointsInput) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetSellInMarketplace

`func (o *EndpointsInput) GetSellInMarketplace() bool`

GetSellInMarketplace returns the SellInMarketplace field if non-nil, zero value otherwise.

### GetSellInMarketplaceOk

`func (o *EndpointsInput) GetSellInMarketplaceOk() (*bool, bool)`

GetSellInMarketplaceOk returns a tuple with the SellInMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellInMarketplace

`func (o *EndpointsInput) SetSellInMarketplace(v bool)`

SetSellInMarketplace sets SellInMarketplace field to given value.

### HasSellInMarketplace

`func (o *EndpointsInput) HasSellInMarketplace() bool`

HasSellInMarketplace returns a boolean if a field has been set.

### GetPriceCredits

`func (o *EndpointsInput) GetPriceCredits() int64`

GetPriceCredits returns the PriceCredits field if non-nil, zero value otherwise.

### GetPriceCreditsOk

`func (o *EndpointsInput) GetPriceCreditsOk() (*int64, bool)`

GetPriceCreditsOk returns a tuple with the PriceCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCredits

`func (o *EndpointsInput) SetPriceCredits(v int64)`

SetPriceCredits sets PriceCredits field to given value.

### HasPriceCredits

`func (o *EndpointsInput) HasPriceCredits() bool`

HasPriceCredits returns a boolean if a field has been set.

### GetName

`func (o *EndpointsInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointsInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointsInput) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *EndpointsInput) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EndpointsInput) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EndpointsInput) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *EndpointsInput) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *EndpointsInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EndpointsInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EndpointsInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EndpointsInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSourceSchemaName

`func (o *EndpointsInput) GetSourceSchemaName() string`

GetSourceSchemaName returns the SourceSchemaName field if non-nil, zero value otherwise.

### GetSourceSchemaNameOk

`func (o *EndpointsInput) GetSourceSchemaNameOk() (*string, bool)`

GetSourceSchemaNameOk returns a tuple with the SourceSchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchemaName

`func (o *EndpointsInput) SetSourceSchemaName(v string)`

SetSourceSchemaName sets SourceSchemaName field to given value.

### HasSourceSchemaName

`func (o *EndpointsInput) HasSourceSchemaName() bool`

HasSourceSchemaName returns a boolean if a field has been set.

### GetSourceTableName

`func (o *EndpointsInput) GetSourceTableName() string`

GetSourceTableName returns the SourceTableName field if non-nil, zero value otherwise.

### GetSourceTableNameOk

`func (o *EndpointsInput) GetSourceTableNameOk() (*string, bool)`

GetSourceTableNameOk returns a tuple with the SourceTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTableName

`func (o *EndpointsInput) SetSourceTableName(v string)`

SetSourceTableName sets SourceTableName field to given value.

### HasSourceTableName

`func (o *EndpointsInput) HasSourceTableName() bool`

HasSourceTableName returns a boolean if a field has been set.

### GetCustomerName

`func (o *EndpointsInput) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *EndpointsInput) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *EndpointsInput) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *EndpointsInput) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetPriceUsd

`func (o *EndpointsInput) GetPriceUsd() float32`

GetPriceUsd returns the PriceUsd field if non-nil, zero value otherwise.

### GetPriceUsdOk

`func (o *EndpointsInput) GetPriceUsdOk() (*float32, bool)`

GetPriceUsdOk returns a tuple with the PriceUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUsd

`func (o *EndpointsInput) SetPriceUsd(v float32)`

SetPriceUsd sets PriceUsd field to given value.

### HasPriceUsd

`func (o *EndpointsInput) HasPriceUsd() bool`

HasPriceUsd returns a boolean if a field has been set.

### GetEndpointSchema

`func (o *EndpointsInput) GetEndpointSchema() map[string]interface{}`

GetEndpointSchema returns the EndpointSchema field if non-nil, zero value otherwise.

### GetEndpointSchemaOk

`func (o *EndpointsInput) GetEndpointSchemaOk() (*map[string]interface{}, bool)`

GetEndpointSchemaOk returns a tuple with the EndpointSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSchema

`func (o *EndpointsInput) SetEndpointSchema(v map[string]interface{})`

SetEndpointSchema sets EndpointSchema field to given value.

### HasEndpointSchema

`func (o *EndpointsInput) HasEndpointSchema() bool`

HasEndpointSchema returns a boolean if a field has been set.

### GetRateLimitRequests

`func (o *EndpointsInput) GetRateLimitRequests() int64`

GetRateLimitRequests returns the RateLimitRequests field if non-nil, zero value otherwise.

### GetRateLimitRequestsOk

`func (o *EndpointsInput) GetRateLimitRequestsOk() (*int64, bool)`

GetRateLimitRequestsOk returns a tuple with the RateLimitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitRequests

`func (o *EndpointsInput) SetRateLimitRequests(v int64)`

SetRateLimitRequests sets RateLimitRequests field to given value.

### HasRateLimitRequests

`func (o *EndpointsInput) HasRateLimitRequests() bool`

HasRateLimitRequests returns a boolean if a field has been set.

### GetRateLimitPeriod

`func (o *EndpointsInput) GetRateLimitPeriod() string`

GetRateLimitPeriod returns the RateLimitPeriod field if non-nil, zero value otherwise.

### GetRateLimitPeriodOk

`func (o *EndpointsInput) GetRateLimitPeriodOk() (*string, bool)`

GetRateLimitPeriodOk returns a tuple with the RateLimitPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPeriod

`func (o *EndpointsInput) SetRateLimitPeriod(v string)`

SetRateLimitPeriod sets RateLimitPeriod field to given value.

### HasRateLimitPeriod

`func (o *EndpointsInput) HasRateLimitPeriod() bool`

HasRateLimitPeriod returns a boolean if a field has been set.

### GetRateLimitGranularity

`func (o *EndpointsInput) GetRateLimitGranularity() string`

GetRateLimitGranularity returns the RateLimitGranularity field if non-nil, zero value otherwise.

### GetRateLimitGranularityOk

`func (o *EndpointsInput) GetRateLimitGranularityOk() (*string, bool)`

GetRateLimitGranularityOk returns a tuple with the RateLimitGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitGranularity

`func (o *EndpointsInput) SetRateLimitGranularity(v string)`

SetRateLimitGranularity sets RateLimitGranularity field to given value.

### HasRateLimitGranularity

`func (o *EndpointsInput) HasRateLimitGranularity() bool`

HasRateLimitGranularity returns a boolean if a field has been set.

### GetAccessMethod

`func (o *EndpointsInput) GetAccessMethod() string`

GetAccessMethod returns the AccessMethod field if non-nil, zero value otherwise.

### GetAccessMethodOk

`func (o *EndpointsInput) GetAccessMethodOk() (*string, bool)`

GetAccessMethodOk returns a tuple with the AccessMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMethod

`func (o *EndpointsInput) SetAccessMethod(v string)`

SetAccessMethod sets AccessMethod field to given value.

### HasAccessMethod

`func (o *EndpointsInput) HasAccessMethod() bool`

HasAccessMethod returns a boolean if a field has been set.

### GetAccessWhitelist

`func (o *EndpointsInput) GetAccessWhitelist() map[string]interface{}`

GetAccessWhitelist returns the AccessWhitelist field if non-nil, zero value otherwise.

### GetAccessWhitelistOk

`func (o *EndpointsInput) GetAccessWhitelistOk() (*map[string]interface{}, bool)`

GetAccessWhitelistOk returns a tuple with the AccessWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessWhitelist

`func (o *EndpointsInput) SetAccessWhitelist(v map[string]interface{})`

SetAccessWhitelist sets AccessWhitelist field to given value.

### HasAccessWhitelist

`func (o *EndpointsInput) HasAccessWhitelist() bool`

HasAccessWhitelist returns a boolean if a field has been set.

### GetStatus

`func (o *EndpointsInput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EndpointsInput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EndpointsInput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EndpointsInput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *EndpointsInput) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EndpointsInput) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EndpointsInput) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EndpointsInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLastAccessed

`func (o *EndpointsInput) GetLastAccessed() time.Time`

GetLastAccessed returns the LastAccessed field if non-nil, zero value otherwise.

### GetLastAccessedOk

`func (o *EndpointsInput) GetLastAccessedOk() (*time.Time, bool)`

GetLastAccessedOk returns a tuple with the LastAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessed

`func (o *EndpointsInput) SetLastAccessed(v time.Time)`

SetLastAccessed sets LastAccessed field to given value.

### HasLastAccessed

`func (o *EndpointsInput) HasLastAccessed() bool`

HasLastAccessed returns a boolean if a field has been set.

### GetMaxRecordsPerRequest

`func (o *EndpointsInput) GetMaxRecordsPerRequest() int64`

GetMaxRecordsPerRequest returns the MaxRecordsPerRequest field if non-nil, zero value otherwise.

### GetMaxRecordsPerRequestOk

`func (o *EndpointsInput) GetMaxRecordsPerRequestOk() (*int64, bool)`

GetMaxRecordsPerRequestOk returns a tuple with the MaxRecordsPerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecordsPerRequest

`func (o *EndpointsInput) SetMaxRecordsPerRequest(v int64)`

SetMaxRecordsPerRequest sets MaxRecordsPerRequest field to given value.

### HasMaxRecordsPerRequest

`func (o *EndpointsInput) HasMaxRecordsPerRequest() bool`

HasMaxRecordsPerRequest returns a boolean if a field has been set.

### GetSampleResponse

`func (o *EndpointsInput) GetSampleResponse() map[string]interface{}`

GetSampleResponse returns the SampleResponse field if non-nil, zero value otherwise.

### GetSampleResponseOk

`func (o *EndpointsInput) GetSampleResponseOk() (*map[string]interface{}, bool)`

GetSampleResponseOk returns a tuple with the SampleResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleResponse

`func (o *EndpointsInput) SetSampleResponse(v map[string]interface{})`

SetSampleResponse sets SampleResponse field to given value.

### HasSampleResponse

`func (o *EndpointsInput) HasSampleResponse() bool`

HasSampleResponse returns a boolean if a field has been set.

### GetActive

`func (o *EndpointsInput) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EndpointsInput) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EndpointsInput) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EndpointsInput) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



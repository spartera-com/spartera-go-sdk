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
**Name** | **string** | Human-readable name for the endpoint | 
**Slug** | Pointer to **string** | Human-readable, URL-safe slug derived from name at creation time. e.g. &#39;NFL Live Moneyline &amp; Spread Odds&#39; → &#39;nfl-live-moneyline-spread-odds&#39;. Never changes after creation. Unique within company (DB constraint). Creation fails with 409 if a duplicate name exists in the same company. | [optional] 
**Description** | Pointer to **string** | Description of what this endpoint provides | [optional] 
**DetailedDescription** | Pointer to **string** | Long-form HTML description for product pages and SEO | [optional] 
**TopQuestions** | Pointer to **string** | Top 3 questions this endpoint can help answer, in English. Stored as JSON array of strings (1-3 items, 15-200 chars each). Strongly encouraged for marketplace endpoints but not required — nudge via UI completeness score, not hard validation. | [optional] 
**SourceSchemaName** | Pointer to **string** | Schema/database name where the table resides | [optional] 
**SourceTableName** | Pointer to **string** | Table name to query from | [optional] 
**CustomerName** | Pointer to **string** | Named customer for B2B deals (pricing handled via asset_price_history) | [optional] 
**EndpointSchema** | Pointer to **map[string]interface{}** | Column configurations including aggregations, filters, and visibility. Format: {columns: [{name, type, aggregation, filter, is_hidden, alias, ...}]}. This is the source of truth — SQL is generated at runtime from this configuration. | [optional] 
**RateLimitNumber** | Pointer to **int64** | Number of requests allowed per rate_limit_period | [optional] 
**RateLimitPeriod** | Pointer to **string** | Time period for rate limiting (HOUR, DAY, MONTH) | [optional] 
**RateLimitGranularity** | Pointer to **string** | How to group rate limits (IP, USER, COMPANY, API_KEY, GLOBAL) | [optional] 
**AccessMethod** | Pointer to **string** | Access control method (OPEN, API_KEY, IP, EMAIL, DOMAIN) | [optional] 
**AccessWhitelist** | Pointer to **map[string]interface{}** | List of allowed IPs, emails, or domains based on access_method. Format: {type: &#39;ip&#39;|&#39;email&#39;|&#39;domain&#39;, values: [&#39;192.168.1.1&#39;, ...]} | [optional] 
**Status** | Pointer to **string** | Current status of the endpoint (ACTIVE, INACTIVE, DEPRECATED) | [optional] 
**DataTimePeriodStart** | Pointer to **time.Time** | Start date of the data time period covered | [optional] 
**DataTimePeriodEnd** | Pointer to **time.Time** | End date of the data time period covered | [optional] 
**DateCollectionStart** | Pointer to **time.Time** | When the seller began actively collecting this data. Distinct from data_time_period_start, which describes when the records themselves begin. Backfilled historical data will have date_collection_start &gt; data_time_period_start. | [optional] 
**GeographicCoverageType** | Pointer to **string** | Type of geographic coverage | [optional] 
**GeographicCoverageDetails** | Pointer to **string** | Specific regions/countries covered (e.g., &#39;United States, Canada, Mexico&#39;) | [optional] 
**DataSourceRefreshFrequency** | Pointer to **string** | How often the source data is refreshed | [optional] 
**Tags** | Pointer to **string** | Comma-separated tags for organizing endpoints | [optional] 
**LastAccessed** | Pointer to **time.Time** | When this endpoint was last called | [optional] 
**MaxRecordsPerRequest** | Pointer to **int64** | Seller-enforced row cap per request. Buyers cannot exceed this. Default 1000. | [optional] 
**ExportEnabled** | Pointer to **bool** | Whether this endpoint supports bulk export to GCS. When True, buyers can request delivery&#x3D;gcs with format&#x3D;csv|parquet. Independent of max_records_per_request, which only governs inline JSON. | [optional] 
**MaxRecordsPerExport** | Pointer to **int64** | Hard ceiling on rows returned per GCS export. Separate from max_records_per_request so sellers can offer larger downloads via file delivery without expanding inline JSON responses. | [optional] 
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

### GetDetailedDescription

`func (o *EndpointsInput) GetDetailedDescription() string`

GetDetailedDescription returns the DetailedDescription field if non-nil, zero value otherwise.

### GetDetailedDescriptionOk

`func (o *EndpointsInput) GetDetailedDescriptionOk() (*string, bool)`

GetDetailedDescriptionOk returns a tuple with the DetailedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedDescription

`func (o *EndpointsInput) SetDetailedDescription(v string)`

SetDetailedDescription sets DetailedDescription field to given value.

### HasDetailedDescription

`func (o *EndpointsInput) HasDetailedDescription() bool`

HasDetailedDescription returns a boolean if a field has been set.

### GetTopQuestions

`func (o *EndpointsInput) GetTopQuestions() string`

GetTopQuestions returns the TopQuestions field if non-nil, zero value otherwise.

### GetTopQuestionsOk

`func (o *EndpointsInput) GetTopQuestionsOk() (*string, bool)`

GetTopQuestionsOk returns a tuple with the TopQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopQuestions

`func (o *EndpointsInput) SetTopQuestions(v string)`

SetTopQuestions sets TopQuestions field to given value.

### HasTopQuestions

`func (o *EndpointsInput) HasTopQuestions() bool`

HasTopQuestions returns a boolean if a field has been set.

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

### GetRateLimitNumber

`func (o *EndpointsInput) GetRateLimitNumber() int64`

GetRateLimitNumber returns the RateLimitNumber field if non-nil, zero value otherwise.

### GetRateLimitNumberOk

`func (o *EndpointsInput) GetRateLimitNumberOk() (*int64, bool)`

GetRateLimitNumberOk returns a tuple with the RateLimitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitNumber

`func (o *EndpointsInput) SetRateLimitNumber(v int64)`

SetRateLimitNumber sets RateLimitNumber field to given value.

### HasRateLimitNumber

`func (o *EndpointsInput) HasRateLimitNumber() bool`

HasRateLimitNumber returns a boolean if a field has been set.

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

### GetDataTimePeriodStart

`func (o *EndpointsInput) GetDataTimePeriodStart() time.Time`

GetDataTimePeriodStart returns the DataTimePeriodStart field if non-nil, zero value otherwise.

### GetDataTimePeriodStartOk

`func (o *EndpointsInput) GetDataTimePeriodStartOk() (*time.Time, bool)`

GetDataTimePeriodStartOk returns a tuple with the DataTimePeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimePeriodStart

`func (o *EndpointsInput) SetDataTimePeriodStart(v time.Time)`

SetDataTimePeriodStart sets DataTimePeriodStart field to given value.

### HasDataTimePeriodStart

`func (o *EndpointsInput) HasDataTimePeriodStart() bool`

HasDataTimePeriodStart returns a boolean if a field has been set.

### GetDataTimePeriodEnd

`func (o *EndpointsInput) GetDataTimePeriodEnd() time.Time`

GetDataTimePeriodEnd returns the DataTimePeriodEnd field if non-nil, zero value otherwise.

### GetDataTimePeriodEndOk

`func (o *EndpointsInput) GetDataTimePeriodEndOk() (*time.Time, bool)`

GetDataTimePeriodEndOk returns a tuple with the DataTimePeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimePeriodEnd

`func (o *EndpointsInput) SetDataTimePeriodEnd(v time.Time)`

SetDataTimePeriodEnd sets DataTimePeriodEnd field to given value.

### HasDataTimePeriodEnd

`func (o *EndpointsInput) HasDataTimePeriodEnd() bool`

HasDataTimePeriodEnd returns a boolean if a field has been set.

### GetDateCollectionStart

`func (o *EndpointsInput) GetDateCollectionStart() time.Time`

GetDateCollectionStart returns the DateCollectionStart field if non-nil, zero value otherwise.

### GetDateCollectionStartOk

`func (o *EndpointsInput) GetDateCollectionStartOk() (*time.Time, bool)`

GetDateCollectionStartOk returns a tuple with the DateCollectionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCollectionStart

`func (o *EndpointsInput) SetDateCollectionStart(v time.Time)`

SetDateCollectionStart sets DateCollectionStart field to given value.

### HasDateCollectionStart

`func (o *EndpointsInput) HasDateCollectionStart() bool`

HasDateCollectionStart returns a boolean if a field has been set.

### GetGeographicCoverageType

`func (o *EndpointsInput) GetGeographicCoverageType() string`

GetGeographicCoverageType returns the GeographicCoverageType field if non-nil, zero value otherwise.

### GetGeographicCoverageTypeOk

`func (o *EndpointsInput) GetGeographicCoverageTypeOk() (*string, bool)`

GetGeographicCoverageTypeOk returns a tuple with the GeographicCoverageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicCoverageType

`func (o *EndpointsInput) SetGeographicCoverageType(v string)`

SetGeographicCoverageType sets GeographicCoverageType field to given value.

### HasGeographicCoverageType

`func (o *EndpointsInput) HasGeographicCoverageType() bool`

HasGeographicCoverageType returns a boolean if a field has been set.

### GetGeographicCoverageDetails

`func (o *EndpointsInput) GetGeographicCoverageDetails() string`

GetGeographicCoverageDetails returns the GeographicCoverageDetails field if non-nil, zero value otherwise.

### GetGeographicCoverageDetailsOk

`func (o *EndpointsInput) GetGeographicCoverageDetailsOk() (*string, bool)`

GetGeographicCoverageDetailsOk returns a tuple with the GeographicCoverageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicCoverageDetails

`func (o *EndpointsInput) SetGeographicCoverageDetails(v string)`

SetGeographicCoverageDetails sets GeographicCoverageDetails field to given value.

### HasGeographicCoverageDetails

`func (o *EndpointsInput) HasGeographicCoverageDetails() bool`

HasGeographicCoverageDetails returns a boolean if a field has been set.

### GetDataSourceRefreshFrequency

`func (o *EndpointsInput) GetDataSourceRefreshFrequency() string`

GetDataSourceRefreshFrequency returns the DataSourceRefreshFrequency field if non-nil, zero value otherwise.

### GetDataSourceRefreshFrequencyOk

`func (o *EndpointsInput) GetDataSourceRefreshFrequencyOk() (*string, bool)`

GetDataSourceRefreshFrequencyOk returns a tuple with the DataSourceRefreshFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceRefreshFrequency

`func (o *EndpointsInput) SetDataSourceRefreshFrequency(v string)`

SetDataSourceRefreshFrequency sets DataSourceRefreshFrequency field to given value.

### HasDataSourceRefreshFrequency

`func (o *EndpointsInput) HasDataSourceRefreshFrequency() bool`

HasDataSourceRefreshFrequency returns a boolean if a field has been set.

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

### GetExportEnabled

`func (o *EndpointsInput) GetExportEnabled() bool`

GetExportEnabled returns the ExportEnabled field if non-nil, zero value otherwise.

### GetExportEnabledOk

`func (o *EndpointsInput) GetExportEnabledOk() (*bool, bool)`

GetExportEnabledOk returns a tuple with the ExportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportEnabled

`func (o *EndpointsInput) SetExportEnabled(v bool)`

SetExportEnabled sets ExportEnabled field to given value.

### HasExportEnabled

`func (o *EndpointsInput) HasExportEnabled() bool`

HasExportEnabled returns a boolean if a field has been set.

### GetMaxRecordsPerExport

`func (o *EndpointsInput) GetMaxRecordsPerExport() int64`

GetMaxRecordsPerExport returns the MaxRecordsPerExport field if non-nil, zero value otherwise.

### GetMaxRecordsPerExportOk

`func (o *EndpointsInput) GetMaxRecordsPerExportOk() (*int64, bool)`

GetMaxRecordsPerExportOk returns a tuple with the MaxRecordsPerExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecordsPerExport

`func (o *EndpointsInput) SetMaxRecordsPerExport(v int64)`

SetMaxRecordsPerExport sets MaxRecordsPerExport field to given value.

### HasMaxRecordsPerExport

`func (o *EndpointsInput) HasMaxRecordsPerExport() bool`

HasMaxRecordsPerExport returns a boolean if a field has been set.

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



# EndpointsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** | Required. | [optional] 
**LastUpdated** | Pointer to **time.Time** | Required. | [optional] 
**UserId** | Pointer to **string** | User who created this endpoint | [optional] 
**CompanyId** | Pointer to **string** | References companies.company_id — A Spartera seller or buyer company account. See GET /companies for valid values. Required. | [optional] 
**ConnectionId** | Pointer to **string** | Connection to the data source | [optional] 
**IndustryId** | Pointer to **int64** | Industry / category for marketplace discovery | [optional] 
**AucId** | Pointer to **int64** | Primary use case for marketplace discovery | [optional] 
**ApprovalStatus** | Pointer to **string** | Approval status before marketplace publication | [optional] 
**ApprovedByUserId** | Pointer to **string** | User who approved this endpoint for marketplace | [optional] 
**ApprovedAt** | Pointer to **time.Time** | Timestamp of marketplace approval | [optional] 
**SellInMarketplace** | Pointer to **bool** | Whether this endpoint appears in the public marketplace | [optional] 
**Name** | Pointer to **string** | Human-readable name for the endpoint | [optional] 
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

### NewEndpointsUpdate

`func NewEndpointsUpdate() *EndpointsUpdate`

NewEndpointsUpdate instantiates a new EndpointsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointsUpdateWithDefaults

`func NewEndpointsUpdateWithDefaults() *EndpointsUpdate`

NewEndpointsUpdateWithDefaults instantiates a new EndpointsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *EndpointsUpdate) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *EndpointsUpdate) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *EndpointsUpdate) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *EndpointsUpdate) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *EndpointsUpdate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EndpointsUpdate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EndpointsUpdate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *EndpointsUpdate) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetUserId

`func (o *EndpointsUpdate) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EndpointsUpdate) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EndpointsUpdate) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EndpointsUpdate) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *EndpointsUpdate) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *EndpointsUpdate) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *EndpointsUpdate) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *EndpointsUpdate) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetConnectionId

`func (o *EndpointsUpdate) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *EndpointsUpdate) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *EndpointsUpdate) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *EndpointsUpdate) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetIndustryId

`func (o *EndpointsUpdate) GetIndustryId() int64`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *EndpointsUpdate) GetIndustryIdOk() (*int64, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *EndpointsUpdate) SetIndustryId(v int64)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *EndpointsUpdate) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetAucId

`func (o *EndpointsUpdate) GetAucId() int64`

GetAucId returns the AucId field if non-nil, zero value otherwise.

### GetAucIdOk

`func (o *EndpointsUpdate) GetAucIdOk() (*int64, bool)`

GetAucIdOk returns a tuple with the AucId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAucId

`func (o *EndpointsUpdate) SetAucId(v int64)`

SetAucId sets AucId field to given value.

### HasAucId

`func (o *EndpointsUpdate) HasAucId() bool`

HasAucId returns a boolean if a field has been set.

### GetApprovalStatus

`func (o *EndpointsUpdate) GetApprovalStatus() string`

GetApprovalStatus returns the ApprovalStatus field if non-nil, zero value otherwise.

### GetApprovalStatusOk

`func (o *EndpointsUpdate) GetApprovalStatusOk() (*string, bool)`

GetApprovalStatusOk returns a tuple with the ApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatus

`func (o *EndpointsUpdate) SetApprovalStatus(v string)`

SetApprovalStatus sets ApprovalStatus field to given value.

### HasApprovalStatus

`func (o *EndpointsUpdate) HasApprovalStatus() bool`

HasApprovalStatus returns a boolean if a field has been set.

### GetApprovedByUserId

`func (o *EndpointsUpdate) GetApprovedByUserId() string`

GetApprovedByUserId returns the ApprovedByUserId field if non-nil, zero value otherwise.

### GetApprovedByUserIdOk

`func (o *EndpointsUpdate) GetApprovedByUserIdOk() (*string, bool)`

GetApprovedByUserIdOk returns a tuple with the ApprovedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedByUserId

`func (o *EndpointsUpdate) SetApprovedByUserId(v string)`

SetApprovedByUserId sets ApprovedByUserId field to given value.

### HasApprovedByUserId

`func (o *EndpointsUpdate) HasApprovedByUserId() bool`

HasApprovedByUserId returns a boolean if a field has been set.

### GetApprovedAt

`func (o *EndpointsUpdate) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *EndpointsUpdate) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *EndpointsUpdate) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *EndpointsUpdate) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetSellInMarketplace

`func (o *EndpointsUpdate) GetSellInMarketplace() bool`

GetSellInMarketplace returns the SellInMarketplace field if non-nil, zero value otherwise.

### GetSellInMarketplaceOk

`func (o *EndpointsUpdate) GetSellInMarketplaceOk() (*bool, bool)`

GetSellInMarketplaceOk returns a tuple with the SellInMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellInMarketplace

`func (o *EndpointsUpdate) SetSellInMarketplace(v bool)`

SetSellInMarketplace sets SellInMarketplace field to given value.

### HasSellInMarketplace

`func (o *EndpointsUpdate) HasSellInMarketplace() bool`

HasSellInMarketplace returns a boolean if a field has been set.

### GetName

`func (o *EndpointsUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointsUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointsUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EndpointsUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *EndpointsUpdate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EndpointsUpdate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EndpointsUpdate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *EndpointsUpdate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *EndpointsUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EndpointsUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EndpointsUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EndpointsUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetailedDescription

`func (o *EndpointsUpdate) GetDetailedDescription() string`

GetDetailedDescription returns the DetailedDescription field if non-nil, zero value otherwise.

### GetDetailedDescriptionOk

`func (o *EndpointsUpdate) GetDetailedDescriptionOk() (*string, bool)`

GetDetailedDescriptionOk returns a tuple with the DetailedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedDescription

`func (o *EndpointsUpdate) SetDetailedDescription(v string)`

SetDetailedDescription sets DetailedDescription field to given value.

### HasDetailedDescription

`func (o *EndpointsUpdate) HasDetailedDescription() bool`

HasDetailedDescription returns a boolean if a field has been set.

### GetTopQuestions

`func (o *EndpointsUpdate) GetTopQuestions() string`

GetTopQuestions returns the TopQuestions field if non-nil, zero value otherwise.

### GetTopQuestionsOk

`func (o *EndpointsUpdate) GetTopQuestionsOk() (*string, bool)`

GetTopQuestionsOk returns a tuple with the TopQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopQuestions

`func (o *EndpointsUpdate) SetTopQuestions(v string)`

SetTopQuestions sets TopQuestions field to given value.

### HasTopQuestions

`func (o *EndpointsUpdate) HasTopQuestions() bool`

HasTopQuestions returns a boolean if a field has been set.

### GetSourceSchemaName

`func (o *EndpointsUpdate) GetSourceSchemaName() string`

GetSourceSchemaName returns the SourceSchemaName field if non-nil, zero value otherwise.

### GetSourceSchemaNameOk

`func (o *EndpointsUpdate) GetSourceSchemaNameOk() (*string, bool)`

GetSourceSchemaNameOk returns a tuple with the SourceSchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchemaName

`func (o *EndpointsUpdate) SetSourceSchemaName(v string)`

SetSourceSchemaName sets SourceSchemaName field to given value.

### HasSourceSchemaName

`func (o *EndpointsUpdate) HasSourceSchemaName() bool`

HasSourceSchemaName returns a boolean if a field has been set.

### GetSourceTableName

`func (o *EndpointsUpdate) GetSourceTableName() string`

GetSourceTableName returns the SourceTableName field if non-nil, zero value otherwise.

### GetSourceTableNameOk

`func (o *EndpointsUpdate) GetSourceTableNameOk() (*string, bool)`

GetSourceTableNameOk returns a tuple with the SourceTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTableName

`func (o *EndpointsUpdate) SetSourceTableName(v string)`

SetSourceTableName sets SourceTableName field to given value.

### HasSourceTableName

`func (o *EndpointsUpdate) HasSourceTableName() bool`

HasSourceTableName returns a boolean if a field has been set.

### GetCustomerName

`func (o *EndpointsUpdate) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *EndpointsUpdate) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *EndpointsUpdate) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *EndpointsUpdate) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetEndpointSchema

`func (o *EndpointsUpdate) GetEndpointSchema() map[string]interface{}`

GetEndpointSchema returns the EndpointSchema field if non-nil, zero value otherwise.

### GetEndpointSchemaOk

`func (o *EndpointsUpdate) GetEndpointSchemaOk() (*map[string]interface{}, bool)`

GetEndpointSchemaOk returns a tuple with the EndpointSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSchema

`func (o *EndpointsUpdate) SetEndpointSchema(v map[string]interface{})`

SetEndpointSchema sets EndpointSchema field to given value.

### HasEndpointSchema

`func (o *EndpointsUpdate) HasEndpointSchema() bool`

HasEndpointSchema returns a boolean if a field has been set.

### GetRateLimitNumber

`func (o *EndpointsUpdate) GetRateLimitNumber() int64`

GetRateLimitNumber returns the RateLimitNumber field if non-nil, zero value otherwise.

### GetRateLimitNumberOk

`func (o *EndpointsUpdate) GetRateLimitNumberOk() (*int64, bool)`

GetRateLimitNumberOk returns a tuple with the RateLimitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitNumber

`func (o *EndpointsUpdate) SetRateLimitNumber(v int64)`

SetRateLimitNumber sets RateLimitNumber field to given value.

### HasRateLimitNumber

`func (o *EndpointsUpdate) HasRateLimitNumber() bool`

HasRateLimitNumber returns a boolean if a field has been set.

### GetRateLimitPeriod

`func (o *EndpointsUpdate) GetRateLimitPeriod() string`

GetRateLimitPeriod returns the RateLimitPeriod field if non-nil, zero value otherwise.

### GetRateLimitPeriodOk

`func (o *EndpointsUpdate) GetRateLimitPeriodOk() (*string, bool)`

GetRateLimitPeriodOk returns a tuple with the RateLimitPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPeriod

`func (o *EndpointsUpdate) SetRateLimitPeriod(v string)`

SetRateLimitPeriod sets RateLimitPeriod field to given value.

### HasRateLimitPeriod

`func (o *EndpointsUpdate) HasRateLimitPeriod() bool`

HasRateLimitPeriod returns a boolean if a field has been set.

### GetRateLimitGranularity

`func (o *EndpointsUpdate) GetRateLimitGranularity() string`

GetRateLimitGranularity returns the RateLimitGranularity field if non-nil, zero value otherwise.

### GetRateLimitGranularityOk

`func (o *EndpointsUpdate) GetRateLimitGranularityOk() (*string, bool)`

GetRateLimitGranularityOk returns a tuple with the RateLimitGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitGranularity

`func (o *EndpointsUpdate) SetRateLimitGranularity(v string)`

SetRateLimitGranularity sets RateLimitGranularity field to given value.

### HasRateLimitGranularity

`func (o *EndpointsUpdate) HasRateLimitGranularity() bool`

HasRateLimitGranularity returns a boolean if a field has been set.

### GetAccessMethod

`func (o *EndpointsUpdate) GetAccessMethod() string`

GetAccessMethod returns the AccessMethod field if non-nil, zero value otherwise.

### GetAccessMethodOk

`func (o *EndpointsUpdate) GetAccessMethodOk() (*string, bool)`

GetAccessMethodOk returns a tuple with the AccessMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMethod

`func (o *EndpointsUpdate) SetAccessMethod(v string)`

SetAccessMethod sets AccessMethod field to given value.

### HasAccessMethod

`func (o *EndpointsUpdate) HasAccessMethod() bool`

HasAccessMethod returns a boolean if a field has been set.

### GetAccessWhitelist

`func (o *EndpointsUpdate) GetAccessWhitelist() map[string]interface{}`

GetAccessWhitelist returns the AccessWhitelist field if non-nil, zero value otherwise.

### GetAccessWhitelistOk

`func (o *EndpointsUpdate) GetAccessWhitelistOk() (*map[string]interface{}, bool)`

GetAccessWhitelistOk returns a tuple with the AccessWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessWhitelist

`func (o *EndpointsUpdate) SetAccessWhitelist(v map[string]interface{})`

SetAccessWhitelist sets AccessWhitelist field to given value.

### HasAccessWhitelist

`func (o *EndpointsUpdate) HasAccessWhitelist() bool`

HasAccessWhitelist returns a boolean if a field has been set.

### GetStatus

`func (o *EndpointsUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EndpointsUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EndpointsUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EndpointsUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDataTimePeriodStart

`func (o *EndpointsUpdate) GetDataTimePeriodStart() time.Time`

GetDataTimePeriodStart returns the DataTimePeriodStart field if non-nil, zero value otherwise.

### GetDataTimePeriodStartOk

`func (o *EndpointsUpdate) GetDataTimePeriodStartOk() (*time.Time, bool)`

GetDataTimePeriodStartOk returns a tuple with the DataTimePeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimePeriodStart

`func (o *EndpointsUpdate) SetDataTimePeriodStart(v time.Time)`

SetDataTimePeriodStart sets DataTimePeriodStart field to given value.

### HasDataTimePeriodStart

`func (o *EndpointsUpdate) HasDataTimePeriodStart() bool`

HasDataTimePeriodStart returns a boolean if a field has been set.

### GetDataTimePeriodEnd

`func (o *EndpointsUpdate) GetDataTimePeriodEnd() time.Time`

GetDataTimePeriodEnd returns the DataTimePeriodEnd field if non-nil, zero value otherwise.

### GetDataTimePeriodEndOk

`func (o *EndpointsUpdate) GetDataTimePeriodEndOk() (*time.Time, bool)`

GetDataTimePeriodEndOk returns a tuple with the DataTimePeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimePeriodEnd

`func (o *EndpointsUpdate) SetDataTimePeriodEnd(v time.Time)`

SetDataTimePeriodEnd sets DataTimePeriodEnd field to given value.

### HasDataTimePeriodEnd

`func (o *EndpointsUpdate) HasDataTimePeriodEnd() bool`

HasDataTimePeriodEnd returns a boolean if a field has been set.

### GetDateCollectionStart

`func (o *EndpointsUpdate) GetDateCollectionStart() time.Time`

GetDateCollectionStart returns the DateCollectionStart field if non-nil, zero value otherwise.

### GetDateCollectionStartOk

`func (o *EndpointsUpdate) GetDateCollectionStartOk() (*time.Time, bool)`

GetDateCollectionStartOk returns a tuple with the DateCollectionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCollectionStart

`func (o *EndpointsUpdate) SetDateCollectionStart(v time.Time)`

SetDateCollectionStart sets DateCollectionStart field to given value.

### HasDateCollectionStart

`func (o *EndpointsUpdate) HasDateCollectionStart() bool`

HasDateCollectionStart returns a boolean if a field has been set.

### GetGeographicCoverageType

`func (o *EndpointsUpdate) GetGeographicCoverageType() string`

GetGeographicCoverageType returns the GeographicCoverageType field if non-nil, zero value otherwise.

### GetGeographicCoverageTypeOk

`func (o *EndpointsUpdate) GetGeographicCoverageTypeOk() (*string, bool)`

GetGeographicCoverageTypeOk returns a tuple with the GeographicCoverageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicCoverageType

`func (o *EndpointsUpdate) SetGeographicCoverageType(v string)`

SetGeographicCoverageType sets GeographicCoverageType field to given value.

### HasGeographicCoverageType

`func (o *EndpointsUpdate) HasGeographicCoverageType() bool`

HasGeographicCoverageType returns a boolean if a field has been set.

### GetGeographicCoverageDetails

`func (o *EndpointsUpdate) GetGeographicCoverageDetails() string`

GetGeographicCoverageDetails returns the GeographicCoverageDetails field if non-nil, zero value otherwise.

### GetGeographicCoverageDetailsOk

`func (o *EndpointsUpdate) GetGeographicCoverageDetailsOk() (*string, bool)`

GetGeographicCoverageDetailsOk returns a tuple with the GeographicCoverageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicCoverageDetails

`func (o *EndpointsUpdate) SetGeographicCoverageDetails(v string)`

SetGeographicCoverageDetails sets GeographicCoverageDetails field to given value.

### HasGeographicCoverageDetails

`func (o *EndpointsUpdate) HasGeographicCoverageDetails() bool`

HasGeographicCoverageDetails returns a boolean if a field has been set.

### GetDataSourceRefreshFrequency

`func (o *EndpointsUpdate) GetDataSourceRefreshFrequency() string`

GetDataSourceRefreshFrequency returns the DataSourceRefreshFrequency field if non-nil, zero value otherwise.

### GetDataSourceRefreshFrequencyOk

`func (o *EndpointsUpdate) GetDataSourceRefreshFrequencyOk() (*string, bool)`

GetDataSourceRefreshFrequencyOk returns a tuple with the DataSourceRefreshFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceRefreshFrequency

`func (o *EndpointsUpdate) SetDataSourceRefreshFrequency(v string)`

SetDataSourceRefreshFrequency sets DataSourceRefreshFrequency field to given value.

### HasDataSourceRefreshFrequency

`func (o *EndpointsUpdate) HasDataSourceRefreshFrequency() bool`

HasDataSourceRefreshFrequency returns a boolean if a field has been set.

### GetTags

`func (o *EndpointsUpdate) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EndpointsUpdate) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EndpointsUpdate) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EndpointsUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLastAccessed

`func (o *EndpointsUpdate) GetLastAccessed() time.Time`

GetLastAccessed returns the LastAccessed field if non-nil, zero value otherwise.

### GetLastAccessedOk

`func (o *EndpointsUpdate) GetLastAccessedOk() (*time.Time, bool)`

GetLastAccessedOk returns a tuple with the LastAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessed

`func (o *EndpointsUpdate) SetLastAccessed(v time.Time)`

SetLastAccessed sets LastAccessed field to given value.

### HasLastAccessed

`func (o *EndpointsUpdate) HasLastAccessed() bool`

HasLastAccessed returns a boolean if a field has been set.

### GetMaxRecordsPerRequest

`func (o *EndpointsUpdate) GetMaxRecordsPerRequest() int64`

GetMaxRecordsPerRequest returns the MaxRecordsPerRequest field if non-nil, zero value otherwise.

### GetMaxRecordsPerRequestOk

`func (o *EndpointsUpdate) GetMaxRecordsPerRequestOk() (*int64, bool)`

GetMaxRecordsPerRequestOk returns a tuple with the MaxRecordsPerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecordsPerRequest

`func (o *EndpointsUpdate) SetMaxRecordsPerRequest(v int64)`

SetMaxRecordsPerRequest sets MaxRecordsPerRequest field to given value.

### HasMaxRecordsPerRequest

`func (o *EndpointsUpdate) HasMaxRecordsPerRequest() bool`

HasMaxRecordsPerRequest returns a boolean if a field has been set.

### GetExportEnabled

`func (o *EndpointsUpdate) GetExportEnabled() bool`

GetExportEnabled returns the ExportEnabled field if non-nil, zero value otherwise.

### GetExportEnabledOk

`func (o *EndpointsUpdate) GetExportEnabledOk() (*bool, bool)`

GetExportEnabledOk returns a tuple with the ExportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportEnabled

`func (o *EndpointsUpdate) SetExportEnabled(v bool)`

SetExportEnabled sets ExportEnabled field to given value.

### HasExportEnabled

`func (o *EndpointsUpdate) HasExportEnabled() bool`

HasExportEnabled returns a boolean if a field has been set.

### GetMaxRecordsPerExport

`func (o *EndpointsUpdate) GetMaxRecordsPerExport() int64`

GetMaxRecordsPerExport returns the MaxRecordsPerExport field if non-nil, zero value otherwise.

### GetMaxRecordsPerExportOk

`func (o *EndpointsUpdate) GetMaxRecordsPerExportOk() (*int64, bool)`

GetMaxRecordsPerExportOk returns a tuple with the MaxRecordsPerExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecordsPerExport

`func (o *EndpointsUpdate) SetMaxRecordsPerExport(v int64)`

SetMaxRecordsPerExport sets MaxRecordsPerExport field to given value.

### HasMaxRecordsPerExport

`func (o *EndpointsUpdate) HasMaxRecordsPerExport() bool`

HasMaxRecordsPerExport returns a boolean if a field has been set.

### GetSampleResponse

`func (o *EndpointsUpdate) GetSampleResponse() map[string]interface{}`

GetSampleResponse returns the SampleResponse field if non-nil, zero value otherwise.

### GetSampleResponseOk

`func (o *EndpointsUpdate) GetSampleResponseOk() (*map[string]interface{}, bool)`

GetSampleResponseOk returns a tuple with the SampleResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleResponse

`func (o *EndpointsUpdate) SetSampleResponse(v map[string]interface{})`

SetSampleResponse sets SampleResponse field to given value.

### HasSampleResponse

`func (o *EndpointsUpdate) HasSampleResponse() bool`

HasSampleResponse returns a boolean if a field has been set.

### GetActive

`func (o *EndpointsUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EndpointsUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EndpointsUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EndpointsUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Assets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** | Optional. | [optional] 
**LastUpdated** | Pointer to **time.Time** | Optional. | [optional] 
**AssetId** | Pointer to **string** | Unique identifier. | [optional] 
**UserId** | Pointer to **string** | References users.user_id — An individual user account within a company. See GET /users for valid values. Optional. | [optional] 
**CompanyId** | **string** | References companies.company_id — A Spartera seller or buyer company account. See GET /companies for valid values. Required. | 
**ConnectionId** | Pointer to **string** | Optional. | [optional] 
**LlmConnectionId** | Pointer to **string** | References connections.connection_id — Secure connections from Spartera to your databases and data warehouses. See GET /connections for valid values. Optional. | [optional] 
**SnippetId** | Pointer to **int64** | References snippets.snippet_id — Predefined code snippets to accelerate insight creation. See GET /snippets for valid values. Optional. | [optional] 
**IndustryId** | Pointer to **int64** | References industries.industry_id — Available industry categories for asset classification. Based on US NAISC codes.. See GET /industries for valid values. Optional. | [optional] 
**AiJobId** | Pointer to **string** | Links to the AutoInsights job that created this asset | [optional] 
**AucId** | Pointer to **int64** | Primary use case for this asset, from clustering analysis | [optional] 
**FunctionId** | Pointer to **string** | Optional identifier for routing to specific functions/models at seller endpoint. For GET: appended to URL path. For POST: included in JSON body. | [optional] 
**ApprovalStatus** | Pointer to **string** | Approval status for AI-generated assets | [optional] 
**ApprovedByUserId** | Pointer to **string** | User who approved this asset for marketplace | [optional] 
**ApprovedAt** | Pointer to **time.Time** | When this asset was approved for marketplace | [optional] 
**Name** | **string** | Required. | 
**Slug** | Pointer to **string** | Optional. | [optional] 
**Description** | Pointer to **string** | Optional. | [optional] 
**DetailedDescription** | Pointer to **string** | Long-form HTML description for product pages and SEO | [optional] 
**Source** | **string** | Required. One of: MANUAL, AUTOMATIC. | 
**AssetType** | Pointer to **string** | Optional. One of: CALCULATION, VISUALIZATION, DATA. | [optional] 
**AssetSchema** | Pointer to **map[string]interface{}** | Stores database table schema data including columns, types, and metadata | [optional] 
**Visibility** | Pointer to **string** | Optional. One of: PRIVATE, SHARED. | [optional] 
**Tags** | Pointer to **string** | Optional. | [optional] 
**ShortCode** | Pointer to **string** | Short code for tera.ac URL shortener (e.g., &#39;f78zq1&#39;) | [optional] 
**RestrictedDomains** | Pointer to **string** | Semicolon or comma-separated list of domains restricted from accessing this asset | [optional] 
**SqlLogic** | Pointer to **string** | Optional. | [optional] 
**SourceSchemaName** | Pointer to **string** | Optional. | [optional] 
**SourceTableName** | Pointer to **string** | Optional. | [optional] 
**SellInMarketplace** | **bool** | Required. | 
**RequireCustomization** | **bool** | Whether this asset requires customization before use | 
**VizChartLibrary** | Pointer to **string** | Optional. One of: PLOTLY, MATPLOTLIB, SEABORN. | [optional] 
**VizChartType** | Pointer to **string** | Optional. One of: LINE, BAR, PIE, DOUGHNUT, POLAR, … (8 total). | [optional] 
**VizDepVarColName** | Pointer to **string** | Optional. | [optional] 
**VizIndepVarColName** | Pointer to **string** | Optional. | [optional] 
**VizSizeColName** | Pointer to **string** | Optional. | [optional] 
**VizColorColName** | Pointer to **string** | Optional. | [optional] 
**VizDataAggregation** | Pointer to **string** | Optional. One of: No Aggregation, Sum, Average, Count, Minimum, … (6 total). | [optional] 
**VizSortDirection** | Pointer to **string** | Optional. One of: No Sorting, Ascending, Descending. | [optional] 
**VizDataLimit** | Pointer to **int64** | Optional. | [optional] 
**VizColorScheme** | Pointer to **string** | Optional. One of: Default, Sequential, Diverging, Categorical, Monochrome, … (8 total). | [optional] 
**VizShowLegend** | Pointer to **bool** | Show/hide chart legend | [optional] 
**VizShowGrid** | Pointer to **bool** | Show/hide grid lines | [optional] 
**VizShowTrendline** | Pointer to **bool** | Show trendline for scatter/line charts | [optional] 
**VizLineSmoothing** | Pointer to **bool** | Enable smoothing for line charts | [optional] 
**VizBarStacked** | Pointer to **bool** | Stack bars instead of grouping | [optional] 
**VizFilterDirection** | Pointer to **string** | Whether data_limit shows TOP or BOTTOM N | [optional] 
**AllowParams** | **bool** | Required. | 
**AcceptTerms** | **bool** | Required. | 
**Cached** | Pointer to **bool** | Optional. | [optional] 
**Schedule** | Pointer to **string** | Optional. | [optional] 
**NextRun** | Pointer to **time.Time** | Optional. | [optional] 
**DataTimePeriodStart** | Pointer to **time.Time** | Start date of the data time period covered | [optional] 
**DataTimePeriodEnd** | Pointer to **time.Time** | End date of the data time period covered | [optional] 
**GeographicCoverageType** | Pointer to **string** | Type of geographic coverage | [optional] 
**GeographicCoverageDetails** | Pointer to **string** | Specific regions/countries covered (e.g., &#39;United States, Canada, Mexico&#39;) | [optional] 
**DataSourceRefreshFrequency** | Pointer to **string** | How often the source data is refreshed | [optional] 
**DataSourceLastRefreshed** | Pointer to **time.Time** | When the source data was last refreshed | [optional] 
**RateLimitNumber** | Pointer to **int64** | Number of requests allowed per period (e.g., 100) | [optional] 
**RateLimitPeriod** | Pointer to **string** | Time period for rate limiting (second, minute, hour, day) | [optional] 
**RateLimitGranularity** | Pointer to **string** | Granularity level for rate limiting (USER, COMPANY, IP) | [optional] 

## Methods

### NewAssets

`func NewAssets(companyId string, name string, source string, sellInMarketplace bool, requireCustomization bool, allowParams bool, acceptTerms bool, ) *Assets`

NewAssets instantiates a new Assets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsWithDefaults

`func NewAssetsWithDefaults() *Assets`

NewAssetsWithDefaults instantiates a new Assets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *Assets) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Assets) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Assets) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Assets) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Assets) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Assets) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Assets) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Assets) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAssetId

`func (o *Assets) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Assets) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Assets) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Assets) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetUserId

`func (o *Assets) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Assets) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Assets) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Assets) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *Assets) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Assets) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Assets) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetConnectionId

`func (o *Assets) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Assets) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Assets) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Assets) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetLlmConnectionId

`func (o *Assets) GetLlmConnectionId() string`

GetLlmConnectionId returns the LlmConnectionId field if non-nil, zero value otherwise.

### GetLlmConnectionIdOk

`func (o *Assets) GetLlmConnectionIdOk() (*string, bool)`

GetLlmConnectionIdOk returns a tuple with the LlmConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmConnectionId

`func (o *Assets) SetLlmConnectionId(v string)`

SetLlmConnectionId sets LlmConnectionId field to given value.

### HasLlmConnectionId

`func (o *Assets) HasLlmConnectionId() bool`

HasLlmConnectionId returns a boolean if a field has been set.

### GetSnippetId

`func (o *Assets) GetSnippetId() int64`

GetSnippetId returns the SnippetId field if non-nil, zero value otherwise.

### GetSnippetIdOk

`func (o *Assets) GetSnippetIdOk() (*int64, bool)`

GetSnippetIdOk returns a tuple with the SnippetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetId

`func (o *Assets) SetSnippetId(v int64)`

SetSnippetId sets SnippetId field to given value.

### HasSnippetId

`func (o *Assets) HasSnippetId() bool`

HasSnippetId returns a boolean if a field has been set.

### GetIndustryId

`func (o *Assets) GetIndustryId() int64`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *Assets) GetIndustryIdOk() (*int64, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *Assets) SetIndustryId(v int64)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *Assets) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetAiJobId

`func (o *Assets) GetAiJobId() string`

GetAiJobId returns the AiJobId field if non-nil, zero value otherwise.

### GetAiJobIdOk

`func (o *Assets) GetAiJobIdOk() (*string, bool)`

GetAiJobIdOk returns a tuple with the AiJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiJobId

`func (o *Assets) SetAiJobId(v string)`

SetAiJobId sets AiJobId field to given value.

### HasAiJobId

`func (o *Assets) HasAiJobId() bool`

HasAiJobId returns a boolean if a field has been set.

### GetAucId

`func (o *Assets) GetAucId() int64`

GetAucId returns the AucId field if non-nil, zero value otherwise.

### GetAucIdOk

`func (o *Assets) GetAucIdOk() (*int64, bool)`

GetAucIdOk returns a tuple with the AucId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAucId

`func (o *Assets) SetAucId(v int64)`

SetAucId sets AucId field to given value.

### HasAucId

`func (o *Assets) HasAucId() bool`

HasAucId returns a boolean if a field has been set.

### GetFunctionId

`func (o *Assets) GetFunctionId() string`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *Assets) GetFunctionIdOk() (*string, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *Assets) SetFunctionId(v string)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *Assets) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### GetApprovalStatus

`func (o *Assets) GetApprovalStatus() string`

GetApprovalStatus returns the ApprovalStatus field if non-nil, zero value otherwise.

### GetApprovalStatusOk

`func (o *Assets) GetApprovalStatusOk() (*string, bool)`

GetApprovalStatusOk returns a tuple with the ApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatus

`func (o *Assets) SetApprovalStatus(v string)`

SetApprovalStatus sets ApprovalStatus field to given value.

### HasApprovalStatus

`func (o *Assets) HasApprovalStatus() bool`

HasApprovalStatus returns a boolean if a field has been set.

### GetApprovedByUserId

`func (o *Assets) GetApprovedByUserId() string`

GetApprovedByUserId returns the ApprovedByUserId field if non-nil, zero value otherwise.

### GetApprovedByUserIdOk

`func (o *Assets) GetApprovedByUserIdOk() (*string, bool)`

GetApprovedByUserIdOk returns a tuple with the ApprovedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedByUserId

`func (o *Assets) SetApprovedByUserId(v string)`

SetApprovedByUserId sets ApprovedByUserId field to given value.

### HasApprovedByUserId

`func (o *Assets) HasApprovedByUserId() bool`

HasApprovedByUserId returns a boolean if a field has been set.

### GetApprovedAt

`func (o *Assets) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *Assets) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *Assets) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *Assets) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetName

`func (o *Assets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Assets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Assets) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Assets) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Assets) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Assets) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Assets) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *Assets) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Assets) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Assets) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Assets) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetailedDescription

`func (o *Assets) GetDetailedDescription() string`

GetDetailedDescription returns the DetailedDescription field if non-nil, zero value otherwise.

### GetDetailedDescriptionOk

`func (o *Assets) GetDetailedDescriptionOk() (*string, bool)`

GetDetailedDescriptionOk returns a tuple with the DetailedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedDescription

`func (o *Assets) SetDetailedDescription(v string)`

SetDetailedDescription sets DetailedDescription field to given value.

### HasDetailedDescription

`func (o *Assets) HasDetailedDescription() bool`

HasDetailedDescription returns a boolean if a field has been set.

### GetSource

`func (o *Assets) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Assets) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Assets) SetSource(v string)`

SetSource sets Source field to given value.


### GetAssetType

`func (o *Assets) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *Assets) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *Assets) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *Assets) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetAssetSchema

`func (o *Assets) GetAssetSchema() map[string]interface{}`

GetAssetSchema returns the AssetSchema field if non-nil, zero value otherwise.

### GetAssetSchemaOk

`func (o *Assets) GetAssetSchemaOk() (*map[string]interface{}, bool)`

GetAssetSchemaOk returns a tuple with the AssetSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSchema

`func (o *Assets) SetAssetSchema(v map[string]interface{})`

SetAssetSchema sets AssetSchema field to given value.

### HasAssetSchema

`func (o *Assets) HasAssetSchema() bool`

HasAssetSchema returns a boolean if a field has been set.

### GetVisibility

`func (o *Assets) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Assets) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Assets) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Assets) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTags

`func (o *Assets) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Assets) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Assets) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Assets) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetShortCode

`func (o *Assets) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *Assets) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *Assets) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.

### HasShortCode

`func (o *Assets) HasShortCode() bool`

HasShortCode returns a boolean if a field has been set.

### GetRestrictedDomains

`func (o *Assets) GetRestrictedDomains() string`

GetRestrictedDomains returns the RestrictedDomains field if non-nil, zero value otherwise.

### GetRestrictedDomainsOk

`func (o *Assets) GetRestrictedDomainsOk() (*string, bool)`

GetRestrictedDomainsOk returns a tuple with the RestrictedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDomains

`func (o *Assets) SetRestrictedDomains(v string)`

SetRestrictedDomains sets RestrictedDomains field to given value.

### HasRestrictedDomains

`func (o *Assets) HasRestrictedDomains() bool`

HasRestrictedDomains returns a boolean if a field has been set.

### GetSqlLogic

`func (o *Assets) GetSqlLogic() string`

GetSqlLogic returns the SqlLogic field if non-nil, zero value otherwise.

### GetSqlLogicOk

`func (o *Assets) GetSqlLogicOk() (*string, bool)`

GetSqlLogicOk returns a tuple with the SqlLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlLogic

`func (o *Assets) SetSqlLogic(v string)`

SetSqlLogic sets SqlLogic field to given value.

### HasSqlLogic

`func (o *Assets) HasSqlLogic() bool`

HasSqlLogic returns a boolean if a field has been set.

### GetSourceSchemaName

`func (o *Assets) GetSourceSchemaName() string`

GetSourceSchemaName returns the SourceSchemaName field if non-nil, zero value otherwise.

### GetSourceSchemaNameOk

`func (o *Assets) GetSourceSchemaNameOk() (*string, bool)`

GetSourceSchemaNameOk returns a tuple with the SourceSchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchemaName

`func (o *Assets) SetSourceSchemaName(v string)`

SetSourceSchemaName sets SourceSchemaName field to given value.

### HasSourceSchemaName

`func (o *Assets) HasSourceSchemaName() bool`

HasSourceSchemaName returns a boolean if a field has been set.

### GetSourceTableName

`func (o *Assets) GetSourceTableName() string`

GetSourceTableName returns the SourceTableName field if non-nil, zero value otherwise.

### GetSourceTableNameOk

`func (o *Assets) GetSourceTableNameOk() (*string, bool)`

GetSourceTableNameOk returns a tuple with the SourceTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTableName

`func (o *Assets) SetSourceTableName(v string)`

SetSourceTableName sets SourceTableName field to given value.

### HasSourceTableName

`func (o *Assets) HasSourceTableName() bool`

HasSourceTableName returns a boolean if a field has been set.

### GetSellInMarketplace

`func (o *Assets) GetSellInMarketplace() bool`

GetSellInMarketplace returns the SellInMarketplace field if non-nil, zero value otherwise.

### GetSellInMarketplaceOk

`func (o *Assets) GetSellInMarketplaceOk() (*bool, bool)`

GetSellInMarketplaceOk returns a tuple with the SellInMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellInMarketplace

`func (o *Assets) SetSellInMarketplace(v bool)`

SetSellInMarketplace sets SellInMarketplace field to given value.


### GetRequireCustomization

`func (o *Assets) GetRequireCustomization() bool`

GetRequireCustomization returns the RequireCustomization field if non-nil, zero value otherwise.

### GetRequireCustomizationOk

`func (o *Assets) GetRequireCustomizationOk() (*bool, bool)`

GetRequireCustomizationOk returns a tuple with the RequireCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCustomization

`func (o *Assets) SetRequireCustomization(v bool)`

SetRequireCustomization sets RequireCustomization field to given value.


### GetVizChartLibrary

`func (o *Assets) GetVizChartLibrary() string`

GetVizChartLibrary returns the VizChartLibrary field if non-nil, zero value otherwise.

### GetVizChartLibraryOk

`func (o *Assets) GetVizChartLibraryOk() (*string, bool)`

GetVizChartLibraryOk returns a tuple with the VizChartLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizChartLibrary

`func (o *Assets) SetVizChartLibrary(v string)`

SetVizChartLibrary sets VizChartLibrary field to given value.

### HasVizChartLibrary

`func (o *Assets) HasVizChartLibrary() bool`

HasVizChartLibrary returns a boolean if a field has been set.

### GetVizChartType

`func (o *Assets) GetVizChartType() string`

GetVizChartType returns the VizChartType field if non-nil, zero value otherwise.

### GetVizChartTypeOk

`func (o *Assets) GetVizChartTypeOk() (*string, bool)`

GetVizChartTypeOk returns a tuple with the VizChartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizChartType

`func (o *Assets) SetVizChartType(v string)`

SetVizChartType sets VizChartType field to given value.

### HasVizChartType

`func (o *Assets) HasVizChartType() bool`

HasVizChartType returns a boolean if a field has been set.

### GetVizDepVarColName

`func (o *Assets) GetVizDepVarColName() string`

GetVizDepVarColName returns the VizDepVarColName field if non-nil, zero value otherwise.

### GetVizDepVarColNameOk

`func (o *Assets) GetVizDepVarColNameOk() (*string, bool)`

GetVizDepVarColNameOk returns a tuple with the VizDepVarColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizDepVarColName

`func (o *Assets) SetVizDepVarColName(v string)`

SetVizDepVarColName sets VizDepVarColName field to given value.

### HasVizDepVarColName

`func (o *Assets) HasVizDepVarColName() bool`

HasVizDepVarColName returns a boolean if a field has been set.

### GetVizIndepVarColName

`func (o *Assets) GetVizIndepVarColName() string`

GetVizIndepVarColName returns the VizIndepVarColName field if non-nil, zero value otherwise.

### GetVizIndepVarColNameOk

`func (o *Assets) GetVizIndepVarColNameOk() (*string, bool)`

GetVizIndepVarColNameOk returns a tuple with the VizIndepVarColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizIndepVarColName

`func (o *Assets) SetVizIndepVarColName(v string)`

SetVizIndepVarColName sets VizIndepVarColName field to given value.

### HasVizIndepVarColName

`func (o *Assets) HasVizIndepVarColName() bool`

HasVizIndepVarColName returns a boolean if a field has been set.

### GetVizSizeColName

`func (o *Assets) GetVizSizeColName() string`

GetVizSizeColName returns the VizSizeColName field if non-nil, zero value otherwise.

### GetVizSizeColNameOk

`func (o *Assets) GetVizSizeColNameOk() (*string, bool)`

GetVizSizeColNameOk returns a tuple with the VizSizeColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizSizeColName

`func (o *Assets) SetVizSizeColName(v string)`

SetVizSizeColName sets VizSizeColName field to given value.

### HasVizSizeColName

`func (o *Assets) HasVizSizeColName() bool`

HasVizSizeColName returns a boolean if a field has been set.

### GetVizColorColName

`func (o *Assets) GetVizColorColName() string`

GetVizColorColName returns the VizColorColName field if non-nil, zero value otherwise.

### GetVizColorColNameOk

`func (o *Assets) GetVizColorColNameOk() (*string, bool)`

GetVizColorColNameOk returns a tuple with the VizColorColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizColorColName

`func (o *Assets) SetVizColorColName(v string)`

SetVizColorColName sets VizColorColName field to given value.

### HasVizColorColName

`func (o *Assets) HasVizColorColName() bool`

HasVizColorColName returns a boolean if a field has been set.

### GetVizDataAggregation

`func (o *Assets) GetVizDataAggregation() string`

GetVizDataAggregation returns the VizDataAggregation field if non-nil, zero value otherwise.

### GetVizDataAggregationOk

`func (o *Assets) GetVizDataAggregationOk() (*string, bool)`

GetVizDataAggregationOk returns a tuple with the VizDataAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizDataAggregation

`func (o *Assets) SetVizDataAggregation(v string)`

SetVizDataAggregation sets VizDataAggregation field to given value.

### HasVizDataAggregation

`func (o *Assets) HasVizDataAggregation() bool`

HasVizDataAggregation returns a boolean if a field has been set.

### GetVizSortDirection

`func (o *Assets) GetVizSortDirection() string`

GetVizSortDirection returns the VizSortDirection field if non-nil, zero value otherwise.

### GetVizSortDirectionOk

`func (o *Assets) GetVizSortDirectionOk() (*string, bool)`

GetVizSortDirectionOk returns a tuple with the VizSortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizSortDirection

`func (o *Assets) SetVizSortDirection(v string)`

SetVizSortDirection sets VizSortDirection field to given value.

### HasVizSortDirection

`func (o *Assets) HasVizSortDirection() bool`

HasVizSortDirection returns a boolean if a field has been set.

### GetVizDataLimit

`func (o *Assets) GetVizDataLimit() int64`

GetVizDataLimit returns the VizDataLimit field if non-nil, zero value otherwise.

### GetVizDataLimitOk

`func (o *Assets) GetVizDataLimitOk() (*int64, bool)`

GetVizDataLimitOk returns a tuple with the VizDataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizDataLimit

`func (o *Assets) SetVizDataLimit(v int64)`

SetVizDataLimit sets VizDataLimit field to given value.

### HasVizDataLimit

`func (o *Assets) HasVizDataLimit() bool`

HasVizDataLimit returns a boolean if a field has been set.

### GetVizColorScheme

`func (o *Assets) GetVizColorScheme() string`

GetVizColorScheme returns the VizColorScheme field if non-nil, zero value otherwise.

### GetVizColorSchemeOk

`func (o *Assets) GetVizColorSchemeOk() (*string, bool)`

GetVizColorSchemeOk returns a tuple with the VizColorScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizColorScheme

`func (o *Assets) SetVizColorScheme(v string)`

SetVizColorScheme sets VizColorScheme field to given value.

### HasVizColorScheme

`func (o *Assets) HasVizColorScheme() bool`

HasVizColorScheme returns a boolean if a field has been set.

### GetVizShowLegend

`func (o *Assets) GetVizShowLegend() bool`

GetVizShowLegend returns the VizShowLegend field if non-nil, zero value otherwise.

### GetVizShowLegendOk

`func (o *Assets) GetVizShowLegendOk() (*bool, bool)`

GetVizShowLegendOk returns a tuple with the VizShowLegend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizShowLegend

`func (o *Assets) SetVizShowLegend(v bool)`

SetVizShowLegend sets VizShowLegend field to given value.

### HasVizShowLegend

`func (o *Assets) HasVizShowLegend() bool`

HasVizShowLegend returns a boolean if a field has been set.

### GetVizShowGrid

`func (o *Assets) GetVizShowGrid() bool`

GetVizShowGrid returns the VizShowGrid field if non-nil, zero value otherwise.

### GetVizShowGridOk

`func (o *Assets) GetVizShowGridOk() (*bool, bool)`

GetVizShowGridOk returns a tuple with the VizShowGrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizShowGrid

`func (o *Assets) SetVizShowGrid(v bool)`

SetVizShowGrid sets VizShowGrid field to given value.

### HasVizShowGrid

`func (o *Assets) HasVizShowGrid() bool`

HasVizShowGrid returns a boolean if a field has been set.

### GetVizShowTrendline

`func (o *Assets) GetVizShowTrendline() bool`

GetVizShowTrendline returns the VizShowTrendline field if non-nil, zero value otherwise.

### GetVizShowTrendlineOk

`func (o *Assets) GetVizShowTrendlineOk() (*bool, bool)`

GetVizShowTrendlineOk returns a tuple with the VizShowTrendline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizShowTrendline

`func (o *Assets) SetVizShowTrendline(v bool)`

SetVizShowTrendline sets VizShowTrendline field to given value.

### HasVizShowTrendline

`func (o *Assets) HasVizShowTrendline() bool`

HasVizShowTrendline returns a boolean if a field has been set.

### GetVizLineSmoothing

`func (o *Assets) GetVizLineSmoothing() bool`

GetVizLineSmoothing returns the VizLineSmoothing field if non-nil, zero value otherwise.

### GetVizLineSmoothingOk

`func (o *Assets) GetVizLineSmoothingOk() (*bool, bool)`

GetVizLineSmoothingOk returns a tuple with the VizLineSmoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizLineSmoothing

`func (o *Assets) SetVizLineSmoothing(v bool)`

SetVizLineSmoothing sets VizLineSmoothing field to given value.

### HasVizLineSmoothing

`func (o *Assets) HasVizLineSmoothing() bool`

HasVizLineSmoothing returns a boolean if a field has been set.

### GetVizBarStacked

`func (o *Assets) GetVizBarStacked() bool`

GetVizBarStacked returns the VizBarStacked field if non-nil, zero value otherwise.

### GetVizBarStackedOk

`func (o *Assets) GetVizBarStackedOk() (*bool, bool)`

GetVizBarStackedOk returns a tuple with the VizBarStacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizBarStacked

`func (o *Assets) SetVizBarStacked(v bool)`

SetVizBarStacked sets VizBarStacked field to given value.

### HasVizBarStacked

`func (o *Assets) HasVizBarStacked() bool`

HasVizBarStacked returns a boolean if a field has been set.

### GetVizFilterDirection

`func (o *Assets) GetVizFilterDirection() string`

GetVizFilterDirection returns the VizFilterDirection field if non-nil, zero value otherwise.

### GetVizFilterDirectionOk

`func (o *Assets) GetVizFilterDirectionOk() (*string, bool)`

GetVizFilterDirectionOk returns a tuple with the VizFilterDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizFilterDirection

`func (o *Assets) SetVizFilterDirection(v string)`

SetVizFilterDirection sets VizFilterDirection field to given value.

### HasVizFilterDirection

`func (o *Assets) HasVizFilterDirection() bool`

HasVizFilterDirection returns a boolean if a field has been set.

### GetAllowParams

`func (o *Assets) GetAllowParams() bool`

GetAllowParams returns the AllowParams field if non-nil, zero value otherwise.

### GetAllowParamsOk

`func (o *Assets) GetAllowParamsOk() (*bool, bool)`

GetAllowParamsOk returns a tuple with the AllowParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowParams

`func (o *Assets) SetAllowParams(v bool)`

SetAllowParams sets AllowParams field to given value.


### GetAcceptTerms

`func (o *Assets) GetAcceptTerms() bool`

GetAcceptTerms returns the AcceptTerms field if non-nil, zero value otherwise.

### GetAcceptTermsOk

`func (o *Assets) GetAcceptTermsOk() (*bool, bool)`

GetAcceptTermsOk returns a tuple with the AcceptTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTerms

`func (o *Assets) SetAcceptTerms(v bool)`

SetAcceptTerms sets AcceptTerms field to given value.


### GetCached

`func (o *Assets) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *Assets) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *Assets) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *Assets) HasCached() bool`

HasCached returns a boolean if a field has been set.

### GetSchedule

`func (o *Assets) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Assets) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Assets) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Assets) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetNextRun

`func (o *Assets) GetNextRun() time.Time`

GetNextRun returns the NextRun field if non-nil, zero value otherwise.

### GetNextRunOk

`func (o *Assets) GetNextRunOk() (*time.Time, bool)`

GetNextRunOk returns a tuple with the NextRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRun

`func (o *Assets) SetNextRun(v time.Time)`

SetNextRun sets NextRun field to given value.

### HasNextRun

`func (o *Assets) HasNextRun() bool`

HasNextRun returns a boolean if a field has been set.

### GetDataTimePeriodStart

`func (o *Assets) GetDataTimePeriodStart() time.Time`

GetDataTimePeriodStart returns the DataTimePeriodStart field if non-nil, zero value otherwise.

### GetDataTimePeriodStartOk

`func (o *Assets) GetDataTimePeriodStartOk() (*time.Time, bool)`

GetDataTimePeriodStartOk returns a tuple with the DataTimePeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimePeriodStart

`func (o *Assets) SetDataTimePeriodStart(v time.Time)`

SetDataTimePeriodStart sets DataTimePeriodStart field to given value.

### HasDataTimePeriodStart

`func (o *Assets) HasDataTimePeriodStart() bool`

HasDataTimePeriodStart returns a boolean if a field has been set.

### GetDataTimePeriodEnd

`func (o *Assets) GetDataTimePeriodEnd() time.Time`

GetDataTimePeriodEnd returns the DataTimePeriodEnd field if non-nil, zero value otherwise.

### GetDataTimePeriodEndOk

`func (o *Assets) GetDataTimePeriodEndOk() (*time.Time, bool)`

GetDataTimePeriodEndOk returns a tuple with the DataTimePeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimePeriodEnd

`func (o *Assets) SetDataTimePeriodEnd(v time.Time)`

SetDataTimePeriodEnd sets DataTimePeriodEnd field to given value.

### HasDataTimePeriodEnd

`func (o *Assets) HasDataTimePeriodEnd() bool`

HasDataTimePeriodEnd returns a boolean if a field has been set.

### GetGeographicCoverageType

`func (o *Assets) GetGeographicCoverageType() string`

GetGeographicCoverageType returns the GeographicCoverageType field if non-nil, zero value otherwise.

### GetGeographicCoverageTypeOk

`func (o *Assets) GetGeographicCoverageTypeOk() (*string, bool)`

GetGeographicCoverageTypeOk returns a tuple with the GeographicCoverageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicCoverageType

`func (o *Assets) SetGeographicCoverageType(v string)`

SetGeographicCoverageType sets GeographicCoverageType field to given value.

### HasGeographicCoverageType

`func (o *Assets) HasGeographicCoverageType() bool`

HasGeographicCoverageType returns a boolean if a field has been set.

### GetGeographicCoverageDetails

`func (o *Assets) GetGeographicCoverageDetails() string`

GetGeographicCoverageDetails returns the GeographicCoverageDetails field if non-nil, zero value otherwise.

### GetGeographicCoverageDetailsOk

`func (o *Assets) GetGeographicCoverageDetailsOk() (*string, bool)`

GetGeographicCoverageDetailsOk returns a tuple with the GeographicCoverageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicCoverageDetails

`func (o *Assets) SetGeographicCoverageDetails(v string)`

SetGeographicCoverageDetails sets GeographicCoverageDetails field to given value.

### HasGeographicCoverageDetails

`func (o *Assets) HasGeographicCoverageDetails() bool`

HasGeographicCoverageDetails returns a boolean if a field has been set.

### GetDataSourceRefreshFrequency

`func (o *Assets) GetDataSourceRefreshFrequency() string`

GetDataSourceRefreshFrequency returns the DataSourceRefreshFrequency field if non-nil, zero value otherwise.

### GetDataSourceRefreshFrequencyOk

`func (o *Assets) GetDataSourceRefreshFrequencyOk() (*string, bool)`

GetDataSourceRefreshFrequencyOk returns a tuple with the DataSourceRefreshFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceRefreshFrequency

`func (o *Assets) SetDataSourceRefreshFrequency(v string)`

SetDataSourceRefreshFrequency sets DataSourceRefreshFrequency field to given value.

### HasDataSourceRefreshFrequency

`func (o *Assets) HasDataSourceRefreshFrequency() bool`

HasDataSourceRefreshFrequency returns a boolean if a field has been set.

### GetDataSourceLastRefreshed

`func (o *Assets) GetDataSourceLastRefreshed() time.Time`

GetDataSourceLastRefreshed returns the DataSourceLastRefreshed field if non-nil, zero value otherwise.

### GetDataSourceLastRefreshedOk

`func (o *Assets) GetDataSourceLastRefreshedOk() (*time.Time, bool)`

GetDataSourceLastRefreshedOk returns a tuple with the DataSourceLastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceLastRefreshed

`func (o *Assets) SetDataSourceLastRefreshed(v time.Time)`

SetDataSourceLastRefreshed sets DataSourceLastRefreshed field to given value.

### HasDataSourceLastRefreshed

`func (o *Assets) HasDataSourceLastRefreshed() bool`

HasDataSourceLastRefreshed returns a boolean if a field has been set.

### GetRateLimitNumber

`func (o *Assets) GetRateLimitNumber() int64`

GetRateLimitNumber returns the RateLimitNumber field if non-nil, zero value otherwise.

### GetRateLimitNumberOk

`func (o *Assets) GetRateLimitNumberOk() (*int64, bool)`

GetRateLimitNumberOk returns a tuple with the RateLimitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitNumber

`func (o *Assets) SetRateLimitNumber(v int64)`

SetRateLimitNumber sets RateLimitNumber field to given value.

### HasRateLimitNumber

`func (o *Assets) HasRateLimitNumber() bool`

HasRateLimitNumber returns a boolean if a field has been set.

### GetRateLimitPeriod

`func (o *Assets) GetRateLimitPeriod() string`

GetRateLimitPeriod returns the RateLimitPeriod field if non-nil, zero value otherwise.

### GetRateLimitPeriodOk

`func (o *Assets) GetRateLimitPeriodOk() (*string, bool)`

GetRateLimitPeriodOk returns a tuple with the RateLimitPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPeriod

`func (o *Assets) SetRateLimitPeriod(v string)`

SetRateLimitPeriod sets RateLimitPeriod field to given value.

### HasRateLimitPeriod

`func (o *Assets) HasRateLimitPeriod() bool`

HasRateLimitPeriod returns a boolean if a field has been set.

### GetRateLimitGranularity

`func (o *Assets) GetRateLimitGranularity() string`

GetRateLimitGranularity returns the RateLimitGranularity field if non-nil, zero value otherwise.

### GetRateLimitGranularityOk

`func (o *Assets) GetRateLimitGranularityOk() (*string, bool)`

GetRateLimitGranularityOk returns a tuple with the RateLimitGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitGranularity

`func (o *Assets) SetRateLimitGranularity(v string)`

SetRateLimitGranularity sets RateLimitGranularity field to given value.

### HasRateLimitGranularity

`func (o *Assets) HasRateLimitGranularity() bool`

HasRateLimitGranularity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



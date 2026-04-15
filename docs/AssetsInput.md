# AssetsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | References users.user_id — An individual user account within a company. See GET /users for valid values. Optional. | [optional] 
**CompanyId** | **string** | References companies.company_id — A Spartera seller or buyer company account. See GET /companies for valid values. Required. | 
**ConnectionId** | Pointer to **string** | Optional. | [optional] 
**IndustryId** | Pointer to **int64** | References industries.industry_id — Available industry categories for asset classification. Based on US NAISC codes.. See GET /industries for valid values. Optional. | [optional] 
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
**Tags** | Pointer to **string** | Optional. | [optional] 
**ShortCode** | Pointer to **string** | Short code for tera.ac URL shortener (e.g., &#39;f78zq1&#39;) | [optional] 
**RestrictedDomains** | Pointer to **string** | Semicolon or comma-separated list of domains restricted from accessing this asset | [optional] 
**SqlLogic** | Pointer to **string** | Optional. | [optional] 
**SourceSchemaName** | Pointer to **string** | Optional. | [optional] 
**SourceTableName** | Pointer to **string** | Optional. | [optional] 
**SellInMarketplace** | Pointer to **bool** | Required. | [optional] 
**RequireCustomization** | Pointer to **bool** | Whether this asset requires customization before use | [optional] 
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
**AllowParams** | Pointer to **bool** | Required. | [optional] 
**AcceptTerms** | Pointer to **bool** | Required. | [optional] 
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

### NewAssetsInput

`func NewAssetsInput(companyId string, name string, source string, ) *AssetsInput`

NewAssetsInput instantiates a new AssetsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsInputWithDefaults

`func NewAssetsInputWithDefaults() *AssetsInput`

NewAssetsInputWithDefaults instantiates a new AssetsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AssetsInput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AssetsInput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AssetsInput) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AssetsInput) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *AssetsInput) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *AssetsInput) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *AssetsInput) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetConnectionId

`func (o *AssetsInput) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *AssetsInput) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *AssetsInput) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *AssetsInput) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetIndustryId

`func (o *AssetsInput) GetIndustryId() int64`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *AssetsInput) GetIndustryIdOk() (*int64, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *AssetsInput) SetIndustryId(v int64)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *AssetsInput) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetAucId

`func (o *AssetsInput) GetAucId() int64`

GetAucId returns the AucId field if non-nil, zero value otherwise.

### GetAucIdOk

`func (o *AssetsInput) GetAucIdOk() (*int64, bool)`

GetAucIdOk returns a tuple with the AucId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAucId

`func (o *AssetsInput) SetAucId(v int64)`

SetAucId sets AucId field to given value.

### HasAucId

`func (o *AssetsInput) HasAucId() bool`

HasAucId returns a boolean if a field has been set.

### GetFunctionId

`func (o *AssetsInput) GetFunctionId() string`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *AssetsInput) GetFunctionIdOk() (*string, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *AssetsInput) SetFunctionId(v string)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *AssetsInput) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### GetApprovalStatus

`func (o *AssetsInput) GetApprovalStatus() string`

GetApprovalStatus returns the ApprovalStatus field if non-nil, zero value otherwise.

### GetApprovalStatusOk

`func (o *AssetsInput) GetApprovalStatusOk() (*string, bool)`

GetApprovalStatusOk returns a tuple with the ApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatus

`func (o *AssetsInput) SetApprovalStatus(v string)`

SetApprovalStatus sets ApprovalStatus field to given value.

### HasApprovalStatus

`func (o *AssetsInput) HasApprovalStatus() bool`

HasApprovalStatus returns a boolean if a field has been set.

### GetApprovedByUserId

`func (o *AssetsInput) GetApprovedByUserId() string`

GetApprovedByUserId returns the ApprovedByUserId field if non-nil, zero value otherwise.

### GetApprovedByUserIdOk

`func (o *AssetsInput) GetApprovedByUserIdOk() (*string, bool)`

GetApprovedByUserIdOk returns a tuple with the ApprovedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedByUserId

`func (o *AssetsInput) SetApprovedByUserId(v string)`

SetApprovedByUserId sets ApprovedByUserId field to given value.

### HasApprovedByUserId

`func (o *AssetsInput) HasApprovedByUserId() bool`

HasApprovedByUserId returns a boolean if a field has been set.

### GetApprovedAt

`func (o *AssetsInput) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *AssetsInput) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *AssetsInput) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *AssetsInput) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetName

`func (o *AssetsInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetsInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetsInput) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *AssetsInput) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AssetsInput) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AssetsInput) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AssetsInput) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *AssetsInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetsInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetsInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetsInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetailedDescription

`func (o *AssetsInput) GetDetailedDescription() string`

GetDetailedDescription returns the DetailedDescription field if non-nil, zero value otherwise.

### GetDetailedDescriptionOk

`func (o *AssetsInput) GetDetailedDescriptionOk() (*string, bool)`

GetDetailedDescriptionOk returns a tuple with the DetailedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedDescription

`func (o *AssetsInput) SetDetailedDescription(v string)`

SetDetailedDescription sets DetailedDescription field to given value.

### HasDetailedDescription

`func (o *AssetsInput) HasDetailedDescription() bool`

HasDetailedDescription returns a boolean if a field has been set.

### GetSource

`func (o *AssetsInput) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AssetsInput) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AssetsInput) SetSource(v string)`

SetSource sets Source field to given value.


### GetAssetType

`func (o *AssetsInput) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AssetsInput) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AssetsInput) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *AssetsInput) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetAssetSchema

`func (o *AssetsInput) GetAssetSchema() map[string]interface{}`

GetAssetSchema returns the AssetSchema field if non-nil, zero value otherwise.

### GetAssetSchemaOk

`func (o *AssetsInput) GetAssetSchemaOk() (*map[string]interface{}, bool)`

GetAssetSchemaOk returns a tuple with the AssetSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSchema

`func (o *AssetsInput) SetAssetSchema(v map[string]interface{})`

SetAssetSchema sets AssetSchema field to given value.

### HasAssetSchema

`func (o *AssetsInput) HasAssetSchema() bool`

HasAssetSchema returns a boolean if a field has been set.

### GetTags

`func (o *AssetsInput) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssetsInput) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssetsInput) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssetsInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetShortCode

`func (o *AssetsInput) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *AssetsInput) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *AssetsInput) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.

### HasShortCode

`func (o *AssetsInput) HasShortCode() bool`

HasShortCode returns a boolean if a field has been set.

### GetRestrictedDomains

`func (o *AssetsInput) GetRestrictedDomains() string`

GetRestrictedDomains returns the RestrictedDomains field if non-nil, zero value otherwise.

### GetRestrictedDomainsOk

`func (o *AssetsInput) GetRestrictedDomainsOk() (*string, bool)`

GetRestrictedDomainsOk returns a tuple with the RestrictedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDomains

`func (o *AssetsInput) SetRestrictedDomains(v string)`

SetRestrictedDomains sets RestrictedDomains field to given value.

### HasRestrictedDomains

`func (o *AssetsInput) HasRestrictedDomains() bool`

HasRestrictedDomains returns a boolean if a field has been set.

### GetSqlLogic

`func (o *AssetsInput) GetSqlLogic() string`

GetSqlLogic returns the SqlLogic field if non-nil, zero value otherwise.

### GetSqlLogicOk

`func (o *AssetsInput) GetSqlLogicOk() (*string, bool)`

GetSqlLogicOk returns a tuple with the SqlLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlLogic

`func (o *AssetsInput) SetSqlLogic(v string)`

SetSqlLogic sets SqlLogic field to given value.

### HasSqlLogic

`func (o *AssetsInput) HasSqlLogic() bool`

HasSqlLogic returns a boolean if a field has been set.

### GetSourceSchemaName

`func (o *AssetsInput) GetSourceSchemaName() string`

GetSourceSchemaName returns the SourceSchemaName field if non-nil, zero value otherwise.

### GetSourceSchemaNameOk

`func (o *AssetsInput) GetSourceSchemaNameOk() (*string, bool)`

GetSourceSchemaNameOk returns a tuple with the SourceSchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchemaName

`func (o *AssetsInput) SetSourceSchemaName(v string)`

SetSourceSchemaName sets SourceSchemaName field to given value.

### HasSourceSchemaName

`func (o *AssetsInput) HasSourceSchemaName() bool`

HasSourceSchemaName returns a boolean if a field has been set.

### GetSourceTableName

`func (o *AssetsInput) GetSourceTableName() string`

GetSourceTableName returns the SourceTableName field if non-nil, zero value otherwise.

### GetSourceTableNameOk

`func (o *AssetsInput) GetSourceTableNameOk() (*string, bool)`

GetSourceTableNameOk returns a tuple with the SourceTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTableName

`func (o *AssetsInput) SetSourceTableName(v string)`

SetSourceTableName sets SourceTableName field to given value.

### HasSourceTableName

`func (o *AssetsInput) HasSourceTableName() bool`

HasSourceTableName returns a boolean if a field has been set.

### GetSellInMarketplace

`func (o *AssetsInput) GetSellInMarketplace() bool`

GetSellInMarketplace returns the SellInMarketplace field if non-nil, zero value otherwise.

### GetSellInMarketplaceOk

`func (o *AssetsInput) GetSellInMarketplaceOk() (*bool, bool)`

GetSellInMarketplaceOk returns a tuple with the SellInMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellInMarketplace

`func (o *AssetsInput) SetSellInMarketplace(v bool)`

SetSellInMarketplace sets SellInMarketplace field to given value.

### HasSellInMarketplace

`func (o *AssetsInput) HasSellInMarketplace() bool`

HasSellInMarketplace returns a boolean if a field has been set.

### GetRequireCustomization

`func (o *AssetsInput) GetRequireCustomization() bool`

GetRequireCustomization returns the RequireCustomization field if non-nil, zero value otherwise.

### GetRequireCustomizationOk

`func (o *AssetsInput) GetRequireCustomizationOk() (*bool, bool)`

GetRequireCustomizationOk returns a tuple with the RequireCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCustomization

`func (o *AssetsInput) SetRequireCustomization(v bool)`

SetRequireCustomization sets RequireCustomization field to given value.

### HasRequireCustomization

`func (o *AssetsInput) HasRequireCustomization() bool`

HasRequireCustomization returns a boolean if a field has been set.

### GetVizChartLibrary

`func (o *AssetsInput) GetVizChartLibrary() string`

GetVizChartLibrary returns the VizChartLibrary field if non-nil, zero value otherwise.

### GetVizChartLibraryOk

`func (o *AssetsInput) GetVizChartLibraryOk() (*string, bool)`

GetVizChartLibraryOk returns a tuple with the VizChartLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizChartLibrary

`func (o *AssetsInput) SetVizChartLibrary(v string)`

SetVizChartLibrary sets VizChartLibrary field to given value.

### HasVizChartLibrary

`func (o *AssetsInput) HasVizChartLibrary() bool`

HasVizChartLibrary returns a boolean if a field has been set.

### GetVizChartType

`func (o *AssetsInput) GetVizChartType() string`

GetVizChartType returns the VizChartType field if non-nil, zero value otherwise.

### GetVizChartTypeOk

`func (o *AssetsInput) GetVizChartTypeOk() (*string, bool)`

GetVizChartTypeOk returns a tuple with the VizChartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizChartType

`func (o *AssetsInput) SetVizChartType(v string)`

SetVizChartType sets VizChartType field to given value.

### HasVizChartType

`func (o *AssetsInput) HasVizChartType() bool`

HasVizChartType returns a boolean if a field has been set.

### GetVizDepVarColName

`func (o *AssetsInput) GetVizDepVarColName() string`

GetVizDepVarColName returns the VizDepVarColName field if non-nil, zero value otherwise.

### GetVizDepVarColNameOk

`func (o *AssetsInput) GetVizDepVarColNameOk() (*string, bool)`

GetVizDepVarColNameOk returns a tuple with the VizDepVarColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizDepVarColName

`func (o *AssetsInput) SetVizDepVarColName(v string)`

SetVizDepVarColName sets VizDepVarColName field to given value.

### HasVizDepVarColName

`func (o *AssetsInput) HasVizDepVarColName() bool`

HasVizDepVarColName returns a boolean if a field has been set.

### GetVizIndepVarColName

`func (o *AssetsInput) GetVizIndepVarColName() string`

GetVizIndepVarColName returns the VizIndepVarColName field if non-nil, zero value otherwise.

### GetVizIndepVarColNameOk

`func (o *AssetsInput) GetVizIndepVarColNameOk() (*string, bool)`

GetVizIndepVarColNameOk returns a tuple with the VizIndepVarColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizIndepVarColName

`func (o *AssetsInput) SetVizIndepVarColName(v string)`

SetVizIndepVarColName sets VizIndepVarColName field to given value.

### HasVizIndepVarColName

`func (o *AssetsInput) HasVizIndepVarColName() bool`

HasVizIndepVarColName returns a boolean if a field has been set.

### GetVizSizeColName

`func (o *AssetsInput) GetVizSizeColName() string`

GetVizSizeColName returns the VizSizeColName field if non-nil, zero value otherwise.

### GetVizSizeColNameOk

`func (o *AssetsInput) GetVizSizeColNameOk() (*string, bool)`

GetVizSizeColNameOk returns a tuple with the VizSizeColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizSizeColName

`func (o *AssetsInput) SetVizSizeColName(v string)`

SetVizSizeColName sets VizSizeColName field to given value.

### HasVizSizeColName

`func (o *AssetsInput) HasVizSizeColName() bool`

HasVizSizeColName returns a boolean if a field has been set.

### GetVizColorColName

`func (o *AssetsInput) GetVizColorColName() string`

GetVizColorColName returns the VizColorColName field if non-nil, zero value otherwise.

### GetVizColorColNameOk

`func (o *AssetsInput) GetVizColorColNameOk() (*string, bool)`

GetVizColorColNameOk returns a tuple with the VizColorColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizColorColName

`func (o *AssetsInput) SetVizColorColName(v string)`

SetVizColorColName sets VizColorColName field to given value.

### HasVizColorColName

`func (o *AssetsInput) HasVizColorColName() bool`

HasVizColorColName returns a boolean if a field has been set.

### GetVizDataAggregation

`func (o *AssetsInput) GetVizDataAggregation() string`

GetVizDataAggregation returns the VizDataAggregation field if non-nil, zero value otherwise.

### GetVizDataAggregationOk

`func (o *AssetsInput) GetVizDataAggregationOk() (*string, bool)`

GetVizDataAggregationOk returns a tuple with the VizDataAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizDataAggregation

`func (o *AssetsInput) SetVizDataAggregation(v string)`

SetVizDataAggregation sets VizDataAggregation field to given value.

### HasVizDataAggregation

`func (o *AssetsInput) HasVizDataAggregation() bool`

HasVizDataAggregation returns a boolean if a field has been set.

### GetVizSortDirection

`func (o *AssetsInput) GetVizSortDirection() string`

GetVizSortDirection returns the VizSortDirection field if non-nil, zero value otherwise.

### GetVizSortDirectionOk

`func (o *AssetsInput) GetVizSortDirectionOk() (*string, bool)`

GetVizSortDirectionOk returns a tuple with the VizSortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizSortDirection

`func (o *AssetsInput) SetVizSortDirection(v string)`

SetVizSortDirection sets VizSortDirection field to given value.

### HasVizSortDirection

`func (o *AssetsInput) HasVizSortDirection() bool`

HasVizSortDirection returns a boolean if a field has been set.

### GetVizDataLimit

`func (o *AssetsInput) GetVizDataLimit() int64`

GetVizDataLimit returns the VizDataLimit field if non-nil, zero value otherwise.

### GetVizDataLimitOk

`func (o *AssetsInput) GetVizDataLimitOk() (*int64, bool)`

GetVizDataLimitOk returns a tuple with the VizDataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizDataLimit

`func (o *AssetsInput) SetVizDataLimit(v int64)`

SetVizDataLimit sets VizDataLimit field to given value.

### HasVizDataLimit

`func (o *AssetsInput) HasVizDataLimit() bool`

HasVizDataLimit returns a boolean if a field has been set.

### GetVizColorScheme

`func (o *AssetsInput) GetVizColorScheme() string`

GetVizColorScheme returns the VizColorScheme field if non-nil, zero value otherwise.

### GetVizColorSchemeOk

`func (o *AssetsInput) GetVizColorSchemeOk() (*string, bool)`

GetVizColorSchemeOk returns a tuple with the VizColorScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizColorScheme

`func (o *AssetsInput) SetVizColorScheme(v string)`

SetVizColorScheme sets VizColorScheme field to given value.

### HasVizColorScheme

`func (o *AssetsInput) HasVizColorScheme() bool`

HasVizColorScheme returns a boolean if a field has been set.

### GetVizShowLegend

`func (o *AssetsInput) GetVizShowLegend() bool`

GetVizShowLegend returns the VizShowLegend field if non-nil, zero value otherwise.

### GetVizShowLegendOk

`func (o *AssetsInput) GetVizShowLegendOk() (*bool, bool)`

GetVizShowLegendOk returns a tuple with the VizShowLegend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizShowLegend

`func (o *AssetsInput) SetVizShowLegend(v bool)`

SetVizShowLegend sets VizShowLegend field to given value.

### HasVizShowLegend

`func (o *AssetsInput) HasVizShowLegend() bool`

HasVizShowLegend returns a boolean if a field has been set.

### GetVizShowGrid

`func (o *AssetsInput) GetVizShowGrid() bool`

GetVizShowGrid returns the VizShowGrid field if non-nil, zero value otherwise.

### GetVizShowGridOk

`func (o *AssetsInput) GetVizShowGridOk() (*bool, bool)`

GetVizShowGridOk returns a tuple with the VizShowGrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizShowGrid

`func (o *AssetsInput) SetVizShowGrid(v bool)`

SetVizShowGrid sets VizShowGrid field to given value.

### HasVizShowGrid

`func (o *AssetsInput) HasVizShowGrid() bool`

HasVizShowGrid returns a boolean if a field has been set.

### GetVizShowTrendline

`func (o *AssetsInput) GetVizShowTrendline() bool`

GetVizShowTrendline returns the VizShowTrendline field if non-nil, zero value otherwise.

### GetVizShowTrendlineOk

`func (o *AssetsInput) GetVizShowTrendlineOk() (*bool, bool)`

GetVizShowTrendlineOk returns a tuple with the VizShowTrendline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizShowTrendline

`func (o *AssetsInput) SetVizShowTrendline(v bool)`

SetVizShowTrendline sets VizShowTrendline field to given value.

### HasVizShowTrendline

`func (o *AssetsInput) HasVizShowTrendline() bool`

HasVizShowTrendline returns a boolean if a field has been set.

### GetVizLineSmoothing

`func (o *AssetsInput) GetVizLineSmoothing() bool`

GetVizLineSmoothing returns the VizLineSmoothing field if non-nil, zero value otherwise.

### GetVizLineSmoothingOk

`func (o *AssetsInput) GetVizLineSmoothingOk() (*bool, bool)`

GetVizLineSmoothingOk returns a tuple with the VizLineSmoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizLineSmoothing

`func (o *AssetsInput) SetVizLineSmoothing(v bool)`

SetVizLineSmoothing sets VizLineSmoothing field to given value.

### HasVizLineSmoothing

`func (o *AssetsInput) HasVizLineSmoothing() bool`

HasVizLineSmoothing returns a boolean if a field has been set.

### GetVizBarStacked

`func (o *AssetsInput) GetVizBarStacked() bool`

GetVizBarStacked returns the VizBarStacked field if non-nil, zero value otherwise.

### GetVizBarStackedOk

`func (o *AssetsInput) GetVizBarStackedOk() (*bool, bool)`

GetVizBarStackedOk returns a tuple with the VizBarStacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizBarStacked

`func (o *AssetsInput) SetVizBarStacked(v bool)`

SetVizBarStacked sets VizBarStacked field to given value.

### HasVizBarStacked

`func (o *AssetsInput) HasVizBarStacked() bool`

HasVizBarStacked returns a boolean if a field has been set.

### GetVizFilterDirection

`func (o *AssetsInput) GetVizFilterDirection() string`

GetVizFilterDirection returns the VizFilterDirection field if non-nil, zero value otherwise.

### GetVizFilterDirectionOk

`func (o *AssetsInput) GetVizFilterDirectionOk() (*string, bool)`

GetVizFilterDirectionOk returns a tuple with the VizFilterDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizFilterDirection

`func (o *AssetsInput) SetVizFilterDirection(v string)`

SetVizFilterDirection sets VizFilterDirection field to given value.

### HasVizFilterDirection

`func (o *AssetsInput) HasVizFilterDirection() bool`

HasVizFilterDirection returns a boolean if a field has been set.

### GetAllowParams

`func (o *AssetsInput) GetAllowParams() bool`

GetAllowParams returns the AllowParams field if non-nil, zero value otherwise.

### GetAllowParamsOk

`func (o *AssetsInput) GetAllowParamsOk() (*bool, bool)`

GetAllowParamsOk returns a tuple with the AllowParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowParams

`func (o *AssetsInput) SetAllowParams(v bool)`

SetAllowParams sets AllowParams field to given value.

### HasAllowParams

`func (o *AssetsInput) HasAllowParams() bool`

HasAllowParams returns a boolean if a field has been set.

### GetAcceptTerms

`func (o *AssetsInput) GetAcceptTerms() bool`

GetAcceptTerms returns the AcceptTerms field if non-nil, zero value otherwise.

### GetAcceptTermsOk

`func (o *AssetsInput) GetAcceptTermsOk() (*bool, bool)`

GetAcceptTermsOk returns a tuple with the AcceptTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTerms

`func (o *AssetsInput) SetAcceptTerms(v bool)`

SetAcceptTerms sets AcceptTerms field to given value.

### HasAcceptTerms

`func (o *AssetsInput) HasAcceptTerms() bool`

HasAcceptTerms returns a boolean if a field has been set.

### GetCached

`func (o *AssetsInput) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *AssetsInput) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *AssetsInput) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *AssetsInput) HasCached() bool`

HasCached returns a boolean if a field has been set.

### GetSchedule

`func (o *AssetsInput) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AssetsInput) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AssetsInput) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AssetsInput) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetNextRun

`func (o *AssetsInput) GetNextRun() time.Time`

GetNextRun returns the NextRun field if non-nil, zero value otherwise.

### GetNextRunOk

`func (o *AssetsInput) GetNextRunOk() (*time.Time, bool)`

GetNextRunOk returns a tuple with the NextRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRun

`func (o *AssetsInput) SetNextRun(v time.Time)`

SetNextRun sets NextRun field to given value.

### HasNextRun

`func (o *AssetsInput) HasNextRun() bool`

HasNextRun returns a boolean if a field has been set.

### GetDataTimePeriodStart

`func (o *AssetsInput) GetDataTimePeriodStart() time.Time`

GetDataTimePeriodStart returns the DataTimePeriodStart field if non-nil, zero value otherwise.

### GetDataTimePeriodStartOk

`func (o *AssetsInput) GetDataTimePeriodStartOk() (*time.Time, bool)`

GetDataTimePeriodStartOk returns a tuple with the DataTimePeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimePeriodStart

`func (o *AssetsInput) SetDataTimePeriodStart(v time.Time)`

SetDataTimePeriodStart sets DataTimePeriodStart field to given value.

### HasDataTimePeriodStart

`func (o *AssetsInput) HasDataTimePeriodStart() bool`

HasDataTimePeriodStart returns a boolean if a field has been set.

### GetDataTimePeriodEnd

`func (o *AssetsInput) GetDataTimePeriodEnd() time.Time`

GetDataTimePeriodEnd returns the DataTimePeriodEnd field if non-nil, zero value otherwise.

### GetDataTimePeriodEndOk

`func (o *AssetsInput) GetDataTimePeriodEndOk() (*time.Time, bool)`

GetDataTimePeriodEndOk returns a tuple with the DataTimePeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimePeriodEnd

`func (o *AssetsInput) SetDataTimePeriodEnd(v time.Time)`

SetDataTimePeriodEnd sets DataTimePeriodEnd field to given value.

### HasDataTimePeriodEnd

`func (o *AssetsInput) HasDataTimePeriodEnd() bool`

HasDataTimePeriodEnd returns a boolean if a field has been set.

### GetGeographicCoverageType

`func (o *AssetsInput) GetGeographicCoverageType() string`

GetGeographicCoverageType returns the GeographicCoverageType field if non-nil, zero value otherwise.

### GetGeographicCoverageTypeOk

`func (o *AssetsInput) GetGeographicCoverageTypeOk() (*string, bool)`

GetGeographicCoverageTypeOk returns a tuple with the GeographicCoverageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicCoverageType

`func (o *AssetsInput) SetGeographicCoverageType(v string)`

SetGeographicCoverageType sets GeographicCoverageType field to given value.

### HasGeographicCoverageType

`func (o *AssetsInput) HasGeographicCoverageType() bool`

HasGeographicCoverageType returns a boolean if a field has been set.

### GetGeographicCoverageDetails

`func (o *AssetsInput) GetGeographicCoverageDetails() string`

GetGeographicCoverageDetails returns the GeographicCoverageDetails field if non-nil, zero value otherwise.

### GetGeographicCoverageDetailsOk

`func (o *AssetsInput) GetGeographicCoverageDetailsOk() (*string, bool)`

GetGeographicCoverageDetailsOk returns a tuple with the GeographicCoverageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicCoverageDetails

`func (o *AssetsInput) SetGeographicCoverageDetails(v string)`

SetGeographicCoverageDetails sets GeographicCoverageDetails field to given value.

### HasGeographicCoverageDetails

`func (o *AssetsInput) HasGeographicCoverageDetails() bool`

HasGeographicCoverageDetails returns a boolean if a field has been set.

### GetDataSourceRefreshFrequency

`func (o *AssetsInput) GetDataSourceRefreshFrequency() string`

GetDataSourceRefreshFrequency returns the DataSourceRefreshFrequency field if non-nil, zero value otherwise.

### GetDataSourceRefreshFrequencyOk

`func (o *AssetsInput) GetDataSourceRefreshFrequencyOk() (*string, bool)`

GetDataSourceRefreshFrequencyOk returns a tuple with the DataSourceRefreshFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceRefreshFrequency

`func (o *AssetsInput) SetDataSourceRefreshFrequency(v string)`

SetDataSourceRefreshFrequency sets DataSourceRefreshFrequency field to given value.

### HasDataSourceRefreshFrequency

`func (o *AssetsInput) HasDataSourceRefreshFrequency() bool`

HasDataSourceRefreshFrequency returns a boolean if a field has been set.

### GetDataSourceLastRefreshed

`func (o *AssetsInput) GetDataSourceLastRefreshed() time.Time`

GetDataSourceLastRefreshed returns the DataSourceLastRefreshed field if non-nil, zero value otherwise.

### GetDataSourceLastRefreshedOk

`func (o *AssetsInput) GetDataSourceLastRefreshedOk() (*time.Time, bool)`

GetDataSourceLastRefreshedOk returns a tuple with the DataSourceLastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceLastRefreshed

`func (o *AssetsInput) SetDataSourceLastRefreshed(v time.Time)`

SetDataSourceLastRefreshed sets DataSourceLastRefreshed field to given value.

### HasDataSourceLastRefreshed

`func (o *AssetsInput) HasDataSourceLastRefreshed() bool`

HasDataSourceLastRefreshed returns a boolean if a field has been set.

### GetRateLimitNumber

`func (o *AssetsInput) GetRateLimitNumber() int64`

GetRateLimitNumber returns the RateLimitNumber field if non-nil, zero value otherwise.

### GetRateLimitNumberOk

`func (o *AssetsInput) GetRateLimitNumberOk() (*int64, bool)`

GetRateLimitNumberOk returns a tuple with the RateLimitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitNumber

`func (o *AssetsInput) SetRateLimitNumber(v int64)`

SetRateLimitNumber sets RateLimitNumber field to given value.

### HasRateLimitNumber

`func (o *AssetsInput) HasRateLimitNumber() bool`

HasRateLimitNumber returns a boolean if a field has been set.

### GetRateLimitPeriod

`func (o *AssetsInput) GetRateLimitPeriod() string`

GetRateLimitPeriod returns the RateLimitPeriod field if non-nil, zero value otherwise.

### GetRateLimitPeriodOk

`func (o *AssetsInput) GetRateLimitPeriodOk() (*string, bool)`

GetRateLimitPeriodOk returns a tuple with the RateLimitPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPeriod

`func (o *AssetsInput) SetRateLimitPeriod(v string)`

SetRateLimitPeriod sets RateLimitPeriod field to given value.

### HasRateLimitPeriod

`func (o *AssetsInput) HasRateLimitPeriod() bool`

HasRateLimitPeriod returns a boolean if a field has been set.

### GetRateLimitGranularity

`func (o *AssetsInput) GetRateLimitGranularity() string`

GetRateLimitGranularity returns the RateLimitGranularity field if non-nil, zero value otherwise.

### GetRateLimitGranularityOk

`func (o *AssetsInput) GetRateLimitGranularityOk() (*string, bool)`

GetRateLimitGranularityOk returns a tuple with the RateLimitGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitGranularity

`func (o *AssetsInput) SetRateLimitGranularity(v string)`

SetRateLimitGranularity sets RateLimitGranularity field to given value.

### HasRateLimitGranularity

`func (o *AssetsInput) HasRateLimitGranularity() bool`

HasRateLimitGranularity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



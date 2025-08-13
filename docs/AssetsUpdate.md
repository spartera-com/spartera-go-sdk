# AssetsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**IndustryId** | Pointer to **int64** |  | [optional] 
**ApprovalStatus** | Pointer to **string** | Approval status for AI-generated assets | [optional] 
**ApprovedByUserId** | Pointer to **string** | User who approved this asset for marketplace | [optional] 
**ApprovedAt** | Pointer to **time.Time** | When this asset was approved for marketplace | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** | Enum type: Source | [optional] 
**AssetType** | Pointer to **string** | Enum type: AssetType | [optional] 
**AssetSchema** | Pointer to **map[string]interface{}** | Stores database table schema data including columns, types, and metadata | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**SqlLogic** | Pointer to **string** |  | [optional] 
**SourceSchemaName** | Pointer to **string** |  | [optional] 
**SourceTableName** | Pointer to **string** |  | [optional] 
**SellInMarketplace** | Pointer to **bool** |  | [optional] 
**VizChartLibrary** | Pointer to **string** | Enum type: PlottingLibrary | [optional] 
**VizChartType** | Pointer to **string** | Enum type: ChartType | [optional] 
**VizDepVarColName** | Pointer to **string** |  | [optional] 
**VizIndepVarColName** | Pointer to **string** |  | [optional] 
**VizSizeColName** | Pointer to **string** |  | [optional] 
**VizColorColName** | Pointer to **string** |  | [optional] 
**VizDataAggregation** | Pointer to **string** | Enum type: AggregationType | [optional] 
**VizSortDirection** | Pointer to **string** | Enum type: SortDirection | [optional] 
**VizDataLimit** | Pointer to **int64** |  | [optional] 
**VizColorScheme** | Pointer to **string** | Enum type: ColorScheme | [optional] 
**AllowParams** | Pointer to **bool** |  | [optional] 
**AcceptTerms** | Pointer to **bool** |  | [optional] 
**Cached** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**NextRun** | Pointer to **time.Time** |  | [optional] 
**DataTimePeriodStart** | Pointer to **time.Time** | Start date of the data time period covered | [optional] 
**DataTimePeriodEnd** | Pointer to **time.Time** | End date of the data time period covered | [optional] 
**GeographicCoverageType** | Pointer to **string** | Type of geographic coverage (Enum type: GeographicCoverage) | [optional] 
**GeographicCoverageDetails** | Pointer to **string** | Specific regions/countries covered (e.g., &#39;United States, Canada, Mexico&#39;) | [optional] 
**DataSourceRefreshFrequency** | Pointer to **string** | How often the source data is refreshed (Enum type: DataRefreshFrequency) | [optional] 
**DataSourceLastRefreshed** | Pointer to **time.Time** | When the source data was last refreshed | [optional] 

## Methods

### NewAssetsUpdate

`func NewAssetsUpdate() *AssetsUpdate`

NewAssetsUpdate instantiates a new AssetsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsUpdateWithDefaults

`func NewAssetsUpdateWithDefaults() *AssetsUpdate`

NewAssetsUpdateWithDefaults instantiates a new AssetsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AssetsUpdate) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AssetsUpdate) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AssetsUpdate) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AssetsUpdate) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *AssetsUpdate) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *AssetsUpdate) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *AssetsUpdate) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *AssetsUpdate) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetConnectionId

`func (o *AssetsUpdate) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *AssetsUpdate) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *AssetsUpdate) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *AssetsUpdate) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetIndustryId

`func (o *AssetsUpdate) GetIndustryId() int64`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *AssetsUpdate) GetIndustryIdOk() (*int64, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *AssetsUpdate) SetIndustryId(v int64)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *AssetsUpdate) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetApprovalStatus

`func (o *AssetsUpdate) GetApprovalStatus() string`

GetApprovalStatus returns the ApprovalStatus field if non-nil, zero value otherwise.

### GetApprovalStatusOk

`func (o *AssetsUpdate) GetApprovalStatusOk() (*string, bool)`

GetApprovalStatusOk returns a tuple with the ApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatus

`func (o *AssetsUpdate) SetApprovalStatus(v string)`

SetApprovalStatus sets ApprovalStatus field to given value.

### HasApprovalStatus

`func (o *AssetsUpdate) HasApprovalStatus() bool`

HasApprovalStatus returns a boolean if a field has been set.

### GetApprovedByUserId

`func (o *AssetsUpdate) GetApprovedByUserId() string`

GetApprovedByUserId returns the ApprovedByUserId field if non-nil, zero value otherwise.

### GetApprovedByUserIdOk

`func (o *AssetsUpdate) GetApprovedByUserIdOk() (*string, bool)`

GetApprovedByUserIdOk returns a tuple with the ApprovedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedByUserId

`func (o *AssetsUpdate) SetApprovedByUserId(v string)`

SetApprovedByUserId sets ApprovedByUserId field to given value.

### HasApprovedByUserId

`func (o *AssetsUpdate) HasApprovedByUserId() bool`

HasApprovedByUserId returns a boolean if a field has been set.

### GetApprovedAt

`func (o *AssetsUpdate) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *AssetsUpdate) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *AssetsUpdate) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *AssetsUpdate) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetName

`func (o *AssetsUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetsUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetsUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetsUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *AssetsUpdate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AssetsUpdate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AssetsUpdate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AssetsUpdate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *AssetsUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetsUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetsUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetsUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSource

`func (o *AssetsUpdate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AssetsUpdate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AssetsUpdate) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AssetsUpdate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAssetType

`func (o *AssetsUpdate) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AssetsUpdate) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AssetsUpdate) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *AssetsUpdate) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetAssetSchema

`func (o *AssetsUpdate) GetAssetSchema() map[string]interface{}`

GetAssetSchema returns the AssetSchema field if non-nil, zero value otherwise.

### GetAssetSchemaOk

`func (o *AssetsUpdate) GetAssetSchemaOk() (*map[string]interface{}, bool)`

GetAssetSchemaOk returns a tuple with the AssetSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSchema

`func (o *AssetsUpdate) SetAssetSchema(v map[string]interface{})`

SetAssetSchema sets AssetSchema field to given value.

### HasAssetSchema

`func (o *AssetsUpdate) HasAssetSchema() bool`

HasAssetSchema returns a boolean if a field has been set.

### GetTags

`func (o *AssetsUpdate) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssetsUpdate) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssetsUpdate) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssetsUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSqlLogic

`func (o *AssetsUpdate) GetSqlLogic() string`

GetSqlLogic returns the SqlLogic field if non-nil, zero value otherwise.

### GetSqlLogicOk

`func (o *AssetsUpdate) GetSqlLogicOk() (*string, bool)`

GetSqlLogicOk returns a tuple with the SqlLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlLogic

`func (o *AssetsUpdate) SetSqlLogic(v string)`

SetSqlLogic sets SqlLogic field to given value.

### HasSqlLogic

`func (o *AssetsUpdate) HasSqlLogic() bool`

HasSqlLogic returns a boolean if a field has been set.

### GetSourceSchemaName

`func (o *AssetsUpdate) GetSourceSchemaName() string`

GetSourceSchemaName returns the SourceSchemaName field if non-nil, zero value otherwise.

### GetSourceSchemaNameOk

`func (o *AssetsUpdate) GetSourceSchemaNameOk() (*string, bool)`

GetSourceSchemaNameOk returns a tuple with the SourceSchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchemaName

`func (o *AssetsUpdate) SetSourceSchemaName(v string)`

SetSourceSchemaName sets SourceSchemaName field to given value.

### HasSourceSchemaName

`func (o *AssetsUpdate) HasSourceSchemaName() bool`

HasSourceSchemaName returns a boolean if a field has been set.

### GetSourceTableName

`func (o *AssetsUpdate) GetSourceTableName() string`

GetSourceTableName returns the SourceTableName field if non-nil, zero value otherwise.

### GetSourceTableNameOk

`func (o *AssetsUpdate) GetSourceTableNameOk() (*string, bool)`

GetSourceTableNameOk returns a tuple with the SourceTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTableName

`func (o *AssetsUpdate) SetSourceTableName(v string)`

SetSourceTableName sets SourceTableName field to given value.

### HasSourceTableName

`func (o *AssetsUpdate) HasSourceTableName() bool`

HasSourceTableName returns a boolean if a field has been set.

### GetSellInMarketplace

`func (o *AssetsUpdate) GetSellInMarketplace() bool`

GetSellInMarketplace returns the SellInMarketplace field if non-nil, zero value otherwise.

### GetSellInMarketplaceOk

`func (o *AssetsUpdate) GetSellInMarketplaceOk() (*bool, bool)`

GetSellInMarketplaceOk returns a tuple with the SellInMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellInMarketplace

`func (o *AssetsUpdate) SetSellInMarketplace(v bool)`

SetSellInMarketplace sets SellInMarketplace field to given value.

### HasSellInMarketplace

`func (o *AssetsUpdate) HasSellInMarketplace() bool`

HasSellInMarketplace returns a boolean if a field has been set.

### GetVizChartLibrary

`func (o *AssetsUpdate) GetVizChartLibrary() string`

GetVizChartLibrary returns the VizChartLibrary field if non-nil, zero value otherwise.

### GetVizChartLibraryOk

`func (o *AssetsUpdate) GetVizChartLibraryOk() (*string, bool)`

GetVizChartLibraryOk returns a tuple with the VizChartLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizChartLibrary

`func (o *AssetsUpdate) SetVizChartLibrary(v string)`

SetVizChartLibrary sets VizChartLibrary field to given value.

### HasVizChartLibrary

`func (o *AssetsUpdate) HasVizChartLibrary() bool`

HasVizChartLibrary returns a boolean if a field has been set.

### GetVizChartType

`func (o *AssetsUpdate) GetVizChartType() string`

GetVizChartType returns the VizChartType field if non-nil, zero value otherwise.

### GetVizChartTypeOk

`func (o *AssetsUpdate) GetVizChartTypeOk() (*string, bool)`

GetVizChartTypeOk returns a tuple with the VizChartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizChartType

`func (o *AssetsUpdate) SetVizChartType(v string)`

SetVizChartType sets VizChartType field to given value.

### HasVizChartType

`func (o *AssetsUpdate) HasVizChartType() bool`

HasVizChartType returns a boolean if a field has been set.

### GetVizDepVarColName

`func (o *AssetsUpdate) GetVizDepVarColName() string`

GetVizDepVarColName returns the VizDepVarColName field if non-nil, zero value otherwise.

### GetVizDepVarColNameOk

`func (o *AssetsUpdate) GetVizDepVarColNameOk() (*string, bool)`

GetVizDepVarColNameOk returns a tuple with the VizDepVarColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizDepVarColName

`func (o *AssetsUpdate) SetVizDepVarColName(v string)`

SetVizDepVarColName sets VizDepVarColName field to given value.

### HasVizDepVarColName

`func (o *AssetsUpdate) HasVizDepVarColName() bool`

HasVizDepVarColName returns a boolean if a field has been set.

### GetVizIndepVarColName

`func (o *AssetsUpdate) GetVizIndepVarColName() string`

GetVizIndepVarColName returns the VizIndepVarColName field if non-nil, zero value otherwise.

### GetVizIndepVarColNameOk

`func (o *AssetsUpdate) GetVizIndepVarColNameOk() (*string, bool)`

GetVizIndepVarColNameOk returns a tuple with the VizIndepVarColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizIndepVarColName

`func (o *AssetsUpdate) SetVizIndepVarColName(v string)`

SetVizIndepVarColName sets VizIndepVarColName field to given value.

### HasVizIndepVarColName

`func (o *AssetsUpdate) HasVizIndepVarColName() bool`

HasVizIndepVarColName returns a boolean if a field has been set.

### GetVizSizeColName

`func (o *AssetsUpdate) GetVizSizeColName() string`

GetVizSizeColName returns the VizSizeColName field if non-nil, zero value otherwise.

### GetVizSizeColNameOk

`func (o *AssetsUpdate) GetVizSizeColNameOk() (*string, bool)`

GetVizSizeColNameOk returns a tuple with the VizSizeColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizSizeColName

`func (o *AssetsUpdate) SetVizSizeColName(v string)`

SetVizSizeColName sets VizSizeColName field to given value.

### HasVizSizeColName

`func (o *AssetsUpdate) HasVizSizeColName() bool`

HasVizSizeColName returns a boolean if a field has been set.

### GetVizColorColName

`func (o *AssetsUpdate) GetVizColorColName() string`

GetVizColorColName returns the VizColorColName field if non-nil, zero value otherwise.

### GetVizColorColNameOk

`func (o *AssetsUpdate) GetVizColorColNameOk() (*string, bool)`

GetVizColorColNameOk returns a tuple with the VizColorColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizColorColName

`func (o *AssetsUpdate) SetVizColorColName(v string)`

SetVizColorColName sets VizColorColName field to given value.

### HasVizColorColName

`func (o *AssetsUpdate) HasVizColorColName() bool`

HasVizColorColName returns a boolean if a field has been set.

### GetVizDataAggregation

`func (o *AssetsUpdate) GetVizDataAggregation() string`

GetVizDataAggregation returns the VizDataAggregation field if non-nil, zero value otherwise.

### GetVizDataAggregationOk

`func (o *AssetsUpdate) GetVizDataAggregationOk() (*string, bool)`

GetVizDataAggregationOk returns a tuple with the VizDataAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizDataAggregation

`func (o *AssetsUpdate) SetVizDataAggregation(v string)`

SetVizDataAggregation sets VizDataAggregation field to given value.

### HasVizDataAggregation

`func (o *AssetsUpdate) HasVizDataAggregation() bool`

HasVizDataAggregation returns a boolean if a field has been set.

### GetVizSortDirection

`func (o *AssetsUpdate) GetVizSortDirection() string`

GetVizSortDirection returns the VizSortDirection field if non-nil, zero value otherwise.

### GetVizSortDirectionOk

`func (o *AssetsUpdate) GetVizSortDirectionOk() (*string, bool)`

GetVizSortDirectionOk returns a tuple with the VizSortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizSortDirection

`func (o *AssetsUpdate) SetVizSortDirection(v string)`

SetVizSortDirection sets VizSortDirection field to given value.

### HasVizSortDirection

`func (o *AssetsUpdate) HasVizSortDirection() bool`

HasVizSortDirection returns a boolean if a field has been set.

### GetVizDataLimit

`func (o *AssetsUpdate) GetVizDataLimit() int64`

GetVizDataLimit returns the VizDataLimit field if non-nil, zero value otherwise.

### GetVizDataLimitOk

`func (o *AssetsUpdate) GetVizDataLimitOk() (*int64, bool)`

GetVizDataLimitOk returns a tuple with the VizDataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizDataLimit

`func (o *AssetsUpdate) SetVizDataLimit(v int64)`

SetVizDataLimit sets VizDataLimit field to given value.

### HasVizDataLimit

`func (o *AssetsUpdate) HasVizDataLimit() bool`

HasVizDataLimit returns a boolean if a field has been set.

### GetVizColorScheme

`func (o *AssetsUpdate) GetVizColorScheme() string`

GetVizColorScheme returns the VizColorScheme field if non-nil, zero value otherwise.

### GetVizColorSchemeOk

`func (o *AssetsUpdate) GetVizColorSchemeOk() (*string, bool)`

GetVizColorSchemeOk returns a tuple with the VizColorScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizColorScheme

`func (o *AssetsUpdate) SetVizColorScheme(v string)`

SetVizColorScheme sets VizColorScheme field to given value.

### HasVizColorScheme

`func (o *AssetsUpdate) HasVizColorScheme() bool`

HasVizColorScheme returns a boolean if a field has been set.

### GetAllowParams

`func (o *AssetsUpdate) GetAllowParams() bool`

GetAllowParams returns the AllowParams field if non-nil, zero value otherwise.

### GetAllowParamsOk

`func (o *AssetsUpdate) GetAllowParamsOk() (*bool, bool)`

GetAllowParamsOk returns a tuple with the AllowParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowParams

`func (o *AssetsUpdate) SetAllowParams(v bool)`

SetAllowParams sets AllowParams field to given value.

### HasAllowParams

`func (o *AssetsUpdate) HasAllowParams() bool`

HasAllowParams returns a boolean if a field has been set.

### GetAcceptTerms

`func (o *AssetsUpdate) GetAcceptTerms() bool`

GetAcceptTerms returns the AcceptTerms field if non-nil, zero value otherwise.

### GetAcceptTermsOk

`func (o *AssetsUpdate) GetAcceptTermsOk() (*bool, bool)`

GetAcceptTermsOk returns a tuple with the AcceptTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTerms

`func (o *AssetsUpdate) SetAcceptTerms(v bool)`

SetAcceptTerms sets AcceptTerms field to given value.

### HasAcceptTerms

`func (o *AssetsUpdate) HasAcceptTerms() bool`

HasAcceptTerms returns a boolean if a field has been set.

### GetCached

`func (o *AssetsUpdate) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *AssetsUpdate) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *AssetsUpdate) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *AssetsUpdate) HasCached() bool`

HasCached returns a boolean if a field has been set.

### GetSchedule

`func (o *AssetsUpdate) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AssetsUpdate) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AssetsUpdate) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AssetsUpdate) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetNextRun

`func (o *AssetsUpdate) GetNextRun() time.Time`

GetNextRun returns the NextRun field if non-nil, zero value otherwise.

### GetNextRunOk

`func (o *AssetsUpdate) GetNextRunOk() (*time.Time, bool)`

GetNextRunOk returns a tuple with the NextRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRun

`func (o *AssetsUpdate) SetNextRun(v time.Time)`

SetNextRun sets NextRun field to given value.

### HasNextRun

`func (o *AssetsUpdate) HasNextRun() bool`

HasNextRun returns a boolean if a field has been set.

### GetDataTimePeriodStart

`func (o *AssetsUpdate) GetDataTimePeriodStart() time.Time`

GetDataTimePeriodStart returns the DataTimePeriodStart field if non-nil, zero value otherwise.

### GetDataTimePeriodStartOk

`func (o *AssetsUpdate) GetDataTimePeriodStartOk() (*time.Time, bool)`

GetDataTimePeriodStartOk returns a tuple with the DataTimePeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimePeriodStart

`func (o *AssetsUpdate) SetDataTimePeriodStart(v time.Time)`

SetDataTimePeriodStart sets DataTimePeriodStart field to given value.

### HasDataTimePeriodStart

`func (o *AssetsUpdate) HasDataTimePeriodStart() bool`

HasDataTimePeriodStart returns a boolean if a field has been set.

### GetDataTimePeriodEnd

`func (o *AssetsUpdate) GetDataTimePeriodEnd() time.Time`

GetDataTimePeriodEnd returns the DataTimePeriodEnd field if non-nil, zero value otherwise.

### GetDataTimePeriodEndOk

`func (o *AssetsUpdate) GetDataTimePeriodEndOk() (*time.Time, bool)`

GetDataTimePeriodEndOk returns a tuple with the DataTimePeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimePeriodEnd

`func (o *AssetsUpdate) SetDataTimePeriodEnd(v time.Time)`

SetDataTimePeriodEnd sets DataTimePeriodEnd field to given value.

### HasDataTimePeriodEnd

`func (o *AssetsUpdate) HasDataTimePeriodEnd() bool`

HasDataTimePeriodEnd returns a boolean if a field has been set.

### GetGeographicCoverageType

`func (o *AssetsUpdate) GetGeographicCoverageType() string`

GetGeographicCoverageType returns the GeographicCoverageType field if non-nil, zero value otherwise.

### GetGeographicCoverageTypeOk

`func (o *AssetsUpdate) GetGeographicCoverageTypeOk() (*string, bool)`

GetGeographicCoverageTypeOk returns a tuple with the GeographicCoverageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicCoverageType

`func (o *AssetsUpdate) SetGeographicCoverageType(v string)`

SetGeographicCoverageType sets GeographicCoverageType field to given value.

### HasGeographicCoverageType

`func (o *AssetsUpdate) HasGeographicCoverageType() bool`

HasGeographicCoverageType returns a boolean if a field has been set.

### GetGeographicCoverageDetails

`func (o *AssetsUpdate) GetGeographicCoverageDetails() string`

GetGeographicCoverageDetails returns the GeographicCoverageDetails field if non-nil, zero value otherwise.

### GetGeographicCoverageDetailsOk

`func (o *AssetsUpdate) GetGeographicCoverageDetailsOk() (*string, bool)`

GetGeographicCoverageDetailsOk returns a tuple with the GeographicCoverageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicCoverageDetails

`func (o *AssetsUpdate) SetGeographicCoverageDetails(v string)`

SetGeographicCoverageDetails sets GeographicCoverageDetails field to given value.

### HasGeographicCoverageDetails

`func (o *AssetsUpdate) HasGeographicCoverageDetails() bool`

HasGeographicCoverageDetails returns a boolean if a field has been set.

### GetDataSourceRefreshFrequency

`func (o *AssetsUpdate) GetDataSourceRefreshFrequency() string`

GetDataSourceRefreshFrequency returns the DataSourceRefreshFrequency field if non-nil, zero value otherwise.

### GetDataSourceRefreshFrequencyOk

`func (o *AssetsUpdate) GetDataSourceRefreshFrequencyOk() (*string, bool)`

GetDataSourceRefreshFrequencyOk returns a tuple with the DataSourceRefreshFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceRefreshFrequency

`func (o *AssetsUpdate) SetDataSourceRefreshFrequency(v string)`

SetDataSourceRefreshFrequency sets DataSourceRefreshFrequency field to given value.

### HasDataSourceRefreshFrequency

`func (o *AssetsUpdate) HasDataSourceRefreshFrequency() bool`

HasDataSourceRefreshFrequency returns a boolean if a field has been set.

### GetDataSourceLastRefreshed

`func (o *AssetsUpdate) GetDataSourceLastRefreshed() time.Time`

GetDataSourceLastRefreshed returns the DataSourceLastRefreshed field if non-nil, zero value otherwise.

### GetDataSourceLastRefreshedOk

`func (o *AssetsUpdate) GetDataSourceLastRefreshedOk() (*time.Time, bool)`

GetDataSourceLastRefreshedOk returns a tuple with the DataSourceLastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceLastRefreshed

`func (o *AssetsUpdate) SetDataSourceLastRefreshed(v time.Time)`

SetDataSourceLastRefreshed sets DataSourceLastRefreshed field to given value.

### HasDataSourceLastRefreshed

`func (o *AssetsUpdate) HasDataSourceLastRefreshed() bool`

HasDataSourceLastRefreshed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



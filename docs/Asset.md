# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**CompanyId** | **string** |  | 
**ConnectionId** | Pointer to **string** |  | [optional] 
**LlmConnectionId** | Pointer to **string** |  | [optional] 
**SnippetId** | Pointer to **string** |  | [optional] 
**IndustryId** | Pointer to **string** |  | [optional] 
**AiJobId** | Pointer to **string** |  | [optional] 
**ApprovalStatus** | Pointer to **string** |  | [optional] 
**ApprovedByUserId** | Pointer to **string** |  | [optional] 
**ApprovedAt** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Source** | **string** |  | 
**AssetType** | Pointer to **string** |  | [optional] 
**AssetSchema** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**SqlLogic** | Pointer to **string** |  | [optional] 
**SourceSchemaName** | Pointer to **string** |  | [optional] 
**SourceTableName** | Pointer to **string** |  | [optional] 
**SellInMarketplace** | Pointer to **string** |  | [optional] 
**VizChartLibrary** | Pointer to **string** |  | [optional] 
**VizChartType** | Pointer to **string** |  | [optional] 
**VizDepVarColName** | Pointer to **string** |  | [optional] 
**VizIndepVarColName** | Pointer to **string** |  | [optional] 
**VizSizeColName** | Pointer to **string** |  | [optional] 
**VizColorColName** | Pointer to **string** |  | [optional] 
**VizDataAggregation** | Pointer to **string** |  | [optional] 
**VizSortDirection** | Pointer to **string** |  | [optional] 
**VizDataLimit** | Pointer to **string** |  | [optional] 
**VizColorScheme** | Pointer to **string** |  | [optional] 
**AllowParams** | Pointer to **string** |  | [optional] 
**AcceptTerms** | Pointer to **string** |  | [optional] 
**Cached** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**NextRun** | Pointer to **string** |  | [optional] 
**DataTimePeriodStart** | Pointer to **string** |  | [optional] 
**DataTimePeriodEnd** | Pointer to **string** |  | [optional] 
**GeographicCoverageType** | Pointer to **string** |  | [optional] 
**GeographicCoverageDetails** | Pointer to **string** |  | [optional] 
**DataSourceRefreshFrequency** | Pointer to **string** |  | [optional] 
**DataSourceLastRefreshed** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **string** |  | [optional] 

## Methods

### NewAsset

`func NewAsset(companyId string, name string, source string, ) *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *Asset) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Asset) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Asset) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Asset) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetUserId

`func (o *Asset) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Asset) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Asset) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Asset) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompanyId

`func (o *Asset) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Asset) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Asset) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetConnectionId

`func (o *Asset) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Asset) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Asset) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Asset) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetLlmConnectionId

`func (o *Asset) GetLlmConnectionId() string`

GetLlmConnectionId returns the LlmConnectionId field if non-nil, zero value otherwise.

### GetLlmConnectionIdOk

`func (o *Asset) GetLlmConnectionIdOk() (*string, bool)`

GetLlmConnectionIdOk returns a tuple with the LlmConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmConnectionId

`func (o *Asset) SetLlmConnectionId(v string)`

SetLlmConnectionId sets LlmConnectionId field to given value.

### HasLlmConnectionId

`func (o *Asset) HasLlmConnectionId() bool`

HasLlmConnectionId returns a boolean if a field has been set.

### GetSnippetId

`func (o *Asset) GetSnippetId() string`

GetSnippetId returns the SnippetId field if non-nil, zero value otherwise.

### GetSnippetIdOk

`func (o *Asset) GetSnippetIdOk() (*string, bool)`

GetSnippetIdOk returns a tuple with the SnippetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetId

`func (o *Asset) SetSnippetId(v string)`

SetSnippetId sets SnippetId field to given value.

### HasSnippetId

`func (o *Asset) HasSnippetId() bool`

HasSnippetId returns a boolean if a field has been set.

### GetIndustryId

`func (o *Asset) GetIndustryId() string`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *Asset) GetIndustryIdOk() (*string, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *Asset) SetIndustryId(v string)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *Asset) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetAiJobId

`func (o *Asset) GetAiJobId() string`

GetAiJobId returns the AiJobId field if non-nil, zero value otherwise.

### GetAiJobIdOk

`func (o *Asset) GetAiJobIdOk() (*string, bool)`

GetAiJobIdOk returns a tuple with the AiJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiJobId

`func (o *Asset) SetAiJobId(v string)`

SetAiJobId sets AiJobId field to given value.

### HasAiJobId

`func (o *Asset) HasAiJobId() bool`

HasAiJobId returns a boolean if a field has been set.

### GetApprovalStatus

`func (o *Asset) GetApprovalStatus() string`

GetApprovalStatus returns the ApprovalStatus field if non-nil, zero value otherwise.

### GetApprovalStatusOk

`func (o *Asset) GetApprovalStatusOk() (*string, bool)`

GetApprovalStatusOk returns a tuple with the ApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatus

`func (o *Asset) SetApprovalStatus(v string)`

SetApprovalStatus sets ApprovalStatus field to given value.

### HasApprovalStatus

`func (o *Asset) HasApprovalStatus() bool`

HasApprovalStatus returns a boolean if a field has been set.

### GetApprovedByUserId

`func (o *Asset) GetApprovedByUserId() string`

GetApprovedByUserId returns the ApprovedByUserId field if non-nil, zero value otherwise.

### GetApprovedByUserIdOk

`func (o *Asset) GetApprovedByUserIdOk() (*string, bool)`

GetApprovedByUserIdOk returns a tuple with the ApprovedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedByUserId

`func (o *Asset) SetApprovedByUserId(v string)`

SetApprovedByUserId sets ApprovedByUserId field to given value.

### HasApprovedByUserId

`func (o *Asset) HasApprovedByUserId() bool`

HasApprovedByUserId returns a boolean if a field has been set.

### GetApprovedAt

`func (o *Asset) GetApprovedAt() string`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *Asset) GetApprovedAtOk() (*string, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *Asset) SetApprovedAt(v string)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *Asset) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetName

`func (o *Asset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Asset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Asset) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Asset) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Asset) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Asset) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Asset) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *Asset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Asset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Asset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Asset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSource

`func (o *Asset) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Asset) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Asset) SetSource(v string)`

SetSource sets Source field to given value.


### GetAssetType

`func (o *Asset) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *Asset) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *Asset) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *Asset) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetAssetSchema

`func (o *Asset) GetAssetSchema() string`

GetAssetSchema returns the AssetSchema field if non-nil, zero value otherwise.

### GetAssetSchemaOk

`func (o *Asset) GetAssetSchemaOk() (*string, bool)`

GetAssetSchemaOk returns a tuple with the AssetSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSchema

`func (o *Asset) SetAssetSchema(v string)`

SetAssetSchema sets AssetSchema field to given value.

### HasAssetSchema

`func (o *Asset) HasAssetSchema() bool`

HasAssetSchema returns a boolean if a field has been set.

### GetVisibility

`func (o *Asset) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Asset) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Asset) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Asset) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTags

`func (o *Asset) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Asset) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Asset) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Asset) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSqlLogic

`func (o *Asset) GetSqlLogic() string`

GetSqlLogic returns the SqlLogic field if non-nil, zero value otherwise.

### GetSqlLogicOk

`func (o *Asset) GetSqlLogicOk() (*string, bool)`

GetSqlLogicOk returns a tuple with the SqlLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlLogic

`func (o *Asset) SetSqlLogic(v string)`

SetSqlLogic sets SqlLogic field to given value.

### HasSqlLogic

`func (o *Asset) HasSqlLogic() bool`

HasSqlLogic returns a boolean if a field has been set.

### GetSourceSchemaName

`func (o *Asset) GetSourceSchemaName() string`

GetSourceSchemaName returns the SourceSchemaName field if non-nil, zero value otherwise.

### GetSourceSchemaNameOk

`func (o *Asset) GetSourceSchemaNameOk() (*string, bool)`

GetSourceSchemaNameOk returns a tuple with the SourceSchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchemaName

`func (o *Asset) SetSourceSchemaName(v string)`

SetSourceSchemaName sets SourceSchemaName field to given value.

### HasSourceSchemaName

`func (o *Asset) HasSourceSchemaName() bool`

HasSourceSchemaName returns a boolean if a field has been set.

### GetSourceTableName

`func (o *Asset) GetSourceTableName() string`

GetSourceTableName returns the SourceTableName field if non-nil, zero value otherwise.

### GetSourceTableNameOk

`func (o *Asset) GetSourceTableNameOk() (*string, bool)`

GetSourceTableNameOk returns a tuple with the SourceTableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTableName

`func (o *Asset) SetSourceTableName(v string)`

SetSourceTableName sets SourceTableName field to given value.

### HasSourceTableName

`func (o *Asset) HasSourceTableName() bool`

HasSourceTableName returns a boolean if a field has been set.

### GetSellInMarketplace

`func (o *Asset) GetSellInMarketplace() string`

GetSellInMarketplace returns the SellInMarketplace field if non-nil, zero value otherwise.

### GetSellInMarketplaceOk

`func (o *Asset) GetSellInMarketplaceOk() (*string, bool)`

GetSellInMarketplaceOk returns a tuple with the SellInMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellInMarketplace

`func (o *Asset) SetSellInMarketplace(v string)`

SetSellInMarketplace sets SellInMarketplace field to given value.

### HasSellInMarketplace

`func (o *Asset) HasSellInMarketplace() bool`

HasSellInMarketplace returns a boolean if a field has been set.

### GetVizChartLibrary

`func (o *Asset) GetVizChartLibrary() string`

GetVizChartLibrary returns the VizChartLibrary field if non-nil, zero value otherwise.

### GetVizChartLibraryOk

`func (o *Asset) GetVizChartLibraryOk() (*string, bool)`

GetVizChartLibraryOk returns a tuple with the VizChartLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizChartLibrary

`func (o *Asset) SetVizChartLibrary(v string)`

SetVizChartLibrary sets VizChartLibrary field to given value.

### HasVizChartLibrary

`func (o *Asset) HasVizChartLibrary() bool`

HasVizChartLibrary returns a boolean if a field has been set.

### GetVizChartType

`func (o *Asset) GetVizChartType() string`

GetVizChartType returns the VizChartType field if non-nil, zero value otherwise.

### GetVizChartTypeOk

`func (o *Asset) GetVizChartTypeOk() (*string, bool)`

GetVizChartTypeOk returns a tuple with the VizChartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizChartType

`func (o *Asset) SetVizChartType(v string)`

SetVizChartType sets VizChartType field to given value.

### HasVizChartType

`func (o *Asset) HasVizChartType() bool`

HasVizChartType returns a boolean if a field has been set.

### GetVizDepVarColName

`func (o *Asset) GetVizDepVarColName() string`

GetVizDepVarColName returns the VizDepVarColName field if non-nil, zero value otherwise.

### GetVizDepVarColNameOk

`func (o *Asset) GetVizDepVarColNameOk() (*string, bool)`

GetVizDepVarColNameOk returns a tuple with the VizDepVarColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizDepVarColName

`func (o *Asset) SetVizDepVarColName(v string)`

SetVizDepVarColName sets VizDepVarColName field to given value.

### HasVizDepVarColName

`func (o *Asset) HasVizDepVarColName() bool`

HasVizDepVarColName returns a boolean if a field has been set.

### GetVizIndepVarColName

`func (o *Asset) GetVizIndepVarColName() string`

GetVizIndepVarColName returns the VizIndepVarColName field if non-nil, zero value otherwise.

### GetVizIndepVarColNameOk

`func (o *Asset) GetVizIndepVarColNameOk() (*string, bool)`

GetVizIndepVarColNameOk returns a tuple with the VizIndepVarColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizIndepVarColName

`func (o *Asset) SetVizIndepVarColName(v string)`

SetVizIndepVarColName sets VizIndepVarColName field to given value.

### HasVizIndepVarColName

`func (o *Asset) HasVizIndepVarColName() bool`

HasVizIndepVarColName returns a boolean if a field has been set.

### GetVizSizeColName

`func (o *Asset) GetVizSizeColName() string`

GetVizSizeColName returns the VizSizeColName field if non-nil, zero value otherwise.

### GetVizSizeColNameOk

`func (o *Asset) GetVizSizeColNameOk() (*string, bool)`

GetVizSizeColNameOk returns a tuple with the VizSizeColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizSizeColName

`func (o *Asset) SetVizSizeColName(v string)`

SetVizSizeColName sets VizSizeColName field to given value.

### HasVizSizeColName

`func (o *Asset) HasVizSizeColName() bool`

HasVizSizeColName returns a boolean if a field has been set.

### GetVizColorColName

`func (o *Asset) GetVizColorColName() string`

GetVizColorColName returns the VizColorColName field if non-nil, zero value otherwise.

### GetVizColorColNameOk

`func (o *Asset) GetVizColorColNameOk() (*string, bool)`

GetVizColorColNameOk returns a tuple with the VizColorColName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizColorColName

`func (o *Asset) SetVizColorColName(v string)`

SetVizColorColName sets VizColorColName field to given value.

### HasVizColorColName

`func (o *Asset) HasVizColorColName() bool`

HasVizColorColName returns a boolean if a field has been set.

### GetVizDataAggregation

`func (o *Asset) GetVizDataAggregation() string`

GetVizDataAggregation returns the VizDataAggregation field if non-nil, zero value otherwise.

### GetVizDataAggregationOk

`func (o *Asset) GetVizDataAggregationOk() (*string, bool)`

GetVizDataAggregationOk returns a tuple with the VizDataAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizDataAggregation

`func (o *Asset) SetVizDataAggregation(v string)`

SetVizDataAggregation sets VizDataAggregation field to given value.

### HasVizDataAggregation

`func (o *Asset) HasVizDataAggregation() bool`

HasVizDataAggregation returns a boolean if a field has been set.

### GetVizSortDirection

`func (o *Asset) GetVizSortDirection() string`

GetVizSortDirection returns the VizSortDirection field if non-nil, zero value otherwise.

### GetVizSortDirectionOk

`func (o *Asset) GetVizSortDirectionOk() (*string, bool)`

GetVizSortDirectionOk returns a tuple with the VizSortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizSortDirection

`func (o *Asset) SetVizSortDirection(v string)`

SetVizSortDirection sets VizSortDirection field to given value.

### HasVizSortDirection

`func (o *Asset) HasVizSortDirection() bool`

HasVizSortDirection returns a boolean if a field has been set.

### GetVizDataLimit

`func (o *Asset) GetVizDataLimit() string`

GetVizDataLimit returns the VizDataLimit field if non-nil, zero value otherwise.

### GetVizDataLimitOk

`func (o *Asset) GetVizDataLimitOk() (*string, bool)`

GetVizDataLimitOk returns a tuple with the VizDataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizDataLimit

`func (o *Asset) SetVizDataLimit(v string)`

SetVizDataLimit sets VizDataLimit field to given value.

### HasVizDataLimit

`func (o *Asset) HasVizDataLimit() bool`

HasVizDataLimit returns a boolean if a field has been set.

### GetVizColorScheme

`func (o *Asset) GetVizColorScheme() string`

GetVizColorScheme returns the VizColorScheme field if non-nil, zero value otherwise.

### GetVizColorSchemeOk

`func (o *Asset) GetVizColorSchemeOk() (*string, bool)`

GetVizColorSchemeOk returns a tuple with the VizColorScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizColorScheme

`func (o *Asset) SetVizColorScheme(v string)`

SetVizColorScheme sets VizColorScheme field to given value.

### HasVizColorScheme

`func (o *Asset) HasVizColorScheme() bool`

HasVizColorScheme returns a boolean if a field has been set.

### GetAllowParams

`func (o *Asset) GetAllowParams() string`

GetAllowParams returns the AllowParams field if non-nil, zero value otherwise.

### GetAllowParamsOk

`func (o *Asset) GetAllowParamsOk() (*string, bool)`

GetAllowParamsOk returns a tuple with the AllowParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowParams

`func (o *Asset) SetAllowParams(v string)`

SetAllowParams sets AllowParams field to given value.

### HasAllowParams

`func (o *Asset) HasAllowParams() bool`

HasAllowParams returns a boolean if a field has been set.

### GetAcceptTerms

`func (o *Asset) GetAcceptTerms() string`

GetAcceptTerms returns the AcceptTerms field if non-nil, zero value otherwise.

### GetAcceptTermsOk

`func (o *Asset) GetAcceptTermsOk() (*string, bool)`

GetAcceptTermsOk returns a tuple with the AcceptTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTerms

`func (o *Asset) SetAcceptTerms(v string)`

SetAcceptTerms sets AcceptTerms field to given value.

### HasAcceptTerms

`func (o *Asset) HasAcceptTerms() bool`

HasAcceptTerms returns a boolean if a field has been set.

### GetCached

`func (o *Asset) GetCached() string`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *Asset) GetCachedOk() (*string, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *Asset) SetCached(v string)`

SetCached sets Cached field to given value.

### HasCached

`func (o *Asset) HasCached() bool`

HasCached returns a boolean if a field has been set.

### GetSchedule

`func (o *Asset) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Asset) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Asset) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Asset) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetNextRun

`func (o *Asset) GetNextRun() string`

GetNextRun returns the NextRun field if non-nil, zero value otherwise.

### GetNextRunOk

`func (o *Asset) GetNextRunOk() (*string, bool)`

GetNextRunOk returns a tuple with the NextRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRun

`func (o *Asset) SetNextRun(v string)`

SetNextRun sets NextRun field to given value.

### HasNextRun

`func (o *Asset) HasNextRun() bool`

HasNextRun returns a boolean if a field has been set.

### GetDataTimePeriodStart

`func (o *Asset) GetDataTimePeriodStart() string`

GetDataTimePeriodStart returns the DataTimePeriodStart field if non-nil, zero value otherwise.

### GetDataTimePeriodStartOk

`func (o *Asset) GetDataTimePeriodStartOk() (*string, bool)`

GetDataTimePeriodStartOk returns a tuple with the DataTimePeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimePeriodStart

`func (o *Asset) SetDataTimePeriodStart(v string)`

SetDataTimePeriodStart sets DataTimePeriodStart field to given value.

### HasDataTimePeriodStart

`func (o *Asset) HasDataTimePeriodStart() bool`

HasDataTimePeriodStart returns a boolean if a field has been set.

### GetDataTimePeriodEnd

`func (o *Asset) GetDataTimePeriodEnd() string`

GetDataTimePeriodEnd returns the DataTimePeriodEnd field if non-nil, zero value otherwise.

### GetDataTimePeriodEndOk

`func (o *Asset) GetDataTimePeriodEndOk() (*string, bool)`

GetDataTimePeriodEndOk returns a tuple with the DataTimePeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimePeriodEnd

`func (o *Asset) SetDataTimePeriodEnd(v string)`

SetDataTimePeriodEnd sets DataTimePeriodEnd field to given value.

### HasDataTimePeriodEnd

`func (o *Asset) HasDataTimePeriodEnd() bool`

HasDataTimePeriodEnd returns a boolean if a field has been set.

### GetGeographicCoverageType

`func (o *Asset) GetGeographicCoverageType() string`

GetGeographicCoverageType returns the GeographicCoverageType field if non-nil, zero value otherwise.

### GetGeographicCoverageTypeOk

`func (o *Asset) GetGeographicCoverageTypeOk() (*string, bool)`

GetGeographicCoverageTypeOk returns a tuple with the GeographicCoverageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicCoverageType

`func (o *Asset) SetGeographicCoverageType(v string)`

SetGeographicCoverageType sets GeographicCoverageType field to given value.

### HasGeographicCoverageType

`func (o *Asset) HasGeographicCoverageType() bool`

HasGeographicCoverageType returns a boolean if a field has been set.

### GetGeographicCoverageDetails

`func (o *Asset) GetGeographicCoverageDetails() string`

GetGeographicCoverageDetails returns the GeographicCoverageDetails field if non-nil, zero value otherwise.

### GetGeographicCoverageDetailsOk

`func (o *Asset) GetGeographicCoverageDetailsOk() (*string, bool)`

GetGeographicCoverageDetailsOk returns a tuple with the GeographicCoverageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicCoverageDetails

`func (o *Asset) SetGeographicCoverageDetails(v string)`

SetGeographicCoverageDetails sets GeographicCoverageDetails field to given value.

### HasGeographicCoverageDetails

`func (o *Asset) HasGeographicCoverageDetails() bool`

HasGeographicCoverageDetails returns a boolean if a field has been set.

### GetDataSourceRefreshFrequency

`func (o *Asset) GetDataSourceRefreshFrequency() string`

GetDataSourceRefreshFrequency returns the DataSourceRefreshFrequency field if non-nil, zero value otherwise.

### GetDataSourceRefreshFrequencyOk

`func (o *Asset) GetDataSourceRefreshFrequencyOk() (*string, bool)`

GetDataSourceRefreshFrequencyOk returns a tuple with the DataSourceRefreshFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceRefreshFrequency

`func (o *Asset) SetDataSourceRefreshFrequency(v string)`

SetDataSourceRefreshFrequency sets DataSourceRefreshFrequency field to given value.

### HasDataSourceRefreshFrequency

`func (o *Asset) HasDataSourceRefreshFrequency() bool`

HasDataSourceRefreshFrequency returns a boolean if a field has been set.

### GetDataSourceLastRefreshed

`func (o *Asset) GetDataSourceLastRefreshed() string`

GetDataSourceLastRefreshed returns the DataSourceLastRefreshed field if non-nil, zero value otherwise.

### GetDataSourceLastRefreshedOk

`func (o *Asset) GetDataSourceLastRefreshedOk() (*string, bool)`

GetDataSourceLastRefreshedOk returns a tuple with the DataSourceLastRefreshed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceLastRefreshed

`func (o *Asset) SetDataSourceLastRefreshed(v string)`

SetDataSourceLastRefreshed sets DataSourceLastRefreshed field to given value.

### HasDataSourceLastRefreshed

`func (o *Asset) HasDataSourceLastRefreshed() bool`

HasDataSourceLastRefreshed returns a boolean if a field has been set.

### GetDateCreated

`func (o *Asset) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Asset) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Asset) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Asset) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Asset) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Asset) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Asset) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Asset) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *Asset) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Asset) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Asset) SetActive(v string)`

SetActive sets Active field to given value.

### HasActive

`func (o *Asset) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



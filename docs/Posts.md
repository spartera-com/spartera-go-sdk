# Posts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** | Optional. | [optional] 
**LastUpdated** | Pointer to **time.Time** | Optional. | [optional] 
**PostId** | Pointer to **string** | Unique identifier. | [optional] 
**UserId** | **string** | User who created this post | 
**CompanyId** | **string** | Company this post belongs to | 
**Title** | **string** | Article title | 
**Category** | Pointer to **string** | Article category (e.g., &#39;Sports&#39;, &#39;Business&#39;) | [optional] 
**Teaser** | Pointer to **string** | Article teaser/subtitle | [optional] 
**ContentHtml** | **string** | Generated article HTML content | 
**InsightsUsed** | Pointer to **map[string]interface{}** | Array of insights used: [{asset_id, asset_name, value, runtime, success}] | [optional] 
**GenerationCreativity** | Pointer to **int64** | Creativity level (0-100), maps to AI temperature | [optional] 
**GenerationTargetWordCount** | Pointer to **int64** | Target word count (null &#x3D; auto) | [optional] 
**GenerationTone** | Pointer to **string** | Writing tone: professional, casual, technical, conversational | [optional] 
**GenerationIncludeCitations** | Pointer to **bool** | Whether to include data source citations | [optional] 
**GenerationSubheadingCount** | Pointer to **int64** | Number of subheadings (2-5) | [optional] 
**GenerationTemperature** | Pointer to **float32** | Actual temperature used for generation (0.0-2.0) | [optional] 
**DataCostCredits** | **int64** | Cost in credits for data insights | 
**ServiceCostCredits** | **int64** | Cost in credits for AI generation service | 
**TotalCostCredits** | **int64** | Total cost in credits (data + service) | 
**SuccessfulInsightsCount** | **int64** | Number of insights that succeeded | 
**FailedInsightsCount** | **int64** | Number of insights that failed | 
**GenerationTimeMs** | Pointer to **int64** | Time taken to generate article in milliseconds | [optional] 
**IsPublished** | **bool** | Whether this post has been published | 
**PublishedAt** | Pointer to **time.Time** | When this post was published | [optional] 
**ViewCount** | **int64** | Number of times this post has been viewed | 
**LastEditedAt** | Pointer to **time.Time** | When this post was last edited | [optional] 

## Methods

### NewPosts

`func NewPosts(userId string, companyId string, title string, contentHtml string, dataCostCredits int64, serviceCostCredits int64, totalCostCredits int64, successfulInsightsCount int64, failedInsightsCount int64, isPublished bool, viewCount int64, ) *Posts`

NewPosts instantiates a new Posts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostsWithDefaults

`func NewPostsWithDefaults() *Posts`

NewPostsWithDefaults instantiates a new Posts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *Posts) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Posts) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Posts) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Posts) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Posts) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Posts) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Posts) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Posts) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPostId

`func (o *Posts) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *Posts) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *Posts) SetPostId(v string)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *Posts) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetUserId

`func (o *Posts) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Posts) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Posts) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCompanyId

`func (o *Posts) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Posts) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Posts) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetTitle

`func (o *Posts) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Posts) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Posts) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCategory

`func (o *Posts) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Posts) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Posts) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Posts) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetTeaser

`func (o *Posts) GetTeaser() string`

GetTeaser returns the Teaser field if non-nil, zero value otherwise.

### GetTeaserOk

`func (o *Posts) GetTeaserOk() (*string, bool)`

GetTeaserOk returns a tuple with the Teaser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeaser

`func (o *Posts) SetTeaser(v string)`

SetTeaser sets Teaser field to given value.

### HasTeaser

`func (o *Posts) HasTeaser() bool`

HasTeaser returns a boolean if a field has been set.

### GetContentHtml

`func (o *Posts) GetContentHtml() string`

GetContentHtml returns the ContentHtml field if non-nil, zero value otherwise.

### GetContentHtmlOk

`func (o *Posts) GetContentHtmlOk() (*string, bool)`

GetContentHtmlOk returns a tuple with the ContentHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHtml

`func (o *Posts) SetContentHtml(v string)`

SetContentHtml sets ContentHtml field to given value.


### GetInsightsUsed

`func (o *Posts) GetInsightsUsed() map[string]interface{}`

GetInsightsUsed returns the InsightsUsed field if non-nil, zero value otherwise.

### GetInsightsUsedOk

`func (o *Posts) GetInsightsUsedOk() (*map[string]interface{}, bool)`

GetInsightsUsedOk returns a tuple with the InsightsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsUsed

`func (o *Posts) SetInsightsUsed(v map[string]interface{})`

SetInsightsUsed sets InsightsUsed field to given value.

### HasInsightsUsed

`func (o *Posts) HasInsightsUsed() bool`

HasInsightsUsed returns a boolean if a field has been set.

### GetGenerationCreativity

`func (o *Posts) GetGenerationCreativity() int64`

GetGenerationCreativity returns the GenerationCreativity field if non-nil, zero value otherwise.

### GetGenerationCreativityOk

`func (o *Posts) GetGenerationCreativityOk() (*int64, bool)`

GetGenerationCreativityOk returns a tuple with the GenerationCreativity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationCreativity

`func (o *Posts) SetGenerationCreativity(v int64)`

SetGenerationCreativity sets GenerationCreativity field to given value.

### HasGenerationCreativity

`func (o *Posts) HasGenerationCreativity() bool`

HasGenerationCreativity returns a boolean if a field has been set.

### GetGenerationTargetWordCount

`func (o *Posts) GetGenerationTargetWordCount() int64`

GetGenerationTargetWordCount returns the GenerationTargetWordCount field if non-nil, zero value otherwise.

### GetGenerationTargetWordCountOk

`func (o *Posts) GetGenerationTargetWordCountOk() (*int64, bool)`

GetGenerationTargetWordCountOk returns a tuple with the GenerationTargetWordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationTargetWordCount

`func (o *Posts) SetGenerationTargetWordCount(v int64)`

SetGenerationTargetWordCount sets GenerationTargetWordCount field to given value.

### HasGenerationTargetWordCount

`func (o *Posts) HasGenerationTargetWordCount() bool`

HasGenerationTargetWordCount returns a boolean if a field has been set.

### GetGenerationTone

`func (o *Posts) GetGenerationTone() string`

GetGenerationTone returns the GenerationTone field if non-nil, zero value otherwise.

### GetGenerationToneOk

`func (o *Posts) GetGenerationToneOk() (*string, bool)`

GetGenerationToneOk returns a tuple with the GenerationTone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationTone

`func (o *Posts) SetGenerationTone(v string)`

SetGenerationTone sets GenerationTone field to given value.

### HasGenerationTone

`func (o *Posts) HasGenerationTone() bool`

HasGenerationTone returns a boolean if a field has been set.

### GetGenerationIncludeCitations

`func (o *Posts) GetGenerationIncludeCitations() bool`

GetGenerationIncludeCitations returns the GenerationIncludeCitations field if non-nil, zero value otherwise.

### GetGenerationIncludeCitationsOk

`func (o *Posts) GetGenerationIncludeCitationsOk() (*bool, bool)`

GetGenerationIncludeCitationsOk returns a tuple with the GenerationIncludeCitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationIncludeCitations

`func (o *Posts) SetGenerationIncludeCitations(v bool)`

SetGenerationIncludeCitations sets GenerationIncludeCitations field to given value.

### HasGenerationIncludeCitations

`func (o *Posts) HasGenerationIncludeCitations() bool`

HasGenerationIncludeCitations returns a boolean if a field has been set.

### GetGenerationSubheadingCount

`func (o *Posts) GetGenerationSubheadingCount() int64`

GetGenerationSubheadingCount returns the GenerationSubheadingCount field if non-nil, zero value otherwise.

### GetGenerationSubheadingCountOk

`func (o *Posts) GetGenerationSubheadingCountOk() (*int64, bool)`

GetGenerationSubheadingCountOk returns a tuple with the GenerationSubheadingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationSubheadingCount

`func (o *Posts) SetGenerationSubheadingCount(v int64)`

SetGenerationSubheadingCount sets GenerationSubheadingCount field to given value.

### HasGenerationSubheadingCount

`func (o *Posts) HasGenerationSubheadingCount() bool`

HasGenerationSubheadingCount returns a boolean if a field has been set.

### GetGenerationTemperature

`func (o *Posts) GetGenerationTemperature() float32`

GetGenerationTemperature returns the GenerationTemperature field if non-nil, zero value otherwise.

### GetGenerationTemperatureOk

`func (o *Posts) GetGenerationTemperatureOk() (*float32, bool)`

GetGenerationTemperatureOk returns a tuple with the GenerationTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationTemperature

`func (o *Posts) SetGenerationTemperature(v float32)`

SetGenerationTemperature sets GenerationTemperature field to given value.

### HasGenerationTemperature

`func (o *Posts) HasGenerationTemperature() bool`

HasGenerationTemperature returns a boolean if a field has been set.

### GetDataCostCredits

`func (o *Posts) GetDataCostCredits() int64`

GetDataCostCredits returns the DataCostCredits field if non-nil, zero value otherwise.

### GetDataCostCreditsOk

`func (o *Posts) GetDataCostCreditsOk() (*int64, bool)`

GetDataCostCreditsOk returns a tuple with the DataCostCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCostCredits

`func (o *Posts) SetDataCostCredits(v int64)`

SetDataCostCredits sets DataCostCredits field to given value.


### GetServiceCostCredits

`func (o *Posts) GetServiceCostCredits() int64`

GetServiceCostCredits returns the ServiceCostCredits field if non-nil, zero value otherwise.

### GetServiceCostCreditsOk

`func (o *Posts) GetServiceCostCreditsOk() (*int64, bool)`

GetServiceCostCreditsOk returns a tuple with the ServiceCostCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCostCredits

`func (o *Posts) SetServiceCostCredits(v int64)`

SetServiceCostCredits sets ServiceCostCredits field to given value.


### GetTotalCostCredits

`func (o *Posts) GetTotalCostCredits() int64`

GetTotalCostCredits returns the TotalCostCredits field if non-nil, zero value otherwise.

### GetTotalCostCreditsOk

`func (o *Posts) GetTotalCostCreditsOk() (*int64, bool)`

GetTotalCostCreditsOk returns a tuple with the TotalCostCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostCredits

`func (o *Posts) SetTotalCostCredits(v int64)`

SetTotalCostCredits sets TotalCostCredits field to given value.


### GetSuccessfulInsightsCount

`func (o *Posts) GetSuccessfulInsightsCount() int64`

GetSuccessfulInsightsCount returns the SuccessfulInsightsCount field if non-nil, zero value otherwise.

### GetSuccessfulInsightsCountOk

`func (o *Posts) GetSuccessfulInsightsCountOk() (*int64, bool)`

GetSuccessfulInsightsCountOk returns a tuple with the SuccessfulInsightsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulInsightsCount

`func (o *Posts) SetSuccessfulInsightsCount(v int64)`

SetSuccessfulInsightsCount sets SuccessfulInsightsCount field to given value.


### GetFailedInsightsCount

`func (o *Posts) GetFailedInsightsCount() int64`

GetFailedInsightsCount returns the FailedInsightsCount field if non-nil, zero value otherwise.

### GetFailedInsightsCountOk

`func (o *Posts) GetFailedInsightsCountOk() (*int64, bool)`

GetFailedInsightsCountOk returns a tuple with the FailedInsightsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedInsightsCount

`func (o *Posts) SetFailedInsightsCount(v int64)`

SetFailedInsightsCount sets FailedInsightsCount field to given value.


### GetGenerationTimeMs

`func (o *Posts) GetGenerationTimeMs() int64`

GetGenerationTimeMs returns the GenerationTimeMs field if non-nil, zero value otherwise.

### GetGenerationTimeMsOk

`func (o *Posts) GetGenerationTimeMsOk() (*int64, bool)`

GetGenerationTimeMsOk returns a tuple with the GenerationTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationTimeMs

`func (o *Posts) SetGenerationTimeMs(v int64)`

SetGenerationTimeMs sets GenerationTimeMs field to given value.

### HasGenerationTimeMs

`func (o *Posts) HasGenerationTimeMs() bool`

HasGenerationTimeMs returns a boolean if a field has been set.

### GetIsPublished

`func (o *Posts) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *Posts) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *Posts) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.


### GetPublishedAt

`func (o *Posts) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *Posts) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *Posts) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *Posts) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetViewCount

`func (o *Posts) GetViewCount() int64`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *Posts) GetViewCountOk() (*int64, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *Posts) SetViewCount(v int64)`

SetViewCount sets ViewCount field to given value.


### GetLastEditedAt

`func (o *Posts) GetLastEditedAt() time.Time`

GetLastEditedAt returns the LastEditedAt field if non-nil, zero value otherwise.

### GetLastEditedAtOk

`func (o *Posts) GetLastEditedAtOk() (*time.Time, bool)`

GetLastEditedAtOk returns a tuple with the LastEditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditedAt

`func (o *Posts) SetLastEditedAt(v time.Time)`

SetLastEditedAt sets LastEditedAt field to given value.

### HasLastEditedAt

`func (o *Posts) HasLastEditedAt() bool`

HasLastEditedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



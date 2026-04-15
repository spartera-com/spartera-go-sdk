# PostsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**DataCostCredits** | Pointer to **int64** | Cost in credits for data insights | [optional] 
**ServiceCostCredits** | Pointer to **int64** | Cost in credits for AI generation service | [optional] 
**TotalCostCredits** | Pointer to **int64** | Total cost in credits (data + service) | [optional] 
**SuccessfulInsightsCount** | Pointer to **int64** | Number of insights that succeeded | [optional] 
**FailedInsightsCount** | Pointer to **int64** | Number of insights that failed | [optional] 
**GenerationTimeMs** | Pointer to **int64** | Time taken to generate article in milliseconds | [optional] 
**IsPublished** | Pointer to **bool** | Whether this post has been published | [optional] 
**PublishedAt** | Pointer to **time.Time** | When this post was published | [optional] 
**ViewCount** | Pointer to **int64** | Number of times this post has been viewed | [optional] 
**LastEditedAt** | Pointer to **time.Time** | When this post was last edited | [optional] 

## Methods

### NewPostsInput

`func NewPostsInput(userId string, companyId string, title string, contentHtml string, ) *PostsInput`

NewPostsInput instantiates a new PostsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostsInputWithDefaults

`func NewPostsInputWithDefaults() *PostsInput`

NewPostsInputWithDefaults instantiates a new PostsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PostsInput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostsInput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostsInput) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCompanyId

`func (o *PostsInput) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PostsInput) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PostsInput) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetTitle

`func (o *PostsInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostsInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostsInput) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCategory

`func (o *PostsInput) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PostsInput) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PostsInput) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PostsInput) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetTeaser

`func (o *PostsInput) GetTeaser() string`

GetTeaser returns the Teaser field if non-nil, zero value otherwise.

### GetTeaserOk

`func (o *PostsInput) GetTeaserOk() (*string, bool)`

GetTeaserOk returns a tuple with the Teaser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeaser

`func (o *PostsInput) SetTeaser(v string)`

SetTeaser sets Teaser field to given value.

### HasTeaser

`func (o *PostsInput) HasTeaser() bool`

HasTeaser returns a boolean if a field has been set.

### GetContentHtml

`func (o *PostsInput) GetContentHtml() string`

GetContentHtml returns the ContentHtml field if non-nil, zero value otherwise.

### GetContentHtmlOk

`func (o *PostsInput) GetContentHtmlOk() (*string, bool)`

GetContentHtmlOk returns a tuple with the ContentHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHtml

`func (o *PostsInput) SetContentHtml(v string)`

SetContentHtml sets ContentHtml field to given value.


### GetInsightsUsed

`func (o *PostsInput) GetInsightsUsed() map[string]interface{}`

GetInsightsUsed returns the InsightsUsed field if non-nil, zero value otherwise.

### GetInsightsUsedOk

`func (o *PostsInput) GetInsightsUsedOk() (*map[string]interface{}, bool)`

GetInsightsUsedOk returns a tuple with the InsightsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsUsed

`func (o *PostsInput) SetInsightsUsed(v map[string]interface{})`

SetInsightsUsed sets InsightsUsed field to given value.

### HasInsightsUsed

`func (o *PostsInput) HasInsightsUsed() bool`

HasInsightsUsed returns a boolean if a field has been set.

### GetGenerationCreativity

`func (o *PostsInput) GetGenerationCreativity() int64`

GetGenerationCreativity returns the GenerationCreativity field if non-nil, zero value otherwise.

### GetGenerationCreativityOk

`func (o *PostsInput) GetGenerationCreativityOk() (*int64, bool)`

GetGenerationCreativityOk returns a tuple with the GenerationCreativity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationCreativity

`func (o *PostsInput) SetGenerationCreativity(v int64)`

SetGenerationCreativity sets GenerationCreativity field to given value.

### HasGenerationCreativity

`func (o *PostsInput) HasGenerationCreativity() bool`

HasGenerationCreativity returns a boolean if a field has been set.

### GetGenerationTargetWordCount

`func (o *PostsInput) GetGenerationTargetWordCount() int64`

GetGenerationTargetWordCount returns the GenerationTargetWordCount field if non-nil, zero value otherwise.

### GetGenerationTargetWordCountOk

`func (o *PostsInput) GetGenerationTargetWordCountOk() (*int64, bool)`

GetGenerationTargetWordCountOk returns a tuple with the GenerationTargetWordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationTargetWordCount

`func (o *PostsInput) SetGenerationTargetWordCount(v int64)`

SetGenerationTargetWordCount sets GenerationTargetWordCount field to given value.

### HasGenerationTargetWordCount

`func (o *PostsInput) HasGenerationTargetWordCount() bool`

HasGenerationTargetWordCount returns a boolean if a field has been set.

### GetGenerationTone

`func (o *PostsInput) GetGenerationTone() string`

GetGenerationTone returns the GenerationTone field if non-nil, zero value otherwise.

### GetGenerationToneOk

`func (o *PostsInput) GetGenerationToneOk() (*string, bool)`

GetGenerationToneOk returns a tuple with the GenerationTone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationTone

`func (o *PostsInput) SetGenerationTone(v string)`

SetGenerationTone sets GenerationTone field to given value.

### HasGenerationTone

`func (o *PostsInput) HasGenerationTone() bool`

HasGenerationTone returns a boolean if a field has been set.

### GetGenerationIncludeCitations

`func (o *PostsInput) GetGenerationIncludeCitations() bool`

GetGenerationIncludeCitations returns the GenerationIncludeCitations field if non-nil, zero value otherwise.

### GetGenerationIncludeCitationsOk

`func (o *PostsInput) GetGenerationIncludeCitationsOk() (*bool, bool)`

GetGenerationIncludeCitationsOk returns a tuple with the GenerationIncludeCitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationIncludeCitations

`func (o *PostsInput) SetGenerationIncludeCitations(v bool)`

SetGenerationIncludeCitations sets GenerationIncludeCitations field to given value.

### HasGenerationIncludeCitations

`func (o *PostsInput) HasGenerationIncludeCitations() bool`

HasGenerationIncludeCitations returns a boolean if a field has been set.

### GetGenerationSubheadingCount

`func (o *PostsInput) GetGenerationSubheadingCount() int64`

GetGenerationSubheadingCount returns the GenerationSubheadingCount field if non-nil, zero value otherwise.

### GetGenerationSubheadingCountOk

`func (o *PostsInput) GetGenerationSubheadingCountOk() (*int64, bool)`

GetGenerationSubheadingCountOk returns a tuple with the GenerationSubheadingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationSubheadingCount

`func (o *PostsInput) SetGenerationSubheadingCount(v int64)`

SetGenerationSubheadingCount sets GenerationSubheadingCount field to given value.

### HasGenerationSubheadingCount

`func (o *PostsInput) HasGenerationSubheadingCount() bool`

HasGenerationSubheadingCount returns a boolean if a field has been set.

### GetGenerationTemperature

`func (o *PostsInput) GetGenerationTemperature() float32`

GetGenerationTemperature returns the GenerationTemperature field if non-nil, zero value otherwise.

### GetGenerationTemperatureOk

`func (o *PostsInput) GetGenerationTemperatureOk() (*float32, bool)`

GetGenerationTemperatureOk returns a tuple with the GenerationTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationTemperature

`func (o *PostsInput) SetGenerationTemperature(v float32)`

SetGenerationTemperature sets GenerationTemperature field to given value.

### HasGenerationTemperature

`func (o *PostsInput) HasGenerationTemperature() bool`

HasGenerationTemperature returns a boolean if a field has been set.

### GetDataCostCredits

`func (o *PostsInput) GetDataCostCredits() int64`

GetDataCostCredits returns the DataCostCredits field if non-nil, zero value otherwise.

### GetDataCostCreditsOk

`func (o *PostsInput) GetDataCostCreditsOk() (*int64, bool)`

GetDataCostCreditsOk returns a tuple with the DataCostCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCostCredits

`func (o *PostsInput) SetDataCostCredits(v int64)`

SetDataCostCredits sets DataCostCredits field to given value.

### HasDataCostCredits

`func (o *PostsInput) HasDataCostCredits() bool`

HasDataCostCredits returns a boolean if a field has been set.

### GetServiceCostCredits

`func (o *PostsInput) GetServiceCostCredits() int64`

GetServiceCostCredits returns the ServiceCostCredits field if non-nil, zero value otherwise.

### GetServiceCostCreditsOk

`func (o *PostsInput) GetServiceCostCreditsOk() (*int64, bool)`

GetServiceCostCreditsOk returns a tuple with the ServiceCostCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCostCredits

`func (o *PostsInput) SetServiceCostCredits(v int64)`

SetServiceCostCredits sets ServiceCostCredits field to given value.

### HasServiceCostCredits

`func (o *PostsInput) HasServiceCostCredits() bool`

HasServiceCostCredits returns a boolean if a field has been set.

### GetTotalCostCredits

`func (o *PostsInput) GetTotalCostCredits() int64`

GetTotalCostCredits returns the TotalCostCredits field if non-nil, zero value otherwise.

### GetTotalCostCreditsOk

`func (o *PostsInput) GetTotalCostCreditsOk() (*int64, bool)`

GetTotalCostCreditsOk returns a tuple with the TotalCostCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostCredits

`func (o *PostsInput) SetTotalCostCredits(v int64)`

SetTotalCostCredits sets TotalCostCredits field to given value.

### HasTotalCostCredits

`func (o *PostsInput) HasTotalCostCredits() bool`

HasTotalCostCredits returns a boolean if a field has been set.

### GetSuccessfulInsightsCount

`func (o *PostsInput) GetSuccessfulInsightsCount() int64`

GetSuccessfulInsightsCount returns the SuccessfulInsightsCount field if non-nil, zero value otherwise.

### GetSuccessfulInsightsCountOk

`func (o *PostsInput) GetSuccessfulInsightsCountOk() (*int64, bool)`

GetSuccessfulInsightsCountOk returns a tuple with the SuccessfulInsightsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulInsightsCount

`func (o *PostsInput) SetSuccessfulInsightsCount(v int64)`

SetSuccessfulInsightsCount sets SuccessfulInsightsCount field to given value.

### HasSuccessfulInsightsCount

`func (o *PostsInput) HasSuccessfulInsightsCount() bool`

HasSuccessfulInsightsCount returns a boolean if a field has been set.

### GetFailedInsightsCount

`func (o *PostsInput) GetFailedInsightsCount() int64`

GetFailedInsightsCount returns the FailedInsightsCount field if non-nil, zero value otherwise.

### GetFailedInsightsCountOk

`func (o *PostsInput) GetFailedInsightsCountOk() (*int64, bool)`

GetFailedInsightsCountOk returns a tuple with the FailedInsightsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedInsightsCount

`func (o *PostsInput) SetFailedInsightsCount(v int64)`

SetFailedInsightsCount sets FailedInsightsCount field to given value.

### HasFailedInsightsCount

`func (o *PostsInput) HasFailedInsightsCount() bool`

HasFailedInsightsCount returns a boolean if a field has been set.

### GetGenerationTimeMs

`func (o *PostsInput) GetGenerationTimeMs() int64`

GetGenerationTimeMs returns the GenerationTimeMs field if non-nil, zero value otherwise.

### GetGenerationTimeMsOk

`func (o *PostsInput) GetGenerationTimeMsOk() (*int64, bool)`

GetGenerationTimeMsOk returns a tuple with the GenerationTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationTimeMs

`func (o *PostsInput) SetGenerationTimeMs(v int64)`

SetGenerationTimeMs sets GenerationTimeMs field to given value.

### HasGenerationTimeMs

`func (o *PostsInput) HasGenerationTimeMs() bool`

HasGenerationTimeMs returns a boolean if a field has been set.

### GetIsPublished

`func (o *PostsInput) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *PostsInput) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *PostsInput) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *PostsInput) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetPublishedAt

`func (o *PostsInput) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PostsInput) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PostsInput) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *PostsInput) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetViewCount

`func (o *PostsInput) GetViewCount() int64`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *PostsInput) GetViewCountOk() (*int64, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *PostsInput) SetViewCount(v int64)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *PostsInput) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### GetLastEditedAt

`func (o *PostsInput) GetLastEditedAt() time.Time`

GetLastEditedAt returns the LastEditedAt field if non-nil, zero value otherwise.

### GetLastEditedAtOk

`func (o *PostsInput) GetLastEditedAtOk() (*time.Time, bool)`

GetLastEditedAtOk returns a tuple with the LastEditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditedAt

`func (o *PostsInput) SetLastEditedAt(v time.Time)`

SetLastEditedAt sets LastEditedAt field to given value.

### HasLastEditedAt

`func (o *PostsInput) HasLastEditedAt() bool`

HasLastEditedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



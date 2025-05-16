# PublicPageViewAnalyticsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoalescingRefineBy** | Pointer to [**PublicFormSubmissionFilterCoalescingRefineBy**](PublicFormSubmissionFilterCoalescingRefineBy.md) |  | [optional] 
**EnableTracking** | Pointer to **bool** |  | [optional] 
**PruningRefineBy** | Pointer to [**PublicFormSubmissionFilterCoalescingRefineBy**](PublicFormSubmissionFilterCoalescingRefineBy.md) |  | [optional] 
**PageUrl** | **string** |  | 
**FilterType** | **string** |  | [default to "PAGE_VIEW"]
**Operator** | **string** |  | 

## Methods

### NewPublicPageViewAnalyticsFilter

`func NewPublicPageViewAnalyticsFilter(pageUrl string, filterType string, operator string, ) *PublicPageViewAnalyticsFilter`

NewPublicPageViewAnalyticsFilter instantiates a new PublicPageViewAnalyticsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPageViewAnalyticsFilterWithDefaults

`func NewPublicPageViewAnalyticsFilterWithDefaults() *PublicPageViewAnalyticsFilter`

NewPublicPageViewAnalyticsFilterWithDefaults instantiates a new PublicPageViewAnalyticsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoalescingRefineBy

`func (o *PublicPageViewAnalyticsFilter) GetCoalescingRefineBy() PublicFormSubmissionFilterCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *PublicPageViewAnalyticsFilter) GetCoalescingRefineByOk() (*PublicFormSubmissionFilterCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *PublicPageViewAnalyticsFilter) SetCoalescingRefineBy(v PublicFormSubmissionFilterCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.

### HasCoalescingRefineBy

`func (o *PublicPageViewAnalyticsFilter) HasCoalescingRefineBy() bool`

HasCoalescingRefineBy returns a boolean if a field has been set.

### GetEnableTracking

`func (o *PublicPageViewAnalyticsFilter) GetEnableTracking() bool`

GetEnableTracking returns the EnableTracking field if non-nil, zero value otherwise.

### GetEnableTrackingOk

`func (o *PublicPageViewAnalyticsFilter) GetEnableTrackingOk() (*bool, bool)`

GetEnableTrackingOk returns a tuple with the EnableTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTracking

`func (o *PublicPageViewAnalyticsFilter) SetEnableTracking(v bool)`

SetEnableTracking sets EnableTracking field to given value.

### HasEnableTracking

`func (o *PublicPageViewAnalyticsFilter) HasEnableTracking() bool`

HasEnableTracking returns a boolean if a field has been set.

### GetPruningRefineBy

`func (o *PublicPageViewAnalyticsFilter) GetPruningRefineBy() PublicFormSubmissionFilterCoalescingRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *PublicPageViewAnalyticsFilter) GetPruningRefineByOk() (*PublicFormSubmissionFilterCoalescingRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *PublicPageViewAnalyticsFilter) SetPruningRefineBy(v PublicFormSubmissionFilterCoalescingRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.

### HasPruningRefineBy

`func (o *PublicPageViewAnalyticsFilter) HasPruningRefineBy() bool`

HasPruningRefineBy returns a boolean if a field has been set.

### GetPageUrl

`func (o *PublicPageViewAnalyticsFilter) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *PublicPageViewAnalyticsFilter) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *PublicPageViewAnalyticsFilter) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.


### GetFilterType

`func (o *PublicPageViewAnalyticsFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicPageViewAnalyticsFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicPageViewAnalyticsFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetOperator

`func (o *PublicPageViewAnalyticsFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicPageViewAnalyticsFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicPageViewAnalyticsFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



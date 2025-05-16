# PublicEventAnalyticsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**CoalescingRefineBy** | Pointer to [**PublicFormSubmissionFilterCoalescingRefineBy**](PublicFormSubmissionFilterCoalescingRefineBy.md) |  | [optional] 
**PruningRefineBy** | Pointer to [**PublicFormSubmissionFilterCoalescingRefineBy**](PublicFormSubmissionFilterCoalescingRefineBy.md) |  | [optional] 
**FilterType** | **string** |  | [default to "EVENT"]
**Operator** | **string** |  | 

## Methods

### NewPublicEventAnalyticsFilter

`func NewPublicEventAnalyticsFilter(eventId string, filterType string, operator string, ) *PublicEventAnalyticsFilter`

NewPublicEventAnalyticsFilter instantiates a new PublicEventAnalyticsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicEventAnalyticsFilterWithDefaults

`func NewPublicEventAnalyticsFilterWithDefaults() *PublicEventAnalyticsFilter`

NewPublicEventAnalyticsFilterWithDefaults instantiates a new PublicEventAnalyticsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PublicEventAnalyticsFilter) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PublicEventAnalyticsFilter) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PublicEventAnalyticsFilter) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCoalescingRefineBy

`func (o *PublicEventAnalyticsFilter) GetCoalescingRefineBy() PublicFormSubmissionFilterCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *PublicEventAnalyticsFilter) GetCoalescingRefineByOk() (*PublicFormSubmissionFilterCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *PublicEventAnalyticsFilter) SetCoalescingRefineBy(v PublicFormSubmissionFilterCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.

### HasCoalescingRefineBy

`func (o *PublicEventAnalyticsFilter) HasCoalescingRefineBy() bool`

HasCoalescingRefineBy returns a boolean if a field has been set.

### GetPruningRefineBy

`func (o *PublicEventAnalyticsFilter) GetPruningRefineBy() PublicFormSubmissionFilterCoalescingRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *PublicEventAnalyticsFilter) GetPruningRefineByOk() (*PublicFormSubmissionFilterCoalescingRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *PublicEventAnalyticsFilter) SetPruningRefineBy(v PublicFormSubmissionFilterCoalescingRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.

### HasPruningRefineBy

`func (o *PublicEventAnalyticsFilter) HasPruningRefineBy() bool`

HasPruningRefineBy returns a boolean if a field has been set.

### GetFilterType

`func (o *PublicEventAnalyticsFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicEventAnalyticsFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicEventAnalyticsFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetOperator

`func (o *PublicEventAnalyticsFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicEventAnalyticsFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicEventAnalyticsFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



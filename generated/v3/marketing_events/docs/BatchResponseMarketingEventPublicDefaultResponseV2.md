# BatchResponseMarketingEventPublicDefaultResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** |  | 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 
**Results** | [**[]MarketingEventPublicDefaultResponseV2**](MarketingEventPublicDefaultResponseV2.md) |  | 
**Status** | **string** |  | 

## Methods

### NewBatchResponseMarketingEventPublicDefaultResponseV2

`func NewBatchResponseMarketingEventPublicDefaultResponseV2(completedAt time.Time, startedAt time.Time, results []MarketingEventPublicDefaultResponseV2, status string, ) *BatchResponseMarketingEventPublicDefaultResponseV2`

NewBatchResponseMarketingEventPublicDefaultResponseV2 instantiates a new BatchResponseMarketingEventPublicDefaultResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseMarketingEventPublicDefaultResponseV2WithDefaults

`func NewBatchResponseMarketingEventPublicDefaultResponseV2WithDefaults() *BatchResponseMarketingEventPublicDefaultResponseV2`

NewBatchResponseMarketingEventPublicDefaultResponseV2WithDefaults instantiates a new BatchResponseMarketingEventPublicDefaultResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetRequestedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetLinks

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) GetResults() []MarketingEventPublicDefaultResponseV2`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) GetResultsOk() (*[]MarketingEventPublicDefaultResponseV2, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) SetResults(v []MarketingEventPublicDefaultResponseV2)`

SetResults sets Results field to given value.


### GetStatus

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



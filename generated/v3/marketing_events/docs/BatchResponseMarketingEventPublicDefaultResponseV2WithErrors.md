# BatchResponseMarketingEventPublicDefaultResponseV2WithErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** |  | 
**NumErrors** | Pointer to **int32** |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 
**Results** | [**[]MarketingEventPublicDefaultResponseV2**](MarketingEventPublicDefaultResponseV2.md) |  | 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewBatchResponseMarketingEventPublicDefaultResponseV2WithErrors

`func NewBatchResponseMarketingEventPublicDefaultResponseV2WithErrors(completedAt time.Time, startedAt time.Time, results []MarketingEventPublicDefaultResponseV2, status string, ) *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors`

NewBatchResponseMarketingEventPublicDefaultResponseV2WithErrors instantiates a new BatchResponseMarketingEventPublicDefaultResponseV2WithErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseMarketingEventPublicDefaultResponseV2WithErrorsWithDefaults

`func NewBatchResponseMarketingEventPublicDefaultResponseV2WithErrorsWithDefaults() *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors`

NewBatchResponseMarketingEventPublicDefaultResponseV2WithErrorsWithDefaults instantiates a new BatchResponseMarketingEventPublicDefaultResponseV2WithErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetNumErrors

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetLinks

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetResults() []MarketingEventPublicDefaultResponseV2`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetResultsOk() (*[]MarketingEventPublicDefaultResponseV2, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) SetResults(v []MarketingEventPublicDefaultResponseV2)`

SetResults sets Results field to given value.


### GetErrors

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseMarketingEventPublicDefaultResponseV2WithErrors) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



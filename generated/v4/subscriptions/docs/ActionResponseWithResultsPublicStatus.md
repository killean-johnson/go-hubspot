# ActionResponseWithResultsPublicStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** |  | 
**NumErrors** | Pointer to **int32** |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 
**Results** | [**[]PublicStatus**](PublicStatus.md) |  | 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewActionResponseWithResultsPublicStatus

`func NewActionResponseWithResultsPublicStatus(completedAt time.Time, startedAt time.Time, results []PublicStatus, status string, ) *ActionResponseWithResultsPublicStatus`

NewActionResponseWithResultsPublicStatus instantiates a new ActionResponseWithResultsPublicStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionResponseWithResultsPublicStatusWithDefaults

`func NewActionResponseWithResultsPublicStatusWithDefaults() *ActionResponseWithResultsPublicStatus`

NewActionResponseWithResultsPublicStatusWithDefaults instantiates a new ActionResponseWithResultsPublicStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *ActionResponseWithResultsPublicStatus) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ActionResponseWithResultsPublicStatus) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ActionResponseWithResultsPublicStatus) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetNumErrors

`func (o *ActionResponseWithResultsPublicStatus) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *ActionResponseWithResultsPublicStatus) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *ActionResponseWithResultsPublicStatus) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *ActionResponseWithResultsPublicStatus) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *ActionResponseWithResultsPublicStatus) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *ActionResponseWithResultsPublicStatus) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *ActionResponseWithResultsPublicStatus) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *ActionResponseWithResultsPublicStatus) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *ActionResponseWithResultsPublicStatus) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ActionResponseWithResultsPublicStatus) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ActionResponseWithResultsPublicStatus) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetLinks

`func (o *ActionResponseWithResultsPublicStatus) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ActionResponseWithResultsPublicStatus) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ActionResponseWithResultsPublicStatus) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ActionResponseWithResultsPublicStatus) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *ActionResponseWithResultsPublicStatus) GetResults() []PublicStatus`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ActionResponseWithResultsPublicStatus) GetResultsOk() (*[]PublicStatus, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ActionResponseWithResultsPublicStatus) SetResults(v []PublicStatus)`

SetResults sets Results field to given value.


### GetErrors

`func (o *ActionResponseWithResultsPublicStatus) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ActionResponseWithResultsPublicStatus) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ActionResponseWithResultsPublicStatus) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ActionResponseWithResultsPublicStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *ActionResponseWithResultsPublicStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionResponseWithResultsPublicStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionResponseWithResultsPublicStatus) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



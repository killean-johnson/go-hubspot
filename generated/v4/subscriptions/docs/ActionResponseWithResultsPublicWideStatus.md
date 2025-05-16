# ActionResponseWithResultsPublicWideStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** |  | 
**NumErrors** | Pointer to **int32** |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 
**Results** | [**[]PublicWideStatus**](PublicWideStatus.md) |  | 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewActionResponseWithResultsPublicWideStatus

`func NewActionResponseWithResultsPublicWideStatus(completedAt time.Time, startedAt time.Time, results []PublicWideStatus, status string, ) *ActionResponseWithResultsPublicWideStatus`

NewActionResponseWithResultsPublicWideStatus instantiates a new ActionResponseWithResultsPublicWideStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionResponseWithResultsPublicWideStatusWithDefaults

`func NewActionResponseWithResultsPublicWideStatusWithDefaults() *ActionResponseWithResultsPublicWideStatus`

NewActionResponseWithResultsPublicWideStatusWithDefaults instantiates a new ActionResponseWithResultsPublicWideStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *ActionResponseWithResultsPublicWideStatus) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ActionResponseWithResultsPublicWideStatus) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ActionResponseWithResultsPublicWideStatus) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetNumErrors

`func (o *ActionResponseWithResultsPublicWideStatus) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *ActionResponseWithResultsPublicWideStatus) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *ActionResponseWithResultsPublicWideStatus) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *ActionResponseWithResultsPublicWideStatus) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *ActionResponseWithResultsPublicWideStatus) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *ActionResponseWithResultsPublicWideStatus) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *ActionResponseWithResultsPublicWideStatus) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *ActionResponseWithResultsPublicWideStatus) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *ActionResponseWithResultsPublicWideStatus) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ActionResponseWithResultsPublicWideStatus) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ActionResponseWithResultsPublicWideStatus) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetLinks

`func (o *ActionResponseWithResultsPublicWideStatus) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ActionResponseWithResultsPublicWideStatus) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ActionResponseWithResultsPublicWideStatus) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ActionResponseWithResultsPublicWideStatus) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *ActionResponseWithResultsPublicWideStatus) GetResults() []PublicWideStatus`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ActionResponseWithResultsPublicWideStatus) GetResultsOk() (*[]PublicWideStatus, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ActionResponseWithResultsPublicWideStatus) SetResults(v []PublicWideStatus)`

SetResults sets Results field to given value.


### GetErrors

`func (o *ActionResponseWithResultsPublicWideStatus) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ActionResponseWithResultsPublicWideStatus) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ActionResponseWithResultsPublicWideStatus) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ActionResponseWithResultsPublicWideStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *ActionResponseWithResultsPublicWideStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionResponseWithResultsPublicWideStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionResponseWithResultsPublicWideStatus) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ActionResponseWithSingleResultURI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** |  | [optional] 
**CompletedAt** | **time.Time** |  | 
**NumErrors** | Pointer to **int32** |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewActionResponseWithSingleResultURI

`func NewActionResponseWithSingleResultURI(completedAt time.Time, startedAt time.Time, status string, ) *ActionResponseWithSingleResultURI`

NewActionResponseWithSingleResultURI instantiates a new ActionResponseWithSingleResultURI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionResponseWithSingleResultURIWithDefaults

`func NewActionResponseWithSingleResultURIWithDefaults() *ActionResponseWithSingleResultURI`

NewActionResponseWithSingleResultURIWithDefaults instantiates a new ActionResponseWithSingleResultURI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ActionResponseWithSingleResultURI) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ActionResponseWithSingleResultURI) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ActionResponseWithSingleResultURI) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ActionResponseWithSingleResultURI) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ActionResponseWithSingleResultURI) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ActionResponseWithSingleResultURI) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ActionResponseWithSingleResultURI) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetNumErrors

`func (o *ActionResponseWithSingleResultURI) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *ActionResponseWithSingleResultURI) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *ActionResponseWithSingleResultURI) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *ActionResponseWithSingleResultURI) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *ActionResponseWithSingleResultURI) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *ActionResponseWithSingleResultURI) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *ActionResponseWithSingleResultURI) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *ActionResponseWithSingleResultURI) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *ActionResponseWithSingleResultURI) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ActionResponseWithSingleResultURI) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ActionResponseWithSingleResultURI) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetLinks

`func (o *ActionResponseWithSingleResultURI) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ActionResponseWithSingleResultURI) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ActionResponseWithSingleResultURI) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ActionResponseWithSingleResultURI) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetErrors

`func (o *ActionResponseWithSingleResultURI) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ActionResponseWithSingleResultURI) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ActionResponseWithSingleResultURI) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ActionResponseWithSingleResultURI) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *ActionResponseWithSingleResultURI) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionResponseWithSingleResultURI) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionResponseWithSingleResultURI) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



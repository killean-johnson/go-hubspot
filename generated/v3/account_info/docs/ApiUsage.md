# ApiUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectedAt** | **time.Time** | Indicates when the cache was last updated. | 
**UsageLimit** | **int32** | Limits by which a single integration can consume the HubSpot public APIs. | 
**FetchStatus** | **string** | Status of fetching the information, including if the data came from the cache. | 
**Name** | **string** | Name of the limit type. | 
**CurrentUsage** | **int32** | How many API calls an account has made for the current day. | 
**ResetsAt** | Pointer to **time.Time** | Time that the limit will reset. | [optional] 

## Methods

### NewApiUsage

`func NewApiUsage(collectedAt time.Time, usageLimit int32, fetchStatus string, name string, currentUsage int32, ) *ApiUsage`

NewApiUsage instantiates a new ApiUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUsageWithDefaults

`func NewApiUsageWithDefaults() *ApiUsage`

NewApiUsageWithDefaults instantiates a new ApiUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectedAt

`func (o *ApiUsage) GetCollectedAt() time.Time`

GetCollectedAt returns the CollectedAt field if non-nil, zero value otherwise.

### GetCollectedAtOk

`func (o *ApiUsage) GetCollectedAtOk() (*time.Time, bool)`

GetCollectedAtOk returns a tuple with the CollectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectedAt

`func (o *ApiUsage) SetCollectedAt(v time.Time)`

SetCollectedAt sets CollectedAt field to given value.


### GetUsageLimit

`func (o *ApiUsage) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *ApiUsage) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *ApiUsage) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.


### GetFetchStatus

`func (o *ApiUsage) GetFetchStatus() string`

GetFetchStatus returns the FetchStatus field if non-nil, zero value otherwise.

### GetFetchStatusOk

`func (o *ApiUsage) GetFetchStatusOk() (*string, bool)`

GetFetchStatusOk returns a tuple with the FetchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchStatus

`func (o *ApiUsage) SetFetchStatus(v string)`

SetFetchStatus sets FetchStatus field to given value.


### GetName

`func (o *ApiUsage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiUsage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiUsage) SetName(v string)`

SetName sets Name field to given value.


### GetCurrentUsage

`func (o *ApiUsage) GetCurrentUsage() int32`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *ApiUsage) GetCurrentUsageOk() (*int32, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *ApiUsage) SetCurrentUsage(v int32)`

SetCurrentUsage sets CurrentUsage field to given value.


### GetResetsAt

`func (o *ApiUsage) GetResetsAt() time.Time`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *ApiUsage) GetResetsAtOk() (*time.Time, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *ApiUsage) SetResetsAt(v time.Time)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *ApiUsage) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



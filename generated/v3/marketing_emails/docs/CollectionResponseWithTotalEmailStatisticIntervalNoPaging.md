# CollectionResponseWithTotalEmailStatisticIntervalNoPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of objects. | 
**Results** | [**[]EmailStatisticInterval**](EmailStatisticInterval.md) | Collection of objects. | 

## Methods

### NewCollectionResponseWithTotalEmailStatisticIntervalNoPaging

`func NewCollectionResponseWithTotalEmailStatisticIntervalNoPaging(total int32, results []EmailStatisticInterval, ) *CollectionResponseWithTotalEmailStatisticIntervalNoPaging`

NewCollectionResponseWithTotalEmailStatisticIntervalNoPaging instantiates a new CollectionResponseWithTotalEmailStatisticIntervalNoPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithTotalEmailStatisticIntervalNoPagingWithDefaults

`func NewCollectionResponseWithTotalEmailStatisticIntervalNoPagingWithDefaults() *CollectionResponseWithTotalEmailStatisticIntervalNoPaging`

NewCollectionResponseWithTotalEmailStatisticIntervalNoPagingWithDefaults instantiates a new CollectionResponseWithTotalEmailStatisticIntervalNoPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CollectionResponseWithTotalEmailStatisticIntervalNoPaging) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CollectionResponseWithTotalEmailStatisticIntervalNoPaging) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CollectionResponseWithTotalEmailStatisticIntervalNoPaging) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetResults

`func (o *CollectionResponseWithTotalEmailStatisticIntervalNoPaging) GetResults() []EmailStatisticInterval`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponseWithTotalEmailStatisticIntervalNoPaging) GetResultsOk() (*[]EmailStatisticInterval, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponseWithTotalEmailStatisticIntervalNoPaging) SetResults(v []EmailStatisticInterval)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



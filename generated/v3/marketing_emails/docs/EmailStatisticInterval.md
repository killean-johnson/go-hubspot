# EmailStatisticInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to [**Interval**](Interval.md) |  | [optional] 
**Aggregations** | Pointer to [**EmailStatisticsData**](EmailStatisticsData.md) |  | [optional] 

## Methods

### NewEmailStatisticInterval

`func NewEmailStatisticInterval() *EmailStatisticInterval`

NewEmailStatisticInterval instantiates a new EmailStatisticInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailStatisticIntervalWithDefaults

`func NewEmailStatisticIntervalWithDefaults() *EmailStatisticInterval`

NewEmailStatisticIntervalWithDefaults instantiates a new EmailStatisticInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *EmailStatisticInterval) GetInterval() Interval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *EmailStatisticInterval) GetIntervalOk() (*Interval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *EmailStatisticInterval) SetInterval(v Interval)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *EmailStatisticInterval) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetAggregations

`func (o *EmailStatisticInterval) GetAggregations() EmailStatisticsData`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *EmailStatisticInterval) GetAggregationsOk() (*EmailStatisticsData, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *EmailStatisticInterval) SetAggregations(v EmailStatisticsData)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *EmailStatisticInterval) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



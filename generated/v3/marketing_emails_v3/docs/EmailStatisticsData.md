# EmailStatisticsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceBreakdown** | **map[string]map[string]int64** | Statistics by device. | 
**QualifierStats** | **map[string]map[string]int64** | Number of emails that were dropped and bounced. | 
**Counters** | **map[string]int64** | Counters like number of &#x60;sent&#x60;, &#x60;open&#x60; or &#x60;delivered&#x60;. | 
**Ratios** | **map[string]float32** | Ratios like &#x60;openratio&#x60; or &#x60;clickratio&#x60; | 

## Methods

### NewEmailStatisticsData

`func NewEmailStatisticsData(deviceBreakdown map[string]map[string]int64, qualifierStats map[string]map[string]int64, counters map[string]int64, ratios map[string]float32, ) *EmailStatisticsData`

NewEmailStatisticsData instantiates a new EmailStatisticsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailStatisticsDataWithDefaults

`func NewEmailStatisticsDataWithDefaults() *EmailStatisticsData`

NewEmailStatisticsDataWithDefaults instantiates a new EmailStatisticsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceBreakdown

`func (o *EmailStatisticsData) GetDeviceBreakdown() map[string]map[string]int64`

GetDeviceBreakdown returns the DeviceBreakdown field if non-nil, zero value otherwise.

### GetDeviceBreakdownOk

`func (o *EmailStatisticsData) GetDeviceBreakdownOk() (*map[string]map[string]int64, bool)`

GetDeviceBreakdownOk returns a tuple with the DeviceBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceBreakdown

`func (o *EmailStatisticsData) SetDeviceBreakdown(v map[string]map[string]int64)`

SetDeviceBreakdown sets DeviceBreakdown field to given value.


### GetQualifierStats

`func (o *EmailStatisticsData) GetQualifierStats() map[string]map[string]int64`

GetQualifierStats returns the QualifierStats field if non-nil, zero value otherwise.

### GetQualifierStatsOk

`func (o *EmailStatisticsData) GetQualifierStatsOk() (*map[string]map[string]int64, bool)`

GetQualifierStatsOk returns a tuple with the QualifierStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifierStats

`func (o *EmailStatisticsData) SetQualifierStats(v map[string]map[string]int64)`

SetQualifierStats sets QualifierStats field to given value.


### GetCounters

`func (o *EmailStatisticsData) GetCounters() map[string]int64`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *EmailStatisticsData) GetCountersOk() (*map[string]int64, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *EmailStatisticsData) SetCounters(v map[string]int64)`

SetCounters sets Counters field to given value.


### GetRatios

`func (o *EmailStatisticsData) GetRatios() map[string]float32`

GetRatios returns the Ratios field if non-nil, zero value otherwise.

### GetRatiosOk

`func (o *EmailStatisticsData) GetRatiosOk() (*map[string]float32, bool)`

GetRatiosOk returns a tuple with the Ratios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatios

`func (o *EmailStatisticsData) SetRatios(v map[string]float32)`

SetRatios sets Ratios field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



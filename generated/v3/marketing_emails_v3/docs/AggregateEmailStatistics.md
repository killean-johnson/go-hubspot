# AggregateEmailStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to **[]int64** | List of email IDs that were sent during the time span. | [optional] 
**CampaignAggregations** | Pointer to [**map[string]EmailStatisticsData**](EmailStatisticsData.md) | The aggregated statistics per campaign. | [optional] 
**Aggregate** | Pointer to [**EmailStatisticsData**](EmailStatisticsData.md) |  | [optional] 

## Methods

### NewAggregateEmailStatistics

`func NewAggregateEmailStatistics() *AggregateEmailStatistics`

NewAggregateEmailStatistics instantiates a new AggregateEmailStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateEmailStatisticsWithDefaults

`func NewAggregateEmailStatisticsWithDefaults() *AggregateEmailStatistics`

NewAggregateEmailStatisticsWithDefaults instantiates a new AggregateEmailStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *AggregateEmailStatistics) GetEmails() []int64`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *AggregateEmailStatistics) GetEmailsOk() (*[]int64, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *AggregateEmailStatistics) SetEmails(v []int64)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *AggregateEmailStatistics) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetCampaignAggregations

`func (o *AggregateEmailStatistics) GetCampaignAggregations() map[string]EmailStatisticsData`

GetCampaignAggregations returns the CampaignAggregations field if non-nil, zero value otherwise.

### GetCampaignAggregationsOk

`func (o *AggregateEmailStatistics) GetCampaignAggregationsOk() (*map[string]EmailStatisticsData, bool)`

GetCampaignAggregationsOk returns a tuple with the CampaignAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignAggregations

`func (o *AggregateEmailStatistics) SetCampaignAggregations(v map[string]EmailStatisticsData)`

SetCampaignAggregations sets CampaignAggregations field to given value.

### HasCampaignAggregations

`func (o *AggregateEmailStatistics) HasCampaignAggregations() bool`

HasCampaignAggregations returns a boolean if a field has been set.

### GetAggregate

`func (o *AggregateEmailStatistics) GetAggregate() EmailStatisticsData`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *AggregateEmailStatistics) GetAggregateOk() (*EmailStatisticsData, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *AggregateEmailStatistics) SetAggregate(v EmailStatisticsData)`

SetAggregate sets Aggregate field to given value.

### HasAggregate

`func (o *AggregateEmailStatistics) HasAggregate() bool`

HasAggregate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



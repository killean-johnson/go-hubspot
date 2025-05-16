# AssociationRecordLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRecordsAtLimit** | **int32** |  | 
**AtLimitFromRecordSamples** | [**[]AtLimitRecordSample**](AtLimitRecordSample.md) |  | 
**TotalRecordsNearLimit** | **int32** |  | 
**Limit** | **int64** |  | 
**NearLimitFromRecordSamples** | [**[]NearLimitRecordSample**](NearLimitRecordSample.md) |  | 

## Methods

### NewAssociationRecordLimitResponse

`func NewAssociationRecordLimitResponse(totalRecordsAtLimit int32, atLimitFromRecordSamples []AtLimitRecordSample, totalRecordsNearLimit int32, limit int64, nearLimitFromRecordSamples []NearLimitRecordSample, ) *AssociationRecordLimitResponse`

NewAssociationRecordLimitResponse instantiates a new AssociationRecordLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationRecordLimitResponseWithDefaults

`func NewAssociationRecordLimitResponseWithDefaults() *AssociationRecordLimitResponse`

NewAssociationRecordLimitResponseWithDefaults instantiates a new AssociationRecordLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRecordsAtLimit

`func (o *AssociationRecordLimitResponse) GetTotalRecordsAtLimit() int32`

GetTotalRecordsAtLimit returns the TotalRecordsAtLimit field if non-nil, zero value otherwise.

### GetTotalRecordsAtLimitOk

`func (o *AssociationRecordLimitResponse) GetTotalRecordsAtLimitOk() (*int32, bool)`

GetTotalRecordsAtLimitOk returns a tuple with the TotalRecordsAtLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordsAtLimit

`func (o *AssociationRecordLimitResponse) SetTotalRecordsAtLimit(v int32)`

SetTotalRecordsAtLimit sets TotalRecordsAtLimit field to given value.


### GetAtLimitFromRecordSamples

`func (o *AssociationRecordLimitResponse) GetAtLimitFromRecordSamples() []AtLimitRecordSample`

GetAtLimitFromRecordSamples returns the AtLimitFromRecordSamples field if non-nil, zero value otherwise.

### GetAtLimitFromRecordSamplesOk

`func (o *AssociationRecordLimitResponse) GetAtLimitFromRecordSamplesOk() (*[]AtLimitRecordSample, bool)`

GetAtLimitFromRecordSamplesOk returns a tuple with the AtLimitFromRecordSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtLimitFromRecordSamples

`func (o *AssociationRecordLimitResponse) SetAtLimitFromRecordSamples(v []AtLimitRecordSample)`

SetAtLimitFromRecordSamples sets AtLimitFromRecordSamples field to given value.


### GetTotalRecordsNearLimit

`func (o *AssociationRecordLimitResponse) GetTotalRecordsNearLimit() int32`

GetTotalRecordsNearLimit returns the TotalRecordsNearLimit field if non-nil, zero value otherwise.

### GetTotalRecordsNearLimitOk

`func (o *AssociationRecordLimitResponse) GetTotalRecordsNearLimitOk() (*int32, bool)`

GetTotalRecordsNearLimitOk returns a tuple with the TotalRecordsNearLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordsNearLimit

`func (o *AssociationRecordLimitResponse) SetTotalRecordsNearLimit(v int32)`

SetTotalRecordsNearLimit sets TotalRecordsNearLimit field to given value.


### GetLimit

`func (o *AssociationRecordLimitResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AssociationRecordLimitResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AssociationRecordLimitResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetNearLimitFromRecordSamples

`func (o *AssociationRecordLimitResponse) GetNearLimitFromRecordSamples() []NearLimitRecordSample`

GetNearLimitFromRecordSamples returns the NearLimitFromRecordSamples field if non-nil, zero value otherwise.

### GetNearLimitFromRecordSamplesOk

`func (o *AssociationRecordLimitResponse) GetNearLimitFromRecordSamplesOk() (*[]NearLimitRecordSample, bool)`

GetNearLimitFromRecordSamplesOk returns a tuple with the NearLimitFromRecordSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearLimitFromRecordSamples

`func (o *AssociationRecordLimitResponse) SetNearLimitFromRecordSamples(v []NearLimitRecordSample)`

SetNearLimitFromRecordSamples sets NearLimitFromRecordSamples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



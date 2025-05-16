# CustomObjectRecordLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverallUsage** | **int32** |  | 
**ByObjectType** | [**[]UsageForObjectType**](UsageForObjectType.md) |  | 
**OverallLimit** | **int32** |  | 
**OverallPercentage** | **float32** |  | 

## Methods

### NewCustomObjectRecordLimitResponse

`func NewCustomObjectRecordLimitResponse(overallUsage int32, byObjectType []UsageForObjectType, overallLimit int32, overallPercentage float32, ) *CustomObjectRecordLimitResponse`

NewCustomObjectRecordLimitResponse instantiates a new CustomObjectRecordLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomObjectRecordLimitResponseWithDefaults

`func NewCustomObjectRecordLimitResponseWithDefaults() *CustomObjectRecordLimitResponse`

NewCustomObjectRecordLimitResponseWithDefaults instantiates a new CustomObjectRecordLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverallUsage

`func (o *CustomObjectRecordLimitResponse) GetOverallUsage() int32`

GetOverallUsage returns the OverallUsage field if non-nil, zero value otherwise.

### GetOverallUsageOk

`func (o *CustomObjectRecordLimitResponse) GetOverallUsageOk() (*int32, bool)`

GetOverallUsageOk returns a tuple with the OverallUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallUsage

`func (o *CustomObjectRecordLimitResponse) SetOverallUsage(v int32)`

SetOverallUsage sets OverallUsage field to given value.


### GetByObjectType

`func (o *CustomObjectRecordLimitResponse) GetByObjectType() []UsageForObjectType`

GetByObjectType returns the ByObjectType field if non-nil, zero value otherwise.

### GetByObjectTypeOk

`func (o *CustomObjectRecordLimitResponse) GetByObjectTypeOk() (*[]UsageForObjectType, bool)`

GetByObjectTypeOk returns a tuple with the ByObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByObjectType

`func (o *CustomObjectRecordLimitResponse) SetByObjectType(v []UsageForObjectType)`

SetByObjectType sets ByObjectType field to given value.


### GetOverallLimit

`func (o *CustomObjectRecordLimitResponse) GetOverallLimit() int32`

GetOverallLimit returns the OverallLimit field if non-nil, zero value otherwise.

### GetOverallLimitOk

`func (o *CustomObjectRecordLimitResponse) GetOverallLimitOk() (*int32, bool)`

GetOverallLimitOk returns a tuple with the OverallLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallLimit

`func (o *CustomObjectRecordLimitResponse) SetOverallLimit(v int32)`

SetOverallLimit sets OverallLimit field to given value.


### GetOverallPercentage

`func (o *CustomObjectRecordLimitResponse) GetOverallPercentage() float32`

GetOverallPercentage returns the OverallPercentage field if non-nil, zero value otherwise.

### GetOverallPercentageOk

`func (o *CustomObjectRecordLimitResponse) GetOverallPercentageOk() (*float32, bool)`

GetOverallPercentageOk returns a tuple with the OverallPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallPercentage

`func (o *CustomObjectRecordLimitResponse) SetOverallPercentage(v float32)`

SetOverallPercentage sets OverallPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CalculatedPropertyLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverallUsage** | **int64** |  | 
**ByObjectType** | [**[]UsageForObjectType**](UsageForObjectType.md) |  | 
**OverallLimit** | **int64** |  | 
**OverallPercentage** | **float32** |  | 

## Methods

### NewCalculatedPropertyLimitResponse

`func NewCalculatedPropertyLimitResponse(overallUsage int64, byObjectType []UsageForObjectType, overallLimit int64, overallPercentage float32, ) *CalculatedPropertyLimitResponse`

NewCalculatedPropertyLimitResponse instantiates a new CalculatedPropertyLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculatedPropertyLimitResponseWithDefaults

`func NewCalculatedPropertyLimitResponseWithDefaults() *CalculatedPropertyLimitResponse`

NewCalculatedPropertyLimitResponseWithDefaults instantiates a new CalculatedPropertyLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverallUsage

`func (o *CalculatedPropertyLimitResponse) GetOverallUsage() int64`

GetOverallUsage returns the OverallUsage field if non-nil, zero value otherwise.

### GetOverallUsageOk

`func (o *CalculatedPropertyLimitResponse) GetOverallUsageOk() (*int64, bool)`

GetOverallUsageOk returns a tuple with the OverallUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallUsage

`func (o *CalculatedPropertyLimitResponse) SetOverallUsage(v int64)`

SetOverallUsage sets OverallUsage field to given value.


### GetByObjectType

`func (o *CalculatedPropertyLimitResponse) GetByObjectType() []UsageForObjectType`

GetByObjectType returns the ByObjectType field if non-nil, zero value otherwise.

### GetByObjectTypeOk

`func (o *CalculatedPropertyLimitResponse) GetByObjectTypeOk() (*[]UsageForObjectType, bool)`

GetByObjectTypeOk returns a tuple with the ByObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByObjectType

`func (o *CalculatedPropertyLimitResponse) SetByObjectType(v []UsageForObjectType)`

SetByObjectType sets ByObjectType field to given value.


### GetOverallLimit

`func (o *CalculatedPropertyLimitResponse) GetOverallLimit() int64`

GetOverallLimit returns the OverallLimit field if non-nil, zero value otherwise.

### GetOverallLimitOk

`func (o *CalculatedPropertyLimitResponse) GetOverallLimitOk() (*int64, bool)`

GetOverallLimitOk returns a tuple with the OverallLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallLimit

`func (o *CalculatedPropertyLimitResponse) SetOverallLimit(v int64)`

SetOverallLimit sets OverallLimit field to given value.


### GetOverallPercentage

`func (o *CalculatedPropertyLimitResponse) GetOverallPercentage() float32`

GetOverallPercentage returns the OverallPercentage field if non-nil, zero value otherwise.

### GetOverallPercentageOk

`func (o *CalculatedPropertyLimitResponse) GetOverallPercentageOk() (*float32, bool)`

GetOverallPercentageOk returns a tuple with the OverallPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallPercentage

`func (o *CalculatedPropertyLimitResponse) SetOverallPercentage(v float32)`

SetOverallPercentage sets OverallPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



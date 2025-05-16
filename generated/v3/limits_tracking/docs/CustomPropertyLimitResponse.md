# CustomPropertyLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverallUsage** | **int64** |  | 
**ByObjectType** | [**[]LimitAndUsageForObjectType**](LimitAndUsageForObjectType.md) |  | 
**OverallLimit** | **int64** |  | 
**OverallPercentage** | **float32** |  | 

## Methods

### NewCustomPropertyLimitResponse

`func NewCustomPropertyLimitResponse(overallUsage int64, byObjectType []LimitAndUsageForObjectType, overallLimit int64, overallPercentage float32, ) *CustomPropertyLimitResponse`

NewCustomPropertyLimitResponse instantiates a new CustomPropertyLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPropertyLimitResponseWithDefaults

`func NewCustomPropertyLimitResponseWithDefaults() *CustomPropertyLimitResponse`

NewCustomPropertyLimitResponseWithDefaults instantiates a new CustomPropertyLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverallUsage

`func (o *CustomPropertyLimitResponse) GetOverallUsage() int64`

GetOverallUsage returns the OverallUsage field if non-nil, zero value otherwise.

### GetOverallUsageOk

`func (o *CustomPropertyLimitResponse) GetOverallUsageOk() (*int64, bool)`

GetOverallUsageOk returns a tuple with the OverallUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallUsage

`func (o *CustomPropertyLimitResponse) SetOverallUsage(v int64)`

SetOverallUsage sets OverallUsage field to given value.


### GetByObjectType

`func (o *CustomPropertyLimitResponse) GetByObjectType() []LimitAndUsageForObjectType`

GetByObjectType returns the ByObjectType field if non-nil, zero value otherwise.

### GetByObjectTypeOk

`func (o *CustomPropertyLimitResponse) GetByObjectTypeOk() (*[]LimitAndUsageForObjectType, bool)`

GetByObjectTypeOk returns a tuple with the ByObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByObjectType

`func (o *CustomPropertyLimitResponse) SetByObjectType(v []LimitAndUsageForObjectType)`

SetByObjectType sets ByObjectType field to given value.


### GetOverallLimit

`func (o *CustomPropertyLimitResponse) GetOverallLimit() int64`

GetOverallLimit returns the OverallLimit field if non-nil, zero value otherwise.

### GetOverallLimitOk

`func (o *CustomPropertyLimitResponse) GetOverallLimitOk() (*int64, bool)`

GetOverallLimitOk returns a tuple with the OverallLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallLimit

`func (o *CustomPropertyLimitResponse) SetOverallLimit(v int64)`

SetOverallLimit sets OverallLimit field to given value.


### GetOverallPercentage

`func (o *CustomPropertyLimitResponse) GetOverallPercentage() float32`

GetOverallPercentage returns the OverallPercentage field if non-nil, zero value otherwise.

### GetOverallPercentageOk

`func (o *CustomPropertyLimitResponse) GetOverallPercentageOk() (*float32, bool)`

GetOverallPercentageOk returns a tuple with the OverallPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallPercentage

`func (o *CustomPropertyLimitResponse) SetOverallPercentage(v float32)`

SetOverallPercentage sets OverallPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



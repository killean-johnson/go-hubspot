# CustomObjectLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usage** | **int64** |  | 
**Percentage** | **float32** |  | 
**Limit** | **int64** |  | 

## Methods

### NewCustomObjectLimitResponse

`func NewCustomObjectLimitResponse(usage int64, percentage float32, limit int64, ) *CustomObjectLimitResponse`

NewCustomObjectLimitResponse instantiates a new CustomObjectLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomObjectLimitResponseWithDefaults

`func NewCustomObjectLimitResponseWithDefaults() *CustomObjectLimitResponse`

NewCustomObjectLimitResponseWithDefaults instantiates a new CustomObjectLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsage

`func (o *CustomObjectLimitResponse) GetUsage() int64`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CustomObjectLimitResponse) GetUsageOk() (*int64, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CustomObjectLimitResponse) SetUsage(v int64)`

SetUsage sets Usage field to given value.


### GetPercentage

`func (o *CustomObjectLimitResponse) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *CustomObjectLimitResponse) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *CustomObjectLimitResponse) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.


### GetLimit

`func (o *CustomObjectLimitResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CustomObjectLimitResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CustomObjectLimitResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



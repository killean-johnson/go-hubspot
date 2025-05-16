# AllPropertyTypesOperationPruningRefineBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | **string** |  | 
**TimeOffset** | [**TimeOffset**](TimeOffset.md) |  | 
**Type** | **string** |  | 
**UpperBoundOffset** | [**TimeOffset**](TimeOffset.md) |  | 
**LowerBoundOffset** | [**TimeOffset**](TimeOffset.md) |  | 
**RangeType** | **string** |  | 
**Timestamp** | **int32** |  | 
**UpperTimestamp** | **int32** |  | 
**LowerTimestamp** | **int32** |  | 
**EndpointBehavior** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**PropertyType** | **string** |  | [default to "rangedtime"]
**PropertyParser** | **string** |  | 
**TimePoint** | [**RangedTimeOperationUpperBoundTimePoint**](RangedTimeOperationUpperBoundTimePoint.md) |  | 
**OperationType** | **string** |  | 
**OperatorName** | **string** |  | 
**Operator** | **string** |  | 
**UpperBoundEndpointBehavior** | **string** |  | 
**UpperBoundTimePoint** | [**RangedTimeOperationUpperBoundTimePoint**](RangedTimeOperationUpperBoundTimePoint.md) |  | 
**LowerBoundEndpointBehavior** | **string** |  | 
**LowerBoundTimePoint** | [**RangedTimeOperationUpperBoundTimePoint**](RangedTimeOperationUpperBoundTimePoint.md) |  | 

## Methods

### NewAllPropertyTypesOperationPruningRefineBy

`func NewAllPropertyTypesOperationPruningRefineBy(comparison string, timeOffset TimeOffset, type_ string, upperBoundOffset TimeOffset, lowerBoundOffset TimeOffset, rangeType string, timestamp int32, upperTimestamp int32, lowerTimestamp int32, endpointBehavior string, includeObjectsWithNoValueSet bool, propertyType string, propertyParser string, timePoint RangedTimeOperationUpperBoundTimePoint, operationType string, operatorName string, operator string, upperBoundEndpointBehavior string, upperBoundTimePoint RangedTimeOperationUpperBoundTimePoint, lowerBoundEndpointBehavior string, lowerBoundTimePoint RangedTimeOperationUpperBoundTimePoint, ) *AllPropertyTypesOperationPruningRefineBy`

NewAllPropertyTypesOperationPruningRefineBy instantiates a new AllPropertyTypesOperationPruningRefineBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllPropertyTypesOperationPruningRefineByWithDefaults

`func NewAllPropertyTypesOperationPruningRefineByWithDefaults() *AllPropertyTypesOperationPruningRefineBy`

NewAllPropertyTypesOperationPruningRefineByWithDefaults instantiates a new AllPropertyTypesOperationPruningRefineBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *AllPropertyTypesOperationPruningRefineBy) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *AllPropertyTypesOperationPruningRefineBy) SetComparison(v string)`

SetComparison sets Comparison field to given value.


### GetTimeOffset

`func (o *AllPropertyTypesOperationPruningRefineBy) GetTimeOffset() TimeOffset`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetTimeOffsetOk() (*TimeOffset, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *AllPropertyTypesOperationPruningRefineBy) SetTimeOffset(v TimeOffset)`

SetTimeOffset sets TimeOffset field to given value.


### GetType

`func (o *AllPropertyTypesOperationPruningRefineBy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AllPropertyTypesOperationPruningRefineBy) SetType(v string)`

SetType sets Type field to given value.


### GetUpperBoundOffset

`func (o *AllPropertyTypesOperationPruningRefineBy) GetUpperBoundOffset() TimeOffset`

GetUpperBoundOffset returns the UpperBoundOffset field if non-nil, zero value otherwise.

### GetUpperBoundOffsetOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetUpperBoundOffsetOk() (*TimeOffset, bool)`

GetUpperBoundOffsetOk returns a tuple with the UpperBoundOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundOffset

`func (o *AllPropertyTypesOperationPruningRefineBy) SetUpperBoundOffset(v TimeOffset)`

SetUpperBoundOffset sets UpperBoundOffset field to given value.


### GetLowerBoundOffset

`func (o *AllPropertyTypesOperationPruningRefineBy) GetLowerBoundOffset() TimeOffset`

GetLowerBoundOffset returns the LowerBoundOffset field if non-nil, zero value otherwise.

### GetLowerBoundOffsetOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetLowerBoundOffsetOk() (*TimeOffset, bool)`

GetLowerBoundOffsetOk returns a tuple with the LowerBoundOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundOffset

`func (o *AllPropertyTypesOperationPruningRefineBy) SetLowerBoundOffset(v TimeOffset)`

SetLowerBoundOffset sets LowerBoundOffset field to given value.


### GetRangeType

`func (o *AllPropertyTypesOperationPruningRefineBy) GetRangeType() string`

GetRangeType returns the RangeType field if non-nil, zero value otherwise.

### GetRangeTypeOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetRangeTypeOk() (*string, bool)`

GetRangeTypeOk returns a tuple with the RangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeType

`func (o *AllPropertyTypesOperationPruningRefineBy) SetRangeType(v string)`

SetRangeType sets RangeType field to given value.


### GetTimestamp

`func (o *AllPropertyTypesOperationPruningRefineBy) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AllPropertyTypesOperationPruningRefineBy) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetUpperTimestamp

`func (o *AllPropertyTypesOperationPruningRefineBy) GetUpperTimestamp() int32`

GetUpperTimestamp returns the UpperTimestamp field if non-nil, zero value otherwise.

### GetUpperTimestampOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetUpperTimestampOk() (*int32, bool)`

GetUpperTimestampOk returns a tuple with the UpperTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperTimestamp

`func (o *AllPropertyTypesOperationPruningRefineBy) SetUpperTimestamp(v int32)`

SetUpperTimestamp sets UpperTimestamp field to given value.


### GetLowerTimestamp

`func (o *AllPropertyTypesOperationPruningRefineBy) GetLowerTimestamp() int32`

GetLowerTimestamp returns the LowerTimestamp field if non-nil, zero value otherwise.

### GetLowerTimestampOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetLowerTimestampOk() (*int32, bool)`

GetLowerTimestampOk returns a tuple with the LowerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerTimestamp

`func (o *AllPropertyTypesOperationPruningRefineBy) SetLowerTimestamp(v int32)`

SetLowerTimestamp sets LowerTimestamp field to given value.


### GetEndpointBehavior

`func (o *AllPropertyTypesOperationPruningRefineBy) GetEndpointBehavior() string`

GetEndpointBehavior returns the EndpointBehavior field if non-nil, zero value otherwise.

### GetEndpointBehaviorOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetEndpointBehaviorOk() (*string, bool)`

GetEndpointBehaviorOk returns a tuple with the EndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointBehavior

`func (o *AllPropertyTypesOperationPruningRefineBy) SetEndpointBehavior(v string)`

SetEndpointBehavior sets EndpointBehavior field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *AllPropertyTypesOperationPruningRefineBy) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *AllPropertyTypesOperationPruningRefineBy) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetDefaultValue

`func (o *AllPropertyTypesOperationPruningRefineBy) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *AllPropertyTypesOperationPruningRefineBy) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *AllPropertyTypesOperationPruningRefineBy) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetPropertyType

`func (o *AllPropertyTypesOperationPruningRefineBy) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *AllPropertyTypesOperationPruningRefineBy) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.


### GetPropertyParser

`func (o *AllPropertyTypesOperationPruningRefineBy) GetPropertyParser() string`

GetPropertyParser returns the PropertyParser field if non-nil, zero value otherwise.

### GetPropertyParserOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetPropertyParserOk() (*string, bool)`

GetPropertyParserOk returns a tuple with the PropertyParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyParser

`func (o *AllPropertyTypesOperationPruningRefineBy) SetPropertyParser(v string)`

SetPropertyParser sets PropertyParser field to given value.


### GetTimePoint

`func (o *AllPropertyTypesOperationPruningRefineBy) GetTimePoint() RangedTimeOperationUpperBoundTimePoint`

GetTimePoint returns the TimePoint field if non-nil, zero value otherwise.

### GetTimePointOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetTimePointOk() (*RangedTimeOperationUpperBoundTimePoint, bool)`

GetTimePointOk returns a tuple with the TimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePoint

`func (o *AllPropertyTypesOperationPruningRefineBy) SetTimePoint(v RangedTimeOperationUpperBoundTimePoint)`

SetTimePoint sets TimePoint field to given value.


### GetOperationType

`func (o *AllPropertyTypesOperationPruningRefineBy) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *AllPropertyTypesOperationPruningRefineBy) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperatorName

`func (o *AllPropertyTypesOperationPruningRefineBy) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *AllPropertyTypesOperationPruningRefineBy) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetOperator

`func (o *AllPropertyTypesOperationPruningRefineBy) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AllPropertyTypesOperationPruningRefineBy) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetUpperBoundEndpointBehavior

`func (o *AllPropertyTypesOperationPruningRefineBy) GetUpperBoundEndpointBehavior() string`

GetUpperBoundEndpointBehavior returns the UpperBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetUpperBoundEndpointBehaviorOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetUpperBoundEndpointBehaviorOk() (*string, bool)`

GetUpperBoundEndpointBehaviorOk returns a tuple with the UpperBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundEndpointBehavior

`func (o *AllPropertyTypesOperationPruningRefineBy) SetUpperBoundEndpointBehavior(v string)`

SetUpperBoundEndpointBehavior sets UpperBoundEndpointBehavior field to given value.


### GetUpperBoundTimePoint

`func (o *AllPropertyTypesOperationPruningRefineBy) GetUpperBoundTimePoint() RangedTimeOperationUpperBoundTimePoint`

GetUpperBoundTimePoint returns the UpperBoundTimePoint field if non-nil, zero value otherwise.

### GetUpperBoundTimePointOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetUpperBoundTimePointOk() (*RangedTimeOperationUpperBoundTimePoint, bool)`

GetUpperBoundTimePointOk returns a tuple with the UpperBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundTimePoint

`func (o *AllPropertyTypesOperationPruningRefineBy) SetUpperBoundTimePoint(v RangedTimeOperationUpperBoundTimePoint)`

SetUpperBoundTimePoint sets UpperBoundTimePoint field to given value.


### GetLowerBoundEndpointBehavior

`func (o *AllPropertyTypesOperationPruningRefineBy) GetLowerBoundEndpointBehavior() string`

GetLowerBoundEndpointBehavior returns the LowerBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetLowerBoundEndpointBehaviorOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetLowerBoundEndpointBehaviorOk() (*string, bool)`

GetLowerBoundEndpointBehaviorOk returns a tuple with the LowerBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundEndpointBehavior

`func (o *AllPropertyTypesOperationPruningRefineBy) SetLowerBoundEndpointBehavior(v string)`

SetLowerBoundEndpointBehavior sets LowerBoundEndpointBehavior field to given value.


### GetLowerBoundTimePoint

`func (o *AllPropertyTypesOperationPruningRefineBy) GetLowerBoundTimePoint() RangedTimeOperationUpperBoundTimePoint`

GetLowerBoundTimePoint returns the LowerBoundTimePoint field if non-nil, zero value otherwise.

### GetLowerBoundTimePointOk

`func (o *AllPropertyTypesOperationPruningRefineBy) GetLowerBoundTimePointOk() (*RangedTimeOperationUpperBoundTimePoint, bool)`

GetLowerBoundTimePointOk returns a tuple with the LowerBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundTimePoint

`func (o *AllPropertyTypesOperationPruningRefineBy) SetLowerBoundTimePoint(v RangedTimeOperationUpperBoundTimePoint)`

SetLowerBoundTimePoint sets LowerBoundTimePoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



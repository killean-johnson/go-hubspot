# PublicFormSubmissionFilterCoalescingRefineBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxOccurrences** | Pointer to **int32** |  | [optional] 
**Type** | **string** |  | 
**MinOccurrences** | Pointer to **int32** |  | [optional] 
**SetType** | **string** |  | 
**Comparison** | **string** |  | 
**TimeOffset** | [**PublicTimeOffset**](PublicTimeOffset.md) |  | 
**UpperBoundOffset** | [**PublicTimeOffset**](PublicTimeOffset.md) |  | 
**RangeType** | **string** |  | 
**LowerBoundOffset** | [**PublicTimeOffset**](PublicTimeOffset.md) |  | 
**Timestamp** | **int64** |  | 
**UpperTimestamp** | **int64** |  | 
**LowerTimestamp** | **int64** |  | 
**EndpointBehavior** | Pointer to **string** |  | [optional] 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**PropertyParser** | Pointer to **string** |  | [optional] 
**OperationType** | **string** |  | 
**TimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 
**Operator** | **string** |  | 
**UpperBoundEndpointBehavior** | Pointer to **string** |  | [optional] 
**UpperBoundTimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 
**LowerBoundEndpointBehavior** | Pointer to **string** |  | [optional] 
**LowerBoundTimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 

## Methods

### NewPublicFormSubmissionFilterCoalescingRefineBy

`func NewPublicFormSubmissionFilterCoalescingRefineBy(type_ string, setType string, comparison string, timeOffset PublicTimeOffset, upperBoundOffset PublicTimeOffset, rangeType string, lowerBoundOffset PublicTimeOffset, timestamp int64, upperTimestamp int64, lowerTimestamp int64, includeObjectsWithNoValueSet bool, operationType string, timePoint PublicTimePointOperationTimePoint, operator string, upperBoundTimePoint PublicTimePointOperationTimePoint, lowerBoundTimePoint PublicTimePointOperationTimePoint, ) *PublicFormSubmissionFilterCoalescingRefineBy`

NewPublicFormSubmissionFilterCoalescingRefineBy instantiates a new PublicFormSubmissionFilterCoalescingRefineBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicFormSubmissionFilterCoalescingRefineByWithDefaults

`func NewPublicFormSubmissionFilterCoalescingRefineByWithDefaults() *PublicFormSubmissionFilterCoalescingRefineBy`

NewPublicFormSubmissionFilterCoalescingRefineByWithDefaults instantiates a new PublicFormSubmissionFilterCoalescingRefineBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxOccurrences

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetMaxOccurrences() int32`

GetMaxOccurrences returns the MaxOccurrences field if non-nil, zero value otherwise.

### GetMaxOccurrencesOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetMaxOccurrencesOk() (*int32, bool)`

GetMaxOccurrencesOk returns a tuple with the MaxOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOccurrences

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetMaxOccurrences(v int32)`

SetMaxOccurrences sets MaxOccurrences field to given value.

### HasMaxOccurrences

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) HasMaxOccurrences() bool`

HasMaxOccurrences returns a boolean if a field has been set.

### GetType

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetType(v string)`

SetType sets Type field to given value.


### GetMinOccurrences

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetMinOccurrences() int32`

GetMinOccurrences returns the MinOccurrences field if non-nil, zero value otherwise.

### GetMinOccurrencesOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetMinOccurrencesOk() (*int32, bool)`

GetMinOccurrencesOk returns a tuple with the MinOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOccurrences

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetMinOccurrences(v int32)`

SetMinOccurrences sets MinOccurrences field to given value.

### HasMinOccurrences

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) HasMinOccurrences() bool`

HasMinOccurrences returns a boolean if a field has been set.

### GetSetType

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetSetType() string`

GetSetType returns the SetType field if non-nil, zero value otherwise.

### GetSetTypeOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetSetTypeOk() (*string, bool)`

GetSetTypeOk returns a tuple with the SetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetType

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetSetType(v string)`

SetSetType sets SetType field to given value.


### GetComparison

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetComparison(v string)`

SetComparison sets Comparison field to given value.


### GetTimeOffset

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetTimeOffset() PublicTimeOffset`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetTimeOffsetOk() (*PublicTimeOffset, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetTimeOffset(v PublicTimeOffset)`

SetTimeOffset sets TimeOffset field to given value.


### GetUpperBoundOffset

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetUpperBoundOffset() PublicTimeOffset`

GetUpperBoundOffset returns the UpperBoundOffset field if non-nil, zero value otherwise.

### GetUpperBoundOffsetOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetUpperBoundOffsetOk() (*PublicTimeOffset, bool)`

GetUpperBoundOffsetOk returns a tuple with the UpperBoundOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundOffset

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetUpperBoundOffset(v PublicTimeOffset)`

SetUpperBoundOffset sets UpperBoundOffset field to given value.


### GetRangeType

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetRangeType() string`

GetRangeType returns the RangeType field if non-nil, zero value otherwise.

### GetRangeTypeOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetRangeTypeOk() (*string, bool)`

GetRangeTypeOk returns a tuple with the RangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeType

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetRangeType(v string)`

SetRangeType sets RangeType field to given value.


### GetLowerBoundOffset

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetLowerBoundOffset() PublicTimeOffset`

GetLowerBoundOffset returns the LowerBoundOffset field if non-nil, zero value otherwise.

### GetLowerBoundOffsetOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetLowerBoundOffsetOk() (*PublicTimeOffset, bool)`

GetLowerBoundOffsetOk returns a tuple with the LowerBoundOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundOffset

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetLowerBoundOffset(v PublicTimeOffset)`

SetLowerBoundOffset sets LowerBoundOffset field to given value.


### GetTimestamp

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetUpperTimestamp

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetUpperTimestamp() int64`

GetUpperTimestamp returns the UpperTimestamp field if non-nil, zero value otherwise.

### GetUpperTimestampOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetUpperTimestampOk() (*int64, bool)`

GetUpperTimestampOk returns a tuple with the UpperTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperTimestamp

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetUpperTimestamp(v int64)`

SetUpperTimestamp sets UpperTimestamp field to given value.


### GetLowerTimestamp

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetLowerTimestamp() int64`

GetLowerTimestamp returns the LowerTimestamp field if non-nil, zero value otherwise.

### GetLowerTimestampOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetLowerTimestampOk() (*int64, bool)`

GetLowerTimestampOk returns a tuple with the LowerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerTimestamp

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetLowerTimestamp(v int64)`

SetLowerTimestamp sets LowerTimestamp field to given value.


### GetEndpointBehavior

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetEndpointBehavior() string`

GetEndpointBehavior returns the EndpointBehavior field if non-nil, zero value otherwise.

### GetEndpointBehaviorOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetEndpointBehaviorOk() (*string, bool)`

GetEndpointBehaviorOk returns a tuple with the EndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointBehavior

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetEndpointBehavior(v string)`

SetEndpointBehavior sets EndpointBehavior field to given value.

### HasEndpointBehavior

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) HasEndpointBehavior() bool`

HasEndpointBehavior returns a boolean if a field has been set.

### GetIncludeObjectsWithNoValueSet

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetPropertyParser

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetPropertyParser() string`

GetPropertyParser returns the PropertyParser field if non-nil, zero value otherwise.

### GetPropertyParserOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetPropertyParserOk() (*string, bool)`

GetPropertyParserOk returns a tuple with the PropertyParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyParser

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetPropertyParser(v string)`

SetPropertyParser sets PropertyParser field to given value.

### HasPropertyParser

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) HasPropertyParser() bool`

HasPropertyParser returns a boolean if a field has been set.

### GetOperationType

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetTimePoint

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetTimePoint() PublicTimePointOperationTimePoint`

GetTimePoint returns the TimePoint field if non-nil, zero value otherwise.

### GetTimePointOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetTimePointOk returns a tuple with the TimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePoint

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetTimePoint(v PublicTimePointOperationTimePoint)`

SetTimePoint sets TimePoint field to given value.


### GetOperator

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetUpperBoundEndpointBehavior

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetUpperBoundEndpointBehavior() string`

GetUpperBoundEndpointBehavior returns the UpperBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetUpperBoundEndpointBehaviorOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetUpperBoundEndpointBehaviorOk() (*string, bool)`

GetUpperBoundEndpointBehaviorOk returns a tuple with the UpperBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundEndpointBehavior

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetUpperBoundEndpointBehavior(v string)`

SetUpperBoundEndpointBehavior sets UpperBoundEndpointBehavior field to given value.

### HasUpperBoundEndpointBehavior

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) HasUpperBoundEndpointBehavior() bool`

HasUpperBoundEndpointBehavior returns a boolean if a field has been set.

### GetUpperBoundTimePoint

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetUpperBoundTimePoint() PublicTimePointOperationTimePoint`

GetUpperBoundTimePoint returns the UpperBoundTimePoint field if non-nil, zero value otherwise.

### GetUpperBoundTimePointOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetUpperBoundTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetUpperBoundTimePointOk returns a tuple with the UpperBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundTimePoint

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetUpperBoundTimePoint(v PublicTimePointOperationTimePoint)`

SetUpperBoundTimePoint sets UpperBoundTimePoint field to given value.


### GetLowerBoundEndpointBehavior

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetLowerBoundEndpointBehavior() string`

GetLowerBoundEndpointBehavior returns the LowerBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetLowerBoundEndpointBehaviorOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetLowerBoundEndpointBehaviorOk() (*string, bool)`

GetLowerBoundEndpointBehaviorOk returns a tuple with the LowerBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundEndpointBehavior

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetLowerBoundEndpointBehavior(v string)`

SetLowerBoundEndpointBehavior sets LowerBoundEndpointBehavior field to given value.

### HasLowerBoundEndpointBehavior

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) HasLowerBoundEndpointBehavior() bool`

HasLowerBoundEndpointBehavior returns a boolean if a field has been set.

### GetLowerBoundTimePoint

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetLowerBoundTimePoint() PublicTimePointOperationTimePoint`

GetLowerBoundTimePoint returns the LowerBoundTimePoint field if non-nil, zero value otherwise.

### GetLowerBoundTimePointOk

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) GetLowerBoundTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetLowerBoundTimePointOk returns a tuple with the LowerBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundTimePoint

`func (o *PublicFormSubmissionFilterCoalescingRefineBy) SetLowerBoundTimePoint(v PublicTimePointOperationTimePoint)`

SetLowerBoundTimePoint sets LowerBoundTimePoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PublicSurveyMonkeyValueFilterValueComparison

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**OperationType** | **string** |  | 
**Value** | **string** |  | 
**Operator** | **string** |  | 
**RequiresTimeZoneConversion** | **bool** |  | 
**Timestamp** | **int32** |  | 
**UpperBound** | **int32** |  | 
**LowerBound** | **int32** |  | 
**DefaultComparisonValue** | Pointer to **string** |  | [optional] 
**ComparisonPropertyName** | **string** |  | 
**NumberOfDays** | **int32** |  | 
**Values** | **[]string** |  | 
**Month** | **string** |  | 
**Year** | **int32** |  | 
**Day** | **int32** |  | 
**UseFiscalYear** | Pointer to **bool** |  | [optional] 
**FiscalYearStart** | Pointer to **string** |  | [optional] 
**TimeUnitCount** | Pointer to **int32** |  | [optional] 
**TimeUnit** | **string** |  | 
**EndpointBehavior** | Pointer to **string** |  | [optional] 
**PropertyParser** | Pointer to **string** |  | [optional] 
**TimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 
**Type** | **string** |  | 
**UpperBoundEndpointBehavior** | Pointer to **string** |  | [optional] 
**UpperBoundTimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 
**LowerBoundEndpointBehavior** | Pointer to **string** |  | [optional] 
**LowerBoundTimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 

## Methods

### NewPublicSurveyMonkeyValueFilterValueComparison

`func NewPublicSurveyMonkeyValueFilterValueComparison(includeObjectsWithNoValueSet bool, operationType string, value string, operator string, requiresTimeZoneConversion bool, timestamp int32, upperBound int32, lowerBound int32, comparisonPropertyName string, numberOfDays int32, values []string, month string, year int32, day int32, timeUnit string, timePoint PublicTimePointOperationTimePoint, type_ string, upperBoundTimePoint PublicTimePointOperationTimePoint, lowerBoundTimePoint PublicTimePointOperationTimePoint, ) *PublicSurveyMonkeyValueFilterValueComparison`

NewPublicSurveyMonkeyValueFilterValueComparison instantiates a new PublicSurveyMonkeyValueFilterValueComparison object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSurveyMonkeyValueFilterValueComparisonWithDefaults

`func NewPublicSurveyMonkeyValueFilterValueComparisonWithDefaults() *PublicSurveyMonkeyValueFilterValueComparison`

NewPublicSurveyMonkeyValueFilterValueComparisonWithDefaults instantiates a new PublicSurveyMonkeyValueFilterValueComparison object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeObjectsWithNoValueSet

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetOperationType

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetValue

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetValue(v string)`

SetValue sets Value field to given value.


### GetOperator

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetRequiresTimeZoneConversion

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetRequiresTimeZoneConversion() bool`

GetRequiresTimeZoneConversion returns the RequiresTimeZoneConversion field if non-nil, zero value otherwise.

### GetRequiresTimeZoneConversionOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetRequiresTimeZoneConversionOk() (*bool, bool)`

GetRequiresTimeZoneConversionOk returns a tuple with the RequiresTimeZoneConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTimeZoneConversion

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetRequiresTimeZoneConversion(v bool)`

SetRequiresTimeZoneConversion sets RequiresTimeZoneConversion field to given value.


### GetTimestamp

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetUpperBound

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetUpperBound() int32`

GetUpperBound returns the UpperBound field if non-nil, zero value otherwise.

### GetUpperBoundOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetUpperBoundOk() (*int32, bool)`

GetUpperBoundOk returns a tuple with the UpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBound

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetUpperBound(v int32)`

SetUpperBound sets UpperBound field to given value.


### GetLowerBound

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetLowerBound() int32`

GetLowerBound returns the LowerBound field if non-nil, zero value otherwise.

### GetLowerBoundOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetLowerBoundOk() (*int32, bool)`

GetLowerBoundOk returns a tuple with the LowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBound

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetLowerBound(v int32)`

SetLowerBound sets LowerBound field to given value.


### GetDefaultComparisonValue

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetDefaultComparisonValue() string`

GetDefaultComparisonValue returns the DefaultComparisonValue field if non-nil, zero value otherwise.

### GetDefaultComparisonValueOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetDefaultComparisonValueOk() (*string, bool)`

GetDefaultComparisonValueOk returns a tuple with the DefaultComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultComparisonValue

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetDefaultComparisonValue(v string)`

SetDefaultComparisonValue sets DefaultComparisonValue field to given value.

### HasDefaultComparisonValue

`func (o *PublicSurveyMonkeyValueFilterValueComparison) HasDefaultComparisonValue() bool`

HasDefaultComparisonValue returns a boolean if a field has been set.

### GetComparisonPropertyName

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetComparisonPropertyName() string`

GetComparisonPropertyName returns the ComparisonPropertyName field if non-nil, zero value otherwise.

### GetComparisonPropertyNameOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetComparisonPropertyNameOk() (*string, bool)`

GetComparisonPropertyNameOk returns a tuple with the ComparisonPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonPropertyName

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetComparisonPropertyName(v string)`

SetComparisonPropertyName sets ComparisonPropertyName field to given value.


### GetNumberOfDays

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetNumberOfDays() int32`

GetNumberOfDays returns the NumberOfDays field if non-nil, zero value otherwise.

### GetNumberOfDaysOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetNumberOfDaysOk() (*int32, bool)`

GetNumberOfDaysOk returns a tuple with the NumberOfDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDays

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetNumberOfDays(v int32)`

SetNumberOfDays sets NumberOfDays field to given value.


### GetValues

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetValues(v []string)`

SetValues sets Values field to given value.


### GetMonth

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetYear

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetYear(v int32)`

SetYear sets Year field to given value.


### GetDay

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetDay(v int32)`

SetDay sets Day field to given value.


### GetUseFiscalYear

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetUseFiscalYear() bool`

GetUseFiscalYear returns the UseFiscalYear field if non-nil, zero value otherwise.

### GetUseFiscalYearOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetUseFiscalYearOk() (*bool, bool)`

GetUseFiscalYearOk returns a tuple with the UseFiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFiscalYear

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetUseFiscalYear(v bool)`

SetUseFiscalYear sets UseFiscalYear field to given value.

### HasUseFiscalYear

`func (o *PublicSurveyMonkeyValueFilterValueComparison) HasUseFiscalYear() bool`

HasUseFiscalYear returns a boolean if a field has been set.

### GetFiscalYearStart

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetFiscalYearStart() string`

GetFiscalYearStart returns the FiscalYearStart field if non-nil, zero value otherwise.

### GetFiscalYearStartOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetFiscalYearStartOk() (*string, bool)`

GetFiscalYearStartOk returns a tuple with the FiscalYearStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearStart

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetFiscalYearStart(v string)`

SetFiscalYearStart sets FiscalYearStart field to given value.

### HasFiscalYearStart

`func (o *PublicSurveyMonkeyValueFilterValueComparison) HasFiscalYearStart() bool`

HasFiscalYearStart returns a boolean if a field has been set.

### GetTimeUnitCount

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetTimeUnitCount() int32`

GetTimeUnitCount returns the TimeUnitCount field if non-nil, zero value otherwise.

### GetTimeUnitCountOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetTimeUnitCountOk() (*int32, bool)`

GetTimeUnitCountOk returns a tuple with the TimeUnitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnitCount

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetTimeUnitCount(v int32)`

SetTimeUnitCount sets TimeUnitCount field to given value.

### HasTimeUnitCount

`func (o *PublicSurveyMonkeyValueFilterValueComparison) HasTimeUnitCount() bool`

HasTimeUnitCount returns a boolean if a field has been set.

### GetTimeUnit

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.


### GetEndpointBehavior

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetEndpointBehavior() string`

GetEndpointBehavior returns the EndpointBehavior field if non-nil, zero value otherwise.

### GetEndpointBehaviorOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetEndpointBehaviorOk() (*string, bool)`

GetEndpointBehaviorOk returns a tuple with the EndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointBehavior

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetEndpointBehavior(v string)`

SetEndpointBehavior sets EndpointBehavior field to given value.

### HasEndpointBehavior

`func (o *PublicSurveyMonkeyValueFilterValueComparison) HasEndpointBehavior() bool`

HasEndpointBehavior returns a boolean if a field has been set.

### GetPropertyParser

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetPropertyParser() string`

GetPropertyParser returns the PropertyParser field if non-nil, zero value otherwise.

### GetPropertyParserOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetPropertyParserOk() (*string, bool)`

GetPropertyParserOk returns a tuple with the PropertyParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyParser

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetPropertyParser(v string)`

SetPropertyParser sets PropertyParser field to given value.

### HasPropertyParser

`func (o *PublicSurveyMonkeyValueFilterValueComparison) HasPropertyParser() bool`

HasPropertyParser returns a boolean if a field has been set.

### GetTimePoint

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetTimePoint() PublicTimePointOperationTimePoint`

GetTimePoint returns the TimePoint field if non-nil, zero value otherwise.

### GetTimePointOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetTimePointOk returns a tuple with the TimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePoint

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetTimePoint(v PublicTimePointOperationTimePoint)`

SetTimePoint sets TimePoint field to given value.


### GetType

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetType(v string)`

SetType sets Type field to given value.


### GetUpperBoundEndpointBehavior

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetUpperBoundEndpointBehavior() string`

GetUpperBoundEndpointBehavior returns the UpperBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetUpperBoundEndpointBehaviorOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetUpperBoundEndpointBehaviorOk() (*string, bool)`

GetUpperBoundEndpointBehaviorOk returns a tuple with the UpperBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundEndpointBehavior

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetUpperBoundEndpointBehavior(v string)`

SetUpperBoundEndpointBehavior sets UpperBoundEndpointBehavior field to given value.

### HasUpperBoundEndpointBehavior

`func (o *PublicSurveyMonkeyValueFilterValueComparison) HasUpperBoundEndpointBehavior() bool`

HasUpperBoundEndpointBehavior returns a boolean if a field has been set.

### GetUpperBoundTimePoint

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetUpperBoundTimePoint() PublicTimePointOperationTimePoint`

GetUpperBoundTimePoint returns the UpperBoundTimePoint field if non-nil, zero value otherwise.

### GetUpperBoundTimePointOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetUpperBoundTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetUpperBoundTimePointOk returns a tuple with the UpperBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundTimePoint

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetUpperBoundTimePoint(v PublicTimePointOperationTimePoint)`

SetUpperBoundTimePoint sets UpperBoundTimePoint field to given value.


### GetLowerBoundEndpointBehavior

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetLowerBoundEndpointBehavior() string`

GetLowerBoundEndpointBehavior returns the LowerBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetLowerBoundEndpointBehaviorOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetLowerBoundEndpointBehaviorOk() (*string, bool)`

GetLowerBoundEndpointBehaviorOk returns a tuple with the LowerBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundEndpointBehavior

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetLowerBoundEndpointBehavior(v string)`

SetLowerBoundEndpointBehavior sets LowerBoundEndpointBehavior field to given value.

### HasLowerBoundEndpointBehavior

`func (o *PublicSurveyMonkeyValueFilterValueComparison) HasLowerBoundEndpointBehavior() bool`

HasLowerBoundEndpointBehavior returns a boolean if a field has been set.

### GetLowerBoundTimePoint

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetLowerBoundTimePoint() PublicTimePointOperationTimePoint`

GetLowerBoundTimePoint returns the LowerBoundTimePoint field if non-nil, zero value otherwise.

### GetLowerBoundTimePointOk

`func (o *PublicSurveyMonkeyValueFilterValueComparison) GetLowerBoundTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetLowerBoundTimePointOk returns a tuple with the LowerBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundTimePoint

`func (o *PublicSurveyMonkeyValueFilterValueComparison) SetLowerBoundTimePoint(v PublicTimePointOperationTimePoint)`

SetLowerBoundTimePoint sets LowerBoundTimePoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



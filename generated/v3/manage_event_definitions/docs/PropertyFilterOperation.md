# PropertyFilterOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**PropertyType** | **string** |  | [default to "rangedtime"]
**OperationType** | **string** |  | 
**Value** | **string** |  | 
**OperatorName** | **string** |  | 
**Operator** | **string** |  | 
**RequiresTimeZoneConversion** | **bool** |  | 
**Timestamp** | **int32** |  | 
**LowerBoundTimestamp** | **int32** |  | 
**UpperBoundTimestamp** | **int32** |  | 
**DefaultComparisonValue** | Pointer to **string** |  | [optional] 
**ComparisonPropertyName** | **string** |  | 
**NumberOfDays** | **int32** |  | 
**Values** | **[]string** |  | 
**CoalescingRefineBy** | [**AllPropertyTypesOperationCoalescingRefineBy**](AllPropertyTypesOperationCoalescingRefineBy.md) |  | 
**PruningRefineBy** | Pointer to [**AllPropertyTypesOperationPruningRefineBy**](AllPropertyTypesOperationPruningRefineBy.md) |  | [optional] 
**UpperBound** | **int32** |  | 
**LowerBound** | **int32** |  | 
**Month** | **string** |  | 
**Year** | **int32** |  | 
**Day** | **int32** |  | 
**UseFiscalYear** | **bool** |  | 
**FiscalYearStart** | Pointer to **string** |  | [optional] 
**TimeUnitCount** | **int32** |  | 
**TimeUnit** | **string** |  | 
**EndpointBehavior** | **string** |  | 
**PropertyParser** | **string** |  | 
**TimePoint** | [**RangedTimeOperationUpperBoundTimePoint**](RangedTimeOperationUpperBoundTimePoint.md) |  | 
**Type** | **string** |  | 
**UpperBoundEndpointBehavior** | **string** |  | 
**UpperBoundTimePoint** | [**RangedTimeOperationUpperBoundTimePoint**](RangedTimeOperationUpperBoundTimePoint.md) |  | 
**LowerBoundEndpointBehavior** | **string** |  | 
**LowerBoundTimePoint** | [**RangedTimeOperationUpperBoundTimePoint**](RangedTimeOperationUpperBoundTimePoint.md) |  | 

## Methods

### NewPropertyFilterOperation

`func NewPropertyFilterOperation(includeObjectsWithNoValueSet bool, propertyType string, operationType string, value string, operatorName string, operator string, requiresTimeZoneConversion bool, timestamp int32, lowerBoundTimestamp int32, upperBoundTimestamp int32, comparisonPropertyName string, numberOfDays int32, values []string, coalescingRefineBy AllPropertyTypesOperationCoalescingRefineBy, upperBound int32, lowerBound int32, month string, year int32, day int32, useFiscalYear bool, timeUnitCount int32, timeUnit string, endpointBehavior string, propertyParser string, timePoint RangedTimeOperationUpperBoundTimePoint, type_ string, upperBoundEndpointBehavior string, upperBoundTimePoint RangedTimeOperationUpperBoundTimePoint, lowerBoundEndpointBehavior string, lowerBoundTimePoint RangedTimeOperationUpperBoundTimePoint, ) *PropertyFilterOperation`

NewPropertyFilterOperation instantiates a new PropertyFilterOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyFilterOperationWithDefaults

`func NewPropertyFilterOperationWithDefaults() *PropertyFilterOperation`

NewPropertyFilterOperationWithDefaults instantiates a new PropertyFilterOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeObjectsWithNoValueSet

`func (o *PropertyFilterOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PropertyFilterOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PropertyFilterOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetDefaultValue

`func (o *PropertyFilterOperation) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *PropertyFilterOperation) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *PropertyFilterOperation) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *PropertyFilterOperation) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetPropertyType

`func (o *PropertyFilterOperation) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *PropertyFilterOperation) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *PropertyFilterOperation) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.


### GetOperationType

`func (o *PropertyFilterOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PropertyFilterOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PropertyFilterOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetValue

`func (o *PropertyFilterOperation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PropertyFilterOperation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PropertyFilterOperation) SetValue(v string)`

SetValue sets Value field to given value.


### GetOperatorName

`func (o *PropertyFilterOperation) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *PropertyFilterOperation) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *PropertyFilterOperation) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetOperator

`func (o *PropertyFilterOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PropertyFilterOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PropertyFilterOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetRequiresTimeZoneConversion

`func (o *PropertyFilterOperation) GetRequiresTimeZoneConversion() bool`

GetRequiresTimeZoneConversion returns the RequiresTimeZoneConversion field if non-nil, zero value otherwise.

### GetRequiresTimeZoneConversionOk

`func (o *PropertyFilterOperation) GetRequiresTimeZoneConversionOk() (*bool, bool)`

GetRequiresTimeZoneConversionOk returns a tuple with the RequiresTimeZoneConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTimeZoneConversion

`func (o *PropertyFilterOperation) SetRequiresTimeZoneConversion(v bool)`

SetRequiresTimeZoneConversion sets RequiresTimeZoneConversion field to given value.


### GetTimestamp

`func (o *PropertyFilterOperation) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PropertyFilterOperation) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PropertyFilterOperation) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetLowerBoundTimestamp

`func (o *PropertyFilterOperation) GetLowerBoundTimestamp() int32`

GetLowerBoundTimestamp returns the LowerBoundTimestamp field if non-nil, zero value otherwise.

### GetLowerBoundTimestampOk

`func (o *PropertyFilterOperation) GetLowerBoundTimestampOk() (*int32, bool)`

GetLowerBoundTimestampOk returns a tuple with the LowerBoundTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundTimestamp

`func (o *PropertyFilterOperation) SetLowerBoundTimestamp(v int32)`

SetLowerBoundTimestamp sets LowerBoundTimestamp field to given value.


### GetUpperBoundTimestamp

`func (o *PropertyFilterOperation) GetUpperBoundTimestamp() int32`

GetUpperBoundTimestamp returns the UpperBoundTimestamp field if non-nil, zero value otherwise.

### GetUpperBoundTimestampOk

`func (o *PropertyFilterOperation) GetUpperBoundTimestampOk() (*int32, bool)`

GetUpperBoundTimestampOk returns a tuple with the UpperBoundTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundTimestamp

`func (o *PropertyFilterOperation) SetUpperBoundTimestamp(v int32)`

SetUpperBoundTimestamp sets UpperBoundTimestamp field to given value.


### GetDefaultComparisonValue

`func (o *PropertyFilterOperation) GetDefaultComparisonValue() string`

GetDefaultComparisonValue returns the DefaultComparisonValue field if non-nil, zero value otherwise.

### GetDefaultComparisonValueOk

`func (o *PropertyFilterOperation) GetDefaultComparisonValueOk() (*string, bool)`

GetDefaultComparisonValueOk returns a tuple with the DefaultComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultComparisonValue

`func (o *PropertyFilterOperation) SetDefaultComparisonValue(v string)`

SetDefaultComparisonValue sets DefaultComparisonValue field to given value.

### HasDefaultComparisonValue

`func (o *PropertyFilterOperation) HasDefaultComparisonValue() bool`

HasDefaultComparisonValue returns a boolean if a field has been set.

### GetComparisonPropertyName

`func (o *PropertyFilterOperation) GetComparisonPropertyName() string`

GetComparisonPropertyName returns the ComparisonPropertyName field if non-nil, zero value otherwise.

### GetComparisonPropertyNameOk

`func (o *PropertyFilterOperation) GetComparisonPropertyNameOk() (*string, bool)`

GetComparisonPropertyNameOk returns a tuple with the ComparisonPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonPropertyName

`func (o *PropertyFilterOperation) SetComparisonPropertyName(v string)`

SetComparisonPropertyName sets ComparisonPropertyName field to given value.


### GetNumberOfDays

`func (o *PropertyFilterOperation) GetNumberOfDays() int32`

GetNumberOfDays returns the NumberOfDays field if non-nil, zero value otherwise.

### GetNumberOfDaysOk

`func (o *PropertyFilterOperation) GetNumberOfDaysOk() (*int32, bool)`

GetNumberOfDaysOk returns a tuple with the NumberOfDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDays

`func (o *PropertyFilterOperation) SetNumberOfDays(v int32)`

SetNumberOfDays sets NumberOfDays field to given value.


### GetValues

`func (o *PropertyFilterOperation) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PropertyFilterOperation) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PropertyFilterOperation) SetValues(v []string)`

SetValues sets Values field to given value.


### GetCoalescingRefineBy

`func (o *PropertyFilterOperation) GetCoalescingRefineBy() AllPropertyTypesOperationCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *PropertyFilterOperation) GetCoalescingRefineByOk() (*AllPropertyTypesOperationCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *PropertyFilterOperation) SetCoalescingRefineBy(v AllPropertyTypesOperationCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.


### GetPruningRefineBy

`func (o *PropertyFilterOperation) GetPruningRefineBy() AllPropertyTypesOperationPruningRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *PropertyFilterOperation) GetPruningRefineByOk() (*AllPropertyTypesOperationPruningRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *PropertyFilterOperation) SetPruningRefineBy(v AllPropertyTypesOperationPruningRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.

### HasPruningRefineBy

`func (o *PropertyFilterOperation) HasPruningRefineBy() bool`

HasPruningRefineBy returns a boolean if a field has been set.

### GetUpperBound

`func (o *PropertyFilterOperation) GetUpperBound() int32`

GetUpperBound returns the UpperBound field if non-nil, zero value otherwise.

### GetUpperBoundOk

`func (o *PropertyFilterOperation) GetUpperBoundOk() (*int32, bool)`

GetUpperBoundOk returns a tuple with the UpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBound

`func (o *PropertyFilterOperation) SetUpperBound(v int32)`

SetUpperBound sets UpperBound field to given value.


### GetLowerBound

`func (o *PropertyFilterOperation) GetLowerBound() int32`

GetLowerBound returns the LowerBound field if non-nil, zero value otherwise.

### GetLowerBoundOk

`func (o *PropertyFilterOperation) GetLowerBoundOk() (*int32, bool)`

GetLowerBoundOk returns a tuple with the LowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBound

`func (o *PropertyFilterOperation) SetLowerBound(v int32)`

SetLowerBound sets LowerBound field to given value.


### GetMonth

`func (o *PropertyFilterOperation) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *PropertyFilterOperation) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *PropertyFilterOperation) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetYear

`func (o *PropertyFilterOperation) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *PropertyFilterOperation) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *PropertyFilterOperation) SetYear(v int32)`

SetYear sets Year field to given value.


### GetDay

`func (o *PropertyFilterOperation) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *PropertyFilterOperation) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *PropertyFilterOperation) SetDay(v int32)`

SetDay sets Day field to given value.


### GetUseFiscalYear

`func (o *PropertyFilterOperation) GetUseFiscalYear() bool`

GetUseFiscalYear returns the UseFiscalYear field if non-nil, zero value otherwise.

### GetUseFiscalYearOk

`func (o *PropertyFilterOperation) GetUseFiscalYearOk() (*bool, bool)`

GetUseFiscalYearOk returns a tuple with the UseFiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFiscalYear

`func (o *PropertyFilterOperation) SetUseFiscalYear(v bool)`

SetUseFiscalYear sets UseFiscalYear field to given value.


### GetFiscalYearStart

`func (o *PropertyFilterOperation) GetFiscalYearStart() string`

GetFiscalYearStart returns the FiscalYearStart field if non-nil, zero value otherwise.

### GetFiscalYearStartOk

`func (o *PropertyFilterOperation) GetFiscalYearStartOk() (*string, bool)`

GetFiscalYearStartOk returns a tuple with the FiscalYearStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearStart

`func (o *PropertyFilterOperation) SetFiscalYearStart(v string)`

SetFiscalYearStart sets FiscalYearStart field to given value.

### HasFiscalYearStart

`func (o *PropertyFilterOperation) HasFiscalYearStart() bool`

HasFiscalYearStart returns a boolean if a field has been set.

### GetTimeUnitCount

`func (o *PropertyFilterOperation) GetTimeUnitCount() int32`

GetTimeUnitCount returns the TimeUnitCount field if non-nil, zero value otherwise.

### GetTimeUnitCountOk

`func (o *PropertyFilterOperation) GetTimeUnitCountOk() (*int32, bool)`

GetTimeUnitCountOk returns a tuple with the TimeUnitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnitCount

`func (o *PropertyFilterOperation) SetTimeUnitCount(v int32)`

SetTimeUnitCount sets TimeUnitCount field to given value.


### GetTimeUnit

`func (o *PropertyFilterOperation) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *PropertyFilterOperation) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *PropertyFilterOperation) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.


### GetEndpointBehavior

`func (o *PropertyFilterOperation) GetEndpointBehavior() string`

GetEndpointBehavior returns the EndpointBehavior field if non-nil, zero value otherwise.

### GetEndpointBehaviorOk

`func (o *PropertyFilterOperation) GetEndpointBehaviorOk() (*string, bool)`

GetEndpointBehaviorOk returns a tuple with the EndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointBehavior

`func (o *PropertyFilterOperation) SetEndpointBehavior(v string)`

SetEndpointBehavior sets EndpointBehavior field to given value.


### GetPropertyParser

`func (o *PropertyFilterOperation) GetPropertyParser() string`

GetPropertyParser returns the PropertyParser field if non-nil, zero value otherwise.

### GetPropertyParserOk

`func (o *PropertyFilterOperation) GetPropertyParserOk() (*string, bool)`

GetPropertyParserOk returns a tuple with the PropertyParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyParser

`func (o *PropertyFilterOperation) SetPropertyParser(v string)`

SetPropertyParser sets PropertyParser field to given value.


### GetTimePoint

`func (o *PropertyFilterOperation) GetTimePoint() RangedTimeOperationUpperBoundTimePoint`

GetTimePoint returns the TimePoint field if non-nil, zero value otherwise.

### GetTimePointOk

`func (o *PropertyFilterOperation) GetTimePointOk() (*RangedTimeOperationUpperBoundTimePoint, bool)`

GetTimePointOk returns a tuple with the TimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePoint

`func (o *PropertyFilterOperation) SetTimePoint(v RangedTimeOperationUpperBoundTimePoint)`

SetTimePoint sets TimePoint field to given value.


### GetType

`func (o *PropertyFilterOperation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertyFilterOperation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertyFilterOperation) SetType(v string)`

SetType sets Type field to given value.


### GetUpperBoundEndpointBehavior

`func (o *PropertyFilterOperation) GetUpperBoundEndpointBehavior() string`

GetUpperBoundEndpointBehavior returns the UpperBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetUpperBoundEndpointBehaviorOk

`func (o *PropertyFilterOperation) GetUpperBoundEndpointBehaviorOk() (*string, bool)`

GetUpperBoundEndpointBehaviorOk returns a tuple with the UpperBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundEndpointBehavior

`func (o *PropertyFilterOperation) SetUpperBoundEndpointBehavior(v string)`

SetUpperBoundEndpointBehavior sets UpperBoundEndpointBehavior field to given value.


### GetUpperBoundTimePoint

`func (o *PropertyFilterOperation) GetUpperBoundTimePoint() RangedTimeOperationUpperBoundTimePoint`

GetUpperBoundTimePoint returns the UpperBoundTimePoint field if non-nil, zero value otherwise.

### GetUpperBoundTimePointOk

`func (o *PropertyFilterOperation) GetUpperBoundTimePointOk() (*RangedTimeOperationUpperBoundTimePoint, bool)`

GetUpperBoundTimePointOk returns a tuple with the UpperBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundTimePoint

`func (o *PropertyFilterOperation) SetUpperBoundTimePoint(v RangedTimeOperationUpperBoundTimePoint)`

SetUpperBoundTimePoint sets UpperBoundTimePoint field to given value.


### GetLowerBoundEndpointBehavior

`func (o *PropertyFilterOperation) GetLowerBoundEndpointBehavior() string`

GetLowerBoundEndpointBehavior returns the LowerBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetLowerBoundEndpointBehaviorOk

`func (o *PropertyFilterOperation) GetLowerBoundEndpointBehaviorOk() (*string, bool)`

GetLowerBoundEndpointBehaviorOk returns a tuple with the LowerBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundEndpointBehavior

`func (o *PropertyFilterOperation) SetLowerBoundEndpointBehavior(v string)`

SetLowerBoundEndpointBehavior sets LowerBoundEndpointBehavior field to given value.


### GetLowerBoundTimePoint

`func (o *PropertyFilterOperation) GetLowerBoundTimePoint() RangedTimeOperationUpperBoundTimePoint`

GetLowerBoundTimePoint returns the LowerBoundTimePoint field if non-nil, zero value otherwise.

### GetLowerBoundTimePointOk

`func (o *PropertyFilterOperation) GetLowerBoundTimePointOk() (*RangedTimeOperationUpperBoundTimePoint, bool)`

GetLowerBoundTimePointOk returns a tuple with the LowerBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundTimePoint

`func (o *PropertyFilterOperation) SetLowerBoundTimePoint(v RangedTimeOperationUpperBoundTimePoint)`

SetLowerBoundTimePoint sets LowerBoundTimePoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



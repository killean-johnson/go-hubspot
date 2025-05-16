# RangedDatePropertyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LowerBoundTimestamp** | **int32** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**PropertyType** | **string** |  | [default to "datetime-ranged"]
**UpperBoundTimestamp** | **int32** |  | 
**RequiresTimeZoneConversion** | **bool** |  | 
**OperationType** | **string** |  | 
**OperatorName** | **string** |  | 
**Operator** | **string** |  | 

## Methods

### NewRangedDatePropertyOperation

`func NewRangedDatePropertyOperation(lowerBoundTimestamp int32, includeObjectsWithNoValueSet bool, propertyType string, upperBoundTimestamp int32, requiresTimeZoneConversion bool, operationType string, operatorName string, operator string, ) *RangedDatePropertyOperation`

NewRangedDatePropertyOperation instantiates a new RangedDatePropertyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangedDatePropertyOperationWithDefaults

`func NewRangedDatePropertyOperationWithDefaults() *RangedDatePropertyOperation`

NewRangedDatePropertyOperationWithDefaults instantiates a new RangedDatePropertyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLowerBoundTimestamp

`func (o *RangedDatePropertyOperation) GetLowerBoundTimestamp() int32`

GetLowerBoundTimestamp returns the LowerBoundTimestamp field if non-nil, zero value otherwise.

### GetLowerBoundTimestampOk

`func (o *RangedDatePropertyOperation) GetLowerBoundTimestampOk() (*int32, bool)`

GetLowerBoundTimestampOk returns a tuple with the LowerBoundTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundTimestamp

`func (o *RangedDatePropertyOperation) SetLowerBoundTimestamp(v int32)`

SetLowerBoundTimestamp sets LowerBoundTimestamp field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *RangedDatePropertyOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *RangedDatePropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *RangedDatePropertyOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetDefaultValue

`func (o *RangedDatePropertyOperation) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *RangedDatePropertyOperation) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *RangedDatePropertyOperation) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *RangedDatePropertyOperation) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetPropertyType

`func (o *RangedDatePropertyOperation) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *RangedDatePropertyOperation) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *RangedDatePropertyOperation) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.


### GetUpperBoundTimestamp

`func (o *RangedDatePropertyOperation) GetUpperBoundTimestamp() int32`

GetUpperBoundTimestamp returns the UpperBoundTimestamp field if non-nil, zero value otherwise.

### GetUpperBoundTimestampOk

`func (o *RangedDatePropertyOperation) GetUpperBoundTimestampOk() (*int32, bool)`

GetUpperBoundTimestampOk returns a tuple with the UpperBoundTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundTimestamp

`func (o *RangedDatePropertyOperation) SetUpperBoundTimestamp(v int32)`

SetUpperBoundTimestamp sets UpperBoundTimestamp field to given value.


### GetRequiresTimeZoneConversion

`func (o *RangedDatePropertyOperation) GetRequiresTimeZoneConversion() bool`

GetRequiresTimeZoneConversion returns the RequiresTimeZoneConversion field if non-nil, zero value otherwise.

### GetRequiresTimeZoneConversionOk

`func (o *RangedDatePropertyOperation) GetRequiresTimeZoneConversionOk() (*bool, bool)`

GetRequiresTimeZoneConversionOk returns a tuple with the RequiresTimeZoneConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTimeZoneConversion

`func (o *RangedDatePropertyOperation) SetRequiresTimeZoneConversion(v bool)`

SetRequiresTimeZoneConversion sets RequiresTimeZoneConversion field to given value.


### GetOperationType

`func (o *RangedDatePropertyOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *RangedDatePropertyOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *RangedDatePropertyOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperatorName

`func (o *RangedDatePropertyOperation) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *RangedDatePropertyOperation) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *RangedDatePropertyOperation) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetOperator

`func (o *RangedDatePropertyOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RangedDatePropertyOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RangedDatePropertyOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



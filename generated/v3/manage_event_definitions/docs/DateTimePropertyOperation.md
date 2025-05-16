# DateTimePropertyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**PropertyType** | **string** |  | [default to "datetime"]
**RequiresTimeZoneConversion** | **bool** |  | 
**OperationType** | **string** |  | 
**OperatorName** | **string** |  | 
**Operator** | **string** |  | 
**Timestamp** | **int32** |  | 

## Methods

### NewDateTimePropertyOperation

`func NewDateTimePropertyOperation(includeObjectsWithNoValueSet bool, propertyType string, requiresTimeZoneConversion bool, operationType string, operatorName string, operator string, timestamp int32, ) *DateTimePropertyOperation`

NewDateTimePropertyOperation instantiates a new DateTimePropertyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimePropertyOperationWithDefaults

`func NewDateTimePropertyOperationWithDefaults() *DateTimePropertyOperation`

NewDateTimePropertyOperationWithDefaults instantiates a new DateTimePropertyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeObjectsWithNoValueSet

`func (o *DateTimePropertyOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *DateTimePropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *DateTimePropertyOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetDefaultValue

`func (o *DateTimePropertyOperation) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *DateTimePropertyOperation) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *DateTimePropertyOperation) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *DateTimePropertyOperation) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetPropertyType

`func (o *DateTimePropertyOperation) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *DateTimePropertyOperation) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *DateTimePropertyOperation) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.


### GetRequiresTimeZoneConversion

`func (o *DateTimePropertyOperation) GetRequiresTimeZoneConversion() bool`

GetRequiresTimeZoneConversion returns the RequiresTimeZoneConversion field if non-nil, zero value otherwise.

### GetRequiresTimeZoneConversionOk

`func (o *DateTimePropertyOperation) GetRequiresTimeZoneConversionOk() (*bool, bool)`

GetRequiresTimeZoneConversionOk returns a tuple with the RequiresTimeZoneConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTimeZoneConversion

`func (o *DateTimePropertyOperation) SetRequiresTimeZoneConversion(v bool)`

SetRequiresTimeZoneConversion sets RequiresTimeZoneConversion field to given value.


### GetOperationType

`func (o *DateTimePropertyOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *DateTimePropertyOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *DateTimePropertyOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperatorName

`func (o *DateTimePropertyOperation) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *DateTimePropertyOperation) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *DateTimePropertyOperation) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetOperator

`func (o *DateTimePropertyOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *DateTimePropertyOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *DateTimePropertyOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetTimestamp

`func (o *DateTimePropertyOperation) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DateTimePropertyOperation) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DateTimePropertyOperation) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



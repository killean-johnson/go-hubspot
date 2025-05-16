# RollingPropertyUpdatedOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**PropertyType** | **string** |  | [default to "rolling-property-updated"]
**OperationType** | **string** |  | 
**NumberOfDays** | **int32** |  | 
**OperatorName** | **string** |  | 
**Operator** | **string** |  | 

## Methods

### NewRollingPropertyUpdatedOperation

`func NewRollingPropertyUpdatedOperation(includeObjectsWithNoValueSet bool, propertyType string, operationType string, numberOfDays int32, operatorName string, operator string, ) *RollingPropertyUpdatedOperation`

NewRollingPropertyUpdatedOperation instantiates a new RollingPropertyUpdatedOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollingPropertyUpdatedOperationWithDefaults

`func NewRollingPropertyUpdatedOperationWithDefaults() *RollingPropertyUpdatedOperation`

NewRollingPropertyUpdatedOperationWithDefaults instantiates a new RollingPropertyUpdatedOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeObjectsWithNoValueSet

`func (o *RollingPropertyUpdatedOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *RollingPropertyUpdatedOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *RollingPropertyUpdatedOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetDefaultValue

`func (o *RollingPropertyUpdatedOperation) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *RollingPropertyUpdatedOperation) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *RollingPropertyUpdatedOperation) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *RollingPropertyUpdatedOperation) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetPropertyType

`func (o *RollingPropertyUpdatedOperation) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *RollingPropertyUpdatedOperation) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *RollingPropertyUpdatedOperation) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.


### GetOperationType

`func (o *RollingPropertyUpdatedOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *RollingPropertyUpdatedOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *RollingPropertyUpdatedOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetNumberOfDays

`func (o *RollingPropertyUpdatedOperation) GetNumberOfDays() int32`

GetNumberOfDays returns the NumberOfDays field if non-nil, zero value otherwise.

### GetNumberOfDaysOk

`func (o *RollingPropertyUpdatedOperation) GetNumberOfDaysOk() (*int32, bool)`

GetNumberOfDaysOk returns a tuple with the NumberOfDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDays

`func (o *RollingPropertyUpdatedOperation) SetNumberOfDays(v int32)`

SetNumberOfDays sets NumberOfDays field to given value.


### GetOperatorName

`func (o *RollingPropertyUpdatedOperation) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *RollingPropertyUpdatedOperation) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *RollingPropertyUpdatedOperation) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetOperator

`func (o *RollingPropertyUpdatedOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RollingPropertyUpdatedOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RollingPropertyUpdatedOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



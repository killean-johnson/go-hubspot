# BoolPropertyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**PropertyType** | **string** |  | [default to "bool"]
**OperationType** | **string** |  | 
**Value** | **bool** |  | 
**OperatorName** | **string** |  | 
**Operator** | **string** |  | 

## Methods

### NewBoolPropertyOperation

`func NewBoolPropertyOperation(includeObjectsWithNoValueSet bool, propertyType string, operationType string, value bool, operatorName string, operator string, ) *BoolPropertyOperation`

NewBoolPropertyOperation instantiates a new BoolPropertyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoolPropertyOperationWithDefaults

`func NewBoolPropertyOperationWithDefaults() *BoolPropertyOperation`

NewBoolPropertyOperationWithDefaults instantiates a new BoolPropertyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeObjectsWithNoValueSet

`func (o *BoolPropertyOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *BoolPropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *BoolPropertyOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetDefaultValue

`func (o *BoolPropertyOperation) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *BoolPropertyOperation) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *BoolPropertyOperation) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *BoolPropertyOperation) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetPropertyType

`func (o *BoolPropertyOperation) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *BoolPropertyOperation) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *BoolPropertyOperation) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.


### GetOperationType

`func (o *BoolPropertyOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *BoolPropertyOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *BoolPropertyOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetValue

`func (o *BoolPropertyOperation) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BoolPropertyOperation) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BoolPropertyOperation) SetValue(v bool)`

SetValue sets Value field to given value.


### GetOperatorName

`func (o *BoolPropertyOperation) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *BoolPropertyOperation) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *BoolPropertyOperation) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetOperator

`func (o *BoolPropertyOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *BoolPropertyOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *BoolPropertyOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



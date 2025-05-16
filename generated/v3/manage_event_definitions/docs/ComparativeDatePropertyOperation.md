# ComparativeDatePropertyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**PropertyType** | **string** |  | [default to "datetime-comparative"]
**DefaultComparisonValue** | Pointer to **string** |  | [optional] 
**OperationType** | **string** |  | 
**ComparisonPropertyName** | **string** |  | 
**OperatorName** | **string** |  | 
**Operator** | **string** |  | 

## Methods

### NewComparativeDatePropertyOperation

`func NewComparativeDatePropertyOperation(includeObjectsWithNoValueSet bool, propertyType string, operationType string, comparisonPropertyName string, operatorName string, operator string, ) *ComparativeDatePropertyOperation`

NewComparativeDatePropertyOperation instantiates a new ComparativeDatePropertyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComparativeDatePropertyOperationWithDefaults

`func NewComparativeDatePropertyOperationWithDefaults() *ComparativeDatePropertyOperation`

NewComparativeDatePropertyOperationWithDefaults instantiates a new ComparativeDatePropertyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeObjectsWithNoValueSet

`func (o *ComparativeDatePropertyOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *ComparativeDatePropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *ComparativeDatePropertyOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetDefaultValue

`func (o *ComparativeDatePropertyOperation) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ComparativeDatePropertyOperation) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ComparativeDatePropertyOperation) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ComparativeDatePropertyOperation) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetPropertyType

`func (o *ComparativeDatePropertyOperation) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *ComparativeDatePropertyOperation) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *ComparativeDatePropertyOperation) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.


### GetDefaultComparisonValue

`func (o *ComparativeDatePropertyOperation) GetDefaultComparisonValue() string`

GetDefaultComparisonValue returns the DefaultComparisonValue field if non-nil, zero value otherwise.

### GetDefaultComparisonValueOk

`func (o *ComparativeDatePropertyOperation) GetDefaultComparisonValueOk() (*string, bool)`

GetDefaultComparisonValueOk returns a tuple with the DefaultComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultComparisonValue

`func (o *ComparativeDatePropertyOperation) SetDefaultComparisonValue(v string)`

SetDefaultComparisonValue sets DefaultComparisonValue field to given value.

### HasDefaultComparisonValue

`func (o *ComparativeDatePropertyOperation) HasDefaultComparisonValue() bool`

HasDefaultComparisonValue returns a boolean if a field has been set.

### GetOperationType

`func (o *ComparativeDatePropertyOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ComparativeDatePropertyOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ComparativeDatePropertyOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetComparisonPropertyName

`func (o *ComparativeDatePropertyOperation) GetComparisonPropertyName() string`

GetComparisonPropertyName returns the ComparisonPropertyName field if non-nil, zero value otherwise.

### GetComparisonPropertyNameOk

`func (o *ComparativeDatePropertyOperation) GetComparisonPropertyNameOk() (*string, bool)`

GetComparisonPropertyNameOk returns a tuple with the ComparisonPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonPropertyName

`func (o *ComparativeDatePropertyOperation) SetComparisonPropertyName(v string)`

SetComparisonPropertyName sets ComparisonPropertyName field to given value.


### GetOperatorName

`func (o *ComparativeDatePropertyOperation) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *ComparativeDatePropertyOperation) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *ComparativeDatePropertyOperation) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetOperator

`func (o *ComparativeDatePropertyOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ComparativeDatePropertyOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ComparativeDatePropertyOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



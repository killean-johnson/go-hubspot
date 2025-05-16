# MultiStringPropertyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoalescingRefineBy** | [**AllPropertyTypesOperationCoalescingRefineBy**](AllPropertyTypesOperationCoalescingRefineBy.md) |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**PruningRefineBy** | Pointer to [**AllPropertyTypesOperationPruningRefineBy**](AllPropertyTypesOperationPruningRefineBy.md) |  | [optional] 
**PropertyType** | **string** |  | [default to "multistring"]
**Values** | **[]string** |  | 
**OperationType** | **string** |  | 
**OperatorName** | **string** |  | 
**Operator** | **string** |  | 

## Methods

### NewMultiStringPropertyOperation

`func NewMultiStringPropertyOperation(coalescingRefineBy AllPropertyTypesOperationCoalescingRefineBy, includeObjectsWithNoValueSet bool, propertyType string, values []string, operationType string, operatorName string, operator string, ) *MultiStringPropertyOperation`

NewMultiStringPropertyOperation instantiates a new MultiStringPropertyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiStringPropertyOperationWithDefaults

`func NewMultiStringPropertyOperationWithDefaults() *MultiStringPropertyOperation`

NewMultiStringPropertyOperationWithDefaults instantiates a new MultiStringPropertyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoalescingRefineBy

`func (o *MultiStringPropertyOperation) GetCoalescingRefineBy() AllPropertyTypesOperationCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *MultiStringPropertyOperation) GetCoalescingRefineByOk() (*AllPropertyTypesOperationCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *MultiStringPropertyOperation) SetCoalescingRefineBy(v AllPropertyTypesOperationCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *MultiStringPropertyOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *MultiStringPropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *MultiStringPropertyOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetDefaultValue

`func (o *MultiStringPropertyOperation) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *MultiStringPropertyOperation) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *MultiStringPropertyOperation) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *MultiStringPropertyOperation) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetPruningRefineBy

`func (o *MultiStringPropertyOperation) GetPruningRefineBy() AllPropertyTypesOperationPruningRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *MultiStringPropertyOperation) GetPruningRefineByOk() (*AllPropertyTypesOperationPruningRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *MultiStringPropertyOperation) SetPruningRefineBy(v AllPropertyTypesOperationPruningRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.

### HasPruningRefineBy

`func (o *MultiStringPropertyOperation) HasPruningRefineBy() bool`

HasPruningRefineBy returns a boolean if a field has been set.

### GetPropertyType

`func (o *MultiStringPropertyOperation) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *MultiStringPropertyOperation) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *MultiStringPropertyOperation) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.


### GetValues

`func (o *MultiStringPropertyOperation) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MultiStringPropertyOperation) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MultiStringPropertyOperation) SetValues(v []string)`

SetValues sets Values field to given value.


### GetOperationType

`func (o *MultiStringPropertyOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *MultiStringPropertyOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *MultiStringPropertyOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperatorName

`func (o *MultiStringPropertyOperation) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *MultiStringPropertyOperation) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *MultiStringPropertyOperation) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetOperator

`func (o *MultiStringPropertyOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MultiStringPropertyOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MultiStringPropertyOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



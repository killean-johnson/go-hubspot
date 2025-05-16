# AllPropertyTypesOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoalescingRefineBy** | [**AllPropertyTypesOperationCoalescingRefineBy**](AllPropertyTypesOperationCoalescingRefineBy.md) |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**DefaultValue** | Pointer to **string** |  | [optional] 
**PruningRefineBy** | Pointer to [**AllPropertyTypesOperationPruningRefineBy**](AllPropertyTypesOperationPruningRefineBy.md) |  | [optional] 
**PropertyType** | **string** |  | [default to "alltypes"]
**OperationType** | **string** |  | 
**OperatorName** | **string** |  | 
**Operator** | **string** |  | 

## Methods

### NewAllPropertyTypesOperation

`func NewAllPropertyTypesOperation(coalescingRefineBy AllPropertyTypesOperationCoalescingRefineBy, includeObjectsWithNoValueSet bool, propertyType string, operationType string, operatorName string, operator string, ) *AllPropertyTypesOperation`

NewAllPropertyTypesOperation instantiates a new AllPropertyTypesOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllPropertyTypesOperationWithDefaults

`func NewAllPropertyTypesOperationWithDefaults() *AllPropertyTypesOperation`

NewAllPropertyTypesOperationWithDefaults instantiates a new AllPropertyTypesOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoalescingRefineBy

`func (o *AllPropertyTypesOperation) GetCoalescingRefineBy() AllPropertyTypesOperationCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *AllPropertyTypesOperation) GetCoalescingRefineByOk() (*AllPropertyTypesOperationCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *AllPropertyTypesOperation) SetCoalescingRefineBy(v AllPropertyTypesOperationCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *AllPropertyTypesOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *AllPropertyTypesOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *AllPropertyTypesOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetDefaultValue

`func (o *AllPropertyTypesOperation) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *AllPropertyTypesOperation) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *AllPropertyTypesOperation) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *AllPropertyTypesOperation) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetPruningRefineBy

`func (o *AllPropertyTypesOperation) GetPruningRefineBy() AllPropertyTypesOperationPruningRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *AllPropertyTypesOperation) GetPruningRefineByOk() (*AllPropertyTypesOperationPruningRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *AllPropertyTypesOperation) SetPruningRefineBy(v AllPropertyTypesOperationPruningRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.

### HasPruningRefineBy

`func (o *AllPropertyTypesOperation) HasPruningRefineBy() bool`

HasPruningRefineBy returns a boolean if a field has been set.

### GetPropertyType

`func (o *AllPropertyTypesOperation) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *AllPropertyTypesOperation) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *AllPropertyTypesOperation) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.


### GetOperationType

`func (o *AllPropertyTypesOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *AllPropertyTypesOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *AllPropertyTypesOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperatorName

`func (o *AllPropertyTypesOperation) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *AllPropertyTypesOperation) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *AllPropertyTypesOperation) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetOperator

`func (o *AllPropertyTypesOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AllPropertyTypesOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AllPropertyTypesOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RollupExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationTypes** | [**[]AssociationSpec**](AssociationSpec.md) |  | 
**EmptyRollupValue** | Pointer to **string** |  | [optional] 
**ConditionalFormula** | Pointer to **string** |  | [optional] 
**SourceObjectTypeId** | **string** |  | 
**ConditionalExpression** | Pointer to [**OrInputsInner**](OrInputsInner.md) |  | [optional] 
**RollupOperator** | **string** |  | 
**SourcePropertyName** | **string** |  | 
**SourceCompareByPropertyName** | Pointer to **string** |  | [optional] 

## Methods

### NewRollupExpression

`func NewRollupExpression(associationTypes []AssociationSpec, sourceObjectTypeId string, rollupOperator string, sourcePropertyName string, ) *RollupExpression`

NewRollupExpression instantiates a new RollupExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollupExpressionWithDefaults

`func NewRollupExpressionWithDefaults() *RollupExpression`

NewRollupExpressionWithDefaults instantiates a new RollupExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationTypes

`func (o *RollupExpression) GetAssociationTypes() []AssociationSpec`

GetAssociationTypes returns the AssociationTypes field if non-nil, zero value otherwise.

### GetAssociationTypesOk

`func (o *RollupExpression) GetAssociationTypesOk() (*[]AssociationSpec, bool)`

GetAssociationTypesOk returns a tuple with the AssociationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypes

`func (o *RollupExpression) SetAssociationTypes(v []AssociationSpec)`

SetAssociationTypes sets AssociationTypes field to given value.


### GetEmptyRollupValue

`func (o *RollupExpression) GetEmptyRollupValue() string`

GetEmptyRollupValue returns the EmptyRollupValue field if non-nil, zero value otherwise.

### GetEmptyRollupValueOk

`func (o *RollupExpression) GetEmptyRollupValueOk() (*string, bool)`

GetEmptyRollupValueOk returns a tuple with the EmptyRollupValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyRollupValue

`func (o *RollupExpression) SetEmptyRollupValue(v string)`

SetEmptyRollupValue sets EmptyRollupValue field to given value.

### HasEmptyRollupValue

`func (o *RollupExpression) HasEmptyRollupValue() bool`

HasEmptyRollupValue returns a boolean if a field has been set.

### GetConditionalFormula

`func (o *RollupExpression) GetConditionalFormula() string`

GetConditionalFormula returns the ConditionalFormula field if non-nil, zero value otherwise.

### GetConditionalFormulaOk

`func (o *RollupExpression) GetConditionalFormulaOk() (*string, bool)`

GetConditionalFormulaOk returns a tuple with the ConditionalFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalFormula

`func (o *RollupExpression) SetConditionalFormula(v string)`

SetConditionalFormula sets ConditionalFormula field to given value.

### HasConditionalFormula

`func (o *RollupExpression) HasConditionalFormula() bool`

HasConditionalFormula returns a boolean if a field has been set.

### GetSourceObjectTypeId

`func (o *RollupExpression) GetSourceObjectTypeId() string`

GetSourceObjectTypeId returns the SourceObjectTypeId field if non-nil, zero value otherwise.

### GetSourceObjectTypeIdOk

`func (o *RollupExpression) GetSourceObjectTypeIdOk() (*string, bool)`

GetSourceObjectTypeIdOk returns a tuple with the SourceObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceObjectTypeId

`func (o *RollupExpression) SetSourceObjectTypeId(v string)`

SetSourceObjectTypeId sets SourceObjectTypeId field to given value.


### GetConditionalExpression

`func (o *RollupExpression) GetConditionalExpression() OrInputsInner`

GetConditionalExpression returns the ConditionalExpression field if non-nil, zero value otherwise.

### GetConditionalExpressionOk

`func (o *RollupExpression) GetConditionalExpressionOk() (*OrInputsInner, bool)`

GetConditionalExpressionOk returns a tuple with the ConditionalExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalExpression

`func (o *RollupExpression) SetConditionalExpression(v OrInputsInner)`

SetConditionalExpression sets ConditionalExpression field to given value.

### HasConditionalExpression

`func (o *RollupExpression) HasConditionalExpression() bool`

HasConditionalExpression returns a boolean if a field has been set.

### GetRollupOperator

`func (o *RollupExpression) GetRollupOperator() string`

GetRollupOperator returns the RollupOperator field if non-nil, zero value otherwise.

### GetRollupOperatorOk

`func (o *RollupExpression) GetRollupOperatorOk() (*string, bool)`

GetRollupOperatorOk returns a tuple with the RollupOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupOperator

`func (o *RollupExpression) SetRollupOperator(v string)`

SetRollupOperator sets RollupOperator field to given value.


### GetSourcePropertyName

`func (o *RollupExpression) GetSourcePropertyName() string`

GetSourcePropertyName returns the SourcePropertyName field if non-nil, zero value otherwise.

### GetSourcePropertyNameOk

`func (o *RollupExpression) GetSourcePropertyNameOk() (*string, bool)`

GetSourcePropertyNameOk returns a tuple with the SourcePropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePropertyName

`func (o *RollupExpression) SetSourcePropertyName(v string)`

SetSourcePropertyName sets SourcePropertyName field to given value.


### GetSourceCompareByPropertyName

`func (o *RollupExpression) GetSourceCompareByPropertyName() string`

GetSourceCompareByPropertyName returns the SourceCompareByPropertyName field if non-nil, zero value otherwise.

### GetSourceCompareByPropertyNameOk

`func (o *RollupExpression) GetSourceCompareByPropertyNameOk() (*string, bool)`

GetSourceCompareByPropertyNameOk returns a tuple with the SourceCompareByPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCompareByPropertyName

`func (o *RollupExpression) SetSourceCompareByPropertyName(v string)`

SetSourceCompareByPropertyName sets SourceCompareByPropertyName field to given value.

### HasSourceCompareByPropertyName

`func (o *RollupExpression) HasSourceCompareByPropertyName() bool`

HasSourceCompareByPropertyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



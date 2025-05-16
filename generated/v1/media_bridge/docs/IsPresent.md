# IsPresent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**ExpressionToEvaluate** | [**OrInputsInner**](OrInputsInner.md) |  | 
**Value** | Pointer to **bool** |  | [optional] 
**Operator** | **string** |  | [default to "IS_PRESENT"]

## Methods

### NewIsPresent

`func NewIsPresent(expressionToEvaluate OrInputsInner, operator string, ) *IsPresent`

NewIsPresent instantiates a new IsPresent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsPresentWithDefaults

`func NewIsPresentWithDefaults() *IsPresent`

NewIsPresentWithDefaults instantiates a new IsPresent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *IsPresent) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *IsPresent) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *IsPresent) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *IsPresent) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *IsPresent) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *IsPresent) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *IsPresent) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *IsPresent) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetExpressionToEvaluate

`func (o *IsPresent) GetExpressionToEvaluate() OrInputsInner`

GetExpressionToEvaluate returns the ExpressionToEvaluate field if non-nil, zero value otherwise.

### GetExpressionToEvaluateOk

`func (o *IsPresent) GetExpressionToEvaluateOk() (*OrInputsInner, bool)`

GetExpressionToEvaluateOk returns a tuple with the ExpressionToEvaluate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionToEvaluate

`func (o *IsPresent) SetExpressionToEvaluate(v OrInputsInner)`

SetExpressionToEvaluate sets ExpressionToEvaluate field to given value.


### GetValue

`func (o *IsPresent) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IsPresent) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IsPresent) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *IsPresent) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *IsPresent) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *IsPresent) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *IsPresent) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



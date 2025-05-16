# OrInputsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Operator** | **string** |  | [default to "FORMAT_FULL_NAME"]
**EnclosedInParentheses** | **bool** |  | 
**StringToCheck** | [**OrInputsInner**](OrInputsInner.md) |  | 
**IfExpression** | [**OrInputsInner**](OrInputsInner.md) |  | 
**ElseExpression** | Pointer to [**OrInputsInner**](OrInputsInner.md) |  | [optional] 
**ExpressionToEvaluate** | [**OrInputsInner**](OrInputsInner.md) |  | 

## Methods

### NewOrInputsInner

`func NewOrInputsInner(operator string, enclosedInParentheses bool, stringToCheck OrInputsInner, ifExpression OrInputsInner, expressionToEvaluate OrInputsInner, ) *OrInputsInner`

NewOrInputsInner instantiates a new OrInputsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrInputsInnerWithDefaults

`func NewOrInputsInnerWithDefaults() *OrInputsInner`

NewOrInputsInnerWithDefaults instantiates a new OrInputsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *OrInputsInner) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *OrInputsInner) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *OrInputsInner) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *OrInputsInner) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *OrInputsInner) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *OrInputsInner) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *OrInputsInner) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *OrInputsInner) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetValue

`func (o *OrInputsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OrInputsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OrInputsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OrInputsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *OrInputsInner) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *OrInputsInner) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *OrInputsInner) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetEnclosedInParentheses

`func (o *OrInputsInner) GetEnclosedInParentheses() bool`

GetEnclosedInParentheses returns the EnclosedInParentheses field if non-nil, zero value otherwise.

### GetEnclosedInParenthesesOk

`func (o *OrInputsInner) GetEnclosedInParenthesesOk() (*bool, bool)`

GetEnclosedInParenthesesOk returns a tuple with the EnclosedInParentheses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosedInParentheses

`func (o *OrInputsInner) SetEnclosedInParentheses(v bool)`

SetEnclosedInParentheses sets EnclosedInParentheses field to given value.


### GetStringToCheck

`func (o *OrInputsInner) GetStringToCheck() OrInputsInner`

GetStringToCheck returns the StringToCheck field if non-nil, zero value otherwise.

### GetStringToCheckOk

`func (o *OrInputsInner) GetStringToCheckOk() (*OrInputsInner, bool)`

GetStringToCheckOk returns a tuple with the StringToCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringToCheck

`func (o *OrInputsInner) SetStringToCheck(v OrInputsInner)`

SetStringToCheck sets StringToCheck field to given value.


### GetIfExpression

`func (o *OrInputsInner) GetIfExpression() OrInputsInner`

GetIfExpression returns the IfExpression field if non-nil, zero value otherwise.

### GetIfExpressionOk

`func (o *OrInputsInner) GetIfExpressionOk() (*OrInputsInner, bool)`

GetIfExpressionOk returns a tuple with the IfExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfExpression

`func (o *OrInputsInner) SetIfExpression(v OrInputsInner)`

SetIfExpression sets IfExpression field to given value.


### GetElseExpression

`func (o *OrInputsInner) GetElseExpression() OrInputsInner`

GetElseExpression returns the ElseExpression field if non-nil, zero value otherwise.

### GetElseExpressionOk

`func (o *OrInputsInner) GetElseExpressionOk() (*OrInputsInner, bool)`

GetElseExpressionOk returns a tuple with the ElseExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseExpression

`func (o *OrInputsInner) SetElseExpression(v OrInputsInner)`

SetElseExpression sets ElseExpression field to given value.

### HasElseExpression

`func (o *OrInputsInner) HasElseExpression() bool`

HasElseExpression returns a boolean if a field has been set.

### GetExpressionToEvaluate

`func (o *OrInputsInner) GetExpressionToEvaluate() OrInputsInner`

GetExpressionToEvaluate returns the ExpressionToEvaluate field if non-nil, zero value otherwise.

### GetExpressionToEvaluateOk

`func (o *OrInputsInner) GetExpressionToEvaluateOk() (*OrInputsInner, bool)`

GetExpressionToEvaluateOk returns a tuple with the ExpressionToEvaluate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionToEvaluate

`func (o *OrInputsInner) SetExpressionToEvaluate(v OrInputsInner)`

SetExpressionToEvaluate sets ExpressionToEvaluate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



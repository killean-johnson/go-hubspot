# IfString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**IfExpression** | [**OrInputsInner**](OrInputsInner.md) |  | 
**Value** | Pointer to **string** |  | [optional] 
**ElseExpression** | Pointer to [**OrInputsInner**](OrInputsInner.md) |  | [optional] 
**Operator** | **string** |  | [default to "IF_STRING"]

## Methods

### NewIfString

`func NewIfString(ifExpression OrInputsInner, operator string, ) *IfString`

NewIfString instantiates a new IfString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIfStringWithDefaults

`func NewIfStringWithDefaults() *IfString`

NewIfStringWithDefaults instantiates a new IfString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *IfString) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *IfString) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *IfString) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *IfString) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *IfString) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *IfString) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *IfString) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *IfString) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetIfExpression

`func (o *IfString) GetIfExpression() OrInputsInner`

GetIfExpression returns the IfExpression field if non-nil, zero value otherwise.

### GetIfExpressionOk

`func (o *IfString) GetIfExpressionOk() (*OrInputsInner, bool)`

GetIfExpressionOk returns a tuple with the IfExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfExpression

`func (o *IfString) SetIfExpression(v OrInputsInner)`

SetIfExpression sets IfExpression field to given value.


### GetValue

`func (o *IfString) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IfString) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IfString) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IfString) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetElseExpression

`func (o *IfString) GetElseExpression() OrInputsInner`

GetElseExpression returns the ElseExpression field if non-nil, zero value otherwise.

### GetElseExpressionOk

`func (o *IfString) GetElseExpressionOk() (*OrInputsInner, bool)`

GetElseExpressionOk returns a tuple with the ElseExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseExpression

`func (o *IfString) SetElseExpression(v OrInputsInner)`

SetElseExpression sets ElseExpression field to given value.

### HasElseExpression

`func (o *IfString) HasElseExpression() bool`

HasElseExpression returns a boolean if a field has been set.

### GetOperator

`func (o *IfString) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *IfString) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *IfString) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



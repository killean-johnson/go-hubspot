# IfNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**IfExpression** | [**OrInputsInner**](OrInputsInner.md) |  | 
**Value** | Pointer to **float32** |  | [optional] 
**ElseExpression** | Pointer to [**OrInputsInner**](OrInputsInner.md) |  | [optional] 
**Operator** | **string** |  | [default to "IF_NUMBER"]

## Methods

### NewIfNumber

`func NewIfNumber(ifExpression OrInputsInner, operator string, ) *IfNumber`

NewIfNumber instantiates a new IfNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIfNumberWithDefaults

`func NewIfNumberWithDefaults() *IfNumber`

NewIfNumberWithDefaults instantiates a new IfNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *IfNumber) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *IfNumber) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *IfNumber) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *IfNumber) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *IfNumber) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *IfNumber) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *IfNumber) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *IfNumber) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetIfExpression

`func (o *IfNumber) GetIfExpression() OrInputsInner`

GetIfExpression returns the IfExpression field if non-nil, zero value otherwise.

### GetIfExpressionOk

`func (o *IfNumber) GetIfExpressionOk() (*OrInputsInner, bool)`

GetIfExpressionOk returns a tuple with the IfExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfExpression

`func (o *IfNumber) SetIfExpression(v OrInputsInner)`

SetIfExpression sets IfExpression field to given value.


### GetValue

`func (o *IfNumber) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IfNumber) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IfNumber) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *IfNumber) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetElseExpression

`func (o *IfNumber) GetElseExpression() OrInputsInner`

GetElseExpression returns the ElseExpression field if non-nil, zero value otherwise.

### GetElseExpressionOk

`func (o *IfNumber) GetElseExpressionOk() (*OrInputsInner, bool)`

GetElseExpressionOk returns a tuple with the ElseExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElseExpression

`func (o *IfNumber) SetElseExpression(v OrInputsInner)`

SetElseExpression sets ElseExpression field to given value.

### HasElseExpression

`func (o *IfNumber) HasElseExpression() bool`

HasElseExpression returns a boolean if a field has been set.

### GetOperator

`func (o *IfNumber) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *IfNumber) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *IfNumber) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



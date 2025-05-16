# MultiplyNumbers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**EnclosedInParentheses** | **bool** |  | 
**Value** | Pointer to **float32** |  | [optional] 
**Operator** | **string** |  | [default to "MULTIPLY_NUMBERS"]

## Methods

### NewMultiplyNumbers

`func NewMultiplyNumbers(enclosedInParentheses bool, operator string, ) *MultiplyNumbers`

NewMultiplyNumbers instantiates a new MultiplyNumbers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiplyNumbersWithDefaults

`func NewMultiplyNumbersWithDefaults() *MultiplyNumbers`

NewMultiplyNumbersWithDefaults instantiates a new MultiplyNumbers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *MultiplyNumbers) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *MultiplyNumbers) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *MultiplyNumbers) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *MultiplyNumbers) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *MultiplyNumbers) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *MultiplyNumbers) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *MultiplyNumbers) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *MultiplyNumbers) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetEnclosedInParentheses

`func (o *MultiplyNumbers) GetEnclosedInParentheses() bool`

GetEnclosedInParentheses returns the EnclosedInParentheses field if non-nil, zero value otherwise.

### GetEnclosedInParenthesesOk

`func (o *MultiplyNumbers) GetEnclosedInParenthesesOk() (*bool, bool)`

GetEnclosedInParenthesesOk returns a tuple with the EnclosedInParentheses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosedInParentheses

`func (o *MultiplyNumbers) SetEnclosedInParentheses(v bool)`

SetEnclosedInParentheses sets EnclosedInParentheses field to given value.


### GetValue

`func (o *MultiplyNumbers) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MultiplyNumbers) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MultiplyNumbers) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *MultiplyNumbers) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *MultiplyNumbers) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MultiplyNumbers) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MultiplyNumbers) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



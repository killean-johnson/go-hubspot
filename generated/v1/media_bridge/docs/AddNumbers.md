# AddNumbers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**EnclosedInParentheses** | **bool** |  | 
**Value** | Pointer to **float32** |  | [optional] 
**Operator** | **string** |  | [default to "ADD_NUMBERS"]

## Methods

### NewAddNumbers

`func NewAddNumbers(enclosedInParentheses bool, operator string, ) *AddNumbers`

NewAddNumbers instantiates a new AddNumbers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddNumbersWithDefaults

`func NewAddNumbersWithDefaults() *AddNumbers`

NewAddNumbersWithDefaults instantiates a new AddNumbers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *AddNumbers) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *AddNumbers) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *AddNumbers) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *AddNumbers) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *AddNumbers) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *AddNumbers) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *AddNumbers) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *AddNumbers) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetEnclosedInParentheses

`func (o *AddNumbers) GetEnclosedInParentheses() bool`

GetEnclosedInParentheses returns the EnclosedInParentheses field if non-nil, zero value otherwise.

### GetEnclosedInParenthesesOk

`func (o *AddNumbers) GetEnclosedInParenthesesOk() (*bool, bool)`

GetEnclosedInParenthesesOk returns a tuple with the EnclosedInParentheses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosedInParentheses

`func (o *AddNumbers) SetEnclosedInParentheses(v bool)`

SetEnclosedInParentheses sets EnclosedInParentheses field to given value.


### GetValue

`func (o *AddNumbers) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddNumbers) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddNumbers) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *AddNumbers) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *AddNumbers) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AddNumbers) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AddNumbers) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



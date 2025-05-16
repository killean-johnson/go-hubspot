# Or

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**EnclosedInParentheses** | **bool** |  | 
**Value** | Pointer to **bool** |  | [optional] 
**Operator** | **string** |  | [default to "OR"]

## Methods

### NewOr

`func NewOr(enclosedInParentheses bool, operator string, ) *Or`

NewOr instantiates a new Or object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrWithDefaults

`func NewOrWithDefaults() *Or`

NewOrWithDefaults instantiates a new Or object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *Or) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *Or) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *Or) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *Or) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *Or) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *Or) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *Or) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *Or) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetEnclosedInParentheses

`func (o *Or) GetEnclosedInParentheses() bool`

GetEnclosedInParentheses returns the EnclosedInParentheses field if non-nil, zero value otherwise.

### GetEnclosedInParenthesesOk

`func (o *Or) GetEnclosedInParenthesesOk() (*bool, bool)`

GetEnclosedInParenthesesOk returns a tuple with the EnclosedInParentheses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosedInParentheses

`func (o *Or) SetEnclosedInParentheses(v bool)`

SetEnclosedInParentheses sets EnclosedInParentheses field to given value.


### GetValue

`func (o *Or) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Or) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Or) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *Or) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *Or) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Or) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Or) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BooleanPropertyVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**Value** | Pointer to **bool** |  | [optional] 
**Operator** | **string** |  | [default to "BOOLEAN_PROPERTY_VARIABLE"]

## Methods

### NewBooleanPropertyVariable

`func NewBooleanPropertyVariable(operator string, ) *BooleanPropertyVariable`

NewBooleanPropertyVariable instantiates a new BooleanPropertyVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooleanPropertyVariableWithDefaults

`func NewBooleanPropertyVariableWithDefaults() *BooleanPropertyVariable`

NewBooleanPropertyVariableWithDefaults instantiates a new BooleanPropertyVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *BooleanPropertyVariable) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *BooleanPropertyVariable) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *BooleanPropertyVariable) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *BooleanPropertyVariable) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *BooleanPropertyVariable) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *BooleanPropertyVariable) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *BooleanPropertyVariable) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *BooleanPropertyVariable) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetValue

`func (o *BooleanPropertyVariable) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BooleanPropertyVariable) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BooleanPropertyVariable) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *BooleanPropertyVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *BooleanPropertyVariable) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *BooleanPropertyVariable) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *BooleanPropertyVariable) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



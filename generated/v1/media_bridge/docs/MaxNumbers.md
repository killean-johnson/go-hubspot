# MaxNumbers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 
**Operator** | **string** |  | [default to "MAX_NUMBERS"]

## Methods

### NewMaxNumbers

`func NewMaxNumbers(operator string, ) *MaxNumbers`

NewMaxNumbers instantiates a new MaxNumbers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxNumbersWithDefaults

`func NewMaxNumbersWithDefaults() *MaxNumbers`

NewMaxNumbersWithDefaults instantiates a new MaxNumbers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *MaxNumbers) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *MaxNumbers) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *MaxNumbers) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *MaxNumbers) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *MaxNumbers) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *MaxNumbers) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *MaxNumbers) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *MaxNumbers) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetValue

`func (o *MaxNumbers) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MaxNumbers) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MaxNumbers) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *MaxNumbers) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *MaxNumbers) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MaxNumbers) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MaxNumbers) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



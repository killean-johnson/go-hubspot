# RoundNearestNumbers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 
**Operator** | **string** |  | [default to "ROUND_NEAREST"]

## Methods

### NewRoundNearestNumbers

`func NewRoundNearestNumbers(operator string, ) *RoundNearestNumbers`

NewRoundNearestNumbers instantiates a new RoundNearestNumbers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoundNearestNumbersWithDefaults

`func NewRoundNearestNumbersWithDefaults() *RoundNearestNumbers`

NewRoundNearestNumbersWithDefaults instantiates a new RoundNearestNumbers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *RoundNearestNumbers) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *RoundNearestNumbers) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *RoundNearestNumbers) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *RoundNearestNumbers) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *RoundNearestNumbers) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *RoundNearestNumbers) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *RoundNearestNumbers) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *RoundNearestNumbers) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetValue

`func (o *RoundNearestNumbers) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RoundNearestNumbers) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RoundNearestNumbers) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *RoundNearestNumbers) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *RoundNearestNumbers) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RoundNearestNumbers) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RoundNearestNumbers) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



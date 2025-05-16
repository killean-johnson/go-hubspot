# NumberToString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Operator** | **string** |  | [default to "NUMBER_TO_STRING"]

## Methods

### NewNumberToString

`func NewNumberToString(operator string, ) *NumberToString`

NewNumberToString instantiates a new NumberToString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberToStringWithDefaults

`func NewNumberToStringWithDefaults() *NumberToString`

NewNumberToStringWithDefaults instantiates a new NumberToString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *NumberToString) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *NumberToString) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *NumberToString) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *NumberToString) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *NumberToString) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *NumberToString) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *NumberToString) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *NumberToString) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetValue

`func (o *NumberToString) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NumberToString) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NumberToString) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NumberToString) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *NumberToString) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *NumberToString) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *NumberToString) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



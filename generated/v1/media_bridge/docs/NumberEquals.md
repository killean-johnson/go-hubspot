# NumberEquals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**Value** | Pointer to **bool** |  | [optional] 
**Operator** | **string** |  | [default to "NUMBER_EQUALS"]

## Methods

### NewNumberEquals

`func NewNumberEquals(operator string, ) *NumberEquals`

NewNumberEquals instantiates a new NumberEquals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberEqualsWithDefaults

`func NewNumberEqualsWithDefaults() *NumberEquals`

NewNumberEqualsWithDefaults instantiates a new NumberEquals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *NumberEquals) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *NumberEquals) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *NumberEquals) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *NumberEquals) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *NumberEquals) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *NumberEquals) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *NumberEquals) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *NumberEquals) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetValue

`func (o *NumberEquals) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NumberEquals) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NumberEquals) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *NumberEquals) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *NumberEquals) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *NumberEquals) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *NumberEquals) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



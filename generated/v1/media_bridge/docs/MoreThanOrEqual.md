# MoreThanOrEqual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**Value** | Pointer to **bool** |  | [optional] 
**Operator** | **string** |  | [default to "MORE_THAN_OR_EQUAL"]

## Methods

### NewMoreThanOrEqual

`func NewMoreThanOrEqual(operator string, ) *MoreThanOrEqual`

NewMoreThanOrEqual instantiates a new MoreThanOrEqual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoreThanOrEqualWithDefaults

`func NewMoreThanOrEqualWithDefaults() *MoreThanOrEqual`

NewMoreThanOrEqualWithDefaults instantiates a new MoreThanOrEqual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *MoreThanOrEqual) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *MoreThanOrEqual) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *MoreThanOrEqual) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *MoreThanOrEqual) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *MoreThanOrEqual) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *MoreThanOrEqual) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *MoreThanOrEqual) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *MoreThanOrEqual) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetValue

`func (o *MoreThanOrEqual) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MoreThanOrEqual) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MoreThanOrEqual) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *MoreThanOrEqual) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *MoreThanOrEqual) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MoreThanOrEqual) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MoreThanOrEqual) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



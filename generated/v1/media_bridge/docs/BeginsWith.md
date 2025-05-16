# BeginsWith

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StringToCheck** | [**OrInputsInner**](OrInputsInner.md) |  | 
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**Value** | Pointer to **bool** |  | [optional] 
**Operator** | **string** |  | [default to "BEGINS_WITH"]

## Methods

### NewBeginsWith

`func NewBeginsWith(stringToCheck OrInputsInner, operator string, ) *BeginsWith`

NewBeginsWith instantiates a new BeginsWith object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeginsWithWithDefaults

`func NewBeginsWithWithDefaults() *BeginsWith`

NewBeginsWithWithDefaults instantiates a new BeginsWith object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStringToCheck

`func (o *BeginsWith) GetStringToCheck() OrInputsInner`

GetStringToCheck returns the StringToCheck field if non-nil, zero value otherwise.

### GetStringToCheckOk

`func (o *BeginsWith) GetStringToCheckOk() (*OrInputsInner, bool)`

GetStringToCheckOk returns a tuple with the StringToCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringToCheck

`func (o *BeginsWith) SetStringToCheck(v OrInputsInner)`

SetStringToCheck sets StringToCheck field to given value.


### GetPropertyName

`func (o *BeginsWith) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *BeginsWith) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *BeginsWith) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *BeginsWith) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *BeginsWith) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *BeginsWith) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *BeginsWith) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *BeginsWith) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetValue

`func (o *BeginsWith) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BeginsWith) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BeginsWith) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *BeginsWith) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *BeginsWith) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *BeginsWith) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *BeginsWith) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



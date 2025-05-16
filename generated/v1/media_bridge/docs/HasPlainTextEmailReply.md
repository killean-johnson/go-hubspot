# HasPlainTextEmailReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**Value** | Pointer to **bool** |  | [optional] 
**Operator** | **string** |  | [default to "HAS_PLAIN_TEXT_EMAIL_REPLY"]

## Methods

### NewHasPlainTextEmailReply

`func NewHasPlainTextEmailReply(operator string, ) *HasPlainTextEmailReply`

NewHasPlainTextEmailReply instantiates a new HasPlainTextEmailReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHasPlainTextEmailReplyWithDefaults

`func NewHasPlainTextEmailReplyWithDefaults() *HasPlainTextEmailReply`

NewHasPlainTextEmailReplyWithDefaults instantiates a new HasPlainTextEmailReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *HasPlainTextEmailReply) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *HasPlainTextEmailReply) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *HasPlainTextEmailReply) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *HasPlainTextEmailReply) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *HasPlainTextEmailReply) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *HasPlainTextEmailReply) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *HasPlainTextEmailReply) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *HasPlainTextEmailReply) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetValue

`func (o *HasPlainTextEmailReply) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HasPlainTextEmailReply) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HasPlainTextEmailReply) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *HasPlainTextEmailReply) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *HasPlainTextEmailReply) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *HasPlainTextEmailReply) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *HasPlainTextEmailReply) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ExtractMostRecentEmailReplyText

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to [**[]OrInputsInner**](OrInputsInner.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Operator** | **string** |  | [default to "EXTRACT_MOST_RECENT_EMAIL_REPLY_TEXT"]

## Methods

### NewExtractMostRecentEmailReplyText

`func NewExtractMostRecentEmailReplyText(operator string, ) *ExtractMostRecentEmailReplyText`

NewExtractMostRecentEmailReplyText instantiates a new ExtractMostRecentEmailReplyText object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractMostRecentEmailReplyTextWithDefaults

`func NewExtractMostRecentEmailReplyTextWithDefaults() *ExtractMostRecentEmailReplyText`

NewExtractMostRecentEmailReplyTextWithDefaults instantiates a new ExtractMostRecentEmailReplyText object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *ExtractMostRecentEmailReplyText) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *ExtractMostRecentEmailReplyText) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *ExtractMostRecentEmailReplyText) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *ExtractMostRecentEmailReplyText) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetInputs

`func (o *ExtractMostRecentEmailReplyText) GetInputs() []OrInputsInner`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ExtractMostRecentEmailReplyText) GetInputsOk() (*[]OrInputsInner, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ExtractMostRecentEmailReplyText) SetInputs(v []OrInputsInner)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *ExtractMostRecentEmailReplyText) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetValue

`func (o *ExtractMostRecentEmailReplyText) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ExtractMostRecentEmailReplyText) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ExtractMostRecentEmailReplyText) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ExtractMostRecentEmailReplyText) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOperator

`func (o *ExtractMostRecentEmailReplyText) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ExtractMostRecentEmailReplyText) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ExtractMostRecentEmailReplyText) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



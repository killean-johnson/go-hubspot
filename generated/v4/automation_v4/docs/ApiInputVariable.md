# ApiInputVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | [**ApiInputVariableValue**](ApiInputVariableValue.md) |  | 

## Methods

### NewApiInputVariable

`func NewApiInputVariable(name string, value ApiInputVariableValue, ) *ApiInputVariable`

NewApiInputVariable instantiates a new ApiInputVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInputVariableWithDefaults

`func NewApiInputVariableWithDefaults() *ApiInputVariable`

NewApiInputVariableWithDefaults instantiates a new ApiInputVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiInputVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiInputVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiInputVariable) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ApiInputVariable) GetValue() ApiInputVariableValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiInputVariable) GetValueOk() (*ApiInputVariableValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiInputVariable) SetValue(v ApiInputVariableValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



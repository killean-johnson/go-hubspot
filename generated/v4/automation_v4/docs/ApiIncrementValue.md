# ApiIncrementValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncrementAmount** | **float32** | The amount be which to increment | 
**Type** | **string** | This is the type of input value. This can be one of: \&quot;FIELD_DATA\&quot;, \&quot;OBJECT_PROPERTY\&quot;, \&quot;STATIC_VALUE\&quot;, \&quot;RELATIVE_DATETIME\&quot;, \&quot;TIMESTAMP\&quot;, \&quot;INCREMENT\&quot;, \&quot;FETCHED_OBJECT_PROPERTY\&quot;, \&quot;APPEND_OBJECT_PROPERTY\&quot;, \&quot;STATIC_APPEND_VALUE\&quot;, \&quot;ENROLLMENT_EVENT_PROPERTY\&quot; | [default to "INCREMENT"]

## Methods

### NewApiIncrementValue

`func NewApiIncrementValue(incrementAmount float32, type_ string, ) *ApiIncrementValue`

NewApiIncrementValue instantiates a new ApiIncrementValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiIncrementValueWithDefaults

`func NewApiIncrementValueWithDefaults() *ApiIncrementValue`

NewApiIncrementValueWithDefaults instantiates a new ApiIncrementValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncrementAmount

`func (o *ApiIncrementValue) GetIncrementAmount() float32`

GetIncrementAmount returns the IncrementAmount field if non-nil, zero value otherwise.

### GetIncrementAmountOk

`func (o *ApiIncrementValue) GetIncrementAmountOk() (*float32, bool)`

GetIncrementAmountOk returns a tuple with the IncrementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementAmount

`func (o *ApiIncrementValue) SetIncrementAmount(v float32)`

SetIncrementAmount sets IncrementAmount field to given value.


### GetType

`func (o *ApiIncrementValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiIncrementValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiIncrementValue) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



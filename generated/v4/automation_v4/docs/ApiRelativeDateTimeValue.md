# ApiRelativeDateTimeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeDelay** | [**ApiTimeDelay**](ApiTimeDelay.md) |  | 
**Type** | **string** | This is the type of input value. This can be one of: \&quot;FIELD_DATA\&quot;, \&quot;OBJECT_PROPERTY\&quot;, \&quot;STATIC_VALUE\&quot;, \&quot;RELATIVE_DATETIME\&quot;, \&quot;TIMESTAMP\&quot;, \&quot;INCREMENT\&quot;, \&quot;FETCHED_OBJECT_PROPERTY\&quot;, \&quot;APPEND_OBJECT_PROPERTY\&quot;, \&quot;STATIC_APPEND_VALUE\&quot;, \&quot;ENROLLMENT_EVENT_PROPERTY\&quot; | [default to "RELATIVE_DATETIME"]

## Methods

### NewApiRelativeDateTimeValue

`func NewApiRelativeDateTimeValue(timeDelay ApiTimeDelay, type_ string, ) *ApiRelativeDateTimeValue`

NewApiRelativeDateTimeValue instantiates a new ApiRelativeDateTimeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRelativeDateTimeValueWithDefaults

`func NewApiRelativeDateTimeValueWithDefaults() *ApiRelativeDateTimeValue`

NewApiRelativeDateTimeValueWithDefaults instantiates a new ApiRelativeDateTimeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeDelay

`func (o *ApiRelativeDateTimeValue) GetTimeDelay() ApiTimeDelay`

GetTimeDelay returns the TimeDelay field if non-nil, zero value otherwise.

### GetTimeDelayOk

`func (o *ApiRelativeDateTimeValue) GetTimeDelayOk() (*ApiTimeDelay, bool)`

GetTimeDelayOk returns a tuple with the TimeDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeDelay

`func (o *ApiRelativeDateTimeValue) SetTimeDelay(v ApiTimeDelay)`

SetTimeDelay sets TimeDelay field to given value.


### GetType

`func (o *ApiRelativeDateTimeValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiRelativeDateTimeValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiRelativeDateTimeValue) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



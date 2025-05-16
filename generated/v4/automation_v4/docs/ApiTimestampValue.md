# ApiTimestampValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimestampType** | **string** | Currently only EXECUTION_TIME is supported. | 
**Type** | **string** | This is the type of input value. This can be one of: \&quot;FIELD_DATA\&quot;, \&quot;OBJECT_PROPERTY\&quot;, \&quot;STATIC_VALUE\&quot;, \&quot;RELATIVE_DATETIME\&quot;, \&quot;TIMESTAMP\&quot;, \&quot;INCREMENT\&quot;, \&quot;FETCHED_OBJECT_PROPERTY\&quot;, \&quot;APPEND_OBJECT_PROPERTY\&quot;, \&quot;STATIC_APPEND_VALUE\&quot;, \&quot;ENROLLMENT_EVENT_PROPERTY\&quot; | [default to "TIMESTAMP"]

## Methods

### NewApiTimestampValue

`func NewApiTimestampValue(timestampType string, type_ string, ) *ApiTimestampValue`

NewApiTimestampValue instantiates a new ApiTimestampValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTimestampValueWithDefaults

`func NewApiTimestampValueWithDefaults() *ApiTimestampValue`

NewApiTimestampValueWithDefaults instantiates a new ApiTimestampValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestampType

`func (o *ApiTimestampValue) GetTimestampType() string`

GetTimestampType returns the TimestampType field if non-nil, zero value otherwise.

### GetTimestampTypeOk

`func (o *ApiTimestampValue) GetTimestampTypeOk() (*string, bool)`

GetTimestampTypeOk returns a tuple with the TimestampType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampType

`func (o *ApiTimestampValue) SetTimestampType(v string)`

SetTimestampType sets TimestampType field to given value.


### GetType

`func (o *ApiTimestampValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiTimestampValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiTimestampValue) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ApiAppendObjectPropertyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | This is the type of input value. This can be one of: \&quot;FIELD_DATA\&quot;, \&quot;OBJECT_PROPERTY\&quot;, \&quot;STATIC_VALUE\&quot;, \&quot;RELATIVE_DATETIME\&quot;, \&quot;TIMESTAMP\&quot;, \&quot;INCREMENT\&quot;, \&quot;FETCHED_OBJECT_PROPERTY\&quot;, \&quot;APPEND_OBJECT_PROPERTY\&quot;, \&quot;STATIC_APPEND_VALUE\&quot;, \&quot;ENROLLMENT_EVENT_PROPERTY\&quot; | [default to "APPEND_OBJECT_PROPERTY"]
**AppendPropertyName** | **string** | The name of the property to append data from | 

## Methods

### NewApiAppendObjectPropertyValue

`func NewApiAppendObjectPropertyValue(type_ string, appendPropertyName string, ) *ApiAppendObjectPropertyValue`

NewApiAppendObjectPropertyValue instantiates a new ApiAppendObjectPropertyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAppendObjectPropertyValueWithDefaults

`func NewApiAppendObjectPropertyValueWithDefaults() *ApiAppendObjectPropertyValue`

NewApiAppendObjectPropertyValueWithDefaults instantiates a new ApiAppendObjectPropertyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiAppendObjectPropertyValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiAppendObjectPropertyValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiAppendObjectPropertyValue) SetType(v string)`

SetType sets Type field to given value.


### GetAppendPropertyName

`func (o *ApiAppendObjectPropertyValue) GetAppendPropertyName() string`

GetAppendPropertyName returns the AppendPropertyName field if non-nil, zero value otherwise.

### GetAppendPropertyNameOk

`func (o *ApiAppendObjectPropertyValue) GetAppendPropertyNameOk() (*string, bool)`

GetAppendPropertyNameOk returns a tuple with the AppendPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendPropertyName

`func (o *ApiAppendObjectPropertyValue) SetAppendPropertyName(v string)`

SetAppendPropertyName sets AppendPropertyName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



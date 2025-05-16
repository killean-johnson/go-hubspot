# ApiStaticAppendValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticAppendValue** | **string** | The value to append | 
**Type** | **string** | This is the type of input value. This can be one of: \&quot;FIELD_DATA\&quot;, \&quot;OBJECT_PROPERTY\&quot;, \&quot;STATIC_VALUE\&quot;, \&quot;RELATIVE_DATETIME\&quot;, \&quot;TIMESTAMP\&quot;, \&quot;INCREMENT\&quot;, \&quot;FETCHED_OBJECT_PROPERTY\&quot;, \&quot;APPEND_OBJECT_PROPERTY\&quot;, \&quot;STATIC_APPEND_VALUE\&quot;, \&quot;ENROLLMENT_EVENT_PROPERTY\&quot; | [default to "STATIC_APPEND_VALUE"]

## Methods

### NewApiStaticAppendValue

`func NewApiStaticAppendValue(staticAppendValue string, type_ string, ) *ApiStaticAppendValue`

NewApiStaticAppendValue instantiates a new ApiStaticAppendValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStaticAppendValueWithDefaults

`func NewApiStaticAppendValueWithDefaults() *ApiStaticAppendValue`

NewApiStaticAppendValueWithDefaults instantiates a new ApiStaticAppendValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticAppendValue

`func (o *ApiStaticAppendValue) GetStaticAppendValue() string`

GetStaticAppendValue returns the StaticAppendValue field if non-nil, zero value otherwise.

### GetStaticAppendValueOk

`func (o *ApiStaticAppendValue) GetStaticAppendValueOk() (*string, bool)`

GetStaticAppendValueOk returns a tuple with the StaticAppendValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticAppendValue

`func (o *ApiStaticAppendValue) SetStaticAppendValue(v string)`

SetStaticAppendValue sets StaticAppendValue field to given value.


### GetType

`func (o *ApiStaticAppendValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiStaticAppendValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiStaticAppendValue) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



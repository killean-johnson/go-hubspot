# ApiFetchedObjectPropertyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyToken** | **string** | The token to use to identify the object property to use | 
**Type** | **string** | This is the type of input value. This can be one of: \&quot;FIELD_DATA\&quot;, \&quot;OBJECT_PROPERTY\&quot;, \&quot;STATIC_VALUE\&quot;, \&quot;RELATIVE_DATETIME\&quot;, \&quot;TIMESTAMP\&quot;, \&quot;INCREMENT\&quot;, \&quot;FETCHED_OBJECT_PROPERTY\&quot;, \&quot;APPEND_OBJECT_PROPERTY\&quot;, \&quot;STATIC_APPEND_VALUE\&quot;, \&quot;ENROLLMENT_EVENT_PROPERTY\&quot; | [default to "FETCHED_OBJECT_PROPERTY"]

## Methods

### NewApiFetchedObjectPropertyValue

`func NewApiFetchedObjectPropertyValue(propertyToken string, type_ string, ) *ApiFetchedObjectPropertyValue`

NewApiFetchedObjectPropertyValue instantiates a new ApiFetchedObjectPropertyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFetchedObjectPropertyValueWithDefaults

`func NewApiFetchedObjectPropertyValueWithDefaults() *ApiFetchedObjectPropertyValue`

NewApiFetchedObjectPropertyValueWithDefaults instantiates a new ApiFetchedObjectPropertyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyToken

`func (o *ApiFetchedObjectPropertyValue) GetPropertyToken() string`

GetPropertyToken returns the PropertyToken field if non-nil, zero value otherwise.

### GetPropertyTokenOk

`func (o *ApiFetchedObjectPropertyValue) GetPropertyTokenOk() (*string, bool)`

GetPropertyTokenOk returns a tuple with the PropertyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyToken

`func (o *ApiFetchedObjectPropertyValue) SetPropertyToken(v string)`

SetPropertyToken sets PropertyToken field to given value.


### GetType

`func (o *ApiFetchedObjectPropertyValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiFetchedObjectPropertyValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiFetchedObjectPropertyValue) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



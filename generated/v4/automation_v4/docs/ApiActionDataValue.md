# ApiActionDataValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataKey** | **string** | The output field name for that action | 
**ActionId** | **string** | Which action to pull data from. | 
**Type** | **string** | This is the type of input value. This can be one of: \&quot;FIELD_DATA\&quot;, \&quot;OBJECT_PROPERTY\&quot;, \&quot;STATIC_VALUE\&quot;, \&quot;RELATIVE_DATETIME\&quot;, \&quot;TIMESTAMP\&quot;, \&quot;INCREMENT\&quot;, \&quot;FETCHED_OBJECT_PROPERTY\&quot;, \&quot;APPEND_OBJECT_PROPERTY\&quot;, \&quot;STATIC_APPEND_VALUE\&quot;, \&quot;ENROLLMENT_EVENT_PROPERTY\&quot; | [default to "FIELD_DATA"]

## Methods

### NewApiActionDataValue

`func NewApiActionDataValue(dataKey string, actionId string, type_ string, ) *ApiActionDataValue`

NewApiActionDataValue instantiates a new ApiActionDataValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiActionDataValueWithDefaults

`func NewApiActionDataValueWithDefaults() *ApiActionDataValue`

NewApiActionDataValueWithDefaults instantiates a new ApiActionDataValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataKey

`func (o *ApiActionDataValue) GetDataKey() string`

GetDataKey returns the DataKey field if non-nil, zero value otherwise.

### GetDataKeyOk

`func (o *ApiActionDataValue) GetDataKeyOk() (*string, bool)`

GetDataKeyOk returns a tuple with the DataKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKey

`func (o *ApiActionDataValue) SetDataKey(v string)`

SetDataKey sets DataKey field to given value.


### GetActionId

`func (o *ApiActionDataValue) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ApiActionDataValue) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ApiActionDataValue) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetType

`func (o *ApiActionDataValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiActionDataValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiActionDataValue) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



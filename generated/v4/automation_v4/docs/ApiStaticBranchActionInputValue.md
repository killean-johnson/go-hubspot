# ApiStaticBranchActionInputValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataKey** | **string** | The output field name for that action | 
**ActionId** | **string** | Which action to pull data from. | 
**Type** | **string** | This is the type of input value. This can be one of: \&quot;FIELD_DATA\&quot;, \&quot;OBJECT_PROPERTY\&quot;, \&quot;STATIC_VALUE\&quot;, \&quot;RELATIVE_DATETIME\&quot;, \&quot;TIMESTAMP\&quot;, \&quot;INCREMENT\&quot;, \&quot;FETCHED_OBJECT_PROPERTY\&quot;, \&quot;APPEND_OBJECT_PROPERTY\&quot;, \&quot;STATIC_APPEND_VALUE\&quot;, \&quot;ENROLLMENT_EVENT_PROPERTY\&quot; | [default to "FIELD_DATA"]
**PropertyName** | **string** | The property name to pull data from. | 
**StaticValue** | **string** | A static value to use as the input | 
**TimeDelay** | [**ApiTimeDelay**](ApiTimeDelay.md) |  | 
**TimestampType** | **string** | Currently only EXECUTION_TIME is supported. | 
**IncrementAmount** | **float32** | The amount be which to increment | 
**PropertyToken** | **string** | The token to use to identify the object property to use | 
**AppendPropertyName** | **string** | The name of the property to append data from | 
**StaticAppendValue** | **string** | The value to append | 
**EnrollmentEventPropertyToken** | **string** |  | 

## Methods

### NewApiStaticBranchActionInputValue

`func NewApiStaticBranchActionInputValue(dataKey string, actionId string, type_ string, propertyName string, staticValue string, timeDelay ApiTimeDelay, timestampType string, incrementAmount float32, propertyToken string, appendPropertyName string, staticAppendValue string, enrollmentEventPropertyToken string, ) *ApiStaticBranchActionInputValue`

NewApiStaticBranchActionInputValue instantiates a new ApiStaticBranchActionInputValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStaticBranchActionInputValueWithDefaults

`func NewApiStaticBranchActionInputValueWithDefaults() *ApiStaticBranchActionInputValue`

NewApiStaticBranchActionInputValueWithDefaults instantiates a new ApiStaticBranchActionInputValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataKey

`func (o *ApiStaticBranchActionInputValue) GetDataKey() string`

GetDataKey returns the DataKey field if non-nil, zero value otherwise.

### GetDataKeyOk

`func (o *ApiStaticBranchActionInputValue) GetDataKeyOk() (*string, bool)`

GetDataKeyOk returns a tuple with the DataKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKey

`func (o *ApiStaticBranchActionInputValue) SetDataKey(v string)`

SetDataKey sets DataKey field to given value.


### GetActionId

`func (o *ApiStaticBranchActionInputValue) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ApiStaticBranchActionInputValue) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ApiStaticBranchActionInputValue) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetType

`func (o *ApiStaticBranchActionInputValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiStaticBranchActionInputValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiStaticBranchActionInputValue) SetType(v string)`

SetType sets Type field to given value.


### GetPropertyName

`func (o *ApiStaticBranchActionInputValue) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *ApiStaticBranchActionInputValue) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *ApiStaticBranchActionInputValue) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.


### GetStaticValue

`func (o *ApiStaticBranchActionInputValue) GetStaticValue() string`

GetStaticValue returns the StaticValue field if non-nil, zero value otherwise.

### GetStaticValueOk

`func (o *ApiStaticBranchActionInputValue) GetStaticValueOk() (*string, bool)`

GetStaticValueOk returns a tuple with the StaticValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticValue

`func (o *ApiStaticBranchActionInputValue) SetStaticValue(v string)`

SetStaticValue sets StaticValue field to given value.


### GetTimeDelay

`func (o *ApiStaticBranchActionInputValue) GetTimeDelay() ApiTimeDelay`

GetTimeDelay returns the TimeDelay field if non-nil, zero value otherwise.

### GetTimeDelayOk

`func (o *ApiStaticBranchActionInputValue) GetTimeDelayOk() (*ApiTimeDelay, bool)`

GetTimeDelayOk returns a tuple with the TimeDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeDelay

`func (o *ApiStaticBranchActionInputValue) SetTimeDelay(v ApiTimeDelay)`

SetTimeDelay sets TimeDelay field to given value.


### GetTimestampType

`func (o *ApiStaticBranchActionInputValue) GetTimestampType() string`

GetTimestampType returns the TimestampType field if non-nil, zero value otherwise.

### GetTimestampTypeOk

`func (o *ApiStaticBranchActionInputValue) GetTimestampTypeOk() (*string, bool)`

GetTimestampTypeOk returns a tuple with the TimestampType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampType

`func (o *ApiStaticBranchActionInputValue) SetTimestampType(v string)`

SetTimestampType sets TimestampType field to given value.


### GetIncrementAmount

`func (o *ApiStaticBranchActionInputValue) GetIncrementAmount() float32`

GetIncrementAmount returns the IncrementAmount field if non-nil, zero value otherwise.

### GetIncrementAmountOk

`func (o *ApiStaticBranchActionInputValue) GetIncrementAmountOk() (*float32, bool)`

GetIncrementAmountOk returns a tuple with the IncrementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementAmount

`func (o *ApiStaticBranchActionInputValue) SetIncrementAmount(v float32)`

SetIncrementAmount sets IncrementAmount field to given value.


### GetPropertyToken

`func (o *ApiStaticBranchActionInputValue) GetPropertyToken() string`

GetPropertyToken returns the PropertyToken field if non-nil, zero value otherwise.

### GetPropertyTokenOk

`func (o *ApiStaticBranchActionInputValue) GetPropertyTokenOk() (*string, bool)`

GetPropertyTokenOk returns a tuple with the PropertyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyToken

`func (o *ApiStaticBranchActionInputValue) SetPropertyToken(v string)`

SetPropertyToken sets PropertyToken field to given value.


### GetAppendPropertyName

`func (o *ApiStaticBranchActionInputValue) GetAppendPropertyName() string`

GetAppendPropertyName returns the AppendPropertyName field if non-nil, zero value otherwise.

### GetAppendPropertyNameOk

`func (o *ApiStaticBranchActionInputValue) GetAppendPropertyNameOk() (*string, bool)`

GetAppendPropertyNameOk returns a tuple with the AppendPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendPropertyName

`func (o *ApiStaticBranchActionInputValue) SetAppendPropertyName(v string)`

SetAppendPropertyName sets AppendPropertyName field to given value.


### GetStaticAppendValue

`func (o *ApiStaticBranchActionInputValue) GetStaticAppendValue() string`

GetStaticAppendValue returns the StaticAppendValue field if non-nil, zero value otherwise.

### GetStaticAppendValueOk

`func (o *ApiStaticBranchActionInputValue) GetStaticAppendValueOk() (*string, bool)`

GetStaticAppendValueOk returns a tuple with the StaticAppendValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticAppendValue

`func (o *ApiStaticBranchActionInputValue) SetStaticAppendValue(v string)`

SetStaticAppendValue sets StaticAppendValue field to given value.


### GetEnrollmentEventPropertyToken

`func (o *ApiStaticBranchActionInputValue) GetEnrollmentEventPropertyToken() string`

GetEnrollmentEventPropertyToken returns the EnrollmentEventPropertyToken field if non-nil, zero value otherwise.

### GetEnrollmentEventPropertyTokenOk

`func (o *ApiStaticBranchActionInputValue) GetEnrollmentEventPropertyTokenOk() (*string, bool)`

GetEnrollmentEventPropertyTokenOk returns a tuple with the EnrollmentEventPropertyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentEventPropertyToken

`func (o *ApiStaticBranchActionInputValue) SetEnrollmentEventPropertyToken(v string)`

SetEnrollmentEventPropertyToken sets EnrollmentEventPropertyToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



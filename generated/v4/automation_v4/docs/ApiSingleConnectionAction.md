# ApiSingleConnectionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionTypeId** | **string** | The ID of the actionType to use. | 
**ActionId** | **string** | The ID for this action. | 
**Connection** | Pointer to [**ApiConnection**](ApiConnection.md) |  | [optional] 
**Type** | **string** | The type of action this is, can be: \&quot;STATIC_BRANCH\&quot;, \&quot;LIST_BRANCH\&quot;, \&quot;AB_TEST_BRANCH\&quot;, \&quot;CUSTOM_CODE\&quot;, \&quot;WEBHOOK\&quot;, or \&quot;SINGLE_CONNECTION\&quot; | [default to "SINGLE_CONNECTION"]
**ActionTypeVersion** | **int32** | The version of this actionType to use. | 
**Fields** | **map[string]map[string]interface{}** | The fields to pass into this action. Different action types accept different fields. | 

## Methods

### NewApiSingleConnectionAction

`func NewApiSingleConnectionAction(actionTypeId string, actionId string, type_ string, actionTypeVersion int32, fields map[string]map[string]interface{}, ) *ApiSingleConnectionAction`

NewApiSingleConnectionAction instantiates a new ApiSingleConnectionAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSingleConnectionActionWithDefaults

`func NewApiSingleConnectionActionWithDefaults() *ApiSingleConnectionAction`

NewApiSingleConnectionActionWithDefaults instantiates a new ApiSingleConnectionAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionTypeId

`func (o *ApiSingleConnectionAction) GetActionTypeId() string`

GetActionTypeId returns the ActionTypeId field if non-nil, zero value otherwise.

### GetActionTypeIdOk

`func (o *ApiSingleConnectionAction) GetActionTypeIdOk() (*string, bool)`

GetActionTypeIdOk returns a tuple with the ActionTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeId

`func (o *ApiSingleConnectionAction) SetActionTypeId(v string)`

SetActionTypeId sets ActionTypeId field to given value.


### GetActionId

`func (o *ApiSingleConnectionAction) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ApiSingleConnectionAction) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ApiSingleConnectionAction) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetConnection

`func (o *ApiSingleConnectionAction) GetConnection() ApiConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ApiSingleConnectionAction) GetConnectionOk() (*ApiConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ApiSingleConnectionAction) SetConnection(v ApiConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ApiSingleConnectionAction) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetType

`func (o *ApiSingleConnectionAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiSingleConnectionAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiSingleConnectionAction) SetType(v string)`

SetType sets Type field to given value.


### GetActionTypeVersion

`func (o *ApiSingleConnectionAction) GetActionTypeVersion() int32`

GetActionTypeVersion returns the ActionTypeVersion field if non-nil, zero value otherwise.

### GetActionTypeVersionOk

`func (o *ApiSingleConnectionAction) GetActionTypeVersionOk() (*int32, bool)`

GetActionTypeVersionOk returns a tuple with the ActionTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeVersion

`func (o *ApiSingleConnectionAction) SetActionTypeVersion(v int32)`

SetActionTypeVersion sets ActionTypeVersion field to given value.


### GetFields

`func (o *ApiSingleConnectionAction) GetFields() map[string]map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ApiSingleConnectionAction) GetFieldsOk() (*map[string]map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ApiSingleConnectionAction) SetFields(v map[string]map[string]interface{})`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



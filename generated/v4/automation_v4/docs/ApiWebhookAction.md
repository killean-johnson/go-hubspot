# ApiWebhookAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The HTTP method to use when calling the webhook URL | 
**QueryParams** | [**[]ApiInputVariable**](ApiInputVariable.md) |  | 
**ActionId** | **string** | The ID for this action. | 
**Connection** | Pointer to [**ApiConnection**](ApiConnection.md) |  | [optional] 
**Type** | **string** | The type of action this is, can be: \&quot;STATIC_BRANCH\&quot;, \&quot;LIST_BRANCH\&quot;, \&quot;AB_TEST_BRANCH\&quot;, \&quot;CUSTOM_CODE\&quot;, \&quot;WEBHOOK\&quot;, or \&quot;SINGLE_CONNECTION\&quot; | [default to "WEBHOOK"]
**AuthSettings** | Pointer to [**ApiWebhookActionAuthSettings**](ApiWebhookActionAuthSettings.md) |  | [optional] 
**WebhookUrl** | **string** | The URL to call each time this action is executed. | 

## Methods

### NewApiWebhookAction

`func NewApiWebhookAction(method string, queryParams []ApiInputVariable, actionId string, type_ string, webhookUrl string, ) *ApiWebhookAction`

NewApiWebhookAction instantiates a new ApiWebhookAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWebhookActionWithDefaults

`func NewApiWebhookActionWithDefaults() *ApiWebhookAction`

NewApiWebhookActionWithDefaults instantiates a new ApiWebhookAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *ApiWebhookAction) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApiWebhookAction) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApiWebhookAction) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetQueryParams

`func (o *ApiWebhookAction) GetQueryParams() []ApiInputVariable`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *ApiWebhookAction) GetQueryParamsOk() (*[]ApiInputVariable, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *ApiWebhookAction) SetQueryParams(v []ApiInputVariable)`

SetQueryParams sets QueryParams field to given value.


### GetActionId

`func (o *ApiWebhookAction) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ApiWebhookAction) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ApiWebhookAction) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetConnection

`func (o *ApiWebhookAction) GetConnection() ApiConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ApiWebhookAction) GetConnectionOk() (*ApiConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ApiWebhookAction) SetConnection(v ApiConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ApiWebhookAction) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetType

`func (o *ApiWebhookAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiWebhookAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiWebhookAction) SetType(v string)`

SetType sets Type field to given value.


### GetAuthSettings

`func (o *ApiWebhookAction) GetAuthSettings() ApiWebhookActionAuthSettings`

GetAuthSettings returns the AuthSettings field if non-nil, zero value otherwise.

### GetAuthSettingsOk

`func (o *ApiWebhookAction) GetAuthSettingsOk() (*ApiWebhookActionAuthSettings, bool)`

GetAuthSettingsOk returns a tuple with the AuthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSettings

`func (o *ApiWebhookAction) SetAuthSettings(v ApiWebhookActionAuthSettings)`

SetAuthSettings sets AuthSettings field to given value.

### HasAuthSettings

`func (o *ApiWebhookAction) HasAuthSettings() bool`

HasAuthSettings returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *ApiWebhookAction) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ApiWebhookAction) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ApiWebhookAction) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



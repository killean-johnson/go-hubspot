# ApiPlatformFlowCreateRequestAllOfActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticBranches** | [**[]ApiStaticBranch**](ApiStaticBranch.md) |  | 
**DefaultBranchName** | Pointer to **string** | The name of the default branch, the branch that gets executed if the object does not match any of the &#x60;listBranch&#x60; criteria. | [optional] 
**DefaultBranch** | Pointer to [**ApiConnection**](ApiConnection.md) |  | [optional] 
**InputValue** | [**ApiStaticBranchActionInputValue**](ApiStaticBranchActionInputValue.md) |  | 
**ActionId** | **string** | The ID for this action. | 
**Type** | **string** | The type of action this is, can be: \&quot;STATIC_BRANCH\&quot;, \&quot;LIST_BRANCH\&quot;, \&quot;AB_TEST_BRANCH\&quot;, \&quot;CUSTOM_CODE\&quot;, \&quot;WEBHOOK\&quot;, or \&quot;SINGLE_CONNECTION\&quot; | [default to "STATIC_BRANCH"]
**ListBranches** | [**[]ApiListBranch**](ApiListBranch.md) |  | 
**TestBranches** | [**[]ApiConnection**](ApiConnection.md) |  | 
**InputFields** | [**[]ApiInputVariable**](ApiInputVariable.md) |  | 
**OutputFields** | [**[]ApiCustomCodeActionOutputFieldsInner**](ApiCustomCodeActionOutputFieldsInner.md) | The list of output fields that this custom action makes available to the rest of the flow. | 
**SourceCode** | **string** | The source code to execute when this action executes. | 
**Runtime** | **string** | The runtime to use to execute the source code. Supported runtimes are: \&quot;NODE16X\&quot;, \&quot;NODE20X\&quot;, \&quot;PYTHON39\&quot; | 
**Connection** | Pointer to [**ApiConnection**](ApiConnection.md) |  | [optional] 
**SecretNames** | **[]string** | The names of any \&quot;secrets\&quot; setup in this portal that will be used in this action. | 
**Method** | **string** | The HTTP method to use when calling the webhook URL | 
**QueryParams** | [**[]ApiInputVariable**](ApiInputVariable.md) |  | 
**AuthSettings** | Pointer to [**ApiWebhookActionAuthSettings**](ApiWebhookActionAuthSettings.md) |  | [optional] 
**WebhookUrl** | **string** | The URL to call each time this action is executed. | 
**ActionTypeId** | **string** | The ID of the actionType to use. | 
**ActionTypeVersion** | **int32** | The version of this actionType to use. | 
**Fields** | **map[string]map[string]interface{}** | The fields to pass into this action. Different action types accept different fields. | 

## Methods

### NewApiPlatformFlowCreateRequestAllOfActions

`func NewApiPlatformFlowCreateRequestAllOfActions(staticBranches []ApiStaticBranch, inputValue ApiStaticBranchActionInputValue, actionId string, type_ string, listBranches []ApiListBranch, testBranches []ApiConnection, inputFields []ApiInputVariable, outputFields []ApiCustomCodeActionOutputFieldsInner, sourceCode string, runtime string, secretNames []string, method string, queryParams []ApiInputVariable, webhookUrl string, actionTypeId string, actionTypeVersion int32, fields map[string]map[string]interface{}, ) *ApiPlatformFlowCreateRequestAllOfActions`

NewApiPlatformFlowCreateRequestAllOfActions instantiates a new ApiPlatformFlowCreateRequestAllOfActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPlatformFlowCreateRequestAllOfActionsWithDefaults

`func NewApiPlatformFlowCreateRequestAllOfActionsWithDefaults() *ApiPlatformFlowCreateRequestAllOfActions`

NewApiPlatformFlowCreateRequestAllOfActionsWithDefaults instantiates a new ApiPlatformFlowCreateRequestAllOfActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticBranches

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetStaticBranches() []ApiStaticBranch`

GetStaticBranches returns the StaticBranches field if non-nil, zero value otherwise.

### GetStaticBranchesOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetStaticBranchesOk() (*[]ApiStaticBranch, bool)`

GetStaticBranchesOk returns a tuple with the StaticBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticBranches

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetStaticBranches(v []ApiStaticBranch)`

SetStaticBranches sets StaticBranches field to given value.


### GetDefaultBranchName

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetDefaultBranchName() string`

GetDefaultBranchName returns the DefaultBranchName field if non-nil, zero value otherwise.

### GetDefaultBranchNameOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetDefaultBranchNameOk() (*string, bool)`

GetDefaultBranchNameOk returns a tuple with the DefaultBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchName

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetDefaultBranchName(v string)`

SetDefaultBranchName sets DefaultBranchName field to given value.

### HasDefaultBranchName

`func (o *ApiPlatformFlowCreateRequestAllOfActions) HasDefaultBranchName() bool`

HasDefaultBranchName returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetDefaultBranch() ApiConnection`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetDefaultBranchOk() (*ApiConnection, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetDefaultBranch(v ApiConnection)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *ApiPlatformFlowCreateRequestAllOfActions) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetInputValue

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetInputValue() ApiStaticBranchActionInputValue`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetInputValueOk() (*ApiStaticBranchActionInputValue, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetInputValue(v ApiStaticBranchActionInputValue)`

SetInputValue sets InputValue field to given value.


### GetActionId

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetType

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetType(v string)`

SetType sets Type field to given value.


### GetListBranches

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetListBranches() []ApiListBranch`

GetListBranches returns the ListBranches field if non-nil, zero value otherwise.

### GetListBranchesOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetListBranchesOk() (*[]ApiListBranch, bool)`

GetListBranchesOk returns a tuple with the ListBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListBranches

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetListBranches(v []ApiListBranch)`

SetListBranches sets ListBranches field to given value.


### GetTestBranches

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetTestBranches() []ApiConnection`

GetTestBranches returns the TestBranches field if non-nil, zero value otherwise.

### GetTestBranchesOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetTestBranchesOk() (*[]ApiConnection, bool)`

GetTestBranchesOk returns a tuple with the TestBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestBranches

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetTestBranches(v []ApiConnection)`

SetTestBranches sets TestBranches field to given value.


### GetInputFields

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetInputFields() []ApiInputVariable`

GetInputFields returns the InputFields field if non-nil, zero value otherwise.

### GetInputFieldsOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetInputFieldsOk() (*[]ApiInputVariable, bool)`

GetInputFieldsOk returns a tuple with the InputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFields

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetInputFields(v []ApiInputVariable)`

SetInputFields sets InputFields field to given value.


### GetOutputFields

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetOutputFields() []ApiCustomCodeActionOutputFieldsInner`

GetOutputFields returns the OutputFields field if non-nil, zero value otherwise.

### GetOutputFieldsOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetOutputFieldsOk() (*[]ApiCustomCodeActionOutputFieldsInner, bool)`

GetOutputFieldsOk returns a tuple with the OutputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFields

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetOutputFields(v []ApiCustomCodeActionOutputFieldsInner)`

SetOutputFields sets OutputFields field to given value.


### GetSourceCode

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetSourceCode() string`

GetSourceCode returns the SourceCode field if non-nil, zero value otherwise.

### GetSourceCodeOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetSourceCodeOk() (*string, bool)`

GetSourceCodeOk returns a tuple with the SourceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCode

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetSourceCode(v string)`

SetSourceCode sets SourceCode field to given value.


### GetRuntime

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.


### GetConnection

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetConnection() ApiConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetConnectionOk() (*ApiConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetConnection(v ApiConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ApiPlatformFlowCreateRequestAllOfActions) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetSecretNames

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetSecretNames() []string`

GetSecretNames returns the SecretNames field if non-nil, zero value otherwise.

### GetSecretNamesOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetSecretNamesOk() (*[]string, bool)`

GetSecretNamesOk returns a tuple with the SecretNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretNames

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetSecretNames(v []string)`

SetSecretNames sets SecretNames field to given value.


### GetMethod

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetQueryParams

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetQueryParams() []ApiInputVariable`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetQueryParamsOk() (*[]ApiInputVariable, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetQueryParams(v []ApiInputVariable)`

SetQueryParams sets QueryParams field to given value.


### GetAuthSettings

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetAuthSettings() ApiWebhookActionAuthSettings`

GetAuthSettings returns the AuthSettings field if non-nil, zero value otherwise.

### GetAuthSettingsOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetAuthSettingsOk() (*ApiWebhookActionAuthSettings, bool)`

GetAuthSettingsOk returns a tuple with the AuthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSettings

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetAuthSettings(v ApiWebhookActionAuthSettings)`

SetAuthSettings sets AuthSettings field to given value.

### HasAuthSettings

`func (o *ApiPlatformFlowCreateRequestAllOfActions) HasAuthSettings() bool`

HasAuthSettings returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### GetActionTypeId

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetActionTypeId() string`

GetActionTypeId returns the ActionTypeId field if non-nil, zero value otherwise.

### GetActionTypeIdOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetActionTypeIdOk() (*string, bool)`

GetActionTypeIdOk returns a tuple with the ActionTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeId

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetActionTypeId(v string)`

SetActionTypeId sets ActionTypeId field to given value.


### GetActionTypeVersion

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetActionTypeVersion() int32`

GetActionTypeVersion returns the ActionTypeVersion field if non-nil, zero value otherwise.

### GetActionTypeVersionOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetActionTypeVersionOk() (*int32, bool)`

GetActionTypeVersionOk returns a tuple with the ActionTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeVersion

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetActionTypeVersion(v int32)`

SetActionTypeVersion sets ActionTypeVersion field to given value.


### GetFields

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetFields() map[string]map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ApiPlatformFlowCreateRequestAllOfActions) GetFieldsOk() (*map[string]map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ApiPlatformFlowCreateRequestAllOfActions) SetFields(v map[string]map[string]interface{})`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



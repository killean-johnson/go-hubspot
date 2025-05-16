# ApiContactFlowAllOfActions

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

### NewApiContactFlowAllOfActions

`func NewApiContactFlowAllOfActions(staticBranches []ApiStaticBranch, inputValue ApiStaticBranchActionInputValue, actionId string, type_ string, listBranches []ApiListBranch, testBranches []ApiConnection, inputFields []ApiInputVariable, outputFields []ApiCustomCodeActionOutputFieldsInner, sourceCode string, runtime string, secretNames []string, method string, queryParams []ApiInputVariable, webhookUrl string, actionTypeId string, actionTypeVersion int32, fields map[string]map[string]interface{}, ) *ApiContactFlowAllOfActions`

NewApiContactFlowAllOfActions instantiates a new ApiContactFlowAllOfActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiContactFlowAllOfActionsWithDefaults

`func NewApiContactFlowAllOfActionsWithDefaults() *ApiContactFlowAllOfActions`

NewApiContactFlowAllOfActionsWithDefaults instantiates a new ApiContactFlowAllOfActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticBranches

`func (o *ApiContactFlowAllOfActions) GetStaticBranches() []ApiStaticBranch`

GetStaticBranches returns the StaticBranches field if non-nil, zero value otherwise.

### GetStaticBranchesOk

`func (o *ApiContactFlowAllOfActions) GetStaticBranchesOk() (*[]ApiStaticBranch, bool)`

GetStaticBranchesOk returns a tuple with the StaticBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticBranches

`func (o *ApiContactFlowAllOfActions) SetStaticBranches(v []ApiStaticBranch)`

SetStaticBranches sets StaticBranches field to given value.


### GetDefaultBranchName

`func (o *ApiContactFlowAllOfActions) GetDefaultBranchName() string`

GetDefaultBranchName returns the DefaultBranchName field if non-nil, zero value otherwise.

### GetDefaultBranchNameOk

`func (o *ApiContactFlowAllOfActions) GetDefaultBranchNameOk() (*string, bool)`

GetDefaultBranchNameOk returns a tuple with the DefaultBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchName

`func (o *ApiContactFlowAllOfActions) SetDefaultBranchName(v string)`

SetDefaultBranchName sets DefaultBranchName field to given value.

### HasDefaultBranchName

`func (o *ApiContactFlowAllOfActions) HasDefaultBranchName() bool`

HasDefaultBranchName returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *ApiContactFlowAllOfActions) GetDefaultBranch() ApiConnection`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *ApiContactFlowAllOfActions) GetDefaultBranchOk() (*ApiConnection, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *ApiContactFlowAllOfActions) SetDefaultBranch(v ApiConnection)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *ApiContactFlowAllOfActions) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetInputValue

`func (o *ApiContactFlowAllOfActions) GetInputValue() ApiStaticBranchActionInputValue`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *ApiContactFlowAllOfActions) GetInputValueOk() (*ApiStaticBranchActionInputValue, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *ApiContactFlowAllOfActions) SetInputValue(v ApiStaticBranchActionInputValue)`

SetInputValue sets InputValue field to given value.


### GetActionId

`func (o *ApiContactFlowAllOfActions) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ApiContactFlowAllOfActions) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ApiContactFlowAllOfActions) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetType

`func (o *ApiContactFlowAllOfActions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiContactFlowAllOfActions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiContactFlowAllOfActions) SetType(v string)`

SetType sets Type field to given value.


### GetListBranches

`func (o *ApiContactFlowAllOfActions) GetListBranches() []ApiListBranch`

GetListBranches returns the ListBranches field if non-nil, zero value otherwise.

### GetListBranchesOk

`func (o *ApiContactFlowAllOfActions) GetListBranchesOk() (*[]ApiListBranch, bool)`

GetListBranchesOk returns a tuple with the ListBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListBranches

`func (o *ApiContactFlowAllOfActions) SetListBranches(v []ApiListBranch)`

SetListBranches sets ListBranches field to given value.


### GetTestBranches

`func (o *ApiContactFlowAllOfActions) GetTestBranches() []ApiConnection`

GetTestBranches returns the TestBranches field if non-nil, zero value otherwise.

### GetTestBranchesOk

`func (o *ApiContactFlowAllOfActions) GetTestBranchesOk() (*[]ApiConnection, bool)`

GetTestBranchesOk returns a tuple with the TestBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestBranches

`func (o *ApiContactFlowAllOfActions) SetTestBranches(v []ApiConnection)`

SetTestBranches sets TestBranches field to given value.


### GetInputFields

`func (o *ApiContactFlowAllOfActions) GetInputFields() []ApiInputVariable`

GetInputFields returns the InputFields field if non-nil, zero value otherwise.

### GetInputFieldsOk

`func (o *ApiContactFlowAllOfActions) GetInputFieldsOk() (*[]ApiInputVariable, bool)`

GetInputFieldsOk returns a tuple with the InputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFields

`func (o *ApiContactFlowAllOfActions) SetInputFields(v []ApiInputVariable)`

SetInputFields sets InputFields field to given value.


### GetOutputFields

`func (o *ApiContactFlowAllOfActions) GetOutputFields() []ApiCustomCodeActionOutputFieldsInner`

GetOutputFields returns the OutputFields field if non-nil, zero value otherwise.

### GetOutputFieldsOk

`func (o *ApiContactFlowAllOfActions) GetOutputFieldsOk() (*[]ApiCustomCodeActionOutputFieldsInner, bool)`

GetOutputFieldsOk returns a tuple with the OutputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFields

`func (o *ApiContactFlowAllOfActions) SetOutputFields(v []ApiCustomCodeActionOutputFieldsInner)`

SetOutputFields sets OutputFields field to given value.


### GetSourceCode

`func (o *ApiContactFlowAllOfActions) GetSourceCode() string`

GetSourceCode returns the SourceCode field if non-nil, zero value otherwise.

### GetSourceCodeOk

`func (o *ApiContactFlowAllOfActions) GetSourceCodeOk() (*string, bool)`

GetSourceCodeOk returns a tuple with the SourceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCode

`func (o *ApiContactFlowAllOfActions) SetSourceCode(v string)`

SetSourceCode sets SourceCode field to given value.


### GetRuntime

`func (o *ApiContactFlowAllOfActions) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *ApiContactFlowAllOfActions) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *ApiContactFlowAllOfActions) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.


### GetConnection

`func (o *ApiContactFlowAllOfActions) GetConnection() ApiConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ApiContactFlowAllOfActions) GetConnectionOk() (*ApiConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ApiContactFlowAllOfActions) SetConnection(v ApiConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ApiContactFlowAllOfActions) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetSecretNames

`func (o *ApiContactFlowAllOfActions) GetSecretNames() []string`

GetSecretNames returns the SecretNames field if non-nil, zero value otherwise.

### GetSecretNamesOk

`func (o *ApiContactFlowAllOfActions) GetSecretNamesOk() (*[]string, bool)`

GetSecretNamesOk returns a tuple with the SecretNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretNames

`func (o *ApiContactFlowAllOfActions) SetSecretNames(v []string)`

SetSecretNames sets SecretNames field to given value.


### GetMethod

`func (o *ApiContactFlowAllOfActions) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApiContactFlowAllOfActions) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApiContactFlowAllOfActions) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetQueryParams

`func (o *ApiContactFlowAllOfActions) GetQueryParams() []ApiInputVariable`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *ApiContactFlowAllOfActions) GetQueryParamsOk() (*[]ApiInputVariable, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *ApiContactFlowAllOfActions) SetQueryParams(v []ApiInputVariable)`

SetQueryParams sets QueryParams field to given value.


### GetAuthSettings

`func (o *ApiContactFlowAllOfActions) GetAuthSettings() ApiWebhookActionAuthSettings`

GetAuthSettings returns the AuthSettings field if non-nil, zero value otherwise.

### GetAuthSettingsOk

`func (o *ApiContactFlowAllOfActions) GetAuthSettingsOk() (*ApiWebhookActionAuthSettings, bool)`

GetAuthSettingsOk returns a tuple with the AuthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSettings

`func (o *ApiContactFlowAllOfActions) SetAuthSettings(v ApiWebhookActionAuthSettings)`

SetAuthSettings sets AuthSettings field to given value.

### HasAuthSettings

`func (o *ApiContactFlowAllOfActions) HasAuthSettings() bool`

HasAuthSettings returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *ApiContactFlowAllOfActions) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ApiContactFlowAllOfActions) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ApiContactFlowAllOfActions) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### GetActionTypeId

`func (o *ApiContactFlowAllOfActions) GetActionTypeId() string`

GetActionTypeId returns the ActionTypeId field if non-nil, zero value otherwise.

### GetActionTypeIdOk

`func (o *ApiContactFlowAllOfActions) GetActionTypeIdOk() (*string, bool)`

GetActionTypeIdOk returns a tuple with the ActionTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeId

`func (o *ApiContactFlowAllOfActions) SetActionTypeId(v string)`

SetActionTypeId sets ActionTypeId field to given value.


### GetActionTypeVersion

`func (o *ApiContactFlowAllOfActions) GetActionTypeVersion() int32`

GetActionTypeVersion returns the ActionTypeVersion field if non-nil, zero value otherwise.

### GetActionTypeVersionOk

`func (o *ApiContactFlowAllOfActions) GetActionTypeVersionOk() (*int32, bool)`

GetActionTypeVersionOk returns a tuple with the ActionTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeVersion

`func (o *ApiContactFlowAllOfActions) SetActionTypeVersion(v int32)`

SetActionTypeVersion sets ActionTypeVersion field to given value.


### GetFields

`func (o *ApiContactFlowAllOfActions) GetFields() map[string]map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ApiContactFlowAllOfActions) GetFieldsOk() (*map[string]map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ApiContactFlowAllOfActions) SetFields(v map[string]map[string]interface{})`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



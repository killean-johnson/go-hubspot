# ApiCustomCodeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFields** | [**[]ApiInputVariable**](ApiInputVariable.md) |  | 
**OutputFields** | [**[]ApiCustomCodeActionOutputFieldsInner**](ApiCustomCodeActionOutputFieldsInner.md) | The list of output fields that this custom action makes available to the rest of the flow. | 
**SourceCode** | **string** | The source code to execute when this action executes. | 
**ActionId** | **string** | The ID for this action. | 
**Runtime** | **string** | The runtime to use to execute the source code. Supported runtimes are: \&quot;NODE16X\&quot;, \&quot;NODE20X\&quot;, \&quot;PYTHON39\&quot; | 
**Connection** | Pointer to [**ApiConnection**](ApiConnection.md) |  | [optional] 
**SecretNames** | **[]string** | The names of any \&quot;secrets\&quot; setup in this portal that will be used in this action. | 
**Type** | **string** | The type of action this is, can be: \&quot;STATIC_BRANCH\&quot;, \&quot;LIST_BRANCH\&quot;, \&quot;AB_TEST_BRANCH\&quot;, \&quot;CUSTOM_CODE\&quot;, \&quot;WEBHOOK\&quot;, or \&quot;SINGLE_CONNECTION\&quot; | [default to "CUSTOM_CODE"]

## Methods

### NewApiCustomCodeAction

`func NewApiCustomCodeAction(inputFields []ApiInputVariable, outputFields []ApiCustomCodeActionOutputFieldsInner, sourceCode string, actionId string, runtime string, secretNames []string, type_ string, ) *ApiCustomCodeAction`

NewApiCustomCodeAction instantiates a new ApiCustomCodeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCustomCodeActionWithDefaults

`func NewApiCustomCodeActionWithDefaults() *ApiCustomCodeAction`

NewApiCustomCodeActionWithDefaults instantiates a new ApiCustomCodeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFields

`func (o *ApiCustomCodeAction) GetInputFields() []ApiInputVariable`

GetInputFields returns the InputFields field if non-nil, zero value otherwise.

### GetInputFieldsOk

`func (o *ApiCustomCodeAction) GetInputFieldsOk() (*[]ApiInputVariable, bool)`

GetInputFieldsOk returns a tuple with the InputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFields

`func (o *ApiCustomCodeAction) SetInputFields(v []ApiInputVariable)`

SetInputFields sets InputFields field to given value.


### GetOutputFields

`func (o *ApiCustomCodeAction) GetOutputFields() []ApiCustomCodeActionOutputFieldsInner`

GetOutputFields returns the OutputFields field if non-nil, zero value otherwise.

### GetOutputFieldsOk

`func (o *ApiCustomCodeAction) GetOutputFieldsOk() (*[]ApiCustomCodeActionOutputFieldsInner, bool)`

GetOutputFieldsOk returns a tuple with the OutputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFields

`func (o *ApiCustomCodeAction) SetOutputFields(v []ApiCustomCodeActionOutputFieldsInner)`

SetOutputFields sets OutputFields field to given value.


### GetSourceCode

`func (o *ApiCustomCodeAction) GetSourceCode() string`

GetSourceCode returns the SourceCode field if non-nil, zero value otherwise.

### GetSourceCodeOk

`func (o *ApiCustomCodeAction) GetSourceCodeOk() (*string, bool)`

GetSourceCodeOk returns a tuple with the SourceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCode

`func (o *ApiCustomCodeAction) SetSourceCode(v string)`

SetSourceCode sets SourceCode field to given value.


### GetActionId

`func (o *ApiCustomCodeAction) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ApiCustomCodeAction) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ApiCustomCodeAction) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetRuntime

`func (o *ApiCustomCodeAction) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *ApiCustomCodeAction) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *ApiCustomCodeAction) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.


### GetConnection

`func (o *ApiCustomCodeAction) GetConnection() ApiConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ApiCustomCodeAction) GetConnectionOk() (*ApiConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ApiCustomCodeAction) SetConnection(v ApiConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ApiCustomCodeAction) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetSecretNames

`func (o *ApiCustomCodeAction) GetSecretNames() []string`

GetSecretNames returns the SecretNames field if non-nil, zero value otherwise.

### GetSecretNamesOk

`func (o *ApiCustomCodeAction) GetSecretNamesOk() (*[]string, bool)`

GetSecretNamesOk returns a tuple with the SecretNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretNames

`func (o *ApiCustomCodeAction) SetSecretNames(v []string)`

SetSecretNames sets SecretNames field to given value.


### GetType

`func (o *ApiCustomCodeAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiCustomCodeAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiCustomCodeAction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



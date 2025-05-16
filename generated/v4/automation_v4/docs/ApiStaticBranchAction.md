# ApiStaticBranchAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticBranches** | [**[]ApiStaticBranch**](ApiStaticBranch.md) |  | 
**DefaultBranchName** | Pointer to **string** | The name of the default branch, the branch that gets executed if &#x60;inputValue&#x60; does not match any of the &#x60;staticBranches&#x60;. | [optional] 
**DefaultBranch** | Pointer to [**ApiConnection**](ApiConnection.md) |  | [optional] 
**InputValue** | [**ApiStaticBranchActionInputValue**](ApiStaticBranchActionInputValue.md) |  | 
**ActionId** | **string** | The ID for this action. | 
**Type** | **string** | The type of action this is, can be: \&quot;STATIC_BRANCH\&quot;, \&quot;LIST_BRANCH\&quot;, \&quot;AB_TEST_BRANCH\&quot;, \&quot;CUSTOM_CODE\&quot;, \&quot;WEBHOOK\&quot;, or \&quot;SINGLE_CONNECTION\&quot; | [default to "STATIC_BRANCH"]

## Methods

### NewApiStaticBranchAction

`func NewApiStaticBranchAction(staticBranches []ApiStaticBranch, inputValue ApiStaticBranchActionInputValue, actionId string, type_ string, ) *ApiStaticBranchAction`

NewApiStaticBranchAction instantiates a new ApiStaticBranchAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStaticBranchActionWithDefaults

`func NewApiStaticBranchActionWithDefaults() *ApiStaticBranchAction`

NewApiStaticBranchActionWithDefaults instantiates a new ApiStaticBranchAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticBranches

`func (o *ApiStaticBranchAction) GetStaticBranches() []ApiStaticBranch`

GetStaticBranches returns the StaticBranches field if non-nil, zero value otherwise.

### GetStaticBranchesOk

`func (o *ApiStaticBranchAction) GetStaticBranchesOk() (*[]ApiStaticBranch, bool)`

GetStaticBranchesOk returns a tuple with the StaticBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticBranches

`func (o *ApiStaticBranchAction) SetStaticBranches(v []ApiStaticBranch)`

SetStaticBranches sets StaticBranches field to given value.


### GetDefaultBranchName

`func (o *ApiStaticBranchAction) GetDefaultBranchName() string`

GetDefaultBranchName returns the DefaultBranchName field if non-nil, zero value otherwise.

### GetDefaultBranchNameOk

`func (o *ApiStaticBranchAction) GetDefaultBranchNameOk() (*string, bool)`

GetDefaultBranchNameOk returns a tuple with the DefaultBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchName

`func (o *ApiStaticBranchAction) SetDefaultBranchName(v string)`

SetDefaultBranchName sets DefaultBranchName field to given value.

### HasDefaultBranchName

`func (o *ApiStaticBranchAction) HasDefaultBranchName() bool`

HasDefaultBranchName returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *ApiStaticBranchAction) GetDefaultBranch() ApiConnection`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *ApiStaticBranchAction) GetDefaultBranchOk() (*ApiConnection, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *ApiStaticBranchAction) SetDefaultBranch(v ApiConnection)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *ApiStaticBranchAction) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetInputValue

`func (o *ApiStaticBranchAction) GetInputValue() ApiStaticBranchActionInputValue`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *ApiStaticBranchAction) GetInputValueOk() (*ApiStaticBranchActionInputValue, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *ApiStaticBranchAction) SetInputValue(v ApiStaticBranchActionInputValue)`

SetInputValue sets InputValue field to given value.


### GetActionId

`func (o *ApiStaticBranchAction) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ApiStaticBranchAction) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ApiStaticBranchAction) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetType

`func (o *ApiStaticBranchAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiStaticBranchAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiStaticBranchAction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



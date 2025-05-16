# ApiListBranchAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultBranchName** | Pointer to **string** | The name of the default branch, the branch that gets executed if the object does not match any of the &#x60;listBranch&#x60; criteria. | [optional] 
**DefaultBranch** | Pointer to [**ApiConnection**](ApiConnection.md) |  | [optional] 
**ActionId** | **string** | The ID for this action. | 
**ListBranches** | [**[]ApiListBranch**](ApiListBranch.md) |  | 
**Type** | **string** | The type of action this is, can be: \&quot;STATIC_BRANCH\&quot;, \&quot;LIST_BRANCH\&quot;, \&quot;AB_TEST_BRANCH\&quot;, \&quot;CUSTOM_CODE\&quot;, \&quot;WEBHOOK\&quot;, or \&quot;SINGLE_CONNECTION\&quot; | [default to "LIST_BRANCH"]

## Methods

### NewApiListBranchAction

`func NewApiListBranchAction(actionId string, listBranches []ApiListBranch, type_ string, ) *ApiListBranchAction`

NewApiListBranchAction instantiates a new ApiListBranchAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiListBranchActionWithDefaults

`func NewApiListBranchActionWithDefaults() *ApiListBranchAction`

NewApiListBranchActionWithDefaults instantiates a new ApiListBranchAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultBranchName

`func (o *ApiListBranchAction) GetDefaultBranchName() string`

GetDefaultBranchName returns the DefaultBranchName field if non-nil, zero value otherwise.

### GetDefaultBranchNameOk

`func (o *ApiListBranchAction) GetDefaultBranchNameOk() (*string, bool)`

GetDefaultBranchNameOk returns a tuple with the DefaultBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchName

`func (o *ApiListBranchAction) SetDefaultBranchName(v string)`

SetDefaultBranchName sets DefaultBranchName field to given value.

### HasDefaultBranchName

`func (o *ApiListBranchAction) HasDefaultBranchName() bool`

HasDefaultBranchName returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *ApiListBranchAction) GetDefaultBranch() ApiConnection`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *ApiListBranchAction) GetDefaultBranchOk() (*ApiConnection, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *ApiListBranchAction) SetDefaultBranch(v ApiConnection)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *ApiListBranchAction) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetActionId

`func (o *ApiListBranchAction) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ApiListBranchAction) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ApiListBranchAction) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetListBranches

`func (o *ApiListBranchAction) GetListBranches() []ApiListBranch`

GetListBranches returns the ListBranches field if non-nil, zero value otherwise.

### GetListBranchesOk

`func (o *ApiListBranchAction) GetListBranchesOk() (*[]ApiListBranch, bool)`

GetListBranchesOk returns a tuple with the ListBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListBranches

`func (o *ApiListBranchAction) SetListBranches(v []ApiListBranch)`

SetListBranches sets ListBranches field to given value.


### GetType

`func (o *ApiListBranchAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiListBranchAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiListBranchAction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ApiABTestBranchAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestBranches** | [**[]ApiConnection**](ApiConnection.md) |  | 
**ActionId** | **string** | The ID for this action. | 
**Type** | **string** | The type of action this is, can be: \&quot;STATIC_BRANCH\&quot;, \&quot;LIST_BRANCH\&quot;, \&quot;AB_TEST_BRANCH\&quot;, \&quot;CUSTOM_CODE\&quot;, \&quot;WEBHOOK\&quot;, or \&quot;SINGLE_CONNECTION\&quot; | [default to "AB_TEST_BRANCH"]

## Methods

### NewApiABTestBranchAction

`func NewApiABTestBranchAction(testBranches []ApiConnection, actionId string, type_ string, ) *ApiABTestBranchAction`

NewApiABTestBranchAction instantiates a new ApiABTestBranchAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiABTestBranchActionWithDefaults

`func NewApiABTestBranchActionWithDefaults() *ApiABTestBranchAction`

NewApiABTestBranchActionWithDefaults instantiates a new ApiABTestBranchAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestBranches

`func (o *ApiABTestBranchAction) GetTestBranches() []ApiConnection`

GetTestBranches returns the TestBranches field if non-nil, zero value otherwise.

### GetTestBranchesOk

`func (o *ApiABTestBranchAction) GetTestBranchesOk() (*[]ApiConnection, bool)`

GetTestBranchesOk returns a tuple with the TestBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestBranches

`func (o *ApiABTestBranchAction) SetTestBranches(v []ApiConnection)`

SetTestBranches sets TestBranches field to given value.


### GetActionId

`func (o *ApiABTestBranchAction) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ApiABTestBranchAction) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ApiABTestBranchAction) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetType

`func (o *ApiABTestBranchAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiABTestBranchAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiABTestBranchAction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



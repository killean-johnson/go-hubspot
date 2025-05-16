# ComboEventRuleBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | **string** |  | 
**ComposingRules** | [**[]ComboEventRule**](ComboEventRule.md) |  | 
**RuleBranches** | [**[]ComboEventRuleBranch**](ComboEventRuleBranch.md) |  | 

## Methods

### NewComboEventRuleBranch

`func NewComboEventRuleBranch(operationType string, composingRules []ComboEventRule, ruleBranches []ComboEventRuleBranch, ) *ComboEventRuleBranch`

NewComboEventRuleBranch instantiates a new ComboEventRuleBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComboEventRuleBranchWithDefaults

`func NewComboEventRuleBranchWithDefaults() *ComboEventRuleBranch`

NewComboEventRuleBranchWithDefaults instantiates a new ComboEventRuleBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *ComboEventRuleBranch) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ComboEventRuleBranch) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ComboEventRuleBranch) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetComposingRules

`func (o *ComboEventRuleBranch) GetComposingRules() []ComboEventRule`

GetComposingRules returns the ComposingRules field if non-nil, zero value otherwise.

### GetComposingRulesOk

`func (o *ComboEventRuleBranch) GetComposingRulesOk() (*[]ComboEventRule, bool)`

GetComposingRulesOk returns a tuple with the ComposingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposingRules

`func (o *ComboEventRuleBranch) SetComposingRules(v []ComboEventRule)`

SetComposingRules sets ComposingRules field to given value.


### GetRuleBranches

`func (o *ComboEventRuleBranch) GetRuleBranches() []ComboEventRuleBranch`

GetRuleBranches returns the RuleBranches field if non-nil, zero value otherwise.

### GetRuleBranchesOk

`func (o *ComboEventRuleBranch) GetRuleBranchesOk() (*[]ComboEventRuleBranch, bool)`

GetRuleBranchesOk returns a tuple with the RuleBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleBranches

`func (o *ComboEventRuleBranch) SetRuleBranches(v []ComboEventRuleBranch)`

SetRuleBranches sets RuleBranches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



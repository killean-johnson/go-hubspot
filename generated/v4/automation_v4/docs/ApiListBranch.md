# ApiListBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | Pointer to **string** | The name of this branch | [optional] 
**Connection** | Pointer to [**ApiConnection**](ApiConnection.md) |  | [optional] 
**FilterBranch** | Pointer to [**ApiListBranchFilterBranch**](ApiListBranchFilterBranch.md) |  | [optional] 

## Methods

### NewApiListBranch

`func NewApiListBranch() *ApiListBranch`

NewApiListBranch instantiates a new ApiListBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiListBranchWithDefaults

`func NewApiListBranchWithDefaults() *ApiListBranch`

NewApiListBranchWithDefaults instantiates a new ApiListBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *ApiListBranch) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *ApiListBranch) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *ApiListBranch) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *ApiListBranch) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetConnection

`func (o *ApiListBranch) GetConnection() ApiConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ApiListBranch) GetConnectionOk() (*ApiConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ApiListBranch) SetConnection(v ApiConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ApiListBranch) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetFilterBranch

`func (o *ApiListBranch) GetFilterBranch() ApiListBranchFilterBranch`

GetFilterBranch returns the FilterBranch field if non-nil, zero value otherwise.

### GetFilterBranchOk

`func (o *ApiListBranch) GetFilterBranchOk() (*ApiListBranchFilterBranch, bool)`

GetFilterBranchOk returns a tuple with the FilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranch

`func (o *ApiListBranch) SetFilterBranch(v ApiListBranchFilterBranch)`

SetFilterBranch sets FilterBranch field to given value.

### HasFilterBranch

`func (o *ApiListBranch) HasFilterBranch() bool`

HasFilterBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



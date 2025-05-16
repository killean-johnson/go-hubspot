# ApiStaticBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchValue** | **string** | If value to check for. If the value of the &#x60;inputValue&#x60; matches this &#x60;branchValue&#x60; than this &#x60;connection&#x60; will get traversed. | 
**Connection** | Pointer to [**ApiConnection**](ApiConnection.md) |  | [optional] 

## Methods

### NewApiStaticBranch

`func NewApiStaticBranch(branchValue string, ) *ApiStaticBranch`

NewApiStaticBranch instantiates a new ApiStaticBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStaticBranchWithDefaults

`func NewApiStaticBranchWithDefaults() *ApiStaticBranch`

NewApiStaticBranchWithDefaults instantiates a new ApiStaticBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchValue

`func (o *ApiStaticBranch) GetBranchValue() string`

GetBranchValue returns the BranchValue field if non-nil, zero value otherwise.

### GetBranchValueOk

`func (o *ApiStaticBranch) GetBranchValueOk() (*string, bool)`

GetBranchValueOk returns a tuple with the BranchValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchValue

`func (o *ApiStaticBranch) SetBranchValue(v string)`

SetBranchValue sets BranchValue field to given value.


### GetConnection

`func (o *ApiStaticBranch) GetConnection() ApiConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ApiStaticBranch) GetConnectionOk() (*ApiConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ApiStaticBranch) SetConnection(v ApiConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ApiStaticBranch) HasConnection() bool`

HasConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



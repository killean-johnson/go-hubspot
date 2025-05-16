# FlagPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverrideState** | Pointer to **string** |  | [optional] 
**DefaultState** | **string** |  | 

## Methods

### NewFlagPutRequest

`func NewFlagPutRequest(defaultState string, ) *FlagPutRequest`

NewFlagPutRequest instantiates a new FlagPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagPutRequestWithDefaults

`func NewFlagPutRequestWithDefaults() *FlagPutRequest`

NewFlagPutRequestWithDefaults instantiates a new FlagPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrideState

`func (o *FlagPutRequest) GetOverrideState() string`

GetOverrideState returns the OverrideState field if non-nil, zero value otherwise.

### GetOverrideStateOk

`func (o *FlagPutRequest) GetOverrideStateOk() (*string, bool)`

GetOverrideStateOk returns a tuple with the OverrideState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideState

`func (o *FlagPutRequest) SetOverrideState(v string)`

SetOverrideState sets OverrideState field to given value.

### HasOverrideState

`func (o *FlagPutRequest) HasOverrideState() bool`

HasOverrideState returns a boolean if a field has been set.

### GetDefaultState

`func (o *FlagPutRequest) GetDefaultState() string`

GetDefaultState returns the DefaultState field if non-nil, zero value otherwise.

### GetDefaultStateOk

`func (o *FlagPutRequest) GetDefaultStateOk() (*string, bool)`

GetDefaultStateOk returns a tuple with the DefaultState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultState

`func (o *FlagPutRequest) SetDefaultState(v string)`

SetDefaultState sets DefaultState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



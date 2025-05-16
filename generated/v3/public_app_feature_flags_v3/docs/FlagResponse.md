# FlagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverrideState** | Pointer to **string** |  | [optional] 
**DefaultState** | **string** |  | 
**AppId** | **int32** |  | 
**FlagName** | **string** |  | 

## Methods

### NewFlagResponse

`func NewFlagResponse(defaultState string, appId int32, flagName string, ) *FlagResponse`

NewFlagResponse instantiates a new FlagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagResponseWithDefaults

`func NewFlagResponseWithDefaults() *FlagResponse`

NewFlagResponseWithDefaults instantiates a new FlagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrideState

`func (o *FlagResponse) GetOverrideState() string`

GetOverrideState returns the OverrideState field if non-nil, zero value otherwise.

### GetOverrideStateOk

`func (o *FlagResponse) GetOverrideStateOk() (*string, bool)`

GetOverrideStateOk returns a tuple with the OverrideState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideState

`func (o *FlagResponse) SetOverrideState(v string)`

SetOverrideState sets OverrideState field to given value.

### HasOverrideState

`func (o *FlagResponse) HasOverrideState() bool`

HasOverrideState returns a boolean if a field has been set.

### GetDefaultState

`func (o *FlagResponse) GetDefaultState() string`

GetDefaultState returns the DefaultState field if non-nil, zero value otherwise.

### GetDefaultStateOk

`func (o *FlagResponse) GetDefaultStateOk() (*string, bool)`

GetDefaultStateOk returns a tuple with the DefaultState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultState

`func (o *FlagResponse) SetDefaultState(v string)`

SetDefaultState sets DefaultState field to given value.


### GetAppId

`func (o *FlagResponse) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *FlagResponse) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *FlagResponse) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetFlagName

`func (o *FlagResponse) GetFlagName() string`

GetFlagName returns the FlagName field if non-nil, zero value otherwise.

### GetFlagNameOk

`func (o *FlagResponse) GetFlagNameOk() (*string, bool)`

GetFlagNameOk returns a tuple with the FlagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagName

`func (o *FlagResponse) SetFlagName(v string)`

SetFlagName sets FlagName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



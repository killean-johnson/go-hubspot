# PublicSequenceStepResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**DelayMillis** | **int32** |  | 
**EmailPattern** | Pointer to [**PublicEmailPatternResponse**](PublicEmailPatternResponse.md) |  | [optional] 
**StepOrder** | **int32** |  | 
**TaskPattern** | Pointer to [**PublicTaskPatternResponse**](PublicTaskPatternResponse.md) |  | [optional] 
**Id** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPublicSequenceStepResponse

`func NewPublicSequenceStepResponse(actionType string, createdAt time.Time, delayMillis int32, stepOrder int32, id string, updatedAt time.Time, ) *PublicSequenceStepResponse`

NewPublicSequenceStepResponse instantiates a new PublicSequenceStepResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSequenceStepResponseWithDefaults

`func NewPublicSequenceStepResponseWithDefaults() *PublicSequenceStepResponse`

NewPublicSequenceStepResponseWithDefaults instantiates a new PublicSequenceStepResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *PublicSequenceStepResponse) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PublicSequenceStepResponse) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PublicSequenceStepResponse) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetCreatedAt

`func (o *PublicSequenceStepResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicSequenceStepResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicSequenceStepResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDelayMillis

`func (o *PublicSequenceStepResponse) GetDelayMillis() int32`

GetDelayMillis returns the DelayMillis field if non-nil, zero value otherwise.

### GetDelayMillisOk

`func (o *PublicSequenceStepResponse) GetDelayMillisOk() (*int32, bool)`

GetDelayMillisOk returns a tuple with the DelayMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMillis

`func (o *PublicSequenceStepResponse) SetDelayMillis(v int32)`

SetDelayMillis sets DelayMillis field to given value.


### GetEmailPattern

`func (o *PublicSequenceStepResponse) GetEmailPattern() PublicEmailPatternResponse`

GetEmailPattern returns the EmailPattern field if non-nil, zero value otherwise.

### GetEmailPatternOk

`func (o *PublicSequenceStepResponse) GetEmailPatternOk() (*PublicEmailPatternResponse, bool)`

GetEmailPatternOk returns a tuple with the EmailPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPattern

`func (o *PublicSequenceStepResponse) SetEmailPattern(v PublicEmailPatternResponse)`

SetEmailPattern sets EmailPattern field to given value.

### HasEmailPattern

`func (o *PublicSequenceStepResponse) HasEmailPattern() bool`

HasEmailPattern returns a boolean if a field has been set.

### GetStepOrder

`func (o *PublicSequenceStepResponse) GetStepOrder() int32`

GetStepOrder returns the StepOrder field if non-nil, zero value otherwise.

### GetStepOrderOk

`func (o *PublicSequenceStepResponse) GetStepOrderOk() (*int32, bool)`

GetStepOrderOk returns a tuple with the StepOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepOrder

`func (o *PublicSequenceStepResponse) SetStepOrder(v int32)`

SetStepOrder sets StepOrder field to given value.


### GetTaskPattern

`func (o *PublicSequenceStepResponse) GetTaskPattern() PublicTaskPatternResponse`

GetTaskPattern returns the TaskPattern field if non-nil, zero value otherwise.

### GetTaskPatternOk

`func (o *PublicSequenceStepResponse) GetTaskPatternOk() (*PublicTaskPatternResponse, bool)`

GetTaskPatternOk returns a tuple with the TaskPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPattern

`func (o *PublicSequenceStepResponse) SetTaskPattern(v PublicTaskPatternResponse)`

SetTaskPattern sets TaskPattern field to given value.

### HasTaskPattern

`func (o *PublicSequenceStepResponse) HasTaskPattern() bool`

HasTaskPattern returns a boolean if a field has been set.

### GetId

`func (o *PublicSequenceStepResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicSequenceStepResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicSequenceStepResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *PublicSequenceStepResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicSequenceStepResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicSequenceStepResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



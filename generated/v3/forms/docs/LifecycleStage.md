# LifecycleStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypeId** | **string** | The objectTypeId for both contact and company | 
**Value** | **string** | The internal name of the contact&#39;s lifecycle stage set when submitting a form | 

## Methods

### NewLifecycleStage

`func NewLifecycleStage(objectTypeId string, value string, ) *LifecycleStage`

NewLifecycleStage instantiates a new LifecycleStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleStageWithDefaults

`func NewLifecycleStageWithDefaults() *LifecycleStage`

NewLifecycleStageWithDefaults instantiates a new LifecycleStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *LifecycleStage) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *LifecycleStage) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *LifecycleStage) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetValue

`func (o *LifecycleStage) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LifecycleStage) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LifecycleStage) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



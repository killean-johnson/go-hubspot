# PublicApiUserActionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActingUser** | [**ActingUser**](ActingUser.md) |  | 
**SubCategory** | Pointer to **string** | The subcategory of the activity. | [optional] 
**OccurredAt** | **time.Time** | The time that the action occurred at. | 
**TargetObjectId** | Pointer to **string** | The ID of the impacted object. | [optional] 
**Action** | **string** | The type of action taken. | 
**Id** | **string** | The unique ID of the activity. | 
**Category** | **string** | The category of the activity. | 

## Methods

### NewPublicApiUserActionEvent

`func NewPublicApiUserActionEvent(actingUser ActingUser, occurredAt time.Time, action string, id string, category string, ) *PublicApiUserActionEvent`

NewPublicApiUserActionEvent instantiates a new PublicApiUserActionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicApiUserActionEventWithDefaults

`func NewPublicApiUserActionEventWithDefaults() *PublicApiUserActionEvent`

NewPublicApiUserActionEventWithDefaults instantiates a new PublicApiUserActionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActingUser

`func (o *PublicApiUserActionEvent) GetActingUser() ActingUser`

GetActingUser returns the ActingUser field if non-nil, zero value otherwise.

### GetActingUserOk

`func (o *PublicApiUserActionEvent) GetActingUserOk() (*ActingUser, bool)`

GetActingUserOk returns a tuple with the ActingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUser

`func (o *PublicApiUserActionEvent) SetActingUser(v ActingUser)`

SetActingUser sets ActingUser field to given value.


### GetSubCategory

`func (o *PublicApiUserActionEvent) GetSubCategory() string`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *PublicApiUserActionEvent) GetSubCategoryOk() (*string, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *PublicApiUserActionEvent) SetSubCategory(v string)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *PublicApiUserActionEvent) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetOccurredAt

`func (o *PublicApiUserActionEvent) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *PublicApiUserActionEvent) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *PublicApiUserActionEvent) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetTargetObjectId

`func (o *PublicApiUserActionEvent) GetTargetObjectId() string`

GetTargetObjectId returns the TargetObjectId field if non-nil, zero value otherwise.

### GetTargetObjectIdOk

`func (o *PublicApiUserActionEvent) GetTargetObjectIdOk() (*string, bool)`

GetTargetObjectIdOk returns a tuple with the TargetObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObjectId

`func (o *PublicApiUserActionEvent) SetTargetObjectId(v string)`

SetTargetObjectId sets TargetObjectId field to given value.

### HasTargetObjectId

`func (o *PublicApiUserActionEvent) HasTargetObjectId() bool`

HasTargetObjectId returns a boolean if a field has been set.

### GetAction

`func (o *PublicApiUserActionEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PublicApiUserActionEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PublicApiUserActionEvent) SetAction(v string)`

SetAction sets Action field to given value.


### GetId

`func (o *PublicApiUserActionEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicApiUserActionEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicApiUserActionEvent) SetId(v string)`

SetId sets Id field to given value.


### GetCategory

`func (o *PublicApiUserActionEvent) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PublicApiUserActionEvent) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PublicApiUserActionEvent) SetCategory(v string)`

SetCategory sets Category field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



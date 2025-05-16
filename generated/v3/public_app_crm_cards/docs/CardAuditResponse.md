# CardAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | **string** |  | 
**ObjectTypeId** | **int32** |  | 
**AuthSource** | **string** |  | 
**ChangedAt** | **int64** |  | 
**ApplicationId** | **int32** |  | 
**InitiatingUserId** | **int32** |  | 

## Methods

### NewCardAuditResponse

`func NewCardAuditResponse(actionType string, objectTypeId int32, authSource string, changedAt int64, applicationId int32, initiatingUserId int32, ) *CardAuditResponse`

NewCardAuditResponse instantiates a new CardAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAuditResponseWithDefaults

`func NewCardAuditResponseWithDefaults() *CardAuditResponse`

NewCardAuditResponseWithDefaults instantiates a new CardAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *CardAuditResponse) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *CardAuditResponse) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *CardAuditResponse) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetObjectTypeId

`func (o *CardAuditResponse) GetObjectTypeId() int32`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *CardAuditResponse) GetObjectTypeIdOk() (*int32, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *CardAuditResponse) SetObjectTypeId(v int32)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetAuthSource

`func (o *CardAuditResponse) GetAuthSource() string`

GetAuthSource returns the AuthSource field if non-nil, zero value otherwise.

### GetAuthSourceOk

`func (o *CardAuditResponse) GetAuthSourceOk() (*string, bool)`

GetAuthSourceOk returns a tuple with the AuthSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSource

`func (o *CardAuditResponse) SetAuthSource(v string)`

SetAuthSource sets AuthSource field to given value.


### GetChangedAt

`func (o *CardAuditResponse) GetChangedAt() int64`

GetChangedAt returns the ChangedAt field if non-nil, zero value otherwise.

### GetChangedAtOk

`func (o *CardAuditResponse) GetChangedAtOk() (*int64, bool)`

GetChangedAtOk returns a tuple with the ChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAt

`func (o *CardAuditResponse) SetChangedAt(v int64)`

SetChangedAt sets ChangedAt field to given value.


### GetApplicationId

`func (o *CardAuditResponse) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CardAuditResponse) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CardAuditResponse) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.


### GetInitiatingUserId

`func (o *CardAuditResponse) GetInitiatingUserId() int32`

GetInitiatingUserId returns the InitiatingUserId field if non-nil, zero value otherwise.

### GetInitiatingUserIdOk

`func (o *CardAuditResponse) GetInitiatingUserIdOk() (*int32, bool)`

GetInitiatingUserIdOk returns a tuple with the InitiatingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatingUserId

`func (o *CardAuditResponse) SetInitiatingUserId(v int32)`

SetInitiatingUserId sets InitiatingUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



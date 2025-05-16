# PublicCardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Fetch** | [**PublicCardFetchBody**](PublicCardFetchBody.md) |  | 
**Display** | [**CardDisplayBody**](CardDisplayBody.md) |  | 
**Id** | **string** |  | 
**Title** | **string** |  | 
**Actions** | [**CardActions**](CardActions.md) |  | 
**AuditHistory** | [**[]CardAuditResponse**](CardAuditResponse.md) |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPublicCardResponse

`func NewPublicCardResponse(fetch PublicCardFetchBody, display CardDisplayBody, id string, title string, actions CardActions, auditHistory []CardAuditResponse, ) *PublicCardResponse`

NewPublicCardResponse instantiates a new PublicCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCardResponseWithDefaults

`func NewPublicCardResponseWithDefaults() *PublicCardResponse`

NewPublicCardResponseWithDefaults instantiates a new PublicCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PublicCardResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicCardResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicCardResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PublicCardResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFetch

`func (o *PublicCardResponse) GetFetch() PublicCardFetchBody`

GetFetch returns the Fetch field if non-nil, zero value otherwise.

### GetFetchOk

`func (o *PublicCardResponse) GetFetchOk() (*PublicCardFetchBody, bool)`

GetFetchOk returns a tuple with the Fetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetch

`func (o *PublicCardResponse) SetFetch(v PublicCardFetchBody)`

SetFetch sets Fetch field to given value.


### GetDisplay

`func (o *PublicCardResponse) GetDisplay() CardDisplayBody`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PublicCardResponse) GetDisplayOk() (*CardDisplayBody, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PublicCardResponse) SetDisplay(v CardDisplayBody)`

SetDisplay sets Display field to given value.


### GetId

`func (o *PublicCardResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicCardResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicCardResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *PublicCardResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PublicCardResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PublicCardResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetActions

`func (o *PublicCardResponse) GetActions() CardActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PublicCardResponse) GetActionsOk() (*CardActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PublicCardResponse) SetActions(v CardActions)`

SetActions sets Actions field to given value.


### GetAuditHistory

`func (o *PublicCardResponse) GetAuditHistory() []CardAuditResponse`

GetAuditHistory returns the AuditHistory field if non-nil, zero value otherwise.

### GetAuditHistoryOk

`func (o *PublicCardResponse) GetAuditHistoryOk() (*[]CardAuditResponse, bool)`

GetAuditHistoryOk returns a tuple with the AuditHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditHistory

`func (o *PublicCardResponse) SetAuditHistory(v []CardAuditResponse)`

SetAuditHistory sets AuditHistory field to given value.


### GetUpdatedAt

`func (o *PublicCardResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicCardResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicCardResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PublicCardResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



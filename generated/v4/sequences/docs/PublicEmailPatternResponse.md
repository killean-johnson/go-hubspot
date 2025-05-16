# PublicEmailPatternResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**TemplateId** | **string** |  | 
**ThreadEmailToStepOrder** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPublicEmailPatternResponse

`func NewPublicEmailPatternResponse(createdAt time.Time, id string, templateId string, updatedAt time.Time, ) *PublicEmailPatternResponse`

NewPublicEmailPatternResponse instantiates a new PublicEmailPatternResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicEmailPatternResponseWithDefaults

`func NewPublicEmailPatternResponseWithDefaults() *PublicEmailPatternResponse`

NewPublicEmailPatternResponseWithDefaults instantiates a new PublicEmailPatternResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PublicEmailPatternResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicEmailPatternResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicEmailPatternResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *PublicEmailPatternResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicEmailPatternResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicEmailPatternResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *PublicEmailPatternResponse) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PublicEmailPatternResponse) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PublicEmailPatternResponse) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetThreadEmailToStepOrder

`func (o *PublicEmailPatternResponse) GetThreadEmailToStepOrder() int32`

GetThreadEmailToStepOrder returns the ThreadEmailToStepOrder field if non-nil, zero value otherwise.

### GetThreadEmailToStepOrderOk

`func (o *PublicEmailPatternResponse) GetThreadEmailToStepOrderOk() (*int32, bool)`

GetThreadEmailToStepOrderOk returns a tuple with the ThreadEmailToStepOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadEmailToStepOrder

`func (o *PublicEmailPatternResponse) SetThreadEmailToStepOrder(v int32)`

SetThreadEmailToStepOrder sets ThreadEmailToStepOrder field to given value.

### HasThreadEmailToStepOrder

`func (o *PublicEmailPatternResponse) HasThreadEmailToStepOrder() bool`

HasThreadEmailToStepOrder returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PublicEmailPatternResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicEmailPatternResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicEmailPatternResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PublicSequenceStepDependencyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReliesOnStepOrder** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**DependencyType** | **string** |  | 
**ReliesOnSequenceStepId** | **string** |  | 
**Id** | **string** |  | 
**RequiredBySequenceStepId** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**RequiredByStepOrder** | **int32** |  | 

## Methods

### NewPublicSequenceStepDependencyResponse

`func NewPublicSequenceStepDependencyResponse(reliesOnStepOrder int32, createdAt time.Time, dependencyType string, reliesOnSequenceStepId string, id string, requiredBySequenceStepId string, updatedAt time.Time, requiredByStepOrder int32, ) *PublicSequenceStepDependencyResponse`

NewPublicSequenceStepDependencyResponse instantiates a new PublicSequenceStepDependencyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSequenceStepDependencyResponseWithDefaults

`func NewPublicSequenceStepDependencyResponseWithDefaults() *PublicSequenceStepDependencyResponse`

NewPublicSequenceStepDependencyResponseWithDefaults instantiates a new PublicSequenceStepDependencyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReliesOnStepOrder

`func (o *PublicSequenceStepDependencyResponse) GetReliesOnStepOrder() int32`

GetReliesOnStepOrder returns the ReliesOnStepOrder field if non-nil, zero value otherwise.

### GetReliesOnStepOrderOk

`func (o *PublicSequenceStepDependencyResponse) GetReliesOnStepOrderOk() (*int32, bool)`

GetReliesOnStepOrderOk returns a tuple with the ReliesOnStepOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliesOnStepOrder

`func (o *PublicSequenceStepDependencyResponse) SetReliesOnStepOrder(v int32)`

SetReliesOnStepOrder sets ReliesOnStepOrder field to given value.


### GetCreatedAt

`func (o *PublicSequenceStepDependencyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicSequenceStepDependencyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicSequenceStepDependencyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDependencyType

`func (o *PublicSequenceStepDependencyResponse) GetDependencyType() string`

GetDependencyType returns the DependencyType field if non-nil, zero value otherwise.

### GetDependencyTypeOk

`func (o *PublicSequenceStepDependencyResponse) GetDependencyTypeOk() (*string, bool)`

GetDependencyTypeOk returns a tuple with the DependencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyType

`func (o *PublicSequenceStepDependencyResponse) SetDependencyType(v string)`

SetDependencyType sets DependencyType field to given value.


### GetReliesOnSequenceStepId

`func (o *PublicSequenceStepDependencyResponse) GetReliesOnSequenceStepId() string`

GetReliesOnSequenceStepId returns the ReliesOnSequenceStepId field if non-nil, zero value otherwise.

### GetReliesOnSequenceStepIdOk

`func (o *PublicSequenceStepDependencyResponse) GetReliesOnSequenceStepIdOk() (*string, bool)`

GetReliesOnSequenceStepIdOk returns a tuple with the ReliesOnSequenceStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliesOnSequenceStepId

`func (o *PublicSequenceStepDependencyResponse) SetReliesOnSequenceStepId(v string)`

SetReliesOnSequenceStepId sets ReliesOnSequenceStepId field to given value.


### GetId

`func (o *PublicSequenceStepDependencyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicSequenceStepDependencyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicSequenceStepDependencyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetRequiredBySequenceStepId

`func (o *PublicSequenceStepDependencyResponse) GetRequiredBySequenceStepId() string`

GetRequiredBySequenceStepId returns the RequiredBySequenceStepId field if non-nil, zero value otherwise.

### GetRequiredBySequenceStepIdOk

`func (o *PublicSequenceStepDependencyResponse) GetRequiredBySequenceStepIdOk() (*string, bool)`

GetRequiredBySequenceStepIdOk returns a tuple with the RequiredBySequenceStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredBySequenceStepId

`func (o *PublicSequenceStepDependencyResponse) SetRequiredBySequenceStepId(v string)`

SetRequiredBySequenceStepId sets RequiredBySequenceStepId field to given value.


### GetUpdatedAt

`func (o *PublicSequenceStepDependencyResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicSequenceStepDependencyResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicSequenceStepDependencyResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRequiredByStepOrder

`func (o *PublicSequenceStepDependencyResponse) GetRequiredByStepOrder() int32`

GetRequiredByStepOrder returns the RequiredByStepOrder field if non-nil, zero value otherwise.

### GetRequiredByStepOrderOk

`func (o *PublicSequenceStepDependencyResponse) GetRequiredByStepOrderOk() (*int32, bool)`

GetRequiredByStepOrderOk returns a tuple with the RequiredByStepOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredByStepOrder

`func (o *PublicSequenceStepDependencyResponse) SetRequiredByStepOrder(v int32)`

SetRequiredByStepOrder sets RequiredByStepOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



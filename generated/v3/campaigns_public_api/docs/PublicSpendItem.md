# PublicSpendItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int32** |  | 
**Amount** | **float32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Order** | **int32** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewPublicSpendItem

`func NewPublicSpendItem(createdAt int32, amount float32, name string, id string, order int32, updatedAt int32, ) *PublicSpendItem`

NewPublicSpendItem instantiates a new PublicSpendItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSpendItemWithDefaults

`func NewPublicSpendItemWithDefaults() *PublicSpendItem`

NewPublicSpendItemWithDefaults instantiates a new PublicSpendItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PublicSpendItem) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicSpendItem) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicSpendItem) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetAmount

`func (o *PublicSpendItem) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PublicSpendItem) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PublicSpendItem) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetName

`func (o *PublicSpendItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicSpendItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicSpendItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PublicSpendItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicSpendItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicSpendItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublicSpendItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PublicSpendItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicSpendItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicSpendItem) SetId(v string)`

SetId sets Id field to given value.


### GetOrder

`func (o *PublicSpendItem) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PublicSpendItem) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PublicSpendItem) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetUpdatedAt

`func (o *PublicSpendItem) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicSpendItem) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicSpendItem) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



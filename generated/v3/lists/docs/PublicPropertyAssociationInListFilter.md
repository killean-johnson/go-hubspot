# PublicPropertyAssociationInListFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListId** | **string** |  | 
**CoalescingRefineBy** | [**PublicFormSubmissionFilterCoalescingRefineBy**](PublicFormSubmissionFilterCoalescingRefineBy.md) |  | 
**PropertyWithObjectId** | **string** |  | 
**FilterType** | **string** |  | [default to "PROPERTY_ASSOCIATION"]
**ToObjectTypeId** | Pointer to **string** |  | [optional] 
**Operator** | **string** |  | 

## Methods

### NewPublicPropertyAssociationInListFilter

`func NewPublicPropertyAssociationInListFilter(listId string, coalescingRefineBy PublicFormSubmissionFilterCoalescingRefineBy, propertyWithObjectId string, filterType string, operator string, ) *PublicPropertyAssociationInListFilter`

NewPublicPropertyAssociationInListFilter instantiates a new PublicPropertyAssociationInListFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPropertyAssociationInListFilterWithDefaults

`func NewPublicPropertyAssociationInListFilterWithDefaults() *PublicPropertyAssociationInListFilter`

NewPublicPropertyAssociationInListFilterWithDefaults instantiates a new PublicPropertyAssociationInListFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListId

`func (o *PublicPropertyAssociationInListFilter) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *PublicPropertyAssociationInListFilter) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *PublicPropertyAssociationInListFilter) SetListId(v string)`

SetListId sets ListId field to given value.


### GetCoalescingRefineBy

`func (o *PublicPropertyAssociationInListFilter) GetCoalescingRefineBy() PublicFormSubmissionFilterCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *PublicPropertyAssociationInListFilter) GetCoalescingRefineByOk() (*PublicFormSubmissionFilterCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *PublicPropertyAssociationInListFilter) SetCoalescingRefineBy(v PublicFormSubmissionFilterCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.


### GetPropertyWithObjectId

`func (o *PublicPropertyAssociationInListFilter) GetPropertyWithObjectId() string`

GetPropertyWithObjectId returns the PropertyWithObjectId field if non-nil, zero value otherwise.

### GetPropertyWithObjectIdOk

`func (o *PublicPropertyAssociationInListFilter) GetPropertyWithObjectIdOk() (*string, bool)`

GetPropertyWithObjectIdOk returns a tuple with the PropertyWithObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyWithObjectId

`func (o *PublicPropertyAssociationInListFilter) SetPropertyWithObjectId(v string)`

SetPropertyWithObjectId sets PropertyWithObjectId field to given value.


### GetFilterType

`func (o *PublicPropertyAssociationInListFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicPropertyAssociationInListFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicPropertyAssociationInListFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetToObjectTypeId

`func (o *PublicPropertyAssociationInListFilter) GetToObjectTypeId() string`

GetToObjectTypeId returns the ToObjectTypeId field if non-nil, zero value otherwise.

### GetToObjectTypeIdOk

`func (o *PublicPropertyAssociationInListFilter) GetToObjectTypeIdOk() (*string, bool)`

GetToObjectTypeIdOk returns a tuple with the ToObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectTypeId

`func (o *PublicPropertyAssociationInListFilter) SetToObjectTypeId(v string)`

SetToObjectTypeId sets ToObjectTypeId field to given value.

### HasToObjectTypeId

`func (o *PublicPropertyAssociationInListFilter) HasToObjectTypeId() bool`

HasToObjectTypeId returns a boolean if a field has been set.

### GetOperator

`func (o *PublicPropertyAssociationInListFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicPropertyAssociationInListFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicPropertyAssociationInListFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PublicList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessingType** | **string** |  | 
**ObjectTypeId** | **string** |  | 
**UpdatedById** | Pointer to **string** |  | [optional] 
**FiltersUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ListId** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ProcessingStatus** | **string** |  | 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**ListVersion** | **int32** |  | 
**Size** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**CreatedById** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPublicList

`func NewPublicList(processingType string, objectTypeId string, listId string, processingStatus string, listVersion int32, name string, ) *PublicList`

NewPublicList instantiates a new PublicList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicListWithDefaults

`func NewPublicListWithDefaults() *PublicList`

NewPublicListWithDefaults instantiates a new PublicList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessingType

`func (o *PublicList) GetProcessingType() string`

GetProcessingType returns the ProcessingType field if non-nil, zero value otherwise.

### GetProcessingTypeOk

`func (o *PublicList) GetProcessingTypeOk() (*string, bool)`

GetProcessingTypeOk returns a tuple with the ProcessingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingType

`func (o *PublicList) SetProcessingType(v string)`

SetProcessingType sets ProcessingType field to given value.


### GetObjectTypeId

`func (o *PublicList) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *PublicList) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *PublicList) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetUpdatedById

`func (o *PublicList) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *PublicList) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *PublicList) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.

### HasUpdatedById

`func (o *PublicList) HasUpdatedById() bool`

HasUpdatedById returns a boolean if a field has been set.

### GetFiltersUpdatedAt

`func (o *PublicList) GetFiltersUpdatedAt() time.Time`

GetFiltersUpdatedAt returns the FiltersUpdatedAt field if non-nil, zero value otherwise.

### GetFiltersUpdatedAtOk

`func (o *PublicList) GetFiltersUpdatedAtOk() (*time.Time, bool)`

GetFiltersUpdatedAtOk returns a tuple with the FiltersUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersUpdatedAt

`func (o *PublicList) SetFiltersUpdatedAt(v time.Time)`

SetFiltersUpdatedAt sets FiltersUpdatedAt field to given value.

### HasFiltersUpdatedAt

`func (o *PublicList) HasFiltersUpdatedAt() bool`

HasFiltersUpdatedAt returns a boolean if a field has been set.

### GetListId

`func (o *PublicList) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *PublicList) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *PublicList) SetListId(v string)`

SetListId sets ListId field to given value.


### GetCreatedAt

`func (o *PublicList) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicList) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicList) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PublicList) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetProcessingStatus

`func (o *PublicList) GetProcessingStatus() string`

GetProcessingStatus returns the ProcessingStatus field if non-nil, zero value otherwise.

### GetProcessingStatusOk

`func (o *PublicList) GetProcessingStatusOk() (*string, bool)`

GetProcessingStatusOk returns a tuple with the ProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStatus

`func (o *PublicList) SetProcessingStatus(v string)`

SetProcessingStatus sets ProcessingStatus field to given value.


### GetDeletedAt

`func (o *PublicList) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PublicList) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PublicList) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PublicList) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetListVersion

`func (o *PublicList) GetListVersion() int32`

GetListVersion returns the ListVersion field if non-nil, zero value otherwise.

### GetListVersionOk

`func (o *PublicList) GetListVersionOk() (*int32, bool)`

GetListVersionOk returns a tuple with the ListVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListVersion

`func (o *PublicList) SetListVersion(v int32)`

SetListVersion sets ListVersion field to given value.


### GetSize

`func (o *PublicList) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PublicList) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PublicList) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PublicList) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetName

`func (o *PublicList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicList) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedById

`func (o *PublicList) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *PublicList) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *PublicList) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *PublicList) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PublicList) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicList) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicList) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PublicList) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



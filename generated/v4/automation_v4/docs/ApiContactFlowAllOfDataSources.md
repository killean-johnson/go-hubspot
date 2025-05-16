# ApiContactFlowAllOfDataSources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypeId** | **string** |  | 
**Name** | **string** |  | 
**AssociationTypeId** | **int32** |  | 
**AssociationCategory** | **string** |  | 
**SortBy** | Pointer to [**ApiSort**](ApiSort.md) |  | [optional] 
**Type** | **string** |  | [default to "ASSOCIATION"]

## Methods

### NewApiContactFlowAllOfDataSources

`func NewApiContactFlowAllOfDataSources(objectTypeId string, name string, associationTypeId int32, associationCategory string, type_ string, ) *ApiContactFlowAllOfDataSources`

NewApiContactFlowAllOfDataSources instantiates a new ApiContactFlowAllOfDataSources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiContactFlowAllOfDataSourcesWithDefaults

`func NewApiContactFlowAllOfDataSourcesWithDefaults() *ApiContactFlowAllOfDataSources`

NewApiContactFlowAllOfDataSourcesWithDefaults instantiates a new ApiContactFlowAllOfDataSources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *ApiContactFlowAllOfDataSources) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ApiContactFlowAllOfDataSources) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ApiContactFlowAllOfDataSources) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetName

`func (o *ApiContactFlowAllOfDataSources) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiContactFlowAllOfDataSources) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiContactFlowAllOfDataSources) SetName(v string)`

SetName sets Name field to given value.


### GetAssociationTypeId

`func (o *ApiContactFlowAllOfDataSources) GetAssociationTypeId() int32`

GetAssociationTypeId returns the AssociationTypeId field if non-nil, zero value otherwise.

### GetAssociationTypeIdOk

`func (o *ApiContactFlowAllOfDataSources) GetAssociationTypeIdOk() (*int32, bool)`

GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypeId

`func (o *ApiContactFlowAllOfDataSources) SetAssociationTypeId(v int32)`

SetAssociationTypeId sets AssociationTypeId field to given value.


### GetAssociationCategory

`func (o *ApiContactFlowAllOfDataSources) GetAssociationCategory() string`

GetAssociationCategory returns the AssociationCategory field if non-nil, zero value otherwise.

### GetAssociationCategoryOk

`func (o *ApiContactFlowAllOfDataSources) GetAssociationCategoryOk() (*string, bool)`

GetAssociationCategoryOk returns a tuple with the AssociationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCategory

`func (o *ApiContactFlowAllOfDataSources) SetAssociationCategory(v string)`

SetAssociationCategory sets AssociationCategory field to given value.


### GetSortBy

`func (o *ApiContactFlowAllOfDataSources) GetSortBy() ApiSort`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ApiContactFlowAllOfDataSources) GetSortByOk() (*ApiSort, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ApiContactFlowAllOfDataSources) SetSortBy(v ApiSort)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ApiContactFlowAllOfDataSources) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetType

`func (o *ApiContactFlowAllOfDataSources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiContactFlowAllOfDataSources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiContactFlowAllOfDataSources) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



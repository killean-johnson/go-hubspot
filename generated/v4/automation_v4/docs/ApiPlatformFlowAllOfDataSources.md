# ApiPlatformFlowAllOfDataSources

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

### NewApiPlatformFlowAllOfDataSources

`func NewApiPlatformFlowAllOfDataSources(objectTypeId string, name string, associationTypeId int32, associationCategory string, type_ string, ) *ApiPlatformFlowAllOfDataSources`

NewApiPlatformFlowAllOfDataSources instantiates a new ApiPlatformFlowAllOfDataSources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPlatformFlowAllOfDataSourcesWithDefaults

`func NewApiPlatformFlowAllOfDataSourcesWithDefaults() *ApiPlatformFlowAllOfDataSources`

NewApiPlatformFlowAllOfDataSourcesWithDefaults instantiates a new ApiPlatformFlowAllOfDataSources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *ApiPlatformFlowAllOfDataSources) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ApiPlatformFlowAllOfDataSources) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ApiPlatformFlowAllOfDataSources) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetName

`func (o *ApiPlatformFlowAllOfDataSources) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPlatformFlowAllOfDataSources) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPlatformFlowAllOfDataSources) SetName(v string)`

SetName sets Name field to given value.


### GetAssociationTypeId

`func (o *ApiPlatformFlowAllOfDataSources) GetAssociationTypeId() int32`

GetAssociationTypeId returns the AssociationTypeId field if non-nil, zero value otherwise.

### GetAssociationTypeIdOk

`func (o *ApiPlatformFlowAllOfDataSources) GetAssociationTypeIdOk() (*int32, bool)`

GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypeId

`func (o *ApiPlatformFlowAllOfDataSources) SetAssociationTypeId(v int32)`

SetAssociationTypeId sets AssociationTypeId field to given value.


### GetAssociationCategory

`func (o *ApiPlatformFlowAllOfDataSources) GetAssociationCategory() string`

GetAssociationCategory returns the AssociationCategory field if non-nil, zero value otherwise.

### GetAssociationCategoryOk

`func (o *ApiPlatformFlowAllOfDataSources) GetAssociationCategoryOk() (*string, bool)`

GetAssociationCategoryOk returns a tuple with the AssociationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCategory

`func (o *ApiPlatformFlowAllOfDataSources) SetAssociationCategory(v string)`

SetAssociationCategory sets AssociationCategory field to given value.


### GetSortBy

`func (o *ApiPlatformFlowAllOfDataSources) GetSortBy() ApiSort`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ApiPlatformFlowAllOfDataSources) GetSortByOk() (*ApiSort, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ApiPlatformFlowAllOfDataSources) SetSortBy(v ApiSort)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ApiPlatformFlowAllOfDataSources) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetType

`func (o *ApiPlatformFlowAllOfDataSources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiPlatformFlowAllOfDataSources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiPlatformFlowAllOfDataSources) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



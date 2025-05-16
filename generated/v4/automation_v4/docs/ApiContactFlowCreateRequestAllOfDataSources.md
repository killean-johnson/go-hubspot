# ApiContactFlowCreateRequestAllOfDataSources

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

### NewApiContactFlowCreateRequestAllOfDataSources

`func NewApiContactFlowCreateRequestAllOfDataSources(objectTypeId string, name string, associationTypeId int32, associationCategory string, type_ string, ) *ApiContactFlowCreateRequestAllOfDataSources`

NewApiContactFlowCreateRequestAllOfDataSources instantiates a new ApiContactFlowCreateRequestAllOfDataSources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiContactFlowCreateRequestAllOfDataSourcesWithDefaults

`func NewApiContactFlowCreateRequestAllOfDataSourcesWithDefaults() *ApiContactFlowCreateRequestAllOfDataSources`

NewApiContactFlowCreateRequestAllOfDataSourcesWithDefaults instantiates a new ApiContactFlowCreateRequestAllOfDataSources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *ApiContactFlowCreateRequestAllOfDataSources) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ApiContactFlowCreateRequestAllOfDataSources) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ApiContactFlowCreateRequestAllOfDataSources) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetName

`func (o *ApiContactFlowCreateRequestAllOfDataSources) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiContactFlowCreateRequestAllOfDataSources) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiContactFlowCreateRequestAllOfDataSources) SetName(v string)`

SetName sets Name field to given value.


### GetAssociationTypeId

`func (o *ApiContactFlowCreateRequestAllOfDataSources) GetAssociationTypeId() int32`

GetAssociationTypeId returns the AssociationTypeId field if non-nil, zero value otherwise.

### GetAssociationTypeIdOk

`func (o *ApiContactFlowCreateRequestAllOfDataSources) GetAssociationTypeIdOk() (*int32, bool)`

GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypeId

`func (o *ApiContactFlowCreateRequestAllOfDataSources) SetAssociationTypeId(v int32)`

SetAssociationTypeId sets AssociationTypeId field to given value.


### GetAssociationCategory

`func (o *ApiContactFlowCreateRequestAllOfDataSources) GetAssociationCategory() string`

GetAssociationCategory returns the AssociationCategory field if non-nil, zero value otherwise.

### GetAssociationCategoryOk

`func (o *ApiContactFlowCreateRequestAllOfDataSources) GetAssociationCategoryOk() (*string, bool)`

GetAssociationCategoryOk returns a tuple with the AssociationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCategory

`func (o *ApiContactFlowCreateRequestAllOfDataSources) SetAssociationCategory(v string)`

SetAssociationCategory sets AssociationCategory field to given value.


### GetSortBy

`func (o *ApiContactFlowCreateRequestAllOfDataSources) GetSortBy() ApiSort`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ApiContactFlowCreateRequestAllOfDataSources) GetSortByOk() (*ApiSort, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ApiContactFlowCreateRequestAllOfDataSources) SetSortBy(v ApiSort)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ApiContactFlowCreateRequestAllOfDataSources) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetType

`func (o *ApiContactFlowCreateRequestAllOfDataSources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiContactFlowCreateRequestAllOfDataSources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiContactFlowCreateRequestAllOfDataSources) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



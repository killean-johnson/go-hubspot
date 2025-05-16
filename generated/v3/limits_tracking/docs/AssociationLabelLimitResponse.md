# AssociationLabelLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllLabels** | **[]string** |  | 
**FromObjectType** | [**ObjectTypeDefinition**](ObjectTypeDefinition.md) |  | 
**ToObjectType** | [**ObjectTypeDefinition**](ObjectTypeDefinition.md) |  | 
**Usage** | **int32** |  | 
**Percentage** | **float32** |  | 
**Limit** | **int32** |  | 

## Methods

### NewAssociationLabelLimitResponse

`func NewAssociationLabelLimitResponse(allLabels []string, fromObjectType ObjectTypeDefinition, toObjectType ObjectTypeDefinition, usage int32, percentage float32, limit int32, ) *AssociationLabelLimitResponse`

NewAssociationLabelLimitResponse instantiates a new AssociationLabelLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationLabelLimitResponseWithDefaults

`func NewAssociationLabelLimitResponseWithDefaults() *AssociationLabelLimitResponse`

NewAssociationLabelLimitResponseWithDefaults instantiates a new AssociationLabelLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllLabels

`func (o *AssociationLabelLimitResponse) GetAllLabels() []string`

GetAllLabels returns the AllLabels field if non-nil, zero value otherwise.

### GetAllLabelsOk

`func (o *AssociationLabelLimitResponse) GetAllLabelsOk() (*[]string, bool)`

GetAllLabelsOk returns a tuple with the AllLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLabels

`func (o *AssociationLabelLimitResponse) SetAllLabels(v []string)`

SetAllLabels sets AllLabels field to given value.


### GetFromObjectType

`func (o *AssociationLabelLimitResponse) GetFromObjectType() ObjectTypeDefinition`

GetFromObjectType returns the FromObjectType field if non-nil, zero value otherwise.

### GetFromObjectTypeOk

`func (o *AssociationLabelLimitResponse) GetFromObjectTypeOk() (*ObjectTypeDefinition, bool)`

GetFromObjectTypeOk returns a tuple with the FromObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromObjectType

`func (o *AssociationLabelLimitResponse) SetFromObjectType(v ObjectTypeDefinition)`

SetFromObjectType sets FromObjectType field to given value.


### GetToObjectType

`func (o *AssociationLabelLimitResponse) GetToObjectType() ObjectTypeDefinition`

GetToObjectType returns the ToObjectType field if non-nil, zero value otherwise.

### GetToObjectTypeOk

`func (o *AssociationLabelLimitResponse) GetToObjectTypeOk() (*ObjectTypeDefinition, bool)`

GetToObjectTypeOk returns a tuple with the ToObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectType

`func (o *AssociationLabelLimitResponse) SetToObjectType(v ObjectTypeDefinition)`

SetToObjectType sets ToObjectType field to given value.


### GetUsage

`func (o *AssociationLabelLimitResponse) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AssociationLabelLimitResponse) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AssociationLabelLimitResponse) SetUsage(v int32)`

SetUsage sets Usage field to given value.


### GetPercentage

`func (o *AssociationLabelLimitResponse) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *AssociationLabelLimitResponse) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *AssociationLabelLimitResponse) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.


### GetLimit

`func (o *AssociationLabelLimitResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AssociationLabelLimitResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AssociationLabelLimitResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



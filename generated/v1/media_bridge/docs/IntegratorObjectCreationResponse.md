# IntegratorObjectCreationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | [**[]PropertyDefinition**](PropertyDefinition.md) |  | 
**ObjectType** | [**InboundDbObjectType**](InboundDbObjectType.md) |  | 
**PropertyGroups** | [**[]Group**](Group.md) |  | 

## Methods

### NewIntegratorObjectCreationResponse

`func NewIntegratorObjectCreationResponse(properties []PropertyDefinition, objectType InboundDbObjectType, propertyGroups []Group, ) *IntegratorObjectCreationResponse`

NewIntegratorObjectCreationResponse instantiates a new IntegratorObjectCreationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegratorObjectCreationResponseWithDefaults

`func NewIntegratorObjectCreationResponseWithDefaults() *IntegratorObjectCreationResponse`

NewIntegratorObjectCreationResponseWithDefaults instantiates a new IntegratorObjectCreationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *IntegratorObjectCreationResponse) GetProperties() []PropertyDefinition`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IntegratorObjectCreationResponse) GetPropertiesOk() (*[]PropertyDefinition, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IntegratorObjectCreationResponse) SetProperties(v []PropertyDefinition)`

SetProperties sets Properties field to given value.


### GetObjectType

`func (o *IntegratorObjectCreationResponse) GetObjectType() InboundDbObjectType`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IntegratorObjectCreationResponse) GetObjectTypeOk() (*InboundDbObjectType, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IntegratorObjectCreationResponse) SetObjectType(v InboundDbObjectType)`

SetObjectType sets ObjectType field to given value.


### GetPropertyGroups

`func (o *IntegratorObjectCreationResponse) GetPropertyGroups() []Group`

GetPropertyGroups returns the PropertyGroups field if non-nil, zero value otherwise.

### GetPropertyGroupsOk

`func (o *IntegratorObjectCreationResponse) GetPropertyGroupsOk() (*[]Group, bool)`

GetPropertyGroupsOk returns a tuple with the PropertyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyGroups

`func (o *IntegratorObjectCreationResponse) SetPropertyGroups(v []Group)`

SetPropertyGroups sets PropertyGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



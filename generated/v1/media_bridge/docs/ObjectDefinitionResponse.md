# ObjectDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to [**InboundDbObjectType**](InboundDbObjectType.md) |  | [optional] 
**ObjectTypeId** | **string** |  | 
**ObjectTypeName** | **string** |  | 
**Properties** | [**[]PropertyDefinition**](PropertyDefinition.md) |  | 
**PropertyGroups** | [**[]GroupView**](GroupView.md) |  | 

## Methods

### NewObjectDefinitionResponse

`func NewObjectDefinitionResponse(objectTypeId string, objectTypeName string, properties []PropertyDefinition, propertyGroups []GroupView, ) *ObjectDefinitionResponse`

NewObjectDefinitionResponse instantiates a new ObjectDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectDefinitionResponseWithDefaults

`func NewObjectDefinitionResponseWithDefaults() *ObjectDefinitionResponse`

NewObjectDefinitionResponseWithDefaults instantiates a new ObjectDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *ObjectDefinitionResponse) GetSchema() InboundDbObjectType`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ObjectDefinitionResponse) GetSchemaOk() (*InboundDbObjectType, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ObjectDefinitionResponse) SetSchema(v InboundDbObjectType)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ObjectDefinitionResponse) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetObjectTypeId

`func (o *ObjectDefinitionResponse) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ObjectDefinitionResponse) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ObjectDefinitionResponse) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetObjectTypeName

`func (o *ObjectDefinitionResponse) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *ObjectDefinitionResponse) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *ObjectDefinitionResponse) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.


### GetProperties

`func (o *ObjectDefinitionResponse) GetProperties() []PropertyDefinition`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ObjectDefinitionResponse) GetPropertiesOk() (*[]PropertyDefinition, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ObjectDefinitionResponse) SetProperties(v []PropertyDefinition)`

SetProperties sets Properties field to given value.


### GetPropertyGroups

`func (o *ObjectDefinitionResponse) GetPropertyGroups() []GroupView`

GetPropertyGroups returns the PropertyGroups field if non-nil, zero value otherwise.

### GetPropertyGroupsOk

`func (o *ObjectDefinitionResponse) GetPropertyGroupsOk() (*[]GroupView, bool)`

GetPropertyGroupsOk returns a tuple with the PropertyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyGroups

`func (o *ObjectDefinitionResponse) SetPropertyGroups(v []GroupView)`

SetPropertyGroups sets PropertyGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



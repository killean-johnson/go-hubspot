# PublicExportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportType** | **string** |  | [default to "LIST"]
**Format** | **string** |  | 
**ExportName** | **string** |  | 
**ObjectProperties** | **[]string** |  | 
**AssociatedObjectType** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | 
**Language** | **string** |  | 
**ExportInternalValuesOptions** | **[]string** |  | 
**OverrideAssociatedObjectsPerDefinitionPerRowLimit** | **bool** |  | 
**ListId** | **string** |  | 
**PublicCrmSearchRequest** | Pointer to [**PublicCrmSearchRequest**](PublicCrmSearchRequest.md) |  | [optional] 

## Methods

### NewPublicExportRequest

`func NewPublicExportRequest(exportType string, format string, exportName string, objectProperties []string, objectType string, language string, exportInternalValuesOptions []string, overrideAssociatedObjectsPerDefinitionPerRowLimit bool, listId string, ) *PublicExportRequest`

NewPublicExportRequest instantiates a new PublicExportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicExportRequestWithDefaults

`func NewPublicExportRequestWithDefaults() *PublicExportRequest`

NewPublicExportRequestWithDefaults instantiates a new PublicExportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportType

`func (o *PublicExportRequest) GetExportType() string`

GetExportType returns the ExportType field if non-nil, zero value otherwise.

### GetExportTypeOk

`func (o *PublicExportRequest) GetExportTypeOk() (*string, bool)`

GetExportTypeOk returns a tuple with the ExportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportType

`func (o *PublicExportRequest) SetExportType(v string)`

SetExportType sets ExportType field to given value.


### GetFormat

`func (o *PublicExportRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PublicExportRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PublicExportRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetExportName

`func (o *PublicExportRequest) GetExportName() string`

GetExportName returns the ExportName field if non-nil, zero value otherwise.

### GetExportNameOk

`func (o *PublicExportRequest) GetExportNameOk() (*string, bool)`

GetExportNameOk returns a tuple with the ExportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportName

`func (o *PublicExportRequest) SetExportName(v string)`

SetExportName sets ExportName field to given value.


### GetObjectProperties

`func (o *PublicExportRequest) GetObjectProperties() []string`

GetObjectProperties returns the ObjectProperties field if non-nil, zero value otherwise.

### GetObjectPropertiesOk

`func (o *PublicExportRequest) GetObjectPropertiesOk() (*[]string, bool)`

GetObjectPropertiesOk returns a tuple with the ObjectProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProperties

`func (o *PublicExportRequest) SetObjectProperties(v []string)`

SetObjectProperties sets ObjectProperties field to given value.


### GetAssociatedObjectType

`func (o *PublicExportRequest) GetAssociatedObjectType() string`

GetAssociatedObjectType returns the AssociatedObjectType field if non-nil, zero value otherwise.

### GetAssociatedObjectTypeOk

`func (o *PublicExportRequest) GetAssociatedObjectTypeOk() (*string, bool)`

GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectType

`func (o *PublicExportRequest) SetAssociatedObjectType(v string)`

SetAssociatedObjectType sets AssociatedObjectType field to given value.

### HasAssociatedObjectType

`func (o *PublicExportRequest) HasAssociatedObjectType() bool`

HasAssociatedObjectType returns a boolean if a field has been set.

### GetObjectType

`func (o *PublicExportRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PublicExportRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PublicExportRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLanguage

`func (o *PublicExportRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PublicExportRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PublicExportRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetExportInternalValuesOptions

`func (o *PublicExportRequest) GetExportInternalValuesOptions() []string`

GetExportInternalValuesOptions returns the ExportInternalValuesOptions field if non-nil, zero value otherwise.

### GetExportInternalValuesOptionsOk

`func (o *PublicExportRequest) GetExportInternalValuesOptionsOk() (*[]string, bool)`

GetExportInternalValuesOptionsOk returns a tuple with the ExportInternalValuesOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportInternalValuesOptions

`func (o *PublicExportRequest) SetExportInternalValuesOptions(v []string)`

SetExportInternalValuesOptions sets ExportInternalValuesOptions field to given value.


### GetOverrideAssociatedObjectsPerDefinitionPerRowLimit

`func (o *PublicExportRequest) GetOverrideAssociatedObjectsPerDefinitionPerRowLimit() bool`

GetOverrideAssociatedObjectsPerDefinitionPerRowLimit returns the OverrideAssociatedObjectsPerDefinitionPerRowLimit field if non-nil, zero value otherwise.

### GetOverrideAssociatedObjectsPerDefinitionPerRowLimitOk

`func (o *PublicExportRequest) GetOverrideAssociatedObjectsPerDefinitionPerRowLimitOk() (*bool, bool)`

GetOverrideAssociatedObjectsPerDefinitionPerRowLimitOk returns a tuple with the OverrideAssociatedObjectsPerDefinitionPerRowLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideAssociatedObjectsPerDefinitionPerRowLimit

`func (o *PublicExportRequest) SetOverrideAssociatedObjectsPerDefinitionPerRowLimit(v bool)`

SetOverrideAssociatedObjectsPerDefinitionPerRowLimit sets OverrideAssociatedObjectsPerDefinitionPerRowLimit field to given value.


### GetListId

`func (o *PublicExportRequest) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *PublicExportRequest) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *PublicExportRequest) SetListId(v string)`

SetListId sets ListId field to given value.


### GetPublicCrmSearchRequest

`func (o *PublicExportRequest) GetPublicCrmSearchRequest() PublicCrmSearchRequest`

GetPublicCrmSearchRequest returns the PublicCrmSearchRequest field if non-nil, zero value otherwise.

### GetPublicCrmSearchRequestOk

`func (o *PublicExportRequest) GetPublicCrmSearchRequestOk() (*PublicCrmSearchRequest, bool)`

GetPublicCrmSearchRequestOk returns a tuple with the PublicCrmSearchRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicCrmSearchRequest

`func (o *PublicExportRequest) SetPublicCrmSearchRequest(v PublicCrmSearchRequest)`

SetPublicCrmSearchRequest sets PublicCrmSearchRequest field to given value.

### HasPublicCrmSearchRequest

`func (o *PublicExportRequest) HasPublicCrmSearchRequest() bool`

HasPublicCrmSearchRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



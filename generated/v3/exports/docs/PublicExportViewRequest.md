# PublicExportViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportType** | **string** |  | [default to "VIEW"]
**Format** | **string** |  | 
**ExportName** | **string** |  | 
**ObjectProperties** | **[]string** |  | 
**AssociatedObjectType** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | 
**Language** | **string** |  | 
**ExportInternalValuesOptions** | **[]string** |  | 
**OverrideAssociatedObjectsPerDefinitionPerRowLimit** | **bool** |  | 
**PublicCrmSearchRequest** | Pointer to [**PublicCrmSearchRequest**](PublicCrmSearchRequest.md) |  | [optional] 
**ListId** | **string** |  | 

## Methods

### NewPublicExportViewRequest

`func NewPublicExportViewRequest(exportType string, format string, exportName string, objectProperties []string, objectType string, language string, exportInternalValuesOptions []string, overrideAssociatedObjectsPerDefinitionPerRowLimit bool, listId string, ) *PublicExportViewRequest`

NewPublicExportViewRequest instantiates a new PublicExportViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicExportViewRequestWithDefaults

`func NewPublicExportViewRequestWithDefaults() *PublicExportViewRequest`

NewPublicExportViewRequestWithDefaults instantiates a new PublicExportViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportType

`func (o *PublicExportViewRequest) GetExportType() string`

GetExportType returns the ExportType field if non-nil, zero value otherwise.

### GetExportTypeOk

`func (o *PublicExportViewRequest) GetExportTypeOk() (*string, bool)`

GetExportTypeOk returns a tuple with the ExportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportType

`func (o *PublicExportViewRequest) SetExportType(v string)`

SetExportType sets ExportType field to given value.


### GetFormat

`func (o *PublicExportViewRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PublicExportViewRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PublicExportViewRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetExportName

`func (o *PublicExportViewRequest) GetExportName() string`

GetExportName returns the ExportName field if non-nil, zero value otherwise.

### GetExportNameOk

`func (o *PublicExportViewRequest) GetExportNameOk() (*string, bool)`

GetExportNameOk returns a tuple with the ExportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportName

`func (o *PublicExportViewRequest) SetExportName(v string)`

SetExportName sets ExportName field to given value.


### GetObjectProperties

`func (o *PublicExportViewRequest) GetObjectProperties() []string`

GetObjectProperties returns the ObjectProperties field if non-nil, zero value otherwise.

### GetObjectPropertiesOk

`func (o *PublicExportViewRequest) GetObjectPropertiesOk() (*[]string, bool)`

GetObjectPropertiesOk returns a tuple with the ObjectProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProperties

`func (o *PublicExportViewRequest) SetObjectProperties(v []string)`

SetObjectProperties sets ObjectProperties field to given value.


### GetAssociatedObjectType

`func (o *PublicExportViewRequest) GetAssociatedObjectType() string`

GetAssociatedObjectType returns the AssociatedObjectType field if non-nil, zero value otherwise.

### GetAssociatedObjectTypeOk

`func (o *PublicExportViewRequest) GetAssociatedObjectTypeOk() (*string, bool)`

GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectType

`func (o *PublicExportViewRequest) SetAssociatedObjectType(v string)`

SetAssociatedObjectType sets AssociatedObjectType field to given value.

### HasAssociatedObjectType

`func (o *PublicExportViewRequest) HasAssociatedObjectType() bool`

HasAssociatedObjectType returns a boolean if a field has been set.

### GetObjectType

`func (o *PublicExportViewRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PublicExportViewRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PublicExportViewRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLanguage

`func (o *PublicExportViewRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PublicExportViewRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PublicExportViewRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetExportInternalValuesOptions

`func (o *PublicExportViewRequest) GetExportInternalValuesOptions() []string`

GetExportInternalValuesOptions returns the ExportInternalValuesOptions field if non-nil, zero value otherwise.

### GetExportInternalValuesOptionsOk

`func (o *PublicExportViewRequest) GetExportInternalValuesOptionsOk() (*[]string, bool)`

GetExportInternalValuesOptionsOk returns a tuple with the ExportInternalValuesOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportInternalValuesOptions

`func (o *PublicExportViewRequest) SetExportInternalValuesOptions(v []string)`

SetExportInternalValuesOptions sets ExportInternalValuesOptions field to given value.


### GetOverrideAssociatedObjectsPerDefinitionPerRowLimit

`func (o *PublicExportViewRequest) GetOverrideAssociatedObjectsPerDefinitionPerRowLimit() bool`

GetOverrideAssociatedObjectsPerDefinitionPerRowLimit returns the OverrideAssociatedObjectsPerDefinitionPerRowLimit field if non-nil, zero value otherwise.

### GetOverrideAssociatedObjectsPerDefinitionPerRowLimitOk

`func (o *PublicExportViewRequest) GetOverrideAssociatedObjectsPerDefinitionPerRowLimitOk() (*bool, bool)`

GetOverrideAssociatedObjectsPerDefinitionPerRowLimitOk returns a tuple with the OverrideAssociatedObjectsPerDefinitionPerRowLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideAssociatedObjectsPerDefinitionPerRowLimit

`func (o *PublicExportViewRequest) SetOverrideAssociatedObjectsPerDefinitionPerRowLimit(v bool)`

SetOverrideAssociatedObjectsPerDefinitionPerRowLimit sets OverrideAssociatedObjectsPerDefinitionPerRowLimit field to given value.


### GetPublicCrmSearchRequest

`func (o *PublicExportViewRequest) GetPublicCrmSearchRequest() PublicCrmSearchRequest`

GetPublicCrmSearchRequest returns the PublicCrmSearchRequest field if non-nil, zero value otherwise.

### GetPublicCrmSearchRequestOk

`func (o *PublicExportViewRequest) GetPublicCrmSearchRequestOk() (*PublicCrmSearchRequest, bool)`

GetPublicCrmSearchRequestOk returns a tuple with the PublicCrmSearchRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicCrmSearchRequest

`func (o *PublicExportViewRequest) SetPublicCrmSearchRequest(v PublicCrmSearchRequest)`

SetPublicCrmSearchRequest sets PublicCrmSearchRequest field to given value.

### HasPublicCrmSearchRequest

`func (o *PublicExportViewRequest) HasPublicCrmSearchRequest() bool`

HasPublicCrmSearchRequest returns a boolean if a field has been set.

### GetListId

`func (o *PublicExportViewRequest) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *PublicExportViewRequest) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *PublicExportViewRequest) SetListId(v string)`

SetListId sets ListId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



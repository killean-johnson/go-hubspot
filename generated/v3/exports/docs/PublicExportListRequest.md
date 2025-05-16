# PublicExportListRequest

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
**PublicCrmSearchRequest** | Pointer to [**PublicCrmSearchRequest**](PublicCrmSearchRequest.md) |  | [optional] 
**ListId** | **string** |  | 

## Methods

### NewPublicExportListRequest

`func NewPublicExportListRequest(exportType string, format string, exportName string, objectProperties []string, objectType string, language string, exportInternalValuesOptions []string, overrideAssociatedObjectsPerDefinitionPerRowLimit bool, listId string, ) *PublicExportListRequest`

NewPublicExportListRequest instantiates a new PublicExportListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicExportListRequestWithDefaults

`func NewPublicExportListRequestWithDefaults() *PublicExportListRequest`

NewPublicExportListRequestWithDefaults instantiates a new PublicExportListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportType

`func (o *PublicExportListRequest) GetExportType() string`

GetExportType returns the ExportType field if non-nil, zero value otherwise.

### GetExportTypeOk

`func (o *PublicExportListRequest) GetExportTypeOk() (*string, bool)`

GetExportTypeOk returns a tuple with the ExportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportType

`func (o *PublicExportListRequest) SetExportType(v string)`

SetExportType sets ExportType field to given value.


### GetFormat

`func (o *PublicExportListRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PublicExportListRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PublicExportListRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetExportName

`func (o *PublicExportListRequest) GetExportName() string`

GetExportName returns the ExportName field if non-nil, zero value otherwise.

### GetExportNameOk

`func (o *PublicExportListRequest) GetExportNameOk() (*string, bool)`

GetExportNameOk returns a tuple with the ExportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportName

`func (o *PublicExportListRequest) SetExportName(v string)`

SetExportName sets ExportName field to given value.


### GetObjectProperties

`func (o *PublicExportListRequest) GetObjectProperties() []string`

GetObjectProperties returns the ObjectProperties field if non-nil, zero value otherwise.

### GetObjectPropertiesOk

`func (o *PublicExportListRequest) GetObjectPropertiesOk() (*[]string, bool)`

GetObjectPropertiesOk returns a tuple with the ObjectProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProperties

`func (o *PublicExportListRequest) SetObjectProperties(v []string)`

SetObjectProperties sets ObjectProperties field to given value.


### GetAssociatedObjectType

`func (o *PublicExportListRequest) GetAssociatedObjectType() string`

GetAssociatedObjectType returns the AssociatedObjectType field if non-nil, zero value otherwise.

### GetAssociatedObjectTypeOk

`func (o *PublicExportListRequest) GetAssociatedObjectTypeOk() (*string, bool)`

GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectType

`func (o *PublicExportListRequest) SetAssociatedObjectType(v string)`

SetAssociatedObjectType sets AssociatedObjectType field to given value.

### HasAssociatedObjectType

`func (o *PublicExportListRequest) HasAssociatedObjectType() bool`

HasAssociatedObjectType returns a boolean if a field has been set.

### GetObjectType

`func (o *PublicExportListRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PublicExportListRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PublicExportListRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLanguage

`func (o *PublicExportListRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PublicExportListRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PublicExportListRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetExportInternalValuesOptions

`func (o *PublicExportListRequest) GetExportInternalValuesOptions() []string`

GetExportInternalValuesOptions returns the ExportInternalValuesOptions field if non-nil, zero value otherwise.

### GetExportInternalValuesOptionsOk

`func (o *PublicExportListRequest) GetExportInternalValuesOptionsOk() (*[]string, bool)`

GetExportInternalValuesOptionsOk returns a tuple with the ExportInternalValuesOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportInternalValuesOptions

`func (o *PublicExportListRequest) SetExportInternalValuesOptions(v []string)`

SetExportInternalValuesOptions sets ExportInternalValuesOptions field to given value.


### GetOverrideAssociatedObjectsPerDefinitionPerRowLimit

`func (o *PublicExportListRequest) GetOverrideAssociatedObjectsPerDefinitionPerRowLimit() bool`

GetOverrideAssociatedObjectsPerDefinitionPerRowLimit returns the OverrideAssociatedObjectsPerDefinitionPerRowLimit field if non-nil, zero value otherwise.

### GetOverrideAssociatedObjectsPerDefinitionPerRowLimitOk

`func (o *PublicExportListRequest) GetOverrideAssociatedObjectsPerDefinitionPerRowLimitOk() (*bool, bool)`

GetOverrideAssociatedObjectsPerDefinitionPerRowLimitOk returns a tuple with the OverrideAssociatedObjectsPerDefinitionPerRowLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideAssociatedObjectsPerDefinitionPerRowLimit

`func (o *PublicExportListRequest) SetOverrideAssociatedObjectsPerDefinitionPerRowLimit(v bool)`

SetOverrideAssociatedObjectsPerDefinitionPerRowLimit sets OverrideAssociatedObjectsPerDefinitionPerRowLimit field to given value.


### GetPublicCrmSearchRequest

`func (o *PublicExportListRequest) GetPublicCrmSearchRequest() PublicCrmSearchRequest`

GetPublicCrmSearchRequest returns the PublicCrmSearchRequest field if non-nil, zero value otherwise.

### GetPublicCrmSearchRequestOk

`func (o *PublicExportListRequest) GetPublicCrmSearchRequestOk() (*PublicCrmSearchRequest, bool)`

GetPublicCrmSearchRequestOk returns a tuple with the PublicCrmSearchRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicCrmSearchRequest

`func (o *PublicExportListRequest) SetPublicCrmSearchRequest(v PublicCrmSearchRequest)`

SetPublicCrmSearchRequest sets PublicCrmSearchRequest field to given value.

### HasPublicCrmSearchRequest

`func (o *PublicExportListRequest) HasPublicCrmSearchRequest() bool`

HasPublicCrmSearchRequest returns a boolean if a field has been set.

### GetListId

`func (o *PublicExportListRequest) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *PublicExportListRequest) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *PublicExportListRequest) SetListId(v string)`

SetListId sets ListId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



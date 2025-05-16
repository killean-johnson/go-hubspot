# PublicConversationsMessageAttachmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUsageType** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | [default to "FILE"]
**Url** | Pointer to **string** |  | [optional] 
**FileId** | **int64** |  | 
**Address** | Pointer to **string** |  | [optional] 
**Latitude** | **float32** |  | 
**Longitude** | **float32** |  | 
**ContactProfile** | [**ContactProfile**](ContactProfile.md) |  | 
**Text** | Pointer to **string** |  | [optional] 
**QuickReplies** | [**[]QuickReply**](QuickReply.md) |  | 
**AllowMultiSelect** | **bool** |  | 
**AllowUserInput** | **bool** |  | 
**CrmObjectIds** | **map[string]int64** |  | 
**MappedTemplateId** | **string** |  | 
**Parameters** | **map[string]string** |  | 

## Methods

### NewPublicConversationsMessageAttachmentsInner

`func NewPublicConversationsMessageAttachmentsInner(fileUsageType string, type_ string, fileId int64, latitude float32, longitude float32, contactProfile ContactProfile, quickReplies []QuickReply, allowMultiSelect bool, allowUserInput bool, crmObjectIds map[string]int64, mappedTemplateId string, parameters map[string]string, ) *PublicConversationsMessageAttachmentsInner`

NewPublicConversationsMessageAttachmentsInner instantiates a new PublicConversationsMessageAttachmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicConversationsMessageAttachmentsInnerWithDefaults

`func NewPublicConversationsMessageAttachmentsInnerWithDefaults() *PublicConversationsMessageAttachmentsInner`

NewPublicConversationsMessageAttachmentsInnerWithDefaults instantiates a new PublicConversationsMessageAttachmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUsageType

`func (o *PublicConversationsMessageAttachmentsInner) GetFileUsageType() string`

GetFileUsageType returns the FileUsageType field if non-nil, zero value otherwise.

### GetFileUsageTypeOk

`func (o *PublicConversationsMessageAttachmentsInner) GetFileUsageTypeOk() (*string, bool)`

GetFileUsageTypeOk returns a tuple with the FileUsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUsageType

`func (o *PublicConversationsMessageAttachmentsInner) SetFileUsageType(v string)`

SetFileUsageType sets FileUsageType field to given value.


### GetName

`func (o *PublicConversationsMessageAttachmentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicConversationsMessageAttachmentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicConversationsMessageAttachmentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicConversationsMessageAttachmentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PublicConversationsMessageAttachmentsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicConversationsMessageAttachmentsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicConversationsMessageAttachmentsInner) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *PublicConversationsMessageAttachmentsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PublicConversationsMessageAttachmentsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PublicConversationsMessageAttachmentsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PublicConversationsMessageAttachmentsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFileId

`func (o *PublicConversationsMessageAttachmentsInner) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *PublicConversationsMessageAttachmentsInner) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *PublicConversationsMessageAttachmentsInner) SetFileId(v int64)`

SetFileId sets FileId field to given value.


### GetAddress

`func (o *PublicConversationsMessageAttachmentsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PublicConversationsMessageAttachmentsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PublicConversationsMessageAttachmentsInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PublicConversationsMessageAttachmentsInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *PublicConversationsMessageAttachmentsInner) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PublicConversationsMessageAttachmentsInner) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PublicConversationsMessageAttachmentsInner) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *PublicConversationsMessageAttachmentsInner) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PublicConversationsMessageAttachmentsInner) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PublicConversationsMessageAttachmentsInner) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetContactProfile

`func (o *PublicConversationsMessageAttachmentsInner) GetContactProfile() ContactProfile`

GetContactProfile returns the ContactProfile field if non-nil, zero value otherwise.

### GetContactProfileOk

`func (o *PublicConversationsMessageAttachmentsInner) GetContactProfileOk() (*ContactProfile, bool)`

GetContactProfileOk returns a tuple with the ContactProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactProfile

`func (o *PublicConversationsMessageAttachmentsInner) SetContactProfile(v ContactProfile)`

SetContactProfile sets ContactProfile field to given value.


### GetText

`func (o *PublicConversationsMessageAttachmentsInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PublicConversationsMessageAttachmentsInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PublicConversationsMessageAttachmentsInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PublicConversationsMessageAttachmentsInner) HasText() bool`

HasText returns a boolean if a field has been set.

### GetQuickReplies

`func (o *PublicConversationsMessageAttachmentsInner) GetQuickReplies() []QuickReply`

GetQuickReplies returns the QuickReplies field if non-nil, zero value otherwise.

### GetQuickRepliesOk

`func (o *PublicConversationsMessageAttachmentsInner) GetQuickRepliesOk() (*[]QuickReply, bool)`

GetQuickRepliesOk returns a tuple with the QuickReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickReplies

`func (o *PublicConversationsMessageAttachmentsInner) SetQuickReplies(v []QuickReply)`

SetQuickReplies sets QuickReplies field to given value.


### GetAllowMultiSelect

`func (o *PublicConversationsMessageAttachmentsInner) GetAllowMultiSelect() bool`

GetAllowMultiSelect returns the AllowMultiSelect field if non-nil, zero value otherwise.

### GetAllowMultiSelectOk

`func (o *PublicConversationsMessageAttachmentsInner) GetAllowMultiSelectOk() (*bool, bool)`

GetAllowMultiSelectOk returns a tuple with the AllowMultiSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiSelect

`func (o *PublicConversationsMessageAttachmentsInner) SetAllowMultiSelect(v bool)`

SetAllowMultiSelect sets AllowMultiSelect field to given value.


### GetAllowUserInput

`func (o *PublicConversationsMessageAttachmentsInner) GetAllowUserInput() bool`

GetAllowUserInput returns the AllowUserInput field if non-nil, zero value otherwise.

### GetAllowUserInputOk

`func (o *PublicConversationsMessageAttachmentsInner) GetAllowUserInputOk() (*bool, bool)`

GetAllowUserInputOk returns a tuple with the AllowUserInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserInput

`func (o *PublicConversationsMessageAttachmentsInner) SetAllowUserInput(v bool)`

SetAllowUserInput sets AllowUserInput field to given value.


### GetCrmObjectIds

`func (o *PublicConversationsMessageAttachmentsInner) GetCrmObjectIds() map[string]int64`

GetCrmObjectIds returns the CrmObjectIds field if non-nil, zero value otherwise.

### GetCrmObjectIdsOk

`func (o *PublicConversationsMessageAttachmentsInner) GetCrmObjectIdsOk() (*map[string]int64, bool)`

GetCrmObjectIdsOk returns a tuple with the CrmObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmObjectIds

`func (o *PublicConversationsMessageAttachmentsInner) SetCrmObjectIds(v map[string]int64)`

SetCrmObjectIds sets CrmObjectIds field to given value.


### GetMappedTemplateId

`func (o *PublicConversationsMessageAttachmentsInner) GetMappedTemplateId() string`

GetMappedTemplateId returns the MappedTemplateId field if non-nil, zero value otherwise.

### GetMappedTemplateIdOk

`func (o *PublicConversationsMessageAttachmentsInner) GetMappedTemplateIdOk() (*string, bool)`

GetMappedTemplateIdOk returns a tuple with the MappedTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedTemplateId

`func (o *PublicConversationsMessageAttachmentsInner) SetMappedTemplateId(v string)`

SetMappedTemplateId sets MappedTemplateId field to given value.


### GetParameters

`func (o *PublicConversationsMessageAttachmentsInner) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PublicConversationsMessageAttachmentsInner) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PublicConversationsMessageAttachmentsInner) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



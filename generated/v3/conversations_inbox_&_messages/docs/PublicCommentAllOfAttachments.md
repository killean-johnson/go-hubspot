# PublicCommentAllOfAttachments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUsageType** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | [default to "FILE"]
**Url** | **string** |  | 
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

### NewPublicCommentAllOfAttachments

`func NewPublicCommentAllOfAttachments(fileUsageType string, type_ string, url string, fileId int64, latitude float32, longitude float32, contactProfile ContactProfile, quickReplies []QuickReply, allowMultiSelect bool, allowUserInput bool, crmObjectIds map[string]int64, mappedTemplateId string, parameters map[string]string, ) *PublicCommentAllOfAttachments`

NewPublicCommentAllOfAttachments instantiates a new PublicCommentAllOfAttachments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCommentAllOfAttachmentsWithDefaults

`func NewPublicCommentAllOfAttachmentsWithDefaults() *PublicCommentAllOfAttachments`

NewPublicCommentAllOfAttachmentsWithDefaults instantiates a new PublicCommentAllOfAttachments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUsageType

`func (o *PublicCommentAllOfAttachments) GetFileUsageType() string`

GetFileUsageType returns the FileUsageType field if non-nil, zero value otherwise.

### GetFileUsageTypeOk

`func (o *PublicCommentAllOfAttachments) GetFileUsageTypeOk() (*string, bool)`

GetFileUsageTypeOk returns a tuple with the FileUsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUsageType

`func (o *PublicCommentAllOfAttachments) SetFileUsageType(v string)`

SetFileUsageType sets FileUsageType field to given value.


### GetName

`func (o *PublicCommentAllOfAttachments) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicCommentAllOfAttachments) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicCommentAllOfAttachments) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicCommentAllOfAttachments) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PublicCommentAllOfAttachments) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicCommentAllOfAttachments) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicCommentAllOfAttachments) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *PublicCommentAllOfAttachments) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PublicCommentAllOfAttachments) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PublicCommentAllOfAttachments) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFileId

`func (o *PublicCommentAllOfAttachments) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *PublicCommentAllOfAttachments) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *PublicCommentAllOfAttachments) SetFileId(v int64)`

SetFileId sets FileId field to given value.


### GetAddress

`func (o *PublicCommentAllOfAttachments) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PublicCommentAllOfAttachments) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PublicCommentAllOfAttachments) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PublicCommentAllOfAttachments) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *PublicCommentAllOfAttachments) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PublicCommentAllOfAttachments) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PublicCommentAllOfAttachments) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *PublicCommentAllOfAttachments) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PublicCommentAllOfAttachments) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PublicCommentAllOfAttachments) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetContactProfile

`func (o *PublicCommentAllOfAttachments) GetContactProfile() ContactProfile`

GetContactProfile returns the ContactProfile field if non-nil, zero value otherwise.

### GetContactProfileOk

`func (o *PublicCommentAllOfAttachments) GetContactProfileOk() (*ContactProfile, bool)`

GetContactProfileOk returns a tuple with the ContactProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactProfile

`func (o *PublicCommentAllOfAttachments) SetContactProfile(v ContactProfile)`

SetContactProfile sets ContactProfile field to given value.


### GetText

`func (o *PublicCommentAllOfAttachments) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PublicCommentAllOfAttachments) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PublicCommentAllOfAttachments) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PublicCommentAllOfAttachments) HasText() bool`

HasText returns a boolean if a field has been set.

### GetQuickReplies

`func (o *PublicCommentAllOfAttachments) GetQuickReplies() []QuickReply`

GetQuickReplies returns the QuickReplies field if non-nil, zero value otherwise.

### GetQuickRepliesOk

`func (o *PublicCommentAllOfAttachments) GetQuickRepliesOk() (*[]QuickReply, bool)`

GetQuickRepliesOk returns a tuple with the QuickReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickReplies

`func (o *PublicCommentAllOfAttachments) SetQuickReplies(v []QuickReply)`

SetQuickReplies sets QuickReplies field to given value.


### GetAllowMultiSelect

`func (o *PublicCommentAllOfAttachments) GetAllowMultiSelect() bool`

GetAllowMultiSelect returns the AllowMultiSelect field if non-nil, zero value otherwise.

### GetAllowMultiSelectOk

`func (o *PublicCommentAllOfAttachments) GetAllowMultiSelectOk() (*bool, bool)`

GetAllowMultiSelectOk returns a tuple with the AllowMultiSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiSelect

`func (o *PublicCommentAllOfAttachments) SetAllowMultiSelect(v bool)`

SetAllowMultiSelect sets AllowMultiSelect field to given value.


### GetAllowUserInput

`func (o *PublicCommentAllOfAttachments) GetAllowUserInput() bool`

GetAllowUserInput returns the AllowUserInput field if non-nil, zero value otherwise.

### GetAllowUserInputOk

`func (o *PublicCommentAllOfAttachments) GetAllowUserInputOk() (*bool, bool)`

GetAllowUserInputOk returns a tuple with the AllowUserInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserInput

`func (o *PublicCommentAllOfAttachments) SetAllowUserInput(v bool)`

SetAllowUserInput sets AllowUserInput field to given value.


### GetCrmObjectIds

`func (o *PublicCommentAllOfAttachments) GetCrmObjectIds() map[string]int64`

GetCrmObjectIds returns the CrmObjectIds field if non-nil, zero value otherwise.

### GetCrmObjectIdsOk

`func (o *PublicCommentAllOfAttachments) GetCrmObjectIdsOk() (*map[string]int64, bool)`

GetCrmObjectIdsOk returns a tuple with the CrmObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmObjectIds

`func (o *PublicCommentAllOfAttachments) SetCrmObjectIds(v map[string]int64)`

SetCrmObjectIds sets CrmObjectIds field to given value.


### GetMappedTemplateId

`func (o *PublicCommentAllOfAttachments) GetMappedTemplateId() string`

GetMappedTemplateId returns the MappedTemplateId field if non-nil, zero value otherwise.

### GetMappedTemplateIdOk

`func (o *PublicCommentAllOfAttachments) GetMappedTemplateIdOk() (*string, bool)`

GetMappedTemplateIdOk returns a tuple with the MappedTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedTemplateId

`func (o *PublicCommentAllOfAttachments) SetMappedTemplateId(v string)`

SetMappedTemplateId sets MappedTemplateId field to given value.


### GetParameters

`func (o *PublicCommentAllOfAttachments) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PublicCommentAllOfAttachments) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PublicCommentAllOfAttachments) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



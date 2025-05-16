# MessageHeaderAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | [default to "MESSAGE_HEADER"]
**FileId** | Pointer to **int64** |  | [optional] 

## Methods

### NewMessageHeaderAttachment

`func NewMessageHeaderAttachment(type_ string, ) *MessageHeaderAttachment`

NewMessageHeaderAttachment instantiates a new MessageHeaderAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageHeaderAttachmentWithDefaults

`func NewMessageHeaderAttachmentWithDefaults() *MessageHeaderAttachment`

NewMessageHeaderAttachmentWithDefaults instantiates a new MessageHeaderAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *MessageHeaderAttachment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MessageHeaderAttachment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MessageHeaderAttachment) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MessageHeaderAttachment) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *MessageHeaderAttachment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageHeaderAttachment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageHeaderAttachment) SetType(v string)`

SetType sets Type field to given value.


### GetFileId

`func (o *MessageHeaderAttachment) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *MessageHeaderAttachment) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *MessageHeaderAttachment) SetFileId(v int64)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *MessageHeaderAttachment) HasFileId() bool`

HasFileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



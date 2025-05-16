# QuickRepliesAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuickReplies** | [**[]QuickReply**](QuickReply.md) |  | 
**Type** | **string** |  | [default to "QUICK_REPLIES"]

## Methods

### NewQuickRepliesAttachment

`func NewQuickRepliesAttachment(quickReplies []QuickReply, type_ string, ) *QuickRepliesAttachment`

NewQuickRepliesAttachment instantiates a new QuickRepliesAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickRepliesAttachmentWithDefaults

`func NewQuickRepliesAttachmentWithDefaults() *QuickRepliesAttachment`

NewQuickRepliesAttachmentWithDefaults instantiates a new QuickRepliesAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuickReplies

`func (o *QuickRepliesAttachment) GetQuickReplies() []QuickReply`

GetQuickReplies returns the QuickReplies field if non-nil, zero value otherwise.

### GetQuickRepliesOk

`func (o *QuickRepliesAttachment) GetQuickRepliesOk() (*[]QuickReply, bool)`

GetQuickRepliesOk returns a tuple with the QuickReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickReplies

`func (o *QuickRepliesAttachment) SetQuickReplies(v []QuickReply)`

SetQuickReplies sets QuickReplies field to given value.


### GetType

`func (o *QuickRepliesAttachment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuickRepliesAttachment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuickRepliesAttachment) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



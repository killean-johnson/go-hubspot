# PublicQuickReplies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuickReplies** | [**[]QuickReply**](QuickReply.md) |  | 
**AllowMultiSelect** | **bool** |  | 
**AllowUserInput** | **bool** |  | 
**Type** | **string** |  | [default to "QUICK_REPLIES"]

## Methods

### NewPublicQuickReplies

`func NewPublicQuickReplies(quickReplies []QuickReply, allowMultiSelect bool, allowUserInput bool, type_ string, ) *PublicQuickReplies`

NewPublicQuickReplies instantiates a new PublicQuickReplies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicQuickRepliesWithDefaults

`func NewPublicQuickRepliesWithDefaults() *PublicQuickReplies`

NewPublicQuickRepliesWithDefaults instantiates a new PublicQuickReplies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuickReplies

`func (o *PublicQuickReplies) GetQuickReplies() []QuickReply`

GetQuickReplies returns the QuickReplies field if non-nil, zero value otherwise.

### GetQuickRepliesOk

`func (o *PublicQuickReplies) GetQuickRepliesOk() (*[]QuickReply, bool)`

GetQuickRepliesOk returns a tuple with the QuickReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickReplies

`func (o *PublicQuickReplies) SetQuickReplies(v []QuickReply)`

SetQuickReplies sets QuickReplies field to given value.


### GetAllowMultiSelect

`func (o *PublicQuickReplies) GetAllowMultiSelect() bool`

GetAllowMultiSelect returns the AllowMultiSelect field if non-nil, zero value otherwise.

### GetAllowMultiSelectOk

`func (o *PublicQuickReplies) GetAllowMultiSelectOk() (*bool, bool)`

GetAllowMultiSelectOk returns a tuple with the AllowMultiSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiSelect

`func (o *PublicQuickReplies) SetAllowMultiSelect(v bool)`

SetAllowMultiSelect sets AllowMultiSelect field to given value.


### GetAllowUserInput

`func (o *PublicQuickReplies) GetAllowUserInput() bool`

GetAllowUserInput returns the AllowUserInput field if non-nil, zero value otherwise.

### GetAllowUserInputOk

`func (o *PublicQuickReplies) GetAllowUserInputOk() (*bool, bool)`

GetAllowUserInputOk returns a tuple with the AllowUserInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserInput

`func (o *PublicQuickReplies) SetAllowUserInput(v bool)`

SetAllowUserInput sets AllowUserInput field to given value.


### GetType

`func (o *PublicQuickReplies) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicQuickReplies) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicQuickReplies) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



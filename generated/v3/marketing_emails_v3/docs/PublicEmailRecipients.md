# PublicEmailRecipients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Include** | Pointer to **[]string** | Included IDs. | [optional] 
**Exclude** | Pointer to **[]string** | Excluded IDs. | [optional] 

## Methods

### NewPublicEmailRecipients

`func NewPublicEmailRecipients() *PublicEmailRecipients`

NewPublicEmailRecipients instantiates a new PublicEmailRecipients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicEmailRecipientsWithDefaults

`func NewPublicEmailRecipientsWithDefaults() *PublicEmailRecipients`

NewPublicEmailRecipientsWithDefaults instantiates a new PublicEmailRecipients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInclude

`func (o *PublicEmailRecipients) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *PublicEmailRecipients) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *PublicEmailRecipients) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *PublicEmailRecipients) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetExclude

`func (o *PublicEmailRecipients) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *PublicEmailRecipients) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *PublicEmailRecipients) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *PublicEmailRecipients) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



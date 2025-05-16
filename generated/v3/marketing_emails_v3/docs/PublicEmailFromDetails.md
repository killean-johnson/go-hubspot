# PublicEmailFromDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomReplyTo** | Pointer to **string** | The reply to recipients will see. | [optional] 
**FromName** | Pointer to **string** | The name recipients will see. | [optional] 
**ReplyTo** | Pointer to **string** | The from address and reply to email address (if no customReplyTo defined) recipients will see. | [optional] 

## Methods

### NewPublicEmailFromDetails

`func NewPublicEmailFromDetails() *PublicEmailFromDetails`

NewPublicEmailFromDetails instantiates a new PublicEmailFromDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicEmailFromDetailsWithDefaults

`func NewPublicEmailFromDetailsWithDefaults() *PublicEmailFromDetails`

NewPublicEmailFromDetailsWithDefaults instantiates a new PublicEmailFromDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomReplyTo

`func (o *PublicEmailFromDetails) GetCustomReplyTo() string`

GetCustomReplyTo returns the CustomReplyTo field if non-nil, zero value otherwise.

### GetCustomReplyToOk

`func (o *PublicEmailFromDetails) GetCustomReplyToOk() (*string, bool)`

GetCustomReplyToOk returns a tuple with the CustomReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomReplyTo

`func (o *PublicEmailFromDetails) SetCustomReplyTo(v string)`

SetCustomReplyTo sets CustomReplyTo field to given value.

### HasCustomReplyTo

`func (o *PublicEmailFromDetails) HasCustomReplyTo() bool`

HasCustomReplyTo returns a boolean if a field has been set.

### GetFromName

`func (o *PublicEmailFromDetails) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *PublicEmailFromDetails) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *PublicEmailFromDetails) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *PublicEmailFromDetails) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### GetReplyTo

`func (o *PublicEmailFromDetails) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *PublicEmailFromDetails) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *PublicEmailFromDetails) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *PublicEmailFromDetails) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



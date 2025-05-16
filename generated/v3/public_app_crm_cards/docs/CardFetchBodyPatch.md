# CardFetchBodyPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerlessFunction** | Pointer to **string** |  | [optional] 
**CardType** | Pointer to **string** |  | [optional] 
**ObjectTypes** | [**[]CardObjectTypeBody**](CardObjectTypeBody.md) | An array of CRM object types where this card should be displayed. HubSpot will call your target URL whenever a user visits a record page of the types defined here. | 
**TargetUrl** | Pointer to **string** | URL to a service endpoint that will respond with details for this card. HubSpot will call this endpoint each time a user visits a CRM record page where this card should be displayed. | [optional] 

## Methods

### NewCardFetchBodyPatch

`func NewCardFetchBodyPatch(objectTypes []CardObjectTypeBody, ) *CardFetchBodyPatch`

NewCardFetchBodyPatch instantiates a new CardFetchBodyPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardFetchBodyPatchWithDefaults

`func NewCardFetchBodyPatchWithDefaults() *CardFetchBodyPatch`

NewCardFetchBodyPatchWithDefaults instantiates a new CardFetchBodyPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerlessFunction

`func (o *CardFetchBodyPatch) GetServerlessFunction() string`

GetServerlessFunction returns the ServerlessFunction field if non-nil, zero value otherwise.

### GetServerlessFunctionOk

`func (o *CardFetchBodyPatch) GetServerlessFunctionOk() (*string, bool)`

GetServerlessFunctionOk returns a tuple with the ServerlessFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerlessFunction

`func (o *CardFetchBodyPatch) SetServerlessFunction(v string)`

SetServerlessFunction sets ServerlessFunction field to given value.

### HasServerlessFunction

`func (o *CardFetchBodyPatch) HasServerlessFunction() bool`

HasServerlessFunction returns a boolean if a field has been set.

### GetCardType

`func (o *CardFetchBodyPatch) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CardFetchBodyPatch) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CardFetchBodyPatch) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *CardFetchBodyPatch) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetObjectTypes

`func (o *CardFetchBodyPatch) GetObjectTypes() []CardObjectTypeBody`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *CardFetchBodyPatch) GetObjectTypesOk() (*[]CardObjectTypeBody, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *CardFetchBodyPatch) SetObjectTypes(v []CardObjectTypeBody)`

SetObjectTypes sets ObjectTypes field to given value.


### GetTargetUrl

`func (o *CardFetchBodyPatch) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *CardFetchBodyPatch) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *CardFetchBodyPatch) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *CardFetchBodyPatch) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



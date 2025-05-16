# ContactPhone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewContactPhone

`func NewContactPhone(phone string, ) *ContactPhone`

NewContactPhone instantiates a new ContactPhone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactPhoneWithDefaults

`func NewContactPhoneWithDefaults() *ContactPhone`

NewContactPhoneWithDefaults instantiates a new ContactPhone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *ContactPhone) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ContactPhone) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ContactPhone) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetType

`func (o *ContactPhone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactPhone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactPhone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContactPhone) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



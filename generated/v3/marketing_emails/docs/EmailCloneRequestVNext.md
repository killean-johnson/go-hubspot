# EmailCloneRequestVNext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneName** | Pointer to **string** | The name to assign to the cloned email. | [optional] 
**Language** | Pointer to **string** | The language code for the cloned email, such as &#39;en&#39; for English. | [optional] 
**Id** | **string** | The unique identifier of the email to be cloned. | 

## Methods

### NewEmailCloneRequestVNext

`func NewEmailCloneRequestVNext(id string, ) *EmailCloneRequestVNext`

NewEmailCloneRequestVNext instantiates a new EmailCloneRequestVNext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailCloneRequestVNextWithDefaults

`func NewEmailCloneRequestVNextWithDefaults() *EmailCloneRequestVNext`

NewEmailCloneRequestVNextWithDefaults instantiates a new EmailCloneRequestVNext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneName

`func (o *EmailCloneRequestVNext) GetCloneName() string`

GetCloneName returns the CloneName field if non-nil, zero value otherwise.

### GetCloneNameOk

`func (o *EmailCloneRequestVNext) GetCloneNameOk() (*string, bool)`

GetCloneNameOk returns a tuple with the CloneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneName

`func (o *EmailCloneRequestVNext) SetCloneName(v string)`

SetCloneName sets CloneName field to given value.

### HasCloneName

`func (o *EmailCloneRequestVNext) HasCloneName() bool`

HasCloneName returns a boolean if a field has been set.

### GetLanguage

`func (o *EmailCloneRequestVNext) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *EmailCloneRequestVNext) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *EmailCloneRequestVNext) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *EmailCloneRequestVNext) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetId

`func (o *EmailCloneRequestVNext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailCloneRequestVNext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailCloneRequestVNext) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



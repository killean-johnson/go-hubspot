# PublicMessageFailureDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessageTokens** | **map[string]string** |  | 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicMessageFailureDetails

`func NewPublicMessageFailureDetails(errorMessageTokens map[string]string, ) *PublicMessageFailureDetails`

NewPublicMessageFailureDetails instantiates a new PublicMessageFailureDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicMessageFailureDetailsWithDefaults

`func NewPublicMessageFailureDetailsWithDefaults() *PublicMessageFailureDetails`

NewPublicMessageFailureDetailsWithDefaults instantiates a new PublicMessageFailureDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessageTokens

`func (o *PublicMessageFailureDetails) GetErrorMessageTokens() map[string]string`

GetErrorMessageTokens returns the ErrorMessageTokens field if non-nil, zero value otherwise.

### GetErrorMessageTokensOk

`func (o *PublicMessageFailureDetails) GetErrorMessageTokensOk() (*map[string]string, bool)`

GetErrorMessageTokensOk returns a tuple with the ErrorMessageTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessageTokens

`func (o *PublicMessageFailureDetails) SetErrorMessageTokens(v map[string]string)`

SetErrorMessageTokens sets ErrorMessageTokens field to given value.


### GetErrorMessage

`func (o *PublicMessageFailureDetails) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PublicMessageFailureDetails) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PublicMessageFailureDetails) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PublicMessageFailureDetails) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PublicMessageStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusType** | **string** |  | 
**FailureDetails** | Pointer to [**PublicMessageFailureDetails**](PublicMessageFailureDetails.md) |  | [optional] 

## Methods

### NewPublicMessageStatus

`func NewPublicMessageStatus(statusType string, ) *PublicMessageStatus`

NewPublicMessageStatus instantiates a new PublicMessageStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicMessageStatusWithDefaults

`func NewPublicMessageStatusWithDefaults() *PublicMessageStatus`

NewPublicMessageStatusWithDefaults instantiates a new PublicMessageStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusType

`func (o *PublicMessageStatus) GetStatusType() string`

GetStatusType returns the StatusType field if non-nil, zero value otherwise.

### GetStatusTypeOk

`func (o *PublicMessageStatus) GetStatusTypeOk() (*string, bool)`

GetStatusTypeOk returns a tuple with the StatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusType

`func (o *PublicMessageStatus) SetStatusType(v string)`

SetStatusType sets StatusType field to given value.


### GetFailureDetails

`func (o *PublicMessageStatus) GetFailureDetails() PublicMessageFailureDetails`

GetFailureDetails returns the FailureDetails field if non-nil, zero value otherwise.

### GetFailureDetailsOk

`func (o *PublicMessageStatus) GetFailureDetailsOk() (*PublicMessageFailureDetails, bool)`

GetFailureDetailsOk returns a tuple with the FailureDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureDetails

`func (o *PublicMessageStatus) SetFailureDetails(v PublicMessageFailureDetails)`

SetFailureDetails sets FailureDetails field to given value.

### HasFailureDetails

`func (o *PublicMessageStatus) HasFailureDetails() bool`

HasFailureDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PartialPublicStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusState** | **string** |  | 
**Channel** | **string** |  | 
**LegalBasis** | Pointer to **string** |  | [optional] 
**SubscriptionId** | **int64** |  | 
**LegalBasisExplanation** | Pointer to **string** |  | [optional] 

## Methods

### NewPartialPublicStatusRequest

`func NewPartialPublicStatusRequest(statusState string, channel string, subscriptionId int64, ) *PartialPublicStatusRequest`

NewPartialPublicStatusRequest instantiates a new PartialPublicStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialPublicStatusRequestWithDefaults

`func NewPartialPublicStatusRequestWithDefaults() *PartialPublicStatusRequest`

NewPartialPublicStatusRequestWithDefaults instantiates a new PartialPublicStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusState

`func (o *PartialPublicStatusRequest) GetStatusState() string`

GetStatusState returns the StatusState field if non-nil, zero value otherwise.

### GetStatusStateOk

`func (o *PartialPublicStatusRequest) GetStatusStateOk() (*string, bool)`

GetStatusStateOk returns a tuple with the StatusState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusState

`func (o *PartialPublicStatusRequest) SetStatusState(v string)`

SetStatusState sets StatusState field to given value.


### GetChannel

`func (o *PartialPublicStatusRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PartialPublicStatusRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PartialPublicStatusRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetLegalBasis

`func (o *PartialPublicStatusRequest) GetLegalBasis() string`

GetLegalBasis returns the LegalBasis field if non-nil, zero value otherwise.

### GetLegalBasisOk

`func (o *PartialPublicStatusRequest) GetLegalBasisOk() (*string, bool)`

GetLegalBasisOk returns a tuple with the LegalBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalBasis

`func (o *PartialPublicStatusRequest) SetLegalBasis(v string)`

SetLegalBasis sets LegalBasis field to given value.

### HasLegalBasis

`func (o *PartialPublicStatusRequest) HasLegalBasis() bool`

HasLegalBasis returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *PartialPublicStatusRequest) GetSubscriptionId() int64`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PartialPublicStatusRequest) GetSubscriptionIdOk() (*int64, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PartialPublicStatusRequest) SetSubscriptionId(v int64)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetLegalBasisExplanation

`func (o *PartialPublicStatusRequest) GetLegalBasisExplanation() string`

GetLegalBasisExplanation returns the LegalBasisExplanation field if non-nil, zero value otherwise.

### GetLegalBasisExplanationOk

`func (o *PartialPublicStatusRequest) GetLegalBasisExplanationOk() (*string, bool)`

GetLegalBasisExplanationOk returns a tuple with the LegalBasisExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalBasisExplanation

`func (o *PartialPublicStatusRequest) SetLegalBasisExplanation(v string)`

SetLegalBasisExplanation sets LegalBasisExplanation field to given value.

### HasLegalBasisExplanation

`func (o *PartialPublicStatusRequest) HasLegalBasisExplanation() bool`

HasLegalBasisExplanation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



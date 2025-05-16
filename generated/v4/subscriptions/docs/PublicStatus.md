# PublicStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionName** | Pointer to **string** |  | [optional] 
**Channel** | **string** |  | 
**SubscriberIdString** | **string** |  | 
**LegalBasis** | Pointer to **string** |  | [optional] 
**SetStatusSuccessReason** | Pointer to **string** |  | [optional] 
**Source** | **string** |  | 
**SubscriptionId** | **int32** |  | 
**LegalBasisExplanation** | Pointer to **string** |  | [optional] 
**BusinessUnitId** | Pointer to **int64** |  | [optional] 
**Status** | **string** |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewPublicStatus

`func NewPublicStatus(channel string, subscriberIdString string, source string, subscriptionId int32, status string, timestamp time.Time, ) *PublicStatus`

NewPublicStatus instantiates a new PublicStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicStatusWithDefaults

`func NewPublicStatusWithDefaults() *PublicStatus`

NewPublicStatusWithDefaults instantiates a new PublicStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionName

`func (o *PublicStatus) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *PublicStatus) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *PublicStatus) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.

### HasSubscriptionName

`func (o *PublicStatus) HasSubscriptionName() bool`

HasSubscriptionName returns a boolean if a field has been set.

### GetChannel

`func (o *PublicStatus) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PublicStatus) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PublicStatus) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetSubscriberIdString

`func (o *PublicStatus) GetSubscriberIdString() string`

GetSubscriberIdString returns the SubscriberIdString field if non-nil, zero value otherwise.

### GetSubscriberIdStringOk

`func (o *PublicStatus) GetSubscriberIdStringOk() (*string, bool)`

GetSubscriberIdStringOk returns a tuple with the SubscriberIdString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberIdString

`func (o *PublicStatus) SetSubscriberIdString(v string)`

SetSubscriberIdString sets SubscriberIdString field to given value.


### GetLegalBasis

`func (o *PublicStatus) GetLegalBasis() string`

GetLegalBasis returns the LegalBasis field if non-nil, zero value otherwise.

### GetLegalBasisOk

`func (o *PublicStatus) GetLegalBasisOk() (*string, bool)`

GetLegalBasisOk returns a tuple with the LegalBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalBasis

`func (o *PublicStatus) SetLegalBasis(v string)`

SetLegalBasis sets LegalBasis field to given value.

### HasLegalBasis

`func (o *PublicStatus) HasLegalBasis() bool`

HasLegalBasis returns a boolean if a field has been set.

### GetSetStatusSuccessReason

`func (o *PublicStatus) GetSetStatusSuccessReason() string`

GetSetStatusSuccessReason returns the SetStatusSuccessReason field if non-nil, zero value otherwise.

### GetSetStatusSuccessReasonOk

`func (o *PublicStatus) GetSetStatusSuccessReasonOk() (*string, bool)`

GetSetStatusSuccessReasonOk returns a tuple with the SetStatusSuccessReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetStatusSuccessReason

`func (o *PublicStatus) SetSetStatusSuccessReason(v string)`

SetSetStatusSuccessReason sets SetStatusSuccessReason field to given value.

### HasSetStatusSuccessReason

`func (o *PublicStatus) HasSetStatusSuccessReason() bool`

HasSetStatusSuccessReason returns a boolean if a field has been set.

### GetSource

`func (o *PublicStatus) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PublicStatus) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PublicStatus) SetSource(v string)`

SetSource sets Source field to given value.


### GetSubscriptionId

`func (o *PublicStatus) GetSubscriptionId() int32`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PublicStatus) GetSubscriptionIdOk() (*int32, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PublicStatus) SetSubscriptionId(v int32)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetLegalBasisExplanation

`func (o *PublicStatus) GetLegalBasisExplanation() string`

GetLegalBasisExplanation returns the LegalBasisExplanation field if non-nil, zero value otherwise.

### GetLegalBasisExplanationOk

`func (o *PublicStatus) GetLegalBasisExplanationOk() (*string, bool)`

GetLegalBasisExplanationOk returns a tuple with the LegalBasisExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalBasisExplanation

`func (o *PublicStatus) SetLegalBasisExplanation(v string)`

SetLegalBasisExplanation sets LegalBasisExplanation field to given value.

### HasLegalBasisExplanation

`func (o *PublicStatus) HasLegalBasisExplanation() bool`

HasLegalBasisExplanation returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *PublicStatus) GetBusinessUnitId() int64`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *PublicStatus) GetBusinessUnitIdOk() (*int64, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *PublicStatus) SetBusinessUnitId(v int64)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *PublicStatus) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### GetStatus

`func (o *PublicStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *PublicStatus) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PublicStatus) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PublicStatus) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



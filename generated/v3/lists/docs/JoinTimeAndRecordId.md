# JoinTimeAndRecordId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordId** | **string** |  | 
**MembershipTimestamp** | **time.Time** |  | 

## Methods

### NewJoinTimeAndRecordId

`func NewJoinTimeAndRecordId(recordId string, membershipTimestamp time.Time, ) *JoinTimeAndRecordId`

NewJoinTimeAndRecordId instantiates a new JoinTimeAndRecordId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinTimeAndRecordIdWithDefaults

`func NewJoinTimeAndRecordIdWithDefaults() *JoinTimeAndRecordId`

NewJoinTimeAndRecordIdWithDefaults instantiates a new JoinTimeAndRecordId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordId

`func (o *JoinTimeAndRecordId) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *JoinTimeAndRecordId) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *JoinTimeAndRecordId) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetMembershipTimestamp

`func (o *JoinTimeAndRecordId) GetMembershipTimestamp() time.Time`

GetMembershipTimestamp returns the MembershipTimestamp field if non-nil, zero value otherwise.

### GetMembershipTimestampOk

`func (o *JoinTimeAndRecordId) GetMembershipTimestampOk() (*time.Time, bool)`

GetMembershipTimestampOk returns a tuple with the MembershipTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipTimestamp

`func (o *JoinTimeAndRecordId) SetMembershipTimestamp(v time.Time)`

SetMembershipTimestamp sets MembershipTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



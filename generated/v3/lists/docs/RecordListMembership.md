# RecordListMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListId** | **string** |  | 
**ListVersion** | **int32** |  | 
**LastAddedTimestamp** | **time.Time** |  | 
**FirstAddedTimestamp** | **time.Time** |  | 

## Methods

### NewRecordListMembership

`func NewRecordListMembership(listId string, listVersion int32, lastAddedTimestamp time.Time, firstAddedTimestamp time.Time, ) *RecordListMembership`

NewRecordListMembership instantiates a new RecordListMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordListMembershipWithDefaults

`func NewRecordListMembershipWithDefaults() *RecordListMembership`

NewRecordListMembershipWithDefaults instantiates a new RecordListMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListId

`func (o *RecordListMembership) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *RecordListMembership) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *RecordListMembership) SetListId(v string)`

SetListId sets ListId field to given value.


### GetListVersion

`func (o *RecordListMembership) GetListVersion() int32`

GetListVersion returns the ListVersion field if non-nil, zero value otherwise.

### GetListVersionOk

`func (o *RecordListMembership) GetListVersionOk() (*int32, bool)`

GetListVersionOk returns a tuple with the ListVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListVersion

`func (o *RecordListMembership) SetListVersion(v int32)`

SetListVersion sets ListVersion field to given value.


### GetLastAddedTimestamp

`func (o *RecordListMembership) GetLastAddedTimestamp() time.Time`

GetLastAddedTimestamp returns the LastAddedTimestamp field if non-nil, zero value otherwise.

### GetLastAddedTimestampOk

`func (o *RecordListMembership) GetLastAddedTimestampOk() (*time.Time, bool)`

GetLastAddedTimestampOk returns a tuple with the LastAddedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAddedTimestamp

`func (o *RecordListMembership) SetLastAddedTimestamp(v time.Time)`

SetLastAddedTimestamp sets LastAddedTimestamp field to given value.


### GetFirstAddedTimestamp

`func (o *RecordListMembership) GetFirstAddedTimestamp() time.Time`

GetFirstAddedTimestamp returns the FirstAddedTimestamp field if non-nil, zero value otherwise.

### GetFirstAddedTimestampOk

`func (o *RecordListMembership) GetFirstAddedTimestampOk() (*time.Time, bool)`

GetFirstAddedTimestampOk returns a tuple with the FirstAddedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAddedTimestamp

`func (o *RecordListMembership) SetFirstAddedTimestamp(v time.Time)`

SetFirstAddedTimestamp sets FirstAddedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



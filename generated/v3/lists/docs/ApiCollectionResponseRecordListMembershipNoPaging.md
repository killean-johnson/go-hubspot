# ApiCollectionResponseRecordListMembershipNoPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Results** | [**[]RecordListMembership**](RecordListMembership.md) |  | 

## Methods

### NewApiCollectionResponseRecordListMembershipNoPaging

`func NewApiCollectionResponseRecordListMembershipNoPaging(results []RecordListMembership, ) *ApiCollectionResponseRecordListMembershipNoPaging`

NewApiCollectionResponseRecordListMembershipNoPaging instantiates a new ApiCollectionResponseRecordListMembershipNoPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCollectionResponseRecordListMembershipNoPagingWithDefaults

`func NewApiCollectionResponseRecordListMembershipNoPagingWithDefaults() *ApiCollectionResponseRecordListMembershipNoPaging`

NewApiCollectionResponseRecordListMembershipNoPagingWithDefaults instantiates a new ApiCollectionResponseRecordListMembershipNoPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ApiCollectionResponseRecordListMembershipNoPaging) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ApiCollectionResponseRecordListMembershipNoPaging) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ApiCollectionResponseRecordListMembershipNoPaging) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ApiCollectionResponseRecordListMembershipNoPaging) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetResults

`func (o *ApiCollectionResponseRecordListMembershipNoPaging) GetResults() []RecordListMembership`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApiCollectionResponseRecordListMembershipNoPaging) GetResultsOk() (*[]RecordListMembership, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApiCollectionResponseRecordListMembershipNoPaging) SetResults(v []RecordListMembership)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



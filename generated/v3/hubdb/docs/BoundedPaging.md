# BoundedPaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**BoundedNextPage**](BoundedNextPage.md) |  | [optional] 

## Methods

### NewBoundedPaging

`func NewBoundedPaging() *BoundedPaging`

NewBoundedPaging instantiates a new BoundedPaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundedPagingWithDefaults

`func NewBoundedPagingWithDefaults() *BoundedPaging`

NewBoundedPagingWithDefaults instantiates a new BoundedPaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *BoundedPaging) GetNext() BoundedNextPage`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *BoundedPaging) GetNextOk() (*BoundedNextPage, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *BoundedPaging) SetNext(v BoundedNextPage)`

SetNext sets Next field to given value.

### HasNext

`func (o *BoundedPaging) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



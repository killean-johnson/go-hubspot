# BoundedNextPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int32** |  | 
**Link** | Pointer to **string** |  | [optional] 

## Methods

### NewBoundedNextPage

`func NewBoundedNextPage(offset int32, ) *BoundedNextPage`

NewBoundedNextPage instantiates a new BoundedNextPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundedNextPageWithDefaults

`func NewBoundedNextPageWithDefaults() *BoundedNextPage`

NewBoundedNextPageWithDefaults instantiates a new BoundedNextPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *BoundedNextPage) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *BoundedNextPage) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *BoundedNextPage) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLink

`func (o *BoundedNextPage) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *BoundedNextPage) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *BoundedNextPage) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *BoundedNextPage) HasLink() bool`

HasLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



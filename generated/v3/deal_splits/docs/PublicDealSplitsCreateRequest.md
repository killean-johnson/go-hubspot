# PublicDealSplitsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Splits** | [**[]PublicDealSplitInput**](PublicDealSplitInput.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewPublicDealSplitsCreateRequest

`func NewPublicDealSplitsCreateRequest(splits []PublicDealSplitInput, id int32, ) *PublicDealSplitsCreateRequest`

NewPublicDealSplitsCreateRequest instantiates a new PublicDealSplitsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicDealSplitsCreateRequestWithDefaults

`func NewPublicDealSplitsCreateRequestWithDefaults() *PublicDealSplitsCreateRequest`

NewPublicDealSplitsCreateRequestWithDefaults instantiates a new PublicDealSplitsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSplits

`func (o *PublicDealSplitsCreateRequest) GetSplits() []PublicDealSplitInput`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *PublicDealSplitsCreateRequest) GetSplitsOk() (*[]PublicDealSplitInput, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *PublicDealSplitsCreateRequest) SetSplits(v []PublicDealSplitInput)`

SetSplits sets Splits field to given value.


### GetId

`func (o *PublicDealSplitsCreateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicDealSplitsCreateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicDealSplitsCreateRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



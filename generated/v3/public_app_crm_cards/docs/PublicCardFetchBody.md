# PublicCardFetchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypes** | [**[]CardObjectTypeBody**](CardObjectTypeBody.md) |  | 
**TargetUrl** | **string** |  | 

## Methods

### NewPublicCardFetchBody

`func NewPublicCardFetchBody(objectTypes []CardObjectTypeBody, targetUrl string, ) *PublicCardFetchBody`

NewPublicCardFetchBody instantiates a new PublicCardFetchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCardFetchBodyWithDefaults

`func NewPublicCardFetchBodyWithDefaults() *PublicCardFetchBody`

NewPublicCardFetchBodyWithDefaults instantiates a new PublicCardFetchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypes

`func (o *PublicCardFetchBody) GetObjectTypes() []CardObjectTypeBody`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PublicCardFetchBody) GetObjectTypesOk() (*[]CardObjectTypeBody, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PublicCardFetchBody) SetObjectTypes(v []CardObjectTypeBody)`

SetObjectTypes sets ObjectTypes field to given value.


### GetTargetUrl

`func (o *PublicCardFetchBody) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *PublicCardFetchBody) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *PublicCardFetchBody) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PublicMergeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectIdToMerge** | **string** | The ID of the company to merge into the primary. | 
**PrimaryObjectId** | **string** | The ID of the primary company, which the other will merge into. | 

## Methods

### NewPublicMergeInput

`func NewPublicMergeInput(objectIdToMerge string, primaryObjectId string, ) *PublicMergeInput`

NewPublicMergeInput instantiates a new PublicMergeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicMergeInputWithDefaults

`func NewPublicMergeInputWithDefaults() *PublicMergeInput`

NewPublicMergeInputWithDefaults instantiates a new PublicMergeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectIdToMerge

`func (o *PublicMergeInput) GetObjectIdToMerge() string`

GetObjectIdToMerge returns the ObjectIdToMerge field if non-nil, zero value otherwise.

### GetObjectIdToMergeOk

`func (o *PublicMergeInput) GetObjectIdToMergeOk() (*string, bool)`

GetObjectIdToMergeOk returns a tuple with the ObjectIdToMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIdToMerge

`func (o *PublicMergeInput) SetObjectIdToMerge(v string)`

SetObjectIdToMerge sets ObjectIdToMerge field to given value.


### GetPrimaryObjectId

`func (o *PublicMergeInput) GetPrimaryObjectId() string`

GetPrimaryObjectId returns the PrimaryObjectId field if non-nil, zero value otherwise.

### GetPrimaryObjectIdOk

`func (o *PublicMergeInput) GetPrimaryObjectIdOk() (*string, bool)`

GetPrimaryObjectIdOk returns a tuple with the PrimaryObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryObjectId

`func (o *PublicMergeInput) SetPrimaryObjectId(v string)`

SetPrimaryObjectId sets PrimaryObjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



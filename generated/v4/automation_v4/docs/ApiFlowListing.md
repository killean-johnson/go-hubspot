# ApiFlowListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionId** | **string** | Deprecated. Will be removed. | 
**CreatedAt** | **time.Time** | The timestamp this flow was created. | 
**ObjectTypeId** | **string** | The CRM object type for objects that can be enrolled into this flow. | 
**IsEnabled** | **bool** | This controls whether or not the flow is \&quot;enabled\&quot; if it&#39;s actively listening for enrollment triggers and executing actions. If this is &#x60;false&#x60; the flow is not accepting any enrollments or executing any actions. | 
**Name** | Pointer to **string** | The user-provided name for this flow. Names get auto-created for workflows that are created without a name. | [optional] 
**Id** | **string** | The unique ID for this flow. This is auto-generated when creating the flow. | 
**Uuid** | Pointer to **string** | An optional unique key for this flow. This is only unique per-portal. | [optional] 
**FlowType** | **string** | Deprecated. Will be removed. | 
**UpdatedAt** | **time.Time** | The timestamp this flow was last updated. | 

## Methods

### NewApiFlowListing

`func NewApiFlowListing(revisionId string, createdAt time.Time, objectTypeId string, isEnabled bool, id string, flowType string, updatedAt time.Time, ) *ApiFlowListing`

NewApiFlowListing instantiates a new ApiFlowListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFlowListingWithDefaults

`func NewApiFlowListingWithDefaults() *ApiFlowListing`

NewApiFlowListingWithDefaults instantiates a new ApiFlowListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionId

`func (o *ApiFlowListing) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *ApiFlowListing) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *ApiFlowListing) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetCreatedAt

`func (o *ApiFlowListing) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiFlowListing) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiFlowListing) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetObjectTypeId

`func (o *ApiFlowListing) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ApiFlowListing) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ApiFlowListing) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetIsEnabled

`func (o *ApiFlowListing) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiFlowListing) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiFlowListing) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetName

`func (o *ApiFlowListing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiFlowListing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiFlowListing) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiFlowListing) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ApiFlowListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiFlowListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiFlowListing) SetId(v string)`

SetId sets Id field to given value.


### GetUuid

`func (o *ApiFlowListing) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiFlowListing) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiFlowListing) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiFlowListing) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetFlowType

`func (o *ApiFlowListing) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *ApiFlowListing) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *ApiFlowListing) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.


### GetUpdatedAt

`func (o *ApiFlowListing) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApiFlowListing) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApiFlowListing) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



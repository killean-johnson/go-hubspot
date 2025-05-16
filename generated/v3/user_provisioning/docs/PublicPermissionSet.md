# PublicPermissionSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiresBillingWrite** | **bool** | Whether this role has a paid seat and requires the billing-write scope to assign/unassign to users | 
**Name** | **string** | The role&#39;s name | 
**Id** | **string** | The role&#39;s unique ID | 

## Methods

### NewPublicPermissionSet

`func NewPublicPermissionSet(requiresBillingWrite bool, name string, id string, ) *PublicPermissionSet`

NewPublicPermissionSet instantiates a new PublicPermissionSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPermissionSetWithDefaults

`func NewPublicPermissionSetWithDefaults() *PublicPermissionSet`

NewPublicPermissionSetWithDefaults instantiates a new PublicPermissionSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiresBillingWrite

`func (o *PublicPermissionSet) GetRequiresBillingWrite() bool`

GetRequiresBillingWrite returns the RequiresBillingWrite field if non-nil, zero value otherwise.

### GetRequiresBillingWriteOk

`func (o *PublicPermissionSet) GetRequiresBillingWriteOk() (*bool, bool)`

GetRequiresBillingWriteOk returns a tuple with the RequiresBillingWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresBillingWrite

`func (o *PublicPermissionSet) SetRequiresBillingWrite(v bool)`

SetRequiresBillingWrite sets RequiresBillingWrite field to given value.


### GetName

`func (o *PublicPermissionSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicPermissionSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicPermissionSet) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *PublicPermissionSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicPermissionSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicPermissionSet) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


